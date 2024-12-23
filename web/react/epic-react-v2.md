# Epic React v2

## React Server Components

> You can [find the course material here](https://react-server-components.epicweb.dev/).

### Warm Up

- In Node, there is a _module resolution algorithm_ that is used to resolve modules when you import them.

  ```js
  import foo from 'bar' -> /* will resolve the import by looking into node_modules */
  ```

  When using ESM in the browser, we can load modules in many different ways.

  1. Using a URL. For example `https://example.com/shape.js`.
  2. Using relative path syntax. For example `./modules/shapes/square.js`.

  Instead, we can leverage the [`importmap`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap) feature. The `importmap` allows us to define what "type" of imports maps to what type of identifier. This introduces consistency in the codebase and creates a central place for us to manage dependencies – akin to `package.json`.

- Back in the old days, when importing scripts to your application, you had to put then at the bottom of the `body` tag.

  - The "bare" `script` identifier will block HTML parsing and execute as soon as the HTML parser encounters it.

  - That is not the case for `script type = "module"`.

    - This one works similar to _deferred_ script – the browser will load it asynchronously, and will execute it AFTER them HTML parsing is complete.

### Server Components

> The idea behind RCS is conceptually simple. Instead of requesting JSON data and handling that off to our components to generate UI, we request the UI itself.

- The RSC format allows for a couple of things.

  1. Mixing payload for interactive components (_client components_) and non-interactive components (_server components_) together.

  2. **Out of order streaming of components**. This is great as it minimizes waiting time for the UI to show up.

- **React has two exports, one for the server and one for the client**.

  - Notice that all the demos are using the same `React` import, no matter where the code is executed. How is that possible?

    - It works by crafting a custom `exports` object configuration in `React` package.json. It allows you to have different entry points for your application based on `--conditions` Node flag.

      ```bash
      node --conditions=react-server your_file.js
      ```

- On the server (wherever that is), we create the RSC payload and send it as a stream. Then you consume that stream on the client and pass the result to `createRoot`.

  - **Keep in mind that React will execute all RSCs first**.

    - If the data you fetch in RSCs takes a while to render, and **you do not use `Suspense`** React will not be able to stream all the other components!

      - `Suspense` is crucial in enabling out-of-order streaming.

      - In Next.js, the `loading.js` file acts as a _Suspense boundary_ for a given route.

- **Since we can't access `Context` in RSCs, passing data around can be painful**.

  - While a bit magical, [the `asyncLocalStorage` Node API](https://nodejs.org/api/async_context.html) is helpful in this regard.

    - **It allows you to access any piece of data you initially "seeded it with" during the request anywhere callbacks and functions in latter parts of the callstack**.

    - This is how the `cookies()` and similar functions in Next.js work under the hood.

- The **presence of the server does not mean we are server-side rendering**.

### Client Components

- You **can't import RSC into a RCC**. There are a couple of reasons.

  1. The RCC can contain secrets and other sensitive data that we use for creating the RCC JSX. We would not want those on the client.

  2. The RCC is not interactive. What if your RCC re-renders? Should we re-fetch the RSC? That would be very inefficient.

  Instead of trying to import RSC into a RCC, **focus on composition – composing RSCs with RCCs via `children` prop**.

- When creating the RSC payload on the server, React will create "placeholders" for RCCs.

  - Then, on the client, React will resolve those "placeholders" to "real" elements.

    - You might need to change how the paths are resolved on the client and on the server to make this work.

- **The `use client` directive is for the bundler**. It tells the bundler that this component is RCC and should have "placeholder" assigned to it in RSC payload.

### Client Router

- When adding the router, it dawned on me that, **so far, we are NOT doing any SSR**.

  - We first load the script that triggers requests to the `rsc` endpoint. The initial response from the server is a white page.

  - **This goes to show that SSR and RSCs are complementary, but one can be used without the other**.

- The basics of the router are as follows.

  1. To change the URL, we call some kind of `navigate` function.
  2. This function, apart from changing the URL, **triggers a new `rsc` request**.
  3. We then `use` that request to get the RSC payload.
  4. React takes care of updating the UI.

- One thing that I was unsure of when going through the _Pending UI_ section, is why did we bother to include `isPending` in the router context, if we are not using it?

  - To determine if something is "pending" we are using `location` and `nextLocation` (by means of `useDeferredValue`). Of course, that is a completely valid way to go about this, but why bother having `isPending` on the context then?

- The solution to handle race conditions was quite fascinating.

  - We used `Symbol` to determine if the promise that just resolved is the _latest_ promise.

  - Of course, we could have used an object with some kind of unique id inside, but the `Symbol` is more ergonomic as each instance is unique.

    - Do not be fooled by the `Symbol()` first parameter. It is not a key or anything like that. It is the symbol description!

- Apparently, the default behavior for React is to _always_ re-trigger Suspense boundaries when users click the forward/back buttons. I could not find any mentions of it, but I trust Ken on this one.

  - Either way, to **cache navigations, we used a central cache with keys stored in `window.history.state`**.

  - I like how we are leveraging the native web functionalities to achieve our goals!

### Server Actions

- The RSC payload also contains the reference to any _action_ you can pass as a prop to a form.

- **You can have both `handleSubmit` and `action` prop defined on the form**.

  - Keep in mind that `event.preventDefault` is not needed as the `action` will do that for you.

- The flow is pretty wild to me.

  1. You create an action.

  2. The bundler/loader will take that action an encode it.

  3. The server will embed that action within the RSC payload.

  4. When action is called, `createFromFetch` will call out special handler which will call a route we have created to handle RSCs.

     - This handler is called `callServer` and it should return the result of the action.

     - You might think that the return data is the RSC payload, but that is not the case. It is simply the result of running the action.

  5. In the route, you import the module via `import()`, execute the function and return the result.

  6. RSC payload contains the updated UI.

  I have to admit – wrapping my head around these concepts will take some time. Especially since the data sent over the wire is NOT that readable.

- To handle _server actions_ properly, we had to update the whole "root" in the UI. The update to the form UI that called the action is taken care of by the `callServer` function.

  - To update the whole UI (the "root"), are are effectively re-invoking our application upon every action. We then pipe that result alongside the action result.

- After finishing the "History Revalidation" exercise, it dawned on me, that **all we are really doing is calling to a server to get next data, and using React to display that data onto the screen**.

  - The payload might look weird (the RSC payload), but how it looks is an implementation detail.

  - There are few gotchas, like having to wait for stream to finish before doing something, but apart from that, the mental model is `action` -> `request` -> `response` -> `update UI`.

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

  - It seems to me that, **the transition caused by the change from "optimistic" -> "server" state will NOT trigger the suspense boundary, but instead will be executed in a transition**.

    - In the exercise, we created a new ship. When the optimistic state switched from `optimisticShip` to `null` (because the server action finished), the code attempted to fetch the new ship. The parents `Suspense` component did not trigger when fetching the newly created ship.

  - It is **critical to understand that, if the "initial value you passed to `useOptimistic` changes, the hook will return that value**. This is quite different to how `useState` works, as there, the initial value is "considered only once", and changes only when you re-mount the component.

    - This **is how the roll-back mechanism is implemented**. If an error occurs, the initial value you passed to the hook will be different than the one you set optimistically. This means that the underlying value returned by the hook will change as well!

      ```ts
      useOptimistic("foo", () => {}); // returns foo
      // update happened to "bar"
      useOptimistic("foo", () => {}); // returns bar
      // error happened
      useOptimistic("foo", () => {}); // return foo
      ```

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

## React Performance

### Element Optimization

> If you give React the same element you gave it on the last render, it wont bother re-rendering that element

- This has huge implications, especially related to how we pass props around.

  ```jsx
  function Message({ greeting }) {
    // some JSX
  }

  const message = <Message greeting={"Hi"} />;

  function Counter() {
    // state
    return <div>{message}</div>;
  }
  ```

  Here, no matter how many times we change the state inside the `Counter`, the `message` will not change. The `Message` component will be invoked only once! This works, because the `message` is referentially equal each time `Counter` runs.

  Contrast this with the following approach.

  ```jsx
  function Message({ greeting }) {
    // some JSX
  }

  function Counter() {
    // state
    return (
      <div>
        <Message greeting={"Hi"} />
      </div>
    );
  }
  ```

  Here, the `Message` component will be invoked every time the state in the `Counter` changes. This happens because the `<Message/>` will create a new object every time the `Counter` runs.

- Sometimes, the component you want to "embed" as a variable takes props. What should we do in this case?

  - **If that is the case, consider passing that component as props!**

  ```jsx
  function Message({ color }) {
    // jsx
  }

  function Counter({ message }) {
    // state
    return <div>{message}</div>;
  }

  function Main() {
    const [color, setColor] = useState("black");

    return (
      <div>
        <Counter message={<Message color={color} />} />
      </div>
    );
  }
  ```

  Here, the `Message` will only change if the `color` changes. From the `Counter` perspective, the `message` is the same element when state changes, because the `Main` did not re-render.

- Another technique would be to leverage `React.Context` API.

  - What if we moved the state from the `Counter` into the `Main`? Then, no matter what we do, the `Message` we are passing as a prop, will be re-created whenever `Main` is re-rendered. The answer is to leverage `React.Context`.

  ```jsx
  const ColorContext = CreateContext(null);

  function Message() {
    const color = use(ColorContext);
  }

  const message = <Message />;

  function Counter({ state, message }) {
    return <div>{message}</div>;
  }

  function Main() {
    const [color, setColor] = useState("black");
    const [state, setState] = useState(1);

    return (
      <ColorContext.Provider value={color}>
        <div>
          <Counter message={message} />
        </div>
      </ColorContext.Provider>
    );
  }
  ```

  Now, we made the `message` a constant variable again! The `Message` will only change when the value of the context changes.

- Of course, we can't forget about the `memo` and `useMemo`.

  - You are more likely to see components wrapped with `memo` than `useMemo`.

    - **It is important to remember that you can wrap components with `useMemo` as well!**

      ```jsx
      function Counter({ color }) {
        const footer = useMemo(() => <Footer color={color} />, [color]);
        // jsx
      }
      ```

      Depending on what you want to do, the above approach might be better, in terms of "visibility" than using `memo`.

      The problem with `memo` is that you have to see the component definition to understand that it is memoized (or name it appropriately.)

      The `useMemo` approach moves the memoization _closer_ to where we actually render the component.

### Optimize Context

- The most important thing to understand about `React.Context` value is that **all components consuming the context will re-render if the value changes**.

  - If your context value is not a primitive, this will happen every time the provider re-renders.

  - The solution here is to **wrap the value of the context with `useMemo`**.

    ```jsx
    const SomeContext = createContext(null);

    function SomeProvider({ name, color }) {
      const value = useMemo(() => ({ name, color }), [name, color]);
      return (
        <SomeContext.Provider value={value}>
          <div>Some elements</div>
        </SomeContext.Provider>
      );
    }
    ```

- Let us not forget about what we learned in the first chapter. **You can combine the `Provider` approach with "element optimization" via the `children` prop**.

  ```jsx
  const SomeContext = createContext(null);

  function SomeProvider({ children }) {
    const [color, setColor] = useState("black");
    const [name, setName] = useState("initial");

    const value = useMemo(() => ({ name, color, setColor, setName }), [color, setColor, name, setName]);
    return <SomeContext.Provider value={value}>{children}</SomeContext.Provider>;
  }
  ```

  This way, **no matter what changes inside the `SomeProvider`, the `children` is still the same as in the previous render**. This means that React will skip re-rendering the `children`.

- Going even further, **you might consider splitting the _setters_ and the _value_ into separate contexts**.

  - In some cases, we have components that only use the _setters_. Why would they need to change if the _value_ changes?

  - The **drawback here is that you will need to create multiple _providers_**.

  ```jsx
  const SomeContextValue = createContext(null);
  const SomeContextDispatch = createContext(null);

  function SomeProvider({ children }) {
    const [name, setName] = useState("hi");
    const [color, setColor] = useState("black");

    const value = useMemo(() => ({ name, color }), [name, color]);
    const dispatch = useMemo(() => ({ setName, setColor }), []);

    return (
      <SomeContextValue.Provider value={value}>
        <SomeContextDispatch.Provider value={dispatch}>{children}</SomeContextDispatch.Provider>
      </SomeContextValue.Provider>
    );
  }
  ```

  Now, if you `use` the `SomeContextDispatch`, your component will NOT re-render when `value` changes!

### Concurrent Rendering

The premise behind the _concurrent rendering_ is the following: instead of rendering everything all at once, React can break up the work into smaller chunks.

There are a couple of implications of this statement:

1. The work that takes a long time must be flexible enough to allow splitting into multiple parts. **If you are not able to split the computation into multiple parts, _concurrent rendering_ won't help you**.

2. React needs to know which operations are considered _high priority_. Luckily for us, we do not have to specify which things are high priority – React will prioritize user interactions.

There is one hook that we could utilize to make the UI _appear_ more fluid – the `useDeferredValue`. This hook is not the silver bullet to performance issues. **The `useDeferredValue` will only make your UI _appear_ faster, it does not improve the "real" performance in any way**.

When using `useDeferredValue`, React will render two times. This might be counterintuitive, but it makes sense.

- On the first render, React will "capture" the JSX of the component you passed the `deferredValue` to. So far, both values are in sync.

- Upon change, React will render the UI with a new value, but **it will re-use the JSX for the component you passed the `deferredValue` to**.

  - **This means you must memoize the component you passed the `useDeferredValue` to, otherwise the second render will be as slow as the first one**.

- Next, **in the background**, React will try to re-render the UI with the `deferredValue` the same as the "fresh" value.

  - This render should be interruptible, but as I mentioned: **if your work can't be split into multiple chunks, the UI will freeze here**.

```tsx
function Component() {
  const query = useQuery();
  const deferredQuery = useDeferredValue(query);

  // First render – values in sync
  // Second render – query changes, values out of sync.
  //   React re-uses JSX for `SlowMemoizedComponent`. The UI appears to be fast.
  //   Background render starts.
  // Third render – background render finished. All values in sync.

  return (
    <div>
      <FastComponent query={query} />
      <SlowMemoizedComponent query={deferredQuery} />
    </div>
  );
}
```

### Code Splitting

The best thing you can do in terms of performance is to _do less stuff_. This also goes for the amount of code you ship to the browser.

One of the mechanism for delaying JavaScript code execution the browser has to do, is to load components only when they are needed. React has a built-in mechanism for that – `lazy` function which interplays with `Suspense`.

```tsx
const LazyGlobe = lazy(() => import("./globe.tsx"));

// Somewhere in the code
<Suspense>
  <LazyGlobe/>
<Suspense>
```

---

In addition, you can **preload** the `LazyGlobe` when user performs some action, like hovering over a link.

```tsx
const loadGlobe = () => import("./globe.tsx");
const LazyGlobe = lazy(loadGlobe);

// Somewhere in the code
<a href="#" onMouseEnter={loadGlobe}>
  SomeLink
</a>;
```

Note that you can call the `loadGlobe` as many times as you want. It will be loaded only once.

---

After adding the `Suspense` and providing the `fallback` prop, you will notice that React, by default, _always_ triggers the _Suspense fallback_, no matter if the module is "ready to go". React will display the `fallback` _for some time_, to avoid flash of "loading" state.

While showing the `fallback` prop while rendering for the first time makes sense, we can do better if the app is already in a "settled" state. We can leverage `useTransition` to display the "old UI" while the module is loading.

```tsx
const LazyGlobe = lazy(() => import("./globe.tsx"));

// Somewhere in the code
const [showGlobe, setShowGlobe] = useState(false);
// Use the `isPending` for loading state.
const [isPending, startTransition] = useTransition()

<Suspense>
  <button onClick = {() => startTransition(() => setShowGlobe(true))}>Click me</button>
  {showGlobe ? <LazyGlobe/> : null}
<Suspense>
```

Apart from this use-case, the _transitions_ are also very useful for data-fetching and changing routes.

### Expensive Calculations

Any calculation you do within the "component function", will, by default, be repeated upon every render. In most cases, this is not problematic at all, because most of the computations you do are really fast.

But, there might be cases, where making such computation over and over again, makes the application feel slow. In such cases, we ought to ensure this computation only runs when it needs to run.

---

You can use the `useMemo` hook to avoid running expensive computations. **Make sure to measure the _before_ and _after_ when using `useMemo`** – every time you introduce an "optimization", you also introduce complexity.

For example, consider this, objectively, unjustified usage of `useMemo`.

```tsx
const someCallback = useMemo(() => {
  return () => {
    console.log("works");
  };
}, []);

<button onClick={someCallback}>Click</button>;
```

You only increased the "visual noise" of the component. There is no benefit in using `useMemo` this way.

Now, consider the following code.

```tsx
const result = useMemo(() => expensiveCalculation(inputValue), [inputValue]);
```

Assuming the `expensiveCalculation` is expensive (names can be misleading!), this usage of `useMemo` makes much more sense.

---

Another way to improve _perceived_ performance is to use WebWorkers. This way, the main thread rendering logic does not have to compete with calculations we have to make.

Working with the "raw" WebWorkers API is not that fun. That is why [you should consider using a library, like `comlink`](https://www.npmjs.com/package/comlink).

There are a few steps to make `comlink` working with your code.

First, you create a file where you isolate the functionality you want to "embed" into the worker.

```ts
// expensive-calc-in-worker.ts

function expensiveCalculation(inputValue) {
  // code
}

const exposed = { expensiveCalculation };

Comlink.expose(exposed);
```

Then, in another file, you initialize the worker with the contents of this file. `comlink` will handle the communication between the worker and the main thread for you!

```ts
const moduleUrl = new URL("./expensive-calc-in-worker.ts", import.meta.url);
const worker = new Worker(moduleUrl, { type: "module" });

const exposedViaComlink = Comlink.wrap(worker).

export async function expensiveCalculation(inputValue) {
  return exposedViaComlink.expensiceCalculation(inputValue)
}
```

After you have all this set up, **consider using the `use` hook in your React code**. The `expensiveCalculation` became a promise (because the communication between the worker and the main thread is asynchronous).

```tsx
const calculationPromise = useMemo(() => expensiveCalculation(inputValue), [inputValue]);

const calculation = use(calculationPromise);
```

You could leverage the `cache` API here as well!

### Optimize Rendering

Kent emphasizes that **using `memo` all over the place is not the silver bullet and might eventually lead to _worse_ performance compared to not using it at all!** The reason, of course, is the increasing amount of computation React has to do to compare the props. Those operations are not free!

---

When using `memo`, it is very important to keep track of all the props your component takes. **If one propr is "unstable" (is not referentially the same across re-renders), the whole memoization will be in vain**.

```tsx
const [state, setState] = useState(null);

const someCallback = (value) => {
  const newValue = setState(newValue); //
};

<MemoizedComponent onClick={someCallback} />;
```

In the example above, even though the `MemoizedComponent` is wrapped with `memo`, the memoization will not work. That is because the `someCallback` is different across re-renders. A solution here would be to wrap the `someCallback` with `useCallback` to ensure it is stable.

```tsx
const [state, setState] = useState(null);

const someCallback = useCallback((value) => {
  const newValue = setState(newValue); //
}, []);

<MemoizedComponent onClick={someCallback} />;
```

Since we have no way of annotating that the prop ought to be memoized (perhaps we could use [TypeScript type branding](https://effect.website/docs/guides/style/branded-types)), it is very easy to make those mistakes.

---

The `memo` function takes a second argument. This argument allows you to customize when React considers the props "new" or "changed". If you wish, you can hardcode that function to return `false` which means "the props did not change", though that would most likely break your application.

But, **instead of writing an elaborate comparator function, consider passing primitives to your components as props**. It is very tempting to pass the whole `user` if you only need the `id` of the user. If you pass only the `user.id`, the string, assuming the `user` object does not change its values, will be same across re-renders!

```tsx
<SomeComponent user = {user}/> // Instead of this.

<SomeComponent userId = {user.id}/> // Consider this.
```

---

### Windowing

React, and your browser, can render a lot of items at the same time, but in some cases, the amount of items is so huge, that the browser and React struggles – think of lists with thousands of items.

Luckily, we have common sense at our side. Do we really need to render all those items _at the same time_? I would argue that is not the case.

If we do not have to render all those items at the same time, we can "cheat" a bit and render a "window" of those items. The "window" would consist of only a handful of items, which we will be replacing as user scrolls.

---

One thing that caught my attention in this section of the workshop was how we **faked the appearance that we are displaying lots of items, but in reality we are displaying only a subset of them at a given time**.

```tsx
<ul>
  <li style={{ height: virtualizer.totalHeight() }} />
  {virtualizer.getItems().map((item) => {
    return; // stuff
  })}
</ul>
```

Without this first `li`, the UI would behave in a weird way – the scrollbar on the element would appear "stuck" in the same place, but we would still have the ability to scroll down.

---

**When working with windowing, consider the fact that the native _search_ functionality of the browser will not work as expected**. This functionality searches across what is inside the DOM. Since you are rendering only a handful of items at a given time, it might happen that the item the user is searching for is not yet rendered.

The solution for this would be to build a custom search box, and prompt the users to use this box instead of the browser native _search_ functionality. Keep in mind that you can focus the search box when user tries to hit `cmd+f` or `ctrl+f`.

Of course, this might be surprising for users. It all depends on your product and your user-base.

## Advanced React APIs

### State Optimization

- React will not trigger a re-render of your component if the state value is already set to the value you are setting it.

  - This works really well for primitive types.

  - This _might_ work with objects and other "complex" types, but **remember that the equality check is a subject to the referential equality here**.

- You might find the _callback form_ of `useState` very handy here.

  - The parameter that gets passed to the callback is the previous state. No matter how complex it is, if you return it, React will NOT trigger a re-render!

  ```jsx
  setState((prevState) => {
    return prevState; // I'm equal to the current state.
  });
  ```

### Custom hooks

Custom hooks are a great way to encapsulate functionality and "give it a name".

Usually, for "giving a name" to the code, we use functions. In React, if we want to use other "foundational" hooks like `useEffect` and `useState`, we can't use regular functions – only hooks can "consume" other hooks.

```tsx
function useCounter() {
  const [] = useState();
  // Other stuff
}
```

**When returning functions from hooks, make sure to wrap them with either `useMemo` or `useCallback`**.

If you do not, if the caller of the hook wants to use that function within the `useEffect`, they will have a bad time – each time component re-renders, their hook will be invoked. In some cases that is desirable, in others, and I would argue in most cases, that is not desirable.

```tsx
function useCounter() {
  const [state, setState] = useState();
  // Other stuff

  // Stable function!
  return useCallback(() => {
    setState();
  }, []);
}
```

### Context Hook

The `Context` API is React is pretty great. It allows us to "share" values between different components in the tree without having to pass props around. The solution is not without tradeoffs – the more layers of indirection, the harder it is to reason about the code.

When using `Context` API, there are things to be mindful of:

1. The `Context` API is not for frequently changing values. This relates to performance and how "consuming" context influences re-rendering of the application.

2. Consider using `SOME_TYPE | null` as the type for your context. This way, you will not have to provide a default value that might be complex.

3. Consider taking the `defaultValue` as a prop on the _provider_ level. This will help you in tests.

```tsx
const SomeContext = createContext<SomeType | null>(null);

function Provider({ children, defaultValue }) {
  const ctxValue = defaultValue ?? deriveValue();

  return <SomeContext.Provider value={ctxValue}>{children}</SomeContext.Provider>;
}

function useSomeContext() {
  const ctxValue = use(SomeContext);
  if (!ctxValue) {
    // This check makes it easier for TypeScript to infer the proper return type.
    throw new Error("boom");
  }

  return ctxValue;
}
```

### Layout Computation

The `useEffect` will run _after_ the browser painted to the screen. This is quite nice, since the code you defined in `useEffect` won't block stuff from appearing on the screen. But that also means, that, any DOM manipulations you make in `useEffect` _could_ create a little bit of flicker – the browser has already painted and displayed the DOM to the user!

This is where `useLayoutEffect` comes in. It will run _before_ browser painted to the screen, but _after_ the DOM was updated. This is an ideal time to perform those DOM manipulations!

**Be mindful of the usage of `useLayoutEffect` as it runs synchronously and could trigger reflows which can negatively impact the performance of your application**.

### Imperative Handles

This one is about `useImperativeHandle` hook. I rarely reach out for this hook, but when I do, **it is mostly when I'm mimicking a certain HTML element in React**.

For example, I would reach for this hook when implementing `collapsible` component which then uses the `details` and `summary` tags. The consumer of such component would, most likely, want to have a way to have an _imperative_ way to open/close it.

**You do not have to wrap your components with `useRef` to pass refs**. Refs can be regular props, but you have to name the prop differently than just `ref`. In some cases, getting the right types for `refs` is quite hard in TypeScript. Passing it as a regular prop makes our life so much easier!

### Focus Management

No matter what technology you use, proper focus management is tricky to achieve. In React, it is no different.

The main "issue" is that **React batches state updates and runs them asynchronously**. This is great for performance, but makes focus management harder.

```ts
onClick(() => {
  setEditing(true);

  inputRef.current.select(); // This will not work!
});
```

In the above snippet, we will most likely call the `.select` BEFORE the DOM got the chance to update. Again, this is because the `setEditing` will run asynchronously. To "force" state to be synchronous, we can use the `flushSync` call.

```ts
onClick(() => {
  flushSync(() => {
    setEditing(true);
  });

  inputRef.current.select(); // Will work!
});
```

Now, we know that the DOM was updated after the `setEditing` call!

**Using `flushSync` has performance implications and should be used sparingly**.

### Sync External State

While in most cases, for _synchronization_ we use `useEffect`, it might be more advantageous to use `useSyncExternalStore`, especially when we want the "state" of that external dependency to influence React rendering lifecycle (to be _reactive_).

```tsx
const mediaQuery = "(max-width: 600px)";

const getSnapshot = () => {
  return window.matchMedia(mediaQuery).matches;
};

const subscribe = (callback: VoidFunction) => {
  const mql = window.matchMedia(mediaQuery);

  const listener = () => {
    callback();
  };

  mql.addEventListener("change", listener);

  return () => mql.removeEventListener("change", listener);
};

function NarrowScreenNotifier() {
  const isNarrow = useSyncExternalStore(subscribe, getSnapshot);
  return isNarrow ? "You are on a narrow screen" : "You are on a wide screen";
}
```

For a very long time, I've used `useEffect` and `useState` to achieve the same functionality. The problem with `useEffect` is that it does not support SSR that well. The **`useSyncExternalStore` will trigger the nearest `Suspense` boundary if you do not provide the server-side snapshot to the hook**. This is quite nice!

---

During the second exercise, Kent instructed us to create a function which then returns a hook. I very rarely do this, so it felt really strange.

```tsx
function makeMediaQueryStore() {
  // data

  return function useMediaQuery() {
    return useSyncExternalStore();
    // data
  };
}
```

---

Making sure `useSyncExternalStore` is working with SSR is quite easy. You can either provide the _server-side snapshot_ OR rely on `Suspense`.

I love the integration with `Suspense` here. Prior to `useSyncExternalStore`, I remember using the `typeof window` checks and `useEffect` with `isMounted` state to get the same behavior I now get with `Suspense`!

Interestingly, if you do SSR, and you rely on `Suspense` for `useSyncExternalStore`, React will still log an error that the _server-side snapshot_ is missing. You can use the `onRecoverableError` in [`hydrateRoot`](https://react.dev/reference/react-dom/client/hydrateRoot) to silence this error.

### Wrapping up

A very nice refresher!

- `useSyncExternalStore` is very useful.

- Returning the same state will not trigger re-rendering.

- Use `useLayoutEffect` with caution.

- `flushSync` is your friend, but be mindful of the tradeoffs.

- `useImperativeHandle` could really simplify your code, but only when used in the right circumstances.

## React Hooks

### Managing UI State

- In the `event.currentTarget`, the `currentTarget` refers to the DOM node **to which the event listener is attached to**.

  - The `.target` would be the element that triggered the event.

  - **In some cases, the `currentTarget` = `target`, but they can be different**. The event might have bubbled "up" from the child to the parent.

  - **The `currentTarget` is only available in the event handler "cycle"**. If you try to access it after performing some async work, it will be null!

    ```js
    onChange={event => {
      const eventTarget = event.currentTarget
      // event.currentTarget is accessible here.

      setTimeout(() => {
        // Prints the DOM node and `null`.
        // We have "captured" the `eventTarget` in a variable so it is accessible.
        // If you access it directly, it will be null.
        console.log('target', eventTarget, event.currentTarget)
      }, 500)

      // event.currentTarget is accessible here.
    }}
    ```

- **Deriving state will drastically simplify your code**.

  - Do not fall into the trap of using `useState` everywhere. If you can, derive it!

  - Having said that, there is nothing wrong with using multiple `useState` calls.

### Side-Effects

- You can clean up event listeners via the `AbortController.signal`. Quite useful when you attach multiple listeners in a single `useEffect`.

- **Make sure you pass the same parameters to `removeEventListener` as you did to `addEventListener`**. If you do not, the listener _might_ (it depends) not be removed.

  - [You can read more about it here](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener#matching_event_listeners_for_removal).

### Lifting State

- Lifting state up is often necessary when both components needs the same state.

- **But you should also refactor the other way around** – if you longer need to share state, **you should _colocate state_** within the component that needs that state.

- Whenever you manipulate where the state is defined, consider what happens when the component unmounts.

  - If you collocated the state, the state will be reset.

  - If you lifted the state up, the state will most likely persist.

### DOM Side-Effects

- There are _refs_ and _refs callbacks_.

  - **In newer versions of React, you can return a _cleanup_ function from the _ref callback_**.

    - The `ref` _cleanup function_ will run **every time the component re-renders**.

- You do not have to (and I strongly believe you should not) put the refs inside the dependency arrays.

### Unique IDs

- The `useId` is very handy – it **creates an ID that is guaranteed to stay consistent across the server-client boundary**.

  - **This is because it leverages the "path" of the component in the "component tree"**.

  - This also means that **for the `useId` to work properly, your component tree needs to be the same on the server and on the client**.

### Tic Tac Toe

- The `.with` function is very useful for inserting data at some position in an array.

  ```js
  const arr = [1, 2, 3];
  const newArr = arr.with(1, 3);
  // [1,3,3]
  ```

### Wrapping up

A great way to refresh some of the basics of handling state. The workshop is definitely tailored more to people who are just starting out, but nevertheless, I found it valuable.

I'm glad that Kent decided to include `useId` in the material. I find this hook very useful!

## React Advanced Patterns

### Composition

- In this exercise we refactored a bunch of components to take `React.ReactNode` props instead of objects or strings.

  - We named those components **_layout components_**.

- I wonder how this technique relates to "slots". IMHO we created bunch of "slots" in this exercise already.

### Latest Ref

- Very useful pattern when you **want to ensure the function you are calling is the "latest" version of that function**.

  - This means that it has captured all the latest variables and functions in its closure.

  - **Super useful in scenarios where you pass functions as callbacks** and **when you take a function as parameter for custom hooks**.

    ```js
    function useTimeout(callback) {
      useEffect(() => {
        setTimeout(() => {
          callback(); // The `setTimeout` created a closure with the INITIAL value of the `callback`.
        }, 300);
      }, []);
    }
    ```

    ```js
    function useTimeout(callback) {
      const latestCallback = useRef(callback);
      useEffect(() => {
        latestCallback.current = callback;
      });

      useEffect(() => {
        setTimeout(() => {
          callback.current(); // Now we are calling the _latest_ version of the callback!
        }, 300);
      }, []);
    }
    ```

### Compound components

This pattern is about having a set of components that compose together to form an interactive UI.

```jsx
<Toggle>
  <ToggleOn>The button is on</Toggle>
  <ToggleOff>The button is off</Toggle>
</Toggle>
```

Usually, we have a "root" component, in this case the `Toggle` that holds the state, and shares that state via context with other, child components.

The **_compound components_ pattern excel at composition**.

```jsx
<Toggle>
  <ToggleOn>The button is on</Toggle>
  <hr/>
<div>
  <ToggleOff>The button is off</Toggle>
</div>
</Toggle>
```

Consider an alternative API:

```jsx
<Select
  options = {
    [
      name: "foo",
      value: "bar"
    ]
  }
/>
```

Here, we are on the other side of the "composability" spectrum. This API is also viable – it introduces different tradeoffs but enforces consistency.

### Slots

This pattern was very new to me. Kent mentioned that, so far, he seen it only used in [`react-aria`](https://react-spectrum.adobe.com/react-aria/Menu.html#text-slots).

This pattern **attempts to enable re-use of different "pieces" a given compound component might consist of**.

Consider the `Toggle` compound component. The JSX might look like as follows.

```jsx
<Toggle>
  <ToggleSwitch/>
  <ToggleOnText>I'm switched on!</ToggleOnText>
  <ToggleOffText>I'm switched off!<ToggleOffText>
</Toggle>
```

Given the amount of components in design systems, we would have to create the same "Text" component for all relevant components.

I do not see anything inherently bad with such "duplication", but if the number of similar "Text" components is large, you might want to consider the following API.

```jsx
function Toggle({ children }) {
  const [on, setOn] = useState(false);

  const slots = {
    onText: { hidden: !on },
    offText: { hidden: on },
    switch: { on, onClick: setOn((on) => !on) },
  };

  return <SlotContext value={slots}>{children}</SlotContext>;
}
```

The `SlotContext` is nothing but a bit of logic to merge props you define in `slots` constant and the ones that are passed to the element.

```jsx
function useSlotProps(props: Record<string, unknown>) {
  const propsFromSlot = use(SlotContext);
  return {...propsFromSlot, ...props}
}
```

Then, your JSX might look like this:

```jsx
<Toggle>
  <Switch/>
  <Text slot = "onText">I'm switched on!</Text>
  <Text slot = "offText">I'm switched off!<Text>
</Toggle>
```

And here is the `Text` component.

```jsx
function Text(props: SpanProps & {slot?: string}) {
  const props = useSlotProps(props)

  return <span {...props}/>
}
```

**As long as the "root" component defines `SlotContext`, we can re-use "slot compatible" components within _any_ context**.

But this **pattern is not free**:

- Another layer of indirection.

- **There is no good way to have type-safety on the `slot` prop**. Since the element, in theory, can be used as a child of any element that defines the `SlotContext`, there is no good way to constrain the `slot` prop.

  - Kent mentioned that `react-aria` performs a runtime check to ensure you did not make any typos while passing value to the `slot` prop.

After [reading this documentation page](https://react-spectrum.adobe.com/react-aria/advanced.html) I _think_ I understand the motivation behind this pattern. If you want extreme levels of reusability and flexibility, which _some_ component libraries might require, it seem like the _slots_ pattern is a way to go.

### Prop Collections and Getters

The main idea is to provide users with necessary props for a component to work. This is usually done via hook that returns either a _prop getter_ or _prop collections_.

#### Props Collections

You can think of them as "peanut butter props". You _spread_ them over a component.

```jsx
const { togglerProps } = useToggler();

return <Toggle {...togglerProps} />;
```

**While this looks nice, it is not really ergonomic**. Consider the case where you want a custom `onClick` logic to run _alongside_ the `onClick` provided by the `togglerProps`.

```jsx
const { togglerProps } = useToggler();

return (
  <Toggle
    {...togglerProps}
    onClick={(...args) => {
      togglerProps.onClick(...args);
      // My custom logic
    }}
  />
);
```

**This "reveals" what kind of props the `togglerProps` contains leaking the internal implementation details**.

#### Prop Getters

**Similar to _prop collections_, but more flexible pattern**.

```jsx
const { getTogglerProps } = useToggler();

return (
  <Toggle
    {...getTogglerProps({
      onClick: () => {
        // my custom code
      },
    })}
  />
);
```

The `getTogglerProps` will take in all the additional props you provided and merge it with its own internal props.

This is the pattern Kent recommends we consider using.

### State Initializer

This pattern is _mostly_ about how to properly handle resetting the state to its original, _initial_ value.

```jsx
const [state, setState] = useState(initialState);
const reset = () => {
  setState(initialState);
};
```

While the code above might look correct to you, **what if the `initialState` changes during the lifecycle of the component?**.
If it does, the `reset` will not reset to `initialState`, but the _updated_ `initialState`. This might, or might not be the behavior you want.

I would argue that, if prop is called `initialX`, then it should reset to value which it was rendered with, not the updated `initialX`.

```jsx
const { current: initialState } = useRef(initialValue);
const [state, setState] = useState(initialState);

const reset = () => {
  setState(initialState);
};
```

The `useRef` allows us to capture the _true_ `initialValue` and ignore any updates to it.
Now, no matter if the `initialValue` is updated or not, we will always reset to that value.

### State Reducer

This one is **about using the _inversion of control_ to enable users of your APIs to implement custom behavior**.

Imagine where you want something "special" to happen when state has certain value. One way to do this, would be to add some kind of prop to the API.

```jsx
const {} = useToggler({ disabledAfterClicks: 4 });
```

The problem is that these kind of properties are very shallow. They have a tendency to multiply and pollute the interface.

Take a look at this alternative API.

```jsx
import { togglerReducer } from "./toggle";

const {} = useToggler({
  reducer(state, action) {
    if (action.type == "toggle" && clickedTooManyTimes) {
      return { ...state, on: state.on };
    }

    return togglerReducer(state, action);
  },
});
```

This way, the **user has complete control over their custom use-case, but also gets to preserve the default behavior**.

Now, in the workshop, Kent decided to _export_ the `togglerReducer`. I'm not a fan. To be, it exposes too much internal details of the hook.

**As an alternative to exporting the `togglerReducer`, you can pass it as a prop to the `reducer` function!**

```jsx
const {} = useToggler({
  reducer(defaultReducer, state, action) {
    if (action.type == "toggle" && clickedTooManyTimes) {
      return { ...state, on: state.on };
    }

    return defaultReducer(state, action);
  },
});
```

I find this version more robust.

### Control Props

**You can think of this pattern as implementing your own `onChange` and `value` props**. Of course, the name of those props is heavily dependant on your use-case.

This pattern is **_State Reducer_ on steroids, as it gives the user full flexibility**. The downside is that the user has to manage some state themselves.

**It turns out, doing this requires a lot of work**. First, you have to determine whether the state in your hook is controlled or uncontrolled.

```jsx
function useToggle({ on: controlledOn, initialOn, onChange }) {
  const [state, dispatch] = useReducer(reducer, { on: initialOn });
  const isControlled = controlledOn != null;
}
```

Next, you will need to implement the `onChange`-like function. You want to call the `onChange` only when it is provided.

```jsx
function useToggle({ on: controlledOn, initialOn, onChange }) {
	const dispatchWithOnChange = (action: ToggleAction) => {
		if (onIsControlled) {
			const suggestedChanges = reducer({ ...state, on }, action)
			return onChange?.(suggestedChanges, action)
		}

		dispatch(action)
	}
}
```

**The complexity grows with the number of "control props"**

## Wrapping up

It is interesting to see that some patterns did not change that much throughout the years.

I first learned about _Prop Collections_, _State Reducer_ and _Control Props_ back in 2019 or so. To me, this means that those patterns are really valuable and robust!

The _Slots_ pattern was very interesting. To me, this pattern steps into over-engineering territory, but I can also see its usefulness in very large organizations.
