# General

## The native `.env` file support

- In v20.6.0 Node.js added [native `.env` file support](https://nodejs.org/en/blog/release/v20.6.0).

- Usage via the `--env-file` flag, like so: `node foo.js --env-file=config.env`

  - **You can define any `NODE_OPTION` in the `.env` file as well!** Pretty neat stuff.

- This **might make the `dotenv` package obsolete**.

### The `loadEnvFile` function

There is also the `loadEnvFile` function (stable in Node.js 24) that you can use. **This replaces the `dotenv()` function from the `dotenv` package**.

According to [this blog post](https://www.stefanjudis.com/today-i-learned/load-env-files-in-node-js-scripts/) the function can throw an error if the specified file is not there – something to be mindful of.

I'm very glad that those additions are _finally_ stable and in Node 🎉.

## The `node:` import prefix

When reading code, you might come across import statements that look like:

```ts
import fs from "node:fs";
import { readFile } from "node:fs/promises";
```

**The `node:` prefix is a convention to signal that you import something built-in to node rather than an external package**.

AFAIK, the benefit is purely related to semantics rather than functionality. Using this convention allows Node.js to introduce new packages which won't conflict with existing, very popular ones, due to that prefixing.

Some packages, like the `node:test` are NOT available via "bare" import without the prefix.

## Watch mode

Node.js can now watch for changes in given files. This is _very_ useful and historically was only possible due to 3rd-party packages!

```ts
// package.json
"dev": "node --watch --env-file=.env app.js"
```

## Native TypeScript support

It happened! Node.js 22.18.0 **allows you to execute `.ts` file WITHOUT any additional flags**.

Keep in mind that **this is not the same thing as running `tsc` on that file and then running it in Node**. There [are some limitations](https://nodejs.org/api/typescript.html#type-stripping).

Having said that, I find myself not using the things excluded from the native support that often, so I'm very happy that I finally can run my applications "natively" without relying on tools like `tsx` or `ts-node`.

## Support for the `using` keyword

I remember first learning about the `using` keyword when reading TypeScript documentation. See [announcement here](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-5-2.html). Then, I kind of forgot about it, since I have not had a need for it.

I **just recently learned that the `using` keyword is available in "vanilla" 24.0.0 Node.js (without nay flags)**. This is a huge win for the ecosystem!

### Why would you want to use it

One thing that I like Golang for is the `defer` keyword. In Golang, you can _defer_ code execution till the "end of the scope".

```go
func foo() {
  // code
  defer file.Close()
  // more code
}
```

**The `using` keyword is as close as we can get to `defer` in JavaScript**.

```js
import fs from "node:fs/promises";
import path from "node:path";

async function main() {
  // Some code
  await using file = await fs.open(path.join(import.meta.dirname, "./example.test.js"));
  // Some other code

  // File will be closed at the end of the scope
}
```

### Custom classes and `using` keyword

The neat thing is that you can use the `using` keyword in the context of a custom `class` or even a function.

```js
import fs from "node:fs/promises";
import path from "node:path";

class CustomResource {
  [Symbol.asyncDispose]() {
    console.log("Dispose custom resource");
  }
}

async function thisWorksAsWell() {
  await using cleanup = new AsyncDisposableStack();

  const file = await fs.open(path.join(import.meta.dirname, "./example.test.js"));
  cleanup.defer(async () => {
    await file.close();
    console.log("Cleanup");
  });
}

async function main() {
  console.log("First log");
  await using resource = new CustomResource();
  console.log("Second log");
  await using foo = await thisWorksAsWell();
  console.log("Third log");
}

await main();
/* 
First log
Second log
Cleanup
Third log
Dispose custom resource
*/
```

The new `AsyncDisposeStack` and `DisposeStack` allow you to _append_ to a disposable stack in different branches of the code

```js
async function conditional({ someCondition }) {
  await using cleanup = new AsyncDisposableStack();

  const file = await fs.open(path.join(import.meta.dirname, "./example.test.js"));
  cleanup.use(file);

  if (someCondition) {
    const otherFile = await fs.open(path.join(import.meta.dirname, "./index.html"));
    cleanup.use(otherFile);
  }
}

async function main() {
  // No need to call `using` here.
  // The `await using cleanup` already does all the cleanup inside the `conditional` function
  await conditional({ someCondition: true });
}

await main();
```

This is pretty magical if you ask me.

**And the best part? It works in most browsers as well!** Currently, Safari is lagging behind, but I believe it's only a matter of time.
