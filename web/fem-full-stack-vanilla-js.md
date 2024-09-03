# FrontendMasters – Full Stack Vanilla JS

Learning from [this course](https://frontendmasters.com/workshops/fullstack-vanilla-js/)

## Learnings

- ESM modules are widely supported.

- The [`ntl`](https://www.npmjs.com/package/ntl) library is quite helpful for running scripts you might have in the application.

  - It will create an interactive menu for your scripts based on the contents of `package.json`!

- **To check which _platform_ we are on ("web" or "console") Erick used a check `globalThis.window`**.

  - I love to see `globalThis` used more frequently.

- If you use ESM, meaning you have `type: "module"` in `package.json`, the JS files you import must have the `.js` extension.

  ```js
  import willWorkAsExpected from "bar.js";
  import willNotWork from "baz";
  ```

- Erick decided to leverage [`jsdoc`](https://jsdoc.app/) in the application.

  - I guess that is the necessary evil, since we do not want to introduce any bundlers.

  - Having said that, **I wonder if `jsdoc` becomes obsolete when Node fully supports TypeScript files**.

- The code for adding the rows could be more optimized.

  - Using `innerHTML` is quite inefficient. Using `appendChild` would be a better choice.

  - We are adding each row to the "table body" in a loop. **Instead of creating X "write" operations to the DOM, consider using a [`DocumentFragment`](https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment) APIs**.

    - First, write a bunch of HTML into that fragment, and then "commit" that fragment to the DOM.

- Rather than using `HTML_ELEMENT.value` we could have leveraged the `formData` to get the various inputs data when form submits.

  ```js
  formElement.addEventListener("submit", (event) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    const data = Object.fromEntries(formData.entries());
  });
  ```

  This approach works very well "validation libraries" like `zod` – you can validate the `data` object and infer types based on the result of validation.

- Similar situation to the above while implementing the "Clear" button on the form.

  - Instead of using the native `<button type = "clear">` Erick chose to use `type="button"` and call `.reset()` on the form HTML node.

    - **While I might be missing some context as to why we did not implement the "Clear" button using the native HTML, this shows that even the most experienced developers are sometimes lacking the knowledge in the basics**.

- Eric created a `launch.json` file with a debugger definition before even starting to write the first test.

  - Very interesting practice. I probably should use the debugging capabilities of VsCode more often.

Finished Part 2 -18:58
