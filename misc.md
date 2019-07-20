# Random Stuff

## Back pressure

Back pressure is when you cannot process data that is coming to you fast enough.
Let's say you are working on a search bar input. You would not want to send http request every keystroke right?. That's back pressure, there is no way for you to process keystroke information fast enough. What do you do? You probably debounce or throttle it.

Back pressure usually stems from the fact that you do not have the ability to control the producer (and are working in a _push_ environment).

Other solutions may include

- buffering
- sampling (giving only a sample of processed data)

## Memoization

Pretty standard technique to prevent unnecessary computation.
You create a cache and store previously computed results there. One catch is that you have to be careful with cache size. It might grow pretty fast and then you do you do.

**You should probably only use memoization with pure functions**.

Simple example

```js
function memoize(func) {
  return function memoized() {
    // we are doing it old school :D
    var args = Array.prototype.slice.call(arguments);
    // cache can be a closed over variable or variable on function itself
    func.cache = func.cache || {};

    var cachedResult = func.cache[args];

    if (cachedResult != null) return cachedResult;

    var computationResult = func.apply(this, args);

    func.cache[args] = computationResult;
    return computationResult;
  };
}
```
