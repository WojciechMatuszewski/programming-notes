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

// TODO
https://www.nearform.com/blog/optimise-node-js-performance-avoiding-broken-promises/?utm_campaign=Adventures%20in%20Nodeland&utm_medium=email&utm_source=Revue%20newsletter
