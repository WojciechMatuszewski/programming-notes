# Web Performance

## Javascript Performance

Javascript is compiled language.
Most browsers use something called **just-in-time (JIT) compilation** (JS is compiled just before it's about to be used).

How it works under the hood.

1. Javascript get send to a browser
2. Files go to the parser
3. Parser turns the code into _AST_
4. That AST goes to the _baseline compiler_
5. _Base compiler_ turns stuff into byte code

There is also second compiler (optimizing compiler).
Any code that can be optimized goes there. That optimizing compiler spits out machine code.

![image info](./assets/js-parsing.png)

### Parsing

Parsing can be very slow.
You should always strive to defer as much parsing as possible.

- eager parsing: parse everything
- lazy parsing: do the bear minimum

```javascript
const a = 1;
const b = 2;
// Wrapping something in parenthesis tells the parser to parse it immediately
(function add(a, b) {
  return x + y;
});
add(a, b);
```

Try to avoid nested functions (this is micro-optimization, take it with a grain of salt)

```js
function sumOfSquares(x, y) {
  // 👇this will  repeatedly be parsed
  function square(n) {
    return n * n;
  }
  return square(x) + square(y);
}

// move the inner function out
function square(n) {
  return n * n;
}
```

### ASTS and Initial Execution

AST is a structure that shows a structure of the actual code.

![simple-ast](./assets/simple-ast.png)

### Optimizing Compiler

_Turbofan_ optimizes code in Chrome. We want to push as much code from _Interpreter_ to _Turbofan_

#### Measuring performance

```js
// Node 8+
const { performance } = require("perf_hooks");

// SETUP 🏁

let iterations = 1e7;

const a = 1;
const b = 2;

const add = (x, y) => x + y;

performance.mark("start");

while (iterations--) {
  add(a, b);
}

performance.mark("end");

performance.measure("My Special Benchmark", "start", "end");

const [measure] = performance.getEntriesByName("My Special Benchmark");
console.log(measure);
```

How do we know if our code went to _optimizing-compiler_?
We can use multiple flags to check it out.

- `node --trace-opt js_file.js`

Code above shows a function which is only used when to add **numbers**.
What would happen if we did this?

```js
...code...
while (iterations--) {
  add(a, b);
}
add('some', 'string')
...code...
```

Run with: `node --trace-deopt --trace-opt`

Well sadly our function got optimized and with the addition of `add('some', 'string')` got deoptimized.

Deoptimizing a function can make it run significantly slower.

You can also use special native-syntax to manipulate optimizations.

```js
...code...
%NeverOptimizeFunction(add)
// or
%OptimizeFunctionOnNextCall(add)
...code...
```

You have to run it with `--allow-natives-syntax`

So the great news is that any tool that allows you to enforce types will effectively speed up your application by allowing for these optimizations to occur faster.

#### Underlying Compiler Types

**Objects can be:**

- Monomorphism: same thing every time
- Polymorphism: generally the same but sometimes mixed bag
- Megamorphism: compiler will not optimize this, wild west of structures

So, how does the browser figure out what type something is?

You can use `%HaveSameMap(obj1, obj2)%` to check if object are 'the same' _compiler type wise_

#### Hidden Classes

- dynamic lookup: this object could be anything, so let me look at the rule book and figure it out (it's slow).

There is secret type system behind your back.

Everything gets something called **hidden class**.
Here is how it looks like in practice.
![hidden-classes](./assets/hidden-classes-theory.png)
The c0,c1 ... are those hidden classes (in reality there are long strings). Therese strings are used for lookup when we are accessing a property. They help compiler to optimize for not doing the lookup process again.

#### Scoping and Prototypes

```js
const makeAPoint = () => {
  class Point {
    constructor(x, y) {
      this.x = x;
      this.y = y;
    }
  }
  return new Point(1, 2);
};
const a = makeAPoint();
const b = makeAPoint();
console.log(%HaveSameMap(a,b)) //false
```

Every time we invoke `makeAPoint` we are creating a fresh reference to a different prototype (see how classes work)
That's why this code is significantly slower than the other case where we would move class definition outside the function.

#### Function inlining

```js
const square = x => x * x;
// which is faster
// (1) this ?
const sumOfSquares = (a, b) => square(a) + square(b);
// (2) or maybe this ?
const sumOfSquares = (a, b) => a * a + b * b;
```

Of course (2) is faster. But does that mean you cannot write functions the (1) way?
Turns out we have this thing called _function inlining_. Optimization compiler can inline functions on your behalf.

### Takeaways

> - The easiest way to reduce parse, compile, and execution times is to ship less code.
> - Use the **User Timing API** to figure out where the biggest amount of hurt is.
> - Use type system if you can to help with hidden classes and optimizations
