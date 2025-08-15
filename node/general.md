# General

## The native `.env` file support

- In v20.6.0 Node.js added [native `.env` file support](https://nodejs.org/en/blog/release/v20.6.0).

- Usage via the `--env-file` flag, like so: `node foo.js --env-file=config.env`

  - **You can define any `NODE_OPTION` in the `.env` file as well!** Pretty neat stuff.

- This **might make the `dotenv` package obsolete**.

## The `node:` import prefix

When reading code, you might come across import statements that look like:

```ts
import fs from "node:fs";
import { readFile } from "node:fs/promises";
```

**The `node:` prefix is a convention to signal that you import something built-in to node rather than an external package**.

AFAIK, the benefit is purely related to semantics rather than functionality. Using this convention allows Node.js to introduce new packages which won't conflict with existing, very popular ones, due to that prefixing.

Some packages, like the `node:test` are NOT available via "bare" import without the prefix.
