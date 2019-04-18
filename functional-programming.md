# Functional programming

## Principles

This section might contain obvious stuff but it's vital to always review the basics.

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
    text
  ),
  all("p")
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

Curry allows us to "feed" a function with one variable at a time. It's very versatile and you can create your `map` function with it and `reduce` (this will be helpful later).

```javascript
var newMap = _.curry(function(transformFunction, list) {
  var concatList = function(acc, elt) {
    // push is impure, it does not return the array, thats why we are using concat here
    return acc.concat(transformFunction(elt));
  };
  return _.reduce(concatList, [], list);
});
```

So now with `newMap` we have a 'better' than native map implementation because its curried which is huge for fp!

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

Function can meld aka compose
Simple compose example

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
var names = _.map(
  _.compose(
    get("name"),
    get("author")
  )
);
```

But do not stress about being `point-free` (will talk about this more later). Sometimes you cannot be 100% `point-free`

```javascript
var isAuthor = function(name, array) {
  return _.compose(
    _.contains(name),
    names
  )(articles);
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

Points means arguments, to be a point free means to omit arguments (passed implicitly)

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

Begin point free can help us with one of the hardest things in programming: Naming functions / variables.

### In review

- make all function inputs explicit as arguments
- these arguments can be provided over time, not just all at once
- try not to modify outside things
- compose without "glue" variables

## Objects (the functional way)

- containers / wrappers for values
- no methods
- not nouns

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

So now we have some kind of container 'convention'. Lets try to use our super-composable functions with it!

```javascript
capitalize("flamethrower"); // Flamethrower

capitalize(Container("flamethrower"));
// => [object Object]
```

Well this did not end up well. How we can solve this problem ?
Lets add some methods to `_Container.prototype`

```javascript
var _Container.prototype.map = function(f) {
  return Container(f(this.val))
}

Container('flamethrower').map(function(s) {
  return capitalize(s)
})
// Container("Flamethrower")
```

So still this seems kinda meh but at least we can use our super-composable functions on the `Container` for now.

For now think about the container as a singleton array (array with one value)

Container allows us to change types,
for example:

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

So far we've been really optimistic with our functions. We did not check any edge cases or apply any null checks. But how to handle them the functional-way?
