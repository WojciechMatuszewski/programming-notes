# React Suspense (the API will change)

## Basic data fetching

What is the most fascinating thing for me about the `React.Suspense` API (at the time of writing this still experimental API), is the fact that we will be **throwing _promises_**. Yup, that's right. So the deal is that by _throwing a promise_ your component will be in so called _suspended_ state. That _promise_ will be caught by `React.Suspense` which will attach a `.then` handler to that _promise_. If the _promise_ resolves, your component will no longer be _suspended_.

Check this out:

```jsx
let data = null;
const pokemonPromise = fetchPokemon("pikachu").then(
    (fetchedPokemon) => (data = fetchedPokemon),
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

## `useTransition`

As of time of writing this, the API is too unstable to even write about it. Just know that it is a hook, and the _transition_ part comes from the fact that you will be transitioning from pre-fallback to fallback state.

## Lazy loading images with _resource_ concept

You can load your data and images at the same time, pretty neat. For that we need to know how to preload an image, it's pretty simple

```js
function preloadImage(src) {
    const img = document.createElement("img");
    img.src = src;
    img.onload = () => console.log("loaded");
}
```

Now par this with the `createResource` function

```js
function createImageResource(src) {
    return createResource(
        new Promise((resolve) => {
            const img = document.createElement("img");
            img.src = src;
            img.onload = () => resolve(src);
        }),
    );
}
```

Then you can create 2 resources when you want your data, an image and the payload

```js
function createPayloadResource(pokemonName, src) {
    return {
        data: createPokemonResource(pokemonName),
        image: createImageResource(src),
    };
}
```

When in your component

```jsx
function PokemonInfo({ pokemonName, src }) {
    const pokemonResource = createPayloadResource(pokemonName, src);
    const pokemon = pokemonResource.data.read();

    return (
        <div>
            <p>{pokemon.name}</p>
            <img src={pokemonResource.image.read()} />
        </div>
    );
}
```

Neat!

## `React.SuspenseList`

Coordinating loading states is not fun. Actually it's a pretty hard problem, and I would even go as far as saying it's probably something you should leave to 3rd party libraries or... the new `React.SuspenseList` API.

So you have a page with multiple sections, maybe you want to load them lazily, but you end up with having like 10 loaders on your page. You can wrap your `React.Suspense` calls with `React.SuspenseList` and fix that.

```jsx
<React.SuspenseList revealOrder = "together">
  <React.Suspense fallback = {<p>loading...</p>}><Foo/><React.Suspense/>
  <React.Suspense fallback = {<p>loading...</p>}><Foo/><React.Suspense/>
  <React.Suspense fallback = {<p>loading...</p>}><Foo/><React.Suspense/>
</React.SuspenseList>
```

The code above makes it so that the components are loaded together - there will be only 1 `loading...` string visible on the screen. Pretty neat right?

The API is not finalized so I will not be going through all the options, take a look into docs (which are probably different than the ones I'm currently look at).
