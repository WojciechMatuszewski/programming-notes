# React 18

Me learning about new features of React 18.

## The problem with tearing

Imagine you have a global store. With _React_ 17 and down, you can freely use the store without any issues.
You update to _React_ 18 and observe a weird behavior, where tree parts have different store values. What gives?

The answer lies in the fact that _React_ 18 uses _concurrent rendering_ technique. The technique is about yielding to the browser – i.e., "pausing" _React_ rendering. **If an update sneaks in between this "pause," one part of the tree might show different values!**

### Why is it not a problem in _React_ 17?

_React_ 17 is synchronous. There is no way for an update to sneak in when _React_ yields to the browser because no yielding occurs.
That is why you have not observed this behavior yet.

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

### Regarding the state updates

Remember that with _React_ 18, the state updates are batched together. In previous versions in _React_, this was not necessarily always the case.
The batching of state updates also applies to the callback of the `startTransition` function.

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

## `useSyncExternalStore`

It seems like the `useSyncExternalStore` is meant to be a drop-in replacement for _subscription-like_ hooks. The idea is to make sure tearing never happens.

I did not find any concrete examples while reading the [discussion about the API](https://github.com/reactwg/react-18/discussions/86). A great excuse to dive into writing my own!
