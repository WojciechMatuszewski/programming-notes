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

## EMF Connecting to the X-Ray daemon, wtf?

This will happen to you whenever you run the logger locally and not set the environment override.
The EMF logger will _probe_ given environments starting from the _Lambda_ one. If some environment variables cannot be found, the next one is _probed_.

It just so happens that the **default environment is using CloudWatch Agent**. Since you did not override anything, the default environment was selected.

This results in you having weird _cannot connect to..._ logs in the console.
To let the logger know that it is running in the local environment, specify `process.env.AWS_EMF_ENVIRONMENT = "Local";`

## Idempotency tokens (ClientRequest Tokens)

While using the AWS SDK you might have noticed that there is this parameter called _ClientRequest Token_. This parameter is responsible for making your calls idempotent.
If you are using the AWS CLI, the CLI is providing that tokens for you automatically.

How this token should be created? You will probably need to combine input parameters into 1 value and compute a token based on that.
Remember, if you provide the same token, the operation will basically be a noop one. The same (with some exceptions regarding WCU/RCU consumption on DDB calls) response will be returned as the original operation (I'm not sure about errors though).

Some services implement a window, e.g. 10 minutes in the case of DDB _TransactWrite_ call, during which if you pass the same token, the operation is guaranteed to be idempotent.

While _ClientRequest Tokens_ are not something that is exclusive to the Node.js SDK.

## SourceMaps - use them

Adding sourcemaps to your lambda functions will help you while debugging. Instead of having to read cryptic stack traces, you will be able to pin-point the file where the issue is much faster.

At the time of writing this, there are 2 ways you could add sourcemaps to lambda functions.

1. Use the `source-map-support/register` with an `inline` or `external` source map

2. Use the `--enable-source-maps` Node.js option

The option no 1. seems to be a better choice since the `--enable-source-maps` is still experimental.

### Adding `source-map-support/register` automatically

The process of adding the import is very human-error prone. To ensure that you have sourcemaps hooked up within every handler you write, you will need to automate the process.

I will be focusing on `aws-cdk` and `esbuild`. IMO there is no point building your infrastructure in any other way (maybe except using a different tool than `aws-cdk`)

#### Unix commands magic

Currently, the `aws-cdk` _NodeJSLambda_ construct does not allow you to write _plugins_ for `esbuild` (but `esbuild` supports plugin API).
To automatically add the import we need, we have to do some manual work.

```ts
const outpath = `./dist/functions/${id}-${scope.node.uniqueId}`;

try {
  require.resolve("source-map-support");
} catch (e) {
  throw new Error("package `source-map-support` is required when using this Construct");
}

/**
 * Add the `source-map-support/register import automatically to the handler
 */

// Create a copy of the original file (backup). This is so that we can preserve the code within the original `entrypoint` file.
const entrypointBackup = `${props.entrypoint}.tmp`;
execSync(`cp -fp '${props.entrypoint}' '${entrypointBackup}'`);

// Add the import statement to the original handler
execSync(`echo 'import "source-map-support/register"' | cat - '${entrypointBackup}' > '${props.entrypoint}'`);

const cmd = `npx esbuild --bundle --target=es2019 --platform=node ....MORE COMMANDS`;

if (props.debug) {
  console.log(cmd);
}

try {
  childProcess.execSync(cmd);
} finally {
  // Replace the original handler with the backup (pre-transformation handler). The handler was compiled with the `source-map-support/register` import
  execSync(`mv -f '${entrypointBackup}' '${props.entrypoint}'`);
}
```

Here is a snippet that adds the import automatically, builds the handler with it, then replaces the handler with non-augmented code (without the import).

Notice the `finally` block, it is crucial. We have to ensure that errors from `esbuild` are forwarded (thus no `catch` block is present) as well as we cleanup after ourselves.
