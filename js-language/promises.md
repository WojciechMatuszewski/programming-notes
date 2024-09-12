# Promises

## The `new Promise` constructor

- **The callback inside the `new Promise` is called synchronously**. It is an anti-pattern to use this callback to "defer" work.

- The **callback inside the `new Promise` can be asynchronous**. Using the `async` function as the callback **is a code smell**. It can lead to errors.

  ```js
  const p = new Promise(async () => {
    throw new Error("boom");
  });
  p.catch((e) => {
    console.log(e); // never called
  });

  const p2 = new Promise(() => {
    throw new Error("boom");
  });

  p2.catch((e) => {
    console.log(e); // called with the error
  });
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

But **why it works** requires deeper understanding of `Promise` API.

### Flattening and Fake Promises

While reading `Async and Performance` by my boi Kyle, I learned a lot about `Promise` API.

One thing that stuck with me is that `Promise` API has a native flattening behavior.

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

We all know that we can chain `.then` and also that stuff returned from `.then` is automatically treated like a `Promise`.
But I never returned another promise from within `.then`. Normally what would result in something like `Promise(Promise(..))`.

Well, as I said before, `Promise` API natively flattens nested Promises. We can even test this theory using _fake promise_.

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

It turns out **any object that has `then` method** inside it will be treated as a `Promise` (at least is expect to behave like one).

No matter the nesting, `Promise` API will drill to the last `.then` and wait for that to resolve.

Pretty neat stuff, huh?

### Reducing

So armed with the knowledge that `Promise` API flattens natively, _how it works_ when reducing promises should be pretty easy for us to understand.

```js
[p1, p2, p3].reduce((acc, nextPromise) => {
  return acc.then(nextPromise);
}, Promise.resolve());
```

Here, just like before we are returning a `Promise` from within a `Promise`. Flattening kicks in, and we wait for the deepest `.then` resolution (being
called) and proceed with another one.

## The "gravity" of `await` statement

> Based on [this great blog post](https://frontside.com/blog/2023-12-11-await-event-horizon/). **Please read the notes as I think one of the examples in the article introduces a memory leak!**

When you think about it, the `await` statement is like a black hole. Once you call it, the scope that holds this statement is _stuck_ until the underlying promise resolves.

```js
async function doSomeWork() {
  console.log("before");
  await work();
  console.log("after");
}
```

The `after` will never appear in the console until the `work` resolves. Usually it does, but there might be times when it does not. As an example, consider a system that invokes the `work` function, and then is prompted to be shut down by the user via the `INTERRUPT` signal. If that happens, we have a leak!

### The `signal` to the rescue?

If you have experience with `abortController` API, you might think that the `signal` is the answer here. The `abortController` API is most often used in the context of data fetching, like so.

```js
const controller = new AbortController();

fetch("foo", {
  signal: controller.signal,
});

// In some other place in the code
controller.abort();
```

Using the `abort` in this context will cause the **browser to ignore the response from the server and throw an error (AbortError)**. This is quite important to understand – **aborting the request does not mean it will not be fulfilled on the server**. Once the server picked up the request, it will be processed. As such **aborting the request is not a silver bullet, you should not consider it as an optimization, but rather a mechanism to control the control flow**.

Okay, so we know the `abort` is quite useful when it comes to cancelling requests on the client-side. But what if the promise you are dealing with does not make a request at all? What if the API that returns the promise does not take in the `signal` parameter?

In this case, you might find the following function handy.

```js
async function safe({ promise, signal }) {
  return await Promise.race([
    promise,
    new Promise((_, reject) => {
      signal.addEventListener("abort", () => {
        reject();
      });
    }),
  ]);
}
```

Here, we are going to **ignore the response from the `promise` when signal aborts**. Again, we are not cancelling anything, we are only ignoring the result. The work that the `promise` kicked-off will still happen.

But what if the `promise` rejects? The semantics of the `Promise.race` tell us, that it will throw as soon as the first promise rejects. **As such we have created a memory leak – the listener created via `addEventListener` will still be in memory when the promise rejects!** Not good. Here is an issue in [kibana, a well known tools for metrics](https://github.com/elastic/kibana/pull/81996) where developers forgot to clean up the listener. **Always clean up the listeners!**

I see two ways to achieve this. One uses another `signal` to cancel the listener.

```js
async function safe({ promise, signal }) {
  const listenerController = new AbortController();

  try {
    return await Promise.race([
      promise,
      new Promise((_, reject) => {
        signal.addEventListener(
          "abort",
          () => {
            reject();
          },
          { signal: listenerController.signal },
        );
      }),
    ]);
  } finally {
    listenerController.abort();
  }
}
```

Another one is to save the listener instance somewhere and then re-use it.

```js
async function safe({ promise, signal }) {
  let abortListener = () => {};

  try {
    return await Promise.race([
      new Promise((_, reject) => {
        abortListener = () => {
          reject();
        };

        signal.addEventListener("abort", abortListener);
      }),
      promise,
    ]);
  } finally {
    signal.removeEventListener("abort", abortListener);
  }
}
```

I personally prefer the `signal` approach for removing the event mainly due to the lack of the "dangling" variable which we mutate. Note that we have to mutate it, since the reference for the listener has to be the same in both `addEventListener` and `removeEventListener` cases.
