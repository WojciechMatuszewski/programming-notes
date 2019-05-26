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
