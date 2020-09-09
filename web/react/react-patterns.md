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
