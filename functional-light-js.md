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

### Pure Functions

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
