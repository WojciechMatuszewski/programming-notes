# Epic Web Dev notes

## Fullstack Foundations

- The **`button` element has the `type="reset"` attribute**.

  - This is super useful if you want to reset the form.

  - Please keep in mind that you can hide it and invoke it programmatically whenever you see fit.

    - This seem to be a good alternative to the `key` prop on the form! Sometimes it is hard to come up with a valid `key`.

- Kent showcases a great example where using plain form-based mutations saves us a lot of trouble.

  - Picture having to perform a network request if user clicks on a button.

    ```jsx
    <button
      onClick={() => {
        fetch("...");
      }}
    >
      Click me
    </button>
    ```

    There is a lot of things to consider.

    1. What if the user clicks the button twice?

    2. What if the request fails?

    3. What if the user clicks the button and then navigates away?

    ...

    There are a lot of edge cases! **Here is where wrapping this `button` in a form comes in handy**. If you do that, you do not have to worry about those things as they will be taken care of by the browser!

    ```jsx
    <Form method="POST">
      <button type="submit">click me</button>
    </Form>
    ```

- While I really enjoy working with the _Actions_ API, I'm not a huge fan of how Remix solved the problems of having multiple actions in the same file.

  - **Using the `name` property on a button is very neat**.

  - My main gripe is that we cannot export different functions, but rather have to perform a switch statement based on the `intent` (or other name that we set).

- The way scroll restoration works in `remix` and `react-router` is quite fascinating.

  - The `react-router` uses the _page hide_ event to save the last known scroll position for a given history entry. Each position is assigned an ID. This ID is derived from a given entry in the history (think a hash function).

  - `react-router` will then search through the saved entries, and if one matches, restore the scroll position for that given entry.

  - `remix` augments this code by injecting inline scripts to ensure the behavior works as expected with SSR.

- Kent demonstrated a very useful technique for sharing ENV variables. Instead of relying on a build tool which, in some cases makes all the values from `process.env` available, he used a _loader_ to inject those to `window`.

  ```jsx
  export async function loader() {
    return json({ env: getEnv() });
  }

  export function App() {
    const data = useLoaderData();

    return (
      <script
        dangerouslySetInnerHTML={{
          __html: `window.ENV=${JSON.stringify(data.ENV)}`,
        }}
      />
    );
  }
  ```

  While this snippet uses remix-specific APIs (the `loader` function), I think it translates well to any other framework.

- Interestingly `remix` will not prefetch all the links on the page by default. This is in contrast to Next.js.

  - What I like the most about `remix` implementation is the control I have. There are different options for `prefetch` like `intent` or `render`. Next.js does not have these options – it only exposes a boolean prop which might be problematic in some cases. For example when you want to prefetch all links on the page but only on hover.

- While it requires a bit more work, I find the `remix` way of exposing the status of the form a bit better.

  - The problem with the `useFormStatus` is that it **needs to be called in a component living inside a `form`**. In some cases, we do not have the luxury of putting the submit button inside the `form` tag. Instead we use the `id` and `formId` props to associate the button with the form.

- The way to handle metadata in `remix` is pretty similar to APIs exposed by Next.js.

  - One thing that I worries me is the crazy route names Kent had to create. He is using some kind of library for `remix` to make the routes "flat". I guess it is just a matter of getting used to those.

- Kent mentions a term I was not familiar with – the **_Splat route_**. The _splat route_ is a _wildcard route_. It sounds kind of cool!
