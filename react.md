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

## Resetting Component State with key property

`key` is used to help React track changes and basically be able to tell whats
changed in between renders.

This may seem trivial but you can actually control this behavior right? Since
you can pass a key to every component / jsx stuff we can manually re-instantiate
a given component / node.

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
  const addCounter = () => setItems(prev => [...prev, { id: prev.length }]);
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
  const incrementCount = () => setCount(prev => prev + 1);
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

would create underlying `React.Element`. When created that way, React is able to associate hooks used inside `Counter` to `React.Element`.

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

### Takeaways

So basically to avoid such errors, **which will probably cause unexpected bugs**

- never call functional components directly when they are using hooks

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

## getSnapshotBeforeUpdate

It sadly seems like this is lesser known React life-cycle method. The method itself is **very useful** in some situations. Not knowing about it may cause a lot of headaches when such situations arrive. So this method will get invoked **before the most recent `render` output**. It can return a value which will be used in `componentDidUpdate`.

Think about all the times you wanted to make chat window scroll nicely on a new message. Dealing with `refs` and `requestAnimationFrame` stuff just to make sure all scrolling logic works nicely. With this you can calculate the `scrollHeight` before an update, send that to `componentDidUpdate` and proceed with logic. SO MUCH EASIER!

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

### Suspense and data fetching

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
  r => {
    status = "success";
    result = r;
  },
  e => {
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
  }
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

## Patterns

### Value/Dispatch Provider Pattern

Sometimes the thing you want to update is complex and you decide to use reducer.
Additionally you want to control the scope of changes to a given value (maybe
state) by other components.

This is where **Value/Dispatch Provider Pattern** shines.

**You can also use state, this would work the same**.

```jsx
const StateContext = React.createContext();
const DispatchContext = React.createContext();
function reducer() {
  // your reducer
}

return function Provider({ children }) {
  const [state, dispatch] = React.useReducer(reducer, initialState);
  return (
    <StateContext.Provider value={state}>
      <DispatchContext.Provider value={dispatch}>
        {children}
      </DispatchContext.Provider>
    </StateContext.Provider>
  );
};
```

Splitting these two makes it possible to skip `useMemo` because the values are
always the same.

#### What would happen if I did not split these two

```jsx
// previous code
function Provider({ children }) {
  const [state, dispatch] = React.useReducer(reducer, initialState);
  let value = { state, dispatch };
  // Creating inline object here would be the same
  // as creating a new variable like shown above.
  return <SomeContext.Provider value={value}>{children}</SomeContext.Provider>;
}
```

Now, with every render, `value` **is different** (as in Object.is notion).

Why is a big deal?

> The propagation from Provider to its descendant consumers is not subject to
> the `shouldComponentUpdate` method. Changes are determined by comparing the
> new and old values using the same algorithm as `Object.is`

Without `useMemo` every time `Provider` rerenders, all of its consumers
rerender.

### Safe Function Call Pattern

Sometimes you have some side-effect you do not want to call on **unmounted**
component.

We can leverage `useRef` to make sure we are safe from memory leaks.

Let's use previous pattern to implement `safeDispatch`

```jsx
const canDispatch = React.useRef(true);
const safeDispatch = React.useCallback(
  (...args) => canDispatch.current && dispatch(...args),
  []
);

React.useEffect(() => () => (canDispatch.current = false), []);

// now you can pass safe dispatch as dispatch
```

Why would I use `useCallback` here?

From the docs we know that `dispatch` never changes (as in Object.is never
changes).

This means that if your component takes only `dispatch` as a prop you do not
need to wrap it in `React.memo` (or use `PureComponent`).

We want to replicate this behavior so we are using `useCallback` to make sure
that `safeDispatch` never changes (as in Object.is never changes).

### Compound Components

#### Non-flexible Compound Components

So basically you want to pass down the necessary props using `React.Children`
and `React.cloneElement` combo.

```jsx
function Counter({ children }) {
  const [count, setCount] = React.useState(0);
  return React.Children.map(children, child =>
    React.cloneElement(child, { count })
  );
}
```

This will make it so props are passed but **unless you have some `displayName`
convention, they are passed to every child**.

#### Flexible Compound Components

##### Using `static` class methods

This can be used for styling purposes

```js
const CardHeading = styled.div``;
class Card {
  static Heading = ({ children }) => <CardHeading>{children}</CardHeading>;
  render() {
    return this.props.children;
  }
}
```

Also using context API

```js
const CounterContext = React.createContext();
class Counter {
  state = {
    count: 0;
  }
  static Display = ({children}) => (
    <CounterContext.Consumer>
      {({count}) => <div>count{children}</div>}
     </CounterContext.Consumer>
  )
  render() {
    return (
      <CounterContext.Provider value = {{count: this.state.count}}>
        {this.props.children}
      </CounterContext.Provider>);
  }
}
```

##### Using hooks

This works basically the same as the class variant but you update classes to
functional components and use hooks. Of course you cannot use static properties.

### Prop Collections and Getters

This pattern was widely used with `render props` now migrated to custom hooks.

The premise is simple, supply custom props in one obj so that consumer can just
spread those without worrying about missing some props.

This can be really helpful (looking at you `react-virtualized` 😉)

#### Prop Collections

```jsx
function useInput(initialValue = undefined) {
  const [value, setValue] = React.useState(initialValue);
  function onChange(e) {
    setValue(e.currentTarget.value);
  }
  function resetValue() {
    setValue(initialValue);
  }
  return {
    // this is prop collection
    inputProps: { value, onChange: setValue },
    reset
  };
}

function Component() {
  const { inputProps } = useInput();
  // now as a consumer I do not have to worry about forgetting a prop
  return <input type="text" {...inputProps} />;
}
```

#### Prop Getters

So the `Prop Collections` are great but they are not flexible. What would happen
if user wanted to provide their own `onChange` on top of ours (this would also
allow the user to pluck off necessary props and implement sort of inversion of
control)?

Lets change an rigid object to a function which can tackle this dilemma.

```jsx
function callAll(...fns) {
  return function allCalledWith(...args) {
    fns.forEach(fn => fn && fn(...args));
  };
}

function useInput(initialValue = undefined) {
  const [value, setValue] = React.useState(initialValue);

  function onChange(e) {
    setValue(e.currentTarget.value);
  }
  function resetValue() {
    setValue(initialValue);
  }

  function getInputProps({ onChange: suppliedOnChange, ...rest } = {}) {
    return {
      onChange: callAll(suppliedOnChange, onChange),
      value,
      ...rest
    };
  }

  return {
    // this is prop collection
    inputProps: { value, onChange: setValue },
    // prop getter
    getInputProps,
    reset
  };
}

function Component() {
  const { getInputProps } = useInput();
  return (
    <input
      type="text"
      // much more flexible solution!
      {...getInputProps({ onChange: () => console.log("changed") })}
    />
  );
}
```

### State Reducer

This pattern enables us to implement _inversion of control_. You are allowing
the consumer to pluck into your internal logic.

```jsx
const callAll = (...fns) => (...args) => fns.forEach(fn => fn && fn(...args));

function toggleReducer(state, { type, initialState }) {
  switch (type) {
    case useToggle.types.toggle: {
      return { on: !state.on };
    }
    case useToggle.types.reset: {
      return initialState;
    }
    default: {
      throw new Error(`Unsupported type: ${type}`);
    }
  }
}

function useToggle({ initialOn = false, reducer = toggleReducer } = {}) {
  const { current: initialState } = React.useRef({ on: initialOn });
  const [state, dispatch] = React.useReducer(reducer, initialState);
  const { on } = state;

  const toggle = () => dispatch({ type: useToggle.types.toggle });
  const reset = () => dispatch({ type: useToggle.types.reset, initialState });
  function getTogglerProps({ onClick, ...props } = {}) {
    return {
      "aria-pressed": on,
      onClick: callAll(onClick, toggle),
      ...props
    };
  }

  function getResetterProps({ onClick, ...props } = {}) {
    return {
      onClick: callAll(onClick, reset),
      ...props
    };
  }

  return {
    on,
    reset,
    toggle,
    getTogglerProps,
    getResetterProps
  };
}
// this will be very helpful
useToggle.reducer = toggleReducer;
// this is also quality of life improvement
useToggle.types = {
  toggle: "toggle",
  reset: "reset"
};

function Usage() {
  // So here we can manipulate external logic by ourselves
  function toggleStateReducer(state, action) {
    if (action.type === useToggle.types.toggle && timesClicked >= 4) {
      return { on: state.on };
    }
    return useToggle.reducer(state, action);
  }
  const { on, getTogglerProps, getResetterProps } = useToggle({
    reducer: toggleStateReducer
  });
  // other code
}
```

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
