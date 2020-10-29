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
