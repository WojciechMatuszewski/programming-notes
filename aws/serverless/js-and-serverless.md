# JS and serverless

## Reusing connections through `callbackWaitsForEmptyEventLoop`

**THE PROBLEM / SOLUTION I'M DESCRIBING ONLY EXISTS WHEN YOU ARE USING `non-async` handlers**.

There is this mysterious property on the `ctx` which is passed to you handler called `callbackWaitsForEmptyEventLoop`.

When you return (or use `callback` for that matter) from your handler, by default, lambda execution environment will wait for any hanging tasks to finish. This is **usually** what you want, but sometimes the `event loop` can be occupied with stuff regarding database connections (usually this is when you are using `mongo` or similar databases).

The easiest example would be with `setTimeout`.

This will timeout (default timeout is 3 seconds):

```js
module.exports.handler = (event, ctx, callback) => {
    setTimeout(() => console.log("timeout"), 10000);

    // or return
    callback(null, 200);
};
```

But this will **not**:

```js
module.exports.handler = (event, ctx, callback) => {
    ctx.callbackWaitsForEmptyEventLoop = false;

    setTimeout(() => console.log("timeout"), 10000);

    // or return
    callback(null, 200);
};
```

### Where would I want to use this

This is an ideal scenario to hold open connection to a database. Usually the setup involves having global `db` variable which corresponds to a given connection:

```js
let connection = null;

function connectToDB() {
    if (connection) return connection;
    // connect and save
}
```

Such connections usually push stuff to the event loop, thus preventing your lambda to _finish_.

### What about the `async` handlers

The behavior is a little bit different when it comes to `async` handlers. There is **no waiting for the event loop**. The async handler behaves as if you specified `ctx.callbackWaitsForEmptyEventLoop = false`.
