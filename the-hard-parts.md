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

#### Call stack

Call stacks keeps track of which _execution context_ we are in at the moment. It's just a stack. You put the execution contexts inside it.
