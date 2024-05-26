# Epic Web Dev V2 notes

## React Suspense

### Data Fetching

- When fetching data, we not only have to think about _how_ to do it, but also what is the user experience _when_ we do it.

  - Do we show some kind of loader?

  - What about errors?

  - What if the network request resolves very fast, and the users sees a "flash" of a loading screen?

  These concerns are not trivial to solve. Luckily for us, React exposes all the necessary abstractions for your data fetching needs!

  Enter the `use` hook, `Suspense` and `ErrorBoundary`!

- The `use` hook takes in a _Promise_ – **any _Promise_ will do**, and will _suspend_ the component until that promise is either resolved or rejected.

  - I put extra emphasis on the "any promise aspect". We often associate the `use`, or `tanstack-query` for that matter, with fetching data from a remote endpoint. While that _appears_ to be the main use case, you can resolve ANY promise using those tools.

  - The **`use` hook works by throwing unresolved _Promise_**. The `Suspense` component will catch it, and **re-render your application either when the promise resolves or rejects**.

    - You can create your own simplified `use` hook. It is NOT magic!

      ```ts
      type UsePromise<Value> = Promise<Value> & {
        status: "pending" | "fulfilled" | "rejected";
        value: Value;
        reason: unknown;
      };

      function use<TResult>(pendingPromise: Promise<TResult>) {
        const usePromise = pendingPromise as UsePromise<TResult>;

        if (usePromise.status === "rejected") {
          throw usePromise.reason;
        }

        if (usePromise.status === "pending") {
          throw usePromise;
        }

        if (usePromise.status === "fulfilled") {
          return usePromise.value;
        }

        usePromise.status = "pending";
        usePromise.then(
          (value) => {
            usePromise.value = value;
            usePromise.status = "fulfilled";
          },
          (error) => {
            usePromise.reason = error;
            usePromise.status = "rejected";
          },
        );

        throw usePromise;
      }
      ```

    The _actual_ implementation is most likely much more complex. But for learning purposes, this version of `use` will suffice.

    **Notice** that we are also throwing an error. **The `ErrorBoundary` will catch any errors and show your "fallback" component**.

    This is why you most likely want to use `Suspense` and `ErrorBoundary` components together!

- Be aware of the **lifecycle of the `use` hook and how it influences the _Promise_ you pass to it**.

  - Since React will re-render the component that _suspended_ AFTER the promise is settled (rejected or resolved), **the _Promise_ you pass to `use` has to be "stable" – cached or created _outside_ of the component**. Otherwise, you will fall into an infinite rendering loop – React will re-render the component which then creates a new promise causing the component to _suspend_ again.

### Dynamic Promises

- To "fix" the issue with infinite loading or _suspending_ unnecessary when using the `use` hook, you most likely want to use some kind of cache.

  - **Using a cache or "stable" promises is especially crucial when the _Promise_ is dependant on the user input**.

  ```ts
  const shipCache = new Map<string, Promise<Ship>>();

  export function getShip(name: string, delay?: number) {
    const pendingPromise = shipCache.get(name) ?? getShipImpl(name, delay);
    shipCache.set(name, pendingPromise);

    return pendingPromise;
  }

  // And here is how you would use it
  const ship = use(getShip(shipName));
  ```

  **Notice that the `getShip` is a _synchronous_ function**. This detail is important.

  If it were an _asynchronous_ function, **we would be effectively creating a new _Promise_ every time we invoke the `getShip`**. Despite returning a _Promise_ from the cache, **it is the "outer" promise that would be brand new**. As such **we would be suspending every time we call this function**.

  When working on the exercise for this section, I was puzzled to find my solution not working, only to discover this was the issue!

- To **preserve the old UI, while we fetch data for the new one** use the `useTransition` hook, or the `startTransition` function.

  - The usage is dependant on the scenario. In some cases, you can't use the `useTransition` because you are not "in a component".

  - The `startTransition` would be useful for "library" code that deals with data fetching. This "library code" would then be used in a given React component.

- There is also the **problem of "pending UI" flash**.

  - As good as the `useTransition` is, what if the `isPending` boolean is `true` only for a couple of milliseconds? In that case, the UX suffers – we are kind of back to square one.

    - **The solution is to show the loading indicator for a minimum amount of time after the _Promise_ fails to resolve after given period**.

      For example: show the loading indicator for MINIMUM of 200ms if the _Promise_ did not resolve after 500ms.

      It does not matter if the _Promise_ resolves after 501ms or 700ms. We will ALWAYS be showing that loading indicator for a MINIMUM of 200ms.

### Optimistic UI

- So, we learned about the _transitions_ and showing the "old" UI. What if we want to show the new UI, but do so _during_ the transition?

  - This is where **`useOptimistic` hook comes in**.

    - The `useOptimistic` hook allows you to **change the state within the component during the _transition_**. You can change the state, and subsequently the UI, as many times as you want!

      - When the **_transition_ is done, it will fall back to the "initial" value, which now should be updated to the "new" value, you provided the hook with**.

- The **`useOptimistic` only makes sense in the context of a _transition_**.

  - If you apply it _outside_ a transition, the **UI will change only for a very split second, and then revert back to the initial value**.

- To enhance the experience, you might want to use `useFormStatus`.

  - It is interesting that **the `form` tag now is an implicit _Context Provider_**.

    - The `useFormStatus` will "read" the nearest form status. **This only works when the form uses `action`**.

    ```tsx
    function SomeForm() {
      return (
        <form action={async () => {}}>
          <CreateButton />
        </form>
      );
    }

    function CreateButton() {
      const { pending } = useFormStatus();

      return (
        <button disabled={pending} type="submit">
          {pending ? "Creating..." : "Create"}
        </button>
      );
    }
    ```

### Suspense img

- As mentioned earlier, the `use` hook and the _Suspense_ component is not only for resolving `fetch` requests.

  - We can use them to resolve any kind of promise – including _preloading_ an image!

```ts
function preloadImage(src: string) {
  return new Promise<string>((resolve, reject) => {
    const img = new Image();
    img.src = src;

    img.onload = () => resolve(src);
    img.onerror = reject;
  });
}

// Then in the component
function Img({ src }: { src: string }) {
  // Remember about the cache!
  const loadedSrc = use(cache(preloadImage(src)));

  // return
}
```

- The biggest advantage of this approach is that you get to decide what happens when the `Img` throws.

```ts
function ShipImg(props: ComponentProps<"img">) {
  return <ErrorBoundary fallback = {/*your stuff*/}><Img {...props}/></ErrorBoundary>
}
```

- This section on the workshop touches on the important concept of **forcing the `Suspense` component to always show the fallback**.

  - **When using _transitions_, React will only show the `Suspense` fallback upon the first load**. After that, any change triggered inside a _transition_ will result in showing the "old" UI while the new one loads.

    - Sometimes, this is not something we want. In our case, we wanted to _always_ show the `fallback` prop when a new image loads.

    - **You can achieve that by using the `key` prop on the `Suspense` or the parent of the `Suspense` component**.

    ```tsx
    function ShipImg(props: React.ComponentProps<"img">) {
      return (
        // Notice the usage of the `key` prop here.
        <ErrorBoundary fallback={<img {...props} />} key={props.src}>
          <Suspense fallback={<img {...props} src={"/img/fallback-ship.png"} />}>
            <Img {...props} />
          </Suspense>
        </ErrorBoundary>
      );
    }
    ```

### Responsive

- The **`useDeferredVale` hook could be used to have certain parts of the application display "old" results while others are up-to-date**.

  - This is **similar to how `useTransition` works**, but **`useDeferredValue` is more granular**.

- The most "famous" example with `useDeferredValue` is search.

  - You want the search bar to have the freshest value – what the user typed in, but the result can lag behind.

    - If you tried to use `useTransition` for this use case, every time user typed a letter, the UI would _suspend_. Not ideal!

- The **critical point to understand** is that the `useDeferredValue` **will cause your component to render twice**.

  - The **first render** is with _deferred value_ as the "old" value.

  - The **second render** is with _deferred value_ as the "current value.

  This has **huge implications**.

  1. React will wait for any _suspended_ components before showing the new UI during the second render.

  2. If you want to use **`useDeferredValue` as optimization technique, the component you pass the _deferred value_ has to be memoized and stable between re-renders**. Why? because React renders the component twice. If it's not memoized, the first render will be slow, defeating the purpose of using `useDeferredValue`.

- React [has great documentation](https://react.dev/reference/react/useDeferredValue) about this hook.

### Optimizations

- Unfortunately, **it is very easy to create network waterfalls with `Suspense`**.

  - I observed that, while a lot of developers understand that they can `Promise.all` or `Promise.allSettled` multiple promises on the backend, the same is not true with frontend developers using React.

  ```tsx
  function Component() {
    // These will load sequentially and not in parallel.
    // We have just created a waterfall!
    const user = use(getUser());
    const post = use(getPost());

    return <div></div>;
  }
  ```

  To "solve" this issue, we could **kick-off _Promises_ BEFORE using the `use` hook**.

  ```tsx
  function Component() {
    // Start fetching the data in parallel.
    // Cache plays a key role here – we assume we can call those functions and they will always return the same promise.
    const userPromise = getUser();
    const postPromise = getPost();

    // Now suspend
    const user = use(userPromise);
    const post = use(postPromise);

    return <div></div>;
  }
  ```

  Sadly, **it is very hard to track all the data-dependencies of a given group of components – most waterfalls are not that easy to fix**.

  I would recommend **using a framework that has the ability to "see" the network requests your application is making**. This way, it would be the framework responsibility to kick-off those requests for you.

### Summary

Such an excellent workshop. It covered the following.

- The usage of the `Suspense` component and how it interplays with `useTransition`.

- The usage of the `use` hook and how to build your own, simplified version of it.

- How to utilize the `Suspense` while loading images and the "key prop trick" to force a new `Suspense` fallback.

- How we could utilize the `useDeferredValue`. First, means of rendering optimization and second, as means of keeping the UI snappy while fetching results.

- How the `Suspense` and the `use` hook could lead to _network waterfalls_ and what to do about it.

  - Here, if you have multiple, non-dependant promises, remember to trigger them BEFORE using the `use` hook.
