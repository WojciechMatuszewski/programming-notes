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

## Placeholder throttling

When I first read about this "feature," my excitement levels were very high. How often have you used `Suspense` to lazy load your component, only for the `placeholder` prop to "flash" for a split second, creating a suboptimal experience? I'm describing a real issue on faster connections, where downloading JS takes a split second.

I thought React 18 would solve that problem for me, as the first pass of the Suspense API introduced the `timeout` property that got removed. It turns out **the "issue" persists in React 18, and if you want to ensure that the fallback does not "flash", you have to implement it yourself**.

Having said that, this is much easier to achieve in React 18 because React 18 will not insert DOM nodes of the half-committed tree, so the structure of your HTML stays intact.

## Data fetching with Suspense
