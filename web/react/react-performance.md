# React performance

## Batching updates

React, by default batches your state updates. That is quite good, you probably had situations where you had multiple `setState` calls after each other right?

```js
function handleSomething() {
  setValue(1);
  setOtherValue(2);
}
```

The 2 calls from the snippet above would result in only one _re-render_ of you component. This is all and fine but have you noticed that **the batching does not work in _async_ or _native-events_** handlers?

```jsx
React.useEffect(() => {
  setTimeout(() => {
    setValue(1);
    setOtherValue(2);
  }, 1000);
}, []);
we;
```

This would result in **two _re-renders_**. This is the way it is and we have to live with it until React team changes the way things are.

One way to make sure your state updates are batched is to use `unstable_batchedUpdates`. The name is scarry right? `unstable` , but it's been somewhat documented on github so I would say it's more or less safe to use.

```jsx
React.useEffect(() => {
  setTimeout(() => {
    React.unstable_batchedUpdates(() => {
      setValue(1);
      setOtherValue(2);
    });
  }, 1000);
}, []);
```

Now you should see only 1 render happening :)

### Internal modes of state

You might think of React working in two modes when it comes to updating / computing your state. The _eager_ and _lazy_ mode.

The lazy mode is the behavior you are most likely familiar with. It's when React batches your updates,
all the state changes are reflected correctly because the computed state is different every time.

But what happens if you set the state to the same value multiple times?

```js
export default function App() {
  const [count, setCount] = useState(0);

  function handleClick() {
    setCount(1);

    setTimeout(() => {
      setCount(1);
    }, 1000);

    setTimeout(() => {
      setCount(1);
    }, 2000);
  }

  console.log("render!");

  return <button onClick={handleClick}>Click me</button>;
}
```

How many times you expect the "render!" to be printed in the console?

> **React will cache the result of your state update, if the next computed state is the same, the render will be skipped**.

Armed with that information, one might guess that the answer to my question would be twice. First for the initial render, second for the first state update to 1.

That is a good guess, but sadly that is not the case. The heuristics of when the update is skipped are not completely known for me. During my testing, the "render!" was printed three times.

All in all, this kind of information is not needed to use React, but I still find it valuable.
How many times someone asked you why did this component "re-rendered" (or not) and you had no clue?

Resources:

- _React Summit Remote Edition 2021_ talk by Adam Klein

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

There is one important part though, **_webpackPreload:true_ will not work for `React.lazy`**. At least I could not get it to work. You will need to manually _preload_ your stuff.
This is mostly likely because the _preloaded module_ has to be imported at module evaluation time, and you are not doing that when using `React.lazy`.

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

Remember that the `setState` is stable, just like the `dispatch` from `useReducer`. Now, the state has to change for the consumers to re-render. This is what we really wanted from the beginning.

Combine this with the _value/dispatch provider_ pattern and you are on your way to create a performant context provider :).

### Memoizing inside the "Provider" component

> That React Component Right Under Your Context Provider Should Probably Use `React.memo`

Why? To avoid re-rendering the whole part of the tree, but you might be already taking advantage of this by using `props.children` instead of rendering a component directly.

So, the following snippet is okay.

```ts
function ChildComponent() {
    return <GrandchildComponent />
}

const MemoizedChildComponent = React.memo(ChildComponent);

function ParentComponent() {
    const [a, setA] = useState(0);
    const [b, setB] = useState("text");

    const contextValue = {a, b};

    return (
      <MyContext.Provider value={contextValue}>
        <MemoizedChildComponent />
      </MyContext.Provider>
    )
}
```

And so is this one (I would argue this one is even better).

```tsx
function ParentComponent({children}) {
    const [a, setA] = useState(0);
    const [b, setB] = useState("text");

    const contextValue = {a, b};

    return (
      <MyContext.Provider value={contextValue}>
      {children}
      </MyContext.Provider>
    )
}
```

If we call the `setA`, in both cases, the children will NOT re-render. In the first scenario, the component is memoized, so React will skip it, in the second, the `children` has the same reference so React will skip it as well!

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

### `useContext` as implicit props

Think about the `useContext` as "implicit" props. **Every time the context changes, your component will re-render, even if it uses `React.memo`**.

So, the following piece of code:

```ts
const GreetUser = React.memo(() => {
  const user = React.useContext(UserContext);
  if (!user) {
    return "Hi there!";
  }
  return `Hello ${user.name}!`;
});
```

It could be thought of as the following piece of code:

```ts
const GreetUser = React.memo(({ user }) => {
  if (!user) {
    return "Hi there!";
  }
  return `Hello ${user.name}!`;
});
```

If you consider the `React.Context` as _implicit_ props, the notion of "context breaks memoization" makes sense.

## State Colocation

This is more of a tip rather than technique. I think you will come to this realization the more code you write. So here it goes:

> Place code as close to where it's relevant as possible

That's all, basically avoid global state, and if you can, put the state as close (preferably in) the component that is consuming / changing that state. This will make it so that when that state changes, only that component re-renders, nothing more, nothing less.

## Passing expensive component as a prop

So you probably know, that React, will "re-render" (be aware of render phases!) any child component (unless it's memoized) upon state / props update in the parent component.

A simple example

```jsx
function Child() {
  console.log("child function invoked");
  return null;
}

export default function App() {
  const [, forceUpdate] = React.useReducer((s) => !s, false);

  return (
    <div>
      <button onClick={forceUpdate}>click me</button>
      <Child />
    </div>
  );
}
```

Every time you click the button, you will see the `child function invoked` log message in your console.

The **reason why** it's happening, is because **React will re-create the _prop object_ of the `Child` component every time the parent changes**. By the _prop object_ I mean the object that is created when JSX is parsed to the object representation.

So let's say the `Child` is expensive to render in some way, or you want to reuse it in multiple places.
Since we do not have `slots` per se (like in angular) you might want to pass it as a prop.

Like this

```jsx
function Child() {
  console.log("child function invoked");
  return null;
}

function Parent({ child }) {
  const [, forceUpdate] = React.useReducer((s) => !s, false);

  return (
    <div>
      <button onClick={forceUpdate}>click me</button>
      {child}
    </div>
  );
}

export default function App() {
  return <Parent child={<Child />} />;
}
```

So what will happen now? **You will only see 1 log in you console, no matter how many times you press the button**.
This is because the `Child` component **props did not change because the `Child` is rendered in the `App`**. The App component in itself will not "change" at all.

So this is a neat optimization technique you might use.

Also do not be afraid of putting components as props. It's completely natural thing to do, especially if you know how JSX works under the hood.

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

## Memoization of children

This can bite you in the ass one day. Remember, when using `React.memo` on the component that takes `children` that `React.memo` will do nothing. `React.memo` **does only shallow compare**, have you looked into how the `ReactElement` obj. looks like? 👍.

### The `children` prop

Sometimes, **by utilizing the children prop** you can skip the re-render of the whole children tree. This works, because **React will NOT re-render a given part of the tree if the children return the same exact reference to given elements**.

In short, if you have a component similar to the following:

```tsx
function SomeProvider({ children }) {
  const [counter, setCounter] = useState(0);

  return (
    <div>
      <button onClick={() => setCounter(counter + 1)}>Count: {counter}</button>
      <OtherChildComponent />
      {children}
    </div>
  );
}
```

React will **skip re-rendering the "children" part of the tree if we update the `counter`**. This is a **good way to stop the "render children recursively" behavior** of React. Note that we **would still re-render children part of the tree if the parent of the `SomeProvider` re-rendered**. That is because, then, the `children` has a new reference, so the `===` equality check returns false.

## Rendering behavior

<https://blog.isquaredsoftware.com/2020/05/blogged-answers-a-mostly-complete-guide-to-react-rendering-behavior/>

I think the most important thing to remember is that **`React` will render all child components unconditionally just because parent re-rendered!**. It is crucial to remember especially when you are using `React.Context`.

```jsx
function App() {
  const forceRerender = React.useReducer((x) => x + 1, 0)[1];
  return (
    <CounterProvider>
      <Parent />
      <button onClick={forceRerender}>Force</button>
    </CounterProvider>
  );
}
```

In above snippet, your _context value_ might be _memoized_ but the `Parent` will still be _invoked_ (no commits to the DOM though) when the button is clicked. This is where the notion of **having `React.memo` in strategic parts of your app** comes in.

### Expensive initial rendering

<https://itnext.io/improving-slow-mounts-in-react-apps-cff5117696dc>

While it would be ideal for the React to be more asynchronous when it comes to rendering (that is coming in React 18! Hurray for _time-slicing_), it's not the case at the moment.

So what do you do when you are faced with a performance issues during during the rendering of your components?
One solution of this would be to defer the rendering to the point where we know that the browser is ready to take in the work without us accidentally freezing the UI.

The article I've linked contains a really neat way to do that. The solutions uses `window.requestIdleCallback` **with a timeout** to ensure that the callback where the computation of what should be rendered next actually happens.

Pretty neat I would say!

## Refs and `useSyncExternalStore`

The new `useSyncExternalStore` allows to create the pub-sub patterns with `useRef`. This **might be a good optimization for the context API** where **instead of holding a state as a value, you hold the `ref` and create a pub-sub pattern for `useSyncExternalStore`**.

It does feel weird to me that we would "sync external store" if that store is created within our own application, using Reacts APIs, but, nevertheless, I think it's a good optimization.

```jsx
const FormContext = createContext(null);

const FormContextProvider = ({
  children,
  initialValue = { username: "", password: "" }
}) => {
  const stateRef = useRef(initialValue);

  const get = useCallback(() => stateRef.current, []);

  const subscribersRef = useRef(new Set());
  const subscribe = useCallback((callback) => {
    subscribersRef.current.add(callback);
    return () => subscribersRef.current.delete(callback);
  }, []);

  const set = useCallback((value) => {
    stateRef.current = { ...stateRef.current, ...value };
    subscribersRef.current.forEach((subscriber) => {
      subscriber(stateRef.current);
    });
  }, []);

  const contextValue = useMemo(() => ({ get, set, subscribe }), [
    get,
    set,
    subscribe
  ]);

  return (
    <FormContext.Provider value={contextValue}>{children}</FormContext.Provider>
  );
};

const useFormState = () => {
  const state = useContext(FormContext);
  if (!state) {
    throw new Error("Boom");
  }

  const stateValue = useSyncExternalStore(state.subscribe, state.get);
  return [stateValue, state.set];
};
```

There is **a lot of additional complexity** using this technique. You **might be better off using a library** – they are using the same `useSyncExternalHook` hook as you (most likely).

## The cost of SVGs in your bundle

There are multiple ways you can use SVGs in your applications.

1. You can "inline" them into your JSX. Use them as components.

2. You can "inline" them into the HTML. That sounds nice but is hard to do with a component-based apps driven by a framework.

3. You can use `img` tag with a link to an SVG.

4. You can use them as sprites, using the `symbol`, `use` and `defs` elements.

All of those have drawbacks, but **inlining the SVG into the JSX appears to be the worst solution in terms of performance**. That is because, **the browser now has to parse the JS that contains SVGs, this really slows down the parsing process**. In addition, **the memory footprint of your application is much bigger** as the JS engine has to hold down onto those SVGs.

The **best solution in terms of performance/tradeoffs ratio appears to be using the SVGs as sprites** (point number 4). There are many resources about this problem, here is the list of the ones I found helpful.

- <https://kurtextrem.de/posts/svg-in-js?ck_subscriber_id=1352906140>
- <https://benadam.me/thoughts/react-svg-sprites/>
- <https://twitter.com/_developit/status/1382838799420514317>

---

In Next.js, **you can use the `Image` tag with a link to the SVG natively**. So there is no excuse to having the SVGs in your bundle, unless you really need to animate it.
