# React Stuff

## Testing

- You can use `debug` with a `DOM` node.

```js
const elem = getElementBy;
debug(elem);
```

- Use `user` from `@testing-library/user-event` for more general events. `fireEvent` only fires a singular event which might not be representative enough.

- Use `jest-axe` for a11y assertions

- considering spying on `console.error` when testing something that can throw. **Remember to restore**, probably using `error.mockRestore()`

- there is something called `toMatchInlineSnapshot`. It basically creates snapshots automatically but the `snapshot` itself is within a test code, not a separate file.

- you should not use `.toHaveBeenCalledTimes` on `mock functions` which are _rendered_ (looking at you `renderProps`). Normally for API calls that is ok, but other than that just use `toHaveBeenCalledTimes`

- asserting on dates is hard. You should create a delta between 2 dates.

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

## Resetting Component State with `key` property

`key` is used to help React track changes and basically be able to tell whats changed in between renders.

This may seem trivial but you can actually control this behavior right? Since you can pass a key to every component / jsx stuff we can manually re-instantiate a given component / node.

### Other use-case for the `key` property

React tries to re-use DOM nodes as much as possible. It might happen that you want to toggle between two separate instance of components while resetting their state every time the instance changes.

```jsx
(
  isCompany ? <Input id = "company"> : <Input id = "person">
)
```

In that case, **switching the `isCompany` variable will not cause the `Input` to reset its state**. From React perspective, the "type" of the node is the same, thus will be re-used. To force a new state, you need to add the `key` property with different value to each of the instance.

```jsx
(
  isCompany ? <Input id = "company" key = "company"> : <Input id = "person" key =  "person">
)
```

With this change, every time you toggle the `isCompany` variable, the `Input` state will reset.

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

## Hooks

### Hooks and their lifecycle

It is quite important to know when a given hook callback runs. So here it is

1. First React will `render` your component. **Keep in mind that RENDERING DOES NOT MEAN CHANGING THE DOM**. It only means invoking your function.

2. Then React will run the _lazy initializers_. These are the functions you pass to `useState`. Also this applies to the _lazy ref pattern_.

3. Then React will **update the DOM if there were any changes**. Keep in mind that **the browser has not painted yet**.

4. Then React runs the cleanup functions for the `useLayoutEffect` hook.

5. Then React runs the `useLayoutEffect` callbacks. **The browser still has not painted to the screen**.

6. After the `useLayoutEffect` callbacks were executed and there are no state updates, then the browser paints.

7. Then React runs the `useEffect` cleanup functions.

8. Then React runs the `useEffect` callbacks.

So, by following this list, you should have a good idea why using expensive calculations inside `useLayoutEffect` is quite dangerous – they delay painting! **This also applies to the _lazy initializers_** callbacks.

### `useReducer` is the cheat mode of hooks

> Based on [this article piece](https://overreacted.io/a-complete-guide-to-useeffect/#why-usereducer-is-the-cheat-mode-of-hooks)

I've always been though to put the `reducer` function outside of the component body. This made sense to me as I though that if I put it inside the component, the reference to the `reducer` would be lost and the dispatch would not be stable between re-renders.

Turns out this is not the case. **You can define the `reducer` function in the component body and close-over any properties**. This I do not have a specific use-case at hand for this technique, it is very interesting to me that we can do this.

As an alternative to what Dan posted, one could solve the issue like this:

```tsx
function reducer(state, action) {
  if (action.type === "tick") {
    return state + action.payload;
  } else {
    throw new Error();
  }
}

function Counter({ step }) {
  const [count, dispatch] = useReducer(reducer, 0);

  useEffect(() => {
    const id = setInterval(() => {
      dispatch({ type: "tick", payload: step });
    }, 1000);
    return () => clearInterval(id);
  }, [dispatch, step]);

  return <h1>{count}</h1>;
}
```

Here, the `reducer` is defined outside the `Counter`. The functionality **appears to behave the same way as before BUT it is NOT**. Every time the `step` changes we will clear the interval (rightfully so). That has an implication that we will start counting from 0 up to the specified time to run the interval callback.

This is not the case when you use the "trick" Dan describes. There, when the step changes, we will never clear the interval, so there wont be this "delay" in counting the number.

### `useLayoutEffect` and performance implications

The **code inside the `useLayoutEffect` and any state changes caused by running that code is guaranteed to run BEFORE the browser paints the screen**. This could be very useful, for example when we want to measure a certain DOM element. **But is can also cause severe performance problems**.

Imagine a scenario where you update some state inside the `useLayoutEffect`. This means that the code that "reacts" to that state change will also run synchronously before the browser gets a chance to paint the screen. **If the state change causes an expensive computation, you will delay paying the screen making the app feel unresponsive**.

Of course, the same issue can happen if the calculations inside the `useLayoutEffect` are expensive.

### Dealing with useEffect complexity

The usage of `useEffect` can introduce a LOT of complexity. The hook is well-known for its ability to produce infinite loops. It is also the source of confusion for many developers, especially developers who are just starting out.

But `useEffect` is, sometimes, a **necessary piece of API to get our job done**. **The `useEffect` is for running SOME _side effects_ your code has**.

**The `useEffect` is not meant to be the _default_ tool for running ALL _side effects_** – in some cases, it is much better to have _side effects_ placed inside _event handlers_.

```tsx
const [value, setNewValue] = useState("");

function onUpdateValue(newValue: string) {
  setNewValue(newValue);
  localStorage.setItem("value", newValue); // <- Side effect in the event handler!
}

<button onClick={onUpdateValue}>Click me</button>;
```

#### Living on a good terms with `useEffect`

1. Use `useEffect` for _synchronizing_ your component with an external system.
2. Push as many _side effects_ into event handlers as reasonably possible.
3. Consider using `useSyncExternalStore` more often.
4. **Question your decision to use `useEffect`**. The mere act of stopping and taking a step back could produce better code.

### Lesser known hooks

#### `useDebugValue`

This hook is tightly integrated with the React dev extension.
The idea is to have more information about given hook when using that extension.

The syntax is pretty simple

```
React.useDebugValue(_value, _transformFn)
```

What it does it displays that value as additional label near given hook. Please note that **this hook is only fired whenever dev tools are open**.

#### `useImperativeHandle`

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

##### Pseudo implementation

One thing I find really useful while learning is to try to re-implement things, this way I'm actually learning how a given abstraction works under the hood (even though my implementation is probably not covering all the use cases and so on).

The simplest way you could implement `useImperativeHandle` would be to just assign methods to ref in the render

```ts
const Component = React.forwardRef((props, ref) => {
  const method = () => {};

  ref.current = {
    method,
  };

  return; // stuff
});
```

The parent who is passing the ref would be able to call the `method` without any problems. This is actually the way you would do this using _class components_.

This method has one drawback and that is the fact that we are performing a _side effect_ in render (also that assignment might not be idempotent). This is something you should avoid, the _component function_ should be a pure function.

What you could do instead is to wrap the assignment with `useLayoutEffect`.

```ts
const Component = React.forwardRef((props, ref) => {
  const method = () => {};

  React.useLayoutEffect(() => {
    ref.current = {
      method,
    };
  });

  return; // stuff
});
```

Now the _side effect_ is contained. Pretty nice!

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

There are 2(or 3?) ways to initialize functional component state.

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

Results in so called `stale-state`?. Well if you log the value inside the `Eager` function you will see the correct value of the `counter` is passed. So the state is not `stale` per se, **the component did not re-rendered**. This is where **3 phases of actually rendering** comes in. Your functional component WAS invoked and the current (correct) value WAS passed into the `useState` hook, but no updates to the dom were committed.

This is happening because the _initial state_ is only set **once**, when the component is created. There is no magical sync between the initial state and the props.

## getSnapshotBeforeUpdate

It sadly seems like this is lesser known React life-cycle method. The method itself is **very useful** in some situations. Not knowing about it may cause a lot of headaches when such situations arrive. So this method will get invoked **before the most recent `render` output**. It can return a value which will be used in `componentDidUpdate`.

Think about all the times you wanted to make chat window scroll nicely on a new message. Dealing with `refs` and `requestAnimationFrame` stuff just to make sure all scrolling logic works nicely. With this you can calculate the `scrollHeight` before an update, send that to `componentDidUpdate` and proceed with logic. SO MUCH EASIER!

## getDerivedStateFromError

You are probably aware of `componentDidCatch`. Maybe you even have an _error boundary_ in which you use the `componentDidCatch` to set the state - usually something similar to `this.setState({error: true})`.

The problem here is that setting state in `componentDidCatch` will be **deprecated** in future React releases.
There is also the fact that the `componentDidCatch` is invoked during the commit phase. Setting the state there would result in, probably, aborting the current commit and starting again with the state set inside the `componentDidCatch`.

So the lesson here is that you **should use `componentDidCatch` for logging and `getDerivedStateFromError` for setting the actual state**.

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
return React.useCallback((...args: any[]) => callback.current.apply(void 0, args), []) as T;
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

`useEffect` will close-over `listener` and hold it's reference event when props change!. You do not have to pass it inside deps array.
Not having the `listener` in the dependency array is not a bug. The linter is smart enough to trace the code inside the `listener` function. **As long as we are not using any outside variables in the `listener` function, the linter will not complain**.

```js
function SomeComponent() {
  const [state, setState] = useState(false);

  function listener() {
    state;
  }

  React.useEffect(() => {
    window.addEventListener("scroll", listener);
    return () => window.removeEventListener("scroll", listener);
  }, []); // the linter will complain as the `listener` uses an "outside" variable – the state
}
```

## Rendering, Commits, Reconciliation

Many people say

> You have wasted renders in your app

> Limit re-renders in you app!

While true, do they really understand what `to rerender` mean?

- `render` is React calling your `render` function or the `FC` itself. It gets the DOM that way (as we know React operates on Virtual DOM and diffs previous DOM with next one)

- `reconciliation` is the phase of diffing the newly acquired DOM with Virtual one.

- `commit` is the phase of react actually updating the DOM.

So when people say that

> My app is slow because I have many components that re-render

They may really say that they have slow `commit` phases (granted re-renders can also cause an issue)

Actually, when component re-renders, DOM actually do not have to be updated. Again, this stems from the fact that `render` is just calling one function and getting new DOM (which can be the same as previous one)

## Future

### Concurrent features

The biggest thing is that **while using concurrent features, React** can **partially render** a tree
without committing to the DOM. And o boi this is huge.

#### Tearing

With great power comes... great amount of bugs. This one is definitely interesting because there are multiple parts at play.
Imagine your component are reading from a shared state. Whey rerender when that state changes but the render can take some time right?.

Now imagine your tree re-rendering and then user decides to trigger the state update again. Naturally React (using _concurrent features_) will yield to a browser because user did something. That will pause the rendering of SOME parts of your tree.

What is happening at that very moment? Some parts of your tree are already ready with PREVIOUS value from the state (pre-user interaction) and SOME are still to be rendered BUT the state already has a different value.

This issue is called tearing, where **some parts of your tree are inconsistent with others when it comes to state**. There is a great resource about this: <https://github.com/dai-shi/will-this-react-global-state-work-in-concurrent-mode#what-is-tearing>

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
  },
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
insert that _user input_ in-between.

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

### Safe refs with concurrent features

So you already know the difference with `Rendering`, `Committing` and
`Reconciliation`.

The deal is that when _concurrent features_ are active React might call your function
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
fine. But the problem occurs when you are using _concurrent features_ and React
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

  return <PersonContext.Provider value={[person, setPerson]}>{children}</PersonContext.Provider>;
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
    [{ familyName }],
    [{ familyName: newFamilyName }],
) {
    return familyName != newFamilyName ? 0b10 : 0b1;
}
const PersonContext;
```

This magical `calculateChangedBits` function is like `shouldComponentUpdate` or
diffing function inside `React.memo`. Instead of returning true or false you are
returning bits, basically creating so-called **bitmask**

This is the mechanism used by redux and mobx to make sure they are only
re-rendering something that changed!

### Manually batching updates

So by now you probably know that React batches updates that relates to state changes.

The question is when is React doing that? Well, certainly not on every call
because that would cause a lot of overhead.

So React does that only in well knows methods like `componentDidUpdate` (and
probably other life-cycle methods) and events callbacks (like `onClick`)

But the more important thing is that **React does not batch state updates in async callbacks and non-browser handlers**. So anything in `setTimeout` or a `Promise` wont batch.

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

### Automatic batching in React 18

As of writing this, the _React 18 alpha_ was released. This version of React introduced _automatic batching_.
With _automatic batching_ you no longer have to use the `unstable_batchUpdates` function. What's important here is that **there has to be a `createRoot` somewhere up the tree**.
The above requirement makes me believe that the _automatic batching_ feature is leveraging the _time slicing_ concept which is pretty cool!

Resources:

- <https://github.com/reactwg/react-18/discussions/21>

#### Opting out

You can opt out batching when calling `setState` by wrapping the `setState` call with `flushSync`.

```js
const [state, setState] = useState()

const syncUpdate = () => {
  flushSync(() => {
    setState(...)
  })
}
```

**This is very useful for manually controlling focus and the [_View Transitions API_](https://developer.mozilla.org/en-US/docs/Web/API/View_Transitions_API)**.

## Inert attribute in React 18

> Based on [this great article](https://www.mayank.co/notes/inert-in-react)

The `inert` attribute is pretty awesome, but at the time of writing this, **it poses a bit of a challenge to use in React**.

Since this is a new attribute, **React does not "recognize it yet"**. You can set it, and it will work as expected, but **as soon as React officially supports it, your JSX might break depending how you set it**.

```jsx
// Will break when updating to React 19.
<div inert=""></div>

// Does not work at all.
<div inert></div>

// MIGHT break, and TypeScript does not like it.
<div inert = "true">
```

So, how do you "properly" set the `inert` attribute on the element? **Use the _callback ref_ pattern!**

```jsx
<div
  ref={(element) => {
    if (!element) {
      return;
    }
    element.inert = true;
  }}
></div>
```

**You might want to make the callback you pass to the `ref` prop stable** since [React will call this function every time it changes](https://react.dev/reference/react-dom/components/common#ref-callback).

```jsx
function applyInert(ref) {
  if (!ref) {
    return;
  }

  ref.inert = true;
}

// Using `useCallback` is also fine, but why bother if it's not necessary?
<div ref={applyInert}></div>;
```

Keep in mind that you can **compose the callback ref functions**. [There are packages that can help with that](https://www.npmjs.com/package/@radix-ui/react-compose-refs/v/1.1.0-rc.7).

## `useActionState`

The `useActionState` (previously `useFormState`) allows you to track the _state_ of the action.

A lot of examples you are going to see will most likely use this hook in a context of _form submission_, but know that you can call the action returned from this hook however you wish – it does not have to be a form!

### Batching of the `dispatch` function

> Based on [this video](https://youtu.be/p_wnN5VR9Ok?t=1928)

This is the coolest thing ever!

So, consider a scenario where someone submits the form twice. Without the `useActionState`, you would have to handle requests that possibly arrive out-of-order on your own. **That is not the case with `useActionState`**.

**When transition is active, the `dispatch` function will queue any additional calls to this function and execute them IN ORDER they occurred** – how awesome is that?

```tsx
const [state, dispatch, isPending] = useActionState(async (previous: unknown, payload: FormData) => {
  const result = await saveAttendee(payload);
  console.log("result", result);
  return result;
}, undefined);

async function action(formData: FormData) {
  dispatch(formData);
}
```

If I pass the `action` prop to the `action={}` on the `form`, I do not have to worry about multiple submissions. All of them will be queued, and I can also use the `previous` to decide what to do!

### Resetting the form

> Based on [this blog post](https://www.nico.fyi/blog/reset-state-from-react-useactionstate)

One of the things we get from `useActionState` is the `state` of the action. The `state` is either the _initial_ state or the value we returned from the action. In your component, you will most likely use this `state` to display something to the user.

```tsx
const initialState = "";

const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
  await submitName();
  return state;
}, initialState);

// Somewhere down in the component

<p>Your name is {state}</p>;
```

**How do we reset this `state` variable**?

When using a form library, like `react-hook-form`, you can call `reset()` and the values of the form will be reset to their default values. The `useActionState` does not give us a `reset` function, so what can we do?

#### First attempt – the `reset` button

My initial inclination was to reach for the `type="reset"` HTML button, like so.

```tsx
export function Page() {
  const initialName = "Wojciech";
  const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
    await submitName();
    return state;
  }, initialName);

  return (
    <Fragment>
      <form
        className={"p-4 border"}
        action={(formData) => {
          dispatch(formData.get("name") as string);
        }}
      >
        <legend>Name form</legend>
        <fieldset>
          <label htmlFor="name"></label>
          <input type="text" name="name" id="name" defaultValue={state} />
          <button className={"block mt-4"} type="submit">
            {isPending ? "Submitting..." : "Submit me"}
          </button>
          <button type="reset">Reset</button>
        </fieldset>
      </form>
      <p>Your name is {state}</p>
    </Fragment>
  );
}
```

The `type="reset"` button works on the same basis as [the `type="reset"` input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/reset) – it will reset all the controls in the form to their initial values.

**But that only works for _uncontrolled_ inputs**. When we click `reset`, the `action` callback is not fired, so the value of the input is preserved.

#### Second attempt – calling the action again

We can call the action again, this time with the `initialName`.

```tsx
export default function Page2() {
  const initialName = "Wojciech";

  const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
    if (state === initialName) {
      return initialName;
    }

    await submitName();
    return state;
  }, initialName);

  return (
    <Fragment>
      <form
        className={"p-4 border"}
        action={(formData) => {
          dispatch(formData.get("name") as string);
        }}
      >
        <legend>Name form</legend>
        <fieldset>
          <label htmlFor="name"></label>
          <input type="text" name="name" id="name" defaultValue={state} />
          <button className={"block mt-4"} type="submit">
            {isPending ? "Submitting..." : "Submit me"}
          </button>
          <button
            type="reset"
            onClick={() => {
              dispatch(initialName);
            }}
          >
            Reset
          </button>
          {/* <input type="reset" value={initialState} /> */}
        </fieldset>
      </form>
      <p>Your name is {state}</p>
    </Fragment>
  );
}
```

Notice the condition within the `useActionState` callback. Every `dispatch` call will trigger the `isPending` if your perform any async-work. This might or might not be what you want.

#### Built-in automatic reset

**If you do not provide the `value` or the `defaultValue`, React will reset the form for you after action is done**.

```tsx
export default function Page2() {
  const initialName = "Wojciech";

  const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
    await submitName();
    return state;
  }, initialName);

  return (
    <Fragment>
      <form
        className={"p-4 border"}
        action={(formData) => {
          dispatch(formData.get("name") as string);
        }}
      >
        <legend>Name form</legend>
        <fieldset>
          <label htmlFor="name"></label>
          <input type="text" name="name" id="name" />
          <button className={"block mt-4"} type="submit">
            {isPending ? "Submitting..." : "Submit me"}
          </button>
        </fieldset>
      </form>
      <p>Your name is {state}</p>
    </Fragment>
  );
}
```

This is quite nice. Prior to that change, you had to manually reset the form via `.reset` called on the ref to the form node. Not that great DX.

## `cache`

> [Based on this great video](https://www.youtube.com/watch?v=MxjCLqdk4G4).

I first **thought that I could use the `cache` function to memoize promises for `Suspense` on the client-side**. It turns out that the **`cache` function is NOT a good candidate for this use case**.

The reason is that the **`cache` function memoizes values only for the duration of the request**. So, usage of this function should be scoped to _server components_. You **can use the `cache` in _client components_ ([reference](https://youtu.be/uttgAAZUdYk?t=3023)), but the function is a noop in that environment**.

I like to think about the `cache` as **_data loader_ WITHOUT the batching**.

[According to this video](https://youtu.be/uttgAAZUdYk?t=1674), the `cache` is internally implemented via `AsyncLocalStorage`.

## The problem with tearing

Imagine you have a global store. With _React_ 17 and down, you can freely use the store without any issues.
You update to _React_ 18 and observe a weird behavior, where tree parts have different store values. What gives?

The answer lies in the fact that _React_ 18 uses _concurrent rendering_ technique. The technique is about yielding to
the browser – i.e., "pausing" _React_ rendering. **If an update sneaks in between this "pause," one part of the tree
might show different values!**

### Note on concurrent rendering

According to [this article](https://vercel.com/blog/how-react-18-improves-application-performance), **React will yield
back to the main thread every 5 MS to check if there is not any more important task to do**. This is quite interesting,
they must use some kind of timer to do the yielding.

### Why is it not a problem in _React_ 17?

_React_ 17 is synchronous. There is no way for an update to sneak in when _React_ yields to the browser because no
yielding occurs.
That is why you have not observed this behavior yet.

## `startTransition` API

The `startTransition` API is meant to be used for **updates that can be deferred**. The most important thing to note
here is that the **callback runs synchronously, but the state update it causes is treated as low priority**.

It seems to me like `startTransition` API should most likely be used for expensive, **local** computations that are
not "important", i.e are not user interactions.

```ts
const [startTransition, isPending] = useTransition();

return (
    <button
        onClick={() => {
            startTransition(() => {
                // states updates here
            });
        }}
    >
    </button>
);
```

When you use `startTransition`, the **React will prepare a new tree in the background**. Once that tree is finished
rendering, the result can be committed into the DOM.

The **`startTransition` API will not help you in the case of CPU-heavy operations**. If the main thread is blocked, then
it will be blocked, regardless if you wrap the computation with `startTransition` or not. According to [this video](https://www.youtube.com/watch?v=T8TZQ6k4SLE), **React yields every 5ms to pool for the user interactions**. If such occur, it will attempt to interrupt the current work. This means that **it is much better to have many small tasks than to have multiple large tasks**, at least from the `startTransition` API perspective.

### `startTransition` and the `hydrateRoot` API

React 18 changed how the entrypoint to the application looks like. If your application uses SSR, now you have to
use `hydrateRoot` instead of `hydrate`. For CSR, you have to switch from `render` to `root.render`. **But did you know
there are taggable benefits in wrapping the `hydrateRoot` with `startTransition`?**

The `startTransition` marks the update as non-blocking. This is quite important as it **allows the browser to more
efficiently manage resources**. If your hydration takes a long time and is resource intensive, the website will not
longer be "locked" or "frozen".

```tsx
startTransition(() => {
    hydrateRoot(...)
})
```

[I've noticed the usage of `startTransition` in one of the the Epic Web Dev codebases](https://github.com/epicweb-dev/web-auth/blob/143e4eea6451397094cc48fa49eb6d0a2ff00fcb/exercises/01.cookies/01.problem.fetcher/app/entry.client.tsx#L5). [See official `remix` docs](https://remix.run/docs/en/main/file-conventions/entry.client)
and also [this twitter thread](https://twitter.com/dan_abramov/status/1567852606642348032).

[Next.js also does this](https://github.com/vercel/next.js/blob/90bfbe72bb79a0f6951c9e0eae24d6aa8a6c342d/packages/next/src/client/app-index.tsx#L308)
which would confirm that this is a performance
optimization, [like this post on GitHub](https://github.com/vercel/next.js/discussions/40691).

Sadly, the official React documentation does not mention using `startTransition` with `hydrateRoot`.

### Regarding the state updates

Remember that with _React_ 18, the state updates are batched together. In previous versions in _React_, this was not
necessarily always the case. The batching of state updates also applies to the callback of the `startTransition` function.

In addition to batching, **React will hold off applying the updates to our current UI until the new tree is fully ready to be rendered**.

```ts
startTransition(() => {
  setState((count) => count + 1);

  router.push("post/123");
});
```

In the code above, since the `router.push` causes a transition (so, in reality, we have a transition nested inside a transition), **the `setState` call will be reflected in the UI AFTER the `router.push` transition ends**.

This mechanism is a base for [building progress bars in Next.js with app router](https://buildui.com/posts/global-progress-in-nextjs).

#### Resetting RSC errors

> [Here is a great video](https://www.youtube.com/watch?v=idEL0dv2V1A) regarding this subject

When RSC errors, Next.js will attempt to render the `error.js` file.

```jsx
'use client'

export function default Error({error, reset}) {
  return <button onClick = {() => {
    reset() // This might not work as you expect!
  }}>Reset</button>
}
```

**If this component is displayed because of an error in RSC, clicking the "Reset" button will not "refresh" the UI**. This is a bit puzzling at first. The **`reset` function only resets the state of ErrorBoundary. It will not re-fetch the RSC**.

So, your next attempt might look like this.

```jsx
'use client'

export function default Error({error, reset}) {
  const router = useRouter()

  return <button onClick = {() => {
    router.refresh()
    reset() // This might not work as you expect!
  }}>Reset</button>
}
```

The `router.refresh` should re-fetch the RSC payload, so in theory, the UI should "refresh" right?

Well, that is not the case either. **By calling `refresh` and `reset` together, we are introducing a race condition**. The `refresh` will make a network request for the new payload. Until that is finished, calling `reset` will have no effect.

Since `refresh` does not return a _Promise_ for us to `await`, we have to use _transitions_!

```jsx
'use client'

export function default Error({error, reset}) {
  const router = useRouter()

  return <button onClick = {() => {
    startTransition(() => {
      router.refresh()
      reset() // This might not work as you expect!
    })
  }}>Reset</button>
}
```

**Using `startTransition` tells React to apply the updates after ALL the changes caused by functions run in _transition_ are done**. This means we no longer have to deal with the race condition I mentioned.

### Behavior in the context of network requests

While many use-cases for the `startTransition` go beyond data fetching, the API can still enhance how users perceive
loaders in the context of network requests.

You are probably familiar with the issue where a loader is shown only for a split second. This creates not-so-great
experiences for the user.

We can wrap the code that updates the "network resource" we are currently working on with `startTransition` to combat
this issue. Doing so will instruct `React` to "defer" state updates until the "network resource" is ready.

```ts
import { startTransition, useState } from "react";
import { suspensify } from "suspensify";

function fetchPokemon(id: number) {
    return fetch(`https://pokeapi.co/api/v2/pokemon/${id}`).then((res) =>
        res.json()
    );
}

const initialPokemonResource = suspensify(() => fetchPokemon(1));

// This component is wrapped with suspense boundary
function PokemonDetail() {
    const [pokemonResource, setPokemonResource] = useState(
        initialPokemonResource
    );

    const pokemon = pokemonResource;
    return (
        <div>
            { pokemon.name }
        < br / >
        <button
            onClick = {()
=>
    {
        // Defer the state update and the suspense placeholder till either this resource is "ready" or some time passed.
        startTransition(() => {
            setPokemonResource(suspensify(() => fetchPokemon(pokemon.id + 1)));
        });
    }
}
>
    Fetch
    < /button>
    < /div>
)
    ;
}
```

#### With `async` function

In React 19, you can pass an `async` function to the `startTransition`. **React will wait for all committed transitions to settle before submitting the update to the DOM**.

```js
startTransition(async () => {
  const data = await getData();
});
```

### Error handling

I know about at least two ways to handle errors with `startTransition`.

- Use `useState` hook.

- Use `ErrorBoundary` components.
  - **When an error occurs within the `useTransition` function, React will propagate the error up**. This is where `ErrorBoundary` comes in handy!

### The problem with `startTransition`

The `startTransition` API is not flexible.

- If used, the child components automatically have to opt into the `concurrent` behaviors.

- Must be used where the state is set. The `startTransition` callback has to contain a state update. This will most
  likely result in prop drilling.

There is one API that solves those issues. Enter the `useDeferredValue`.

### Deferring updates and UX

> Based on [this blog post](https://www.charpeni.com/blog/dont-blindly-use-usetransition-everywhere)

When used properly, `useTransition` can really improve the UX of an application. **However, there are also situations where it might _seem_ like using `useTransition` is a good idea, but doing so might result in poor UX**.

The blog post I linked critiques an example from the React documentation in which `useTransition` is used to "prepare" loading a "slow view" while still showing the current UI after interaction.

The way it’s implemented, only the "link" in the navigation is aware of the transition—not the "outlet" component that renders different views. So, we have an `isPending` UI on the "link," but the "outlet" still shows the old content without any indicators.

**One solution would be to _lift_ the usage of the `useTransition` hook to the "outlet" that renders different routes**. Then you could make the whole UI aware of the `isPending` boolean.

**Another, even better solution would be to use `useOptimistic` on the state within the "outlet"** ([see this comment](https://github.com/charpeni/charpeni.com/discussions/112#discussioncomment-14972056)).

**`useOptimistic` can be used for urgent state updates during an ongoing transition**. Spot on!

## `useDeferredValue`

> [Here is a great blog post](https://www.joshwcomeau.com/react/use-deferred-value/) about this topic.

This API is for **scheduling a low-priority render** for **components that are slow to render**. It will **not help you if the computation of the value you want to deffer is CPU heavy** – React has control over _rendering_, not computing the values that go into components as props.

### Under the hood

So, what happens when you click _"Increment"_ in the following code:

```jsx
function App() {
  const [count, setCount] = React.useState(0);
  const deferredCount = React.useDeferredValue(count);

  return (
    <>
      <ImportantStuff count={count} />
      <SlowStuff count={deferredCount} />
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </>
  );
}
```

- React will perform the state update.

- React will **schedule a high-priority render** with **`count` set to 1** and **`deferredCount` set to 0**.

  - Notice that the `deferredCount` "lags" behind the "real" value.

- **After** rendering the UI with `count` set to 1, **React will re-render the UI with `deferredCount` set to 1**.

  - Since the high-priority update is already done, **React can interrupt the low-priority render** if necessary.

  - The ability for React to interrupt the render is crucial to performance improvements.

    - **Always keep in mind that, if the prop change results in a heavy IO computation, there still be lag**.

### Use `React.Memo` on "slow" components

> Other gotchas are also work keeping in mind! Consult [official React documentation](https://react.dev/reference/react/useDeferredValue#caveats) to learn more.

**The component you pass the `deferredValue` has to be memoized and stable**.

In our case, the `SlowStuff` has to be wrapped with `React.Memo`. **If you do not do that, the "high priority" update will also re-render the `SlowStuff` component making the UI choppy**.

Remember, the **mental model** here is **first, the high-priority fast update, THEN the low-priority slow update**. This is why React will re-render the UI twice!

The `React.Memo` allows the first render to "skip" rendering the `SlowStuff` as it was already rendered during the "initial" render.

### The `isPending` like indicator

The `useTransition` exposes the `isPending` boolean so we can implement our loading states. Can we do the same with `useDeferredValue`? **Yes we can**!

```jsx
function App() {
  const [count, setCount] = React.useState(0);
  const deferredCount = React.useDeferredValue(count);

  const isPending = count !== deferredCount;

  return (
    <>
      <ImportantStuff count={count} />
      <SlowStuff count={deferredCount} />
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </>
  );
}
```

If we are in the middle of the low-priority render, the `deferredCount` will be different than the `count`. At this time we are "loading" or "computing".

### The initial render problem

We have one problem in our code – the first render is choppy!

React renders both the `ImportantStuff` and the `SlowStuff` at the same priority level. Understandable, as there is no "old" UI to fallback to yet.

**This is where the second parameter of the `useDeferredValue` comes in handy**.

```jsx
function App() {
  const [count, setCount] = React.useState(0);
  const deferredCount = React.useDeferredValue(count, null);

  const isPending = count !== deferredCount;

  return (
    <>
      <ImportantStuff count={count} />
      {deferredCount != null ? <SlowStuff count={deferredCount} /> : null}
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </>
  );
}
```

Now, we first render the most important parts of the UI. Then React can take care of the low-priority render and interrupt it if necessary.

### Integration with Suspense

The `useDeferredValue` integrates with `Suspense`.

**When you pass a deferred prop to a component that suspense, React will show the old UI rather than the fallback**.

This is aligned with how the `setTransition` works!

## The `use` hook

At the time of writing, there are two main use cases for the `use` hook.

- Trigger a _suspense boundary_ when passing a promise to the `use` hook.

```js
// Has to be stable in-between re-renders.
// So, using a cache, or creating it _outside_ of React component
const stablePromise = ...

function Component() {
  const result = use(stablePromise);
}
```

- **Conditionally** get the value of `React.Context`.

```js
if (someCondition) {
  const context = use(FormContext);
}
```

## `flushSync` (from `react-dom`)

Most of the `setState` calls are queued. For example

```jsx
<button
  onClick={() => {
    setName("foo");
    setNumber(1);
  }}
>
  Click me
</button>
```

The two `setXX` calls in the `onClick` handler will be queued together and processed together resulting in only one
re-render pass. **While this behavior is mostly what you want, sometimes it makes certain things hard to do. One of them
being focus management**.

Imagine a scenario where you are flipping between `input` and a `button`. You want to focus the input when you click the
button and vice-versa. **However, you do not want to focus the button when the `blur` event fires on the `input`**. This
could be achieve by adding additional state to the application. You would track the _last action_ and then
in `layoutEffect` conditionally call the `.focus` on the right element reference.

Would it be nice to get rid of the `layoutEffect` altogether? We have to use it, because the changes
to `setEditing(true/false)` are reflected asynchronously.

```jsx
<button
  onClick={() => {
    setEditing(true);
    // You cannot call the `focus` here as the UI has not been updated yet.
  }}
>
  Edit
</button>
```

**The `flushSync` API allows us to achieve just that**. If we wrap the `setEditing` in the `flushSync`, we tell React to
**update the UI BUT NOT THE STATE synchronously**.

```jsx
<button
  onClick={() => {
    flushSync(() => {
      setEditing(true);
    });
    // You CAN call the `focus` here. The UI was updated.
    inputRef.current.focus();
  }}
>
  Edit
</button>
```

**Wrapping state updates within `flushSync` is a _deoptimization_**. Since the update cannot be interrupted as it is
synchronous, **React will not perform transitions**. In most cases that is a big deal, but **it is completely fine if
the state update only affects a small portion of the tree**. In our case, we only are flipping between an `input` and
a `button` HTML tags so the _deoptimization_ introduced by `flushSync` is not a big deal.

Check out [this great tweet explaining the API based on the example I alluded to above](https://twitter.com/ryanflorence/status/1722358755499913582).

## `useSyncExternalStore`

It seems like the `useSyncExternalStore` is meant to be a drop-in replacement for _subscription-like_ hooks. The idea is
to make sure tearing never happens. Let us write `useIntervalHook` that utilizes the `useSyncExternalStore`.

```jsx
let now = new Date().toISOString();
const subscribers = new Set();

setInterval(() => {
  now = new Date().toISOString();
  subscribers.forEach((notify) => notify());
}, 1000);

const onSubscribe = (notify) => {
  subscribers.add(notify);
  return () => {
    subscribers.delete(notify);
  };
};

const onSnapshot = () => now;

function App() {
  const value = useSyncExternalStore(onSubscribe, onSnapshot);
  return <div>{value}</div>;
}
```

The `now` is the _external store_ value. The `setInterval` simulates changes. The `useSyncExternalStore` is a bridge
between the _module scope_ and _React_ rendering lifecycle.

---

You might wonder why the `notify` function is not taking any parameters? Would not that be more straightforward?
Instead, we have to create the `onSnapshot` function. The answer to this question lies in understanding how rendering
works in React 18.

Before React 18, the rendering was synchronous. If React started rendering the tree, it had to finish in one go. With
React 18, that is no longer the case – the rendering is interruptible.

Interruptible rendering means that, in extreme cases, if not taken into account, React could render part of your tree
with state X and part of the tree with state Y (the update of the state happened in-between the interruption). To
prevent such occurrences, as they relate to external stores, **instead of using the "live store value", React takes
the "snapshot" value and performs the rendering cycle using that particular value for the whole process, even if it is
interrupted**.

### The usefulness of `useSyncExternalStore`

#### Subscribing to events

It turns out the `useSyncExternalStore` hook is useful in the global context and not only for a library maintainers. In
particular, [this blog post](https://thisweekinreact.com/articles/useSyncExternalStore-the-underrated-react-api#link2)
has two examples which really speak to me.

Here is one for the scroll position state.

```jsx
import { useSyncExternalStore } from "react";

const useOptimizedScroll = (selector = () => null) => {
  const subscribe = (notify) => {
    window.addEventListener("scroll", notify);
    return () => {
      window.removeEventListener("scroll", notify);
    };
  };

  const getSnapshot = () => {
    /**
     * If the selector returns the same value multiple times, the React will NOT update the subscriber.
     * Only unique values count.
     */
    return selector(window.scrollY);
  };

  return useSyncExternalStore(subscribe, getSnapshot);
};

function App() {
  const scrollPosition = useOptimizedScroll((value) => {
    return Math.floor(value / 100) * 100;
  });

  return (
    <div style={{ height: "300vh" }}>
      <div style={{ position: "fixed" }}>{scrollPosition}</div>
    </div>
  );
}

export default App;
```

The alternative being using `useRef` and `useState`. I would say the `useSyncExternalState` version is much easier to
reason about (especially since the pub-sub model is so widely used).

##### `useLocalStorage` hook

> Base on [this great blog post](https://www.nico.fyi/blog/ssr-friendly-local-storage-react-custom-hook)

The `useSyncExternalStore` is a perfect fit for `useLocalStorage` hook.

```tsx
function useLocalStorage({ key }: { key: string }) {
  const storageEventTarget = useMemo(() => {
    return new EventTarget();
  }, []);

  const value = useSyncExternalStore(
    (onStoreChange) => {
      const controller = new AbortController();

      window.addEventListener(
        // Only fires when change is made in another tab or window
        "storage",
        () => {
          onStoreChange();
        },
        { signal: controller.signal },
      );

      storageEventTarget.addEventListener(
        // This fires when we make a change through this hook.
        "storage",
        () => {
          onStoreChange();
        },
        { signal: controller.signal },
      );

      return () => {
        controller.abort();
      };
    },
    () => {
      return localStorage.getItem(key);
    },
    () => {
      return null;
    },
  );

  const setValue = useCallback(
    (value: string) => {
      localStorage.setItem(key, value);

      storageEventTarget.dispatchEvent(new Event("storage"));
    },
    [key, storageEventTarget],
  );

  return [value, setValue] as const;
}
```

**The `storage` event on the window will only fire if change was made by another window or tab**. Calling `setItem` in the same window that is listening to the event will not trigger this event. When testing, I also observed that **the `storage` event fires when you change the value directly through the web console**.

So, to cover all possible cases, we also create the `EventTarget` we can send events to. Depending on the use-case, you could even make this `EventTarget` a global variable.

**Notice the usage of `AbortController`**. I wager, the most well-known use-case for it is to cancel `fetch` calls, but **you can use it to cancel multiple listeners via single `.abort` call**.

#### Preventing hydration mismatches

Let us say that the following component is server-side rendered. Can you spot the issue?

```jsx
function Component(event) {
  const lastUpdated = getLastUpdated();
  return <span>{lastUpdated.toLocaleDateString()}</span>;
}
```

The **output of this component will, most likely, be different from on the client**. Unless the client resides in the
same timezone as the server, the output will be different due to timezones.

I've seen many ways developer "fix" this issue ranging from using `useEffect` to the `supressHydrationWarning` attribute
on the element.

It turns out, that **the `useSyncExternalStore` is quite useful in this situation**.

```jsx
function Component(event) {
  const lastUpdated = getLastUpdated();
  const date = useSyncExternalStore(
    () => {},
    // on the client
    lastUpdated.toLocaleDateString(),
    // on the server
    null,
  );

  if (!date) {
    return null;
  }

  return <span>{lastUpdated.toLocaleDateString()}</span>;
}
```

This is an alternative to the state + `useEffect` solution. Every time I do NOT have to write `useEffect` I deem a
situation pure win.

### At odds with concurrent features

Here is the sad part: the updates you trigger via the `useSyncExternalStore` will not work with `useTransition` and will
**cause React to bail-out out of the concurrent features**.

**The only way, at the time of writing this, to hold state and make it work with concurrent features is to use `useState` and `useReducer`**.

- [Here is Tanner talking about reactivity and concurrent features](https://twitter.com/tannerlinsley/status/1732474127712481371)
- [Here is the creator of `zustand` talking about the de-opt behavior of `useSyncExternalStore`](https://blog.axlight.com/posts/why-use-sync-external-store-is-not-used-in-jotai/)

It seems like we cannot have the cake and eat it too. At least not now. I wonder how this discussion/issue will progress as larger community is relying more and more on signals/fine-grain reactivity primitives.

## `useId`

Imagine you are using SSR. You create some ID in the body of the component and pass that ID as a prop. Since your
component will be invoked twice (once on a server, once on a client), **you will most likely face hydration mismatches
due to different ID values on the server and the client**.

The following is the example usage of the `useId` hook.

```jsx
function Checkbox() {
  const id = useId();
  return (
    <>
      <label htmlFor={id}>Do you like React?</label>
      <input type="checkbox" name="react" id={id} />
    </>
  );
}
```

### How do they maintain the "stability"

The million-dollar question is: how the hell do they maintain the stability of the ID between SSR render and hydration.
The API is designed to be called inside the component body, so it must be called twice and return the same unique value.

From what I was able to gather in [this PR](https://github.com/facebook/react/pull/22644), the **`useId` uses the
components tree position (which should not change between SSR and hydration), to generate a stable identifier**.
Literally 200 IQ move.

## `useFormState`

> For more information, check [this great blog post](https://www.epicreact.dev/react-forms).

The `useFormState` is a handy hook that exposes the `actions` submit value as well as the function to call submit the
form.

```tsx
const [submitValue, action] = useFormState(serverOrClientAction, initialValue, permaLink)

<form action = { action }></form>
```

**Note that the `action` could be used in either _client action_ or _server action_ context**.

You might also think that one could use the `action` in any kind of function and not only in the context of forms. **That is the case – you can invoke it manually as long as you have `formData` object handy**

```tsx
<button
  type="button"
  onClick={() => {
    const formData = new FormData();
    formData.set("text", "value");
    dispatch(formData);
  }}
>
  Submit
</button>
```

Whether you should do that is another matter entirely. I think it would be better to have a form with a submit button here.

### The `permalink` prop

The `useFormState` accepts three arguments. The last one is an _optional_ "permalink". The **`permalink` prop is for progressive enhancement**.

If the form is submitted BEFORE React has a chance to load, the regular browser behavior will take over and submit the form with the full-page refresh functionality.

## `useFormStatus`

This one is interesting. I'm not a fan of the API as, at least to me, is a bit magical. Check this out.

```jsx
<form>
  <Button>Click me</Button>
</form>;

function Button() {
  const { pending } = useFormStatus();

  // stuff
}
```

**The most frustrating thing is that the `useFormStatus` can only be used within a `form` element**. This means that the
**`form` element magically becomes a "context provider"**. This could be quite **misleading** and also **makes it
impossible to have a submit button that lives outside of the `form` element which is a valid pattern**.

Have not found a good alternative yet,
but [others proposed theirs](https://allanlasser.com/posts/2024-01-26-avoid-using-reacts-useformstatus).

## `useInsertionEffect`

## My `fallback` prop in `React.Suspense` is not rendering

After playing around with `useDeferredValue` and `useTransition`, you might have noticed that the `fallback` prop you
pass to `React.Suspense` is only rendered "the first time" you change the application state.

This is because **in a low-priority render, React will render previously committed UI instead of discarding it in favor
of the `fallback` prop**. Imagine how annoying it would be for the UI to transition from `fallback` prop UI to the "
proper" one – so many layout shifts!

It is crucial to understand this concept as you might have been taught that React will always render the `fallback` prop
while it suspends – that is not the case!

## Offscreen component

With the `useTransition` API, you can mark a given update as a _low priority_. But what about marking the rendering of
the whole sub-tree as _low priority_? This is where we could use the `Offscreen` component.

> Keep in mind that the API is _unstable_ and will most likely change.

```jsx
function Component() {
  const [hidden, setHidden] = useState(true);

  return (
    <div>
      <button onClick={() => setHidden((_) => !_)}>toggle</button>
      <Offscreen mode={hidden ? "hidden" : "visible"}>
        <MyInitiallyHiddenUI />
      </Offscreen>
    </div>
  );
}
```

- The tree is **hidden via the CSS**.

- React **renders the tree wrapped by the `Offscreen` component with the lowest priority**.

- Allows you to **"prepare the UI" before revealing it to the user**. It works well with Suspense.

### An example use case

So far, I've seen only one use case for the `Offscreen` component (bear in mind that the API is experimental). Folks
at [replay.io](https://www.replay.io/) use the `Offscreen` API to "cache" the result of the rendering of some of the
components. Check out [this video](https://www.loom.com/share/69b18fb36bfb4ab6b70a2bda49afa499). Around the 4:45 mark,
Brian talks about using the `Offscreen` API.

```tsx
function App() {
  const [visible, setVisible] = useState(false);
  return (
    <div>
      <button onClick={() => setVisible(!visible)}>Toggle</button>
      <Offscreen mode={visible ? "visible" : "hidden"}>
        <VeryHeavyComponent id="with-offscreen" />
      </Offscreen>
      {/* Always rendered first. The render wrapped with Offscreen is marked as low prio. */}{" "}
      <VeryHeavyComponent id="pure" />
    </div>
  );
}

function fib(n) {
  if (n <= 1) return 1;
  return fib(n - 1) + fib(n - 2);
}

function VeryHeavyComponent({ id }) {
  useMemo(() => fib(40), []);
  useEffect(() => {
    console.log(`heavy mounted ${id}`);
    return () => {
      console.log(`heavy unmounted ${id}`);
    };
  }, [id]);
  return <div>Heavy!</div>;
}
```

## The different component types

React 18 introduced the `use client` directive and with it, it brought _server components_. This means that one can now
use either components with the `use client` directive or leverage the fresh _server components_ when building their UIs.
The following dives into a bit of detail about what each type of component does and what are their limitations.

### Before Client and Server Components

Before React 18 existed, you had two choices

1. Render your whole application on the client. This could lead to "white page" of first content.

2. Render your application on the server. Perform the hydration on the client. **This means that your application is
   executed twice**. Not ideal. **Frameworks like Qwik address this**.

### React Server Components

Before you learn anything about how they work, you should know that **this is a spec rather than a technology**. It just
so happens that Next.js is the first meta-framework to implement it.

Whenever I think about _RSC_ I also think about the **_React Server_ and _Components_**. We used to have only "React",
now we have _React Server_ and the _Components_ which could be either _Client_ or
_Server_. [This blog post explains this topic further](https://bobaekang.com/blog/rsc-is-react-server-plus-component).

---

Here you **stream non-interactive serialized representation of _virtual DOM_ from the server to the client**. This is **similar to `getServerSideProps` in Next.js**, but it is **different**. The main difference between _React Server Components_ and `getServerSideProps` are.

- With `getServerSideProps` you could create components that were interactive. That is not possible with _React Server
  Components_.

  - **You cannot use any React hooks with _React Server Components_**.

  - Using `getServerSideProps` is **to display a non-interactive version of the _client_ component** and then hydrate
    it for interactivity. There is **no hydration using _React Server Components_**.

  - The `getServerSideProps` could be problematic in cases where you have conditionals based on the props you pass down from the server.
    - See [this article](https://www.rexforde.com/blog/conditional-render-problem) for an example where this could be problematic.

- With _React Server Components_ you can **fetch as your render**, where the component definition is asynchronous.

- The **dependencies you use to render _Server Components_ do not add to your overall bundle**.

  - Since there is **no hydration when using SSR**, there is no need to push that code to the client.

    > See [this tweet](https://twitter.com/sebmarkbage/status/1341142110385410049).

  - The _React Server Components_ have **automatic bundle splitting**. As in you do not have to use `React.lazy` for
    code splitting.

    - If the page is not using some of the components, they will not be send to the client.

- The **_Server Components_ allow you to use native Node.js functions as they only run on the backend**.

- The **_Server Components_ are always "rendered", no matter if they are lazily loaded**.

  - This is something I learned from [this video](https://www.youtube.com/watch?v=AGAax7WzStc) and also
    from [the next.js docs](https://beta.nextjs.org/docs/optimizing/lazy-loading#example-importing-server-components).

    - According to the docs, "If you dynamically import a Server Component, only the client components that are
      children of the Server Component will be lazy loaded - not the Server Component itself.". This **is very
      important to keep in mind**.

    - As your page grows, you might need to stream more and more data. I must be missing something, but this
      strategy does not sound right to me. What if I have a huge number of components?

      - After a bit of googling, I came to a conclusion that it does not matter as you most likely will split
        everything per page. What I worry about are waterfalls while fetching client components JS.

        - This [blog from the remix team](https://remix.run/blog/react-server-components#the-react-teams-demo)
          confirms my suspicions. Unless you kickoff all the promises to start fetching and pass them down to
          components, you will get into fetch-render-fetch-render cycle that causes waterfalls.

  - Here is [an additional video on the subject of RSCs](https://portal.gitnation.org/contents/simplifying-server-components)

#### Notes from the "React Server Components with Dan Abramov, Joe Savona, and Kent C. Dodds"

[Full link to the video here](https://www.youtube.com/watch?v=h7tur48JSaw).

> It seems like it is very easy to cause waterfalls with RSCs. Since you are streaming the response, there might be a
> lot of back-and-fourth between the server and, for example, a database.

1. The **architecture of server components is separate from SSR**. You can put SSR on the edge, but run the server
   components close to the data layer.

2. Regardless of the framework, you can introduce waterfalls. The answer is observability and performance tracing.

---

1. According to Dan, [RSC automatically de-duplicate requests](https://youtu.be/h7tur48JSaw?t=2257).

   - I'm not sure that is true for _native_ RSC? It is
     a [feature of Next.js 13](https://beta.nextjs.org/docs/data-fetching/fundamentals#automatic-fetch-request-deduping).

   - And [here, Kent talks about overloading the fetch](https://youtu.be/h7tur48JSaw?t=2441). I think that he is
     referring to the Next.js 13 implementation?

     - It [turns out there is a new "fetch" exposed by React](https://youtu.be/h7tur48JSaw?t=2517). **They are
       talking about `react-fetch` package that leverages the cache API**.

---

1. As it stands now, you cannot have one server component and client component live in the same file.

   - This is not a limitation of the architecture. It is a conscious decision.

---

1. [Here Kent talks about how we should structure the application that uses server components](https://youtu.be/h7tur48JSaw?t=5145).

   - This is a shift in how we usually write React apps, where the `children` prop is used but not to that extend.

- Server components as island architecture?

#### Notes from the "Into the Depths with Server Components and Functions"

> You can [find the source here](https://www.youtube.com/watch?v=QS9yAsv1czg).

- Server Components as islands. The root is on the server. This allows for optimization on the data-serialization level.

  - The _server_ tree is continuous, while the _client_ tree can be split by the server components. This makes it hard
    to communicate between different client-components (use client context for that).

- Caching (mostly de-duping) moves from the client to the server. Of course you can cache on the client, but keep in
  mind that the client components are mixed with server components. Since the server is the root, it makes sense to
  cache on the server.

- Nested routing and the ability to deduce which data lives where allows you to skip waterfalls. You can fetch data for
  components you are about to render in parallel while rendering them.

  - That is not the case in most of the apps today. Currently we "fetch on render" most of the times.

#### Notes from "Server Components are NOT islands" part of the Ryan Carniato stream

> [Source](https://youtu.be/2zhYwg_nBqQ?t=9913).

- The static data appears twice in the downloaded HTML. Once in the script, once in the HTML markup itself.

  - Imagine a situation that the static content is _initially hidden with a client toggle_. You **would not want to
    make a server request when we toggle the content on the client**. That is why **even the "path not taken" is
    included in the initial markup**.

    - **This is why Server Components render all the "server tree", no matter if it's visible or not**.

- `Suspense` allows out-of-order streaming

  - When streaming, **you do not know which components are going to be used**. This means you have to serialize all
    the props for all the components that you steam.

    - As a solution, one might **delay streaming some content until JS loads**. This way, you will know which
      components are used, and serialize the props accordingly.

- The bottom line is that the problem space is very hard to reason about. This most likely means that we are looking at
  the problem from the wrong angle.

#### Infinite Scroll and React Server Components

I've [tried to make the _infinite scroll_ to work](https://github.com/WojciechMatuszewski/rsc-rcc-playing-around) with Apollo Client and RSCs a while back. Due to the complexity of the solution, I decided to move all the fetching to the client instead.

Recently, after [encountering this repository](https://github.com/gabrielelpidio/next-infinite-scroll-server-actions/blob/main/src/app/page.tsx#L51) I've started to think about this "problem" again.

Notice the author returning a JSX node and passing the node to `setState` call. That is a very interesting approach! The initial fetch happens on the server, but then, the client takes over.

That is somewhat similar to the approach I took while playing around with the functionality initially, but I have not though about populating the state with JSX nodes.

### React "universal components"

If you use the `use server` pragma, the component is an RSC. If you use the `use client` pragma, the component is RCC.

What if you do not use any pragma at all? **Then, we call these components _universal_**.

They are _universal_ since they can be, in theory, rendered as a RSC or as a RCC.

In practice, since you do not have to explicitly annotate the components (since they follow the "context" in which the parent was defined), they might not be so universal after all.

### React Client Components

These components are **the regular components you have been using so far**. **In the context of Next.js** these are the
components **that get executed on the server (either statically built or via SSR) and hydrated**. This means that **when
using RCCs you pay the cost of shipping JS to the client**.

**You cannot import RSC into RCC** because **RSC never "re-renders"**. Imagine a scenario where you would be able to render RSC in client components
What should happen if the state in the client component updates and you are passing this prop into the RSC? **The RSC would not update, and your app would look like it is broken!** That is why you can only import RSC in other RSC.

This "limitation" promotes composability. If you cannot import components, you have to compose them. Composability is a great way to ensure your code is scalable and responds to change in requirements well.

I also like to think about this restriction in terms of **_owner_ and _parent_ components**. The **RCC can be _parent_ of the RSC but not the _owner_ of RSC**.

> You can read more about the _parent_ and _owner_ relationship [here](https://reacttraining.com/blog/react-owner-components).

#### What does 'use client' do

> Taking notes while reading [this blog post](https://overreacted.io/what-does-use-client-do/).

- **Think of the `use client` and `use server` pragmas as opening a "door" to another "side" of the same program spanning two environments**.

  - You have the "backend" side of the program, and the "frontend" side of the program. Still, it is _the same program_.

  - The **`use client` and `use server` pragmas are _NOT_ about "marking" code as being on the _client_ and on the _server_**.

    - They are about opening "doors" to a different environment.

- When you use `use server`, you mark all the exports in a file as "callable" from the frontend.

- When you use `use client`, you mark all the exports in a file as "renderable" from the backend.

All of this is powered by abstractions on the module level. The `use server` will create necessary HTTP endpoints. The `use client` will produce the necessary references.

When calling a _server function_ from the frontend environment, you implicitly make an HTTP call, but you preserve all the nice things about "just calling a function":

1. You can "find references" to it.

2. The input arguments are typed (if you use TypeScript).

The same can be said about "rendering" a _client component_ on the backend environment.

Again, notice the quotes. Bundler does A LOT of things behind the scenes to make it appear as you are "rendering" a component or "calling" a function.

#### Notes from the "Dan Abramov explores React Server Components with us!"

> You can [find the source here](https://www.youtube.com/watch?v=Fctw7WjmxpU).

- The **term server and client** is a bit **misleading**. You do **not need a server to use server components**.

  - If you do not use the server, the "server" components would be built during the app build.

    - In fact, in the video, they started with the client-only architecture.

- The **response of an RSC is like a "JSON with holes"**. These are not instructions of any kind. This is streamable
  JSON .

- The data-fetching story gets interesting when you take `Suspense` into the mix. Keep in mind that **`Suspense` now
  works on the server and with server components!**.

  - With `Suspense` you can **achieve out-of-order streaming**. This is nice as some server components might take more
    time to resolve. You would not want to wait for ALL of them to resolve before showing content.

- You **cannot import server components into client components**.

  - This does make sense. If your server component uses a node-specific API, it would explode on the client.

  - To **compose, use the `children` prop**.

- Server components allow **for automatic code splitting of client components**.

  - **The JSON data of RSC contains the location of the client components file**. If the server component does not
    include the client components, there is nothing to download.

- The **`startTransition` tells the React that it is okay for the screen to be delayed while we wait for the RSC to
  refresh**.

  - This allows you to skip the `Suspense` loading screens when the part of the tree update.

    - Dan says that the `startTransition` allows you to **wait till React has something to show**.

### The bottom line

1. The **_Server and Client Components_ do NOT replace SSR**. Keep in mind that **SSR can render HTML output of client
   components**. That is how they are implemented in Next.js.

> Client Components enable you to add client-side interactivity to your application. In Next.js, they are pre-rendered
> on the server and hydrated on the client. You can think of Client Components as how components in the Pages Router
> have
> always worked.

2. Now you **have a greater control over what runs where**. I'm not sure if that is a good thing or not. Most likely not
   since it should be an "opt-in" rather a "must-do". These concerns are addressed by frameworks like Qwik and Marko
   where the place where component executes is opaque to the developer.

3. **React can stream parent components' output before their children finish rendering**. If it was unable to do so, you
   would be blocking the rendering every time you created an async RCS.

## Notes from "RSC From Scratch"

> [Here is the link to the first entry](https://github.com/reactwg/server-components/discussions/5) in the series.

- The SSR is about sending the HTML as the initial request. The RSC is about sending serialized JSX upon subsequent
  navigations so that we can navigate without destroying the state of the application.

- While using RSCs, the navigation **will fetch, by default, only the parts that could have changed**. There is no point
  in returning the serialized JSX for the "Layout" if you know that it could not have changed.

- The RSC have a special format to them because returning the "raw JSX" is not possible and even if it would be, the "
  raw JSX" is quite large.

  - The "raw JSX" contains symbols that correspond to the element type. These get stripped when
    performing `JSON.stringify`.

> Waiting for the part 2 as the part 1 was a fascinating read.

## More about RSCs and RCCs

> Based on [this great blog post](https://demystifying-rsc.vercel.app/).

- The **SSR output of the RSCs is the HTML and the encoded _virtual DOM_**.

  - The _virtual DOM_ is needed for future updates and to ensure we can mix RCCs with RSCs.

  - The data is encoded in a "special" new format that allows streaming.

- In **Next.js, RCCs are, by default, pre-rendered on the server**. That is why you see static HTML when you view the
  page source.

  - This is the SSR mechanism that we have been using for a while now.

- **Every time you use `use client`, you tell the bundler to put the component into a separate file**.

  - Then, React can reference the file in the streaming RSC output.

- You can **control whether the RCC runs on the server or not via the `next/dynamic` and the `ssr: true/false` option**.

- The **RCC can have RSC as `children`, but keep in mind that updating props passed to RCS will NOT cause a re-render**!

- If you **import a component inside a RCC, the component becomes RCC**.

  - This means that you can skip the `use server` on some occasions, but that might lead to a mistake where you want
    the component to explicitly be a RSC, but it becomes RCC.

    - One can **use the `server-only` module** to ensure that developers do not import RSCs into RCCs by accident.

- **Asynchronous RSCs are rendered in parallel if they are on the same nesting level**.

- The `Suspense` allows streaming. This can speed up the perceived performance of the page because React will render
  something, be it the fallback, as soon as possible.

## Server actions

- At the time of writing, they are marked as _alpha_ in Next.js

- Allow you to create **ad-hoc backend endpoints** which then you can use to **submit form data or use them as RPC calls
  from the frontend**.

- While I like the premise, **the creation of ad-hoc backend endpoints scares me**.

  - People usually **ignore the fact that these could be an entry point to your system when attacked**.

  - Reading blog posts and other materials on these, **people fail to think about rate-limiting** on those endpoints.

- There is a **real danger of leaking secrets or other sensitive data** if you are not careful.

  - The framework has to serialize the underlying parameters you pass to the _server action_. If you pass a secret
    from the frontend, you have leaked it! (of course having the access to secrets on the frontend is a whole another
    discussion).

- You can either import a _server action_ into a RCC or define a _server action_ as a function in RCC.

```ts
'use client'

import { myServerAction } from './actions';

function SomeForm() {
    return (<form action = { myServerAction } >
    <label htmlFor = "name" > Name < /label>
        < input
    name = "name"
    id = "name"
    type = "text" >
        </form>)
}

// Or you could do

async function myInlineServerAction(userId: string) {
    'use server'

    assertValidUserId(userId);

    await db.get(userId)
}
```

## Client Actions

- They have **the same syntax as _server actions_, but they differ in behavior**.

  - They **do not create ad-hoc backend endpoints**.

  - They **integrate with _Suspense_ and _Error Boundaries_**.

- They are useful for streamlining the form handling and **integrating with the `useFormStatus` hook**.

## `server-only` and `client-only` packages

- These packages allow you to mark a given file to be accessible only on the client or the server.

  - This is an additional protective layer against unwanted data transition from the server and the client.

- **They work on the basis of _conditional `package.json` exports_. I find this mechanism pretty interesting**.
  - The **key to make that work is the `--conditions` flag in Node.js**.

```json5
// server-only package.json
{
  "exports": {
    ".": {
      "react-server": "./empty.js",
      "default": "./index.js"
      // this file throws an error
    }
  }
```

Now, if someone tries to use the file with `server-only` import outside the `react-server` "condition" (check out the `--conditions` Node.js flag), the bundler will throw an error! Pretty smart.
