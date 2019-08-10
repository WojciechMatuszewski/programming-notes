# Mostly Adequate Javascript

This is me taking notes from the book. Many chapters will be skipped and notes will probably be chaotic.

## Chapter 07: Hindley-Milner and Me

So functional paradigm has types. These are used for, well, knowing that's the shape of a given function. If you know `typescript` or `flow` you should be good here.

```js
// capitalize :: String -> String
const capitalize = s => toUpperCase(head(s)) + toLowerCase(tail(s));

capitalize('smurf'); // 'Smurf'
```

Pretty easy stuff right? `name :: TYPES`

We can also do a _generic_ variation of this. Just like in typescript, but the syntax is a little bit different (it's more like a convention)

```js
// identity :: a -> a
const identity = x => x;
```

So here, `a` can be any type.

You can also _group_ types. This is nothing special, just helps with readability.

```js
// match :: Regex -> (String -> [String])
const match = curry((reg, s) => s.match(reg));
```

So the act of matching in this part is the `(...)` syntax.

Taking a function as a parameter is interesting when it comes to HM types.

```js
// map :: (a -> b) -> [a] -> [b]
const map = curry((f, array) => array.map(f));
```

Here, `map` takes a function which _maps type `a`_ to _type `b`_. And since map always return arrays we project that in the type itself.

Typing reduce is fun, check it out:

```js
// reduce :: ((b -> a) -> b) -> b -> [a] -> b
const reduce = curry((reducer, acc, array) => array.reduce(reducer, acc));
```

This may seem cryptic at first but it's not that hard tbh.

- `((b -> a) -> b)` this is our reducer function, `b` is the accumulator and `a` is the next item

* `b -> [a] -> b` `b`, again is our accumulator (look at the curry signature), `[a]` is the array and `b` notes that we return the shape of accumulator. `b` is a generic type as in it could be an array or an object or anything really.

### Constrains

Just like in any type-language instead of using generics or built-in types you can constrain parameters to a given shape (like `interface`)

```js
// sort :: Ord a => [a] -> [a]
```

Syntax is a little bit different. We are using `=>` first then normal `->`.

## Chapter 08: Tupperware

So a container. This will help us:

- manage error handling
- async actions
- state
- effects

All of these are not really possible with pure tiny composable functions.

Lets create a container

```js
class Container {
  constructor(x) {
    this.$value = x;
  }
  static of(x) {
    return new Container(x);
  }
}
Container.of(3);
// Container(3)

Container.of('hotdogs');
// Container("hotdogs")

Container.of(Container.of({ name: 'yoda' }));
// Container(Container({ name: 'yoda' }))
```

So far so good. This is like a structure that holds our data.
One thing that is very important: **once data goes into the container IT STAYS THERE**, this concept is usually hard for people to grasp.

### Functor

So a functor is, MAINLY, structure that you can map over (has `map` method).
Since our container holds value, we can make it a _Functor_

```js
// map :: (a -> b) -> Container a -> Container b
Container.prototype.map = function(f) {
  return Container.of(f(this.$value));
};
```

With types, we are constraining the result to `Container` type.

### Maybe

This concept shows the power of the notion of a container. Basically _container-mentality_ allows us to implement other structures that are somewhat complex.

```js
class Maybe {
  static of(x) {
    return new Maybe(x);
  }

  get isNothing() {
    return this.$value === null || this.$value === undefined;
  }

  constructor(x) {
    this.$value = x;
  }

  map(fn) {
    return this.isNothing ? this : Maybe.of(fn(this.$value));
  }

  inspect() {
    return this.isNothing ? 'Nothing' : `Just(${inspect(this.$value)})`;
  }
}
```

`Maybe` is pretty similar to our previous container. It has added benefit of checking if the value is _falsy_ and handling that case.

```js
Maybe.of({ age: 10 })
  .map(prop('someRandomProp'))
  .map(addTen); // Nothing
```

No errors, it just returns `Nothing` in this case since we have null checking bult-in to our container :).
