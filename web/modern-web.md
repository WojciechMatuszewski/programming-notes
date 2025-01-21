# The modern web

## Server side rendering & time

> Related to [this great blog post](https://next-intl-docs.vercel.app/blog/date-formatting-nextjs)

As soon as you decide to SSR your page, you have created a problem – the server is most likely in different timezone than the user who is requesting the page. If you do not address this difference, the time you display on the website might be incorrect or users might experience a "flicker" of UI related to displaying time.

What can we do about it? **The most pragmatic solution seem to be keeping the timezone as UTC and and formatting the dates accordingly on client-side**.

1. Make sure that your server, and the initial HTML payload _always_ returns dates in UTC timezone, using the same format.

2. (Optional) On the client, take that date, and re-format it accordingly to the client current timezone.

The step number two is optional, as you might choose to always display the time in UTC timezone. It is up to you!

**To avoid hydration mismatches when performing step two, make sure you run this logic _after_ hydration is finished**. Consider adding a placeholder element when the UI is hydrating.

### Example

Consider the following

```tsx
"use client";

function ClientComponent() {
  const date = new Date().toISOString();
  return <span>{date}</span>;
}
```

If we were to SSR this component, the initial HTML output would contain the _server_ formatted time, and the HTML after hydration would contain the _client_ formatted time. This is so-called **_hydration mismatch_**.

If displaying the date based on the location of the client is not important, consider passing the date as prop!

```tsx
function ServerComponent() {
  const date = new Date().toISOString();
  return <ClientComponent date={date} />;
}

// --- //

("use client");

function ClientComponent({ date }: { date: string }) {
  return <span>{date}</span>;
}
```

Now, the date is consistent with the server time.

What happens if you have multiple such components? **If you initialize the date in multiple _server components_, there might be a slight misalignment in time in client components, as different _server components_ might take longer/shorter to render**.

To prevent this from occurring, use some-kind of per-request cache!

## Hydration

> Taking notes based on [this video](https://www.youtube.com/watch?v=iR5T2HefqKk).

### Defining hydration

- The definition on wikipedia is incomplete. While it talks about attaching events, there is much more going on.

  - The definition fails to mention to most crucial bit – the state of the application and how it is populated when the website is loaded.

- You can think of hydration as **the work required by the client to bring the SSR app to the same state as if it was client-side rendered**.

  - This means sending the serialized state, HTML AND the JS required to make the application work. This means we have to send a lot of stuff twice!

- **Hydration is ordered**. It must start at the entry point of the application and work itself towards leaves. Hydration **can not skip some parts of the tree**. It can prioritize different parts, but in the end, the whole tree has to be hydrated.

  - **Hydration might cause your lazy-loaded components to load eagerly**. Unless the _lazy-loaded_ component is not on the current page (maybe different route, or is not initially rendered), hydration has to consume all components to uncover the components hierarchy.

  [You can read more about this topic here](https://www.builder.io/blog/hydration-sabotages-lazy-loading).

### Why efficient hydration is hard

> Notes from [this article](https://dev.to/this-is-learning/why-efficient-hydration-in-javascript-frameworks-is-so-challenging-1ca3)

- Hydration and state preservation are complex. The framework has to encode the application state into the HTML and then send it to the browser. This makes the HTML document size bigger.

- Initializing the framework from the state embedded in the HTML takes time. If the amount of state is large, you might freeze the browser.

- There are various ways to defer the hydration work. Those include generating static sites, using progressive or out-of-order hydration or creating "islands of reactivity".

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

## CloseWatcher API

> [Based on this great blog post](https://logaretm.com/blog/fix-your-annoying-popups-with-the-closewatcher-api/)

Have you ever found yourself having to implement "close the thing on ESC" logic? You most likely did. Have you ever tested that logic on mobile and via screen readers? **You most likely did not**. And, sadly, it most likely does not work there.

Here is where the `CloseWatcher` API comes to aid – this API exposes an _universal_ way to "listen" to _close request_ and react accordingly.

No need to have a special logic for `ESC` key for inputs. No need to have special logic for handling multiple modals and closing them sequentially when multiple of them are opened.

```ts
const watcher = new CloseWatcher();

function onOpenClick() {
  watcher.onclose = () => {
    // code to close the thing
    watcher.destroy();
  };
}

function onClose() {
  watcher.requestClose();
}
```

Notice that we are **creating the _watcher_ instance every time we open something**. This gives us the ability to handle "stacking" – closing only the "latest" thing that was opened. In addition, **you do not have to setup any key-specific listeners yourself**. All is handled by the `CloseWatcher` API.

You can [read more about the API on MDN](https://developer.mozilla.org/en-US/docs/Web/API/CloseWatcher). At the time of writing, this API is experimental, but given how useful it is, I bet it will not be experimental for long.

## View Transition API

The _View Transition API_ allows you to _transition_ between the _old_ and the _new_ content of the webpage.

**The best part of this API is that you have control over HOW the content transitions**. It could be a subtle fade animation, or something more fancy.

You can go really deep with configuration for this API, but you can also get started with a couple of lines of code.

### SPA View Transitions

By default, the browser will cross-fade the elements. All you have to do is to use the `startViewTransition` API in JS.

> Read more about [SPA View Transitions here](https://developer.chrome.com/docs/web-platform/view-transitions/same-document#the_default_transition_cross-fade).

```js
function performAction() {
  startViewTransition(() => {
    updateDOM();
  });
}
```

### MPA View Transitions

At the time of writing, the support for those is not that good, but nevertheless, it is worth exploring this option.

> Read more about [MPA View Transitions here](https://developer.chrome.com/docs/web-platform/view-transitions/cross-document).

To opt-in, each page has to have the following CSS:

```css
@view-transition {
  navigation: auto;
}
```

While I have zero experience with this API, it seems like the effort involved to make MPA _View Transitions_ to work is a bit greater than the SPA.
