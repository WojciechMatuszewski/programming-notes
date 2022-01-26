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

### Confusion around `useDeferredValue`

After playing around with `useDeferredValue`, I concluded that I do not understand its functionality one bit.

My main point of confusion is around how `useDeferredValue` relates to _Suspense boundraries_. To illustrate, let us look at an example.

```tsx
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

The **`RenderPokemon` fetches data and suspends**. It uses the `deferredValue` as a prop to fetch a pokemon.
What confuses me here is the **behavior of the `div` that has the `style` prop defined**. It lives outside of the _Suspense boundary_, but, despite that, it **holds onto the "stale" value just as if it was wrapped with `Suspense`**.

If you click on the button several times, the `isLoading` prop will be `false` after a couple of milliseconds, **but the `div` will still have the `opacity` set to `0.4`**. The **`opacity` will not change until the `RenderPokemon` component catches up**.

**It seems like passing the `deferredValue` as a prop to a component that suspends "entangles" the `deferredValue` with the "suspense context" of that component**.

I found a clue regarding the behavior in the [_React Hooks in Action_ book](https://livebook.manning.com/book/react-hooks-in-action/chapter-13/59).

> If React can successfully render a new UI with the new value, and no children suspend or delay rendering, the hook returns the new value, and React updates the UI. If the new value causes React to wait for a process to complete before finishing rendering, the hook returns the old value, and React displays the UI with the old value (while working in memory on the UI with the new value).

I have no ide why is that.

## `useSyncExternalStore`

It seems like the `useSyncExternalStore` is meant to be a drop-in replacement for _subscription-like_ hooks. The idea is to make sure tearing never happens.

I did not find any concrete examples while reading the [discussion about the API](https://github.com/reactwg/react-18/discussions/86). A great excuse to dive into writing my own!
