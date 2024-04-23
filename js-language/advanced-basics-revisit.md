# Advanced basics revisited

## Primitive types

- there are primitives types (7 kinds) in Javascript
- the _every thing is an object_ is a meme, not true at all

## Boxing

- this is the thing that allows you to use `.` notation on primitive types:

```js
"1,2,3".split(",");
```

- when you are using `.`, the engine will automatically box the primitives for you

## `__proto__`

- used in inheritance
- by default, points to the global `Object`
- you can create `__proto__` chains using `Object.create`

```js
const a = {};
console.log(a.__proto__); // global Object

const b = Object.create(a, { speak: () => "woof" });
console.log(b.__proto__); // {speak: () => "woof"} and then global Object (__proto__.__proto__)
```

## `.prototype`

- **arrow functions does not have the `.prototype`**
- automatically created for functions (the normal ones)
- **not used in inheritance**

- when instance is created with `new`, that `.prototype` is the top most `__.proto__`.

```js
function someFunc() {}

someFunc.prototype.speak = () => "woof";
console.log(someFunc.__proto__); // global Object

const instance = new someFunc();
console.log(instance.__proto__); // {speak: () => "woof"} and then global Object (__proto__.__proto__)
```

## `var` is function-scoped based

- very simple concept. With the `block` scope:

```js
var name = "Wojtek";
{
  var name = "Mateusz";
  console.log(name); // Mateusz
}
console.log(name); // Mateusz
```

- with the function scope:

```js
var name = "Wojtek";
function someFunc() {
  var name = "Mateusz";
  console.log(name); // Mateusz
}
console.log(name); // Wojtek
```

## Arrow function and lexical scope

- arrow function always looks up to the next non-arrow function for it's scope
- this means that you might end up with a global scope.
- the scope is **determined at author time**. The declaration place matters.

```js
const Person = {
  firstName: "Wojtek",
  getName: () => this.firstName,
};
console.log(Person.getName); // undefined

const Person2 = {
  firstName: "Wojtek",
  getName: function () {
    // placed inside a function with a context
    return () => this.firstName;
  },
};
console.log(Person2.getName()()); // Wojtek
```
