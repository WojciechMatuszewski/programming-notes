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
