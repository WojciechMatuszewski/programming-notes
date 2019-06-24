# JS: The New Hard Parts

## Async JS

We all know by now that JS is single-threaded and has synchronous execution model. What if we need to **wait some time before we can execute certain bits of code?**

### Solution 1 (a.k.a garbage solution)

```js
function display(data) {
  console.log(data);
}

// thread blocking function which waits for remote response
const dataFromAPI = fetchAndWait('someurl');

// user can do NOTHING here
// could be 300ms could be 10s

display(dataFromAPI);
console.log('LATER! LOSERZZZZ');
```

### Web Browser APIs

These are features that happens outside of JS thread. These features exists inside Web Browser.

- console.log
- local.storage
- XHR
- setTimeout
  ... and more

All of these are not JS features.

```js
function printHello() {
  console.log('Hello');
}

// setTimeout does not guarantee to call that function exactly after 1000ms. It will be invoked after no less than 1000ms!
setTimeout(printHello, 1000);

console.log('Me first');

// me first
// hello
```

`setTimeout` is facade for functionality happening outside JS. **JS has no timer!!**

`setTimeout` spins a `Timer` browser API. `printHello` is invoked _onCompletion_.

What happens if we have main thread blocking function running at the same time?

```js
function printHello() {
  console.log('hello');
}
function blockFor1Second() {
  // self explanatory
}

setTimeout(printHello, 0);
blockFor1Second();
console.log('meFirst');

// meFirst
// after 1000ms
// hello
```

There is a simple rule when stuff from _browser APIs_ are allowed back to JS land. **Main call stack has to be empty!**

#### Callback Queue (task Queue, macro task Queue)

_Callback Queue_ is a **Javascript engine feature**.
Along with _call stack_ we also have many queues. One of them is the _Callback Queue_. This is where our `printHello` is waiting to be placed back to _call stack_ (it's placed here after browser API is done with the timer)

#### Event Loop

_Event Loop_ is a **Javascript engine feature**.
_Event loop_ checks if _Callback Queue_ is empty and places stuff from _Callback queue_ to _Call Stack_

### Promises

Promises do two things at the same time. One in JS-land one in the Browser.

```js
function display(data) {
  console.log(data);
}
const futureData = fetch('someurl');
futureData.then(display);
console.log('meFist');
```

#### Running `fetch`

`fetch` is facade function. It kicks off _XHR_ request browser API.
`fetch` immediately returns an object:

```js
{
    status: 'pending' | 'resolved' | 'rejected',
    value: ... // will be populated with, well result
    onFulfillment: [] // ever called .then? this is where your callback ends up,
    onRejection: [] // .catch callbacks or second fetch prop
}
```

When `.value` gets updated all functions passed to `onFulfillment` array gets triggered with updated value.
To add stuff to `onFulfillment` array we use `.then` method.

#### XHR Browser API

`fetch` is also spins _XHR Browser API_ feature. When the data comes back it updates `value` prop on the _returned fetch object_ using `onCompletion` callback

#### `Micro`task Queue

We learned about `setTimeout` and `fetch` so far so let's combine them.

```js
function display(data) {
  console.log(data);
}
function printHello() {
  console.log('Hello');
}
function blockFor300ms() {}

setTimeout(printHello, 0);

var futureData = fetch('...');
futureData.then(display);
blockFor3000ms();

console.log('meFirst');

// meFirst
// console.log(data)
// hello
```

Js has additional queue. It's called **Microtask(job) Queue**.
**Event Loop prioritizes stuff inside Microtask Queue**.

## Iterators

What if we wanted to create a _stream_ of data which we can control? Like asking for another item in that stream?

> _Iterators_ automate the accessing of each element - so we can focus on what to do to each element and make it available to us in a smooth way

### Tracking current element

To start things out let's write a simple function that tracks current element and is able to return to us the next element in _the stream of data_

```js
function createFunction(array) {
  var currentIndex = 0;
  // closure
  function inner() {
    var element = array[currentIndex];
    currentIndex++;
    return element;
  }
  return inner;
}

const returnNextElement = createFunction([1, 2, 3, 4, 5]);
returnNextElement(); // 1
returnNextElement(); // 2
// ...
```

Implementation is quite simple but it shows how **powerful closure can really be**.

This _backpack_ that you can put stuff into (when creating closure) has a **very serious sounding name: _closed over variable environment_**.

### Manually creating iterators

When creating _iterators_ using _generators_ returned object has `.next` method on it. Let's do that manually (we have omitted `{value, done}` for simplicity reasons).

```js
function createFlow(array) {
  var i = 0;
  return {
    next: function nextFn() {
      var element = array[i];
      i++;
      return element;
    }
  };
}
const returnNextElement = createFlow([4, 5, 6]);
const element1 = returnNextElement.next(); // 4
```

## Generators

Generators create iterators.

```js
function* createFlow() {
  yield 4;
  yield 5;
  yield 6;
}
const returnNextElement = createFlow();
const element1 = returnNextElement.next(); // {value: 4, done: false}
```

But how we are able to get that next value?

### Power of `yield`

`yield` keyword can be interpreted as return statement but it's so much more!

> `yield` pauses generator function execution and the value of the expression following the yield is returned to the generator's caller

#### Dynamically setting what data flows to us

This is very nice. Check it out:

```js
function* createFlow() {
  const num = 10;
  // return 10
  // this expression never had a chance to assign anything to newNum
  // since the execution context got 'paused' when we yielded
  const newNum = yield num;
  yield 5 + newNum;
  yield 6;
}
const returnNextElement = createFlow();
const element1 = returnNextElement.next(); // 10

// this line is pure magic
// so we 'paused' execution context on line newNum = yield num
// now, whatever we pass as an argument will be the result of that assignment
const element2 = returnNextElement.next(2); // 7 = 5 + 2
```
