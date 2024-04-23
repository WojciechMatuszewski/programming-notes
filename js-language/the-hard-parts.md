# Javascript: Hard Parts

## Principles of Javascript

### What happens when JS executes (runs) my code ?

```javascript
const num = 3;
function multiplyBy2(inputNumber) {
  const result = inputNumber * 2;
  return result;
}
const name = "Will";
```

JS will store stuff in _global memory_ (for example `num` or declaration `function multiplyBy2`).

During the first pass JS talks with _scope manager_ and declares scopes and variables:

- `var` => `undefined`
- `let/const` => `TDZ`

As soon as we start running our code, we create **global execution context**.

When you run a function (in this case `multiplyBy2`) you create **local execution context**. That execution context has local memory.

```js
function multiplyBy2(inputNumber) {}
```

`inputNumber` is a **parameter**. That parameter is stored in _local execution context memory_

When we get 'out' of a function we go back to (**in this case**) _global execution context_.JS actually knows that because of **call stack** (more on that later). Garbage collector will probably collect that functions local memory.

### Scope

Scope is about access, it gives function the ability to modify or change given variable. It defines which variables are "visible" (and where). There are two types of scope:

- global scope
- local scope (function)

Example (without strict mode!):

```js
// global scope  - includes firstNum, secondNum, and the
// function number
var firstNum = 1;
function number() {
  // local scope for number - only thirdNum is local to number()
  // because it was explicitly declared. secondNum is implicitly declared in the
  // the global scope.
  secondNum = 2;
  var thirdNum = 3;
  return firstNum + secondNum;
}
// what do we have access to in the global scope?
number(); // 3
firstNum; // 1
secondNum; // 2
thirdNum; // Reference Error: thirdNum is not defined
```

I think the best metaphor to think about the scope is to think about the russian dolls. Only the inner-most scope has access (in theory) to every variable (because of scope-chains).

In practice since everyone is parsing with BABEL we would get reference error (strict-mode auto applied)

### Execution Context

Whenever a function is called it's placed on **execution (call) stack**. This structure helps JS to keep track _where are we now in our code_.

![execution-context](./assets/execution-context.png)

As in the case of scope we have two types of _execution context_

- global execution context (always running, stops when we close given tab / browser)
- function execution context (local one, stops when a given function is finished)

#### Activation / Variable object

Each execution context has a special object associated with it.
That object holds:

- declared variables
- functions
- parameters

Try console.logging a function. See that **[[SCOPE\]]** prop?. That is where for example variables from closure go.

### Closure

Closure is (imo) most powerful concept in all JS.
Closure is the ability for a function to "remember" it's surroundings. It's like having a backpack with you at all times and being able to put stuff inside that backpack.

It's also a way to keep variables private. You would not want your backpack stolen right?

Simple example:

```js
function firstName(first) {
  function fullName(last) {
    // im going to put "first" inside my backpack :)
    console.log(first + " " + last);
  }
  return fullName;
}
var name = firstName("Mister");
name("Smith"); // Mister Smith
name("Jones"); // Mister Jones
```
