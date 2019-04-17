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

## Curry

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

## Compose

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
