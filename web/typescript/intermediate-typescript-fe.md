# Intermediate TypeScript FrontendMasters

## Declaration merging

- Multiple things stacking into each other

- Notion of things declared in _type space_ and _value space_. You can have types named the same as your JavaScript identifiers.

```ts
interface Fruit {
  name: string;
}

const Fruit = {
  name: "banana",
};

export { Fruit }; // Both type and JavaScript value
```

- This "overloading" for a given identifier also works with _namespaces_

```ts
interface Fruit {
  name: string;
}

namespace Fruit {
  createFruit(): Fruit {
    return {name: "banana"}
  }
}

const Fruit = {
  name: "banana",
};

export { Fruit }; // Both type and JavaScript value
```

- Please note that **unless explicitly silenced, the above will result in TS errors**.

- You can test whether an identifier is referring to a type or a JavaScript value by using an assignment operator.
  You cannot assign a JavaScript value to a type.

- One notable JavaScript construct is a class. It turns out that the **classes are considered to be both types and values**

```ts
class Fruit {
  public name?: string;
  static getFruit() {
    return { name: "banana" };
  }
}

const fruitClass = Fruit; // no error, it means that the Fruit is a value

const instance: Fruit = {} as any; // no error, it means that the Fruit is a type
```

## Modules and CJS

- Due to historical reasons, a lot of code out there is using CJS. The format Node.js uses but slowly deprecates to ESM (thank god).

### CJS Interop

- Usually you can import CJS using the `import * as foo from require("bar")` syntax which would be equivalent to `const foo = require("bar")`

- Sometimes it's not possible. This might be the case whenever you are not importing a namespace (like `fs`) but a single exported function / variable.
  As an example, one might do `module.exports = function createBanana(){}`. **The exported value is a function, but the `import * as ...` syntax demands the exported variable to be an object**.

- There are two ways you can make the `module.exports = function createBanana(){}` work with TypeScript:

  1. Use the `import createBanana = require("./create-banana")` syntax (not standard but will work)
  2. Use the `esModuleInterop` compiler flag

  Usually, people go for the latter option. While it seems like it's the easiest route to take, it also has it's own drawbacks.
  Mainly - **whenever you set `esModuleInterop` compiler flag, people consuming your code will also have to turn that flag on**.
  Mike calls it a "viral option".

### Namespace imports and bundlers

- The namespace style import (`import * as foo from 'bar'`) is tempting, but if you are worried about the _bundle size_ of your application, you might want to think twice before using this syntax within your code.

- The problem with namespace style imports is that, **usually**, bundlers will eagerly import everything from the file you are referencing.
  The tree shaking might not occur.

## Type Queries

- A way of obtaining a type from a _JavaScript value_

### The `keyof` operator

- Used for retrieving keys of a given interface / type.

- You **can use type primitives for narrowing the collection of keys**. For example `keyof Date & number` would yield all `Date` keys that are numbers.

```ts
type foo = {
  1: number;
  2: number;
  bar: string;
};

type FooNumberKeys = keyof foo & number; // 1 | 2
type FooStringKeys = keyof foo & string; // "bar"
type FooKeys = keyof foo; // 1 | 2 | "bar"
```

### The `typeof` operator

- Used for retrieving a type that describes a JavaScript value.

- Mainly used for creating _type aliases_ for things that are not already described by types.

```ts
async function main() {
  const apiResponse = await Promise.all([
    fetch("http://example.com"),
    Promise.resolve("white"),
  ]);

  type ApiResponseType = typeof apiResponse; // [Response, string]
}
```

- Can be useful for things that are dynamic and change often. Instead of having to change the implementation and the underlying typings,
  the `typeof` can be used. With the typings declared using `typeof`, the typings will be updated automatically as the implementation changes.

## Conditional types

- Ternary operators but used purely in TypeScript world.

- Think of the `extends` keyword as a `>=` comparison (so _this type has to be AT LEAST this type_) and not as an equality check.
  I've also found thinking in terms of narrowing down helpful:

  - _Is the type on the left more specific than the type on the right?_
  - _Is the type on the left more specialized flavour of the type on the right?_

- You should read the condition from left to right and ask yourself: _does type X fits within type Y_.

```ts
type q1 = 64 extends number ? true : false; // Does literal type 64 "is included" in a type number? Sure.

type q2 = number extends 64 ? true : false; // Does number "is included" in a literal type 64  ? Nope. The number type also includes 63, 62 and other numbers
```

## `Extract` and `Exclude`

- The `Extract` utility type is like querying a given type and checking if it contains another type.

```ts
type Foo = number | string | object;

type QueryResult = Extract<Foo, string>; // Does Foo "is included" in a string? Sure, here are the results: string
```

- the `Exclude` is the inverse of `Extract`.

```ts
type Foo = number | string | object;

type QueryResult = Exclude<Foo, string>; // Please give me a type that DOES NOT "include" type string. Sure, here are the results: number | object
```

## Interference with conditional types

- The `infer` keyword can only be used within the context of a _conditional type expression_.

- Used for obtaining a given type. Do not mix the `infer` keyword with the `typeof` keyword. The `infer` keyword is used within the context of generics with conditional types.

```ts
type PromiseResult<T> = T extends Promise<infer R> ? R : never;

const booleanPromise = new Promise<boolean>((r) => r(true));

type Result = PromiseResult<typeof booleanPromise>; // true;
```

- Please note that `infer` and _conditional types_ are not free in terms of performance.

- One example given in the course was obtaining the type of the first constructor argument of a class.

```ts
type FirstConstructorArgType<C> = C extends {
  new (arg: infer A, ...rest: any[]): any;
}
  ? A
  : never;

class MySuperClass {
  constructor(firstParam: string, secondParam: string) {}
}

type Result = FirstConstructorArgType<typeof MySuperClass>;
```

Note that I've used the `typeof` instead of directly using the `MySuperClass`. It's the JavaScript class that is _newable_, not the type.

## Index Access Types

- Used for accessing part of a given type via an _index_.

```ts
type Foo = {
  property: string;
};

type Accessed = Foo["property"]; // string
```

- Not to be mistaken with the _dot notation_. We are not operating on values here. Instead we operate on types.

- **The key can be an union type**.

## Mapped Types

- Characterized by the usage of `in` keyword.

```ts
// Considered a mapped type
type MyRecord = { [NameMattersCanBeReferenced in "apple" | "cherry"]: string };

// Not considered a mapped type
type MyDict = { [key_name_does_not_matter: string]: string };
```

### Key Mappings

- Manipulate the keys during the mapping.

- In terms of _type literals_ you have additional utility types at your disposal.

```ts
type CapitalizeKeys<T, K extends keyof T> = {
  [Key in keyof T as Key extends K & string
    ? `${Capitalize<Key>}`
    : Key]: T[Key];
};

type Properties = {
  prop1: string;
  0: string;
  prop2: number;
};

type Result = CapitalizeKeys<Properties, "prop1">; // {Prop1: string, 0: string, prop2: number}
```

### Filtering properties out

- _Key Mappings_ can be used to filter keys.

- **Filtering on the keys level is usually a much better solution than trying to annotate a property with `never`**.
  This is because annotating with `never` does not work.

```ts
type OnlyFunctions<T> = {
  [Key in keyof T as T[Key] extends Function ? Key : never]: T[Key];
};

type Foo = {
  funcProp: Function;
  numberProps: number;
};

type Result = OnlyFunctions<Foo>; // {funcProps: Function}
```
