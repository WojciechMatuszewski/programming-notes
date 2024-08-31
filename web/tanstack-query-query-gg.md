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

- The React Query **implements _stale-while-revalidate_ model**.

  - This means that in some cases, you might present stale data to the user.

    - Remember, **in MOST circumstances, stale data is better than no data**.

  - The **`staleTime` setting DOES NOT remove the data from the cache**. It only tells the library when it should consider _refetching_ the data!

    - To **control when the data is removed from the cache, use the `gcTime` property**.

- The `refetchInterval` accepts a function. This is very useful if you wish to stop pooling after certain condition is met.

- There are at least _two_ ways you could tackle _dependant queries_.

  Consider the following code.

  ```ts
  function fetchBook(bookId: string) {
    // code
  }

  function fetchAuthor(authorId: string) {
    // code
  }
  ```

  We first have to fetch the book, then we can fetch the author.

  ```ts
  const {} = useQuery({
    key: [],
    queryFn: () => {
      const book = await fetchBook("someId");
      const author = await fetchAuthor(book.authorId);
    },
  });
  ```

  We have consolidated the fetches into a single `useQuery` call.

  - You do not have to worry about multiple _loading_ and _error_ states.

  - **This approach disables the data re-use for `book` and `author` fetches**. If another component only needs to fetch the book, it will NOT use the cache as the data for the book is NOT in the cache.

  In most cases, it is better to split the fetches into _dependant queries_.

  ```ts
  const { data: book } = useQuery({
    key: [],
    queryFn: () => {
      const book = await fetchBook("someId");
      return book;
    },
  });

  const {} = useQuery({
    key: [],
    queryFn: () => {
      const author = await fetchAuthor(book.authorId);
    },
    enabled: book != null,
  });
  ```

  Now, each query is independent from the cache perspective.

  **The same applies to running multiple queries in parallel**.

  - You could use `Promise.all` in a single query, but that approach suffers from inflexibility at the cache level.

  - You could use `useQueries` which will run the queries in parallel, and also cache the results of the queries separately!

- There is a **difference between `initialData` and `placeholderData`**.

  - The `initialData` is treated as data returned from the `queryFn` function. This means that **React Query will NOT refetch the data until the `staleTime` is up**.

  - The `placeholderData` is treated as "incomplete" data and, no matter what the data is, React Query will trigger `queryFn` to replace this data.

- The course **showcased two ways of doing pagination** – via the `useQuery` and via the `useInfiniteQuery`. There is a difference between those two.

  - First, **there is a difference with how data is stored in the cache**.

    - For the `useInfiniteQuery` the new data is appended to existing cache entry.

    - For the `useQuery` the new data is a **separate cache entry**.

  - The **`useQuery` hook really shines for UIs with explicit "next page" / "previous page" buttons**.

    - Here, **using the `previousData` parameter is critical**. Otherwise, every time the user hits one of those buttons, the whole UI will "vanish" and display the loading state.

      ```ts
      const [currentPage, setCurrentPage] = useState(0);

      const { data, isLoading, isError, isPlaceholderData } = useQuery({
        queryKey: ["posts", { currentPage }],
        queryFn: () => fetchPosts({ page: currentPage }),

        placeholderData: (prevData) => {
          return prevData;
        },
      });
      ```

      Use the `isPlaceholderData` to display loading state when transitioning from one page to another. Use the `isLoading` to display the "initial" loading state.

      **This technique gets even more powerful if you introduce prefetching into the mix**.

      ```ts
      const getPostsOptions = ({ currentPage }: { currentPage: number }) => ({
        queryKey: ["posts", { currentPage }],
        queryFn: () => fetchPosts({ page: currentPage }),
        placeholderData: (prevData) => {
          return prevData;
        },
      });

      // In a component

      const [currentPage, setCurrentPage] = useState(0);

      const { data, isLoading, isError } = useQuery(getPostsOptions({ currentPage }));

      const queryClient = useQueryClient();
      useEffect(() => {
        void queryClient.prefetchQuery(getPostsOptions({ currentPage: currentPage + 1 }));
      }, [currentPage, queryClient]);
      ```

      Now paginating through the list will feel near instantaneous.

  - The **`useInfiniteQuery` hook really shines for UIs with infinite lists that we append to** when user scrolls or clicks a button.

    - It is important to understand that, **`useInfiniteQuery` will refetch data for all the page entries**.

      - If you have a lot of pages in the cache, that might be a lot of requests. **You can control this behavior via `maxPages` parameter**.

  - **I've seen people trying to use `useInfiniteQuery` in place of `useQuery` for pagination and that does not end up well**.

    - Be mindful about the difference here.

TODO: Talk about mutations, the `currentTarget` "missing" and the fact that one can specify two callbacks for `onSuccess`.

Finished 29
