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

### Behavior in the context of network requests

While many use-cases for the `startTransition` go beyond data fetching, the API can still enhance how users perceive loaders in the context of network requests.

You are probably familiar with the issue where a loader is shown only for a split second. This creates not-so-great experiences for the user.

We can wrap the code that updates the "network resource" we are currently working on with `startTransition` to combat this issue. Doing so will instruct `React` to "defer" state updates until the "network resource" is ready.

```ts
import { startTransition, useState } from "react";
import { suspensify } from "suspensify";

function fetchPokemon(id: number) {
  return fetch(`https://pokeapi.co/api/v2/pokemon/${id}`).then((res) =>
    res.json()
  );
}

const initialPokemonResource = suspensify(() => fetchPokemon(1));

// This component is wrapped with suspense boundary
function PokemonDetail() {
  const [pokemonResource, setPokemonResource] = useState(
    initialPokemonResource
  );

  const pokemon = pokemonResource;
  return (
    <div>
      {pokemon.name}
      <br />
      <button
        onClick={() => {
          // Defer the state update and the suspense placeholder till either this resource is "ready" or some time passed.
          startTransition(() => {
            setPokemonResource(suspensify(() => fetchPokemon(pokemon.id + 1)));
          });
        }}
      >
        Fetch
      </button>
    </div>
  );
}
```

### Confusion around `startTransition`

TODO: write about the fact that the callback passed to `startTransition` seem to be invoked multiple times(?).

- Tested on production build as well.

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

### How the `useDeferredValue` works

The behavior of the `useDeferredValue` can be confusing. How come the hook returns the "stale" value for some renders while returning the "up-to-date" version for others (see the example below)?

```jsx
function App() {
  const [value, setValue] = React.useState(1);
  const deferredValue = React.useDeferredValue(value);

  const isLoading = deferredValue !== value;
  return (
    <div>
      <button
        onClick={() => {
          setValue((v) => v + 1);
        }}
      >
        Next pokemon ({value + 1})
      </button>
      <div style={{ opacity: isLoading ? 0.4 : 1 }}>
        <React.Suspense fallback={<span>Loading...</span>}>
          <RenderPokemon id={deferredValue}></RenderPokemon>
        </React.Suspense>
      </div>
    </div>
  );
}
```

To understand how `useDeferredValue` works, we must understand one of the following: **React can now mark a render as "low priority" and return a "stale" value for that render for a given hook. In this case, the `useDeferredValue` hook**.

- The `setValue` update is a high-priority one.

- The button text updates and the **`useDeferredValue` returns a "stale" value of the initial value (1)**.

- React "remembers" that the deferred value will need to transition to `1` in a later, low-priority render.

- The `opacity` is applied as the `deferredValue` is NOT equal to `value`.

- React has nothing better to do, so it works on the deferred update.

  - The `value` is set to two (after the first update).

  - The `useDeferredValue` returns two (just like the `value`) in this render.

  - Since we do not have results for `RenderPokemon` with `id` of 2 yet, **React suspens**. **Usually, this would cause the `fallback` to render, but since we are in the low-priority render, React can keep the previously committed result visible**.

- React commits the result.

I think [this GitHub comment](https://github.com/reactwg/react-18/discussions/129#discussioncomment-2440646) is the best explanation of this feature one can ever get.

## `useSyncExternalStore`

It seems like the `useSyncExternalStore` is meant to be a drop-in replacement for _subscription-like_ hooks. The idea is to make sure tearing never happens. Let us write `useIntervalHook` that utilizes the `useSyncExternalStore`.

```jsx
let now = new Date().toISOString();
const subscribers = new Set();
setInterval(() => {
  now = new Date().toISOString();
  subscribers.forEach((notify) => notify());
}, 1000);

const onSubscribe = (notify) => {
  subscribers.add(notify);
  return () => {
    subscribers.delete(notify);
  };
};

const onSnapshot = () => now;

function App() {
  const value = useSyncExternalStore(onSubscribe, onSnapshot);
  return <div>{value}</div>;
}
```

The `now` is the _external store_ value. The `setInterval` simulates changes. The `useSyncExternalStore` is a bridge between the _module scope_ and _React_ rendering lifecycle.

---

You might wonder why the `notify` function is not taking any parameters? Would not that be more straightforward? Instead, we have to create the `onSnapshot` function. The answer to this question lies in understanding how rendering works in React 18.

Before React 18, the rendering was synchronous. If React started rendering the tree, it had to finish in one go. With React 18, that is no longer the case – the rendering is interruptable.

Interruptable rendering means that, in extreme cases, if not taken into account, React could render part of your tree with state X and part of the tree with state Y (the update of the state happened in-between the interruption). To prevent such occurrences, as they relate to external stores, **instead of using the "live store value", React takes the "snapshot" value and performs the rendering cycle using that particular value for the whole process, even if it is interrupted**.

## `useId`

Imagine you are using SSR. You create some ID in the body of the component and pass that ID as a prop. Since your component will be invoked twice (once on a server, once on a client), **you will most likely face hydration mismatches due to different ID values on the server and the client**.

The following is the example usage of the `useId` hook.

```jsx
function Checkbox() {
  const id = useId();
  return (
    <>
      <label htmlFor={id}>Do you like React?</label>
      <input type="checkbox" name="react" id={id} />
    </>
  );
);
```

### How do they maintain the "stability"

The million-dollar question is: how the hell do they maintain the stability of the ID between SSR render and hydration. The API is designed to be called inside the component body, so it must be called twice and return the same unique value.

From what I was able to gather in [this PR](https://github.com/facebook/react/pull/22644), the **`useId` uses the components tree position (which should not change between SSR and hydration, to generate a stable identifier**. Literally 200 IQ move.

## `useInsertionEffect`

## My `fallback` prop in `React.Suspense` is not rendering

After playing around with `useDeferredValue` and `useTransition`, you might have noticed that the `fallback` prop you pass to `React.Suspense` is only rendered "the first time" you change the application state.

This is because **in a low-priority render, React will render previously committed UI instead of discarding it in favor of the `fallback` prop**. Imagine how annoying it would be for the UI to transition from `fallback` prop UI to the "proper" one – so many layout shifts!

It is crucial to understand this concept as you might have been taught that React will always render the `fallback` prop while it suspends – that is not the case!
