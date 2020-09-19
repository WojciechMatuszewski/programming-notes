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

### Keeping the default behavior while using comparator

You might be tempted to escape early (I know i'm) inside the `React.memo` comparator function.

```js
React.memo(Comp, (prev, current) => {
  return prev.PROP == current.PROP;
});
```

This might help you wil optimization but you still have to remember that **_React_ will compare every prop for you inside the default comparator function**. This matters since by escaping early you might introduce some bugs by not re-rendering when some other prop changes (maybe you added a new one).

So you should always try to preserve the default behavior while using your custom _comparator_ function, either by using some kind of library or by performing the `===` (`Object.is`) check yourself.

## Context

So you want to create context for (hopefully) some part of your application. You might write it like so

```jsx
const Context = React.createContext();

function Provider({ children }) {
  const [state, setState] = React.useState();
  return (
    <Context.Provider value={{ state, setState }}>{children}</Context.Provider>
  );
}
```

While this way of creating the _provider_ component is completely valid, what will happen if the parent of this component will re-render? You guessed it, the `value` will have brand new reference, thus **every consumer of that context will re-render**. This is quite bad.

What you can do here is to memoize the `value` that you pass onto the _provider_.

```jsx
const Context = React.createContext();

function Provider({ children }) {
  const [state, setState] = React.useState();
  const memoizedValue = React.useMemo(() => ({ state, setState }), [state]);
  return <Context.Provider value={memoizedValue}>{children}</Context.Provider>;
}
```

Remember that the `setState` is stable, just like the `dispatch` from `useReducer`.

Now, the state has to change for the consumers to re-render. This is what we really wanted from the begging.

Combine this with the _value/dispatch provider_ pattern and you are on your way to create a performant context provider :).

### When NOT to use `React.useMemo`

While using the `React.useMemo` is usually justified, sometimes it's not needed. Mainly, the **`React.useMemo` is not needed when your provider does not re-render (triggered from the parent)**. This usually means that the _provider_ is a global level provider. Maybe some kind of configuration that is static, or maybe something that you wrap your whole app with.

When you have such situation, the `children` prop will not change (usually).

```jsx
function Provider({ children }) {
  // NO NEED TO MEMOIZE CONTEXT HERE!!!
  return <Consumer>{children}</Consumer>;
}

function App() {
  return <p>My whole app here</p>;
}

export function AppWithProviders() {
  return (
    // nothing to trigger re-render of the Provider - no need to memoize the context.
    <Provider>
      <App />
    </Provider>
  );
}
```

## State Colocation

This is more of a tip rather than technique. I think you will come to this realization the more code you write. So here it goes:

> Place code as close to where it's relevant as possible

That's all, basically avoid global state, and if you can, put the state as close (preferably in) the component that is consuming / changing that state. This will make it so that when that state changes, only that component re-renders, nothing more, nothing less.

## Profiling with `React.Profiler`

There is the `React.Profiler` for profiling performance. It exposes timings of commit and render phase.

**BEWARE** It cannot be all sunshine and rainbows. **By default react does not
include this API in production bundle**. You have to opt in yourself. This is
achieved through _webpack aliases_. The impact itself is minimal, but still
worth considering.

[Consult the docs for more info](https://gist.github.com/bvaughn/25e6233aeb1b4f0cdb8d8366e54a3977)

### Note on `unstable_wrap`

At least for me, the provided gist was not clear enough so that I could understand what `wrap` is for. So the deal with wrap is that it allows you to _associate_ the _wrapped_ interaction with the current _trace_ segment.

If your _trace_ callback kicks off side effects that change state, you would want to wrap those functions using the `wrap` API so that they are _associated_ with the `trace` segment.

The notion of `wrap` and `trace` is pretty similar to X-Ray traces.
