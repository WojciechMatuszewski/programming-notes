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

Finished Part 2 -12:52
