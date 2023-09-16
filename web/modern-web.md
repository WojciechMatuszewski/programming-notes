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
