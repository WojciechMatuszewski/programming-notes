# React Stuff

## Testing

- You can use `debug` with a `DOM` node.

```js
const elem = getElementBy...
debug(elem)
```

- Use `user` from `@testing-library/user-event` for more general events. `fireEvent` only fires a singular event which might not be representative enough.

* Use `jest-axe` for a11y assertions

- considering spying on `console.error` when testing something that can throw. **Remember to restore**, probably using `error.mockRestore()`

* there is something called `toMatchInlineSnapshot`. It basically creates snapshots automatically but the `snapshot` itself is within a test code, not a separate file.

- you should not use `.toHaveBeenCalledTimes` on `mock functions` which are _rendered_ (looking at you `renderProps`). Normally for API calls that is ok, but other than that just use `toHaveBeenCalledTimes`

* asserting on dates is hard. You should create a delta between 2 dates.

```js
const preDate = new Date();

// code doing stuff

const postDate = new Date();

// here you should assert on mockCalls date probably

const dateFromMockCalls = ...mock.calls[0]..

expect(dateFromMockCalls.getTime()).toBeGreaterOrEqual(preDate.getTime())

expect(dateFromMockCalls.getTime()).toBeLessOrEqual(post.getTime())
```

- for generating test data you can you `test-data-bot` library. With that you can communicate clearly whats important in your tests and whats not by creating random fake data and specific strings that stand out, specifically created by you, not by `test-data-bot`

## Getting all elements of form from `event`

Use `event.target.elements` to get all form elements. **You have to add name prop to the form elements**

## `JSX` pragma

Sometimes when scrolling through some library (usually UI-related) documentation, you might encounter this syntax:

```js
/** @jsx jsx */
// imports

// code
```

This is usually done to **customize the transition from JSX to JS by using custom `React.createElement` function**.

One usage of this would be emotion `css` prop.

And do not look surprised, you already know one pragma (a directive really) that exists in vanilla JS, the **strict mode**.

```js
"use strict";

```

## Resetting Component State with key property

`key` is used to help React track changes and basically be able to tell whats
changed in between renders.

This may seem trivial but you can actually control this behavior right? Since
you can pass a key to every component / jsx stuff we can manually re-instantiate
a given component / node.

## Referential identity and `React.memo`

Lets re-introduce some JS basics. You probably already know this:

```js
true === true
"string" === "string"

{} !== {}
```

Pretty straightforward right? The last equality check has a huge impact on how `React.memo` works.

As you probably know, `React.memo` uses `Object.is` under the hood. It's more or less glorified `===` under the hood (works for `NaN` values and the `-0` and `+0` edge cases).

Either way, when you are passing an object as a prop to a component which uses `React.memo` you might find that component performing the whole render cycle again.

```jsx
<Component article={article} />;

const Component = React.memo((article) => {
  // still re-renders :CC
});
```

So the problem is that `article` object. The _referential identity_ is not the same between re-renders.

### Lesser known optimization

Yes, you could pass the second parameter to the `React.memo`, but you might also do something like this

```jsx
<Component {...article} />
```

Why would this work? Well, know `React.memo` will be diffing between **primitives** (granted `article` does not have an object property). Since `React.memo` diffs between _primitives_, there is no referential identity to worry about.

## Lesser known hooks

### `useDebugValue`

This hook is tightly integrated with the React dev extension.
The idea is to have more information about given hook when using that extension.

The syntax is pretty simple

```
React.useDebugValue(_value, _transformFn)
```

What it does it displays that value as additional label near given hook. Please note that **this hook is only fired whenever dev tools are open**.

### `useImperativeHandle`

This one allows you to control how what `ref` can do.

You are probably familiar with `React.forwardRef` right? The one that allows you to use `ref`s for you functional components as they were HTML elements themselves.

Now, you might want to do something different than just place that `ref` on the underlying HTML element within the component like so:

```jsx
const Component = React.forwardRef((props, ref) => {
  return <div ref={ref}>Hi there</div>;
});
```

With `useImperativeHandle` it gives you control over how the ref is used and what capabilities it exposes.
Let's say you only want your users to be able to `focus` on a particular input within your element using ref.

```jsx
const Component = React.forwardRef((props, ref) => {
  const inputRef = React.useRef()
  React.useImperativeHandle(
    ref,
    () => ({
      focus: () => inputRef.current.focus();
    }),
    []
  );

  return (
    <div>
      <input />
      <input ref = {inputRef}/>
    </div>
  );
});
```

Whenever someone decides to use `ref` on your `Component` the only method he will see will be the `focus` method.

```jsx
function Other() {
  // ref can only do .focus now
  const ref = React.useRef();
  return <Component ref={ref} />;
}
```

I think this hook is not that useful in day-to-day work, but there are probably some use cases where if you do not use it, you might have a hard time doing something.

The `useImperativeHandle` hook allows you too implement _bidirectional_ flow of the data. Just like you could with class components and `React.ref`.

## Why does `useRef` have the `current` property

When it comes to `React` there are a lot of things that can go wrong, but one of them is something called a **stale closure problem**. This can especially be a problem while working with hooks.

The simplest example would be with an `useEffect`

```jsx
function FetchData() {
  const [count, setCount] = React.useState(0);

  React.useEffect(() => {
    setTimeout(() => {
      setCount(count + 1);
    }, 2000);
  }, []);

  // increment by clicking a button
}
```

Classical example where your dependency list is wrong.

So we clearly have an stale closure problem here, this is because the `count` is a primitive value, there is no reference.

With the `ref` having the `current` property, they are passed by reference, thus the closure is on that reference, not the value itself.

This means that **the `current` property ensures that the value of the `ref` is an object, thus being passed, referenced by reference**.

So this would work

```jsx
function FetchData() {
  const [count, setCount] = React.useState(0);
  const counterRef = React.useRef(count);

  React.useEffect(() => {
    counterRef.current = count;
  });

  React.useEffect(() => {
    setTimeout(() => {
      setCount(count + 1);
    }, 2000);
  }, []);

  // increment by clicking a button
}
```

We are basically mutating an object here.

## Components as Functions Gotcha

We are now in `React.useSomething` era (_React Hooks_) so you probably are writing functional components. There is one gotcha you might encounter that will result in serious bugs.

### Rules of Hooks

There are some internal rules of usage when it comes to _React Hooks_. One of them being that you cannot use them conditionally. Another one is that hooks have to be called in the same order every _call cycle_. This one actually comes into play in this gotcha.

### The Gotcha

Lets say you have this setup

Some kind of list component

```jsx
function CountersList() {
  const [items, setItems] = React.useState([]);
  const addCounter = () => setItems((prev) => [...prev, { id: prev.length }]);
  return (
    <div>
      <button onClick={addCounter}>Add Counter</button>
      {items.map(Counter)}
    </div>
  );
}
```

And the counter:

```jsx
function Counter() {
  const [_, setCount] = React.useState(0);
  const incrementCount = () => setCount((prev) => prev + 1);
  return <button onClick={incrementCount}> Increment </button>;
}
```

Pay very close attention to the `CountersList`:

```jsx
<div>
  <button onClick={addCounter}>Add Counter</button>
  {items.map(Counter)}
</div>
```

It seems like we are mapping over items and calling `Counter`, which is a function which returns `JSX`. So everything should work right?

Well, not quite. When you click on the button you will get an error saying:

> React has detected a change in the order of Hooks called by CountersList.

To understand this message lets follow the render cycle:

1. Initial render to `CountersList` sets `items` to an empty array, nothing to _map over_.

2. You click on the button, `addCounter` fires, updating `items`. Now array is not empty, map calls `Counter`. `Counter` calls initial `setState`.

There does not seem to be a change in the order of Hooks right? In the first cycle `CountersList` items were set once, in the second also only once.

### Associating Hooks with Components

It turns out that React **cannot associate a hook with a given component when that component is called as a function**.

Normally such syntax:

```jsx
return <Counter />;
```

Would create underlying `React.Element`. When created that way, React is able to associate hooks used inside `Counter` to `React.Element`.

But when you are trying to invoke the `Counter` directly:

```js
Counter();
```

There is no `React.Element` created, thus nothing to associate hooks with.
Using the form from above is basically the same as doing

```jsx
  {items.map(() => {
    const [_, setCount] = React.useState(0)
    // rest of Counter code
  }}
```

Which is not _rules of hooks_ compliant.

#### Takeaways

So basically to avoid such errors, **which will probably cause unexpected bugs**

- never call functional components directly when they are using hooks

### Component gets invoked multiple times

TODO: talk about `React.Strict`

## SSR

SSR is getting quite popular nowadays. While developing your SSR solutions you might encounter a problem where `React` seem to be behaving weirdly, almost bug-like, when your SSR HTML structure is not aligning with the one on the server.

This is mostly due to the fact that **React is only diffing on HTML Tags level and will not try to patch up inconsistencies for you**.

This is one case I encountered at work. Suppose you have a function witch renders a `a HTML tag` using a config.

```jsx
function renderLink({to, ...restOfConfig}) {
  return <a href={to} {...restOfConfig}>
}
```

Now, what happens if the `to` prop is different on the server than on the client?. This might be the case due to eg. servers inability to pick up locale data. `config` is enriched with that data on the client.

In such case you will get a warning that HTML structure is different on the client and the server. **You might think that since it's a prop change React will just re-render that tag with a correct, newest prop. But that is not the case**.

What will happen is that the `prop` becomes stale, and your functionality will not work. Like I said before, React will not try to patch up inconsistencies, **its up to you to make sure that the HTML structure and props are the same on the server and on the client**

### Gotcha with 2 pass render

When thinking about SSR you should think about **2 pass render**. Just like in Vanilla Javascript where 2 pass happens (compilation and interpretation).

As you know Reacts `renderToString` will not be able to handle dynamic parts of your APP. This is sometimes apparent on websites where user data is displayed on the navigation or other parts of the site. That is because the lifecycle hooks do not run on the server. One approach is to include SSR state within the `script` tag inside the HML and pick it up from there. I think this is a nice solution but maybe not a fit for a simple app.

Other solutions involve checking for a `window`.

```js
function Component() {
  if (window == undefined) {
    return null;
  }
  return <div>Some content</div>;
}
```

This approach might seem like a good solution but **there is a problem with this solution**. We are violating one of the React rules, where data returned from the server is different from the data React sees during rehydration process.

Now, during the rehydration process, the Javascript is there, we are inside the browser environment so the output from the `Component` will look as follows:

```js
<div>Some content</div>
```

Meanwhile **the output of `renderToString` for that Component is just empty**.
This is a recipe for a disaster. Sometimes, React will be able to handle the differences and keep the DOM structure but when there are multiple nested elements, **it will probably lead to bugs.**

### Solution

The solution is quite simple. It relies on the fact that **useEffect will ONLY AND ONLY be run when COMPONENT IS MOUNTED**. This mean that **useEffect DOES NOT RUN ON REHYDRATION**. This is a crucial realization when dealing with SSR.

So now your component looks as follows:

```js
function Component() {
  const [mounted, setMounted] = React.useState(false);

  React.useEffect(() => setMounted(true), []);

  if (!mounted) return null;

  return <div>Some content</div>;
}
```

Now **the output from `renderToString`**:

```js
// empty
```

**During the rehydration `React` sees**:

```js
// empty
```

Again, that is because `useEffect` does not run during rehydration.

When your component mounts the output becomes

```html
<div>Some content</div>
```

## Memoization and semantic guarantee

You are probably using hooks by now. That's great. Also you probably know about
`useMemo` and `useCallback` hooks which help you _stabilize_ or _memoize_
(basically the same meaning, you would just use different terms in different
contexts) stuff.

Yet, there is a hidden caveat that comes with them. According to React docs:

> You may rely on useMemo as a performance optimization, not as a semantic
> guarantee. In the future, React may choose to “forget” some previously
> memoized values and recalculate them on next render...

This information is **huge**.

You may be thinking to yourself:

> But docs only mention `useMemo` right?

Well we are out of luck on this one too, `useCallback` can be implemented using
`useMemo`:

```jsx
const useCallback = (fn, deps) => React.useMemo(() => fn, deps);
```

There is raging discussion about this on git. Many popular libraries relay on
the fact that React does not _forget_ previously memoized values. That would be
kind meh. Wonder what `React.Suspense` will bring to the table when it comes to
that...

## Lazy State initialization (FC)

There are 2(or 3?) ways to initialize functionals component state.

First variant would be to just pass value. This sometimes results in a gotcha where state is stale when that value is coming from a prop.

```jsx
function Counter() {
  React.useState(1);
}
```

Nothing spectacular there, move along

Second way would be to use a function. Maybe you need to compute something. I would call this **eager initialization**. This is because the function is invoked every time when the functional component is invoked.

```jsx
function eager() {
  console.log("invoked");
  return 42;
}

function Component() {
  React.useState(eager());
  return null;
}
```

You will see `invoked` **whenever the parent of this component changes state**. This is quite important to understand.

Finally, there is a third way. This is where you **declare a function as initial state**.

```js
function Component() {
  React.useState(() => {
    console.log("lazy");
    return 42;
  });
  return null;
}
```

**React will not invoke the function when the parent state changes**. You might see 2 logs here and there, and this is due to nature of `React` but other than that no logs should be displayed.

### Eager with stale state (or is it?)

Ever wondered why this setup:

```js
function Eager(v) {
  return v;
}
const Component = ({ counter }) => {
  const [count, setCount] = React.useState(Eager(counter));
  return <pre>{count}</pre>;
};
```

Or simply this:

```js
const Component = ({ counter }) => {
  const [count, setCount] = React.useState(counter);
  return <pre>{count}</pre>;
};
```

Results in so called `stale-state`?. Well if you log the value inside the `Eager` function you will see the correct value of the `counter` is passed. So the state is not `stale` per se, **the component did not rerendered**. This is where **3 phases of actually rendering** comes in. Your functional component WAS invoked and the current (correct) value WAS passed into the `useState` hook, but no updates to the dom were committed.

This is happening because the _initial state_ is only set **once**, when the component is created. There is no magical sync between the initial state and the props.

## getSnapshotBeforeUpdate

It sadly seems like this is lesser known React life-cycle method. The method itself is **very useful** in some situations. Not knowing about it may cause a lot of headaches when such situations arrive. So this method will get invoked **before the most recent `render` output**. It can return a value which will be used in `componentDidUpdate`.

Think about all the times you wanted to make chat window scroll nicely on a new message. Dealing with `refs` and `requestAnimationFrame` stuff just to make sure all scrolling logic works nicely. With this you can calculate the `scrollHeight` before an update, send that to `componentDidUpdate` and proceed with logic. SO MUCH EASIER!

## getDerivedStateFromError

You are probably aware of `componentDidCatch`. Maybe you even have an _error boundary_ in which you use the `componentDidCatch` to set the state - usually something similar to `this.setState({error: true})`.

The problem here is that setting state in `componentDidCatch` will be **deprecated** in future React releases.
There is also the fact that the `componentDidCatch` is invoked during the commit phase. Setting the state there would result in, probably, aborting the current commit and starting again with the state set inside the `componentDidCatch`.

So the lesson here is that you **should use `componentDidCatch` for logging and `getDerivedStateFromError` for setting the actual state**.

## Profiling

React has it's great _profiler_ as an browser extension. That tool is really
good. But what if you want to send profiling data to your server?.

Well, _React_ has this `React.Profiler` component, which I never heard of. You
can wrap any piece of your tree with it and it gives you various info about
_render-related_ timings.

```jsx
<React.Profiler id="counter" onRender={reportProfile}>
  <Counter />
</React.Profiler>
```

`reportProfile` usually would be a queue sending data to server each X seconds.

**BEWARE** It cannot be all sunshine and rainbows. **By default react does not
include this API in production bundle**. You have to opt in yourself. This is
achieved through _webpack aliases_. The impact itself is minimal, but still
worth considering.

[Consult the docs for more info](https://gist.github.com/bvaughn/25e6233aeb1b4f0cdb8d8366e54a3977)

## Stale Closure

You are most likely familiar with this issue, when using hooks sometimes values
gets lock inside a closure and become _stale_ after a while (eg. state change
causes them to become stale).

Usually you should just listen to hooks linter and write your code normally. But
there is also an escape hatch you can use (but probably sparingly).

Let's look at `Formik`s piece of code:

```ts
function useEventCallback<T extends (args: ...any[]) => any>(callback: T): T {
  const callbackRef = React.useRef(callback);
  React.useLayoutEffect(() => {
    callbackRef.current = callback
  })
  return React.useCallback((...args: any[]) => callback.current.apply(void 0, args), []) as T;
}
```

So what does this piece of code do?

- save ref to callback
- which each component call save fresh `callback` (with new variables it uses
  because of re-render)
- return memoized callback that closes over fresh `callback` (this eliminates
  stale closure problem)

One crucial piece of code from this snippet is the following:

```ts
return React.useCallback(
  (...args: any[]) => callback.current.apply(void 0, args),
  []
) as T;
```

Lets try to return the `callback` with _point-free_ style.

```ts
return React.useCallback(callback.current, []) as T;
```

Would this work? **Nope**.

> but if you attempt to overwrite the reference it will not affect the copy of
> the reference held by the caller - i.e. the reference itself is passed by
> value

And this is the key to why it will not work. Here we are trying to overwrite the
reference (passed as `callback.current`). And since reference itself is passed
by value it will not change and always be stale.

## Synthetics Events

React is using browser-agnostic wrapper for events that are passed to you when you. There are called `Synthetic Events`. You can still get to the underlying, native events by using `.nativeEvent` property.

### Mysterious `.persist` property

Tbh, there is nothing mysterious about it. Those `Synthetic Events` are nullified synchronously after the callback (your callback) has been invoked. If you want to use events in a asynchronous context you probably want to use `.persis` method. This will give the control of the event back to you and that event will not be synchronously nullified.

## Events Listeners and Hooks

While using `useEffect` you have to remember about deps array, that is pretty
obvious. But have you ever wondered how reference capture work with event
`callbacks` functions.

Usually I defined event `callbacks` following way:

```js
React.useEffect(() => {
  function listener() {}

  window.addEventListener("scroll", listener);
  return () => window.removeEventListener("scroll", listener);
}, []);
```

Reasoning behind this is to make sure that we are holding the same reference to
the function when adding and removing that listener.

It turns out you can actually define `event listeners` outside `useEffect`.

```js
function SomeComponent() {
  function listener() {}
  React.useEffect(() => {
    window.addEventListener("scroll", listener);
    return () => window.removeEventListener("scroll", listener);
  }, []);
}
```

`useEffect` will close-over `listener` and hold it's reference event when props
change!. You do not have to pass it inside deps array.

## Rendering, Commits, Reconciliation

Many people say

> You have wasted renders in your app

> Limit re-renders in you app!

While true, do they really understand what `to rerender` mean?

- `render` is React calling your `render` function or the `FC` itself. It gets
  the DOM that way (as we know React operates on Virtual DOM and diffs previous
  DOM with next one)

- `reconciliation` is the phase of diffing the newly acquired DOM with Virtual
  one.

- `commit` is the phase of react actually updating the DOM.

So when people say that

> My app is slow because I have many components that re-render

They may really say that they have slow `commit` phases (granted re-renders can
also cause an issue)

Actually, when component re-renders, DOM actually do not have to be updated.
Again, this stems from the fact that `render` is just calling one function and
getting new DOM (which can be the same as previous one)

## Future

### Concurrent React

The biggest thing is that **concurrent react** can **partially render** a tree
without committing to the DOM. And o boi this is huge.

#### Tearing

With great power comes... great amount of bugs. This one is definitely interesting because there are multiple parts at play.
Imagine your component are reading from a shared state. Whey rerender when that state changes but the render can take some time right?.

Now imagine your tree re-rendering and then user decides to trigger the state update again. Naturally React (with concurrent mode) will yield to a browser because user did something. That will pause the rendering of SOME parts of your tree.

What is happening at that very moment? Some parts of your tree are already ready with PREVIOUS value from the state (pre-user interaction) and SOME are still to be rendered BUT the state already has a different value.

This issue is called tearing, where **some parts of your tree are inconsistent with others when it comes to state**. There is a great resource about this: https://github.com/dai-shi/will-this-react-global-state-work-in-concurrent-mode#what-is-tearing

#### Suspense and data fetching

So, when writing this, there is a trick you can use to leverage `Suspense` for
your data fetching. That is to `throw` given promise.

(this example does not really work it's just for demonstration)

```jsx
function Component() {
    throw aPromise
    return <p>Works</p>
}
function App() {
    return (
        <React.Suspense>
            <Component/>
        <React.Suspense>
    )
}
```

Ok so whats the deal?

- `Suspense` is catching that promise
- **Seems to be calling your component again on `.then` of that promise** (?)

The second part is very important to understand. This is where this notion of
`not safe` for `Suspense` comes from.

#### Breaking-down code sandbox `wrap promise`

So it's React Conf 2019 and examples are flying everywhere. Where is this one
particularly interesting function called `wrapPromise` that seems to do all the
magic.

```js
function wrapPromise(promise) {
  let status = "pending";
  let result;
  // ...
}
```

So

- `status` is needed here. `Suspense` will **call our render function more than
  once!**. This variable decides if we should return results or to `suspend`
- `result`: well this is the result of our `promise`

```js
// ..
let suspender = promise.then(
  (r) => {
    status = "success";
    result = r;
  },
  (e) => {
    status = "error";
    result = e;
  }
);
// ..
```

This is the very important part. `suspender` is the thing we are going to throw.
Notice it has, sort of, pretended the _promise resolution_ call before
`Suspense` can do anything with it.

```js
// ..
return {
  read() {
    if (status === "pending") {
      throw suspender;
    } else if (status == "error") {
      throw result;
    } else if (status == "success") {
      return result;
    }
  },
};
```

And the return statement. Inside our _component_ we are performing
`resource.read()` call. This of course can, literally, have any other name but i
think React team is doing the naming on purpose here to get us familiar with
actual `createResource` API from React.

#### Step By Step

So lets follow what happens step by step

- `wrapPromise` 'prepends' a callback when `promise` gets resolved. This is
  crucial to understand that there is a closure in play here.

- initially we throw the `promise`. **Now Suspense takes over**.

- `Suspense` waits for the promise to resolve, maybe with `.then` call, i do not
  actually know. **That `.then` call happens AFTER our 'pretended' `.then` call
  inside `wrapPromise`**.
- `Suspense` calls our _component_ again, the _component_ calls `.read()` again.

- now inside `read` data is already set by `.then` call created on `promise`,
  the one that happened before `Suspense` called our _component_ again. `status`
  is either `error` or `success` by now so we have our data in our _component_.

#### Time-Slicing

So whats the what are we _slicing_ ?

> Time-Slicing means that React has the ability to split work into chunks and
> spread it's execution over time.

Imagine a huge `render` method. It takes a lot of time to process. When writing
this react is fully sync, that means that when you have a sync task running no
other _user inputs_ can be processed until that `render` task is done.

This is kinda a bummer.

With `Time-Slicing` React will be able to _slice_ the main `render` task and
insert that _user input_ inbetween.

### Moving away from `.defaultProps`

Looking at the discussions on git about React 16.9 there is a proposal to move
away from `.defaultProps`, at lest on functional components.

The reason is pretty justified in my opinion, just look at the code snippet
here:

```js
export function createElement(type, config, children) {
  let propName;

  // Reserved names are extracted
  const props = {};

  if (config != null) {
    // Handling ref and keys
    // Assign props to prop object
  }

  // Transfer children to newly allocated props object

  // Resolve default props
  if (type && type.defaultProps) {
    const defaultProps = type.defaultProps;
    for (propName in defaultProps) {
      if (props[propName] === undefined) {
        props[propName] = defaultProps[propName];
      }
    }
  }

  return ReactElement(/*stuff*/);
}
```

The problem is that _resolving default props_ happens on every
`React.createElement` call. This may seem insignificant but image how many times
this function gets invoked.

The answer for this problem would be to use default properties available in ES6,
but the are also problems with that solution.

Bundle size bloating may occur when using destructuring and default values
(especially object ones) when transpiling to other versions.

React team still have to asses different choices but one is almost certain that
`.defaultProps` will go out of favour very soon.

### Safe refs with Concurrent Mode

So you already know the difference with `Rendering`, `Committing` and
`Reconciliation`.

The deal is that when `Concurrent Mode` is active React might call your function
(or your render method) multiple times.

This is all and good but might cause some problems with mutable objects like
`refs`

Contrived Example

```jsx
function ComponentWithRef() {
  const counter = React.useRef(0);

  counter.current++;

  return null;
}

function App() {
  const [_, setState] = React.useState(0);

  return (
    <React.Fragment>
      <ComponentWithRef />
      {/* set new object as state*/}
      <button onClick={() => setState({})}>Click me</button>
    </React.Fragment>
  );
}
```

Ok, so now, whenever you click the button the component will rerender. Whenever
you rerender the `counter.current` will increment. With 'normal' mode it works
fine. But the problem occurs when you are using `Concurrent Mode` and React
renders multiple times before committing.

Then, instead of seeing your counter incremented once, you will probably see
different number.

So the trick is to place it inside `useEffect` without any dependencies, like
so:

```js
React.useEffect(() => {
  counter.current++;
});
```

Some rules to make your life easier

- **DO NOT mutate refs current value in render if they rely oon the previous
  ref's value**

Easy as that.

## Memoization of children

This can bite you in the ass one day. Remember, when using `React.memo` on the component that takes `children` that `React.memo` will do nothing. `React.memo` **does only shallow compare**, have you looked into how the `ReactElement` obj. looks like? 👍.

One thing you might do is to `memoize` the `children` themselves before passing them as prop

```jsx
function Component({ children }) {
  const MemoizedChildren = React.useMemo(() => children, []);
  return <MemoizedComponent>{MemoizedChildren}</MemoizedComponent>;
}
```

But this is pretty bad. In such situations you such probably **reach for second argument within `React.memo`**.

## Undocumented features

### Contexts' bits

So there is undocumented feature in `Context` API. From the docs you probably
know that when `Context` changes every component that is under that uses that
`Context` will re-render.

That's why it is so important to make sure you are `memoizing` your context
value.

But there is a hidden feature that allows you to **hand pick components that you
want to be rerendered**

Lets create simple context

```jsx
const initialState = { firstName: "Harry", familyName: "Potter" };
const PersonContext = React.createContext(null);

function PersonProvider({ children }) {
  const [person, setPerson] = React.useState(initialState);

  return (
    <PersonContext.Provider value={[person, setPerson]}>
      {children}
    </PersonContext.Provider>
  );
}
```

Now lets say we have 2 components that display `firstName` and `familyName`

```jsx
function DisplayFirstName() {
  return (
    <PersonContext.Consumer unstable_observedBits={0b1}>
      {([person]) => <div>{person.firstName}</div>}
    </PersonContext.Consumer>
  );
}
function DisplayFamilyName() {
  return (
    <PersonContext.Consumer unstable_observedBits={0b10}>
      {([person]) => <div>{person.familyName}</div>}
    </PersonContext.Consumer>
  );
}
```

We've added this mysterious `unstable_observedBits`. This is kind of
_identificatior_ for given consumer.

Now how do we distinguish between those to skip on some re-rendering?

```jsx
function calculateChangedBits(
  [{familyName}],
  [{familyName: newFamilyName}]
) {
  return familyName != newFamilyName ? 0b10 : 0b1
}
const PersonContext(null, calculateChangedBits)
```

This magical `calculateChangedBits` function is like `shouldComponentUpdate` or
diffing function inside `React.memo`. Instead of returning true or false you are
returning bits, basically creating so-called **bitmask**

This is the mechanism used by redux and mobx to make sure they are only
re-rendering something that changed!

### BatchUpdates

So by now you probably know that React batches updates (calls like `setState` or
`setWhatever`[hooks])

The question is when is React doing that? Well, certainly not on every call
because that would cause a lot of overhead.

So React does that only in well knows methods like `componentDidUpdate` (and
probably other life-cycle methods) and events callbacks (like `onClick`)

But the more important thing is that **React does not batch state updates in
async callbacks**. So anything in `setTimeout` or a `Promise` wont batch.

There is a reliable way to batch state update though that method is marked as
`_unstable`.

`unstable_batchUpdates` is that method. Usage:

```javascript
ReactDOM.unstable_batchedUpdates(() => {
  setState(/**/);
  setState(/**/);
  setState(/**/);
});
```

This will make it so only one `setState` will fire.
