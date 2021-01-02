# Production TypeScript FrontendMasters

## New notable features

### _Recursive type aliases_

This one is pretty nice, you no longer have to use weird hacks and know that something is evaluated eagerly and other things are evaluated lazily.
For example, this is how you could type a _type_ which allows only _json values_

```ts
type JSONValue =
    | number
    | string
    | null
    | boolean
    | JSONValue[]
    | { [k: string]: JSONValue };
```

Notice that I'm referring to `JSONValue` twice within it's declaration. Previously this was not allowed.

### _Template type literals_

This is pretty neat, I still have to play around with it but it allow you to basically create permutations of _string literals_.

```ts
type Corner = `${'top' | 'bottom'}-${'left'|'right'}`
```

I think it main usage will be in automatically typing _query parameters_ for you. Since it works with generics, you can do a lot of cool stuff here.

### _@ts-expect-error_

This is pretty cool. In the era of linters we sometimes use _SOME_LINTER-ignore_ to just, well, ignore the linter for that line or file.
TypeScript also has something like this, called `@ts-ignore`. That _directive_ is basically saying

> There might be an error on the next line, or might not. Either way, just ignore it

So what is `@ts-expect-error` for?

`@ts-expect-error` is just a better version of `@ts-ignore` (most of the times). Let's say I have situation like so

```ts
// bar.ts
type Bar = number;

// implementation.ts

// @ts-expect-error
const num1: Bar = "hello";

// @ts-ignore
const num2: Bar = "hello";
```

In the snippet above, both directives smilingly behave the same, now, what will happen if we change the typings of `Bar`?

```ts
// bar.ts
type Bar = string;

// implementation.ts

// @ts-expect-error
const num1: Bar = "hello";

// @ts-ignore
const num2: Bar = "hello";
```

In this case, the `@ts-expect-error` will yell at you that the directive is not used. This is correct as **there was a change in typings, some type information might be lost**.
While it's not the case here, you can imagine that having something that will alert you that typings have changed is pretty useful.

Basically as a rule of thumb I would **always prefer `@ts-expect-error` over `@ts-ignore`**.

### Typing `try/catch` blocks

As you probably know, the `error` that is passed to the `catch` block is of type `any`. This is suboptimal because you can basically _throw_ anything as an error (`React` throws `Promise` instances in concurrent mode).

There was a neat addition to TypeScript, where you can type the error that is passed to `catch` block as `unknown` and **only** `unknown`. This forces you to do necessary checks before accessing any properties on that error

```ts
declare function somethingRisky(): number;
try {
    somethingRisky();
    // e can literally be anything, strings, numbers, objects, you name it!
} catch (e: unknown) {
    if (e instanceof Error) {
        console.log(e.stack);
    }
}
```

### _Type assertions / assert signatures_

This is, IMO, pretty niche but still useful when I'm writing tests.
So, you are probably familiar with a notion of _type guard_. To give you an example, the `if` check from the snipped above could have been written like this

```ts
function isError(err: unknown): err is Error {
    return err instanceof Error;
}
```

With _type assertions_ or some might call them _assert signatures_ we will not rely on the boolean value returned from that function, instead we will throw an error if the assertion does not pass

```ts
function assertIsError(err: unknown): asserts err is Error {
    if (!(err instanceof Error)) throw new Error("not an error!");
}
```

This looks pretty weird but I use it all the time while testing code that is using _feature flags_. They usually have a name, some kind of identifier, and I want to make sure that the identifier passed to the function that gets the _feature flag_ value is correct.

So, let's say I want to mock the function that gets the feature flag value, I would do it like so

```ts
function assertIsCorrectName(name: string): asserts name is "MY_FEATURE_FLAG" {
    if (!name == "MY_FEATURE_FLAG") {
        throw new Error(`expected MY_FEATURE_FLAG, got: ${name}`);
    }
}

jest
    .spyOn(featureFlagService, "getFeatureFlag")
    .mockImplementation((featureFlagName) => {
        assertIsCorrectName(featureFlagName);
        return true;
    });
```

This increases the confidence I have in my test suite. Every time I change the identifier for the feature flag, I should also be changing my tests.

### _type only imports_

While working with TypeScript you might have noticed that you can import types along with exported symbols without any problems, like so

```ts
// type, javascript symbol (function)
import { getUser, UserType } from "./user";
```

This is completely ok, but sometimes, the type has the same name as the symbol

```ts
// type, javascript symbol (object)
import { User, User } from "./user";
```

This is possible because TypeScript _lives in different namespace_ if you will from your JavaScript code, this is why having types named the same as your symbols is possible.
Normally this is not a problem, but sometimes, the compiler might get confused which thing to import and which thing to just strip away while bundling / compiling code.
This used to be a problem while working with _firebase_. (or really working with code splitting in general, where you import types from modules that you are splitting)

```ts
// are we importing the namespace or the symbol?
import { Firebase } from "firebase";
```

A lot of complains were filled that even though some people used only _modular imports_ of the _firebase_ library, they had the whole library included in their bundle. This is because the bundler could not
determine (sometimes) if the import is for a typing or a JS symbol.

To combat this issue a _type only imports_ were introduced.

```ts
import type { Firebase } from "firebase";
```

Now there is no ambiguity with the import. The bundler exactly knows that this should be stripped, because of the `import type` syntax.

## Creating a library - takeaways

- use `npx gitignore node` for generating `.gitignore` files for node environment.

- use `volta` to pin engines and libraries. `volta` is like better `nvm`.

- consider opting-in for `types` option. Since this will force you to explicitly list the _typings_ that are available to you, you are reducing a chance of someone referencing a `devDependency` module in the `src` folder. If you need to split things between testing and development, consider using 2 `tsconfig.json` with different `include` statements.

- use `eslint --init` to setup `eslint` config fast

- consider adding separate `tsconfig` for `eslint`. You can specify the `projects` option

- it seems like the `babel-jest` (so really _babel_ + _jest_) is faster than `ts-jest`

- `api-extractor` seems like a great tool for automatically generating documentation. It's pretty legit, try it!

- `api-extractor` without a `--local` flag can be used like a validation mechanism for changes in _API surface area_.

- you can use `api-documenter` to automatically create docs from the output of the `api-extractor`. **You can then host those on github pages with like 0 work**. This is so great!

- `api-extractor` and `api-documenter` is maintained by Microsoft. This makes this library no-brainier for creating documentation.

- consider using `noImplicitReturns`. This I hate this rule personally, Mike made a good point that it's too easy to change the intent of given function. I would consider using it in any `.ts` file. The `.tsx` seems to be a bit overkill.

- there is this weird syntax which is equivalent to `module.exports = callableValue`

  ```js
  export = callableValue;
  ```

  I would avoid it as much as I can, but this is how you would migrate the `module.exports` as a _non-namespace_

### Notable `tsconfig` options

- `esModuleInterop` treats `cjs` modules as if they have _default export_. To avoid using this flag, which again is kinda parasitic since if you enable it, your consumers also have to, use the `import something = require('./something')` syntax.

- `skipLibCheck` can be dangerous. If you have this on and _you have to have this on for your library to compile_ any consumer of your library also has to do it. This is something you might not want to happen. Be careful here. I would turn this on, only within apps context.

## Converting from JS to TS

Here just look at the notes [in the repo](https://github.com/mike-north/professional-ts).

## Pure type information

This is where you have to define types for yours app entities. Mike decided to simply use `types.ts` file. I'm all up for it, seems reasonable.

## 3rd party typings and overwrites

_Definitely Typed_ is interesting. By default any merge request that is not marked as _needs work_ will be merged. Due to this, you are taking upon a lot of _mystery code_. You do not want bugs from such code block you. To make sure this is not the case, you will need to know how to overwrite 3rd party typings.

1. Create _types_ directory in the root of your project. This is where we will be putting our overwrites.

2. Use `tsconfig` _paths_ option

```json
"paths": {
    "*": ["types/*"]
  }
```

This tells the compiler to also include the `types` directory when resolving module types. So let's say you are importing _React_ module

```js
import React from "react";
```

- compiler will first look into the `types` directory for `react.d.ts` file

- if the file is not there, compiler will look into `@types/react` for `react.d.ts` file

Please note that this will **completely overwrite _@types_ definitions**.

## Handling runtime data

Use _type guards_ with _assert signatures_. One neat trick Mike shown is how to type the _type guard_ itself when it's passed as a parameter

```ts
function isTypedArray<T>(
    arg: any,
    check: (val: any) => val is T,
): asserts arg is T[] {}
```

You would mostly use this when using _fetch_ API or similar. Basically anything that comes from the backend, maybe expect _GraphQL_ :)

The method presented is considerable faster than `propTypes`.

## Tests for types

This is very important topic. This is how you would **keep your _type-guards_ up to date**.

There are two tools you can use

- _dtslint_: used mostly in the _Definitely Typed_ repo. **Your test will run against multiple _TypeScript_ versions**. This is great stuff.
  One catch here is that the setup is not fun.

- _tsd_: much more straightforward setup than _dtslint_. Instead of comments it uses special generic types for assertions.

A simple test using `dtslint` would look as following

```ts
const team: ITeam = null; // $ExpectError
```

Notice the comment here. This is how you make assertions using _dtslint_.

Contrast that with _tsd_.

```ts
import { expectAssignable, expectNotAssignable } from "tsd";
import { ITeam } from "../../src/types";

expectNotAssignable<ITeam>(null);
```

The _tsd_ exposes a lot of assertions. The **only thing _tsd_ is missing is the ability to test multiple _TypeScript_ versions**.
