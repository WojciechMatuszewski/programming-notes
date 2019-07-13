# Functional Programming with Professor Frisby

## Your first `Functor` (or maybe even Monad? 🤔)

So lets say you want to transform some string.

```js
const nextCharForNumberString = str => {
  const trimmed = str.trim();
  const number = parseInt(trimmed);
  const nextNumber = number + 1;
  return String.fromCharCode(nextNumber);
};

const result = nextCharForNumberString(' 64 ');

console.log(result);
```

This seems all good and would probably make through CR (not mine tho!).
There is one problem with this function. **It's full of side-effects!**.

Let's try to avoid side-effects using composable operations like `map`

```js
const nextCharForNumberString = str =>
  // putting stuff into `container`
  [str]
    .map(s => s.trim())
    .map(s => parseInt(s))
    .map(i => i + 1)
    .map(i => String.fromCharCode(i));

console.log(result[0]);
```

Much better now right? You might say work done, let's grab a coffee and go home. But we are not done here.

Turns out that `container` is a data type called `Functor`. To be exact **functor is a structure that you can `.map` over**

```js
// this is going to be our functor
const Box = x => ({
  map: f => Box(f(x)),
  // used for extracting values
  fold: f => f(x)
});

const nextCharForNumberString = str =>
  // putting stuff into `container`
  Box(str)
    .map(s => s.trim())
    .map(s => parseInt(s))
    .map(i => i + 1)
    .fold(i => String.fromCharCode(i));
```

Now we have algebraic structure to work with. **Your first functor**👍

## Either type

`Either` type is defined as `Left` or `Right`.

## Multiple Containers

You feel really powerful with your `Functor` but you encountered a problem.

```js
const applyDiscount = (price, discount) =>
  // 1 functor deep
  moneyToFloat(price).map(cost =>
    // 2 functors deep
    percentToFloat(discount).map(savings => cost - cost * savings)
  );
// Box(Box(4))
console.log(applyDiscount('$5.00', '20%'));
```

Oh no! You have **a Box within a Box**. That is bad. You might say _ok Wojtek, I'm just going to fold twice instead of mapping_

```js
const applyDiscount = (price, discount) =>
  // 1 functor deep
  moneyToFloat(price).fold(cost =>
    // 2 functors deep
    percentToFloat(discount).fold(savings => cost - cost * savings)
  );
```

But that's not really the solution. It just feels bad.
