# Advanced Asynchronous Javascript

## Minimal Observable Implementation

`Observable` can be defined as ES6 class.

```js {.line-numbers}
class Observable {
  constructor(subscribe) {
    this._subscribe = subscribe;
  }
  subscribe(observer) {
    return this._subscribe(observer);
  }
  // simple timeout API (factory function)
  static timeout(time) {
    // look at the constructor
    return new Observable(function subscribe(observer) {
      const timeoutRef = setTimeout(function() {
        observer.next();
        observer.complete();
      }, time);
      // returning unsubscribe so consumer can unsubscribe
      return {
        unsubscribe() {
          clearTimeout(timeoutRef);
        }
      };
    });
  }
}
```

We wrote only 24 lines of code and yet, this implementation is already powerful.

```js
const subscription = Observable.timeout(1000).subscribe({
  next: () => console.log("next"),
  complete: () => console.log("complete")
});

// you can also unsubscribe
subscription.unsubscribe();

// after 1 second
// "next"
// "complete"
```

## Hot Observable vs Cold Observable

- _Cold Observable_: no matter how many times you subscribe to it, **you always gets the same result**. Good example of hot observables are event streams (think of Node.js callback APIs for reading files or similar)

- _Hot Observable_ (A.K.A me 😉): when you subscribe to it multiple times **you might get different data**. Good example of these are event handlers in browser API. When you start a listener couple of seconds later than another one you might miss some of the events that came before you started listening for the second time.

Our simple timeout method on the Observable is a _Cold Observable_. We will never miss any notification because **we kick off the timeout when we actually call `.subscribe`**.That's called **lazy evaluation** (not doing work till you actually need to do it).

## Data Stream into Hot Observable

```js {.line-numbers}
... inside Observable class ...
  static fromEvent(domElement, eventName) {
    return new Observable(function subscribe(observer) {
      domElement.addEventListener(eventName, eventHandlerFunction);
      return {
        unsubscribe() {
          domElement.removeEventListener(eventName, eventHandlerFunction);
        }
      };
      function eventHandlerFunction(event) {
        observer.next(event);
      }
    });
  }
```

How we could use this method?

```js {.line-numbers}
const btn = document.querySelector("button");
Observable.fromEvent(btn, "click").subscribe({
  next: () => console.log(this)
});
```

What about `.this binding` ?
Well, remember what Kyle was drilling into your mind? **The only thing that matters is how the function is called**. In this case we preserve this binding because we have _arrow function_ inside our subscribe as next call. This function will go up looking for the correct context.

```js {.line-numbers}
class Component {
  constructor() {
    this.someAPI = 3;
  }

  onInit() {
    const btn = document.querySelector("button");
    Observable.fromEvent(btn, "click").subscribe({
      // logs Component context
      next: () => console.log(this)
    });
  }
}

const comp = new Component();
compOnInit();
```

### Building Map Operator

The `.map` operator is a staple of functional programming.

In newer versions of RxJs we have `.pipe` operator which allows us, well, to pipe operators. Let's start with pipe first.

### Pipe

Let's think about what's the _shape_ of an operator. To chain operators freely we have to return observable each time we apply an operator. So the shape is:

> operator(someArgument) => new Observable

```js
... Other Observable code ...
 pipe(...operators) {
     // reduce , reading from left to right
    return operators.reduce(
        // passing current observable to an operator (look at the shape)
      (currentObservable, operator) => operator(currentObservable),
      this
    );
  }
```

#### Map

Well, map takes a _projection function_ an applies that function to a argument passed to inner function. Lets see how we would implement that.

```js
function map(projectionFn) {
  // this is the inputObservable from .pipe
  return function(inputObservable) {
    // returning new observable so we can chain other methods
    return new Observable(function subscribe(outputObserver) {
      const subscription = inputObservable.subscribe({
        // you could wrap this call into try catch also
        next: x => outputObserver.next(projectionFn(x)),
        complete: () => outputObserver.complete(),
        error: err => outputObserver.error(err)
      });
      return {
        unsubscribe: subscription.unsubscribe
      };
    });
  };
}
```

It's not that scary after all, is it?

Now you can event _compose map operators_

```js
Observable.fromEvent(btn, "click")
  .pipe(
    compose(
      map(evt => 3),
      map(num => num + 1),
      map(num => num / 2)
    )
  )
  .subscribe({
    next: console.log // 2...
  });
```

#### Filter

Very similar to '.map'

```js
function filter(predicateFn) {
  return function(inputObservable) {
    return new Observable(function subscribe(outputObserver) {
      return inputObservable.subscribe({
        next: x => {
          if (predicateFn(x)) {
            return outputObserver.next(x);
          }
        }
        // other methods
      });
    });
  };
}
```

## "Animations Allowed" Problem

Let's say you want to disable animations on very low-end devices when there is a lot of things going on in the background.
How would we tackle this problem?

First thing first let's refresh our memory on when it comes to basic `RxJs` operators

**`Observable.of(value)`**
It creates an observable which calls `next(value)` and then completes

```js
Observable.of(5); // => 5
```

**`Observable.concat(...[observables])`**
Concatenates observables just like you can concatenate a string. It waits till current completes and subscribes to the next one.

**`Observable.distinctUntilChanged()`**
Very useful when creating typeahead. Basically diffs the prev value with current. Emits only when they are different.

**`Observable.scan((acc, current) => acc + current)`**
Scan is similar to reduce but it gives you all the intermediate values as well

```js
// RxJS v6+
import { of } from "rxjs";
import { scan } from "rxjs/operators";

const source = of(1, 2, 3);
// basic scan example, sum over time starting with zero
const example = source.pipe(scan((acc, curr) => acc + curr, 0));
// log accumulated values
// output: 1,3,6
const subscribe = example.subscribe(val => console.log(val));
```

### Solution

Solution is actually so brilliant that \*_you should watch it yourself because this guy is amazing_

## Catching errors

Catching errors is quite important. There are various ways to do so in RxJs.

### `catchError`

`catchError` basically catches the error in the chain. **You have to return new Observable from this operator**.

### `retry`

This operator allows you to retry some operation X number of times.
Simple implementation:

(I'm assuming this is a method on a class called _Observable_)

```js
  retry(num) {
    return new Observable(observer => {
      let currentSub = null;
      function processRequest(currentAttemptNumber) {
        currentSub = this.subscribe({
          next: v => observer.next(v),
          complete: () => observer.complete(),
          error: err => {
            if (currentAttemptNumber == 0) {
              observer.error(err);
            } else {
              processRequest(currentAttemptNumber - 1);
            }
          }
        });
      }
      processRequest(num);
      return {
        unsubscribe: currentSub.unsubscribe
      };
    });
  }
```

## Image preload

We can use RxJs to implement our own image preload function.

Maybe we could implement it like so :

```js
function preloadImage(src) {
  const img = new Image();
  const success = fromEvent(img, "load").pipe(map(() => src));
  const failure = fromEvent(img, "error").pipe(map(() => LOADING_ERROR_URL);
  img.src = src;
  return merge(success, failure);
}
```

Where is one big gotcha in this implementation. We are doing work **before** someone might subscribe to that observable. There is a possibility that **image loads before someone calls .subscribe**.

Much better implementation would use `defer` operator.

```js
function preloadImage(src) {
  return defer(() => {
    const img = new Image();
    const success = fromEvent(img, "load").pipe(map(() => src));
    const failure = fromEvent(img, "error").pipe(map(() => LOADING_ERROR_URL));
    img.src = src;
    return merge(success, failure);
  });
}
```

`defer` operator make it so that the work begins only when someone actually subscribes, creating _lazy observable_. `defer` basically is an **observable factory**
