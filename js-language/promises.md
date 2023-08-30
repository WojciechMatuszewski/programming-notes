# Promises

## The `new Promise` constructor

- **The callback inside the `new Promise` is called synchronously**. It is an anti-pattern to use this callback to "defer" work.

- The **callback inside the `new Promise` can be asynchronous**. Using the `async` function as the callback **is a code smell**. It can lead to errors.

  ```js
    const p = new Promise(async () => {
      throw new Error("boom")
    })
    p.catch(e => {
      console.log(e) // never called
    })

    const p2 = new Promise(() => {
      throw new Error("boom")
    })

    p2.catch(e => {
      console.log(e) // called with the error
    })
  ```

- The **value you return from the callback is IGNORED**.

## The `.then(callback, errorCallback)` syntax

- It is a valid way for handling errors, but **it does only work "locally" to a given `.then` callback**.

  - Using it might be dangerous – **you might forget about providing the second parameter for another `.then` callback**. This will result in an unhandled promise rejection.

- Since the error handling does not apply globally when using the `errorCallback`, it is recommended to use **.catch**.

  - The `.catch` works very similar to `.then(null, errorCallback)`.

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
    then: function () {
      console.log("IM CALLED");
    },
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
