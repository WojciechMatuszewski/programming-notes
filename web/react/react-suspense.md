# React Suspense (the API will change)

## Basic data fetching

What is the most fascinating thing for me about the `React.Suspense` API (at the time of writing this still experimental API), is the fact that we will be **throwing _promises_**. Yup, that's right. So the deal is that by _throwing a promise_ your component will be in so called _suspended_ state. That _promise_ will be caught by `React.Suspense` which will attach a `.then` handler to that _promise_. If the _promise_ resolves, your component will no longer be _suspended_.

Check this out:

```jsx
let data = null;
const pokemonPromise = fetchPokemon("pikachu").then(
  (fetchedPokemon) => (data = fetchedPokemon)
);

function Component() {
  if (!data) {
    throw pokemonPromise;
  }

  return <div>pokemon is fetched</div>;
}

function App() {
  return (
    <React.Suspense fallback={<p>loading</p>}>
      <Component />
    </React.Suspense>
  );
}
```

Since this is the basics of the basics, we are fetching the data at _import time_. This is usually not the case, but nevertheless it should give you some intuition on how the `React.Suspense` works.

### Handling errors

Something will inevitably go wrong with fetching data. There might be a timeout on a lambda function or just network issues related to users device.

Assuming the previous example we could do something like this

```jsx
// code from previous exercise
let error = null;
const pokemonPromise = fetchPokemon("pikachu").then(
  (fetchedPokemon) => (data = fetchedPokemon),
  (err) => (error = err)
);

function Component() {
  if (error) {
    throw error;
    // OR
    return <div>something went wrong</div>;
  }
  if (!data) {
    throw pokemonPromise;
  }

  return <div>pokemon is fetched</div>;
}
```

If you embark on the `throw` route, remember to add the error boundary **before your `React.Suspense`**

## The `createResource` utility

While looking at code which uses `React.Suspense` you might encounter code which looks similar to this:

```jsx
const resource = createResource(fetchPokemon("pikachu"));

function Component() {
  const pokemon = resource.read();
  return <div>{pokemon.name}</div>;
}
```

So the `resource.read` basically does what we did previously, it _throws_ the promise or error or just returns _resolved_ data.

What's nice about it is that the implementation most likely contains a caching mechanism so you are not making a http call every time you fetch something (or do something that is asynchronous)

## Fetch as you render

Tbh I was struggling to understand the name of the pattern. When you look at the components that use the `resource.read` the name _fetch as you render_ does not really makes sense

```jsx
const resource = createResource(fetchPokemon("pikachu"));

function Component() {
  const pokemon = resource.read();
  return <div>{pokemon.name}</div>;
}
```

Like what is _React_ going to render when `resource.read` is still resolving? In this case... _nothing_. But think about a bit more sophisticated example

```jsx
const resource = createResource(fetchPokemon("pikachu"));
const imageResource = createResource(fetchImage("pikachu"));

function Component() {
  const pokemon = resource.read();
  return (
    <div>
      <p>{pokemon.name}</p>
      // uses `.read`
      <LazyImage resource={imageResource} />
    </div>
  );
}
```

Now the _render as you fetch_ makes more sense. You might be able to render the pokemon part while the `LazyImage` is still being resolved. Each time something resolves, `React` is able to go _deeper_ and render more stuff. This is the essence of
_render as you fetch_ pattern.
