# Errors

Dealing with errors in any language is critical. Some of the most hard-to-debug bugs I've came across were related to error handling (or lack of it).

## The `cause` property AKA error wrapping

In Golang, we can "wrap" errors using the `fmt.Errorf`, like so:

```go
result, err := downloadFile()
if err != nil {
    return fmt.Errorf("error downloading file: %w", err)
}
```

We can also use our own custom types.

```go
type ErrUserNotFound struct {
    User string
}

func (e ErrUserNotFound) Error() string {
    return fmt.Sprintf("user %q not found", e.User)
}
```

---

In Rust, we can either "box" the errors, or create our own custom types.

- [Boxing errors](https://web.mit.edu/rust-lang_v1.25/arch/amd64_ubuntu1404/share/doc/rust/html/rust-by-example/error/multiple_error_types/boxing_errors.html).

- [Wrapping errors with custom types](https://web.mit.edu/rust-lang_v1.25/arch/amd64_ubuntu1404/share/doc/rust/html/rust-by-example/error/multiple_error_types/wrap_error.html).

---

In JavaScript, we _finally_ have the ability to "wrap" the errors similar to how we do it in Golang.

```js
const err1 = new Error("first");

const err2 = new Error("second", { cause: err1 });
```

**It is imperative that your logger can correctly "resolve" the cause chain** – you can nest the `cause` keyword usage.

```js
const err1 = new Error("first");

const err2 = new Error("second", { cause: err1 });

const err3 = new Error("third", { cause: err2 });
```

Using the `cause` property is a good way to "enhance" the original error with more context.

Of course, the longer the chain, the bigger the stacktrace. This might or might not have an impact on your observability system. Perhaps you want to truncate those?

## The `Error.isError`

> [Based on this blog post](https://allthingssmitty.com/2026/02/23/from-instanceof-to-error-iserror-safer-error-checking-in-javascript/).

In JavaScript, you can `throw` anything.

```js
throw 3;
throw "foo";
```

But most of our code operates on the notion that we are throwing _instances of_ the `Error` class.

```js
try {
  await someFunction();
} catch (error) {
  if (error.message.includes("....")) {
    // do something
  }
}
```

When you are using TypeScript, you _have_ to perform some runtime checks on the `error` because it is typed as `unknown` by default (a very good decision by the TypeScript team!)

In most cases, those runtime checks are:

1. Quite complex.
2. Not comprehensive enough.
3. Repeated through the codebase.

**We now have the `Error.isError` at our disposal** which can help us with all three of those issues!

```js
Error.isError(new Error("Oops!")); // true
Error.isError(new TypeError("Bad type")); // true
Error.isError("just a string"); // false
Error.isError({ message: "Not really an error" }); // false
Error.isError(Object.create(Error.prototype)); // false
```

Please check out the article for more information. The most important, to me, part is that you can reliably use this method to check if a variable is an error.
