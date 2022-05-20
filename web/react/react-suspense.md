# React Suspense

It is safe to say that most people writing react code daily get sick of writing code familiar with the following.

```ts
const { loading, data, error } = fetchResource();
if (error) {
  return <p>Error</p>;
}

if (loading) {
  return <p>Loading</p>;
}

// Do something with the data
```

Luckily for us, React team decided to introduce the Suspense API. The idea behind the Suspense API is to make a "_loading state_" a first-class citizen of React programming paradigm.

```jsx
<Menu>
  <Suspense fallback={<p>Loading...</p>}>
    <Account />
  </Suspense>
</Menu>
```

The scope and the functionality changed in a significant way in-between React 16 and React 18 releases. Let us examine the differences next.

## How Suspense worked in React 16

In React 16, the Suspense supported a single-use case – code splitting via the `React.lazy` API, and it **did not work on the server**.

The Suspense API controlled the visibility of the content **via `display:none` CSS property**. This behavior caused some people a lot of issues, especially in cases where parent components of the suspending child are listening to the "mount" event. For more information regarding this behavior, refer to [this link](https://github.com/reactjs/rfcs/blob/main/text/0213-suspense-in-react-18.md#behavior-change-committed-trees-are-always-consistent).

Since this version of React guaranteed that if a component started rendering, it would finish rendering no matter what, the Suspense API was limited in scope. It does not perform _placeholder throttling_(a feature where the Suspense wrapper will not render the placeholder immediately) or streaming.

As per [this post](https://github.com/reactwg/react-18/discussions/7) React team referred to this version of Suspense as _"legacy Suspense"_. Interesting.

## How Suspense works in React 18

React 18 improved the capabilities of the Suspense feature.

The most notable change is that **the Suspense works on the server!**. This change is thanks to the new server-side renderer that is asynchronous and **capable of streaming data to the browser** (what SSR streaming in React context refers to).

Next, we have **new features** like **_placeholder throttling_ and transitions**. The new semantics of "I can interrupt any render and do other work" makes all these features possible.

I'm most excited about **data fetching with Suspense** and **_placeholder throttling_**. All of these and more are the subject of further discussions below.

### Effects and Suspense

It turns out that **effects will NOT fire unless the Suspense boundary the component is wrapped with finishes suspending**. This is an important detail as it guarantees `useEffect` and `useLayoutEffect` stability.

```jsx
function MyComponent() {
  const [finishedSuspending, setFinishedSuspending] = useState(false);

  return (
    <Suspense fallback={<p>Loading...</p>}>
      <Suspender />
      <Lifecycle onEffectFired={() => setFinishedSuspending(true)} />
    </Suspense>
  );
}

function Lifecycle({ onEffectFired }) {
  useEffect(() => {
    onEffectFired();
  }, [onEffectFired]);

  return null;
}
```

This guarantee makes sense as it would suck if the effect would fire before React is done with the Suspense Boundary.

#### Where would I use this fact

[This video](https://www.youtube.com/watch?v=sOkgIa560qM) walks through one fascinating use case – rendering the `Suspense` component conditionality based on the lifecycle state (if it is a first "render" or not).

The idea is that you could have a different "loading experience" if it is the first time you visit the application (one single big spinner) vs. when some of the resources are already available and when the data fetching is done in a given part of an application.

## Placeholder throttling

When I first read about this "feature," my excitement levels were very high. How often have you used `Suspense` to lazy load your component, only for the `placeholder` prop to "flash" for a split second, creating a suboptimal experience? I'm describing a real issue on faster connections, where downloading JS takes a split second.

I thought React 18 would solve that problem for me, as the first pass of the Suspense API introduced the `timeout` property that got removed. It turns out **the "issue" persists in React 18, and if you want to ensure that the fallback does not "flash", you have to implement it yourself**.

Having said that, this is much easier to achieve in React 18 because React 18 will not insert DOM nodes of the half-committed tree, so the structure of your HTML stays intact.

### Making sure placeholders do not flash manually

So, the scenario is as follows: you use the `lazy` API to code split your application. While testing the application locally, you have noticed these annoying "flashes" of the `placeholder` content. You are in your office, and the internet there is high-speed, so the bundles download in a split second.

```jsx
const LazyList = lazy(() => import("./List"));

return (
  <Suspense fallback={<p>Loading...</p>}>
    <LazyList />
  </Suspense>
);
```

What can we do about it? We can do two things.

1. Instead of rendering the `<p>Loading...</p>` we could render `null`. Users will not observe any "flash" of content because you render `null` as the fallback.

2. Add artificial delay inside the callback function passed to the `lazy` function.

The following is the first option.

```jsx
const LazyList = lazy(() => import("./List"));

return (
  <Suspense fallback={null}>
    <LazyList />
  </Suspense>
);
```

Problem solved for high-speed connections, but what about slower ones? If it does take some time to download the `List` bundle, the user will not see any visual feedback that this is happening. Such experience might leave him confused and ask whether the application is working.

To solve this particular problem, one might look into implementing the second point.

```jsx
const wait = (ms) => {
  return new Promise((resolve) => {
    setTimeout(() => resolve(undefined), ms);
  });
};

const List = lazy(async () => {
  const result = await Promise.all([wait(1000), import("./List")]);

  return result[1];
});

return (
  <Suspense fallback={null}>
    <LazyList />
  </Suspense>
);
```

I think this implementation is an excellent compromise between _not rendering anything_ and _making sure the placeholder does not "flash"_.

### When does placeholder throttling occurs

> Before we start, know this – I'm not sure whether what I'm about to talk about is the so-called _placeholder throttling_. I've tried searching for an example but could not find any. I'm basing this section on my gut instinct and understanding of the React 18 features.

Before we start, know this – I'm not sure whether what I'm about to talk about is the so-called _placeholder throttling_. I've tried searching for an example but could not find any. I'm basing this section on my gut instinct and understanding of the React 18 features.

If you have ever used the new `useTransition` (and the `startTransition` function) hook, you might have noticed that the old content stays on the screen while the new content loads. After a certain period (which is NOT configurable), the new content "reveals" itself – all due to the _transitions_ and the ability to "render multiple screens at once".

**React will "keep" the previous screen while rendering the next one. No `placeholder` is shown during this operation as React uses the "previous" screen as the placeholder**. I **think** this is the **_placeholder throttling_** that React team mentions in GitHub discussions.

```jsx
import { Suspense, lazy, useState, useTransition } from "react";
import { ErrorBoundary } from "react-error-boundary";

const wait = (ms) => {
  return new Promise((resolve) => {
    setTimeout(() => resolve(undefined), ms);
  });
};

const getComponentResource = () => {
  const max = 1000;
  const min = 200;
  const delay = Math.random() * (max - min) + min;

  return lazy(async () => {
    const result = await Promise.all([
      wait(delay),
      Promise.resolve({ default: EagerList }),
    ]);

    return result[1];
  });
};

function App() {
  const [resource, setResource] = useState(() => getComponentResource());
  const onNewResource = () => {
    startTransition(() => setResource(getComponentResource()));
  };

  const [isPending, startTransition] = useTransition();

  const Comp = resource;
  return (
    <div>
      <button onClick={onNewResource}>New resource</button>
      <div style={{ opacity: isPending ? 0.4 : 1 }}>
        <ErrorBoundary fallback={<p>Error!</p>}>
          <Suspense fallback={<p>Loading...</p>}>
            <Comp />
          </Suspense>
        </ErrorBoundary>
      </div>
    </div>
  );
}

const EagerList = () => {
  const number = Math.random();
  return <div>{number}</div>;
};
```

React will render the `fallback` prop when you first load the application since there is no "previous" screen to fall back to. **Upon clicking on the _New resource_ React will use the "previous" screen as the fallback**. Pretty neat!

### "Previous" screen as the fallback

I've alluded to the _transitions_ behavior in previous sections where React would keep the "previous" screen as the fallback. Let us consider what would happen if we were to have a single _Suspense boundary_ at the top of our application.

```jsx
return (
  <ErrorBoundary fallback={<p>Error!</p>}>
    <Suspense fallback={<p>Loading...</p>}>
      <div>
        <MoreComponents />
        <button onClick={onNewResource}>New resource</button>
        <div style={{ opacity: isPending ? 0.4 : 1 }}>
          <Comp />
        </div>
      </div>
    </Suspense>
  </ErrorBoundary>
);
```

The application would still be reactive, and all the events would work. This is the "magic" (albeit I do not like this term in the context of software development).

## Data fetching and Suspense

We have only considered code-splitting as the use case for the Suspense, but the Suspense API can also help us with data fetching and getting rid of these pesky `if(isLoading)` checks.

The last time I checked, Suspense's internal implementation relied on throwing promises caught by the nearest Suspense boundary. If the library you are using supports Suspense, you can use that to your advantage!

```jsx
const User = () => {
  const { data } = useFetchUser();

  return <p>User name is {data.name}</p>;
};

const App = () => {
  return (
    <ErrorBoundary>
      <Suspense fallback={<p>Loading user...</p>}>
        <User />
      </Suspense>
    </ErrorBoundary>
  );
};
```

Notice that I'm accessing the `data` as if the `useFetchUser` was synchronous. I'm able to do this because **React can now "pause" rendering of the `User` component until the `useFetchUser` resolves**. No more `if (isLoading)` checks. Now it's the role of the `fallback` prop on the `Suspense` component to take care of the loading for us. Pretty neat!

One additional detail – notice the `ErrorBoundary`. If the `useFetchUser` fails, the error will propagate to the nearest `ErrorBoundary` (or crash your application).
