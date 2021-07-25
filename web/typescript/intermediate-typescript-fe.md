# Intermediate TypeScript FrontendMasters

## Declaration merging

- Multiple things stacking into each other

- Notion of things declared in _type space_ and _value space_. You can have types named the same as your JavaScript identifiers.

```ts
interface Fruit {
  name: string;
}

const Fruit = {
  name: "banana",
};

export { Fruit }; // Both type and JavaScript value
```

- This "overloading" for a given identifier also works with _namespaces_

```ts
interface Fruit {
  name: string;
}

namespace Fruit {
  createFruit(): Fruit {
    return {name: "banana"}
  }
}

const Fruit = {
  name: "banana",
};

export { Fruit }; // Both type and JavaScript value
```

- Please note that **unless explicitly silenced, the above will result in TS errors**.

- You can test whether an identifier is referring to a type or a JavaScript value by using an assignment operator.
  You cannot assign a JavaScript value to a type.

- One notable JavaScript construct is a class. It turns out that the **classes are considered to be both types and values**

```ts
class Fruit {
  public name?: string;
  static getFruit() {
    return { name: "banana" };
  }
}

const fruitClass = Fruit; // no error, it means that the Fruit is a value

const instance: Fruit = {} as any; // no error, it means that the Fruit is a type
```

## Modules and CJS

- Due to historical reasons, a lot of code out there is using CJS. The format Node.js uses but slowly deprecates to ESM (thank god).

### CJS Interop

- Usually you can import CJS using the `import * as foo from require("bar")` syntax which would be equivalent to `const foo = require("bar")`

- Sometimes it's not possible. This might be the case whenever you are not importing a namespace (like `fs`) but a single exported function / variable.
  As an example, one might do `module.exports = function createBanana(){}`. **The exported value is a function, but the `import * as ...` syntax demands the exported variable to be an object**.

- There are two ways you can make the `module.exports = function createBanana(){}` work with TypeScript:

  1. Use the `import createBanana = require("./create-banana")` syntax (not standard but will work)
  2. Use the `esModuleInterop` compiler flag

  Usually, people go for the latter option. While it seems like it's the easiest route to take, it also has it's own drawbacks.
  Mainly - **whenever you set `esModuleInterop` compiler flag, people consuming your code will also have to turn that flag on**.
  Mike calls it a "viral option".

### Namespace imports and bundlers

- The namespace style import (`import * as foo from 'bar'`) is tempting, but if you are worried about the _bundle size_ of your application, you might want to think twice before using this syntax within your code.

- The problem with namespace style imports is that, **usually**, bundlers will eagerly import everything from the file you are referencing.
  The tree shaking might not occur.

## Type Queries

- A way of obtaining a type from a _JavaScript value_

### The `keyof` operator
