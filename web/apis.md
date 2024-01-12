# Useful Web APIs

## `structuredClone`

How many times you had to _deep clone_ an object in JavaScript? Probably quite a lot, right?

```js
const deepClone = (data) => {
  return JSON.parse(JSON.stringify(data));
};
```

We have been taught to use the `JSON` family of APIs for a long time. This is justified because **there were no better tools to do this until recently**.

The `JSON.parse(JSON.stringify)` methods works, **but it has a few key problems**.

- It does not parse `Date` or similar objects.

- It does **not handle circular references**.

- It does not work with [_transferable objects_](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API/Transferable_objects) – those are used to share data between the main thread and workers.

Read more about the `structuredClone` capabilities [here](https://www.builder.io/blog/structured-clone). It has been supported for a while, and, unless you are shipping a very old JS, consider using it whenever you need to deeply clone some data!
