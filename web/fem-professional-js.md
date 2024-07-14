# Professional JS – Features You Need to Know

- JavaScript is not versioned.

  - That is different from other languages.

  - We used to have JavaScript versions very long time ago.

- We use ECMA specifications for versioning.

  - TC39 is the committee that establishes new ECMA specification versions.

- There are various tools you could use to see if the feature you want to use is available in given browser / Node version.

  - For Node, use the [node.green](https://node.green/).

  - For web, use the [caniuse](https://caniuse.com/).

- **ECMA will never introduce breaking changes**.

  - More changes are, in most cases, a newer syntax.

  - They do not want to break anyone.

- If you wish to use a new feature and the environment you are targeting does not support it, you have two options:

  1. Use _polyfills_: libraries that aim to "fill a hole" and emulate the new APIs.

  2. Use _transpilers_: will take your code and do their best to _convert_ your code to code that works in older browsers.

- **Before deciding what version of ECMA to ship, check what browsers your users are using**.

  - There is no point in _transpiling_ to older versions if you do not need to.

    - **The older the version you _transpile_ to, the bigger the JavaScript footprint**.

## ES6 recap

- ESModules

  - When importing modules, think URLs.

    - You most likely are importing files without their extension. You most likely have a _transpiler_ add the file extension for you.

  - One thing that I do not appreciate that Max highlighted is that **no matter how many times you import a module, it will only be executed once**.

    - You can import `script3.js` in _module A_ and _module B_. That `script3.js` will only execute once.

## Language Enhancements

This section only includes things that caught my attention. See [this website](https://firtman.github.io/projs/) for full list.

### `GlobalThis`

The "global" might be different depending on the environment your code runs in.

1. In Node, the "global" is `global`.

2. In browsers, the "global" is `window`.

3. In workers, the "global" is `self`.

To make accessing the "global" more ergonomic, you can use `globalThis` which **will always point to the "correct global" object**.

### Optional Catch Binding

You you do not have to "get" the error in the `try/catch` block.

```js
try {
} catch {
  // you do not have to use `catch(error)`
}
```

### Object Management

- The `Object.hasOwn` checks if the property is declared on the object rather than the _prototype chain_.

## Array and Collection Enhancements

### New Collections

- The `Set` will not contain any duplicates.

  - This is pretty neat when you need to de-duplicate data.

  - In some cases, it can also provide better ergonomics than plain old _array_.

- The `Map` allows for keys to be objects or even functions.

  - You can only use strings for keys when working with _objects_.

### Array `At`

I find this API most useful as alternative to `array[array.length-1]`.

```js
const array = [1, 2, 3, 4];
const last = array.at(-1); // 4;
```

### Change by Copy

The `sort` and `splice` mutate the underlying data. It is quite easy to forget that this is the case. If you do, you could mutate the parameters your function is receiving and that could be very confusing.

Instead of using `sort` and `splice` **consider using `toSorted`, `toReversed` and `toSpliced`. These methods WILL NOT mutate the underlying array**.

```js
const array = [1, 2, 3, 4];
const sorted = array.toSorted(); // the `array` is not mutated!
```

## Asynchronous Programming

### Promise Improvements

**Keep in mind that these APIs will NOT cancel the promises that you fired. If any other "ignored" promises fail, you might be looking at `uncaught exception` error**.

- We have `Promise.any` which will return the first resolved promise.

- Use the `Promise.fulfill` to return a _"promisified"_ data.

  - I do not buy the argument of "if you, in the future, would like to change the source of the data to be `async`, it will be easier to do", but the API is there!

## Advanced Techniques

### PTC - Proper Tail Calls

This is an internal change that, **if you structure your functions in a given way**, will make your recursive calls take up less frames. The less frames you use, the faster your code will execute. In addition, the danger of running into _stack overflow_ is pretty much zero.

```js
function factorial(n, acc = 1) {
  if (n === 0) {
    return 1;
  }

  return factorial(n - 1, n * acc); // Structured in a way to take advantage of PTC.
}
```

**The name "tail calls" refers to what the function returns. If the function returns _only_ a function call, it will be "optimized by the engine**.

Contrast the above example with the following:

```js
function factorial(n) {
  if (n === 0) {
    return 1;
  }

  return n * factorial(n - 1);
}
```

Notice that the "tail", so the return statement, is a multiplication.

### Error cause

This allows you to "chain" errors.

```js
try {
} catch (error) {
  throw new Error("some error", { cause: error });
}
```

**This is much better than stringing errors together in the message of the error**. In addition, the debugging experience is much better as you know can get the `.cause` of the error!

### Import Metadata

I'm very much used to using `__dirname` and similar APIs. These APIs do not work in ESModules. In ESModules you have to refer to other modules as URLs.

The `import.meta` is an object containing bunch of very useful properties that would allow you to re-create the `__dirname` and similar globals you are familiar with.

```js
import fs from "node:fs/promises";

const fileURL = new URL("./file.txt", import.meta.url); // join(__dirname, "file.txt");
```

## Wrapping up

The most exciting features for me are

- The ability to provide `cause` in the `Error` constructor.

- The `.at` function.

- The new functions to manipulate arrays without mutating the original array – `.toReversed`, `toSorted` and `toSpliced`.

- The new collections – `Map` and `Set`. I feel like I'm not using them enough and I default to an array when I would benefit from using either `Map` or `Set`.

- I finally understood what "tail" means in "PTC". Thank you Max!

  - As a reminder, the "proper" in "PTC" refers to the tail which is a function invocation rather than some operation.
