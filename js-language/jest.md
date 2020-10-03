# Random `jest` bits

## `mockName` property on a mock

This one is something that is often overlook, but may help you debug failing tests.
While we mock, we often do not think about how your test output will look like if the tests fail. This is a mistake! The more information (useful) information you can provide while your test fails, the better the chances you will fix it in time.

This is where `mockName` on `jest.fn` comes in.

When you create a `spy` you often end up writing something like this:

```js
const spy = jest.fn(); // maybe with mockImplementation
```

While this is completely valid way of using `jest.fn` when the assertion on that spy fails, the output will look similar to this

```js
expect(jest.fn()).not.toHaveBeenCalledWith(...expected);
```

Not that good right? When you have multiple mocks, this can hinder your ability to debug failing test.

You can add `mockName` to the `jest.fn` so that the output will look similar to this

```js
// declaration
const spy = jest.fn().mockName("json");
// failing output
expect(json).not.toHaveBeenCalledWith(...expected);
```

Much better!

## Snapshots

Snapshots **can be useful** but only **when used correctly**. There are many traps you can fall into while using snapshots for your tests, but where snapshots shine are **error messages**.
In such cases snapshots can be really helpful. **Use snapshots for error messages, otherwise avoid them**.

### Snapshots with dynamic values

Imagine this: you are trying to delete an item from a database, and the `id` that your provided does not exist. The error message might contain that `id`. Within tests, you most likely generate that `id` at random, so whenever you try to create an `inlineSnapshot` your tests will fail.

The solution here is pretty simple. Just replace the `id` with something static

```js
const id = "1"; // from the item
const idlessMessage = error.message.replace(id, "SOME_ID");
expect(idlessMessage).toMatchInlineSnapshot();
```

## Mocks on object which only have getters

This usually happens when you have a file which re-exports methods from other files, a barrel module if you will.

```js
export * from "./foo";
```

This is the crux of the issue since this file will be transpiled to an object with only `get` property defined. **Without the `set` property *jest*s `spyOn` does not work**.
To combat this, mock the entire module, probably using `jest.requireActual` along the way.

```js
jest.mock("module", () => ({
  ...jest.requireActual("module"),
  method: jest.fn(),
}));
```
