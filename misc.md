# Random Stuff

## Back pressure

Back pressure is when you cannot process data that is coming to you fast enough.
Let's say you are working on a search bar input. You would not want to send http
request every keystroke right?. That's back pressure, there is no way for you to
process keystroke information fast enough. What do you do? You probably debounce
or throttle it.

Back pressure usually stems from the fact that you do not have the ability to
control the producer (and are working in a _push_ environment).

Other solutions may include

- buffering
- sampling (giving only a sample of processed data)

## Memoization

Pretty standard technique to prevent unnecessary computation. You create a cache
and store previously computed results there. One catch is that you have to be
careful with cache size. It might grow pretty fast and then you do you do.

**You should probably only use memoization with pure functions**.

Simple example

```js
function memoize(func) {
  return function memoized() {
    // we are doing it old school :D
    var args = Array.prototype.slice.call(arguments);
    // cache can be a closed over variable or variable on function itself
    func.cache = func.cache || {};

    var cachedResult = func.cache[args];

    if (cachedResult != null) return cachedResult;

    var computationResult = func.apply(this, args);

    func.cache[args] = computationResult;
    return computationResult;
  };
}
```

## Reducing promises

So the question is

> How do i sequentially resolve promises ?

The answer is simple:

> Use reduce

But the **why it works** requires deeper understanding of `Promise` API.

### Flattening and Fake Promises

While reading `Async and Performance` by my boi Kyle, I learned a lot about
`Promise` API.

One thing that stuck with me is that `Promise` API has a native flattening
behavior.

Check this out:

```js
function wait(ms) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve();
    }, ms);
  });
}

Promise.resolve(42)
  // automatically resolve
  .then(() => wait(500))
  // im going to get to that handler \/ after more-less 500ms
  .then(console.log);
```

We all know that we can chain `.then` and also that stuff returned from `.then`
is automatically treated like a `Promise`. But I never returned another promise
from within `.then`. Normally what would result in something like
`Promise(Promise(..))`.

Well, as I said before, `Promise` API natively flattens nested Promises. We can
event test this theory using _fake promise_.

```js
function fakePromise() {
  return {
    then: function() {
      console.log('IM CALLED');
    }
  };
}

Promise.resolve(42).then(fakePromise);

// logs: 'IM CALLED'
```

It turns out **any object that has `then` method** inside it will be treated as
a `Promise` (at least is expect to behave like one).

No matter the nesting, `Promise` API will drill to the last `.then` and wait for
that to resolve.

Pretty neat stuff huh?

### Reducing

So armed with the knowledge that `Promise` API flattens natively, the _how it
works_ when reducing promises should be pretty easy for us to understand.

```js
[p1, p2, p3].reduce((acc, nextPromise) => {
  return acc.then(nextPromise);
}, Promise.resolve());
```

Here, just like before we are returning a `Promise` from within a `Promise`.
Flattening kicks in and we wait for the deepest `.then` resolution (being
called) and proceed with another one.

## CORS

CORS is the mechanism which enables one webapp (like your localhost) to share
some resource with another webapp (like your endpoint API).

If those 2 apps have **the same origin** they can easily share those resources
with no hassle at all.

Problems begin when they are on different origins.

So what does _different origin_ mean?

- different domain like `google.com` and `twitter.com`
- different subdomains like `localhost:3000` and `localhost:3000.api/v1`
- different ports like `:3000` and `:4000`
- different protocols like `http` and `https`

To make it work you have to follow the CORS standard.

So how does it work?

Suppose we have 2 apps: A and B. They want to share resources. App A makes a
POST request to app B:

- `preflight` request is made (before the actual request) **also known as
  OPTIONS call**
- app B now have the responsibility of verifying either this request is valid or
  not.
- app B sets some additional headers to that request and sends it back.
- now browser knows if the request is valid or not. The actual `POST` request is
  made

### Simple Request vs Preflight Request

So we've seen how the preflight mechanism works. But the next question on your
mind probably is: is this happening every time I send a request?

Well, no.

Some request are marked as `simple` by the browser and the preflight (`OPTIONS`)
request is not send.

### Caching

Browsers can actually cache preflight responses. You usually specify that in a
header.

## Observers API

You probably know of `intersectionObserver`. An API to check if given object is
visible at the screen currently.

But there are other _observers_ too:

- `mutationObserver`: used to watch for `DOM` _mutations_.

* `resizeObserver`: a new kid on the block. Tries to achieve the holy grail of
  being notified when **given element** resizes to given width/height.

There is actually very interesting article on `resizeObserver` by Philip Walton
[Link to article](https://philipwalton.com/articles/responsive-components-a-solution-to-the-container-queries-problem/)

## Check for Idle period

This would be nice would not it? Having a way to know when browser is finished
doing stuff so that we can fire off some kind of computation.

There is an API for that: `requestIdleCallback`, but sadly is not all green when
it comes to browsers.

That said, you can actually do very interesting stuff with this API, described
on Philip Walton's blog
[Link to article](https://philipwalton.com/articles/idle-until-urgent/)

## Magic Webpack Comments

There soon may be deprecated due to `webpack 5` releasing but did you know that you can use multiple of them? like

```js
const Tilt = React.lazy(() =>
  import(/* webpackChunkName: "tilt", webpackPrefetch: true */ '../tilt')
);
```

This works !

## Web workers with webpack

So web workers are great. They enable you to offload the work on different thread. Nice!.

But did you know that there is an webpack plugin which enables you to turn any js file into _web worker_?. That plugin is called `workerize-loader`.
Let's say you have a module called `expensive.js` and you want to import that module as _web worker_.

```js
import expensiveWorkerized from 'workerize-loader!./expensive';
```

**BOOM!**. Thats all. Granted now methods exposed by `expensiveWorkerized` are _async_ but that should not be a problem.
