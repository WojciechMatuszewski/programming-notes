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
