# React performance

## Lazy loading

Standard technique, using `React.lazy` and the `React.Suspense` is a standard thing to do.

What you should know here is that **webpack will only load your module once**. That means that you can **call the `import` function how many times you want**. This is because there is some internal cache that keeps tract of the loaded _promises_

So what you can do here is to **preload your chunk on user interaction**. Like _hover_ or _focus_

```jsx
function loadGlobe() {
  return import("../globe");
}
const Globe = React.lazy(loadGlobe);

function Component() {
  return (
    <div onFocus={loadGlobe} onMouseOver={loadGlobe}>
      <React.Suspense fallback={<p>loading..</p>}>
        {showGlobe && <Globe />}
      </React.Suspense>
    </div>
  );
}
```

### Magic Comments

Use the _magic comments_ feature of webpack. It supports the browser script loading hint.

### `React.Suspense` and _concurrent_ mode

Lets say you are toggling a component which is imported using `React.lazy` API. Depending on your situation, there are 2 ways to do that

First would be to wrap the component in question in `React.Suspense` and toggle that

```jsx
{
  showGlobe && (
    <React.Suspense fallback={<div>loading...</div>}>
      <Globe />
    </React.Suspense>
  );
}
```

This is completely **fine today, but could be troublesome later on**. See, in the _concurrent_ version of _React_, the `React.Suspense` will wait X amount of time before it will show the fallback UI. This is **only true if no component is suspended when the `React.Suspense` mounts**. That means that you will ALWAYS see the _loading..._ fallback, it does not matter how fast the component itself loads.

So to prevent this, **do not toggle the `React.Suspense`, toggle the children**. Like that

```jsx
{
  <React.Suspense fallback={<div>loading...</div>}>
    {showGlobe && <Globe />};
  </React.Suspense>;
}
```

## Web workers

Web workers are fun, usually you will not need them, but when you do, there are few ways to go about introducing them to your codebase. One of my favorites is the `workerize` webpack loader.

All you have to do (given you have you webpack setup) is to just the loader, usually you do that inline, with the `!` syntax.

```js
import getWorkerForModule from "workerize!./module";
const module = getWorkerForModule();
```

You have to **remember though** that **introducing web workers will make your functions asynchronous**.

## `React.memo`

I'm not going to be talking about that you should not use this everywhere in your APP. This should be obvious.

One thing that is really useful while making optimizations, especially with `React.memo` is to make sure that **you are passing stable / primitive props or you implement custom comparator function**. The comparator function is especially tricky for me since you have to return `false` when the re-render should happen. Maybe I'm thinking in the context of `shouldComponentUpdate` all the time 🤔

Either way, use `useCallback` for callbacks, and try to pass only primitives for props (or stable values from _state_ or `useMemo`).
