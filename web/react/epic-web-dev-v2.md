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

TODO: should the `setOptimistic` be only applied within the `startTransition`?

Finished "Optimistic UI" -> "Optimistic UI"
