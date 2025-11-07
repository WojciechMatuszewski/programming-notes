# Fem React Performance V2 2025

> [Course material](https://stevekinney.com/courses/react-performance).

- Doing stuff is more expensive than not doing stuff.

- Memoization and caching is also _doing stuff_. We just assume that memoization and caching is _less_ expensive than the "stuff" without it.

- Not doing stuff is the best optimization you can make.

- The goal of performance optimization should be to: **first eliminate the unnecessary work, then to prioritize the necessary work**.

- Commits to the DOM _can't_ be interrupted. Makes sense.

- If it's possible, _pushing state down_ can be quite effective optimization technique.

  - For example, a component that renders a form plus other components. It might be worth making sure the form component handles it's own state.

- Using the _transitions_ API is a way of telling React to schedule the work for later.

  - `useDeferredValue`

  - `startTransition` / `useTransition`

- While there is some overlap between the two, the `useDeferredValue` is quite useful when you do not control the incoming value, so `props`.

  - The `startTransition` or `useTransition` is great since it allows you to control the code-flow a bit better.

- One one about `useDeferredValue`. **React will schedule a render with the "old" and the "new"**.

  - This means that, initially, it will trigger two re-renders, where the non-deferred is the same as deferred. **This means that the component your component renders better render fast**. Otherwise you will pay a penalty of the slower render twice.

- Again and again, I'm forgetting that **`useOptimistic` only works in the context of an async**.

  - After calling the `useOptimistic` you must trigger an async transition for the UI to update.

    - The form actions trigger async transitions.
    - You can use `startTransition` with an async callback.

    ```
    addOptimistically()

    startTransition(async() => {})
    ```

## Wrapping up

While not that comprehensive, it was a good refresher.

The `useOptimistic` got me again!
