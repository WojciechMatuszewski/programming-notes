# The modern web

## Hydration

> Taking notes based on [this video](https://www.youtube.com/watch?v=iR5T2HefqKk).

### Defining hydration

- The definition on wikipedia is incomplete. While it talks about attaching events, there is much more going on.

  - The definition fails to mention to most crucial bit – the state of the application and how it is populated when the website is loaded.

- You can think of hydration as **the work required by the client to bring the SSR app to the same state as if it was client-side rendered**.

  - This means sending the serialized state, HTML AND the JS required to make the application work. This means we have to send a lot of stuff twice!

### Why efficient hydration is hard

> Notes from [this article](https://dev.to/this-is-learning/why-efficient-hydration-in-javascript-frameworks-is-so-challenging-1ca3)

- Hydration and state preservation are complex. The framework has to encode the application state into the HTML and then send it to the browser. This makes the HTML document size bigger.

- Initializing the framework from the state embedded in the HTML takes time. If the amount of state is large, you might freeze the browser.

- There are various ways to defer the hydration work. Those include  generating static sites, using progressive or out-of-order hydration or creating "islands of reactivity".

> A note regarding the island architecture. As good as it might seem, if separate islands needs to share the state, you might run into problems. You cannot pass the state down, since that would make the parent of those island and island as well. The most common solution to this problem I've seen is to use a store for a given set of islands. Read [this blog for more information](https://frontendatscale.com/blog/islands-architecture-state/).

### Hydrating SPAs

- Deferring hydration is not an option. Hydration can take a lot of time. If the user interacts with a page, would you want them to wait?

  - I guess it's about tradeoffs.

- People argue that progressive enhancement is enough. It is not. Okay, your website might work without JS, but it is most likely a bit suboptimal.

  - When the app hydrates, the user will "feel" the hydration (a minor freeze, or some kind of glitch) if you defer it.

### Different kinds of hydration

There are spectrums of hydration. One can mix them accordingly. For example, you can have _eager_ and _partial_ hydration implemented in the same framework.

#### Eager vs Progressive hydration

- _Progressive hydration_ would be loading the JS and hydrating on event/interaction.

- _Eager_ would be to hydrate everything at the beginning.

#### Partial vs Full

- _Partial hydration_ is to send only the JS needed for interaction for a given page. This combines the knowledge that the server and the client has.

  - If you do not need the JS, the framework will only send the serialized HTML/data for given pieces of the page.

- _Full hydration_ is to send all the JS for every component.

#### Replayable vs Resumable

- _Resumable hydration_ is when you **do NOT** repeat any work that the server already done.

  - Keep in mind that so far we have been running the JS that already been run on the server. This is quite revolutionary.

  - This technique requires the serialization of the whole state. Both the app state and the framework state.

- _Replayable hydration_ is when you re-store the state of the components as the server seen them.

  - This means you only ship the app state and "replay" the framework to a given state.

## Data fetching

> Notes from [this video](https://youtu.be/8ObxzMSIqKA)

- It seems like we are making a full circle in terms of data fetching.

  1. API as HTML

  2. API as XML/JSON – AJAX

  3. API as State – Redux

  4. API as Normalized Cache – GraphQL

  5. API as Query Cache – React Query

  6. API as Page Cache – gSSP / Loaders

  7. API as JSX(HTML) – RSCs, HTMX

  Notice that the `no.7` is really `no.1` but better.

- Data fetching **ties very closely with the router the application uses**. Here is a very brief evolution of client-side routing.

  1. Simple client router – no consideration for data fetching

  2. Ember Router - MVC for the Browser – had hooks to fetch the data when the route changes. It was a game-changer

  3. React router 4 regression - "we are not responsible for data fetching anymore"

      - One could do _Link preloading_, but that did not mean _data preloading_

  4. Sapper Router, Nuxt Router, Solid's App Router, React Router 6

- Data fetching and the SSR also evolved. At the beginning, the APIs were synchronous (`renderJSXtoString`). How does data-fetching fit in this scenario?

  1. _Prepass_ would render till it hits a data-fetching call. Then it would wait for the data to finish fetching and continue rendering. This is very slow and creates a lot of waterfalls.

  2. _gSSP/Loader_ would parallelize the data fetching BEFORE rendering. This means that you avoid SOME waterfalls, and the rendering is not interleaved with data-fetching.

  3. _Suspense with hoisting_ allowed us to start rendering and data-fetch AT THE SAME TIME. This means we are shifting the first render left. Unlike the _gSSP/Loaders_ where we first had to fetch all the data to start rendering.

      - Of course, **one can also make a mistake here**. The key is to start data-fetching **at the route level or even higher**. If you are fetching in your components, then we are pretty much back to the _prepass_ scenario.

  4. _RSCs_ (blocking) is not better than the _Suspense with hoisting_. In fact, it is much worse. That is why the Remix people were skeptical about RSCs in the first place. Since the data-fetching is not hoisted at the top, we cannot make parallel requests (one could do them, but only inside a single RSC which is a best practice).

- In an ideal world, we would be able to use _Context_ API and RSCs at the same time. If that would be the case, we could hoist the data fetching to the top, and not have to deal with prop-drilling.

  - Ryan thinks about solving this problem with signals. **Signal will only block when read**, so you can "mark" the resource as fetching, but then continue down the tree. Only stop rendering and fetch when the signal is read in the JSX.
