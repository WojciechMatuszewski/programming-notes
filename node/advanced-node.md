# Advanced node

Course material from [node university](https://node.university/).

## How does `require` work

- Imports given module

* There is **require cache**. That means that **module will be imported only once**

- The **cache is based on all imports within a program**

* Cache can be cleared

- The **`require` will actually execute the code in the file**

* If you **do not want the code to be executed** – **use `require.resolve`**
  You would use this method to check if a module exists. The **`require.resolve` also uses the cache mechanism**

- The `require` statement is **synchronous**. It uses `fs.readFileSync` under the hood.

  ```js
  // code
  if (cached && cached.source) {
    content = cached.source;
    cached.source = undefined;
  } else {
    content = fs.readFileSync(filename, "utf8");
  }
  // code
  ```

* Remember that each module is **wrapped within a function which exposes `__dirname`, `__filename`, `module`, `exports` to you**.
  This is why you do not have to import anything when you write `__dirname` within your module.

### `require` vs `import`

- The **`require` statement is dynamic**. That means that you can compute the final parameter.

* The **`import` statement is static**. While you cannot compute the module name, it's much better in terms of tree-shaking and dead code elimination.

### `require` cache

- You can delete cached imports

  ```js
  require("./module-4.js"); // execute code in the `module-4.js`
  delete require.cache[require.resolve("./module-4.js")];
  require("./module-4.js"); // execute code AGAIN since the item is not in the cache
  ```

## Globals

- Instead of `window` there is a `global` variable (sometimes you see it used while checking out Jest code)

## ESM

- you can opt into native ESM support by specifying `type: module` in your `package.json`

* the **most important thing to remember** is that **the globals from the CJS wrapper are thing of the past**.
  This means that `__filename` , `__dirname` and others are **gone**

- as an **alternative to CJS wrapper globals** consider using **`import.meta.XX`**

## The `v8` API

One neat trick I discovered with the v8 API is the fact that you can deeply clone an object without using the `JSON.stringify` stuff

```ts
function cloneDeep(values: Record<string, unknown>): Record<string, unknown> {
  return v8.deserialize(v8.serialize(values));
```

Look at this! So nice. Also much more readable. I bet that with the `JSON.stringify` you will get a question about what is going on.

## Event loop and promises

It turns out you can block the event loop with promises. Yup, you've heard me right. And all of this is possible because how the promises interact with the event loop.

### The native layer

Remember about all the Node.js queues that exist? The _macrotask queue_ and the _task queue_ and other queues? Turns out some of them - mainly the **microtask queue** and the **nextTick queue** are executed by so called _native layer_.

The _native layer_ is nothing more than the place where _libuv_ resides and executes. It's the **_native layer_ functions that drain the _microtask queue_ and the _nextTick queue_, not the Node.js event loop!**.

### When does native layer run

The _native layer_ **runs in-between every event loop cycle**. When **_native layer_ runs, the event loop is BLOCKED**.
So if the _native layer_ can block the event loop, **it can happen that your promises block the event loop, since they are drained by the native layer**.

Here is an illustration that shows how the event loop interacts with the _native layer_.
![event loop interacting with the native layer](../assets/native-layer-event-loop.png)

### Code example

To drive the point further, let us look at some code to make sure we understand how the _native layer_ operates.
Here is an, albeit very contrived, example of a sample Node.js program.

```js
new Promise((r) => {
  r();
})
  .then(() => {
    console.log("1");
  })
  .then(() => {
    console.log("2");
  })
  .then(() => {
    return new Promise((r) => setTimeout(r, 100));
  })
  .then(() => {
    console.log("3");
  });

setTimeout(() => {
  console.log("timer");
}, 0);

process.nextTick(() => {
  console.log("next tick");
});

setImmediate(() => {
  console.log("immediate");
});

console.log("regular code");
```

The goal is to guess what will be the order of the `console.log` statements. Please remember what we have talked about before. The **_native layer_ drains the microtask and nextTick queues in-between event loop cycles. This is a blocking process**.

The order of the log statements will be as follows

```shell
regular code
next tick
1
2
timer
immediate
3
```

1. The event loop starts, regular JS code is executed, thus first thing that is logged is the `regular code` log.
2. During the event loop cycle, the `new Promise` callback is executed because it's a synchronous call (**not completely sure here**)
3. During the event loop cycle, the `setTimeout` callback is scheduled to execute (timers queue) and the `setImmediate` is scheduled do execute (check phase, _immediate queue_)
4. The _native layer_ gets to work. The `nextTick` callback and two `then` callbacks are executed. **During this time, the _event loop_ is blocked**.
5. During this _native layer_ cycle, a `setTimeout` callback is scheduled to executed.
6. Control is yielded back to the event loop, since the _timers queue_ is run before the _immediate queue_, we see the `timer` log first, then the `setImmediate` callback is invoked
7. The _native layer_ gets to work, the last `then` callback is invoked.

TODO : https://www.youtube.com/watch?v=yRyfr1Qcf34
