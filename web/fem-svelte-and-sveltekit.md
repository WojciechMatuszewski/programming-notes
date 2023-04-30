# Svelte & SvelteKit

Going trough [this course](https://frontendmasters.com/workshops/svelte-sveltekit/).

## What is Svelte

- Svelte is **both a framework and a language**. In that sense it is pretty unique.

  - Other frameworks do not have an _advanced language compiler_.

## Components in Svelte

- The file **encapsulates both the markup, styles, and the JavaScript**.

  - I still remember the old days of Angular 2.x where these were separate things.

  - The **styles are scoped to a given component**. This is pretty nice!

- Writing everything in a single file **feels weird without having a function that denotes the component boundary**.

  - In JSX you have the "component function", in Svelte that is not the case.

### Props

- Even weirder than the `$` symbol.

### Markup

- Ok, it seems like the **Svelte components use some kind of extended version of HTML**. Again Svelte is also a language!

  - It looks similar to handlebars but it is not the same thing.

## Reactivity

- The model of reactivity is quite wild. You **can directly mutate variables**. The mutation will be reflected in the output.

  - Completely different than the "vanilla" React. Of course nowadays something similar is possible with signals.

- The model of _computed values_ uses labels – `$`. Again, the syntax is pretty weird but mostly because I'm not used to it.

  - Keep in mind that Svelte is a language. It is not JavaScript per-se.

- **Effects also use the `$` symbol. Holy moly this is crazy**. Look at the example below. Mind blown.

    ```text
    <script>

      let count = 0;

      $: if (count >= 10) { // LIKE WHAT!
        console.log(`count is too large!`);
        count = 0;
      }

      function handleClick() {
        count += 1;
      }

    </script>
    ```

- Keep in mind that **reactivity is triggered by assignment**. This works well for objects and primitives, but **does not work well for arrays and objects**.

  - A simple `push` will not update the output.

  - The same applies to a nested object where you assign to a given property.

  - In that sense, this logic is different than signals / proxies which should track updates in a nested fashion.

- Svelte works a bit differently than React where each time you change something in the parent, all the children re-render.

  - In Svelte you must assign variables to be _reactive_ using the `$`.

    - I'm personally not a fan of this approach as it can introduce subtle bugs, especially across different components.

### Asynchronous events and rendering

- Svelte seem to have a concept similar to _Suspense_ where you can render a fallback while a promise is pending.

  - Of course this is done via custom markup directive(?) – the `{#await}` block.

  - It also handles errors via the `{:catch}`, similar how you can use the `ErrorBoundary` in React.

Finished day 1 part 1 55:46
<https://learn.svelte.dev/tutorial/await-blocks>
