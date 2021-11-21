# React 18

Me learning about new features of React 18.

## `startTransition` API

The `startTransition` API is meant to be used for **updates that can be deferred**. The most important thing to note here is that
the **callback runs synchronously, but the state update it causes is treated as low priority**.

It seems to me like `startTransition` API should most likely be used for expensive, **local** computations that are not "important", i.e are not user interactions.

```ts
const [startTransition, isPending] = useTransition();

return (
  <button
    onClick={() =>
      startTransition(() => {
        // states updates here
      })
    }
  ></button>
);
```

### The problem with `startTransition`

The `startTransition` API is not flexible.

- If used, the child components automatically have to opt into the `concurrent` behaviors.
- Must be used where the state is set. The `startTransition` callback has to contain a state update. This will most likely result in prop drilling.

There is one API that solves those issues. Enter the `useDeferredValue`.

## `useDeferredValue`

This API is meant to be used where the semantics of `useTransition` does not make sense. For example, where we do not have access to the function that updates the state.

```ts
const deferredState = useDeferredValue(expensiveState);

return <Component state={deferredState} />;
```

Here, the `deferredState` is a piece of state that might or might not cause a re-render.
By wrapping the `expensiveState` in the `useDeferredValue`, we tell _React_ to postpone updates to components that take this state if needed.

Pretty good API if you ask me.
