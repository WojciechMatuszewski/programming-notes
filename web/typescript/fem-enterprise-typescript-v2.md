# FEM Enterprise TypeScript v2

- The `stripInternal` option makes it so any method you marked as `@internal` via JSDoc comment, will NOT emit the types.

  - This means that the customers will not, in theory, be able to import it into their project.

    - One use-case for this are test-only modules.

- Interestingly Mike sets `skipLibCheck` to `false`.

  - The argumentation is that the type-checking should pass for all the `d.ts` file as this is what you are shipping to customers of the package.

    - I tend to agree, but only in the context of a standalone packages. In applications, I would set to `false` to speed up type-checking

- **`yarn 3.x` supports plugins which can aid you when using the `yarn` command**.

  - For example, the project we were working in had a plugin to automatically install `@types/x` dependencies. Pretty neat!

- Some `eslint` rules require TypeScript to work. This makes sense as TypeScript greatly enriches the information linter can depend on.

  - One of such rules is `'plugin:@typescript-eslint/recommended-requiring-type-checking'`

  - [Check out the documentation here](https://typescript-eslint.io/linting/typed-linting/).

- Very interesting choice of using `jest` and `babel` for testing.

  - I think Mike is going for the most battle-tested tools for the workshop.

    - I would personally use `swc` and `vitest` as those tools seem to be much more modern.

      - That being said, `babel` and `jest` is used everywhere so betting on those is pretty safe.

  - **One thing that came to my mind: by having different compilers for tests and build, we are NOT testing what our customers will be using**.

    - It not it bad? I bet there could be bugs we are not aware of, but we would ship them since our tests passed.

- Mike uses a tool called `api-extractor` for documentation, but **also for creating various d.ts files that correspond to "versions" of the package**.

  - The `version` could be "beta", "alpha" or "public".

  - This is pretty neat as it allows you to control what kind of API you are exposing.

    - Think about publishing the "beta" `d.ts` files in a `@beta` release.

  - I have to admit, **the documentation the tool generates is pretty neat**.

- The `alwaysStrict` option in `tsconfig.json` could be misleading.

  - It is not associated with TypeScript at all! It relates to the `use strict` pragma in files.

    - Nowadays, this setting is rarely used. The ES6+ implies the usage of `use strict` pragma – you do not have to specify it.

- The `strictNullChecks` does not seem so "strict" to me. Rather it seems logical.

  - Without this setting, TypeScript will ignore cases where the value could be `null` or `undefined`.

    - A good example would be the `find` function. The `find` function could return the element or `undefined`. If you have `strictNullChecks: false`, TypeScript thinks the value is always there. A bit scarry if you ask me.

  - **Luckily all the settings with `strict` prefix are already included in `strict` setting, so you do not have to worry about them**.

- The `exactOptionalPropertyTypes` is interesting and it deals with the difference between the following definitions.

  ```ts
  interface Person {
    SSN?: number;
  }

  interface Person2 {
    SSN: number | undefined;
  }
  ```

  The difference is that the `Person2` has to have the `SSN` key defined. It could be `undefined`, but the key has to be there. This means that doing `Object.values` or `Object.entires` will yield this key.

  The `Person` is different. There you can set the `SSN` to `undefined`, but you can also completely omit it from the object. If you omit the key, it will not show up when doing `Object.values` or `Object.entries`.

  I much better prefer the `Person2` version. It is much more explicit and enables me to see the WHOLE API.

  Having said that, I deem `Person`-style interfaces much more appropriate for function parameters. Instead of putting the burden of specifying every key on the caller, the can specify only those keys that they need.

- As always, Mike reminds of _viral options_ we would not want to specify.

- Later in the course Mike prefixes interfaces with `I`, for example `IMessage`.

  - I'm unsure I understand the point of this convention. While there are differences between the `type` and `interface` keywords, does it really warrant the convention?

    - Any convention that is not automatically enforced increases cognitive load. I bet there there is a linting rule for this, but does it really matter to disambiguate between the `type` and `interface` keyword?

- TIL that you can explicitly type _type-guards_ as _type parameters_.

  ```ts
  function isTypedArray<T>(
    arr: unknown,
    // A regular function returning a boolean wont do here!
    guard: (element: T) => element is T,
  ): arr is T[] {
    if (!Array.isArray(arr)) {
      return false;
    }

    return arr.every(guard);
  }
  ```

- Later on Mike showcases how to test types.

  - No real changes in-between v1 and v2 of this workshop.

  - In the workshop, we are using `tsd` for testing types.

    - Keep in mind that "testing types" means running `tsc` rather than any kind of test runner.
