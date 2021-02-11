# Typescript Stuff

## `tsconfig.json`

### `esModuleInterop` shenanigans

The whole purpose of this option is to enable you to write `ESM` compliant imports when you are using `CJS` modules.

```ts
// instead of this
const express = require("express");

// you can import it this way
import express from "express";
```

The problem starts whenever you generate a _declaration file_ from a `ESM` module transpiled down to `CJS`.

Let's say your module looks as follows

```ts
// foo.ts
export function foo() {
  return 1;
}

// index.ts
export * from "./foo";
```

The _declaration file_ would look like this

```ts
export * from "./foo";
```

So it does look like a regular `ESM` barrel file. Nothing wrong with that right?
Well, **if you have `esModuleInterop: true`, TypeScript will not complain at you, if you import modules as if they had `export default` defined but in reality they do not**.

The `index.ts` clearly does not have `export default` defined, nor the _declaration file_. Well imagine my surprise when having something like this

```ts
import lib from "lib/foo";
```

does not result in TypeScript errors. It should, because if you do

```ts
import lib from "lib/foo";
lib.foo();
```

you will be greeted with a runtime error.

This whole issue stems from the fact that **sometimes TypeScript cannot be sure if _synthetic default exports_ should be allowed for a given _declaration file_**

As a rule of thumb, you should always check how the declaration file is looking before attempting to import any 3rd party module. **Test might not help you. If you forgot to set `__esModule:true` while mocking, the wrong import will still work due to interop settings**.

## _type space_ vs _value space_

While working with Typescript, you will be operating in 2 different _spaces_.

**The _type space_ is where the typings live**

**The _value space_ is where the JS constructs live**

It is important to be aware of this. You might see code that looks as follows

```ts
// declared in the 'type space'
type Person = { age: number };

// declared in the 'value space'
const Person = { firstName: "Wojtek" };
```

Notice that we have a sort of overlap here. You might expect some kind of shadowing issue to occur, but that is not the case. These 2 declarations are _isolated_ from each other, they live in 2 different _spaces_.

## Augmenting global declarations

Lets say you are building a NodeJs app and you want to have strongly typed `process.env` object.

All you have to do is to create some `.d.ts` file (could be `.ts` file but I would go for `d.ts` for clarity) and use the fact that **namespaces are merged just like interfaces**.

```ts
namespace NodeJS {
  interface ProcessEnv {
    MY_GLOBAL_ENV_VARIABLE: string;
  }
}
```

That is all.

### My augmentations does not work when I import something to the declaration file

> By default when you start typing code in a new TypeScript file your code is in a global namespace

This means that you can influence types (make augmentations) which apply globally.

If you add an `import` statement to your declaration file though

```ts
import type fs from "fs";
namespace NodeJS {
  interface ProcessEnv {
    MY_GLOBAL_ENV_VARIABLE: string;
  }

  interface global {
    myGlobalFS: fs;
  }
}
```

It will not work, as in syntax like this

```ts
const test = global.myGlobalFS; // `myGlobalFS` is not defined on the global
const globalVariable: string = process.env.MY_GLOBAL_ENV_VARIABLE; // `MY_GLOBAL_ENV_VARIABLE` is a string | undefined, not string
```

This is because **the act of adding an import, made the typings in the file module scoped**

### Global scope and importing types

What if you want to import types from some module and augment global namespace at the same time?

Not everything is lost, there are 2 ways of doing so.

1. Scope the import to the `namespace` / `module` declaration
   Instead of using a _top-level_ import, use the import inside the declaration itself

   ```ts
   declare namespace NodeJS {
     // You can also use `import type` syntax here
     import fs from "fs";
     interface Global {
       globalReadStream: fs.ReadStream;
     }
   }

   // in another file
   global.globalReadStream; // ReadStream
   ```

2. Use the dynamic `import` syntax, also scoped to the `namespace` / `module` declaration

   ```ts
   declare namespace NodeJS {
     interface Global {
       globalReadStream: import("fs").ReadStream;
     }
   }

   // in another file
   global.globalReadStream; // ReadStream
   ```

### Do I have to write `declare XXX`

You might have noticed that I'm using the `declare module` or `declare namespace` syntax while working with declaration files.

Usually, you would use the `declare module` or `declare namespace` syntax to tell TypeScript compiler that a given variable / class etc.. is declared somewhere else, probably in a `.js` file.

Since Node.js has it's own typings defined already, you can skip the `declare` keyword and rely on _declaration merging_ while augmenting Node.js globals, but it may not be the case for a 3rd party library that does not have TypeScript typings.

```ts
declare module "3rd-party-lib" {}
```

As a rule of thumb I'm always sticking to `declare module` syntax, just to make things consistent

### Typescript ignores my `d.ts` file

First of all check if that file matches the `include` pattern that you specified within your `tsconfig`.

If that's the case, we are dealing with something very strange that I've discovered only recently.

You **have a file named the same way as you `d.ts` file**, eg. `env.ts` and `env.d.ts` file.
The way typescript works is that **the `env.d.ts` file will be ignored since typescript things it was derived from `env.ts` file**. Pretty strange right?

https://github.com/microsoft/TypeScript/issues/31397#issuecomment-492269754

There are 2 solutions here:

1. Rename your `d.ts` file
2. Specify the `d.ts` file within the `file` block inside your `tsconfig`.

## TripleSlash aka Reference

You have seen them, the weird `/// <reference types|lib=...>` syntax. This is mostly relic of the past but still can be useful in day-to-day work.

So, before `tsconfig.json` existed, you had to use the `/// <reference types|lib= >` syntax to tell the compiler

> hey, whenever you parse this file, also include _the files I referenced_ in the compilation

Nowadays, we would most likely use the `include` property in the `tsconfig.json` to do so.

But, there are some use cases for them, mainly when you want to make sure that given _declarations_ are imported when the compiler imports your file.
So let's say you want to augment the `process.env` typings. This is done by creating a _declaration file_ and _extending the ProcessEnv interface_

```ts
declare namespace NodeJS {
  interface ProcessEnv {
    MY_VALUE: string;
  }
}
```

And this is completely fine, if you do it for your project only, and no one is going to consume it, you might leave it as it is. But sometimes when you are exposing your project as a package, you might want to separate your project typings from the global augmentations

```ts
// index.d.ts
interface MyLibrary {}

// global.d.ts
declare module NodeJS {
  // code
}
```

Usually, only the `index.d.ts` is consumed, with the `/// <reference types|lib="code">` syntax you can include the `global.d.ts` file inside the `index.d.ts` file

```ts
// index.d.ts
/// <reference path = "./global.d.ts">

interface MyLibrary {}
```

You are doing this because **there is nothing to import from the `global.d.ts` file**. You could import the file itself, but the `/// <reference types|lib|path = "">` syntax is more common.

## Type-Only imports and import elision

Being able to import a concrete `JavaScript` implementation and the type in the same import statement is awesome

```ts
import { doThing, OptionsType } from "./foo.js";
```

When bundling, babel (or ts) will get rid of that import, everything should be working fine. This process is working because of something called _import elision_. (That's the removal of ts-typings from the file part).

But sometimes, these imports can get a bit ambiguous

```ts
import { Something } from "./module.js";

export { Something };
```

Is `Something` a type or a concrete JavaScript code?. Before `TypeScript 3.8` this situation could result in bugs and wrong code being emitted (looking at you `firebase`).

There was a way to guard yourself against such behavior using special `import` syntax. Lets

```ts
const Something: import("./module.js").Something;

export { Something };
```

This basically tells `TypeScript` that `Something` is only a type.

With `TypeScript 3.8` there is another, more streamlined way of doing this, somewhat taken from `flow`. There will be new `import type` syntax

```ts
import type { Bar, Baz } from "module";
```

Now, there are some restrictions to it, one of which is that you cannot mix default and named exports. This is to ensure that the import statements are non-ambiguous

```ts
import type Foo, { Bar, Baz } from "module";
// ^ this will fail
```

## Testing

What can we test in the realm of TypeScript? Code ? sure, but what about the
Types.

Is it possible to write units for TypeScript Types? Well indeed it is possible.

There are number of libraries that check types, either with special _generic
type_ or a comment inside the code.

The strategy is as follows:

- write a test for type declaration -> this just means using that function on a
  very simple data (but that data has to check type-implementation)
- run `tsc` on test files

### dtslint

This tool was built by Microsoft. A sample test:

```ts
var stooges = [
  { name: "moe", age: 40 },
  { name: "larry", age: 50 },
  { name: "curly", age: 60 },
];
_.pluck(stooges, "name"); // $ExpectType string[]
```

In this example we are using a special comment that ensures that this is the
return type.

You can even test with different versions of TS:

```ts
// TypeScript Version: 2.1
export function pluck<K extends keyof T, T>(array: T[], key: K): Array<T[K]>;
```

[More on this topic](https://medium.com/hackernoon/testing-types-an-introduction-to-dtslint-b178f9b18ac8)

## Assert Signatures

So with `type guards` you are returning `true` or `false`. This then determines
the outcome of the type.

But `assert signatures` **are quite different**

### Different schematics

Type Guard:

```ts
function isDefined<T>(x: T): x is NonNullable<T> {
  return x != undefined;
}
```

Assert Signature:

```ts
function isDefined<T>(x: T): asserts x is NonNullable<T> {
  if (x == undefined) {
    throw AssertionError("Not defined!");
  }
}
```

The signature differs greatly and there is actually more to the
`assert signatures` signature than presented here.

### Two types of Assert Signatures

There are actually 2 variants

- for checking a condition
- for telling TypeScript that specific variable or property has a different type

So it all basically boils down to that `Assert Signatures` does not return
anything, they throw this `AssertionError` whenever something is wrong.

`Type Guards` on the other hand return `true` or `false` based on they inputs.

The signature `asserts something` or `asserts x is something` tells the reader
of the code that **that function will only return if the assertion holds**

```ts
function checkIfString(input:any) asserts input is string {
    if (typeof input != 'string') throw Error('must be a string')
}

function doSomething(val: number | string) {
    checkIfString(val)
    val // string here!
}
```

## Nullish Coalescing

This is more of a JavaScript thingy but hey, we are all probably writing only
TypeScript now :)

So, do you remember the deal with `&&` and `||` ?

- with `&&` you guard the right value with left value (checking _truthiness_)

- with `||` you either will get left or right value depending on their
  _truthiness_

And _truthiness_ is the key-word here.

So the deal with `Null Coalescing` is that it only checks for `null` and
`undefined`

```js
console.log(0 || "something"); // something
console.log(0 ?? "something"); // 0
```

This can help in cases where you have valid non-truthy values as your _guardian
values_ but you still want to check for `null` and `undefined`

## Logical assignment operator

Just like the previous feature, this one is more of a JavaScript thingy. With _nullish coalescing_ you are returning given value,
with _logical assignment operator_ you can assign given value using the nullish operators.

```ts
type Obj = {
  prop: {
    value?: string;
  };
};

function doWork(obj: Obj) {
  obj.prop.value ??= "default value";
  return obj;
}
```

I personally do not use this feature that often but, nevertheless I think its nice to have.

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

With combine power of `Exclude` and `Pick` we can do some nice stuff (especially
with HOC's). Let's say we want to remove a prop from something in a generic way.

```typescript
// from Root Pick...
type Omit<Root, PropsToOmit> = Pick<
  Root,
  // Exclude these props from Root which can be found in PropsToOmit
  Exclude<keyof Root, keyof PropsToOmit>
>;
```

### Caution warning

What happens if `PropsToOmit` is a single value, let's say `string`. Well then
bad things will happen. `keyof string` will actually look at it's prototype
chain.

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
type SomeType = ReturnType<typeof something>; // {x: string}

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

Now with helpers (we are creating those because `ReturnType` does not take
second argument)

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

This feature is very useful. Basically you can `pluck` a type from generic using
conditional types.

**We can place the infer keyword at the position where we want the type to be
inferred.**

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

#### Inferring multiple types

Nothing is stopping your from using the _infer_ keyword multiple times. Check this out

```ts
type AppendArgument<Fn, A> = Fn extends (...args: infer Args) => infer R
  ? (...args: [...Args, A]) => R
  : never;
```

Here I've used _infer_ to both get the hold of the function arguments, but also the return type of the `Fn` type. While I could use the `ReturnType` generic, using _infer_ is also nice ;)

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

### Altering mapped types keys

You can use `as _something_` syntax to alter (probably perform conditional operation) on the _mapped types_ keys.

```ts
type NoNumbers<T extends Record<string, unknown>> = {
  [K in keyof T as T[K] extends number ? never : K]: T[K];
};

type Test = NoNumbers<{ prop1: string; prop2: number }>; // {prop1: string}
```

As of writing this, this is relatively new addition to the language.

### Keyof and `[keyof]`

Differences are quite big

```typescript
type Test1 = keyof Something; // "id" | "name" | "property"
type Test2 = Something[keyof Something]; // string | number | undefined
```

It's very similar to accessing object values and `Object.keys` in JS. **It's
just that the value is the type itself**

```js
var someObj = {
  prop1: 1,
  prop2: 2,
  prop3: "someString",
};

Object.keys(someObj); // 'prop1' , 'prop2' ...
someObj["prop1"]; // 1
```

### Caution warning

Sometimes typescript is very strange. It seems that `prop?: number` is not the
same as `prop: number | undefined`?. Let's consider the following

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type UndefinedAsNever<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Type[Key];
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

This difference stems from the fact that `type | undefined` allows for value to be _skipped_.

```ts
declare function foo1(prop: number | undefined);
foo1(); // Ok.
```

While the optional parameter syntax does not

```ts
declare function foo1(prop?: number);
foo1(); // Error!
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

First thing first we probably should _mark_ `property` somehow so that we know
that we want to _pluck_ this prop.

Remember our `[keyof Something]` notation?

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
}[keyof Type];

type Test1 = RemoveUndefinableKeys<Something>; // "id" | "name" | undefined
```

How does `RemoveUndefinableKeys` work?

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
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

Let's add `[keyof Something]` notation (we will basically get only the values
from the interface).

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
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

So by marking `property` as `undefined` we basically _plucked_ it from the
interface.

No we just need to make `Identity` type generic and name it somehow.

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
}[keyof Type];

type RemoveUndefinable<Type> = {
  // this is the same as Key in "id" | "name" | undefined
  // undefined will be omitted
  [Key in RemoveUndefinableKeys<Type>]: Type[Key];
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

### Record

Remember typing dictionaries awkwardly like:

```typescript
type Dict = { [key: string]: number };
```

Ahh... sad times. As you want to be _the cool kid_ you probably should use this
_leet hackrzz_ `Record` stuff :

```typescript
type Dict = Record<string, number>;
```

**Remember that in Javascript all keys are `string`'s !!**

## This keyword

When using strictest possible typescript settings (as you always should) you
might night to type `this` keyword. Let's see how this can be done:

```typescript
interface SomeObj {
  someFn: (num: number) => number;
  numberToAdd: number;
}

const someObj: SomeObj = {
  someFn,
  numberToAdd: 4,
};

function someFn(num: number) {
  return num + this.numberToAdd; // might cause an error
  // typescript sometimes has problems with inferring the right this
}

// much better implementation would be
function someFn(this: SomeObj, num: number) {
  return num + this.numberToAdd; // much better, we even get autocomplete
}
```

This may look weird, it may seem like `someFn` now takes 2 arguments but that's
not the case. First argument (`this` typing) will get compiled away.

### `ThisType`

This one is pretty wild. This utility type allows you to ensure that all methods on a given object have correct this.

```ts
const someObj: Record<string, any> & ThisType<{ foo: string }> = {
  method() {
    // strongly typed this!
    console.log(this.foo);
  },
};
```

The implementation is also interesting..., being an empty interface. Kinda magical if you ask me.

## Typeof

Here there is distinct difference between Javascript world and Typescript world.

When using Javascript `typeof` will return underlying type as in the type that
you can create in vanilla Javascript. This is familiar territory

```js
typeof []; // "object"
typeof "something"; // "string"
typeof 3; // "number"
```

But in Typescript `typeof` behaves a little bit differently. Instead of
returning underlying vanilla JS types it will return us the Typescript type.

```typescript
const person = {
  age: 22,
  name: "Wojtek",
};

type Person = typeof person; // {age: number, name: string}
```

This is very powerful especially with `ReturnType`.

## Type Guards

Using a _Type Guard_ you can tell Typescript which type something is.

### Typeof Type Guard

This is very simple guard. Check this out:

```typescript
function someFn(arg: number | string) {
  if (typeof arg == "number") {
    // typescript knows we are dealing with a number here
    return arg.toExponential(); // ok
  }
  // typescript knows we are dealing with a string here
  // BUT BEWARE!
  // if we did not return above type would be number | string here
  arg.toLowerCase(); // ok
}
```

### `instanceof` Type Guard

#### In vanilla JS

> it tests if a `.prototype` property of a constructor exists somewhere in
> another object

Example:

```js
class Foo {
  bar() {}
}

const foo = new Foo();

// we all know this is true
Object.getPrototypeOf(foo) == Foo.prototype)
// above is essentially the same as:
foo instanceof Foo
```

#### In Typescript

This works basically the same as `typeof`. Rarely used (used mainly with
classes, brr)

### User Defined Type Guard

Now we are talking. Check it out:

```typescript
interface Response {
  result: any;
  doSmth: any;
}

interface OkResponse extends Response {
  status: 'OK';
}
interface BadResponse extends Response {
  status: 'NOT_OK';
}

// now whenever we use this function typescript is going to set that variable as OkResponse
// if this function returns true, otherwise it will be BadResponse
function isGoodResponse(response: OkResponse | BadResponse) response is OkResponse {
  return response.status == 'OK'
}
```

### `in` Type Guards

You can also use `in` operator as a _boolean checks_ just like you sometimes
want to check if some browser feature is available.

```typescript
interface Athlete {
  speed: 99;
  age: 30;
}
interface NormalPerson {
  age: 30;
}
function isAthlete(subject: Athlete | NormalPerson): subject is Athlete {
  return "speed" in subject;
}
```

## Intersection Types

Instead of `extend`ing interfaces you can use `&` to _merge_ them.

```typescript
interface Order {
  amount: number;
}
interface Stripe {
  cvc: string;
  card: string;
}
interface PayPal {
  email: string;
}

// i think this is much better than interface Stripe extends Order {}
type OrderWithStripe = Order & Stripe;
type OrderWithPayPal = Order & PayPal;

// typescript is great at inferring as well!
const stripeOrder = Object.assign({}, order, stripeData); // OrderWithStripe
```

### Distributive properties

While working with unions you might want to check some condition for each member of the union.

For example, I might want to check if this union

```ts
type MyUnion = string | number | boolean;
```

contains the `string` type. Thankfully for us, the _union_ type has this property where when you use the _conditional types_ the condition will be applied to all members of the union.

```ts
type MyUnion = "ala" | 0 | true;

type HasString = Check<MyUnion, "ala">; // 'ala'
```

The condition inside the `HasString` type is expanded to something like this

```ts
type HasString =
  | ("ala" extends "ala" ? "ala" : never)
  | (0 extends "ala" ? "ala" : never)
  | (true extends "ala" ? "ala" : never);
```

Now, typescript is doing this for you, but this is something worth knowing.

### `Extract` utility

So now you know that union members are distributed among the condition when using _conditional types_, but did you know that you do not have to write the _condition_ itself? You can rely on the `Extract<U, T>`.

```ts
type MyUnion = string | number | boolean;

type HasString = Extract<MyUnion, string>;
```

Pretty neat right?

## Discriminant union

Ever used reducer? You probably used `action.type` or similar property to
distinguish between different actions.

To 'gather' all actions you probably did this:

```typescript
type Actions = ADD | DELETE | SOME_ACTION;
```

Thats the _union_ part. Now the _discriminant_ is the **thing that enables
typescript (and you) to distinguish between different actions**

```typescript
reducer(state, action) {
    // type is a common property
    // that lives on all of the actions
    switch(action.type) {
      case DELETE:
      // type inference works because of discriminant unions!
    }
  }
```

## Interface vs Type

- You cannot use `extend` keyword with types but you can use `&` instead

```typescript
interface Item {
  name: string;
}

interface Artist extends Item {
  songs: number;
}

type Artist2 = {
  songs: number;
} & Item;
```

- You can merge declarations with interfaces (you cannot have two types with the
  same name)

```typescript
interface Artist {
  name: string;
}
interface Artist {
  songs: number;
}
// /\ merged together

// now interface Artist contains name and songs
```

## Function Overloads

You can provide different implementations based on the arguments that we supply.
It makes stuff more readable

```typescript
// these are virtual, they will get compiled away
function reverse(dataToReverse: string): string;
function reverse<T>(dataToReverse: T[]): T[];
// real implementation
function reverse<T>(dataToReverse: string | T[]): string | T[] {
  if (typeof dataToReverse == "string") {
    return dataToReverse.split("").reverse().join("");
  }
  return dataToReverse.slice().reverse();
}
```

## Declare keyword

This keyword is used for telling typescript that a **javascript construct**
(like a function, variable etc) has already been defined elsewhere. (as a part
of runtime environment)

```ts
declare function add(x: number, y: number): number;

// somewhere in js file for example

function add(x, y) {
  return x + y;
}
```

This allows you to have JS codebase covered with types that are separate. Users
who use typescript can benefit from type completion while users using vanilla
still have access to your library.

## Readonly

The `readonly` keyword is used to make sure you are not mutating the data **within a given context**. If parameters are annotated with `readonly`, even if you pass something that previously was not annotated with `readonly`, it **will be _promoted_ to `readonly` within that context**.

```ts
type Person = {
  age: number;
  name: string;
};

const p: Person = {
  age: 12,
  name: "Wojtek",
};

declare function foo(person: Readonly<Person>): any;

foo(p); // Ok, note that I did not declare P as `Readonly<Person>`
```

Of course, if I declare the variable (cannot be primitive) as `readonly` there is no _promotion_ process.

```ts
type Person = {
  age: number;
  name: string;
};

const p: Readonly<Person> = {
  age: 12,
  name: "Wojtek",
};

declare function foo(person: Person): any;

foo(p); // Error, you cannot "downgrade" from readonly
```

## Enums

Enums are quite popular with _Ngrx_. They are not all sunshine and rainbows
though.

- they are typescript only concept
- can cause bundle bloat

```typescript
enum Something {}
// you just introduced this to your bundle
("use strict");
var Something;
(function (Something) {})(Something || (Something = {}));
```

Not looking to hot right? Well, there is a solution. A very simple one. Use
`const` before `enum`. That way the whole `enum` construct will get compiled
away and only picked properties will stay as normal variables. Enum props _get
inlined_

```typescript
const enum Something {
  yes = "Yes",
  no = "No",
}
let selected = Something.no;

// gets compiled to
("use strict");
let selected = "No"; /* no */
```

Much better now!

**BUT BEWARE** Enums cannot be used with _plugin-transform-typescript_ which you
are probably using.

## Mocking with Typescript

When testing sometimes you have to mock stuff. It's pretty common procedure, but
typescript sometimes makes it difficult.

```typescript
import { Link as MockLink } from 'react-router-dom';

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  Link: ({ children }: { children: React.ReactNode }) => children
}));

test(/* some test name */, () => {
  // TypeError: !
  MockLink.mockImplementationOnce(() => {/* ... */})
})
```

It is frustrating , we have to help typescript a little bit by casting to a
`mock type`

```typescript
import { Link as LinkDep } from "react-router-dom";

const MockLink = LinkDep as jest.Mock<LinkDep>;

// now you can test in peace
```

Same technique applies to _global mocks_.

## Typing `get` function

Typing such functions is a nightmare. But we can make our lives easier with a
couple of tricks.

[### All credit goes to this article](https://medium.com/@jamesscottmcnamara/type-yoga-typing-flexible-functions-with-typescripts-advanced-features-b5a282878b74)

And btw, we are going all in when it comes to functional programming so our
`get` function will be fp ready :)

### HasKey

Turns out you can create object types out of thin air. Check this out:

```typescript
type HasKey<K extends string, V> = { [_ in K]: V };
type Testing = HasKey<"wojtek", number>;
/*
  {
    wojtek: number
  }
*/
```

This is just mind bending stuff. Very clever usage of the `in` keyword. How does
it work?

- We are extending `string` because generic will be type literal (like `wojtek`
  or `ala`)

- `{[_ in K]: V}` means an object with keys in K with value V

Lets say you use `|` then typing `K` what will happen? Well you will get 2 props
on an object with value: `any` (or any value you passed to generic).

Again, very clever stuff

### Basic Implementation

With `HasKey` we can start our basic implementation.

```typescript
declare function get<K extends string>(
  key: K
): <Obj extends HasKey<K>>(obj: Obj) => Obj[K];
```

Our function could be used as such

```ts
get("name")({ name: "wojtek" }); // all ok
get("name")({ someprop: "someprop" }); // Typescript is not happy, error!
```

### KeyAt

You might think that we achieved what we wanted:

> You just have to declare more overloads right?

Not really, sadly this function is far from complete. The inferring system might
have problems with more complex types.

To fix this we introduce another type: `KeyAt`

```typescript
type KeyAt<Obj, K extends string> = Obj extends HasKey<K> ? Obj[K] : never;
interface SomeInterface {
  wojtek: 'ala 123'
}
KeyAt<SomeInterface, 'wojtek'> // 'ala 123', literal type!
```

This allows us to _pluck a given type_ out of object. This makes sure that
return our function has return value correctly typed.

```ts
declare function get<K extends string>(
  key: K
): <Obj extends HasKey<K>>(obj: Obj) => KeyAt<Obj, K>;
```

Personally i would name this type `TypeAt` but I'm going to roll with this name
paying an homage to original author :).

### Traversals

Our function can also filter stuff. We basically want to work _lenses-like_.

Example:

```ts
get(
  matching((friend) => friend.friends > 5),
  "name"
)(obj.friends);
```

With our current implementation this operation is impossible. How would we
enable such functionality?

Let's type `matching` first:

```ts
interface Traversal<Item> {}

declare function matching<A>(
  filteringFunction: (a: A) => boolean
): Traversal<A>;
```

We have to change our implementation a bit to introduce `matching`.

```ts
declare function get<Item, K extends string>(
  traversal: Traversal<Item>,
  key: K
): <Obj extends Array<HasKey<K>>>(obj: Obj) => Array<KeyAt<Obj, K>>;

const popularFriends = get(
  matching((user: User) => user.friends.length > 5),
  "name"
)(user.friends);
```

But there is a problem. Our `popularFriends` are typed as `never[]`.

Going back to our declaration of `KeyAt` we typed it so that `Obj` has to be
`HasKey<>` not `Array` of that type.

That is easily fixable. Just change `obj: Obj` to `obj: Obj[]`.

### Unpacking

Very useful stuff for our function (which we want to be able to accept multiple
containers) and overall (I really wonder why they would not put it inside TS
utility types already).

```ts
// power of conditionals and infer baby!
export type Unpack<F> = F extends Array<infer Item>
  ? Item
  : F extends Set<infer Item>
  ? Item
  : F extends Map<any, Item>
  ? Item
  : F extends { [n: string]: infer Item }
  ? Item
  : F extends { [n: number]: infer Item }
  ? Item
  : never;
```

## Tuples and Currying

This is going to be wild ride so strap on.

### Head

This type will let us pluck off the head of the `tuple Type`. Will come in handy
later

```ts
type Head<A extends any[]> = A extends [infer First, ...any[]] ? First : never;
type Test = Head<[1, 2, 3, 4]>; // 1
```

This type is using `infer` to get the correct type.

### Tail

We implemented `Head` it's time for `Tail` now. As of writing this we cannot
just get the last type out of the tuple.

Lets try the naive approach

```ts
type Tail<A extends any[]> = A extends [any, ...infer tail] ...
```

### HasTail

Since classical curried functions are taking one argument at a time we have to
know when we should stop and return the return type. `HasTail` type will allows
us to do so

```ts
type HasTail<A extends any[]> = A extends [] | [any] ? false : true;
type Test = HasTail<[]>; // false
type Test2 = HasTail<[1, 2]>; // true
```

Pretty straight forward right? Unless our tuple is empty or only has 1 element
left we can keep going with currying.

### CurryV0

With those simple types we can define our **strict curry** type

```ts
type CurryV0<Parameters extends any[], ReturnType> = (
  arg: Head<Parameters>
) => HasTail<Parameters> extends true
  ? CurryV0<Tail<Parameters>, ReturnType>
  : ReturnType;

declare function curry<P extends any[], R>(f: (...args: P) => R): CurryV0<P, R>;

function addTwo(x: number, y: number) {
  return x + y;
}

const curried = curry(addTwo);

curried(1)(2); // works like a charm!
```

I specifically am very verbose with names to make this type clear. This type is
using recursion to gradually (with each call) pluck off one parameter at a time.

### Last

Our curry implementation is great! But, we can always improve on things right?.
So what if we want to support _loose curry_ ? (like partial application). This
would prove to be very difficult using our current tools.

One type that might help us reach that goal of partial application is the `Last`
type.

Instead of plucking off tail from tuple we will only pluck the last type.

```ts
type Last<P extends any[]> = {
  0: Last<Tail<P>>;
  1: Head<P>;
}[HasTail<P> extends true ? 1 : 0];
type Test = Last<[1, 2, 3, 4]>; // 4
```

So this might be hard to digest but stay with me. This type is using recursion

- if there is a tail, pass that tail recursively to `Last`
- if there is no tail use that value, stop recursion

The picking if we need to stop the recursion happens in `[]`. This is what's
called `indexed type accessor`.

You might think we can do the type using normal turnery like :

```ts
type Last2<P extends any[]> = HasTail<P> ? Last2<Tail<P>> : Head<P>
```

This restriction stems from TS itself, you can though reference a type from
within an object type just like we are doing with our first `Last`
implementation.

#### 1000 IQ big brain Last

Why do you bother with `Head` and `Tail`. You probably forgot about the `length` property on a array.
In _JavaScript_ we would usually subtract one from the length to get the last element right?

```js
const last = arr[arr.length - 1];
```

You could also use the `slice` method (I still think the `length` method is much more readable)

```js
const last = arr.slice(-1)[0];
```

Either way, you subtract one from the length. So how we can do that in `typescript`. Well glad you asked!

```ts
type Last<T extends any[]> = [any, ...T][T["length"]];
```

Just appreciate how nice it is and how much less _noise_ there is. Since we cannot _subtract_ per se in `typescript`, we can add one throwaway type and just call `length`. It is that simple!.

### Length

This type will allow us to have basic information about arguments that are
passed in and such. This, in terms, will allows us to implement partial
application.

```ts
type Length<T extends any[]> = T["length"];

type Test = Length<[1, 2, 3, 4]>; // 4
```

`Length` type will work as a pseudo-counter.

### Prepend

This will allow us to prepend a type to a tuple type, which will allow us to
know which parameters has already been supplied. To implement this type we will,
again, make us of `function types` trick.

```ts
type Prepend<TypeToPrepend, Tuple extends any[]> = ((
  head: TypeToPrepend,
  ...tail: Tuple
) => any) extends (...args: infer U) => any
  ? U
  : never;

type Test = Prepend<number, [1, 2, 3]>; // [number, 1,2,3]
```

Just to make sure you know how this is working. As mentioned before, we cannot
use array destructuring (or spread for that matter) to assign types or use
`infer` keyword.

To circumvent this restrictions we have to operate on functions parameters.

In this type we are basically checking if `Function == Function` (which will
always be the case) but in the act of checking we can merge `TypeToPrepend` and
`TupleType` using `infer U` on the second function.

### Drop

Just like we can `Prepend` type to a tuple type we are also in need of the
ability to remove a number of arguments from the tuple type.

To achieve such functionality we will use **recursive indexed types** (already
seen before)

```ts
type Drop<
  ElementsToDrop extends Number,
  TupleToDropFrom extends any[],
  Iterator extends any[] = []
> = {
  0: Drop<ElementsToDrop, Tail<TupleToDropFrom>, Prepend<Iterator, any>>;
  1: TupleToDropFrom;
}[Length<Iterator> extends ElementsToDrop ? 1 : 0];

type Test = Drop<2, [1, 2, 3, 4]>; // [3,4]
```

One might be curious about that `any` type passed to `Prepend`. Well this
`Iterator` type only accts as to-be-thrown-away accumulator, one might say:
recursion stop predicate. Since we do not care about the type passed to iterator
we default to `any`.

## Matching exact shape

> TypeScript is a structural type system. This means as long as your data
> structure satisfies a contract, TypeScript will allow it. Even if you have too
> many keys declared.

This definitely can be a problem. Example

```typescript
type Person = {
  first: string;
  last: string;
};

const tooFew = { first: "Stefan" };
const tooMany = { first: "Stefan", last: "Joe", other: "something" };

declare function savePerson(person: Person): void;

savePerson(tooFew); // Error!
savePerson(tooMany); // OK -> WTF ;C
```

We can use clever generic type along with `Exclude` to make sure our params are
the exact shape of given `type / interface`

```ts
type ValidateShape<T, Shape> = T extends Shape
  ? Exclude<keyof T, keyof Shape> extends never
    ? T
    : never
  : never;
```

- check if T is _`Shape`-compatible_ with `extends`
- see if there is 1:1 match between `T` and `Shape` when it comes to properties.
- otherwise return never

Pretty neat stuff!. With this we can re-write our example to:

```ts
declare function savePerson<T>(person: ValidateShape<T, Person>): void;

savePerson(tooFew); // Error -> returns never
savePerson(tooMany); // Error -> returns never
```

## Optional Chaining && Null Coalescing

As of writing this, these count as _future_.

### Optional Chaining

If you ever worked with Angular you are probably familiar with optional
chaining. It is very interesting that Angular had that from the version 2
onwards and Typescript is getting them now 🤔.

So an example:

```js
const someObj = {
    prop1: {
        prop2: undefined,
    },
};

const value = someObj.prop1 && someObj.prop1.prop2
        // ...
        & const; // ..
```

Syntax with `?` is much cleaner, especially with nested objects and properties.
You no longer have to worry about checks with `&&`. `?` operator takes care that
for you.

## `const` assertion

There are _literal_ and _primitive_ types in typescript. You know, there is a difference between these 2 types:

```ts
type Literal = "string";

type Primitive = string;
```

Usually, we do not want to _mutate_ anything, especially while working with objects and so on. Sine objects and arrays are mutable, typescript was wideing the type of the property to accommodate for that (also for `let` declarations)

```ts
const circle = {
  radius: 10,
};
// radius: number :C
```

Now, we can tell typescript that we will not be mutating that value, using so called _const assertion_

```ts
const circle = {
  radius: 10,
} as const;
// radius: 10 => literal type
```

## Functional types

### Function parameters

Typescript can now infer the _function parameters_ from type variable.

So previously to type this simple signature:

```ts
function call(fn, ...args) {}
```

You either went for weak `any[]` for parameters with generic as return type:

```ts
declare function call<R>(fn: (...args: any[]) => R, ...args: any[]): R;
```

Or played around with `infer` keyword

```ts
type GetFunctionArguments<F> = F extends (...args: infer A) => any ? A : never;
declare function call<R, F extends (...args: any[]) => any>(
  fn: F,
  ...args: GetFunctionArguments<F>
): R;
```

But now since there were improvements to _touple types_ you can just specify the parameters as a generic type.

```ts
declare function call<A extends any[], R>(fn: (...args: A) => R, ...args: A): R;
```

Much better!

## Typescript being too generous

I love typescript, but sometimes, in my opinion, it's just too permissive, even on the strictest settings.

### The `VoidFunction`

`VoidFunction` type, function that should not return anything. It turns out, you often can return stuff from that function, _Typescript_ will not shout at you (but it should!)

```ts
const doWork = (work: VoidFunction) => work();

const arr = [];
doWork(() => arr.push(1));
```

The `arr.push` will return the amount of elements that were pushed into the array. So in this case, the `work` function annotated as `VoidFunction` actually returns some data.

This behavior is pretty weird. I would love it so that _Typescript_ will not let me do that.

### The matching parameter shape

As long as the data you are passing matches the parameter type, it will be allowed. I'm mainly talking about this

```ts
type Props = {
  key: string;
};

const doWork = (props: Props) => {};

const param = { key: "1", somethingElse: "2" };
doWork(param);
```

It would be super nice for _Typescript_ to complain here.

## Testing for `never`

Let's say that for some reason, you want to test if the _type parameter_ that you defined is of type `never`. Normally I would do something like this

```ts
type CheckForNever<T> = T extends never ? 1 : 0;

type Test1 = CheckForNever<"hi">; // 0 => Ok.
type Test2 = CheckForNever<never>; // never => Wtf?
```

Tbh, I have no idea why it happens. The way the article describes is

> Once one of the types in your expression is `never` it will **poison** the rest of the expression to evaluate to `never`

You can trick the type system into behaving as you first though (returning `1` when `never` is passed in) by checking on _touple types_

```ts
type CheckForNever<T> = [T] extends [never] ? 1 : 0;

type Test1 = CheckForNever<"hi">; // 0 => Ok.
type Test2 = CheckForNever<never>; // 1 => Ok.
```
