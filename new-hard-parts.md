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
