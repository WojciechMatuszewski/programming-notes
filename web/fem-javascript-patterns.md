# JavaScript Patterns FEM

## Workshop

[Here is the link](https://frontendmasters.com/workshops/javascript-patterns/)

## Design patterns

### Module Pattern

While it is a common practice now, it did not use to be the case that we split our JavaScripet files into multiple _modules_. From a time perspective, this practice is relatively new.

#### Using a bundler

A modern bundler, like webpack, will wrap the code in your files with IFFIE. That is the OG way of providing encapsulation.

#### In HTML

Apart from having the bundler wrap your code into IFFIE, you could add the `type="module"` to the `script` tag within the HTML.

#### In Node

You can either **add the `"type": "module"`** to your `package.json` or **use the `.mjs`** extension.

One important thing – **it seems like you have to specify the file extension in the import path**. [According to MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import), this is to distinguish between `node_module` modules and your locally defined modules.

## Singleton pattern

We are in luck, as by default, the **browser treats modules as non-modifiable singletons**. That being said, there are a few things you might want to consider when creating a singleton.

1. Consider **freezing** the singleton to avoid "out-of-scope mutations". Here, the `Object.freeze` will do – **remember that the `Object.freeze` is shallow!**.

2. To my best knowledge, you can create singletons via objects of classes. Pick your poison.

3. Since modules are singletons by default, you most likely **do not need to use classes or objects** to create singletons (though it might be a good idea to do so, to be more explicit).

```js
// counter.js

let count = 0;

export function add(x) {
  count += x;
  return count;
}
```

```js
// counter.js

let count = 0;
export const counter = Object.freeze({
  add: function (x) {
    count += x;
    return count;
  },
});
```

```js
// counter.js

export const counter = {
  count: 0
  add: function (x) {
    count += x;
    return count;
  },
};
```

If I were to use `Object.freeze` in the last code snippet, the `count` would always be zero, no matter how many times you call the `add` function. Remember that the "frozen" properties cannot change!

## Proxy pattern

The `Proxy` acts as a middleman between you and a given object. With it, one can know whether the code accesses or changes a given object's property. The `Proxy` is widely used in React as a rendering optimization technique, where libraries only re-render the components that use a given piece of the state.

```js
const person = {
  name: "John Doe",
  age: 42,
  email: "john@doe.com",
  country: "Canada",
};

const personProxy = new Proxy(person, {
  get: (target, prop) => {
    console.log(`The value of ${prop} is ${target[prop]}`);
    return Reflect.get(target, prop);
  },
  set: (target, prop, value) => {
    console.log(`Changed ${prop} from ${target[prop]} to ${value}`);
    return Reflect.set(target, prop, value);
  },
});
```

<!-- TODO: write about the Reflect -->

<!-- Finished part_1, 47:25 -->
