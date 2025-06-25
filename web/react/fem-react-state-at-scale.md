# FEM React State At Scale

Watching [this course](https://frontendmasters.com/workshops/react-state-at-scale/) and taking notes.

## Learnings

### Anti-patterns

#### Using `useEffect` instead of deriving state

This is a very common anti-pattern that you will see in all codebases.

```tsx
const [orders, setOrders] = useState([]);
const [total, setTotal] = useState(0);

useEffect(() => {
  setTotal(/* logic to calculate total */);
}, [orders]);
```

By leveraging `useEffect` to calculate the `total`, we are:

1. Making the code more complex and harder to read. `useEffect` is an abstraction that adds yet another "layer" someone has to append to their "context" of the code.

2. Introducing unnecessary re-render of the component by setting the `total` state.

React already re-renders the component when state changes, so we should use that to calculate the `total` _inline_ instead of using `useEffect`.

```tsx
const [orders, setOrders] = useState([]);
const total = // calculate total here!
```

##### What about `useMemo`?

You might be tempted to wrap the `total` calculation with `useMemo` to "make the operation more performant".

Before you do, please keep in mind that the `useMemo` also introduces an overhead:

1. Similarly to using `useEffect`, you are using yet another abstraction.

2. The logic that powers the `useMemo` has to perform calculations and decide whether the output is different or not.

**`useMemo` should be reserved for operations you know are "expensive"**.

#### `refs` vs `state`

The `useRef` hook is not only for DOM nodes. You can also hold values there!

Sometimes, you do not need a change in value to re-render the component.

```tsx
function Timer() {
  const [timeLeft, setTimeLeft] = useState(60);
  const [timerId, setTimerId] = useState<NodeJS.Timeout | null>(null); // ❌ Causes re-renders

  const startTimer = () => {
    const id = setInterval(() => {
      setTimeLeft((prev) => prev - 1);
    }, 1000);
    setTimerId(id); // ❌ Triggers unnecessary re-render
  };

  useEffect(() => {
    return () => timerId && clearInterval(timerId);
  }, [timerId]); // ❌ Effect runs every time timerId changes

  return <div>{timeLeft}s remaining</div>;
}
```

Notice that we do not need to store the `timerId` in the state because **when `timerId` changes, the UI is the same as it previously was**.

Think of the `timerId` as _metadata_ about the state or component that we can put in `useRef`.

```tsx
function Timer() {
  const [timeLeft, setTimeLeft] = useState(60);
  const timerIdRef = useRef<NodeJS.Timeout | null>(null); // ✅ No re-renders

  const startTimer = () => {
    const id = setInterval(() => {
      setTimeLeft((prev) => prev - 1);
    }, 1000);
    timerIdRef.current = id; // ✅ No re-render triggered
  };

  useEffect(() => {
    return () => timerIdRef.current && clearInterval(timerIdRef.current);
  }, []); // ✅ Effect runs only once

  return <div>{timeLeft}s remaining</div>;
}
```

---

In one of the exercises, we use `useEffect` to subscribe to the `scroll` event on the window to calculate the direction of the scroll.

I remembered that the `useSyncExternalStore` might be a good choice for this functionality, so I've tried to replicate what we already have with that hook.

It actually turned out to be quite challenging, because the `useSyncExternalStore` does not pass "previous" state value to the `getSnapshot` function.

```ts
function subscribeToScrollEvents(subscriber: VoidFunction) {
  window.addEventListener("scroll", subscriber);
  return () => window.removeEventListener("scroll", subscriber);
}

// These ones are the key!
const lastScrollY = useRef<number | null>(null);
const lastDirection = useRef<string | undefined>(undefined);

const direction = useSyncExternalStore(
  subscribeToScrollEvents,
  () => {
    const currentScrollY = window.scrollY;

    if (lastScrollY.current === null) {
      lastScrollY.current = currentScrollY;
      return undefined;
    }

    const difference = currentScrollY - lastScrollY.current;

    if (difference !== 0) {
      lastScrollY.current = currentScrollY;
      lastDirection.current = difference > 0 ? "down" : "up";
    }

    return lastDirection.current;
  },
  () => {
    return undefined;
  },
);
```

We do limit the number of re-renders here, which is nice, but we introduce a lot of complexity to the code.

It seems like the `useSyncExternalStore` is very good hook when you do not require the "previous state" for your computation. If you do, you will need to leverage refs which adds a bit of complexity.

### Many different types of diagrams

You can map various relationships within your applications using _diagrams_. As a visual learner, I approve of this method!

You do not have to go super deep or robust here. A simple diagram will often suffice.

**The main point behind making diagrams for you application state or entities is to surface any complexity that might be "hidden" behind the code**.
I like to think about it in terms of **making implicit explicit**.

### Combining and Optimizing State

- When it comes to state changes, _events_ are the source of truth.

  - Think about it - you mostly update the state when user performs some action.

- Deriving and updating the state via _pure functions_.

- Consider _thinking in state machines_.

  - Completely agree here. How many times have you seen the UI jump from "no results" to "here are the results" when clicking on a search bar?

    - This is an example of code debt that stems from lack of proper state handling.

---

The most impactful refactors you can do it so _derive_ boolean flags rather than have them embedded in state.

```tsx
const [status, setStatus] = useState("idle");

const isSubmitting = status === "submitting";
```

**Boolean flags are problematic because they can lead to "impossible states"**. Think `isError: true` and `isSuccess: true`.

On top of that, booleans have a tendency to "multiply". It is very continent to add _yet another boolean flag_ to state of function signature.

### Forms and `useActionState`

You can get rid of a LOT of state by using the `useActionState` to handle form submissions.

Firstly, you most likely do not need a state value for each of the form fields. The `action` prop will pass you the `formData` which you can transform to values.

Secondly, you do not need to define explicit loading states. One of the things the `useActionState` returns is the `isPending` boolean. You can use that value to implement the "loading UI".

### `useReducer`

The `useReducer` hook is quite awesome for _collapsing_ the state into a single object.

Why would you want to do that? The main benefit is that you have **more control over state transitions**. You can define actions and only allow the state to transition to another state given a certain action.

One pattern David showcases in the workshop is putting the state and dispatch function in the context and using those to make state changes instead of drilling "setXX" functions through many layers of components.

```tsx
<Provider value={{ state, dispatch }}>{children}</Provider>
```

**This is nice pattern, but we can make it better!**

The problem with this pattern is that the component consuming the context only to use the `dispatch` function will also re-render when the `state` changes.

Instead of providing the `state` and `dispatch` to a single context, we can split the contexts. This way, when I need only the `dispatch` function, I can skip reading `state` which means that, if the `state` changes, the component consuming the "dispatch context" won't re-render.

```tsx
<DispatchProvider value={{ dispatch }}>
  <StateProvider value={{ state }}>{children}</StateProvider>
</DispatchProvider>
```

### External State Management Libraries

The main motivation for using a library instead of leveraging the built-in hooks is the ability to benefit from the robustness and ease of use of those libraries when your state becomes complex.

All those libraries use the built-in hooks under the hood, but they abstract _a lot_ of complexity and handle edge cases.

**Using a library also eliminates all the conflicts or inconsistencies that rolling your own abstractions might create** – if you are working with other people, it is inevitable for others to create their own abstractions that differ from yours.

- The "central store" approach.

  - A centralized place where all the state lives.

  - You change the state by creating _events_ that then cause the store to change.

  - **Consider using a store approach when you have complex state and complex requirements**.

    - In such cases, having a single place where everything "stems from" is usually a very good thing.

- The "atomic" approach.

  - Pieces of data that can change from anywhere.

  - Atoms are reactive!

  - **Consider this approach for simpler state where you do not care about what or how this value changes**.

### Data Normalization

What does it even mean to _normalize_ data (or state)?

**In this context, _normalization_ usually means "flattening" the state so you can perform O(1) lookups to get a given piece of data**.

Why is this important?

- It makes updating the state much easier, since you do not have to "drill down" into parent to update the child.

- It makes reading the state faster, because you do not have to "drill down" into parent to read the child.

Overall, _data normalization_ makes your life easier.

```tsx
const data = {
  reservations: {
    id: "...",
    flights: [
      {
        id: "...",
        passengers: [
          {
            id: "..",
          },
        ],
      },
    ],
  },
};

const normalizedData = {
  reservations: {
    "...": {
      id: "...",
      passengers: {
        "...": {
          id: "..."
        }
      }
      flights: {
        "...": {
          id: "...",
          passengers: ["..."],
        },
      },
    },
  },
};
```

Notice how much easier it would be to update a given passenger or flight.

Finished part 5 -30:29
