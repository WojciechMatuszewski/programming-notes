# React Stuff

## Future

### Moving away from `.defaultProps`

Looking at the discussions on git about React 16.9 there is a proposal to move away from `.defaultProps`, at lest on functional components.

The reason is pretty justified in my opinion, just look at the code snippet here:

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

The problem is that _resolving default props_ happens on every `React.createElement` call. This may seem insignificant but image how many times this function gets invoked.

The answer for this problem would be to use default properties available in ES6, but the are also problems with that solution.

Bundle size bloating may occur when using destructuring and default values (especially object ones) when transpiling to other versions.

React team still have to asses different choices but one is almost certain that `.defaultProps` will go out of favour very soon.

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

### Compound Components

#### Non-flexible Compound Components

So basically you want to pass down the necessary props using `React.Children` and `React.cloneElement` combo.

```jsx
function Counter({ children }) {
  const [count, setCount] = React.useState(0);
  return React.Children.map(children, child =>
    React.cloneElement(child, { count })
  );
}
```

This will make it so props are passed but **unless you have some `displayName` convention, they are passed to every child**.

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

This works basically the same as the class variant but you update classes to functional components and use hooks. Of course you cannot use static properties.

### Prop Collections and Getters

This pattern was widely used with `render props` now migrated to custom hooks.

The premise is simple, supply custom props in one obj so that consumer can just spread those without worrying about missing some props.

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

So the `Prop Collections` are great but they are not flexible.
What would happen if user wanted to provide their own `onChange` on top of ours (this would also allow the user to pluck off necessary props and implement sort of inversion of control)?

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
      {...getInputProps({ onChange: () => console.log('changed') })}
    />
  );
}
```

### State Reducer

This pattern enables us to implement _inversion of control_. You are allowing the consumer to pluck into your internal logic.

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
      'aria-pressed': on,
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
  toggle: 'toggle',
  reset: 'reset'
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
