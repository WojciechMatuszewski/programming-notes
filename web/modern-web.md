# The modern web

## Hydration

> Taking notes based on [this video](https://www.youtube.com/watch?v=iR5T2HefqKk).

### Defining hydration

- The definition on wikipedia is incomplete. While it talks about attaching events, there is much more going on.

  - The definition fails to mention to most crucial bit – the state of the application and how it is populated when the website is loaded.

- You can think of hydration as **the work required by the client to bring the SSR app to the same state as if it was client-side rendered**.

  - This means sending the serialized state, HTML AND the JS required to make the application work. This means we have to send a lot of stuff twice!

### Hydrating SPAs

- Deferring hydration is not an option. Hydration can take a lot of time. If the user interacts with a page, would you want them to wait?

  - I guess it's about tradeoffs.

- People argue that progressive enhancement is enough. It is not. Okay, your website might work without JS, but it is most likely a bit suboptimal.

  - When the app hydrates, the user will "feel" the hydration (a minor freeze, or some kind of glitch) if you defer it.

### Eager vs Progressive hydration

- _Progressive hydration_ would be loading the JS and hydrating on event/interaction.

- _Eager_ would be to hydrate everything at the beginning.

[Finished here](https://youtu.be/iR5T2HefqKk?t=6391)
