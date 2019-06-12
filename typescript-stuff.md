# Typescript Stuff

## Pick and Exclude

### Pick<T,K>

> From T, pick a set of properties whose keys are in the union K

You can, well, _Pick_ specific properties from an interface.

```typescript
interface User {
  name: string;
  email: string;
  password: string;
}

type OnlyName = Pick<User, "name">; // {name: string}
```

### Exclude<T, U>

> Exclude from T those types that are assignable to U

It works by diffing two types. That's a common gotcha (at least for me).

```typescript
interface User {
  name: string;
  email: string;
  password: string;
}

// remember, we are diffing 2 types, code below will not do what we want it to do
type WithoutName = Exclude<User, "name">; // User, because T is not extending U so Exclude returns User

// now we are talking, we are diffing each key with 'name'
type WithoutName = Exclude<keyof User, "name">; // 'email' | 'password'
```

## Combining Exclude And Pick

With combine power of `Exclude` and `Pick` we can do some nice stuff (especially with HOC's). Let's say we want to remove a prop from something in a generic way.

```typescript
// from Root Pick...
type Omit<Root, PropsToOmit> = Pick<
  Root,
  // Exclude these props from Root which can be found in PropsToOmit
  Exclude<keyof Root, keyof PropsToOmit>
>;
```

### Caution warning

What happens if `PropsToOmit` is a single value, let's say `string`. Well then bad things will happen. `keyof string` will actually look at it's prototype chain.

```typescript
type Test = keyof "something"; // "toString" | "charAt" | "charCodeAt" | "concat" | "indexOf" | "lastIndexOf"
```

## Return Type

You can actually get the return type of a given function. Quite neat!

```typescript
function add(x: number) {
  return x + 1;
}

type SomeType = ReturnType<typeof add>; // number
```

### Caution warning

- you have to cast directly if you want to get _value_ as the type

```typescript
const None = "None";
function something() {
  return { x: None };
}
type HeyTs = typeof None; // "None"
type SomeType = ReturnType<typeof something>; //{x: string}

function something() {
  return { x: None as typeof None };
}
// .. same operations
// return type is now {x: "None"}
```

- you have to create helpers when dealing with generic functions

```typescript
function identity<T>(prop: T): T {
  return prop;
}
type SomeType = ReturnType<typeof identity>; // unknown
```

Now with helpers (we are creating those because `ReturnType` does not take second argument)

```typescript
type Callable<T> = (...args: any[]) => T;

type MyOwnReturnType<ReturnType, F> = F extends Callable<ReturnType>
  ? ReturnType
  : never;

type SomeType = MyOwnReturnType<string, typeof identity>; // string
```

## Conditional Types

### Very basic example

Just like ternary but with types

```typescript
type SomeType<T> = T extends string ? T : never;
interface Obj {
  name: string;
  age: number;
}

type t1 = SomeType<"ala">; // 'ala'
type t2 = SomeType<Obj["name"]>; // string
type t3 = SomeType<Obj["age"]>; // never
```

Using this feature you can create (most of them already ship by default) types

```typescript
// already built-in
type NonNullable<T> = T extends undefined | null ? never : T;
type t1 = NonNullableMy<null>; // never
```

### Inferring the type

This feature is very useful. Basically you can `pluck` a type from generic using conditional types.

**We can place the infer keyword at the position where we want the type to be inferred.**

```typescript
// you could also do args: infer U
type GetFunctionArgumentTypes<F> = F extends (...args: Array<infer U>) => void
  ? U
  : never;

function numberArg(x: number) {}

function arrayMixed(x: [1, "a", {}]) {}

type t1 = GetFunctionArgumentTypes<typeof numberArg>; // number
type t2 = GetFunctionArgumentTypes<typeof arrayMixed>; // [1, 'a', {}]
```
