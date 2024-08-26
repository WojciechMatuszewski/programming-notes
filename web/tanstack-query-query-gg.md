# React Query course

> Going through [this course](https://ui.dev/c/react-query).

## Notes / learnings

- I'm positively surprised that the course does not present the React Query as _data-fetching library_.

  - So many developers only think about the React Query as _data-fetching library_ and, while they are partly correct, it can do so much more!

    - **You can use React Query with _anything_ that returns a promise**.

- By default, results of the queries are cached and will be _shared_ amongst queries.

  - As long as the `queryKey` matches, and you did not change the cache-related settings on a query, requests will be de-duped.

    - **Consider abstracting the invocations of `useQuery` to custom hooks**. This pattern is quite useful as the name of the hook can add semantic meaning to the React Query code inside the function!

- **It would be awesome if the `status` or `isError`/`isLoading` were _type guards_**.

  - After writing all those conditionals for `if(isError)` and `if(isLoading)` I would expect the `data` to be `T` not the `T | undefined`.

- **By default, all the data is considered _stale_ out of the box**. This means that the `staleTime` parameter is `0`.

  - If you think about it, this makes sense. The `staleTime` is very context-dependent, and should be set by developers on case-by-case basis.

- After playing around with React Query, you might notice that, by default, the `useQuery` will fire as soon as it runs. In some cases, you want to defer making the network request till some condition is met.

  - For that, use the **`enabled` prop on `useQuery`**, but **ensure you handle the _loading states_ correctly**.

    - The `status === "pending"` only tells you that the data is not within the cache and there are no errors.

    - The `isLoading` only tells you that the `status === "pending"` and that the data is currently fetching.

    Notice that, in both of these cases, you are not handling the case of `status === "pending"` and `isLoading === false`.

    **The `status === "success" tells you that the data is in the cache!**.

Finished 14
