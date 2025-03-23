# Frontend Masters Intermediate React v6

Learning from [this course](https://intermediate-react-v6.holt.courses/).

## React Render Modes

- **Client Side**: the traditional single page application approach.

  - The server does little work. You might have an API running on the server or not. You might have some assets served from the server or not.

  - Historically, we have been writing applications this way for very long time.

  - **Still viable**. All depending on the use-case.

- **Static Site Generation**: render the JSX as static markup.

  - During the build step, you extract all the JSX as static markup and put it into a `.html` file.

    - This greatly speeds up the _perceived_ performance, as the user is able to see the HTML content right away.

    - This also helps with SEO, but might be less relevant nowadays, since Google is now able to index SPAs.

  - **The initial shell will NOT be "personalized" to the given user** since the HTML is extracted at _build time_.

    - This is the main drawback of SSG as compared to SSR.

- **Server-side Rendering**: **execute** the application when the request happens and return HTML.

  - It is not a silver bullet. In some cases you can make the application slower!

    - Hydration _might_ take a long time. It all depends.

## RSCs

- **RSCs can work with or without SSR**.

  - The "server" meaning is fluid. It could be a "real" server, or it could be a build-step.

- **The server components are NOT HYDRATED**. This means that they can contain code that is node-specific. This code will not be re-executed in the browser.

- In the course, we had scripts to produce two "versions" of our application. One for the server, one for the client.

  - For the "server version", we used the [conditions API](https://nodejs.org/api/cli.html#-c-condition---conditionscondition) which is quite fascinating feature.

    - Package authors can specify which file is the "entry" file for a given condition.

- When working on the internals of RSCs, you might encounter a term "React Flight".

  - To the best of my knowledge, this term was used to describe what is now knows as RSCs.

## Transitions

- The purpose of _transitions_ is to enable React to do something "big" behind the scenes, while also allowing for the user input.

  - **This only works if the work that happens in the background is interruptible. If you lock the main thread, nothing going to save you**.

- React added support for _async transitions_ which is nice, but they come with a [one potential edge case](https://react.dev/reference/react/useTransition#react-doesnt-treat-my-state-update-after-await-as-a-transition).

  ```js
  startTransition(async () => {
    await someAsyncFunction();

    // This update is NOT going to be treated as transition.
    setPage("/about");
  });

  startTransition(async () => {
    await someAsyncFunction();

    // Correct way of doing it.
    startTransition(() => {
      setPage("/about");
    });
  });
  ```

  Why would we need to nest the `startTransition` calls? I'm **unsure**, but it might have to do with slicing updates into interruptible chunks. React might need to interrupt the work _after_ the async call happened, but BEFORE the `setPage` run.

## Optimistic UI

- The optimistic UI is now built-in into React!

  - Use the `useOptimistic` hook.

  - **The updates you make to the optimistic value MUST happen within a transition or action**.

    - Remember that form actions implicitly mark the code as transitions!

- The nice thing about the `useOptimistic` is that it will re-synchronize with the "initial" value AFTER the transition is finished.

  - This means you you get rollback functionality for free. If you do not update the "main" value in the transition, the UI will rollback to the initial value before the transition happened.

## Deferred Values

https://intermediate-react-v6.holt.courses/lessons/deferred-values/what-are-deferred-values

Part 7
