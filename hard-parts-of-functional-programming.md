# Hard parts of functional programming

##### This material is from frontend masters workshop

- Tiny functions
- No consequences except within a given function (no side effects)
- Recombine/compose - build up our application by using small function blocks

There are many techniques for combining or creating building blocks, mainly:

- function composition
- higher order functions
- currying

## Higher order functions

Suppose we have a function `copyArrayAndMultiplyBy2`

```js
function copyArrayAndMultiplyBy2(array) {
  const output = [];
  for (let i = 0; i < array.length; i++) {
    output.push(array[i] * 2);
  }
  return output;
}
const myArray = [1, 2, 3];
const result = copyArrayAndMultiplyBy2(myArray);
```

There is one glaring problem about this implementation:

- it's true that the name is very specific but the function itself is not flexible at all (we would have to create another one if we wanted to multiply by another number)

To make the function more flexible we can pass `instructions` (basically what we want to do with given array item).

```js
function instructions(input) {
  // manipulate input
  // return manipulatedInput
}
function copyArrayAndManipulate(array, instructions) {
  // ...
}
```

### Functions as first class objects

How is it that we can pass functions around? We can do that because they are _first class objects_ (to be more precise they are just callable objects). That we can do stuff with them as if they were objects.

Looking back at the example, which of these functions is the `higher order function` and which is the `callback` function?

- `copyArrayAndManipulate` is the _higher order function_ because it takes a function as a parameter

- `instructions` is the _callback function_ since it's passed as an argument.

Granted to be _higher order function_ you do not have to take function as an parameter. Your function could also return a function or do both of the things at once (return and take as a parameter).

## Composing functions

As an alternative (in my opinion much better one) to chaining we can compose functions together. This not only compose multiple functions into one but also allows us to gain some, miniscule but still, performance improvements.

### Composing using reduce

Reduce is very versatile, in fact its the most versatile function in fp.

Below style of composition is still **not** a _true composition_ as in it does not build the one meta function from the functions but instead it steps through every function passing down the result.

```js
const multiplyBy3 = x => x * 2;
const add3 = x => x + 3;
const reducer = (input, fn) => fn(input);
[multiplyBy3, add3].reduce(reducer, 11);
```

To reap every benefit of composition and lazy evaluation we have would have to glue functions together then run the final "glued" function against values.

```js
function compose(...fns) {
  // 'glue' functions together
  return fns.reduce((f, g) => (...args) => f(g(...arg)));
}
compose(
  fn1,
  fn2,
  fn3
)(/*some value*/);
```
