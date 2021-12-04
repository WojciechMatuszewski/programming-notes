# React gotchas

## `ErrorBoundary` will not catch asynchronous errors (and many other stuff)

So just to set the stage here. Lets say you have a simple `Component` component.

```jsx
function Component() {
  // code
}
```

This component is wrapped in an `ErrorBoundary`

```jsx
// app
<ErrorBoundary>
  <Component />
</ErrorBoundary>
```

You would expect all errors be cough by that `ErrorBoundary` right? Well it's sadly not the case.
There is a list of things that `ErrorBoundary` will not catch

- errors originating from event handlers
- asynchronous errors
- server side rendering related errors
- errors thrown in the `ErrorBoundary` itself

So if you were to write the `Component` like this

```jsx
function Component() {
  React.useEffect(() => {
    fetch("flanky-response")
      .then(handleResponse)
      .catch(() => throw new Error("boom!"));
  }, []);
}
```

You will not see the `ErrorBoundary` fallback :C

What you can do is to use `setState`s second parameter, the setter.

```jsx
const [_, setError] = React.useState();
React.useEffect(() => {
  fetch("flanky-response")
    .then(handleResponse)
    .catch(() => setError(() => throw new Error("boom!")));
}, []);
```

I'm not completely sure why it works. Maybe something with queuing? No idea. It seems like I will have to ask Dan himself one day.

## You can break the rule of hooks safely (maybe)

I want to show you one particular situation where it is safe to break the _rule of hooks_.
We will be wrapping the `useEffect` with an condition block. As scary it may sound, this will produce no bugs and is completely safe.

So here is the code

```jsx
if (process.env.NODE_ENV != "production") {
  useEffect(() => {
    // logging or anything really
  }, []);
}
```

This works because the condition will always evaluate to one value, and will be computed whenever your app starts.
This technique is very useful while creating anything related to logging, especially while building reusable components.

Say you are building a custom `Input` component.
It is assumed that whenever you pass the `value` prop, that decision has to be preserved throughout the lifecycle of a given component.
You cannot pass `value` on one render, then pass `undefined` on another.

We can produce a warning if this happens

```js
const { current: wasControlled } = React.useRef(value != undefined);
const isControlled = value != undefined;

React.useEffect(() => {
  if (wasControlled && !isControlled) {
    console.error("You are breaking the rules mate!");
  }
}, [wasControlled, isControlled]);
```

Now, we would not want this message to be produced whenever our app is running in production mode, thus we wrap it with an `if` block.

```jsx
if (process.env.NODE_ENV != "production") {
  React.useEffect(() => {
    if (wasControlled && !isControlled) {
      console.error("You are breaking the rules mate!");
    }
  }, [wasControlled, isControlled]);
}
```

## Cleaning up Refs DOM handles

Sometimes while we work with _Refs_, especially the ones that are bound to DOM elements, we
attach handlers to them. Since we probably do not want to introduce memory leaks within our apps, we also clean after ourselves.

You are, most likely, carrying out the cleanup process within the cleanup function of the `useEffect` hook.
While doing so, you might be shooting yourself in a foot - here is how.

Let's say the `useEffect` where you bind the events and are cleaning up after yourself looks as follows

```jsx
useEffect(() => {
  ref.current.addEventListener("click", ...)

  return () => {
    ref.current.removeEventListener("click", ...)
  }
}, []);
```

If you are using linting rules for hooks, you will be greeted with an warning

> [...] will likely have changed by the time this effect cleanup function runs. If this ref ...

This is completely reasonable message to throw here as **cleanup function is ran after the new view was rendered**.
This means that, exactly as the message said, the `ref` might have been mutated - could possibly be `null` since the element that _Ref_ is attached to,
might no longer exist.

I'm aware of two ways to handle this issue

### Capturing the Ref value within the closure

First solution would be to do what the linting rule is telling you to do - capture the `current` value of the _Ref_ within the cleanup function closure.

```jsx
useEffect(() => {
  const capturedRef = ref.current
  capturedRef.addEventListener("click", ...)

  return () => {
    capturedRef.removeEventListener("click", ...)
  }
}, []);
```

Here, the cleanup function will have access to the _Ref_ value that was not mutated by any changes.

### Using `useCallback` and `callback Refs`

This technique is lesser known, but very important in some situations - when we want to know the underlying _Ref_ changed (looking at you `useIntersectionObserver` hooks that are poorly written)

```jsx

const cleanupRef = useRef(() => {});

const callbackRef = useCallback((ref) => {
  if (!ref) {
    cleanupRef.current()
    cleanupRef = () => {}
    return
  }

  const listener = () => {...}
  ref.addEventListener("click", listener)
  cleanupRef.current = () => ref.removeEventListener("click", listener)
}, []);

return <div ref={callbackRef} />;
```

It is **very important** that you use **`useCallback`** here. Otherwise you might be in for an infinite loop.

As you can see, the complexity of this solution is a bit higher (at least from the _familiarity_ point of view) than the `useEffect` one.

### Bottom line

When I need to know that the underlying `ref` changed (remember that using `ref.current` on the useEffect array is pointless) I will reach out for the
_callback refs_, otherwise I'm sticking with `useEffect` way of doing things.

## `useEffect` fires before pain event

You probably heard that the `useEffect` should fire **some time** after the paint event. Here is an excerpt from the [React documentation](https://reactjs.org/docs/hooks-reference.html#useeffect).

> Unlike componentDidMount and componentDidUpdate, the function passed to useEffect fires after layout and paint, during a deferred event.

Many people stop here, but just after that, the docs state.

> Although useEffect is deferred until after the browser has painted, it’s guaranteed to fire before any new renders. React will always flush a previous render’s effects before starting a new update.

And this is where things get interesting. It means that **in theory**, the **`useEffect` can run before the paint event!**.

Let us explore how.

### State updates in `useLayoutEffect`

If you update a piece of state in `useLayoutEffect` you will trigger a render. What has the guarantee of being flushed before a new render? That is right – the `useEffect`.

```ts
const [state, setState] = React.useState(1);

React.useLayoutEffect(() => setState((state) => state + 1));
// Will run before "paint" event \/
React.useEffect(() => console.log("before Paint!"));
```

### This does not matter

In the end, this is a technical minutia rather than something you need to be aware of. I would strongly recommend you follow the guidelines of "read/mutate DOM in `useLayoutEffect`, everything else – use `useEffect`".

Remember, in the end, your code has to be straightforward. You could use the behavior I'm explaining here to merge some of your `useEffect` callbacks with the `useLayoutEffect` ones into one, but what's the point?

Keep it simple.
