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

- **Where you put the generic parameter matters!**

Let's say you want to write a function that takes a selector and returns a result.

```ts
function useStateSlice(selector: (state: State) => ??) {}
```

So what should we do here? What should we write in place of `??`

We can start by using `unknown`. It's not optimal since TypeScript won't be able to infer the result.

```ts
function useStateSlice(selector: (state: State) => unknown) {}

// Unknown.
const { bar } = useStateSelector((state) => {
  return {
    bar: "123",
  };
});
```

Next, you can attempt to add the generic parameter at the level of the selector.

```ts
function useStateSlice(selector: <TReturn>(state: State) => TReturn) {}

const { isOpen } = useStateSelector((state) => {
  // Type '{ isOpen: boolean; }' is not assignable to type 'TReturn'.
  // 'TReturn' could be instantiated with an arbitrary type which could be unrelated to '{ isOpen: boolean; }'
  return { isOpen: state.isOpen };
});
```

The problem with this approach is that **the caller does not control the type**.

Think about it, the _real_ return value could be changed _within_ the caller of the `selector` – that is why TypeScript is complaining here.

If you declare the generic at the level of the caller though, all will work as expected.

```ts
function useStateSlice<TReturn>(selector: (state: State) => TReturn) {}

// {isOpen: boolean }
const { isOpen } = useStateSelector((state) => {
  return { isOpen: state.isOpen };
});
```

- You can narrow unions down to a single type via

  - The `in` keyword, checking for a property. `if ('meow' in cat)`

  - The `typeof` keyword.

  - The `instanceof` keyword.

  You **can also use a discriminant property if you are dealing with objects which might be the best way to go about it**.

  ```ts
  type cat = {
    kind: "CAT";
  };

  type dog = {
    kind: "DOG";
  };

  type animal = cat | dog;

  // Now you can check the `kind` property. TypeScript will narrow automatically.
  ```

  - Also, there are serval "helper" functions that you might find useful, like `Array.isArray` and `Number.isNaN`

- You can ensure that you exhausted all possible cases of value by assigning to `never` at the end. This is useful for `switch` statements as well as `if` statements

  ```ts
  switch (foo) {
    case "bar": {
    }
    default: {
      const exhaustiveCheck: never = foo;
    }
  }
  ```

- The `satisfies` keyword is pretty neat. Use it to **verify that a given object _satisfies_ a given type, but do NOT allow any additional extra properties**.

  The satisfies will do the narrowing for you, especially useful for `string`-based properties that have a literal type equivalent in the type.

- Conditional types and unions are pretty interesting.

  Consider the following code:

  ```ts
  type IsArray<T> = T extends Array<infer I> ? I : never;
  // The `Weird` return `number`
  type Weird = IsArray<Array<number> | string>;
  ```

  The **conditional type will "run" with all the members of the union separately and then "union" the result type at the end**

  Since the `number | never` is `number`, that is what you see returned.

- Conditional types are pretty useful for writing _tests for types_.

  The idea is to create a conditional type that returns either `true` or `false` and then assign the result of the "type test" to a variable.

  If the result type is different, you will get a TypeScript error.

- You can _overload_ functions in TypeScript. Keep in mind that **function overloading is a "type-land" feature, not a _runtime_ feature**.

  No matter how many overloads you have, you have to have a single implementation that handles those overloads.

  There is no runtime feature that would "pick" the matching implementation for you!

- **Function overloads are useful in preventing "impossible parameters"**.

  Consider the following function:

  ```ts
  function useSearch<T extends string | number>(arg: T): T extends string ? Array<User> : User {}

  const stringOrNumber: string | number = Math.random() < 0.5 ? "foo" : 1;

  // ??
  useSearch(stringOrNumber);
  ```

  Notice that the `stringOrNumber` is an "impossible" parameter from the types perspective.

  What if we used function overloads here?

  ```ts
  function useSearch(arg: string): Array<User>;
  function useSearch(arg: number): User;
  function useSearch(arg: string | number): Array<User> | User {}

  const stringOrNumber: string | number = Math.random() < 0.5 ? "foo" : 1;

  // TypeScript throws an error!
  useSearch(stringOrNumber);
  ```

  TypeScript will throw an error. That's much better!

Start Part 5 19:32
