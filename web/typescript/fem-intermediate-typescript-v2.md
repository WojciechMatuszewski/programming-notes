# FEM Intermediate TypeScript V2

> You can find the link to the [course here](https://frontendmasters.com/courses/intermediate-typescript-v2).

## Declaration merging

- Mike showcases how `interfaces` could merge with function declarations and `namespaces`.

  - I always knew that `interfaces` merge together, but I was not aware that this concept also applies to other constructs.

- Mike showcases that **`namespaces` are treated as values by TypeScript compiler**. Check this out

  ```ts
  namespace foo {
    function add(x: number, y: number): number {
      return x + y;
    }
  }

  const as_type: foo = {}; // Cannot use namespace as a type

  const as_value = foo; // This one is okay!
  ```

- Mike showcases one use-case for `namespaces` – how JQuery used to be used.

  ```ts
    function $(selector: string): NoeListOf<Element> {
      return document.querySelectorAll(selector)
    }

    namespace $ {
      export function ajax() {
        return ...
      }
    }

    // So now, you can use the famous `$(..)` syntax
    const foo = $("#bar")

    // Or you can use specific methods on $.
    const result = $.ajax({})
  ```

  In short, using namespaces allows you to encapsulate functionality to a certain identifier and then have it be merged with the same identifier, but implemented in totally different way (in the example above, we are merging function with a concrete `$` value).

  **You can also think of `namespaces` as a tool to organize global values**.

- Have you ever wondered how it is possible to have a class be a type, but also something concrete? You most likely seen things like

  ```ts
  class MyCustomError {
    constructor(private name: string, private statusCode: number) {
      super(name);
    }
  }

  // We are using a class as a value and a type!
  const myError: MyCustomError = new MyCustomError("foo", 403);
  ```

  **This is possible due to declaration merging on the TypeScript compiler level**. This use-case of having the class be a both a type and a value was so common that TypeScript implemented it internally.

## Top and Bottom types

- Those are the types that accept nothing (the bottom) and anything (the top). There are also interesting set of types that land in-between being an _almost top_ or _almost bottom_ type.

- One of the more interesting top types is the `unknown` type.

  - **The difference between `unknown` and `any` is that with `unknown` you cannot USE the value before introspecting it**.

    ```ts
    let value: any = 4;
    value = Function;
    value = document.documentElement;
    const anotherValue = value.i.can.use.it.easily;

    let valueUnknown: any = 4;
    valueUnknown = Function;
    valueUnknown = document.documentElement;
    const anotherValue = value.i.cannot.use; // errors.
    ```

  - One use-case for the `unknown` is error handling. This is because in JavaScript we can throw anything. As such we cannot trust that what we caught in `try/catch` is an `Error` instance.

- The `object` (mind the casing here!) is _almost a top type_. It will accept anything except primitive values.

- Another one is the `{}` type. This one will accept everything except `null` or `undefined`.

  - One use-case for this type is to **remove nullability from the type**.

    ```ts
    type NullableStringOrNumber = string | number | null | undefined;
    type StringOrNumber = NullableStringOrNumber & {}; // string | number
    ```

  - Another use-case for this type is to add autocomplete to a function that takes an union of well defined symbols and a `string` type.

    ```ts
    function foobar1(prop: "bar" | "baz" | string) {}
    foobar1("bar"); // no auto-complete

    function foobar2(prop: "bar" | "baz" | (string & {})) {}
    foobar2("bar"); // with autocomplete!
    ```

    Check out [this StackOverflow reply for more information](https://stackoverflow.com/a/61048124).
