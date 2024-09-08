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

  - Very interesting practice. I should probably get more accustomed using _launch configurations_ in VSCode.

- **TIL that node has the `--watch` flag**. It seem to be working pretty great.

  - I bet `nodemon` has some interesting features, but for the most basic apps, you might not need it anymore.

- We had to add mocks for many DOM APIs manually. Erick mentioned that, we could have also used the `JSDOM` library.

  - While I agree that `JSDOM` is a valid choice here, **I would argue that using `Playwright` or any tool that runs our code in a real browser would be even better**.

    - The _component testing_ exists in `Cypress` and in `Playwright` and the DX is pretty good!

- The error you get when trying to compare _objects_, but I think this will be true for any non-primitive type, is misleading.

  - It tells you that _"Values have same structure but are not reference-equal"_. The message is true, but it should point you to another method on `assert` for non-primitive types.

    - **Use `assert.deepStrictEqual` when comparing non-primitive types**.

- Our "web" tests are pretty brittle. They test the implementation details of the methods, and the mocking set up we have is pretty involved.

- **Node test runner has mocking capabilities**. It seems like it is on pair with Jest?

  - I really like the fact that the `it` or `test` function callback gets the `context` object with `mock` method.

    - You do not have to clear any mocks since they are _local_ to the test!

- **TIL that you can make a JavaScript file an executable**.

  - Add the `#!/usr/bin/env node` to the top of the file.

  - Add the `bin` entry to the `package.json`.

  - Run `npm link`.

    Now, you can run the app as if it was installed on your machine. Interesting!

- To build the graphical CLI, Erick decided to use `blessed` and `blessed-contrib` packages.

  - It appears that these packages are no longer maintained.

---

- **TIL that you can parse "incomplete" URLs via `parse` exposed from `node:url`**.

  - It seems like the `request.url` returns "partial" urls, like `/users`.

    - The `new URL` would not be a good fit here, as you would have to "fake" the base path – `new URL(request.url, 'http://something.com')`

    - The `parse(url)` returns the `URL` object which many properties set to `null`.

Finished Part 8 -42:33
