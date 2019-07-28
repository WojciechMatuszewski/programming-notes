# React Stuff

## Patterns

### Value/Dispatch Provider Pattern

Sometimes the thing you want to update is complex and you decide to use reducer. Additionally you want to control the scope of changes to a given value (maybe state) by other components.

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

Splitting these two makes it possible to skip `useMemo` because the values are always the same.

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

> The propagation from Provider to its descendant consumers is not subject to the `shouldComponentUpdate` method.
> Changes are determined by comparing the new and old values using the same algorithm as `Object.is`

Without `useMemo` every time `Provider` rerenders, all of its consumers rerender.

### Safe Function Call Pattern

Sometimes you have some side-effect you do not want to call on **unmounted** component.

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

From the docs we know that `dispatch` never changes (as in Object.is never changes).

This means that if your component takes only `dispatch` as a prop you do not need to wrap it in `React.memo` (or use `PureComponent`).

We want to replicate this behavior so we are using `useCallback` to make sure that `safeDispatch` never changes (as in Object.is never changes).
