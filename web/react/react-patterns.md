# React patterns

## `isMounted` and functional components

While reading github issues I've noticed a that even Dan himself proposes this pattern:

```js
  useEffect(() => {
    let cancelled = false;
    async function() {
      await somethingAsync();
      if (!cancelled) {
        setState({/*state*/})
      }
    }
    return () => cancelled = true;
  }, [])
```

But is not it the famous `isMounted` pattern? Pretty weird right?.

Either way, in such cases you should probably use **`AbortController` API**. **It is now widely supported, go and see for yourself**.

## Value/Dispatch Provider Pattern

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

### What would happen if I did not split these two

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

Without `useMemo` every time `Provider` re-renders, all of its consumers
rerender

## Safe Function Call Pattern

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

## Context Module Functions

The name might be scary but the implementation is straightforward - really.

When you create context, especially with `useReducer` you might want to create helper functions inside the _consumer_ hook or in the _provider_ directly. This is so that you do not have to call the `dispatch` directly from within your components

```jsx
function useCounter() {
  const [state, dispatch] = React.useContext(CounterContext);

  const increment = () => dispatch("increment");
  const decrement = () => dispatch("decrement");

  return { increment, decrement };
}
```

Pretty straightforward right? Now, **what happens when that helper function is a dependency of a hook?**
Well now you have to wrap then in `useCallback`

```jsx
function useCounter() {
  // code

  const increment = React.useCallback(() => dispatch("increment"));

  // code
}
```

While this is OK for a simple _useCounter_, having a lot of memoized functions for bigger context might be a problem. And the **solution is really simple!**

Instead of creating the helper functions inside the _consumer hook_ or _provider_, just define them at file level.

```js
function increment(dispatch) {
  dispatch("increment");
}
```

**No need for memoization!**. This also brings other benefits like _tree-shaking_ and such. Overall, it's a much cleaner solution than having this function defined inside the _consumer hook_ or the _provider_ directly.

## Compound Components

### Non-flexible Compound Components

So basically you want to pass down the necessary props using `React.Children`
and `React.cloneElement` combo.

```jsx
function Counter({ children }) {
  const [count, setCount] = React.useState(0);
  return React.Children.map(children, (child) =>
    React.cloneElement(child, { count })
  );
}
```

This will make it so props are passed but **unless you have some `displayName`
convention, they are passed to every child**.

### Flexible Compound Components

#### Using `static` class methods

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

#### Using hooks

This works basically the same as the class variant but you update classes to
functional components and use hooks. Of course you cannot use static properties.

##### The `useProvider` hook

As an addition to _flexible compound components_ pattern you will often see an `useXXX` hook defined. Usually it will look something like this

```jsx
function useProvider() {
  const context = React.useContext(ProviderContext);
  if (!context)
    throw Error(
      "`useProvider` cannot be used outside of the `ProviderContext`"
    );

  return context;
}
```

You can even make it so that the hook takes parameters and perform some calculations / derives state.

What is important here to note that **`React.useContext`** will always return the **default context value** when **there is no `Provider` up in the tree**. This is why you should **always set your `Provider` default value to null / undefined**.

## Prop Collections and Getters

This pattern was widely used with `render props` now migrated to custom hooks.

The premise is simple, supply custom props in one obj so that consumer can just
spread those without worrying about missing some props.

This can be really helpful (looking at you `react-virtualized` 😉)

### Prop Collections

This is you basically creating a bag with properties , which you spread on elements. These usually fulfil very common use cases

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
    reset,
  };
}

function Component() {
  const { inputProps } = useInput();
  // now as a consumer I do not have to worry about forgetting a prop
  return <input type="text" {...inputProps} />;
}
```

Then the consumer, can just spread the `inputProps` on the component he wants to behave like a _controlled input_.

### Prop Getters

This is more flexible version on the `Prop Collections` pattern. This is where you create a _function_ instead of an object. The one benefit here is that user can specify merge their implementation with implementation of your hook.

```js
function getInputProps({ onChange: suppliedOnChange, ...rest } = {}) {
  return {
    onChange: callAll(suppliedOnChange, onChange),
    value,
    ...rest,
  };
}
```

Now instead of returning `inputProps` you will return a function. This will allow the user to do something like this

```jsx
<input
  type="text"
  // much more flexible solution!
  {...getInputProps({ onChange: () => console.log("changed") })}
/>
```

The `getInputProps` is responsible for _pseudo-composing_ the `onChange` handlers.

## State Reducer

This is an implementation of _inversion of control_ principle. Your component / hook is using a reducer for managing it's state.
You want to expose the ability for the user to integrate with your state and influence it as the user sees fit.

```jsx
const callAll = (...fns) => (...args) => fns.forEach((fn) => fn && fn(...args));

function toggleReducer(state, { type, initialState }) {
  switch (
    type
    // code
  ) {
  }
}

function useToggle({ initialOn = false, reducer = toggleReducer } = {}) {
  const { current: initialState } = React.useRef({ on: initialOn });
  const [state, dispatch] = React.useReducer(reducer, initialState);
  const { on } = state;

  const toggle = () => dispatch({ type: useToggle.types.toggle });
  const reset = () => dispatch({ type: useToggle.types.reset, initialState });
  function getTogglerProps({ onClick, ...props } = {}) {
    //code
  }

  function getResetterProps({ onClick, ...props } = {}) {
    return {
      onClick: callAll(onClick, reset),
      ...props,
    };
  }

  return {
    on,
    reset,
    toggle,
    getTogglerProps,
    getResetterProps,
  };
}
// this will be very helpful
useToggle.reducer = toggleReducer;
// this is also quality of life improvement
useToggle.types = {
  toggle: "toggle",
  reset: "reset",
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
    reducer: toggleStateReducer,
  });
  // other code
}
```

Look at `toggleStateReducer` and see how easy it is for the user to fulfil his need. You do not have to implement anything internally. Pretty great!

## Control props

You are probably aware of the notion of _controlled_ and _uncontrolled_ inputs. This is where you either pass the `value` to an input (non-empty string) and `onChange` function or you do not (there is also option for `onChange` + `readOnly` prop)

Either way, **React has to know, if either you or the framework controls the state of a given element**. This **has to be true for the whole lifecycle of that element**.

When you are writing custom reusable components, you often should do the same, as in check if the consumer is passing any properties that would make your component _controlled_. This is usually done by creating `isControlled` variable.

```jsx
function MyComponent({ on, onChange }) {
  const isControlled = on != null;

  function handleChange() {
    if (isControlled) {
      const suggestedChange = .// something you would set your state to
      return onChange(suggestedChange)
    }
  // do not call setState when you are controlled! This will result in unnecessary renders.
    setState(value)
  }
}
```

This is a powerful technique, I might go as far as argue that it is much more powerful than the `state reducer` itself. But the real benefit comes from combining those 2 patterns.

### Warnings

You might have seen them, especially while working with `input` components. They usually scream at you from switching from _controlled_ to _uncontrolled_ and vice-versa.

For your custom components, you should do the same. There is a pattern which utilizes `React.Ref` which enables you to easily check if your component is transitioning from one state to another.

```js
const { current: wasControlled } = React.useRef(isControlled);
React.useEffect(() => {
  // code
}, [isControlled]);
```

Notice that I do not use the `prevState` pattern. There is no `useEffect` which saves the current state as the `prevState`. This is because the **component state (either `controlled` or `uncontrolled`) has to be the same for the entire lifecycle of that component**. There is no need to save the previous state, our point of reference should be the initial state.
