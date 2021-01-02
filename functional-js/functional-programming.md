# Functional programming

## Principles

This section might contain obvious stuff but it's vital to always review the
basics.

### Separate inputs from environment

Secret input: time

```javascript
function daysThisMonth() {
    var date = new Date(),
    y = date.getFullYear()
    ....
}
```

This works always the same. More testable and overall better

```javascript
function daysInMonth(y, m) {
    var start = new Date(y, m - 1, 1),
        end = new Date(y, m, 1);
}
```

### Separating mutation from calculation

```javascript
// this function calculates and mutates stuff, this is not the best
function teaser(size, elt) {
    setText(elt, slice(0, size, text(elt)));
}
map(teaser(50), all("p"));

// this is much better
var teaser = slice(0);
map(
    compose(
        // setText is only one of the operations in chain, easily removable
        setText,
        teaser(50),
        text,
    ),
    all("p"),
);
```

### Recognize pure functions

There are many benefits to pure functions, they are:

- testable
- portable
- memoizable
- parallelizable

### Separate functions from rules

Here there is a bit of talk about currying and how it helps with fp.

## Crucial concepts

### Curry

Curry implementation

```javascript
function curry(fn) {
    return function curried() {
        if (arguments.length >= fn.length) return fn.apply(null, arguments);
        var args = Array.prototype.slice.call(arguments);
        return curried.bind(null, ...args);
    };
}
```

Curry allows us to "feed" a function with one variable at a time. It's very
versatile and you can create your `map` function with it and `reduce` (this will
be helpful later).

```javascript
var newMap = _.curry(function(transformFunction, list) {
  var concatList = function(acc, elt) {
    // push is impure, it does not return the array, thats why we are using concat here
    return acc.concat(transformFunction(elt));
  };
  return _.reduce(concatList, [], list);
});
```

So now with `newMap` we have a 'better' than native map implementation because
its curried which is huge for fp!

So to recap

```javascript
// with currying instead of writing functions like this
var words = function(str) {
  return split(" ", str);
};

// write something like this (curried example)
var words = split(" ");
```

### Compose

Function can meld aka compose Simple compose example

```javascript
// the convention is to read composed function from right to left
function compose(g, f) {
    return function(x) {
        return g(f(x));
    };
}
```

Compose can be used in `map` (curried)

```javascript
var names = _.map(_.compose(get("name"), get("author")));
```

But do not stress about being `point-free` (will talk about this more later).
Sometimes you cannot be 100% `point-free`

```javascript
var isAuthor = function(name, array) {
    return _.compose(_.contains(name), names)(articles);
};
```

And of course there is more than compose. Lets use curry with point-free style

```javascript
var fork = _.curry(function(lastly, f, g, x) {
    return lastly(f(x), g(x));
});
```

Look how easily I can now calculate the average

```javascript
var avg = fork(_.divide, _.sum, _.size);
```

### Point free

Points means arguments, to be a point free means to omit arguments (passed
implicitly)

Example

```javascript
// this code is NOT point free, look how many times we've repeated 'error'
// error is a "glue" variable
function onError(function(error) {
  console.log(error.message)
})
// this is a point free implementation, look how beautiful it looks :)
onError(compose(log, get('message')))
```

Begin point free can help us with one of the hardest things in programming:
Naming functions / variables.

### In review

- make all function inputs explicit as arguments
- these arguments can be provided over time, not just all at once
- try not to modify outside things
- compose without "glue" variables

## Objects (the functional way)

- containers / wrappers for values
- no methods
- not nouns

### Lenses

So apart from being awesome, `lenses` allow you to create safe, reusable getters
and setters where you do not have to worry about null values and such.

With ramda creating a lense is very easy

```js
const XLens = R.lens(R.prop("x"), R.assoc("x"));
```

So there is the `R.lens` where you can take control over how your `lens` works
by passing _getter_ and a _setter_ (these preserve immutability).

There are also shortcut method called `R.lensProp`or a `R.lensPath` which
basically provides a `lens` with a getter and a setter for you, you just have to
specify property name.

There are various methods that you can use with `lenses` like

- `R.view`
- `R.over` => apply a function over a given `lens` returned value
- `R.set`

**`Lenses` also compose nicely with each other!**

### Container convention

```javascript
var _Container = function(val) {
    this.val = val;
};

var Container = function(x) {
    return new _Container(x);
};

Container(3); // _Container {val: 3}
```

So now we have some kind of container 'convention'. Lets try to use our
super-composable functions with it!

```javascript
capitalize("flamethrower"); // Flamethrower

capitalize(Container("flamethrower"));
// => [object Object]
```

Well this did not end up well. How we can solve this problem ? Lets add some
methods to `_Container.prototype`

```javascript
var _Container.prototype.map = function(f) {
  return Container(f(this.val))
}

Container('flamethrower').map(function(s) {
  return capitalize(s)
})
// Container("Flamethrower")
```

So still this seems kinda meh but at least we can use our super-composable
functions on the `Container` for now.

For now think about the container as a singleton array (array with one value)

Container allows us to change types, for example:

```javascript
Container([1, 2, 3])
  .map(reverse)
  .map(first); // Container(3)

Container("flamethrower")
  .map(length)
  .map(add(1)); // Container(13)
```

### Functors

Functor is an object or data structure you can map over

#### Pesky null values

So far we've been really optimistic with our functions. We did not check any
edge cases or apply any null checks. But how to handle them the functional-way?

#### Maybe Functor

- captures a null check
- the value inside may not be there
- sometimes has two subclasses Just / Nothing

Simple example

```javascript
var _Maybe.prototype.map = function(f) {
  return this.val ? Maybe(f(this.val)): Maybe(null)
}

map(capitalize, Maybe('flamethrower'))
// => Maybe("Flamethrower")
```

Handling null value with `Maybe`

```javascript
var firstMatch = compose(first, match(/cat/g));
firstMatch("dogsup"); // ERROR!

// now with Maybe

firstMatch = compose(map(first), Maybe, match(/cat/g));
firstMatch("dogsup"); // => Maybe(null), no errors :)
```

Exercises: [Jsbin link](https://jsbin.com/yumog/edit?js,console)

```javascript
// Exercise 1
// ==========
// Use _.add(x,y) and map(f,x) to make a function that increments a value inside a functor
console.log("--------Start exercise 1--------");

var ex1 = map(_.add(1));

assertDeepEqual(Identity(3), ex1(Identity(2)));
console.log("exercise 1...ok!");
// Exercise 2
// ==========
// Use _.head to get the first element of the list
var xs = Identity(["do", "ray", "me", "fa", "so", "la", "ti", "do"]);
console.log("--------Start exercise 2--------");

var ex2 = map(_.head);

assertDeepEqual(Identity("do"), ex2(xs));
console.log("exercise 2...ok!");

// Exercise 3
// ==========
// Use safeGet and _.head to find the first initial of the user
var safeGet = _.curry(function(x, o) {
    return Maybe(o[x]);
});
var user = { id: 2, name: "Albert" };
console.log("--------Start exercise 3--------");

var ex3 = compose(map(_.head), safeGet("name"));

assertDeepEqual(Maybe("A"), ex3(user));
console.log("exercise 3...ok!");

// Exercise 4
// ==========
// Use Maybe to rewrite ex4 without an if statement
console.log("--------Start exercise 4--------");

var ex4 = function(n) {
    return Maybe(n).map(parseInt);
};

// or

var ex4 = compose(map(parseInt), Maybe);

assertDeepEqual(Maybe(4), ex4("4"));
console.log("exercise 4...ok!");
```

#### Either

- typically used for pure error handling
- like `Maybe` but with an error message embedded
- has two subclasses `Left` and `Right`

```javascript
var determineAge = function(user) {
  return user.age ? Right(user.age) : Left("could not get age");
};
var yearOlder = compose(map(add(1)), determineAge);

yearOlder({ age: 22 });
// Right(23)

yearOlder({ age: null });
// Left('could not get age')
```

#### I/O

- lazy computation "builder"
- typically used to contain side effects
- you must runIO to perform the operation

Exercises [Jsbin link](https://output.jsbin.com/zegat)

```javascript
console.clear();
var _ = R;
var P = PointFree;
var map = P.fmap;
var compose = P.compose;
var Maybe = P.Maybe;
var Identity = P.Id;

var Either = folktale.data.Either;
var Left = Either.Left;
var Right = Either.Right;
var IO = P.IO.IO;
var runIO = P.IO.runIO;
P.IO.extendFn();

// Exercise 1
// ==========
// Write a function that uses checkActive() and showWelcome() to grant access or return the error
console.log("--------Start exercise 1--------");

var showWelcome = compose(_.add("Welcome "), _.get("name"));

var checkActive = function(user) {
    return user.active ? Right(user) : Left("Your account is not active");
};

var ex1 = compose(map(showWelcome), checkActive);

assertDeepEqual(
    Left("Your account is not active"),
    ex1({ active: false, name: "Gary" }),
);
assertDeepEqual(
    Right("Welcome Theresa"),
    ex1({ active: true, name: "Theresa" }),
);
console.log("exercise 1...ok!");

// Exercise 2
// ==========
// Write a validation function that checks for a length > 3. It should return Right(x) if it is greater than 3 and Left("You need > 3") otherwise
console.log("--------Start exercise 2--------");

var ex2 = function(x) {
    return x.length > 3 ? Right(x) : Left("You need > 3");
};

assertDeepEqual(Right("fpguy99"), ex2("fpguy99"));
assertDeepEqual(Left("You need > 3"), ex2("..."));
console.log("exercise 2...ok!");

// Exercise 3
// ==========
// Use ex2 above and Either as a functor to save the user if they are valid

var save = function(x) {
    console.log("SAVED USER!");
    return x;
};

var ex3 = compose(map(save), ex2);

console.log("--------Start exercise 2--------");
assertDeepEqual(Right("fpguy99"), ex3("fpguy99"));
assertDeepEqual(Left("You need > 3"), ex3("duh"));
console.log("exercise 3...ok!");

// Exercise 4
// ==========
// Get the text from the input and strip the spaces
console.log("--------Start exercise 4--------");

var getValue = function(x) {
    return document.querySelector(x).value;
}.toIO();
var stripSpaces = function(s) {
    return s.replace(/\s+/g, "");
};

var ex4 = compose(map(stripSpaces), getValue);

assertEqual("honkeytonk", runIO(ex4("#text")));
console.log("exercise 4...ok!");

// Exercise 5
// ==========
// Use getHref() / getProtocal() and runIO() to get the protocal of the page.
var getHref = function() {
    return location.href;
}.toIO();
var getProtocal = compose(_.head, _.split("/"));
var ex5 = compose(map(getProtocal), getHref);

console.log("--------Start exercise 5--------");
assertEqual("https:", runIO(ex5(null)));
console.log("exercise 5...ok!");

// Exercise 6*
// ==========
// Write a function that returns the Maybe(email) of the User from getCache(). Don't forget to JSON.parse once it's pulled from the cache so you can _.get() the email

// setup...
localStorage.user = JSON.stringify({ email: "george@foreman.net" });

var getCache = function(x) {
    return Maybe(localStorage[x]);
}.toIO();
var ex6 = compose(map(map(compose(_.get("email"), JSON.parse))), getCache);

assertDeepEqual(Maybe("george@foreman.net"), runIO(ex6("user")));
console.log("exercise 6...ok!");

// TEST HELPERS
// =====================
function inspectIt(x) {
    return (
        (x.inspect && x.inspect()) || (x.toString && x.toString())
        || x.valueOf()
    ); // hacky for teachy.
}

function assertEqual(x, y) {
    if (x !== y) {
        throw "expected " + x + " to equal " + y;
    }
}
function assertDeepEqual(x, y) {
    if (x.val !== y.val) {
        throw "expected " + inspectIt(x) + " to equal " + inspectIt(y);
    }
}
```

#### Other misc functors

##### EventStream

- an infinite list of results
- dual of array
- it's map is sometimes lazy
- calls the mapped function each time an event happens

(here Brian talks about basically something really similar to rxjs)

##### Future

- has an eventual value
- similar to a promise but it's 'lazy'
- you must fork it to kick it off
- it takes a function as it's value
- calls the function with it's result once it's there

### Monads & Pointed Functors

Any Functor that has an `.of` method is a pointed Functor. Why is this
important? This concept is used in `Monads`. Monads help when you have nested
functors like `Maybe(Maybe(value))`

Monad is a `Pointed Functor` with extra function called `mjon` or `chain`

Example

```javascript
mjoin(Container(Container(2)));
// => Container(2), it's no longer nested

var getTrackingId = compose(Maybe, get("tracking_id"));
var findOrder = compose(Maybe, Api.findOrder);
var getOrderTracking = compose(
  mjoin, // flatten 2 nested Maybes into 1
  map(getTrackingId), // mapping inside Maybe returns Maybe(Maybe)
  findOrder // returns Maybe
);
```

## Lifting

Lifting just means transforming to a given type. Let's reflect on how we are writing `async` code.

```js
async function compute() {
    const v1 = await Promise.resolve(1);
    const v2 = await Promise.resolve(2);

    return v1 + v2;
}
```

There is something wrong with it right? It feels like we are `unpacking` from a structure (in this case a `promise`) and then packing again. The code above is like writing something like this:

```js
function capitalizeFirst(arr) {
    const [first, ...rest] = arr;
    const changed = first.toUpperCase();
    return [changed, ...rest];
}
```

Again, `unpacking` from a structure (in this case an `array`) and packing again.

To avoid such situations we can create `helper functions` which are `lifted` to a structure we are operating on. The `unpacking` part was clearly the fault on operators we wanted to use.

```js
func liftP(func) {
  return function(...promises) {
    return Promise.all(promises).then(func)
  }
}
```

This way we can `add` using `async` code without unpacking

```js
const adder = liftP(addTwo);
async function compute() {
    return await adder(Promise.resolve(1), Promise.resolve(2));
}
```
