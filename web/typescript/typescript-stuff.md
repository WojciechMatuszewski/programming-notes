# Typescript Stuff

## `tsconfig.json`

### Multiple `tsconfig.json` files and VS Code

As of the time of writing this, **VS Code DOES NOT support custom `tsconfig` file names**. If you wish to have a configuration with multiple `tsconfig.json` files which utilize the _project references_, you **need** to have **at least one `tsconfig.json` somewhere in your project** for the VS Code to pick your configuration up.

**Please note that, even with _project references_ VSCode might not show up TS Errors in the IDE**. Some errors might only show up after you run the `tsc`. This is supper annoying and **works as expected in WebStorm**.

#### Scoping given `compilerOptions` to a set of files

Let us imagine you want to add the [`noUncheckedIndexedAccess`](https://www.typescriptlang.org/tsconfig#noUncheckedIndexedAccess) option to your TypeScript config.

One might want to add it **only to application files** and **skip the test files**. The application files where a runtime bug could cause the website to crash. We do not care about runtime issues within the test files – a failing test is not a user-facing issue.

**First**, create **a separate `tsconfig` file**. In my case, I will name mine `tsconfig.jest.json`. The "test only" `tsconfig` file would extend the base `tsconfig`. Inside that "test only" file, I'm going **to disable the `noUncheckedIndexedAccess` setting**.

```
{
 "extends" ...,
"compilerOptions": {
  "noUncheckedIndexedAccess": false,
....
},
include: [ALL_TEST_FILES, ALL_APPLICATION_FILES]
}
```

Next, **add the `noUncheckedIndexedAccess` to the root `tsconfig.json` and EXCLUDE all test files**. This way, TypeScript will not scream at you whenever you do not explicitly perform null-checks when accessing array items in your test files.

**The good rule to follow is this**: every layer in your application should have a separate `tsconfig.json` file. The tests, the e2e tests, the application code etc... Otherwise, you will "leak" global types, or types that should only be available in the test files to other parts of your application.

#### The problem

Notice that, in the previous example, we "downgraded" the strictness of the type checker for a subset of files. **The reverse would NOT be possible**.

Since the `tsconfig.jest.json` must also include the application files (you are testing the application files after all), specifying `noUncheckedIndexedAccess` for the "test only" config would automatically apply it to the application files. This would likely result in many TypeScript errors forcing you to adopt the `noUncheckedIndexedAccess` in the whole codebase.

The inverse (the example above) is possible because the application files have "stricter" type settings than the test files.

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

### The `moduleDetection` setting

Go and do this in your project now.

1. Create an empty `.ts` file. The name does not matter.

2. Add the following code to that file:

```ts
const name = "FooBar";
```

Are you staring at a TypeScript error? I was just as baffled when you are now.

> Cannot redeclare block-scoped variable 'name'.ts(2451)
> lib.dom.d.ts(27241, 15): 'name' was also declared here.

What is going on??

It turns out, **TypeScript, by default, assumes that each file lives in "global" namespace**.

If you do not export anything from the file, or set the `moduleDetection` to a specific value, it thinks that all the files in your project will merge with the "global" environment. It just happens that the `name` variable is already defined as a global variable when using `lib: "dom"`.

**We have three potential solutions to the problem**.

1. Add an empty export to the file, like so:

   ```ts
   const name = "FooBar";
   export {};
   ```

2. Change the `moduleDetection` to `force`.

3. Add the `type:"module"` to your `package.json`.

**You most likely never seen this error popup in your codebase**. Why? Because when creating a file, you most likely export from it or you have the `type:"module"` set in your `package.json` file – otherwise your bundler wont work.

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

## Namespaces

_Namespaces_ are used to **organize the global interfaces / types**. It is to ensure that **multiple libraries do not override each other types**.
The most notable namespace is the `NodeJS` namespace. Please note that the **namespaces, like interfaces can merge together**. It all depends on how you structure the files.

**_Namespaces_ are considered legacy feature – something to be used before ESM modules came along**. They also **compile to runtime code** which is far less than ideal.

## Augmenting global declarations

In the section below, we will be using the `declare` keyword. Before diving into each problem, it is vital you understand what the `declare` keyword is about.

The **`declare` keyword tells TypeScript that a variable/function/module exists and might be used inside the code**. Think of it as declaring global variables with the benefit of autocompletion. Of course that is not the sole purpose of this keyword. Read more below!

### Augmenting NodeJs `process.env`

Lets say you are building a NodeJs app and you want to have strongly typed `process.env` object.

All you have to do is to create some `.d.ts` file (could be `.ts` file but I would go for `d.ts` for clarity) and use the fact that **namespaces are merged just like interfaces**.

```ts
declare global {
  namespace NodeJS {
    interface ProcessEnv {
      MY_GLOBAL_ENV_VARIABLE: string;
    }
  }
}
```

That is all.

### Augmenting the `window`

Some libraries add things into the `window` object. For example, the `Cypress` library or the `monaco-editor` library.
In most cases, these should be ignored, but in some, like debugging, they are very handy.

If you want to access them without ignoring TypeScript errors, you need to type them – augment the `window` property.
Here is how to do it.

```ts
declare global {
  // Uses interface merging
  interface Window {
    Cypress: MY_TYPE;
    Monaco: MY_TYPE;
  }
}
```

#### With dynamic types

Sometimes, the list of the types you want to add is dynamic – it might be coming from an object that will change over time. Imagine the following scenario.

```ts
const addAllOfThisToWindow = {
  add: (a: number, b: number) => a + b,
  subtract: (a: number, b: number) => a - b,
  multiply: (a: number, b: number) => a * b,
  divide: (a: number, b: number) => a / b,
};

Object.assign(window, addAllOfThisToWindow);
```

We could add those properties one by one, and that would be my preferred way (depending on the team experience with TypeScript).
We can also leverage the fact that the `window` is an interface which we can extend.

```ts
const addAllOfThisToWindow = {
  add: (a: number, b: number) => a + b,
  subtract: (a: number, b: number) => a - b,
  multiply: (a: number, b: number) => a * b,
  divide: (a: number, b: number) => a / b,
};

declare global {
  // Keep in mind that for the `extends` to work, you have to declare either a separate type or an interface.
  // You cannot do it "inline".
  type ToAdd = typeof addAllOfThisToWindow;
  interface Window extends ToAdd {}
}
```

That is it!

### Augmenting NodeJs `global`

```ts
declare global {
  var myGlobalVariable: string;
  function myGlobalFunction(): string;
}
```

Keep in mind that this is not the same as augmenting the `process.env` or the `window`. This code augments the `globalThis`.

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

### Colocation of global interfaces

I'm not sold on globally-accessible interfaces, but if you were to do this pattern, I would **recommend using colocation when augmenting the `global`**.

Here is what I mean

```ts
// file foo.ts

declare global {
  interface MyGlobalInterface {
    // PROPERTIES THAT ONLY THIS FILE USES
  }
}

// file bar.ts
declare global {
  interface MyGlobalInterface {
    // PROPERTIES THAT ONLY THIS FILE USES
  }
}
```

If you were to delete the `bar.ts` file, all the properties appended to the `MyGlobalInterface` would no longer be available. This is pretty nice, as the worst part of refactoring code is often checking if something is used or not, specially if the interface is quite complex.

### Typescript ignores my `d.ts` file

First of all check if that file matches the `include` pattern that you specified within your `tsconfig`.

If that's the case, we are dealing with something very strange that I've discovered only recently.

You **have a file named the same way as you `d.ts` file**, eg. `env.ts` and `env.d.ts` file.
The way typescript works is that **the `env.d.ts` file will be ignored since typescript things it was derived from `env.ts` file**. Pretty strange right?

<https://github.com/microsoft/TypeScript/issues/31397#issuecomment-492269754>

There are 2 solutions here:

1. Rename your `d.ts` file
2. Specify the `d.ts` file within the `file` block inside your `tsconfig`.

#### When using d.ts file I accidentally overwritten all definitions for a library

Depending on how your tsconfig is set up, by writing a `declare module 'XX'` in the `d.ts` file you will be **completely overwriting the types of a given library**. This option is a bit nuclear, but sometimes it is completely valid, for example when the library is written in JS and does not export any types.

### I cannot override method inside an interface

**To my best knowledge**, overriding a method declared inside an interface is **not possible**. Why?

Because TypeScript **will pick the "latest" evaluated interface** for that method signature.

```ts
interface Foo {
  bar(): string;
}

interface Foo {
  bar(): number;
}

declare const foo: Foo;

const result = foo.bar();

// result: number
```

Now, if I were to switch the interfaces.

```ts
interface Foo {
  bar(): number;
}

interface Foo {
  bar(): string;
}

declare const foo: Foo;

const result = foo.bar();

// result: string
```

Kind of makes sense does not it?

It can be frustrating though – in situations where you want to override a 3rd party library interface property. **I do not have a good answer for doing that yet**.

### What does "ambient context" mean?

When working with global types, you might encounter a world _ambient_. For example, TypeScript is going to complain when you try to use the `declare` keyword within a `declare` block, like so:

```ts
declare global {
  declare function myFunc(): boolean; // ERROR: A 'declare' modifier cannot be used in an already ambient context.
}
```

So what does the _ambient context_ mean here?

**You can think of the _ambient context_ as ["without implementation"](https://github.com/Microsoft/TypeScript-Handbook/issues/180#issuecomment-195452691) or that the implementation of the code is somewhere else**.

In the example above, the implementation of the `global` does not live within the file I'm editing. The types will be erased at runtime. Thus the _declaration_ I'm working with exists within an _ambient context_.

## "Loose" autocomplete

> Check out [this tip](https://www.totaltypescript.com/tips/create-autocomplete-helper-which-allows-for-arbitrary-values).

Sometimes you might want to have a property which can take either a well defined value or any value of a given type. In such situations, ideally, we want to keep the autocomplete functionality we have, when the type is scoped to only a couple of well-defined values.

```js
function getFontSize(size: "sm" | "xs") {
  // Here I would like to also allow for `any` size where I fallback to some value
}
```

One way to do that, would be to expand the `size` prop to accept the `string` type.

```js
function getFontSize(size: "sm" | "xs" | string) {}
```

The **problem with this approach is that we are going to loose autocomplete**. TypeScript will **expand the `size` to accept strings**. There is no way to provide autocomplete on "all" strings.

**One solution is to scope the `string` type**, by omitting the well-defined values.

```ts
function getFontSize(size: "sm" | "xs" | Omit<string, "sm" | "xs">) {}
```

Now, the autocomplete works as expected. It's either `sm`, `xs` or all the strings (except the `sm` or `xs` value).

**Another solution is to use the `{}` type**.

```ts
function getFontSize(size: "sm" | "xs" | (string & {})) {}
```

Check out [this StackOverflow reply for more information](https://stackoverflow.com/a/61048124).

## What the `d.ts` file is good for and what it is NOT good for

The `d.ts` files are for

- Providing global types that are available in all your TS files without the need to import them.

- Augmenting the library types.

The `d.ts` files are **not for**

- **Putting types that are related to modules and importing those types in different files**.

The reason why you do not want to do this, is that, in most cases, you will have the `skipLibCheck:true` turned on (makes sense, since it is TypeScript that usually generates those files. Why bother checking them?). If you do, you will not notice the type errors in those files.

**TypeScript team also is against putting module-types into `d.ts` files. See [this comment](https://github.com/microsoft/TypeScript/issues/52593#issuecomment-1419505081)**.

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

### TypeScript 4.5 improvements

The `import type` syntax is very helpful but it forces you to create "type-only" imports (which I'm a fan of).
The TypeScript team decided to add "squashed type-only" import. The syntax is as follows

```ts
import { someFunc, type someFuncType } from "./a";
```

The functionality and the intent stays the same, but instead of two import statements you do one.

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

## The `register` function

The `ts-node` and similar tools expose a very neat mechanism that parses TypeScript files on the fly, at the very moment you `require` them!
All you have to do is to use the `.register` function. Here is an example:

```ts
const tsNode = require("ts-node");

tsNode.register({
  transpileOnly: true,
});
```

### Where would I use this?

Imagine you want to load a file that is provided by a 3rd party, then execute that file. Of course, keeping all operations within the realm of JavaScript.

Well, since you cannot execute TypeScript files, you will not be able to just import them into your script. You can execute JavaScript files, do not you?

With the help of the `register` call, all of your imports will be transpiled by `ts-node`. This ensures that whatever you are importing
is a valid JavaScript, without any types and other TypeScript features.

## Assert Signatures

So with `type guards` you are returning `true` or `false`. This then determines
the outcome of the type. But `assert signatures` **are quite different**.

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

The signature differs greatly and there is actually more to the `assert signatures` signature than presented here.
On top of that, the **assert-type functions cannot be using the arrow function syntax**.

```ts
// THIS WILL NOT WORK!
const assertValue = (value: boolean): asserts value => {
  // ...
};

// Error when using it: Assertions require every name in the call target to be declared with an explicit type annotation
```

I'm not sure why is this the case. It seems to have something to do with _control flow analysis_, but I'm not 100% sure.
See [this link](https://github.com/microsoft/TypeScript/issues/34523) for more info.

### Two types of Assert Signatures

There are actually 2 variants

- for checking a condition
- for telling TypeScript that specific variable or property has a different type

So it all basically boils down to that `Assert Signatures` does not return anything, they throw this `AssertionError` whenever something is wrong.

`Type Guards` on the other hand return `true` or `false` based on they inputs.

The signature `asserts something` or `asserts x is something` tells the reader of the code that **that function will only return if the assertion holds**

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

## Getting the type out of the array

Let us say you have are working with the following array.

```ts
type PossibleCombinations = ["foo", "bar"] | ["baz", "quix"];
```

How would you go about getting all the values as type union from the `PossibleCombinations` type? You could use the index operator, like so.

```ts
type AllValues = PossibleCombinations[0 | 1];
```

That would work, but if the length of the arrays are different, you might be in trouble.

```ts
type PossibleCombinations = ["foo", "bar"] | ["baz", "quix"] | ["a", "b", "c"];
type AllValues = PossibleCombinations[0 | 1 | 2]; // the union contains `undefined` type
```

Since the size of the arrays varies, we cannot use hardcoded indexes. For this, **use the `number` keyword**.

```ts
type PossibleCombinations = ["foo", "bar"] | ["baz", "quix"] | ["a", "b", "c"];
type AllValues = PossibleCombinations[number]; // union of all the values
```

The `number` keyword here is the **union of all the possible numbers**. You can **think of the `number` keyword as the "numeric index" of a given type**.

So, the `AllValues` is a _type at a numeric index of the `PossibleCombinations` type_.

### Another example

The example above was a bit contrived. Most of the time we do not have to deal with union of arrays with unknown lengths.
In most cases, the type definition you have to deal with would look like similar to the following.

```ts
type People = { name: string; age: number }[];
```

How can we extract the _inner_ type from the `People` array type? – there are at least three ways I'm aware of.

```ts
type Person_v1 = People[0]; // Works because all the items in the array are the same.
type Person_v2 = People[number];
type Person_v3 = People extends Array<infer Inner> ? Inner : never;
```

I'm leaning towards the first two options. The last one is a bit of an overkill if you ask me.

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

#### Exclude an non-primitive types

You can use the `Exclude` utility type with interfaces and types. Check this out.

```ts
type Event =
  | {
      type: "click";
      x: number;
      y: number;
    }
  | {
      type: "focus";
    }
  | {
      type: "change";
      value: string;
    };

type ClickAndFocusEvent = Exclude<Event, { type: "click" }>; // { type: 'focus' } | { type: 'change', value: string }
type OnlyChangeEvent = Exclude<Event, { type: "click" | "focus" }>; // { type: 'change', value: string }
```

In the above example, we were dealing with a discriminant on the `type` property. **But the `Exclude` also works without a discriminant**.

```ts
type Evnt =
  | {
      x: number;
      y: number;
    }
  | {
      name: string;
    }
  | {
      type: "change";
      value: string;
    };

type ClickAndFocusEvent = Exclude<Evnt, { name?: string; value?: string }>; // { x:number, y:number }
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

type MyOwnReturnType<ReturnType, F> = F extends Callable<ReturnType> ? ReturnType : never;

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
type GetFunctionArgumentTypes<F> = F extends (...args: Array<infer U>) => void ? U : never;

function numberArg(x: number) {}

function arrayMixed(x: [1, "a", {}]) {}

type t1 = GetFunctionArgumentTypes<typeof numberArg>; // number
type t2 = GetFunctionArgumentTypes<typeof arrayMixed>; // [1, 'a', {}]
```

#### Inferring multiple types

Nothing is stopping your from using the _infer_ keyword multiple times. Check this out

```ts
type AppendArgument<Fn, A> = Fn extends (...args: infer Args) => infer R ? (...args: [...Args, A]) => R : never;
```

Here I've used _infer_ to both get the hold of the function arguments, but also the return type of the `Fn` type. While I could use the `ReturnType` generic, using _infer_ is also nice ;)

##### Placing a constraint on the inferred type

Did you know you can **place a constraint on what you infer**? The syntax looks as follows.

```ts
type FakeReturnTypeString<F> = F extends ((...args: any[]) => infer R extends string) ? R : never;

function returnNumber() {
  return 3;
}

type test = FakeReturnTypeString<typeof returnNumber>; // never
```

It is as if you were to say: **infer this type ONLY WHEN it matches my constraint**. Seems pretty useful, **especially if the alternative is to perform a nested ternary check**.

```ts
type FakeReturnTypeString<F> = F extends (...args: any[]) => infer R ? (R extends string ? R : never) : never;
```

The version with a constraint on the `infer` seems much cleaner.

### Conditional return type

There comes a time where you might need to put a conditional type as a return type for a given function. When doing so **you will most likely encounter an error**. The error would be saying that TypeScript cannot assign the return type to the values you are returning from the function. **This is to be expected**, at least at the time of writing this – TypeScript is not smart enough to figure everything out!

```ts
function youSayGoodbyeISayHello<
  TGreeting extends string,
  TReturn = TGreeting extends "hello" ? "goodbye" : "hello"
>(greeting: TGreeting): TReturn {
  return greeting === "goodbye" ? "hello" : "goodbye"; // error
```

The **solution is to perform as type cast on the return value**. I hope that, one day, this will not be necessary, but at the time of writing this, it is.

```ts
function youSayGoodbyeISayHello<TGreeting extends string, TReturn = TGreeting extends "hello" ? "goodbye" : "hello">(
  greeting: TGreeting,
) {
  return (greeting === "goodbye" ? "hello" : "goodbye") as TReturn; // Notice the cast here.
}
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

### Altering mapped types keys

You can use `as _something_` syntax to alter (probably perform conditional operation) on the _mapped types_ keys.

```ts
type NoNumbers<T extends Record<string, unknown>> = {
  [K in keyof T as T[K] extends number ? never : K]: T[K];
};

type Test = NoNumbers<{ prop1: string; prop2: number }>; // {prop1: string}
```

As of writing this, this is relatively new addition to the language.

Keep in mind that **you can also alter what we "iterate" over**. This is **very handy when dealing with an union**.

```ts
type MyUnion =
  | {
      type: "banana";
      color: "yellow";
    }
  | { type: "apple"; color: "red" };

type ToObject = { [K in MyUnion as K["type"]]: K }; // banana: {type: "banana", color: "yellow"} and so on...
```

If you try to "iterate" over the union without the `as` keyword, you will get an error. It is super handy since `K` is the whole object (member of the union), and not a primitive value.

### Reverse Mapped Types

Usually when using mapped types, you work on a result. The pipeline is `A -> MappedType<A> -> B`. This means that there are no "backwards lookups" of types because the `A` is already established.

The **idea behind the reverse mapped types is that we are instructing TypeScript to "look back" and transform a given type**.
I think its best illustrated on a example.

```ts
export function makeEventHandlers<T extends Record<string, unknown> = {}>(obj: {
  [K in keyof T]: (prop: K) => void;
}) {
  return obj;
}

const obj = makeEventHandlers({
  click: (name) => {
    console.log(name); // name is `click`

    type test = Expect<Equal<typeof name, "click">>;
  },
  focus: (name) => {
    console.log(name); // name is `focus`

    type test = Expect<Equal<typeof name, "focus">>;
  },
});
```

Notice that, in the example, I'm **not providing the shape of the obj upfront, yet TypeScript is still able to compute the property names and pass them to the function argument**. Again, what I pass to the `makeEventHandlers` function is not known upfront. **I'm able to transform the inferred type**. This is pretty cool!

Inspired by [this talk](https://youtu.be/fjPzlywUH0M?t=3020).

### Keyof and `[keyof]`

Differences are quite big

```typescript
type Test1 = keyof Something; // "id" | "name" | "property"
type Test2 = Something[keyof Something]; // string | number | undefined
```

It's very similar to accessing object values and `Object.keys` in JS. **It's just that the value is the type itself**

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

## Type Guards (predicates)

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

> it tests if a `.prototype` property of a constructor exists somewhere in another object

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

This works basically the same as `typeof`. Rarely used (used mainly with classes, brr)

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

### Type Guards in callback functions

This technique is **very useful for `filter` or `reduce` callbacks where you want the return type to match your filter predicate without using `as` keyword**. Here is an example.

```ts
const values = [1, 2, undefined, 3];

const theRegularFilter = values.filter((val) => val != null); // (number | undefined)[]
const theTypeCastFilter = values.filter((val) => val != null) as number[]; // number[]

const theTypeGuardFilter = values.filter((val): val is number => {
  return val != null;
}); // number[]
```

The `theTypeGuardFilter` is quite nice. It works on the similar basis as the `as` keyword **in terms of type-safety**, but is is much robust solution.
This is a great way to ensure the filtered data has the type you want it to have!

### Type Guards with generics

Depending on how you structure the generic functions, you **can use the type guards with generic signatures**.

```ts
type Confirmer<T> = (node: unknown) => node is T;
declare function confirm<T>(confimer: Confirmer<T>);
```

This also works for properties within the interface.

```ts
interface DOMNodeExtractorConfig<T, Result> {
  isNode(node: unknown): node is T;
  // As alternative one might write
  // isNode: (node: unknown) => node is T;
  transform: (node: T) => Result;
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

Let us say you have this collection of fruits.

```ts
type Fruits = "apple" | "banana" | "orange";
```

Now, you would like to create a type that only contains the `"apple"` or `"banana"`. To do so, you might be tempted to write the type like this.

```ts
type AppleOrBanana = Fruit extends "apple" | "banana" ? "apple" | "banana" : never;
```

Sadly this **will not work as you would expect**. The underlying type is still `never`. This is because **the _distributive_ part of the types is only applied when we are dealing with generics**.

In the example above, TypeScript will **compare the WHOLE `Fruits` type to the WHOLE union type you created by hand**. This is not want you want. You want to compare each member of `Fruits` with each member of the union type. To do so, you must use generics.

```ts
type AppleOrBanana<TFruit> = TFruit extends "apple" | "banana" ? "apple" | "banana" : never;
```

This time, TypeScript will compare each member of `TFruit` with each member of the `"apple"| "banana"` union.

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

- It seems like **there is no performance difference between the `types` and `interface` keyword**. [Source](https://youtu.be/zM9UPcIyyhQ?t=58).

  > For the most part, you can choose based on personal preference, and TypeScript will tell you if it needs something to be the other kind of declaration. If you would like a heuristic, use interface until you need to use features from type.

- Check [the official documentation](https://www.typescriptlang.org/docs/handbook/2/everyday-types.html#differences-between-type-aliases-and-interfaces).

All in all it boils down to what you prefer. Do you prefer the name `interface`? Go with the `interface`. Do you prefer the name`type`? Go with `type`. **The only exception is when building a library. There you should be exposing interfaces so that users can extend your types!**.

## Function Overloads

You can provide different implementations based on the arguments that we supply. It makes stuff more readable. Remember that **to use overloads you have to use the `function` keyword**.

```typescript
// these are virtual, they will get compiled away
function reverse(dataToReverse: string): string;
function reverse<T>(dataToReverse: T[]): T[];
// real implementation. You should add an explicit return type here for safety.
function reverse<T>(dataToReverse: string | T[]): string | T[] {
  if (typeof dataToReverse == "string") {
    return dataToReverse.split("").reverse().join("");
  }
  return dataToReverse.slice().reverse();
}
```

It is **important to put your "narrowest" definition on the top**. Overloads are read from the top to bottom. If you were to reverse this rule, the code using the overloading function will always land on the "widest" overload, making the DX bad (the "widest" overloads are usually there as a fallback).

### Function Overloads and the implementation signature

When writing function overloads, **make sure to list ALL possible specific cases the implementation function has to handle**. Since the **signature of the implementation function is NOT exposed to TS, your overloads have to handle all possible cases**. You cannot "cheat" by doing something like this:

```ts
function getRolePrivileges(role: "admin"): AdminPrivileges;
function getRolePrivileges(role: "user"): UserPrivileges;
function getRolePrivileges(role: string): AnonymousPrivileges {
  switch (role) {
    case "admin":
      return {
        sitesCanDelete: [],
        sitesCanEdit: [],
        sitesCanVisit: [],
      };
    case "user":
      return {
        sitesCanEdit: [],
        sitesCanVisit: [],
      };
    default:
      return {
        sitesCanVisit: [],
      };
  }
}
```

Notice that the implementation function returns specific type rather than an union of types. **This is wrong and TypeScript will complain**. The fix would be to ensure our overloads take care of all the values.

```ts
function getRolePrivileges(role: "admin"): AdminPrivileges;
function getRolePrivileges(role: "user"): UserPrivileges;
function getRolePrivileges(role: string): AnonymousPrivileges;
function getRolePrivileges(role: string): AnonymousPrivileges | AdminPrivileges | UserPrivileges {
  switch (role) {
    case "admin":
      return {
        sitesCanDelete: [],
        sitesCanEdit: [],
        sitesCanVisit: [],
      };
    case "user":
      return {
        sitesCanEdit: [],
        sitesCanVisit: [],
      };
    default:
      return {
        sitesCanVisit: [],
      };
  }
}
```

### Generics and Function Overloads

Keep in mind that you can use generic signature in the function overload signature.

```ts
function returnWhatIPassInExceptFor1<T extends string>(t: T): T;
function returnWhatIPassInExceptFor1<T extends number>(t: T): T extends 1 ? 2 : T;
function returnWhatIPassInExceptFor1(t: unknown): unknown {
  if (t === 1) {
    return 2;
  }

  return t;
}
```

### Variadic functions

Depending on the type user provided, **you can make the function take different amount of parameters, without any overloads**.

```ts
interface Events {
  click: {
    x: number;
    y: number;
  };
  focus: undefined;
}

export const sendEvent = <TName extends keyof Events>(
  event: TName,
  ...args: Events[TName] extends undefined ? [] : [payload: Events[TName]]
) => {
  // Send the event somewhere!
};

it("Should force you to pass a second argument when you choose an event with a payload", () => {
  // @ts-expect-error
  sendEvent("click");

  sendEvent("click", {
    // @ts-expect-error
    x: "oh dear",
  });

  sendEvent(
    "click",
    // @ts-expect-error
    {
      y: 1,
    },
  );

  sendEvent("click", {
    x: 1,
    y: 2,
  });
});

it("Should prevent you from passing a second argument when you choose an event without a payload", () => {
  sendEvent("focus");

  sendEvent(
    "focus",
    // @ts-expect-error
    {},
  );
});
```

Depending on the first parameter, the user might or might not have to provide the second parameter. **Notice the trick with an empty array and the spread operator**.

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

## The `object` type

The `object` type is meant to represent **all non-primitive types** in TypeScript. This is in **difference to the `Object` type** which **represents all primitive AND non-primitive types**.

### The usage

I would reach out for `object` type in the case where **I want to pass a "shallow" type to a "wider" type**. Here is an example.

```ts
interface User {
  id: string;
  name: string;
}

type IsUser<O extends Record<string, unknown>> = O extends { id: string } ? true : false;
type Result = IsUser<User>; // Type 'User' does not satisfy the constraint 'Record<string, unknown>'. Index signature for type 'string' is missing in type 'User'
```

The error happens because we are trying to **pass a very strict definition to a more generic one**. In this case **an interface with well defined keys into a `Record` type with unknown keys**. TypeScript will not let us to do that. Now, if I specify the `O` to extends the `object`, the `IsUser` generic type will work as expected.

```ts
interface User {
  id: string;
  name: string;
}

type IsUser<O extends object> = O extends { id: string } ? true : false;
type Result = IsUser<User>; // true
```

**Interestingly, if I were to type the `User` as a `type`, TypeScript would not complain**.

```ts
type User = {
  id: string;
  name: string;
};

type IsUser<O extends Record<string, unknown>> = O extends { id: string } ? true : false;
type Result = IsUser<User>;
```

It turns out that **this is the intended behavior**. As I understand it, since the `types` cannot be _augmented_ in place, it is safe to allow them to be "indexed". You can [read more about this here](https://github.com/microsoft/TypeScript/issues/15300#issuecomment-332366024).

### The problem

The example we have looked so far was about objects. Both the `Record` and the `object` type allow for objects. The problem is that **the `object` type allows for arrays and functions as well!**. In most cases, such behavior is undesirable, **but how could we create a "more strict" version of the `object` type**.

```ts
type MyObject = object;
const fakeObject: MyObject = () => null; // No errors. Behaves as per spec, but undesirable in our case.
```

There is a way to do so, but it is a bit hacky.

```ts
type MyObject = object;
const fakeObject: MyObject = () => null;

type MyStrictObject = object & { call?: never } & { bind?: never } & {
  push?: never;
};
const fakeObject2: MyStrictObject = () => null; // Error
const fakeObject3: MyStrictObject = []; // Error

const obj: Record<string, unknown> = {};
const realObject: MyStrictObject = obj; // Ok
```

We explicitly annotate some of the properties available to a function and an array as `never`. This ensures that we cannot assign them to the type.
This, of course, is very hacky, but I could not find any other way. The `Exclude` type did not work when using with `object` type.

## The `{}` type

This type **represents all values except `null` or `undefiled` values**. This type is **different than the `object` type**.

```ts
type Maybe<T extends {}> = T | null | undefined;

Maybe<null>; // error
Maybe<undefined>; // error
Maybe<false>; // ok
Maybe<() => null>; // ok
```

If you go by the definition from above, it should not be surprising that this works. Strings, Booleans and other types have multiple of properties on them, thanks to _boxing_.

```js
true.valueOf();
"ss".includes("xx");

// and so on...
```

As for another definition – **the `{}` type represents all the types that will not throw if you access any kind of property on them. Only `null` or `undefined` will throw in such case**. You can also think of this type **as top-level TypeScript type**. It allows you to describe every possible type, just like `unknown`.

## Enums

Enums are quite popular with _Ngrx_. They are not all sunshine and rainbows though.

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

**BUT BEWARE** Enums cannot be used with _plugin-transform-typescript_ which you are probably using.

### Enums considered harmful

You learned that the `enum` keyword is not ideal as it could bloat your bundle and that the `const enum` is a better alternative. Now the question is **should you even use enums at all?**. Keep in mind that enums are **not native to JS** so you are introducing a feature that is purely TS related. This might or might not be a problem when `enum` keyword is introduced to the JS language.

**But I think the biggest argument AGAINST using the `enum` keyword** is the fact that you can leverage the native JS object to achieve the same result.

```ts
enum MyEnum {
  foo,
  bar,
}

declare function withTSEnum(enumValue: MyEnum): void;
withTSEnum("foo"); // error!
withTSEnum(MyEnum.foo); // ok

const nativeEnum = {
  foo: "foo",
  bar: "bar",
} as const;
type NativeEnumValues = keyof typeof nativeEnum;

declare function withObjectEnum(enumValue: NativeEnumValues): void;
withObjectEnum("foo"); // ok
withObjectEnum(nativeEnum.foo); // ok
```

I would argue that the `enum object` pattern is even better than the TS version. Notice that you can provide literal values as well as the property of a given object. **If you are not convinced, consider looking at the following resources**.

1. The TS documentation about [`const enum` keyword pitfalls](https://www.typescriptlang.org/docs/handbook/enums.html#const-enum-pitfalls).

2. This [youtube video explaining the state of the `enum` keyword](https://www.youtube.com/watch?v=jjMbPt_H3RQ).

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
declare function get<K extends string>(key: K): <Obj extends HasKey<K>>(obj: Obj) => Obj[K];
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
  wojtek: "ala 123";
}
KeyAt<SomeInterface, "wojtek">; // 'ala 123', literal type!
```

This allows us to _pluck a given type_ out of object. This makes sure that return our function has return value correctly typed.

```ts
declare function get<K extends string>(key: K): <Obj extends HasKey<K>>(obj: Obj) => KeyAt<Obj, K>;
```

Personally i would name this type `TypeAt` but I'm going to roll with this name
paying an homage to original author :).

### Traversals

Our function can also filter stuff. We basically want to work _lenses-like_.

Example:

```ts
get(
  matching((friend) => friend.friends > 5),
  "name",
)(obj.friends);
```

With our current implementation this operation is impossible. How would we
enable such functionality?

Let's type `matching` first:

```ts
interface Traversal<Item> {}

declare function matching<A>(filteringFunction: (a: A) => boolean): Traversal<A>;
```

We have to change our implementation a bit to introduce `matching`.

```ts
declare function get<Item, K extends string>(
  traversal: Traversal<Item>,
  key: K,
): <Obj extends Array<HasKey<K>>>(obj: Obj) => Array<KeyAt<Obj, K>>;

const popularFriends = get(
  matching((user: User) => user.friends.length > 5),
  "name",
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
export type Unpack<F> =
  F extends Array<infer Item>
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
  arg: Head<Parameters>,
) => HasTail<Parameters> extends true ? CurryV0<Tail<Parameters>, ReturnType> : ReturnType;

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

This restriction stems from TS itself, you can though reference a type from within an object type just like we are doing with our first `Last` implementation.

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
type Prepend<TypeToPrepend, Tuple extends any[]> = ((head: TypeToPrepend, ...tail: Tuple) => any) extends (
  ...args: infer U
) => any
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
type Drop<ElementsToDrop extends Number, TupleToDropFrom extends any[], Iterator extends any[] = []> = {
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
type ValidateShape<T, Shape> = T extends Shape ? (Exclude<keyof T, keyof Shape> extends never ? T : never) : never;
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

### At least one item within the array

You are probably familiar with multiple ways one might declare an _array_ type.

```ts
type ArrayOfNumbers1 = number[];

type ArrayOfNumbers2 = Array<number>;
```

There is no semantic difference between these two declarations.
They work the same apart from their combination with the `readonly` keyword (the `Array` cannot be used with `readonly`. To achieve the desired effect one need to use `ReadonlyArray`).

The `number[]` and `Array<number>` can be assigned to variables holding empty arrays. I would argue that in most cases this is a desired outcome.

```ts
type ArrayOfNumbers1 = number[];

const numbers: ArrayOfNumbers1 = [];
```

But what if I wanted to ensure that the `numbers` variable holds at least one number, in other worlds is not empty.
Here is how one might do this.

```ts
type AtLeastOneNumber = [number, ...number[]];

const emptyArray: AtLeastOneNumber = []; // error

const exactlyOneNumber: AtLeastOneNumber = [1]; // ok
const oneOreMoreNumbers: AtLeastOneNumber = [1, 2]; // ok
```

The _spread_ syntax does not have the same limitations as the one in JavaScript. You can spread first and then use a concrete type.

```ts
type LastString = [...number[], string];
```

Pretty cool!

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

Syntax with `?` is much cleaner, especially with nested objects and properties. You no longer have to worry about checks with `&&`. `?` operator takes care that for you.

## `const` assertion

There are _literal_ and _primitive_ types in typescript. You know, there is a difference between these 2 types:

```ts
type Literal = "string";

type Primitive = string;
```

Usually, we do not want to _mutate_ anything, especially while working with objects and so on. Sine objects and arrays are mutable, typescript was widening the type of the property to accommodate for that (also for `let` declarations)

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

## The `satisfies` keyword

The `satisfies` keyword allows us to check if a variable annotated with the `satisfies` adheres to any given subset of a given type.
You would mainly use this **for objects where the TS infers the widest possible type** instead of a narrow slice that your object adheres to.

Here is what I mean. Let us say we have an object where the keys can be of different values.

```ts
type Colors = "red" | "green" | "blue";
type RGB = [red: number, green: number, blue: number];

const palette: Record<Colors, string | RGB> = {
  red: [255, 0, 0],
  green: "#00ff00",
  blue: [0, 0, 255],
};

const redComponent = palette.red; // string | RGB
```

The `redComponent` variable points to a correct type – the widest possible definition the property can have within the `palette` object. But it would be nice if we could narrow it down to only the `RGB` tuple type.

```ts
type Colors = "red" | "green" | "blue";
type RGB = [red: number, green: number, blue: number];

const palette = {
  red: [255, 0, 0],
  green: "#00ff00",
  blue: [0, 0, 255],
} satisfies Record<Colors, string | RGB>;

const redComponent = palette.red; // [number, number, number]
```

Note that **the `satisfies` keyword works DIFFERENTLY THAN the `as` keyword**. Using the `as` keyword instead of `satisfy` would yield the original result – the `redComponent` having the widest possible type definition.

### What about the `const` assertion?

The `const` assertion works well in situations where you want to narrow down the inferred types but where you do not have a concrete type defined.
If you have a concrete type defined, like the `RGB` in the example prior, the `satisfies` keyword is the preferred way to narrow down types.

### Interaction with functions

TODO: not implemented yet. I suspect this will allow us to infer function types and also provide generics? (the problem with `typeof`)

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
declare function call<R, F extends (...args: any[]) => any>(fn: F, ...args: GetFunctionArguments<F>): R;
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

What is interesting is that when you pass the parameters like so

```ts
type Props = {
  key: string;
};

const doWork = (props: Props) => {};

// TypeScript complains
doWork({ key: "1", somethingElse: "2" });
```

TypeScript would complain.

You can learn more about this behavior (and others, especially for `Object.keys`) in [this video](https://portal.gitnation.org/contents/understanding-types-as-sets).

## Testing for `never`

Let's say that for some reason, you want to test if the _type parameter_ that you defined is of type `never`. Normally I would do something like this

```ts
type CheckForNever<T> = T extends never ? 1 : 0;

type Test1 = CheckForNever<"hi">; // 0 => Ok.
type Test2 = CheckForNever<never>; // never => Wtf?
```

To understand what is happening, we need to talk about three things.

### `never` is a subtype of every type

While introducing the `never` type, Anders Hejlsberg wrote that the _never is subtype of every type_.
Implication of this is that `never extends T` is always true, no matter what `T` is.

### naked vs `clothed` type parameters

Whenever you use conditional types, eg.

```ts
type ConditionalCheck<T> = T extends string ? true : false;

type Result = ConditionalCheck<number>;
```

you pass _type parameters_ to the _type function_ (I'm not sure if this is an official name, but I like to think of it that way).

There you pass _naked type parameter_ if that parameter is not wrapped within an other structure (in TS land). An example of a _clothed_ type parameter would be passing `[number]` as `T` to the _type function_.

When conditional types are used with the _naked type parameter_, they are called _distributive conditional types_.
The name sounds scary, but the underlying functionality is rather not that complicated. Let us jump in!

### Distributive conditional types

The notion of _distributive conditional types_ is best explained using an example.

```ts
type IsNumber<T> = T extends number ? true : false;

type Test1 = IsNumber<number>; // => true

type Test2 = IsNumber<number | string>; // => boolean
```

So why is the `Test2` boolean? Let us unpack what TypeScript is doing under the hood when `IsNumber<number | string>` is evaluated

```ts
type IsNumber<T> = T extends number ? true : false;

type Test2 = IsNumber<number | string>;
// Is the same as
type Test2 = IsNumber<number> | IsNumber<string>;
// Is the same as
type Test2 = true | false;
// Is the same as
type Test2 = boolean;
```

Notice the transformation from `IsNumber<number | string>` to `IsNumber<number> | IsNumber<string>`. **This is called distributing over an union type**.

If the `number | string` would be _clothed_, ie. wrapped in a type, the union of `IsNumber<number> | IsNumber<string>` would not be evaluated.

So when you see _distributive conditional types_, just think of the underlying type parameter being split into multiple unions (or a single union if the type parameter is a singular type, eg. `number`)

### Distributive conditional types and the `never` type

What about `never` in this context? Going back to our original example

```ts
type CheckForNever<T> = T extends never ? 1 : 0;

type Test1 = CheckForNever<"hi">; // 0 => Ok.
type Test2 = CheckForNever<never>; // never => Wtf?
```

The `Test2` is evaluated to type `never` **due to the fact that when `never` is passed as a type parameter, TypeScript distributes over an empty union**. Since the union is _empty_ and there is nothing to distribute over, the result is a `never` type.

This is not the case with other, less _special_ types like `number` or `string` (or others), where if you pass them as a type parameter, TypeScript distributes of a union with one member - that type.

Here is a good mental model to think about the `never` type in terms of an union

```ts
type never = | // explicitly marked as an empty union. NOT VALID TS SYNTAX!

type boolean = true | false
type string = string // union with 1 member
// ...
```

### The solution to our problem

It should be clear to us that as long as we use the _naked_ `never` type in a conditional type context, the result will always be `never`.
How can we express our intent without distributing over an empty union? (See _distributive conditional types_ if you are still unsure what that means)

Well, **we learned about _clothed_ types right**? When a type parameter is a singular _clothed_ type, it will not be a subject to the _distribution_.

With that information, all we have to do is to amend our existing snippet just a tiny bit

```ts
type CheckForNever<T> = [T] extends [never] ? 1 : 0;

type Test1 = CheckForNever<"hi">; // 0 => Ok.
type Test2 = CheckForNever<never>; // 1 => Ok.
```

Notice that I opted to wrap the type parameter within the type function itself. I view it as an encapsulation mechanism, but you could also write it like so

```ts
type CheckForNever<T> = T extends [never] ? 1 : 0;

type Test1 = CheckForNever<"hi">; // 0 => Ok.
type Test2 = CheckForNever<[never]>; // 1 => Ok.
```

As long as the type parameter is wrapped at some point, the inner type will not be distributed and the "issue" with an empty union will not occur.

## Generics and Inference

Whenever you use _generic type parameters_ with default types, TS compiler will (in most cases) to infer that _type parameter_ from the values that you provided.

To give you an example

```ts
function foo<T>(arg: T): T {
  return arg;
}

foo(1); // 1
foo("name"); // "name"
```

That is OK in most cases, but what happens if you want semantics like these to be at your disposal

```ts
type MyObj = {
  code: "MY_CODE";
};

type BaseObj = {
  code: string;
};

function foo<Obj extends BaseObj = MyObj>(obj: Obj): Obj {
  return obj;
}

foo({ code: "MY_CODE" }); // Ok

foo({ code: "SOMETHING_ELSE" }); // Should not be allowed, but is - through inference

foo<{ code: "MY_CUSTOM_CODE" }>({ code: "MY_CUSTOM_CODE" }); // Ok

foo<{ code: "MY_CUSTOM_CODE" }>({ code: "SOMETHING_ELSE" }); // Error as it should be
```

As you can see, with our naive implementation of `foo` function, one use-case was not met. The `foo({ code: "SOMETHING_ELSE" });` snippet is not producing TS errors because of _type parameter_ inference. TS compiler sees that you provide a `string`, thus the `code` type will be inferred as string.

In other worlds, **the TS compiler will always expand the generic parameter to the widest possible type available**.

### The golden rule of Inference and Generics

Like in the case of function parameters, **one can declare too many _generic type parameters_**. This is a smell and should be fixed. The reason is that **it can make it really hard for other people to understand what the function type parameters are for**. Imagine seeing a lot of generics slots with multiple duplicate values inferred. Not that fun right?

So, **always declare the generics as close to where they are assigned** and **keep the number of generic parameters as low as possible**. Like in programming and in life – less is more!

### Lazy type evaluation - prevent type parameter inference

In this situation, what we want is for **TypeScript to defer the inference of the `Obj` until we provide it explicitly**. This is not possible without tricks as **the default behavior of TypeScript is to always infer the type parameter from it's first occurrence**.

The best example would be a `compare` function.

```ts
declare function compare<A>(a: A, b: A): boolean;

compare(1, "123"); // "123" is not assignable to type number
compare("123", 1); // 1 is not assignable to type string
```

Notice that the type parameter `A` got inferred from it's first occurrence. We can **change this behavior by leveraging lazy evaluation of the type parameters**. This is possible **with conditional types, as when TypeScript sees `?`, it will defer the inference until after the T is resolved**.

```ts
type NoInfer<T> = [T][T extends any ? 0 : never];

declare function compare<A>(a: NoInfer<A>, b: A): boolean;

compare(1, "123"); // 1 is not assignable to type string
compare("123", 1); // "123" is not assignable to type number
```

Now it's the second parameter that is used for the inference. How this method could help us solve our problem?

---

Let us use the `NoInfer` type and apply it to our problem. All of this is inspired by [this StackOverflow answer](https://stackoverflow.com/a/56688073).

```ts
type NoInfer<T> = [T][T extends any ? 0 : never];
type MyObj = {
  code: "MY_CODE";
};

type BaseObj = {
  code: string;
};

function foo<Obj extends BaseObj = never>(
  // See the section for comparing with `never` type
  obj: [Obj] extends [never] ? MyObj : NoInfer<Obj>,
): [Obj] extends [never] ? MyObj : NoInfer<Obj> {
  return obj;
}
```

As I mentioned, this _workaround_ is leveraging the fact that if a _type parameter_ is used in a context of conditional type, it will be evaluated lazily as in the inference will not occur.

### Type narrowing

Sometimes all you work with is a "narrow" type, but you want TS to infer the most strict type possible – think a `string[]` and a tuple of strings.

For this case, consider using **the `F.Narrow` function from `ts-toolbelt`**.

```ts
const makeRouter3 = <TConfig extends Record<string, { search?: string[] }>>(config: TConfig) => {
  return config;
};

const t = makeRouter3({ foo: { search: ["bar", "baz"] } });
t.foo["search"]; // string[]
```

With `F.Narrow` the situation changes.

```ts
const makeRouter3 = <TConfig extends Record<string, { search?: string[] }>>(config: F.Narrow<TConfig>) => {
  return config;
};

const t = makeRouter3({ foo: { search: ["bar", "baz"] } });
t.foo["search"]; // ["bar", "baz"]
```

#### Type narrowing by providing a default generic parameter

In some cases, especially when you perform a recursive build-up of types (think builder pattern where each method call "enhances" the given object), it is vital to provide a default type for the generic. **If you do not, TypeScript will infer the widest possible type from the generic, which will ruin the constraints on your type**.

Here is what I mean.

```ts
type Foo = Record<string, string> & Record<"foo", string>; // The type will allow for any string
type Bar = {} & Record<"foo", string>; // The type will allow only for the "foo" key.
```

This is very important in generic signatures.

```ts
class Builder<TMap extends Record<string, string> = {}> {
  constructor(private map: TMap) {

  }

  set<K extends string>(key: K, value: string): Builder<TMap & Record<K, string>> {
    (this.map[key] as any) = value;
    return this.
  }
}
```

Without the `= {}` TypeScript would always use the `Record<string, string>` type. No matter how many times you call set, the widest possible type will always win with more specific one.

#### Type narrowing with _const annotations_

When creating a variable, you have the ability to add `as const` to it. This will make the TypeScript **infer the literal types of the values**.

```ts
const foo = "bar" as const; // foo is of type "bar"
const foo2 = "bar"; // foo2 is of type "string"
```

This is very handy – it enables you to be strict and explicit with the values you are passing (keep in mind that the `as const` also adds `readonly` to object properties). **What is nice is the ability to use the _const annotations_ in the context of generic parameters**. Starting from TypeScript v5, the following is a valid code.

```ts
declare function router<const R extends readonly string[]>(routes: R): Record<R[number], unknown>;
const foo = router(["a", "b"]); // Record<"a" | "b", unknown>
```

If I were to create a similar function signature in TypeScript v4.x, the inference would fallback to the `string`.

```ts
declare function router<R extends string[]>(routes: R): Record<R[number], unknown>;
const foo = router(["a", "b"]); // Record<"string", unknown>
```

Keep in mind that, **for the _const annotation_ to take an effect here, one has to use the `readonly` modifier on the array**. Otherwise TypeScript will fall back to the old behavior.

In addition, if you use the `const R` generic syntax, all the properties of R will become `readonly`. This might or might not be what you want. Keep that in mind.

### The `could be instantiated with an arbitrary type which could be unrelated to` problem

When working with generics, you might encounter the following error message:

> TGenericParameter could be instantiated with an arbitrary type which could be unrelated to OtherTGenericParameter

This mostly happens because:

1. **TypeScript is not capable of reasoning about your code**.
2. TypeScript prevents you from making a mistake in the types.

The situation number 1 often happens when you are typing a _higher order function_.

```ts
export const withRouter = <TProps>(Component: React.ComponentType<TProps>) => {
  const NewComponent = (props: Omit<TProps, "router">) => {
    const router = useRouter();
    /**
     * TypeScript is not able to deduce that
     * Exclude<T, k> & { [k]: T[k] } === T
     *
     * https://github.com/microsoft/TypeScript/issues/35858#issuecomment-573909154
     * We have to past props to `props as TProps`
     */
    return <Component {...props} router={router} />;
  };

  NewComponent.displayName = `withRouter(${Component.displayName})`;

  return NewComponent;
};
```

Like I wrote in the comment, TypeScript is not capable of understanding that `Omit<TProps, 'router'> & {router: Router}` is the same as `TProps`.

--

As for the option number 2. This one ensures that you do not make a silly mistake with the types.

```ts
// Type '"bar"' is not assignable to type 'T'.
// '"bar"' is assignable to the constraint of type 'T', but 'T' could be instantiated with a different subtype of constraint '"foo" | "bar"'.
const foo = <T extends "foo" | "bar">(arg: T = "bar") => {};
```

This makes sense as you could, in theory, in your code, do the following:

```ts
foo<"foo">();
```

Which then creates a really weird situation for the default `arg` value.

## Type branding (AKA _opaque types_)

Imagine you have a function that converts EURO to USD. Here is how one might write the type declaration for this function.

```ts
declare function euroToUSD(euro: number): number;
```

You most likely want to ensure that the `euro` parameter is not a negative number. If it is, we could throw an error (or gracefully convert it to a positive number, which might make it impossible to find potential bugs in our code).

```ts
function euroToUSD(euro: number): number {
  if (euro < 0) {
    throw new Error("Cannot be a negative number");
  }

  // Logic...
}
```

From the developer ergonomics perspective, would it not be nice to have TypeScript do some of the work of checking whether something is a positive number for us? Enter **type branding**. The concept is similar to Go type aliases, where you can create more specific versions of a type derived from the base type.

```go
type MyCustomIntType = int
```

The `MyCustomIntType` **is NOT** the same as the `number` type. They are different, even though the base type is the same. The compiler will not let you pass a regular `int` type where the `MyCustomIntType` is required. Sadly, TypeScript does not have a native way to represent "different" types that derive from the same base type.

```ts
type MyCustomNumber = number;
let foo: MyCustomNumber = 3;
let bar: number = 3;

foo = bar; // OK, but it should throw an error
```

But, do not lose hope! We can still make it work. Almost as if we had the same type aliases Go has – by **branding the derivative type using a unique identifier**.

```ts
type MyCustomNumber = number & { __brand: "MyCustomNumber" };
let foo = 3 as MyCustomNumber;
let bar: number = 3;

foo = bar; // Error!
```

Of course, the ergonomics of the type branding are not ideal. You have to cast the underlying value to a given type. But, with some casting, we now have quasi Go (and possibly other languages) type aliases working.

Now, we can go back to our `euroToUSD` function definition and apply the type branding technique we have just learned about.

```ts
type Euro = number & { __brand: "Euro" };
declare function euroToUSD(euro: Euro): number;
```

For ultimate type safety, you can scope the returned value as well!

```ts
type Euro = number & { __brand: "Euro" };
type USD = number & { __brand: "USD" };

declare function euroToUSD(euro: Euro): USD;
```

Remember that you will have to cast the returned value to the `USD` type. Otherwise TypeScript will yell at you that it is impossible to convert a `number` to `USD`,

### Branding non-primitive types

The technique of type-branding also allows you to _brand_ any complex type, like a shape of an object. To give you an example.

```ts
interface Credentials {
  password: string,
  email: string
}

type ValidCredentials = Credentials & {_brand_: "Valid"};
const validateCredentials = (credentials: Credentials) => ValidCredentials {
  // ... logic
}
```

Here, I would _brand_ each member of the `Credentials` interface, but I wanted to show you that branding "complex" shapes is also possible. In the end this technique uses _type intersections_ so it applies to every type.

### Type branding for objects

Keep in mind that you can also **use branded types as keys of an object**. This creates an interesting dynamic where, based on the key you provided, a different type is returned.

```ts
type UserId = ...
type User = ...

type PostId = ...
type Post = ...

type DB = Record<UserId, User> & Record<PostId, Post>
```

If I access the `DB` via the `PostId`, I will get a `Post`. If I access it via the `UserId`, I will get a `User`. How cool is that!?

### Type branding and type guards (predicates)

For the ultimate type safety, one could combine both techniques. Check this out.

```ts
type Valid<T> = Brand<T, "Valid">;

interface PasswordValues {
  password: string;
  confirmPassword: string;
}

const isValidPassword = (values: PasswordValues): values is Valid<PasswordValues> => {
  if (values.password !== values.confirmPassword) {
    return false;
  }

  return true;
};
```

When we call the `isValidPassword` we get the correct type – the `PasswordValues` wrapped within the `Brand` type. Pretty neat!

## Making the Branded Types suck less (subjective)

If you decide to use the _type branding_ technique without any abstractions, you will be type-casting all over the place. Not ideal. In my humble opinion, the **_type branding_ combines well with a function that explicitly casts the type for you**. While that function is very shallow, and violates one of the core principles of good design, I like it much better than having `as XX` scattered thought the codebase.

```ts
type Email = string & { _brand_: "EMAIL" };

// The name is very contrived. Change it according to your use-case.
const asBrandedEmail = (email: string): Email => {
  return email as Email;
};

const foo: Email = asBrandedEmail("wojciech@stedi.com");
```

## TypeScript and Classes

The `class` **can act both as a runtime construct and a type**. This is quite rare in TypeScript (the only other construct I can think of, that behaves the same way, is the `enum` and `const enum`).

This means you can do something like the following.

```ts
class MyError extends Error {}

function foo(error: MyError) {} // used as a type

const error = new MyError(); // used as a runtime value
```

### Type Predicates and Classes

You can **leverage the type predicates (guards) on the methods of the class**. This is pretty wild. Look at the following example.

```ts
class UserValues {
  error?: string;
  constructor(private values: Record<string, string>) {}

  // Notice this wild piece of syntax
  public areValid(): this is this & { error: string } {
    // validates that `this.values` are valid
  }
}

// On the caller side

const userValues = new UserValues({ firstName: "", lastName: "" });

if (!userValues.areValid()) {
  userValues.error; // string
  return;
}

userValues.error; // string | undefined
```

Since we do not have a parameter for the `areValid` function, the only way to refer to the "current scope" is through the use of `this` keyword.
I have to say, the `this is this &` syntax is quite funny :D.

Please note that the **type predicates can also influence the types inside the class itself**.

```ts
class UserValues {
  error?: string;
  constructor(private values: Record<string, string>) {}

  // Notice this wild piece of syntax
  public areValid(): this is this & { error: string } {
    // validates that `this.values` are valid
  }

  public login() {
    const areValid = this.areValid();
    if (!areValid) {
      this.error; // string
      return;
    }

    this.error; // string | undefined
  }
}
```

This is quite amazing. You can change the type of `this` dynamically. It is like "self mutation".

### Assertion functions and Classes

**Like in the case of type predicates, you can use the assertion functions inside classes**. The syntax is basically the same as the regular assertion function and works on the same basis as type predicate.
