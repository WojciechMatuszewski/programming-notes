# FEM TypeScript in the Age of AI

> Notes [based on this course](https://frontendmasters.com/workshops/typescript-ai/).

- Nominal vs. structural typing.

  - TypeScript uses _structural typing_. **The type defines the _shape_ of the value**

  - With **nominal typing, the type defines what the value can hold**.

    ```ts
    class Foo {
      bar: number;
      baz: string;
    }

    // This would not be valid in language that uses nominal types
    const foo: Foo = {
      bar: 3,
      baz: "baz",
    };
    ```

- **Types are sets of valid (matching) values**.

  - The `unknown` type encompasses all types, including `any`. **The `unknown` is so-called _top-type_**.

  - **The `never` type is so-called "bottom" type**. You can think of it as an empty set, since there is "nothing below it" in the hierarchy.

- Types can be _literal values_.

- Since I default to using `const`, I never payed attention to it, but using `let` will use the _widest_ possible type of the value that the variable is initialized with, and `const` will use the _narrowest_ type the value is initialized with.

  - **This only applies to primitive values, for objects, the keys of the object will always infer to the _widest_ possible type**. This makes sense, given that you can override the properties of `const`-defined object.

    - At the time of writing, this applies even when using the `Object.freeze` on the object.

    - Remember about `as const` here!

- Using `typeof <variable>` is quite nice to derive types and prevent duplication in your code.

- Something that seems obvious but I often forget about it.

  **Consider a situation where you want to infer a very complex type based on variable defined in a component as a prop to another component**.

  ```tsx
  // 3rd-party library with a very complex type
  const form = useForm()

  <div>
      <SomeComponent form={form}>
  </div>

    // What is the type of `form` here?
  function SomeComponent({form}: {form: any}) {

  }
  ```

  Of course, one way would be to copy & paste the type you see when you hover over a `form` variable, but that's not ideal.

  **You can solve this problem by introducing a layer of indirection**. In most cases, it's a worthy tradeoff.

  ```tsx
  function useSpecificForm() {
    return useForm()
  }

    const form = useSpecificForm()

  <div>
      <SomeComponent form={form}>
  </div>

  function SomeComponent({form}: {form: ReturnType<typeof useSpecificForm>}) {

  }
  ```

- The same problem as 👆 can apply to a _parameters_ a given function takes.

  - You can use `Parameters<typeof xxx>`.
  - The library author (or you) should export the types!

- Using explicit return types might be beneficial when you are dealing with different return types.

  ```ts
  function foo () {
      if (!success) {
          return ...
      }

      return {
          ...
      }
  }
  ```

  Usually, the resulting type might not be what you want. In such cases, consider typing the return union manually.

- The `class` keyword can also take _type parameters_ (so it's a generic class).

- The `extends` keyword _constraints_ the type argument you can pass to the generic function.

  - Remember about the structural typing and the "sets". You can pass a "higher" type as parameter. You can't pass a "lower" type as parameter.

Consider the following type:

```ts
type LoadingPacket<TArgs extends Array<unknown>> = {
  load: (...args: TArgs) => Promise<unknown>;
};
```

Now, let's implement the code adhering to this type.

```ts
function createLoadingPacket<TArgs extends Array<unknown>>(packet: LoadingPacket<TArgs>) {
  // The return here does not matter.
  return {};
}

// Inferred as
// function createLoadingPacket<[page: number, search: string]>(packet: LoadingPacket<[page: number, search: string]>): {}
const loader = createLoadingPacket({
  async load(page: number, search: string) {
    return {};
  },
});
```

**Notice that TypeScript inferred the args as tuple, instead of `string | number` array**.

Ok, so now let's add another property to our type and attempt to implement it.

```ts
type LoadingPacket<TArgs extends Array<unknown>> = {
  load: (...args: TArgs) => Promise<unknown>;
  getArgs: (cookies: Record<string, string>) => TArgs;
};

function createLoadingPacket<TArgs extends Array<unknown>>(packet: LoadingPacket<TArgs>) {
  // The return here does not matter.
  return {};
}

const loader = createLoadingPacket({
  async load(page: number, search: string) {
    return {};
  },
  // Type Error!
  // Type '() => (string | number)[]' is not assignable to type '(cookies: Record<string, string>) => [page: number, search: string]'.
  // Type '(string | number)[]' is not assignable to type '[page: number, search: string]'.
  getArgs() {
    return [1, "foo"];
  },
});
```

Notice that the return type of the `getArgs` is inferred as a "wider" version of the tuple `TArgs`. Why is that?

**TypeScript will "widen" the inferred return type unless you explicitly annotate it as tuple**.

```ts
type LoadingPacket<TArgs extends Array<unknown>> = {
  load: (...args: TArgs) => Promise<unknown>;
  // Explicitly annotate the return type of `getArgs` as tuple (by using spread operator).
  // Now the types match. The tuple inferred from `load` is the same as the tuple annotated in `getArgs`
  getArgs: (cookies: Record<string, string>) => [...TArgs];
};
```

TODO: Describe the fact that _where_ you put the generic arg, it matters! (end of the part 2)
