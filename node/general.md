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
