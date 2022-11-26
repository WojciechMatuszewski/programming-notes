# FrontendMasters React Performance v2

- Not doing stuff is faster than doing stuff.

## How does React do what it does?

- There are three phases to the _rendering cycle_

  - The _render phase_ where **React invokes your components**.

  - The _commit phase_ where **React puts the outputs of your components into the DOM**.

  - The _clean-up phase_ where **React invokes the cleanup functions of `useEffect` and `layoutEffect` hooks**.

- Keep in mind that **React will re-render all children of a given parent**.

- In React 18 **all the state changes are batched**. This was **not the case in prior versions of React**.

  - There was _partial_ batching support for events for change handlers and hooks (synchronous).

    - Any changes originated from an async piece of code were not subject to batching.

## First exercise – wordman

- The instructor showcases how **slow compute of an initial value of the `useState` can negatively impact your app performance**.

  - Keep in mind that **the initializer of your state will be called every render, unless you use the _callback_ form of the initializer**.
    This is **NOT A REACT THING, THIS IS JAVASCRIPT EXECUTING**.

    ```jsx
    const [state, setState] = useState(slowToComputeValue()) // Will cause performance issues when the component re-renders
    const [state, setState] = useState(() => slowToComputeValue()) // Only the initial render is slow
    ```

## Second exercise – packing list & wordman

- It is often said that you should _lift your state up_ when necessary.

  - The key word here is _"when necessary"_. **Often, the act of _listing the state up_ causes performance problems**. The reason is that one state change will trigger the re-render of every children. If that state lives at the root of your application, the whole app will re-render!

- Sometimes, the best thing you can do is to **create a new component which would encapsulate all the necessary state**. This component could render a couple of other components.

  - Another solution would be to **use the `children` prop to render expensive components in a component which can change**. It is important to **pass the children to that component in a component that does not change often**. It works because the **children component belongs to the tree where we declare it in JSX and NOT where we render it in the DOM**.

    - This is a great optimization technique which works similarly to encapsulating the state as close where it's needed.

- The instructor talked about the **shallow comparison** and the fact that **the non-primitives compare by reference**.

  - This is a common pitfall one can come across while working with React.

    - One can use `useMemo` and `useCallback` hooks to ensure references stay the same given the same inputs.

      - Please **make sure not to abuse these – favour the correct composition first!**.

      - Another technique would be to **use the `useReducer` and pass dispatch around**. This way **you do not have to create functions wrapped with `useCallback`** as you would pass the dispatch closest to where it is used. All the dependencies are in the reducer function already!

## Third exercise – the `context` API

- The instructor showed the two most common gotchas with the `useContext` and the `context` API in general.

  1. First is the fact that, if your provider re-renders often, the value you pass to the `Provider` component might change. If that happens, all the consumers will re-render.

      - The solution here is **to use two `Context` objects, one for the data, one for the `dispatch` or the setters**.

        - Keep in mind that the **order in which you render the providers matters**. You want to render the "actions" provider first.
          All parents trigger re-render on their children. You would not want to re-render the "actions" provider when `items` changed.

          ```jsx
            <ActionsContext.Provider value = {dispatch}>
              <ItemsContext.Provider value = {items}>
                {children}
              </ItemsContext.Provider>
            </ActionsContext.Provider>
          ```

  2. The `useContext` API injects "hidden" props into a given component. Those "hidden" props will always bust the memoization.

      - This is not really a problem so there is no solution. It is how the API works. If you make your context granular, you should be good.

## Fourth exercise – hottest-takes

- Here we are dealing with a situation where the "state" context has two pieces of state which are disjointed from each other. If one changes, a new context "object" value is created.

  ```jsx
  <ActionsContext.Provider value={actions}>
    <StateContext.Provider value={{ posts, users }}> // two separate `useReducer` calls
      {children}
    </StateContext.Provider>
  </ActionsContext.Provider>
  ```

    So, if the `posts` change, all the components "subscribed" to the `users` will change as well. The **solution is to make the `value` prop more granular, by adding additional context providers**.

- When using **memoization, and your component takes in `children` prop, ensure that the `children` render in a place that does not re-render frequently**. Otherwise you will be busting memoization. Every time the place where you render those `children` re-render, the memoized component will re-render.

    ```jsx
    const [users, setUsers] = useState([])
    return (
      <div>
        {users.map(user => {
          return (
            <User user = {user}>
              <SomeExpensiveComponent/> // This one will re-render every time the `users` change. If possible, rendering this component elsewhere (either inside User or at the parent level).
            </User>
          )
        })}
      </div>
    )
    ```

  - This is because the `children` are a complex object data structure.

- Regardless of the optimization techniques, **consider normalizing your state data structures**.

  - We have a tendency to use arrays for global state, but **an object with keys is usually much better from the "read" use-case** which is **the main access pattern in the UIs**.

## Fifth exercise – lots-to-do

- Here the instructor showed us two React 18 hooks, the `useDeferredValue` and `useTransition`. **Both are useful when we want to make some things "low priority"**.

  - In React 17 it is not possible to differentiate between low-priority and high-priority updates. With React 18, such distinction is possible.

## My key takeaways

- **You can gain a lot in terms of performance by moving stuff around**.

  - Co-locating and encapsulating state is one of the best ways to improve performance without introducing complexity.

    - As a bonus, you will most likely make your app structure more readable.

- **Using `memo` is okay, but do not overuse it**.

  - Keep in mind the `memo` gotcha with the `children` prop – where you render the `children` matters!

- Ensure that **your `context`s are granular!**.

  - Use the `state`/`dispatch` context pattern.

- React 18 exposes hooks to make the UI feel more responsive – the `useDeferredValue` and the `useTransition` hooks.
