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

## How does Qwik work under the hood?

- By default, all frameworks similar to React have to **execute all the code to figure out if there are listeners attached to a given elements**.

  - This is very wasteful. Even if you **lazy load, the code that you lazy loading (on the same route) will be executed instantly**.

    - The only case where lazy loading makes sense is if the lazy-loaded code is behind a condition or on a different route.

  - The **fundamental assumption with existing framework is: all code is available when they run**.

    - This is NOT the case in Qwik. Such assumption creates performance constraints.

- Everything in Qwik is asynchronous. You can embed _promises_ right into the JSX and they will be resolved.

  - This means you can _fetch as you render_ without any ceremony. This is **huge**.

- Since you can lazy-load functions, **the serializer must be quite complex since serializing functions is very hard**.

  - Think of all the variables in closures and the logical scope.

- **Due to all the magic happening behind the scenes** you might end up in a scenario where a seemingly regular usage of the variables might trigger a deoptimization.

  - I'm not talking about the `deopt` call from Node. I'm talking about Qwik downloading the whole component instead of skipping it.

  - That is the price you pay for having so many great features?

  - **Even with deoptimization behavior, the baseline is much better than the "current" frameworks**.

- **Qwik is a truly reactive framework**. If you update the `signal`, only the parts that use the signal will update.

- Qwik has the notion of an `useEffect` (without dependencies) and `useLayoutEffect` (without dependencies).

  - The main difference are the **name of the hooks, and their granularity**. Some run before render on the server, some on the client and such.

  - Note that `useEffect` runs on the client and on the server AFTER rendering. Running `useEffect` BEFORE rendering on the server and on the client is not possible. Qwik have you covered here.

<!-- Finished day 2 part 2 -->
