# ESM in Node

It seems that, at the time of writing this, we are finally embracing the ESM module system for Node.JS on a bigger scale. This is great news! Having the interop between browser and Node is insane!

## Enabling ESM modules

You can enable ESM modules within your Node environment using one of three ways

- use `.mjs` extension files
- add `"type": "module"` to your `package.json` file - this is my preferred option
- use `--input-type=module` while working with STDIN data

## Biggest changes

One thing that caught me off guard is the fact that some constructs that we all know and love are missing.
I'm mainly talking about `__dirname`, `module`, `__filename`, `exports` and `require`.

Thankfully, there is a way to "get them back", the syntax is just different.

### Missing `__dirname` and `__filename`

```js
import { fileURLToPath } from "url";
import { dirname } from "path";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

console.log(__dirname);
console.log(__filename);
```

The most interesting thing for me here is that we are leveraging a property pulled from the `import` object.
This is something I would not expect to work as I'm very much used to thinking about `import` as a keyword not an object.

### Top-level await

Recall all the situations where you had to wrap your function within an IFFIE because it was async.
While completely valid, would it be nice if we do not have to do it that way?

Well, within the new ESM enabled Node.JS environment you no longer have to do that.

```js
import { promises as fs } from "fs";

// Look ma, no async function wrapper!
console.log(JSON.parse(await fs.readFile("./package.json")).type);
```

So neat!

### Missing require - how to import `json` files

> I really suggest you get familiar with the new `module` API. You can [view it here](https://nodejs.org/api/module.html#module_module_createrequire_filename)

Since the `require` is gone, how to do we import `json` files?
It turns out the `require` is not completely gone, you just need to construct it yourself.

```js
import { createRequire } from "module";
const require = createRequire(import.meta.url);
```

Personally, I'm not sure I love the way we let the `require` statement still be present in the "new world" of ESM modules. While I completely understand that having it will make the migration easier, why not enable importing `json` files through `import` directly?

It turns out that is on the horizon, [there is the JSON modules flag that enables you to do that](https://nodejs.org/api/esm.html#esm_experimental_json_modules).

### The ability to import from URL

Native in browsers, **experimental in NodeJs**. If you use the `--experimental-loader` flag, you **can import the modules from a remote source, like a `HTTP` endpoint**. This allows you to create a `RPC` mechanism, like [demonstrated here](https://betterprogramming.pub/http-modular-my-node-js-library-for-converting-server-side-functions-into-es-modules-ac78799899ce).

```js
// Only an example...
import foo from "https://google.com";
foo();
```
