# Functional-light JS with Kyle Simpson

## Imperative vs Declarative

### What does **Imperative** mean?

- code that is focused, primarily, on **how to do something**.
- future reader of the code, has to read all the code and, in a sense, mentally execute it (imply from the code what the code is doing)

Computer is pretty god at executing a code, not our brains.
Any time we force a person to 'execute' a code inside their minds to enable them to see how it's working we are dealing with a code that is hard to understand.

### What does **Declarative** mean?

- code that is focused, primarily on **the outcome, the why.**

## Functions

Functions, by default (in JavaScript) return `undefined`

### What is a function really

```javascript
// is this really a function?
// just because it uses a *function* keyword does not make it a function
// Kyle calls it procedure
function addNumbers(x = 0, y = 0, z = 0, w = 0) {
    var total = x + y + z + w;
    console.log(total);
}

// it definitely has return keyword
// is it a function?
// not really :C
function extraNumbers(x = 2, ...args) {
    return addNumbers(x, 40, ...args);
}
```

Why does Kyle calls `addNumbers` a _procedure_ instead of a _function_ ?

- **function has, not only, take some input but also return some output**

Is `extraNumbers` a function then?

- **functions can only call other functions**

These definitions are not complete, they will be improved upon as we progress :)

### True functions

The example below presents a true, in spirit, function.

```javascript
function tuple(x, y) {
    return [x + 1, y - 1];
}

var [a, b] = tuple(...[5, 10]);
a; // 6
b; // 9
```

A function is a **relationship** between input and output.

```javascript
f(x) = 2x^2 + 3

// we could write it like so
function f(x){
    // there is clear relationship between inputs and the outputs
    return 2* Math.pow(x,2) + 3
}

// another example
function shippingRate(size, weight, speed) {
    return ((size + 1) * weight ) + speed
}
```

So it seems that for it to be a function it cannot have any **side effects**.

But what does that really mean?

### Side Effects

```javascript
function shippingRate() {
    rate = (size + 1) * weight + speed;
}
var rate;
var size = 12;
var weight = 4;
var speed = 5;
shippingRate();
rate; // 57
```

The above code works, but inputs and the outputs are indirect (even though there is semantic relationship between them).

So it seems that function should

- take direct inputs (arguments passed to parameters)
- compute and return a value without assigning or accessing anything outside itself

So what are a **side effects** really (not a complete list) ?

- I/O (console, files, etc)
- Database Storage
- Network Calls
- DOM
- Timestamps
- Random Number (generation)
- _CPU Heat_
- _CPU Time Delay_

As you can see a program without **side effects** could not exist.
We sometimes have to do **side effects**. But with doing so we should **make them obvious for the reader of the code**.

## Pure Functions

_Pure Function_ is a function that obey the previously stated terms,
and **has no side effects**.

```javascript
// pure
function addTwo(x, y) {
    return x + y;
}

// impure
function addAnother(x, y) {
    // accessing z is an side effect!
    return addTwo(x, y) + z;
}
```

But does adding `z` really invalidates the function?

```javascript
const z = 1;
function addTwo(x, y) {
    return x + y;
}

// this is pure function
function addAnotherPure(x, y) {
    return addTwo(x, y) + 1;
}
// this, well by definition is impure but is it really
// it does not cause any side effects
// it does not use any side effects other that constant z
function addAnother(x, y) {
    return addTwo(x, y) + z;
}
```

Kyle argues here that `addAnother` is not really impure. Since the semantics of `const z = 1` tell us that it will never change and since it's kind of a 'placeholder' for a value we should treat it as `addAnotherPure`

```javascript
function addAnother(z) {
    return function addTwo(x, y) {
        return x + y + z;
    };
}
addAnother(1)(20, 21);
```

In above example, we reduced the 'surface area' of a program from 11 lines that could modify `const/var z` to only 2 : `1.5 and 2.5`.
That's how we should increase readability and confidence in our programs.

What about predictability ?

- **pure function should always return the same output given the same input**

### Extracting Impurity

Sometimes we cannot avoid impurity. We should **extract impurity** to make tem more obvious.

```javascript
function addComment(userId, comment) {
  var record = {
    // side effect
    id: uniqueID(),
    ...
  }
  var elem = buildCommentElement(record);
  // side effect
  commentsList.appendChild(elem);
}
// impure call
addComment(42, "...")
```

Is it possible to extract impurity of the above snippet?

```javascript
function newComment(userID, commentID, comment) {
  var record = {
    id: commentID,
    ...
  }
  return buildCommentElement(record);
}

var commentID = uniqueID();
var elem = newComment(....)
commentsList.appendChild(elem);
```

We extracted impurity but it's not really a great solution.
We polluted global scope with impurity.
We should **contain impurity** along with extracting it.

### Containing Impurity

#### Wrapping side effects inside other functions

```javascript
var SomeAPI = {
  threshold: 13,
  isBelowThreshold(x) {
    return x <= this.threshold;
  }
};

var numbers = [];

function insertSortedDesc(v) {
  SomeAPI.threshold = v;
  ...
  idx = ...
  ...
  // mutate array
  // this side effect affects global scope
  numbers.splice(idx, 0, v);
}
```

And the _contained_ version

```javascript
var SomeAPI = ....
var numbers = [];

// containing side effect inside this function
function getSortedNums(nums,v) {
  // create a copy of an array
  var numbers = nums.slice();
  insertSortedDesc(v);
  return numbers;

  function insertSortedDesc(v) {/* same as above */}
}
```

Now we contained side effect inside `getSortedNums`. This does not occur very frequently but it's a good technique to know.

But still we are still modifying `SomeAPI.threshold` inside `getSortedNums`

There is one technique we can use as an escape hatch. It's not pretty but it's better than nothing.

```javascript
var SomeAPI = ....
var numbers = [];
function insertSortedDesc(v) {
  /* same as above */
}
function getSortedNums(nums,v) {
  // make a copy of 'current state'
  var [originalNumbers, originalThreshold] = [numbers, SomeAPI.threshold];
  numbers = nums.slice();
  // side effects occur here!
  insertSortedDesc(v);
  // 'capture' new changed state
  nums = numbers;
  // restore state to it's original values
  [numbers, SomeAPI.threshold] = [originalNumbers, originalThreshold];
  return nums;
}
```

This is not pretty, you should not really do this. I would advise doing it as last resort.

## Arguments

- **parameter is the thing inside the function definition**

```javascript
// x,y are parameters
function addNumber(x,y) {
  ...
}
```

- **argument is the thing that is passed to function parameters when calling given function**

```javascript
// 3,4 are the arguments
addNumber(3, 4);
```

### Unrary & Binary

There is a notion of a _shape_ when talking about functions.

We can define _shape_ of a given function as:

- the number and a kinds of things you pass into it
- the number and a kinds of things that come out of it

### Unrary shape

Unrary function takes a single value

```javascript
function increment(x) {
    return x + 1;
}
```

### Binary shape

Binary shape takes 2 parameters

```javascript
function sum(x,y){...}
```

## Higher order functions

_higher order function_ is a function which

- either receives as inputs one or more functions and / or returns one or more functions

```javascript
function unary(fn) {
    return function one(arg) {
        return fn(arg);
    };
}
```

## Point-Free

This is a _style_ of writing functions.
What does that even mean?

- defining a function without the need to define it's inputs (points)

```javascript
// person parameter here is the 'point'
getPerson(function onePerson(person) {
    return renderPerson(person);
});

// so instead we can do
getPerson(renderPerson);
```

### Using Point-Free style with Higher Order Functions

```javascript
function isOdd(v) {
    return v % 2 == 1;
}
function isEven(v) {
    return !isOdd(v);
}

// but we can improve this solution with
function not(fn) {
    return function negated(...args) {
        return !fn(...args);
    };
}

var isEven = not(isOdd);
```

### Advanced Point-Free

```javascript
function mod(y) {
    return function forX(x) {
        return x % y;
    };
}

function eq(y) {
    return function forX(x) {
        return x === y;
    };
}

/// How would we defined isOdd using point-free
var mod2 = mod(2);
var eq1 = eq(1);

// 1 step
function isOdd(x) {
    return eq1(mod2(x));
}

// to make it completely point-free we can use composition
var isOdd = compose(
    eq1,
    mod2,
);
```

## Closure

> Closure is when a function "remembers" the variables around it even when that function is executed elsewhere

Example

```javascript
function makeCounter() {
    var counter = 0;
    // increment is closed over counter variable
    return function increment() {
        return ++counter;
    };
}
var c = makeCounter();
// these are NOT pure function calls!
c();
c();
c();
```

### Lazy vs Eager

#### Lazy Computation

Given below snippet, when does the actual work of constructing the string occurs?

```javascript
function repeater(count) {
    return function allTheAs() {
        return "".padStart(count, "A");
    };
}

var A = repeater(10);
A(); // "AAAAAAAAAA"
A();
```

Is it on line `7` or line `9`?
Its on line `9`.
We deferred the work from line `7` to line `9`. That is called `lazy computation`

#### Eager Computation

```javascript
function repeater(count) {
    var str = "".padStart(count, "A");
    // allTheAs is closing over str variable
    return function allTheAs() {
        return str;
    };
}
var A = repeater(10);
A();
A();
```

Now we moved all the work to line `2` Now works occurs on line `8`. We are doing the work only once.

#### Combining Eager and Lazy

What if we could combine the best of both words?

```javascript
function repeater(count) {
    var str;
    return function allTheAs() {
        if (str == undefined) {
            str = "".padStart(count, "A");
        }
        return str;
    };
}

var A = repeater(10);
// we deferred work
A(); // we are doing the work only here
A(); // now we are pulling for existing variable
```

Before we said that to be functionally pure when using closure we **should not close over the thing that changes**.
`A` is a pure function call, but the code is not obvious.

#### Using memoization

```javascript
function repeater(count) {
    return memoize(function allTheAs() {
        return "".padStart(count, "A");
    });
}
var A = repeater(10);
A();
A();
```

This style of code is very obvious for the reader.

## Pure function the complete definition (function call)

> If you could take the return value of that function call and replace the function call with the returned value and not affect the rest of the program you have **pure function call**.

In other ways the function call is pure when it has **referential transparency**.

## Generalized to Specialized

```javascript
function ajax(url, data, cb) {...}
// we are passing a lot of arguments here
ajax(CUSTOMER_API, {id: 42}, renderCustomer)
```

Lets consider some intermediately steps.

```javascript
function ajax(url, data,cb) {...}
ajax(CUSTOMER_API, {id:32}, renderCustomer)

function getCustomer(data, cb) {
  return ajax(CUSTOMER_API, data,cb);
}
// this function is much better than the ajax function
getCustomer({id:32}, renderCustomer)
```

We just split generalized function to more specialized sub-function.
This technique allows for better semantics and code readability.

Can we somehow made the specialized versions more _point-free_?

### Partial Application

```javascript
function ajax(url, data, cb) {...}
// pre-set functions
var getCustomer = partial(ajax, CUSTOMER_API);
var getCurrentUser = partial(getCustomer, {id: 42});
// much better right ?
getCustomer({id:42}, renderCustomer);
getCurrentUser(renderCustomer);
```

### Currying

Much more common form of specialization.
_currying_ and _partial application_ both accomplish the same goal, they both specialize a generalized function. But they do it differently

```javascript
// manual curry
function ajax(url) {
  return function getData(data) {
    return function getCB(cb){..}
  }
}
ajax(CUSTOMER_API)({id: 42})(renderCustomer);
var getCustomer = ajax(CUSTOMER_API);
var getCurrentUser = getCustomer({id: 42});
```

#### ES6 curry

```javascript
function curry(fn) {
  return function curried(...args) {
    if (fn.length <= args.length) {
      return fn.apply(null, args)
    }
    return function curried.bind(null, ...args)
  }
}
```

#### ES5 curry

This one is a bit harder since we cannot use spread and gather

```javascript
function curry(fn) {
    return function curried() {
        var args = Array.prototype.slice.call(arguments);
        if (fn.length <= args.length) {
            return fn.apply(null, args);
        }
        // here we are appending previous arguments to currently passed ones
        // we basically did this with .bind before
        return function partiallyApplyCurriedArguments() {
            return curried.apply(
                null,
                args.concat(Array.prototype.slice.call(arguments)),
            );
            // You could also if you really want use push here
            // make a copy so not to mutate args variable
            var argsToApply = args.slice();
            Array.prototype.push.apply(
                argsToApply,
                Array.prototype.slice.call(arguments),
            );
            return curries.apply(null, argsToApply);
        };
    };
}
```

### Partial Application vs Currying (strict)

- both are specialization techniques
- _partial application_ presets some arguments now, receives the rest on the next call
- _currying_ does not preset any arguments, receives each argument one at a time.

## Composition

Composition is when one function takes other function output as input.
Lets consider this not very functional example:

```javascript
function minus2(x) {
    return x - 2;
}
function triple(x) {
    return x * 3;
}
function increment(x) {
    return x + 1;
}

var tmp = increment(4);
tmp = triple(tmp);
totalCost = basePrice + minus2(tmp);
```

How we can make it more functional?

Lets just get rid of `tmp` variable. That should help right?

```javascript
... functions from above ...
totalCost = basePrice + minus2(triple(increment(4)));
```

Technically this is called `composition` but still looks ugly (and hard to read).
Soo maybe let's abstract the ugly part to a function?

```javascript
... functions from above ...
function shippingRate(x) {
  return minus2(triple(increment(4)));
}
totalCost = basePRice + shippingRate(4)
```

This solution is like putting all your dirty clothes into a closet, not good (I know people who do such things 😂)

We can make a function which makes us a function composed of different functions. Let's try that!

```javascript
function composeThree(fn3, fn2, fn1) {
    return function composed(v) {
        return fn3(fn2(fn1(v)));
    };
}
```

Now we can define as many shipping rates as we want.

```javascript
... functions from above ...
// also look! point free function
var shippingRate = composeThree(minus2, triple, increment);
// other shipping rates can be easily defined
totalCost = basePRice + shippingRate(4)
```

**composition is RIGHT TO LEFT!!**
**pipe is LEFT TO RIGHT!!**

### Combining Currying with Composition

We only should compose unary functions, otherwise it would be very hard to make all the functions compatible shape-wise

```javascript
function sum(x, y) {
    return x + y;
}
function triple(x) {
    return x * 3;
}
function divBy(y, x) {
    return x / y;
}

divBy(2, triple(sum(3, 5))); // 12
// lets use curry to solve this problem
sum = curry(2, sum);
divBy = curry(2, divBy);
composeThree(divBy(2), triple, sum(3))(5); // 12
// much better & cleaner 👌
```

## Lists (data structures)

### Map: transformation

- does not mutate original data structure
- values are _projected_ onto new values
- amount of values stays the same, data structure also stays the same

Simple implementation
(There is also built-in implementation on `Array.prototype`)

```js
function map(mapper, arr) {
  var newList = [];
  for (let value of arr) {
    newList.push(mapper(value));
  }
  return newList;
}
```

### Filter: exclusion (or maybe actually inclusion?)

- does not mutate original data structure
- return `true` if you want to keep a value and `false` if not
- amount of values can change, data structure stays the same

Simple implementation
(There is also built-in implementation on `Array.prototype`)

```js
function filter(predicate, arr) {
    var newList = [];
    for (let value of arr) {
        if (predicate(elem)) {
            newList.push(elem);
        }
    }
    return newList;
}
```

### Reduce: combining

- you can implement `.filter` or/and `.map` with reduce (used with transducing)
- data structure can change
- amount of values can change
- starts with initial value
- be very careful not to mutate stuff

Built-in implementation has many different overloads (much more complex than this)

```js
function reduce(reducer, initialVal, arr) {
  var ret = initialVal;
  for (let elem of arr) {
    ret = reducer(ret, elem);
  }
  return ret;
}
```

**`.reduce` goes left to right, `.reduceRight` goes right to left**

### Composition Revisited

How can we use `.reduce` to implement `compose` better?

```js
// this has the same shape as reducer!
function composeTwo(fn2, f1) {
  return function composed(v) {
    return fn2(fn1(v));
  };
}
// look how sexy it is 😻
var f = [div3, mul2, add1].reduce(composeTwo);
```

Composition allows us to _combine operations_ (instead of traversing array multiple times we can do it only once)

## Composing with different shapes (transducing)

It's all nice and easy when function's has the same shape. But what about about combining map & filter & reduce together?

Transducing is a _composition of reducers_

High-level API example

```js
var transducer = compose(
    // these functions might look scary but they are not that hard to write
    mapReducer(add1),
    filterReducer(isOdd),
);
transduce(transducer, sum, 0, [1, 2, 3, 4]);
// or
into(transducer, 0, [1, 2, 3, 4]);
```

### Deriving transduction

Writing `mapReducer` and `filterReducer`

```js
var mapReducer = curry(2, function mapReducer(mappingFn, combineFn) {
    return function reducer(currentAccumulatorValue, v) {
        return combineFn(currentAccumulatorValue, mappingFn(v));
    };
});

var filterReducer = curry(2, function filterReducer(predicateFn, combineFn) {
    return function reducer(currentAccumulatorValue, v) {
        if (predicateFn(v)) return combineFn(currentAccumulatorValue, v);
        return list;
    };
});

var transducer = compose(
    mapReducer(add1),
    filterReducer(isOdd),
);
// now we can use 1 reduce

list.reduce(transducer(sum), 0);
```

- `sum` combiner goes to filterReducer
- it creates a reducer which can call sum if it matches `predicateFn`
- then that reducer goes into `mapReducer` as `combineFn`

So when we are calling `combineFn` inside `mapReducer` we are, in reality, calling `filterReducer` with `combineFn` set to `sum`!

## FP data structures

Any value that we can map-over is a `functor`

### Monad

Monad is a pattern for pairing data with a set of predictable behaviors that ley it interact with other data + behavior pairings (monads). There is a lot of different ways to code a `monad data structure` in JS. You probably you want to use library for this.

Calling `.map` on a `monad` will give you back a `monad`

#### Just - wrapper around a single value

That value can be array or object or anything pretty much.
Very basic implementation of `Just monad`:

```js
function Just(val) {
    return { map, chain, ap };

    function map(fn) {
        return Just(fn(val));
    }

    // aka: bind, flatMap
    function chain(fn) {
        return fn(val);
    }

    function ap(anotherMonad) {
        return anotherMonad.map(val);
    }
}
```

Example operations

```js
var fortyOne = Just(41);
var fortyTwo = fortyOne.map(function inc(v) {
  return v + 1;
});

function identity(v) {
  return v;
}

// very loose implementation, normally you should not return value from chain
// once it's in the monad it has to stay there
fortyOne.chain(identity); // 41
fortyOne.chain(identity); // 42
```

`.ap` function

```js
var user1 = Just("kyle");
var user2 = Just("susan");
var tuple = curry(2, function tuple(x, y) {
    return [x, y];
});
var users = user1.map(tuple).ap(user2);

users.chain(identity); // [kyle, susan]
```

Monad we've made with `user1.map(tuple)` is a monad which value is a **function** which waits for another argument (in this case _y_)

Result of `.ap(user2)` is a monad with a value returned from `tuple` function

#### Maybe - safe operations

```js
var someObj = { something: { else: { entirely: 42 } } };
```

Works great but what if `someObj.something` was `undefined` or `null`?
(or any other prop on that object for that matter).

```js
// black hole of monads. Every method returns Nothing
function Nothing() {
  return { map: Nothing, chain: Nothing, ap: Nothing };
}

var Maybe = { Just, Nothing, of: Just };

function fromNullable(val) {
  if (val == null) return Maybe.Nothing();
  else return Maybe.of(val);
}

var prop = curry(2, function prop(prop, obj) {
  return fromNullable(obj[prop]);
});
```

Armed with above code we can make safe property access function

```js
Maybe.of(someObj)
    // prop('something') returns a function which will return Just or Nothing
    // chain prevents Monads from nesting
    .chain(prop("something"))
    .chain(prop("else"))
    .chain(prop("entirely"));
```

There are many more monads. Kyle showed only basic ones.

## Async (observables)

This is a very brief mention.

```js
var a = new Rx.Subject();

setInterval(function everySecond() {
    a.next(Math.random());
}, 1000);

var b = a.map(function double(v) {
    return v * 2;
});

b.subscribe(function onValue(v) {
    console.log(v);
});
```
