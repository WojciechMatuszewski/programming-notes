# RxJS

## Basics

### Hot and Cold

With `RxJs` there is a notion of `hot` and `cold` observables.

- with `cold` observables value-producer is created when `subscribing`. Making a
  `http` call is a good example (in _Angular_).
- with `hot` observables value-producer is closed over by a given `observable`.
  Listening to a `scroll` event with `fromEvent` is a good example

So, the notion of `hot` and `cold` carries with itself some gotchas.

```js
const obs$ = of(1).pipe(from(() => fetch("something")));
```

So a network request will be made **whenever you subscribe to `obs$`** since
`obs$` is a `cold` observable... and this CAN be a problem.

### Multicast vs Unicast

So this concept goes along with `hot` and `cold` observables.

- **multicast** means that all subscribers **get the data from the same
  producer** (they share the source)

- **unicast** means that each subscriber **gets the data from different
  producer**

Again, this is pretty obvious when we look on _natively-hot_ observables like
`fromEvent` streams.

### Making cold hot

There are a couple of ways to do this.

#### Multicast operator

This operator is used to take control over the whole process of making an `cold`
observable hot. You have to pass the `subject` as new _provider_, since subjects
are `multicast` (and hot) by default.

```js
this.user$ = this.http.get(`api/user/1`).pipe(multicast(new Subject()));

this.users$.connect();
```

One big drawback of this approach is that you have to manually call `.connect`
on a given _source_. It also does not take care of counting existing subscribers
which can introduce memory leaks and such.

#### Other Multicast-like operators

`RxJs` provides a lot of ways to achieve `hot` observable without you having to
worry about counting subscribers and such:

- `publish`
- `share` ... and many more (with different subjects passed into `multicast`,
  etc..)

https://medium.com/@benlesh/hot-vs-cold-observables-f8094ed53339

https://itnext.io/the-magic-of-rxjs-sharing-operators-and-their-differences-3a03d699d255

## Properly handling errors using `catchError`

So you've been using `catchError` just treating it like `.catch` in a
promise-base API and it all seemed all good and sweet. But sometimes you
encountered a bug where a stream would not be called again after an error. _But
you caught the error with `catchError` and returned a new stream_, what could go
wrong?🤔. Well just know that: **`catchError` replaces whole stream, WHOLE
STREAM**. Now lets see an example:

```typescript
source$.pipe(
  // switchMap can fail
  switchMap((something) => from(resourceGetterFn(something))).pipe(
    // resolveResourceResponse can fail
    mergeMap((response) => resolveResourceResponse(response)),
  ),
);
```

Now, what would happen when we did this:

```typescript
source$.pipe(
  // previous code with switchMap etc
  catchError((_) => {
    return of();
  }),
);
```

So, error is propagated and is caught by `catchError`, that's all and good. But
again **CATCH ERROR REPLACES WHOLE STREAM!**(and we are returning an empty
Observable). That means, **after an error, that operator is just an empty
Observable**.

### Solution

Solution would be... well reading the docs and such (and actually understanding
what code you are writing). To solve this problem we just need to move
`catchError` **inside switchMap**.

```typescript
source$.pipe(
  // switchMap can fail
  switchMap((something) => from(resourceGetterFn(something))).pipe(
    mergeMap((response) => resolveResourceResponse(response)),
    catchError((_) => {
      return of();
    }),
  ),
);
```

There, no magic, no weird copy-paste from stack. That's all.

_Reference:
[this great article](https://medium.com/city-pantry/handling-errors-in-ngrx-effects-a95d918490d9)_

## Operators

### `fromFetch` and `ajax`

You do not have to use `from(fetch...)` to get obs. back anymore. RxJs ships
with 2 operators that do the fetching for you and produce observable. There is
one catch though.

**`fromFetch` will automatically setup abort controller**. Which is great news!
But... it may cause errors while working on older versions of browsers, like
IE 11.

`fromFetch` is lazy. It will only fire when you subscribe.

```js
const data$ = from(fetch("")); // fired right away

const data2$ = defer(() => from(fetch(""))); // fired when subscribed to

const data3$ = fromFetch(""); // fired when subscribed to
```

The fact that `from(fetch)` fires right away is huge. I wonder how many bugs
that caused in my applications 🤔

## Recipes

### Subscription Sink

There are various ways to handle `cleanup a.k.a unsubscribing` with `RxJS`.

- using _async pipe_ in _Angular_
- using plain old `unsubscribe`
- using `takeUntil` watcher pattern
- using **`subscription sink`**

So what is this `subscription sink` ?

Well, it turns out you can add multiple subscribers to one _meta_ subscription
and unsubscribe only from this one giant _met_ subscription.

```js
const subSink = /* some cone that produces Subscription */
    const; /* some cone that produces Subscription */

subSink.unsubscribe();
```

Actually you can even create standalone `subscription` using `new` keyword.

```js
const subSink = new Subscription();
```

Pretty neat stuff!

### Teardown with `takeUntil`

`takeUntil` is really great when working with DOM and events. You want to drop
`Observable` completely when something happens, eg. a _touchend_ event or smth.

This is all and good but did you know that you can actually place operators
**AFTER** `takeUntil`?.

```js
const touchStart$ = fromEvent(element, "touchstart");
const touchMove$ = fromEvent(element, "touchmove");
const touchEnd$ = fromEvent(element, "touchend");

const drag$ = touchStart.pipe(switchMap(() => touchMove$.pipe(takeUntil(touchEnd$), concat() /*some observable*/)));
```

This pattern is really powerful when combined with `defer`. (Remember that
whatever you put into `concat` will be run at _define_ time !!. That is why you
usually want to use `defer`)

### Pooling

When you want to update stuff every X seconds

```js
const timer$ = timer(0, 5000);

timer$.pipe(
  exhaustMap(() => /* your http call */)
)
```

`exhaustMap` will not fire another request till previous request in-flight is
not finished. If you want to drop that request and start another one you
probably need `switchMap`.

### Drag and Drop

```js
const element = querySelector("element");

const mouseDown$ = fromEvent(element, "mousedown");
const mouseUp$ = fromEvent(element, "mouseup");
const mouseLeave$ = fromEvent(element, "mouseleave");
const mouseMove$ = fromEvent(element, "mousemove");

const stop$ = merge(mouseLeave$, mouseUp$);

const dragAndDrop$ = mouseDown$.pipe(exhaustMap(() => mouseMove$.pipe(takeUntil($stop))));
```

### Prevent double click

This is very useful :).
[A **very basic** implementation could be found here.](https://codesandbox.io/s/old-feather-rz6xm)

```js
const button = querySelector("element");
const buttonClick$ = fromEvent(button, "click");

const preventDoubleSubmit$ = buttonClick$.pipe(exhaustMap(() => http.post(/* something */)));
```

## Uncommon Types

### Notification

Aside from all the `Subject`-y related types there is also notification.

`Notification` **does not create an observer**. It wraps it annotating it with
additional metadata.

Example:

```js
of(1)
  .pipe(mapTo(new Notification("E")))
  .subscribe(console.log);
/*
    Notification {kind: "E", value: undefined, error: undefined, hasValue: false, constructor: Object}
    kind: "E"
    value: undefined
    error: undefined
    hasValue: false
    <constructor>: "Notification"
*/
```

As you see it can swallow original values. Notification can be of a different
type (type corresponds to observable life cycle).

## Uncommon Operators

### Materialize / dematerialize

So you know about `Notification` type. It's all good and great but you probably
wonder how you could use it.

Lets say you have a source that can error out. Sure, that could happen but when
that does happen, **error bubbles up and ignores every operator that is yet to
come**. This might be a problem.

To prevent this you might want to turn the error into `Notification` using
`materialize` operator. If a given source errors out you can check if
`Notification` is of type error and act accordingly without skipping operators.

Example:

```js
// sample stream
interval(500)
  .pipe(
    mapTo("normal value"),
    // sometimes value, sometimes throw
    map((v) => {
      if (randomInt() > 50) {
        throw new Error("boom!");
      } else return v;
    }),
    materialize(),
    // turns Observable<T> into Notification<Observable<T>>
    // so we can delay or use other operators.
    delay(500),
    // Notification of value (error message)
    map((n) => (n.hasValue ? n : new Notification("N", n.error.message, null))),
    // back to normal
    dematerialize(),
  )
  // now it never throw so in console we will have
  // `normal value` or `boom!` but all as... normal values (next() emission)
  // and delay() works as expected
  .subscribe((v) => console.log(v));
```
