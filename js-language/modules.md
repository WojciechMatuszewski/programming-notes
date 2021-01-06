# Modules

It seems to me like we are finally going to see `ESM` being used everywhere. The days of `CJS` are pretty much over given the fact that you can use them while working with Node without much hassle.

One thing that `ESM` brings to the table, which was not obvious to me at all, is the ability to have _cyclic dependency imports_. I do not want to discuss whether this is good or a bad thing. I'm more interested in the tech that enables us to do that.

> Looking back, It never occurred to me that I'm creating a cyclic dependency, especially in big codebases. Well, now I know how it's even possible.

## ESM _live bindings_

It is the concept of _live bindings_ that allows us to create _cyclic dependency imports_. There is this [great article on the subject matter](https://hacks.mozilla.org/2018/03/es-modules-a-cartoon-deep-dive/).

The high level overview is that, whenever you `export` something, the `import` for that `export` points to the same memory. I personally view it like a pointer.

Since we are talking _pointers_ now, the one who holds the reference can observe changes to it. So in the case of _cyclic dependency_ all the variables and values will _eventually_ be resolved.

```js
// index.js

import { count } from "./counter.js";

console.log(count);

export const message = "foo";

// counter.js
import { message } from "./index.js";

// Wait for module resolution to be over
setTimeout(() => {
  console.log(message);
}, 0);

export const count = 5;
```

You would not be able to run this code using `CJS`. This is because the `require` statement just copies over the `export` object at the time of resolution.
