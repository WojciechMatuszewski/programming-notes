# Qwik

## My thoughts after the pitch

- Qwik tries to bridge good DX with good performance.

  - A hard thing to do.

- **Qwik is NOT about lazy loading JS**. It is **about lazy execution of JS**.

  - The creator argues that in todays world, it is not about the speed of downloading the JS, but the speed of execution.

  - I'm not sure how I feel about this one. I somewhat agree, but also disagree in some respect.

- The **graph of hydration store is insane and looks very wasteful**.

  - When you learn how it works it really looks like a workaround rather than a "feature".

- **Qwik uses the concept of resumability**. This means **pausing the execution on the server and resuming on the client**.

  - It encodes the state just like in the case of React SCs.

  - The difference is that **it does no replay that state on the client**. The serialized payload is enough to provide interactivity.

- **Qwik makes lazy loading frictionless**.

  - You can lazy-load closures and components. Pretty good.

  - It uses special functions with particular names. You do not have to define the lazy-loading boundaries.

## Working with Qwik

- Similar to the Next 13, Qwik has the concept of layouts.

  - Instead of _children_, the framework uses _slots_.

- Qwik uses the concept of loaders from Remix, but they have a different name for it. Very interesting that the framework mixes and adds it is own twist on the things from the other frameworks/meta-frameworks.

  - You can only have a single _loader_ in a given component file. Like in Remix, you can import other loaders into your loader.

  - There is **a loader that responds to the POST submit for a form**. Again, similar to Remix.

  - There is also **a loader to responds to both POST and GET requests**. Very useful as a middleware, especially for authorization.

- The **most interesting thing about Qwik** for me is the fact that **the framework automatically lazy-executes the JS**.

  - It is almost as if the components were lazy loaded. **If the component is NOT visible on the screen, its JS will NOT run**.

- There is **no clear client vs server boundary for components code**.

  - Some parts of the component might run on the server and resume on the client.

  - Some method have "client" or "server" within the name.

- You can control if you want to navigate using SSR or client-side via the usage of `a` tag (server-side) or `Link` component (client-side).

- The **ability to defer JS execution in Qwik is insane**.

  - By using the `$` you can defer functions inside the component. This is very new to me.

  - As the creator of the Qwik explained, it is due to a rust pre-compile step that goes in and splits the code in the chunks.

  - This pre-compile step alongside with a heuristic on the service worker side of things **produces a very optimized bundles that the framework will fetch automatically based on user interactions**.

    - If you thought that the React mental-model is hard to grasp, Qwik makes React look like a toy.

<!-- Finished day 1 -->
