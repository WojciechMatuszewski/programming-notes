# React 18

Me learning about new features of React 18.

## The problem with tearing

Imagine you have a global store. With _React_ 17 and down, you can freely use the store without any issues.
You update to _React_ 18 and observe a weird behavior, where tree parts have different store values. What gives?

The answer lies in the fact that _React_ 18 uses _concurrent rendering_ technique. The technique is about yielding to
the browser – i.e., "pausing" _React_ rendering. **If an update sneaks in between this "pause," one part of the tree
might show different values!**

### Note on concurrent rendering

According to [this article](https://vercel.com/blog/how-react-18-improves-application-performance), **React will yield
back to the main thread every 5 MS to check if there is not any more important task to do**. This is quite interesting,
they must use some kind of timer to do the yielding.

### Why is it not a problem in _React_ 17?

_React_ 17 is synchronous. There is no way for an update to sneak in when _React_ yields to the browser because no
yielding occurs.
That is why you have not observed this behavior yet.

## `startTransition` API

The `startTransition` API is meant to be used for **updates that can be deferred**. The most important thing to note
here is that
the **callback runs synchronously, but the state update it causes is treated as low priority**.

It seems to me like `startTransition` API should most likely be used for expensive, **local** computations that are
not "important", i.e are not user interactions.

```ts
const [startTransition, isPending] = useTransition();

return (
    <button
        onClick = {()
=>
startTransition(() => {
    // states updates here
})
}
>
</button>
)
;
```

When you use `startTransition`, the **React will prepare a new tree in the background**. Once that tree is finished
rendering, the result can be committed into the DOM.

The **`startTransition` API will not help you in the case of CPU-heavy operations**. If the main thread is blocked, then
it will be blocked, regardless if you wrap the computation with `startTransition` or not.

### `startTransition` and the `hydrateRoot` API

React 18 changed how the entrypoint to the application looks like. If your application uses SSR, now you have to
use `hydrateRoot` instead of `hydrate`. For CSR, you have to switch from `render` to `root.render`. **But did you know
there are taggable benefits in wrapping the `hydrateRoot` with `startTransition`?**

The `startTransition` marks the update as non-blocking. This is quite important as it **allows the browser to more
efficiently manage resources**. If your hydration takes a long time and is resource intensive, the website will not
longer be "locked" or "frozen".

```tsx
startTransition(() => {
    hydrateRoot(...)
})
```

[I've noticed the usage of `startTransition` in one of the the Epic Web Dev codebases](https://github.com/epicweb-dev/web-auth/blob/143e4eea6451397094cc48fa49eb6d0a2ff00fcb/exercises/01.cookies/01.problem.fetcher/app/entry.client.tsx#L5). [See official `remix` docs](https://remix.run/docs/en/main/file-conventions/entry.client)
and also [this twitter thread](https://twitter.com/dan_abramov/status/1567852606642348032).

[Next.js also does this](https://github.com/vercel/next.js/blob/90bfbe72bb79a0f6951c9e0eae24d6aa8a6c342d/packages/next/src/client/app-index.tsx#L308)
which would confirm that this is a performance
optimization, [like this post on GitHub](https://github.com/vercel/next.js/discussions/40691).

Sadly, the official React documentation does not mention using `startTransition` with `hydrateRoot`.

### Regarding the state updates

Remember that with _React_ 18, the state updates are batched together. In previous versions in _React_, this was not
necessarily always the case.
The batching of state updates also applies to the callback of the `startTransition` function.

### Behavior in the context of network requests

While many use-cases for the `startTransition` go beyond data fetching, the API can still enhance how users perceive
loaders in the context of network requests.

You are probably familiar with the issue where a loader is shown only for a split second. This creates not-so-great
experiences for the user.

We can wrap the code that updates the "network resource" we are currently working on with `startTransition` to combat
this issue. Doing so will instruct `React` to "defer" state updates until the "network resource" is ready.

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
            { pokemon.name }
        < br / >
        <button
            onClick = {()
=>
    {
        // Defer the state update and the suspense placeholder till either this resource is "ready" or some time passed.
        startTransition(() => {
            setPokemonResource(suspensify(() => fetchPokemon(pokemon.id + 1)));
        });
    }
}
>
    Fetch
    < /button>
    < /div>
)
    ;
}
```

### Confusion around `startTransition`

TODO: write about the fact that the callback passed to `startTransition` seem to be invoked multiple times(?).

- Tested on production build as well.

### The problem with `startTransition`

The `startTransition` API is not flexible.

- If used, the child components automatically have to opt into the `concurrent` behaviors.
- Must be used where the state is set. The `startTransition` callback has to contain a state update. This will most
  likely result in prop drilling.

There is one API that solves those issues. Enter the `useDeferredValue`.

## `useDeferredValue`

This API is meant to be used where the semantics of `useTransition` does not make sense. For example, where we do not
have access to the function that updates the state.

```ts
const deferredState = useDeferredValue(expensiveState);

return <Component state = { deferredState }
/>;
```

Here, the `deferredState` is a piece of state that might or might not cause a re-render. By wrapping
the `expensiveState` in the `useDeferredValue`, we tell _React_ to postpone updates to components that take this state
if needed.

The **`useDeferredValue` will not help you if the computation in React components is very CPU-heavy**. It should be *
*used when the rendering takes a lot of time, not computing values**.

### How the `useDeferredValue` works

The behavior of the `useDeferredValue` can be confusing. How come the hook returns the "stale" value for some renders
while returning the "up-to-date" version for others (see the example below)?

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

To understand how `useDeferredValue` works, we must understand one of the following: **React can now mark a render as "
low priority" and return a "stale" value for that render for a given hook. In this case, the `useDeferredValue` hook**.

- The `setValue` update is a high-priority one.

- The button text updates and the **`useDeferredValue` returns a "stale" value of the initial value (1)**.

- React "remembers" that the deferred value will need to transition to `1` in a later, low-priority render.

- The `opacity` is applied as the `deferredValue` is NOT equal to `value`.

- React has nothing better to do, so it works on the deferred update.

    - The `value` is set to two (after the first update).

    - The `useDeferredValue` returns two (just like the `value`) in this render.

    - Since we do not have results for `RenderPokemon` with `id` of 2 yet, **React suspens**. **Usually, this would
      cause the `fallback` to render, but since we are in the low-priority render, React can keep the previously
      committed result visible**.

- React commits the result.

I think [this GitHub comment](https://github.com/reactwg/react-18/discussions/129#discussioncomment-2440646) is the best
explanation of this feature one can ever get.

## `flushSync` (from `react-dom`)

Most of the `setState` calls are queued. For example

```jsx
<button
    onClick={() => {
        setName("foo");
        setNumber(1);
    }}
>
    Click me
</button>
```

The two `setXX` calls in the `onClick` handler will be queued together and processed together resulting in only one
re-render pass. **While this behavior is mostly what you want, sometimes it makes certain things hard to do. One of them
being focus management**.

Imagine a scenario where you are flipping between `input` and a `button`. You want to focus the input when you click the
button and vice-versa. **However, you do not want to focus the button when the `blur` event fires on the `input`**. This
could be achieve by adding additional state to the application. You would track the _last action_ and then
in `layoutEffect` conditionally call the `.focus` on the right element reference.

Would it be nice to get rid of the `layoutEffect` altogether? We have to use it, because the changes
to `setEditing(true/false)` are reflected asynchronously.

```jsx
<button
    onClick={() => {
        setEditing(true);
        // You cannot call the `focus` here as the UI has not been updated yet.
    }}
>
    Edit
</button>
```

**The `flushSync` API allows us to achieve just that**. If we wrap the `setEditing` in the `flushSync`, we tell React to
**update the UI synchronously**.

```jsx
<button
    onClick={() => {
        flushSync(() => {
            setEditing(true);
        });
        // You CAN call the `focus` here. The UI was updated.
        inputRef.current.focus();
    }}
>
    Edit
</button>
```

**Wrapping state updates within `flushSync` is a _deoptimization_**. Since the update cannot be interrupted as it is
synchronous, **React will not perform transitions**. In most cases that is a big deal, but **it is completely fine if
the state update only affects a small portion of the tree**. In our case, we only are flipping between an `input` and
a `button` HTML tags so the _deoptimization_ introduced by `flushSync` is not a big deal.

Check
out [this great tweet explaining the API based on the example I alluded to above](https://twitter.com/ryanflorence/status/1722358755499913582).

## `useSyncExternalStore`

It seems like the `useSyncExternalStore` is meant to be a drop-in replacement for _subscription-like_ hooks. The idea is
to make sure tearing never happens. Let us write `useIntervalHook` that utilizes the `useSyncExternalStore`.

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

The `now` is the _external store_ value. The `setInterval` simulates changes. The `useSyncExternalStore` is a bridge
between the _module scope_ and _React_ rendering lifecycle.

---

You might wonder why the `notify` function is not taking any parameters? Would not that be more straightforward?
Instead, we have to create the `onSnapshot` function. The answer to this question lies in understanding how rendering
works in React 18.

Before React 18, the rendering was synchronous. If React started rendering the tree, it had to finish in one go. With
React 18, that is no longer the case – the rendering is interruptible.

Interruptible rendering means that, in extreme cases, if not taken into account, React could render part of your tree
with state X and part of the tree with state Y (the update of the state happened in-between the interruption). To
prevent such occurrences, as they relate to external stores, **instead of using the "live store value", React takes
the "snapshot" value and performs the rendering cycle using that particular value for the whole process, even if it is
interrupted**.

### The usefulness of `useSyncExternalStore`

#### Subscribing to events

It turns out the `useSyncExternalStore` hook is useful in the global context and not only for a library maintainers. In
particular, [this blog post](https://thisweekinreact.com/articles/useSyncExternalStore-the-underrated-react-api#link2)
has two examples which really speak to me.

Here is one for the scroll position state.

```jsx
import { useSyncExternalStore } from "react";

const useOptimizedScroll = (selector = () => null) => {
    const subscribe = (notify) => {
        window.addEventListener("scroll", notify);
        return () => {
            window.removeEventListener("scroll", notify);
        };
    };

    const getSnapshot = () => {
        /**
         * If the selector returns the same value multiple times, the React will NOT update the subscriber.
         * Only unique values count.
         */
        return selector(window.scrollY);
    };

    return useSyncExternalStore(subscribe, getSnapshot);
};

function App() {
    const scrollPosition = useOptimizedScroll((value) => {
        return Math.floor(value / 100) * 100;
    });

    return (
        <div style={{ height: "300vh" }}>
            <div style={{ position: "fixed" }}>{scrollPosition}</div>
        </div>
    );
}

export default App;
```

The alternative being using `useRef` and `useState`. I would say the `useSyncExternalState` version is much easier to
reason about (especially since the pub-sub model is so widely used).

#### Preventing hydration mismatches

Let us say that the following component is server-side rendered. Can you spot the issue?

```jsx
function Component(event) {
    const lastUpdated = getLastUpdated()
    return <span>{lastUpdated.toLocaleDateString()}</span>
}
```

The **output of this component will, most likely, be different from on the client**. Unless the client resides in the
same timezone as the server, the output will be different due to timezones.

I've seen many ways developer "fix" this issue ranging from using `useEffect` to the `supressHydrationWarning` attribute
on the element.

It turns out, that **the `useSyncExternalStore` is quite useful in this situation**.

```jsx
function Component(event) {
    const lastUpdated = getLastUpdated()
    const date = useSyncExternalStore(
        () => {
        },
        // on the client
        lastUpdated.toLocaleDateString(),
        // on the server
        null
    )

    if (!date) {
        return null
    }

    return (
        <span>{lastUpdated.toLocaleDateString()}</span>
    )
}
```

This is an alternative to the state + `useEffect` solution. Every time I do NOT have to write `useEffect` I deem a
situation pure win.

### At odds with concurrent features

Here is the sad part: the updates you trigger via the `useSyncExternalStore` will not work with `useTransition` and will
**cause React to bail-out out of the concurrent features**. **The only way, at the time of writing this, to hold state
and make it work with concurrent features is to use `useState` and `useReducer`**.

- [Here is Tanner talking about reactivity and concurrent features](https://twitter.com/tannerlinsley/status/1732474127712481371)
- [Here is the creator of `zustand` talking about the de-opt behavior of `useSyncExternalStore`](https://blog.axlight.com/posts/why-use-sync-external-store-is-not-used-in-jotai/)

It seems like we cannot have the cake and eat it too. At least not now. I wonder how this discussion/issue will progress
as larger community is relying more and more on signals/fine-grain reactivity primitives.

## `useId`

Imagine you are using SSR. You create some ID in the body of the component and pass that ID as a prop. Since your
component will be invoked twice (once on a server, once on a client), **you will most likely face hydration mismatches
due to different ID values on the server and the client**.

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
)
    ;
```

### How do they maintain the "stability"

The million-dollar question is: how the hell do they maintain the stability of the ID between SSR render and hydration.
The API is designed to be called inside the component body, so it must be called twice and return the same unique value.

From what I was able to gather in [this PR](https://github.com/facebook/react/pull/22644), the **`useId` uses the
components tree position (which should not change between SSR and hydration), to generate a stable identifier**.
Literally 200 IQ move.

## `useFormState`

The `useFormState` is a handy hook that exposes the `actions` submit value as well as the function to call submit the
form.

```ts
const [submitValue, action] = useFormState(serverOrClientAction, initialValue)

< form
action = { action } > </form>
```

**Note that the `action` could be used in either _client action_ or _server action_ context**.

You might also think that one could use the `action` in any kind of function and not only in the context of forms. *
*That is the case – you can invoke it manually as long as you have `formData` object handy**

```tsx
<button
    type="button"
    onClick={() => {
        const formData = new FormData();
        formData.set("text", "value");
        dispatch(formData);
    }}
>
    Submit
</button>
```

Whether you should do that is another matter entirely. I think it would be better to have a form with a submit button
here.

## `useFormStatus`

This one is interesting. I'm not a fan of the API as, at least to me, is a bit magical. Check this out.

```jsx
<form>
    <Button>Click me</Button>
</form>;

function Button() {
    const { pending } = useFormStatus();

    // stuff
}
```

**The most frustrating thing is that the `useFormStatus` can only be used within a `form` element**. This means that the
**`form` element magically becomes a "context provider"**. This could be quite **misleading** and also **makes it
impossible to have a submit button that lives outside of the `form` element which is a valid pattern**.

Have not found a good alternative yet,
but [others proposed theirs](https://allanlasser.com/posts/2024-01-26-avoid-using-reacts-useformstatus).

## `useInsertionEffect`

## My `fallback` prop in `React.Suspense` is not rendering

After playing around with `useDeferredValue` and `useTransition`, you might have noticed that the `fallback` prop you
pass to `React.Suspense` is only rendered "the first time" you change the application state.

This is because **in a low-priority render, React will render previously committed UI instead of discarding it in favor
of the `fallback` prop**. Imagine how annoying it would be for the UI to transition from `fallback` prop UI to the "
proper" one – so many layout shifts!

It is crucial to understand this concept as you might have been taught that React will always render the `fallback` prop
while it suspends – that is not the case!

## Offscreen component

With the `useTransition` API, you can mark a given update as a _low priority_. But what about marking the rendering of
the whole sub-tree as _low priority_? This is where we could use the `Offscreen` component.

> Keep in mind that the API is _unstable_ and will most likely change.

```jsx
function Component() {
    const [hidden, setHidden] = useState(true);

    return (
        <div>
            <button onClick={() => setHidden((_) => !_)}>toggle</button>
            <Offscreen mode={hidden ? "hidden" : "visible"}>
                <MyInitiallyHiddenUI />
            </Offscreen>
        </div>
    );
}
```

- The tree is **hidden via the CSS**.

- React **renders the tree wrapped by the `Offscreen` component with the lowest priority**.

- Allows you to **"prepare the UI" before revealing it to the user**. It works well with Suspense.

### An example use case

So far, I've seen only one use case for the `Offscreen` component (bear in mind that the API is experimental). Folks
at [replay.io](https://www.replay.io/) use the `Offscreen` API to "cache" the result of the rendering of some of the
components. Check out [this video](https://www.loom.com/share/69b18fb36bfb4ab6b70a2bda49afa499). Around the 4:45 mark,
Brian talks about using the `Offscreen` API.

```tsx
function App() {
    const [visible, setVisible] = useState(false);
    return (
        <div>
            <button onClick={() => setVisible(!visible)}>Toggle</button>
            <Offscreen mode={visible ? "visible" : "hidden"}>
                <VeryHeavyComponent id="with-offscreen" />
            </Offscreen>
            {/* Always rendered first. The render wrapped with Offscreen is marked as low prio. */}{" "}
            <VeryHeavyComponent id="pure" />
        </div>
    );
}

function fib(n) {
    if (n <= 1) return 1;
    return fib(n - 1) + fib(n - 2);
}

function VeryHeavyComponent({ id }) {
    useMemo(() => fib(40), []);
    useEffect(() => {
        console.log(`heavy mounted ${id}`);
        return () => {
            console.log(`heavy unmounted ${id}`);
        };
    }, [id]);
    return <div>Heavy!</div>;
}
```

## The different component types

React 18 introduced the `use client` directive and with it, it brought _server components_. This means that one can now
use either components with the `use client` directive or leverage the fresh _server components_ when building their UIs.
The following dives into a bit of detail about what each type of component does and what are their limitations.

### Before Client and Server Components

Before React 18 existed, you had two choices

1. Render your whole application on the client. This could lead to "white page" of first content.

2. Render your application on the server. Perform the hydration on the client. **This means that your application is
   executed twice**. Not ideal. **Frameworks like Qwik address this**.

### React Server Components

Before you learn anything about how they work, you should know that **this is a spec rather than a technology**. It just
so happens that Next.js is the first meta-framework to implement it.

Whenever I think about _RSC_ I also think about the **_React Server_ and _Components_**. We used to have only "React",
now we have _React Server_ and the _Components_ which could be either _Client_ or
_Server_. [This blog post explains this topic further](https://bobaekang.com/blog/rsc-is-react-server-plus-component).

---

Here you **stream non-interactive serialized representation of _virtual DOM_ from the server to the client**. This is *
*similar to `getServerSideProps` in Next.js**, but it is **NOT the same**. The main difference between _React Server
Components_ and `getServerSideProps` are.

- With `getServerSideProps` you could create components that were interactive. That is not possible with _React Server
  Components_.

    - **You cannot use any React hooks with _React Server Components_**.

    - Using `getServerSideProps` is **to display a non-interactive version of the _client_ component** and then hydrate
      it for interactivity. There is **no hydration using _React Server Components_**.

- With _React Server Components_ you can **fetch as your render**, where the component definition is asynchronous.

- The **dependencies you use to render _Server Components_ do not add to your overall bundle**.

    - Since there is **no hydration when using SSR**, there is no need to push that code to the client.

      > See [this tweet](https://twitter.com/sebmarkbage/status/1341142110385410049).

    - The _React Server Components_ have **automatic bundle splitting**. As in you do not have to use `React.lazy` for
      code splitting.

        - If the page is not using some of the components, they will not be send to the client.

- The **_Server Components_ allow you to use native Node.js functions as they only run on the backend**.

- The **_Server Components_ are always "rendered", no matter if they are lazily loaded**.

    - This is something I learned from [this video](https://www.youtube.com/watch?v=AGAax7WzStc) and also
      from [the next.js docs](https://beta.nextjs.org/docs/optimizing/lazy-loading#example-importing-server-components).

        - According to the docs, "If you dynamically import a Server Component, only the client components that are
          children of the Server Component will be lazy loaded - not the Server Component itself.". This **is very
          important to keep in mind**.

        - As your page grows, you might need to stream more and more data. I must be missing something, but this
          strategy does not sound right to me. What if I have a huge number of components?

            - After a bit of googling, I came to a conclusion that it does not matter as you most likely will split
              everything per page. What I worry about are waterfalls while fetching client components JS.

                - This [blog from the remix team](https://remix.run/blog/react-server-components#the-react-teams-demo)
                  confirms my suspicions. Unless you kickoff all the promises to start fetching and pass them down to
                  components, you will get into fetch-render-fetch-render cycle that causes waterfalls.

    - Here
      is [an additional video on the subject of RSCs](https://portal.gitnation.org/contents/simplifying-server-components)

#### Notes from the "React Server Components with Dan Abramov, Joe Savona, and Kent C. Dodds"

[Full link to the video here](https://www.youtube.com/watch?v=h7tur48JSaw).

> It seems like it is very easy to cause waterfalls with RSCs. Since you are streaming the response, there might be a
> lot of back-and-fourth between the server and, for example, a database.

1. The **architecture of server components is separate from SSR**. You can put SSR on the edge, but run the server
   components close to the data layer.

2. Regardless of the framework, you can introduce waterfalls. The answer is observability and performance tracing.

---

1. According to Dan, [RSC automatically de-duplicate requests](https://youtu.be/h7tur48JSaw?t=2257).

    - I'm not sure that is true for _native_ RSC? It is
      a [feature of Next.js 13](https://beta.nextjs.org/docs/data-fetching/fundamentals#automatic-fetch-request-deduping).

    - And [here, Kent talks about overloading the fetch](https://youtu.be/h7tur48JSaw?t=2441). I think that he is
      referring to the Next.js 13 implementation?

        - It [turns out there is a new "fetch" exposed by React](https://youtu.be/h7tur48JSaw?t=2517). **They are
          talking about `react-fetch` package that leverages the cache API**.

---

1. As it stands now, you cannot have one server component and client component live in the same file.

    - This is not a limitation of the architecture. It is a conscious decision.

---

1. [Here Kent talks about how we should structure the application that uses server components](https://youtu.be/h7tur48JSaw?t=5145).

    - This is a shift in how we usually write React apps, where the `children` prop is used but not to that extend.

- Server components as island architecture?

#### Notes from the "Into the Depths with Server Components and Functions"

> You can [find the source here](https://www.youtube.com/watch?v=QS9yAsv1czg).

- Server Components as islands. The root is on the server. This allows for optimization on the data-serialization level.

    - The _server_ tree is continuous, while the _client_ tree can be split by the server components. This makes it hard
      to communicate between different client-components (use client context for that).

- Caching (mostly de-duping) moves from the client to the server. Of course you can cache on the client, but keep in
  mind that the client components are mixed with server components. Since the server is the root, it makes sense to
  cache on the server.

- Nested routing and the ability to deduce which data lives where allows you to skip waterfalls. You can fetch data for
  components you are about to render in parallel while rendering them.

    - That is not the case in most of the apps today. Currently we "fetch on render" most of the times.

#### Notes from "Server Components are NOT islands" part of the Ryan Carniato stream

> [Source](https://youtu.be/2zhYwg_nBqQ?t=9913).

- The static data appears twice in the downloaded HTML. Once in the script, once in the HTML markup itself.

    - Imagine a situation that the static content is _initially hidden with a client toggle_. You **would not want to
      make a server request when we toggle the content on the client**. That is why **even the "path not taken" is
      included in the initial markup**.

        - **This is why Server Components render all the "server tree", no matter if it's visible or not**.

- `Suspense` allows out-of-order streaming

    - When streaming, **you do not know which components are going to be used**. This means you have to serialize all
      the props for all the components that you steam.

        - As a solution, one might **delay streaming some content until JS loads**. This way, you will know which
          components are used, and serialize the props accordingly.

- The bottom line is that the problem space is very hard to reason about. This most likely means that we are looking at
  the problem from the wrong angle.

### React Client Components

These components are **the regular components you have been using so far**. **In the context of Next.js** these are the
components **that get executed on the server (either statically built or via SSR) and hydrated**. This means that **when
using RCCs you pay the cost of shipping JS to the client**.

**You cannot import RSC into RCC** because **RSC never "re-renders"**. Imagine a scenario where you would be able to
render RSC in client components. What should happen if the state in the client component updates and you are passing
this prop into the RSC? **The RSC would not update, and your app would look like it is broken!** That is why you can
only import RSC in other RSC.

This "limitation" promotes composability. If you cannot import components, you have to compose them. Composability is a
great way to ensure your code is scalable and responds to change in requirements well.

#### Notes from the "Dan Abramov explores React Server Components with us!"

> You can [find the source here](https://www.youtube.com/watch?v=Fctw7WjmxpU).

- The **term server and client** is a bit **misleading**. You do **not need a server to use server components**.

    - If you do not use the server, the "server" components would be built during the app build.

        - In fact, in the video, they started with the client-only architecture.

- The **response of an RSC is like a "JSON with holes"**. These are not instructions of any kind. This is streamable
  JSON .

- The data-fetching story gets interesting when you take `Suspense` into the mix. Keep in mind that **`Suspense` now
  works on the server and with server components!**.

    - With `Suspense` you can **achieve out-of-order streaming**. This is nice as some server components might take more
      time to resolve. You would not want to wait for ALL of them to resolve before showing content.

- You **cannot import server components into client components**.

    - This does make sense. If your server component uses a node-specific API, it would explode on the client.

    - To **compose, use the `children` prop**.

- Server components allow **for automatic code splitting of client components**.

    - **The JSON data of RSC contains the location of the client components file**. If the server component does not
      include the client components, there is nothing to download.

- The **`startTransition` tells the React that it is okay for the screen to be delayed while we wait for the RSC to
  refresh**.

    - This allows you to skip the `Suspense` loading screens when the part of the tree update.

        - Dan says that the `startTransition` allows you to **wait till React has something to show**.

### The bottom line

1. The **_Server and Client Components_ do NOT replace SSR**. Keep in mind that **SSR can render HTML output of client
   components**. That is how they are implemented in Next.js.

> Client Components enable you to add client-side interactivity to your application. In Next.js, they are pre-rendered
> on the server and hydrated on the client. You can think of Client Components as how components in the Pages Router
> have
> always worked.

2. Now you **have a greater control over what runs where**. I'm not sure if that is a good thing or not. Most likely not
   since it should be an "opt-in" rather a "must-do". These concerns are addressed by frameworks like Qwik and Marko
   where the place where component executes is opaque to the developer.

3. **React can stream parent components' output before their children finish rendering**. If it was unable to do so, you
   would be blocking the rendering every time you created an async RCS.

## Notes from "RSC From Scratch"

> [Here is the link to the first entry](https://github.com/reactwg/server-components/discussions/5) in the series.

- The SSR is about sending the HTML as the initial request. The RSC is about sending serialized JSX upon subsequent
  navigations so that we can navigate without destroying the state of the application.

- While using RSCs, the navigation **will fetch, by default, only the parts that could have changed**. There is no point
  in returning the serialized JSX for the "Layout" if you know that it could not have changed.

- The RSC have a special format to them because returning the "raw JSX" is not possible and even if it would be, the "
  raw JSX" is quite large.

    - The "raw JSX" contains symbols that correspond to the element type. These get stripped when
      performing `JSON.stringify`.

> Waiting for the part 2 as the part 1 was a fascinating read.

## More about RSCs and RCCs

> Based on [this great blog post](https://demystifying-rsc.vercel.app/).

- The **SSR output of the RSCs is the HTML and the encoded _virtual DOM_**.

    - The _virtual DOM_ is needed for future updates and to ensure we can mix RCCs with RSCs.

    - The data is encoded in a "special" new format that allows streaming.

- In **Next.js, RCCs are, by default, pre-rendered on the server**. That is why you see static HTML when you view the
  page source.

    - This is the SSR mechanism that we have been using for a while now.

- **Every time you use `use client`, you tell the bundler to put the component into a separate file**.

    - Then, React can reference the file in the streaming RSC output.

- You can **control whether the RCC runs on the server or not via the `next/dynamic` and the `ssr: true/false` option**.

- The **RCC can have RSC as `children`, but keep in mind that updating props passed to RCS will NOT cause a re-render**!

- If you **import a component inside a RCC, the component becomes RCC**.

    - This means that you can skip the `use server` on some occasions, but that might lead to a mistake where you want
      the component to explicitly be a RSC, but it becomes RCC.

        - One can **use the `server-only` module** to ensure that developers do not import RSCs into RCCs by accident.

- **Asynchronous RSCs are rendered in parallel if they are on the same nesting level**.

- The `Suspense` allows streaming. This can speed up the perceived performance of the page because React will render
  something, be it the fallback, as soon as possible.

## Server actions

- At the time of writing, they are marked as _alpha_ in Next.js

- Allow you to create **ad-hoc backend endpoints** which then you can use to **submit form data or use them as RPC calls
  from the frontend**.

- While I like the premise, **the creation of ad-hoc backend endpoints scares me**.

    - People usually **ignore the fact that these could be an entry point to your system when attacked**.

    - Reading blog posts and other materials on these, **people fail to think about rate-limiting** on those endpoints.

- There is a **real danger of leaking secrets or other sensitive data** if you are not careful.

    - The framework has to serialize the underlying parameters you pass to the _server action_. If you pass a secret
      from the frontend, you have leaked it! (of course having the access to secrets on the frontend is a whole another
      discussion).

- You can either import a _server action_ into a RCC or define a _server action_ as a function in RCC.

```ts
'use client'

import { myServerAction } from './actions';

function SomeForm() {
    return (<form action = { myServerAction } >
    <label htmlFor = "name" > Name < /label>
        < input
    name = "name"
    id = "name"
    type = "text" >
        </form>)
}

// Or you could do

async function myInlineServerAction(userId: string) {
    'use server'

    assertValidUserId(userId);

    await db.get(userId)
}
```

## Client Actions

- They have **the same syntax as _server actions_, but they differ in behavior**.

    - They **do not create ad-hoc backend endpoints**.

    - They **integrate with _Suspense_ and _Error Boundaries_**.

- They are useful for streamlining the form handling and **integrating with the `useFormStatus` hook**.

## `server-only` and `client-only` packages

- These packages allow you to mark a given file to be accessible only on the client or the server.

    - This is an additional protective layer against unwanted data transition from the server and the client.

- **They work on the basis of _conditional `package.json` exports_. I find this mechanism pretty interesting**.

  ```json
  // server-only package.json
  "exports": {
    ".": {
      "react-server": "./empty.js",
      "default": "./index.js" // this file throws an error
    }
  ```

  Now, if someone tries to use the file with `server-only` import outside of the `react-server` "condition" (check out
  the `--conditions` Node.js flag), the bundler will throw an error! Pretty smart.
