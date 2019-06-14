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

## Mapped And Lookup Types

You can use `in` and `keyof` to transform interfaces and type-objects

```typescript
type MyPartial<Type> = { [Key in keyof Type]+?: Type[Key] };

interface Something {
  id: number;
  name: string;
  property?: string;
}
type MyPartial<Type> = { [Key in keyof Type]+?: Type[Key] };

type Test = MyPartial<Something>;
/*
  {
    id?: number | undefined
    ...
  }
*/
```

### Keyof and `[keyof]`

Differences are quite big

```typescript
type Test1 = keyof Something; // "id" | "name" | "property"
type Test2 = Something[keyof Something]; // string | number | undefined
```

It's very similar to accessing object values and `Object.keys` in JS.
**It's just that the value is the type itself**

```js
var someObj = {
  prop1: 1,
  prop2: 2,
  prop3: "someString"
};

Object.keys(someObj); // 'prop1' , 'prop2' ...
someObj["prop1"]; // 1
```

### Caution warning

Sometimes typescript is very strange.
It seems that `prop?: number` is not the same as `prop: number |undefined`?.
Let's consider the following

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type UndefinedAsNever<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Type[Key]
};

type Test1 = UndefinedAsNever<Something>;
/*
WTF ???
{
    id: number;
    name: string;
    property?: undefined;
}
*/
// you can merge interfaces btw :)
interface Something {
  id: number;
  name: string;
  // changed this
  property: string | undefined;
}
type Test1 = UndefinedAsNever<Something>;
/*
WEIRD STUFF HUH?
{
    id: number;
    name: string;
    property: never;
}
*/
```

### Plucking nullable (also undefined) keys

Let's say you have an interface

```typescript
interface Something {
  id: number;
  name: string;
  // we want to remove this \/
  property: string | undefined | null;
  // can also be written like this
  property?: string;
}
```

First thing first we probably should _mark_ `property` somehow so that we know that we want to _pluck_ this prop.

Remember our `[keyof Something]` notation?

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
}[keyof Type];

type Test1 = RemoveUndefinableKeys<Something>; //"id" | "name" | undefined
```

How does `RemoveUndefinableKeys` work?

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
};
// would return
/*
  {
    id: "id",
    name: "name",
    property: undefined
  }
*/
```

Now we _marked_ `property` as the one to be deleted (by undefined type)

Let's add `[keyof Something]` notation (we will basically get only the values from the interface).

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
}[keyof Type];
// would return "id" | "name" | undefined
```

See ? no `property` prop.

We can also do this

```typescript
interface Something {
  name: string;
  age: number;
}

type Identity = { [Key in "name" | "age"]: Something[Key] };
// would return the same Something type
```

So by marking `property` as `undefined` we basically _plucked_ it from the interface.

No we just need to make `Identity` type generic and name it somehow.

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
}[keyof Type];

type RemoveUndefinable<Type> = {
  // this is the same as Key in "id" | "name" | undefined
  // undefined will be omitted
  [Key in RemoveUndefinableKeys<Type>]: Type[Key]
};

type Test = RemoveUndefinable<Something>;
/*
would return
{
  id: number;
  name: string;
}
*/
```
