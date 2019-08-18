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
