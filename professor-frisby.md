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
    .map(i => i + 1)s
    .fold(i => String.fromCharCode(i));
```

Now we have algebraic structure to work with. **Your first functor**👍

## Either type

`Either` type is defined as `Left` or `Right`.

```js
const Right = x => ({
  map: f => Right(f(x)),
  fold: (f, g) => g(x)
});

const Left = x => ({
  map: f => Left(x),
  // first function is treated as error callback
  fold: (f, g) => f(x)
});
```

This allows us to branch out our code.
Let's say you have a method that can return `null` or `undefined` but you still want to _dot chain_ operations on returned data.

```js
// can return undefined
function getColor(name) {
  return { green: 'a', blue: 'b', yellow: 'c' }[name];
}

// protecting against null or undefined
function fromNullable(x) {
  return x != null ? Right(x) : Left(x);
}

fromNullable(getColor('someColor'))
  .map(/*some kind of operation*/)
  .fold(
    e => console.log('undefined or null, all maps ignored'),
    data => console.log(data)
  );
```

### Another example

Let's say you want to read from file (using `fs` module) but you also want to guard against errors

```js
const fs = require('fs');

const getPort = () => {
  try {
    const str = fs.readFileSync('config.json');
    const config = JSON.parse(str);
    return config.port;
  } catch (e) {
    return 3000;
  }
};

const result = getPort();
```

The code above **is OK!**, but it's imperative and could be improved.
Let's use our previous knowledge to solve fight this imperative dragon.

```js
// this is our only try/catch
function tryCatch(f) {
  try {
    Right(f());
  } catch (e) {
    Left(e);
  }
}

const getPort = () =>
  tryCatch(fs.readFileSync('config.json'))
    // but hey this returns Right(Left) or Left(Right) or Right(Left)....
    .map(config => tryCatch(JSON.parse(config)))
    .fold(e => 3000, config => config.port);
```

Yea there is one problem, and that problem has to do with **multiple container types**. How to solve that problem?

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

### Chain method

Chain method allows us to **escape nested containers**.

```js
// defined inside Monad
function chain(f) {
  // x is from closure
  return f(x);
}
```

Now everywhere we have nested containers we can use chain to make everything _flat_ again :).

## Semi-groups

> Semi group is a type with a `.concat` method that is associative

Very basic example of such type is `Sum` type.

```js
const Sum = x => ({
  x,
  concat: ({ x: y }) => Sum(x + y)
});
```

Using this would look something like this:

```js
Sum(3).concat(Sum(3)); // and so one
```

You might find it useless now , but later on this type will become helpful.

## Monoids

> Monoid is a semi-group with neutral element

Ok, but what does that really mean?
Well let's promote our `Sum` _semi-group_ to _monoid_

```js
// ..
// previous implementation
Sum.empty = () => Sum(0);
```

We only had to add `Sum.empty` and now he have _monoid_.
But what are benefits?

Well, monoids ensure our **combinations are failsafe**. That is even without _starting value_ we can _dot chain_ away :).

## Lazy Boxes

Delayed computation is great. It allows for nice optimizations.
Using the power of Monad types we can achieve delayed computation pretty easily.

```js
const LazyBox = (g = {
  map: f => LazyBox(() => f(g())),
  fold: f => f(g())
});
```

With this implementation, as long as we are `.map`'ing the functions will not run. **Using fold is like pulling a trigger**. It makes everything run.

```js
const result = LazyBox(() => ' 64 ')
  .map(abba => abba.trim())
  .map(trimmed => Number(trimmed))
  .map(number => number + 1);
  // composition of functions , nothing ran.
  console.log(result);

  // imagine console.log is not there

  // now we run stuff
  .fold(/* */)
```

## Using Task (folktale)

Task allows us to do _lazy computation_, just like `Lazy Boxes`.
Let's say we want to read / write stuff from/to file. We want to do this in a declarative manner since we are hooked on functional programming.

```js
// firs, we have to change readFile and writeFile implementations
const readFile = (filename, encoding) =>
  new Task(
    (reject, res) => fs.readFile(filename, encoding),
    (err, contents) => (err ? reject(err) : res(contents))
  );

const writeFile = (filename, contents) =>
  new Task(
    (reject, res) => fs.writeFile(filename, contents),
    (err, success) => (err ? reject(err) : res(success))
  );

const app = fs
  // returns Task
  .readFile('config.json', 'utf-8')
  .map(contents => contents.replace(/8/g, '6'));
  // chaining here so that we do not return Task(Task)
  .chain(contents => writeFile('config1.json', contents))

  // now we can fork it
  app.fork(/* error case */, /* success case */)
```

## Lifting value into a type

Remember `of` and `from` methods from RxJs library?
There is also a notion of `lifting` when writing custom operators.
Well know it all comes full circle.
`of` allows you to, well place given value into given type so you can start using properties of that type (either Monadic or Functor stuff).

```js
Box.of(3).map(...)
Either.of(true).map(...)
Task.of('something').fold(...)
```

## You've been learning about Monads

> Monad is `Functor` with a `.of` and `.chain` method (sometimes called flatMap, bind , >>=)
