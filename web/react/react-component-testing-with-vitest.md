# React Component Testing with Vitest

Learnings based on [this course](https://github.com/epicweb-dev/react-component-testing-with-vitest)

## Learnings

- Historically, we used libraries like JSDOM to _emulate_ browser DOM in Node.js.

  - This worked to some extent, but there are problems with that approach.

    - The main issue is _incompatibility_. JSDOM will never be a "real" DOM for Node.js. It is a library maintained by the open-source community. Many things behave differently than in a "real" DOM.

      - A good example is the `.text` method on the `File` instance. It exists in Node.js and DOM, but it's nowhere to be found in JSDOM. **You could write correct JavaScript code, but your test might still fail** due to this difference.

- Instead of using JSDOM and other layers related to it, we **want to run our tests in a real browser – just like your users would run your application**.

  - This has huge benefits. If you run your tests that way, you have much greater confidence that the code will also work for users!

    - In addition, you can run tests across multiple browsers.

- In the course, we use the Vite browser mode. This allows us to test on a real browser and leverage all the nice things about Vite.

  - Why would we use Vite browser mode instead of Playwright Component tests?

    - I'm a bit torn on this one. On the one hand, you could say that this allows you to power your unit and component/integration tests via the same framework. But then if you use Playwright Component tests, you power your end-to-end and component/integration tests via the same framework.

      - But, if you use Vitest for unit/integration tests for FE and BE already, using the Vitest browser mode for component tests might be the best choice.

- Vitest allows you to **create "test workspaces"** which can target different environments.

  - [Read this documentation page](https://vitest.dev/guide/workspace.html#defining-a-workspace) to learn more.

  - **When doing this, consider using different `tsconfig.json` files for each environment**.

    - This allows you to control what global types are available in which files. You probably do not want to allow others to use DOM APIs in the context of a Node.js specific test, right?

- When querying for elements, make sure to use **_accessible selectors_** whenever possible.

  - This will force you to make your application more accessible by applying correct labels to HTML inputs and using semantically-sound HTML tags.

- In the course, we provide the MSW worker as a fixture. I'm quite fond of doing it this way.

  - Having said that, I'm not a fan of mocking the default happy paths for each test. IMO it makes it harder to reason about what happens in each test.

- Sometimes, **we want to assert that certain side effects _did not_ happen**. For example, that a notification _did not appear_ after clicking a button.

  - How can we do that? The notification might appear after a delay, so writing `.not.toBeInTheDocument` will not work as expected - the assertion will instantly pass.

    - **In such situations, consider using [inverse assertions](https://www.epicweb.dev/inverse-assertions)**.

- Earlier, I wrote about Playwright Component Tests versus Vite Browser Mode.

  - **I see one clear advantage of using Vite Browser Mode for component tests: the ability to create new components in the test files**.

    - If you use Playwright, you cannot do that. Usually, this is not a necessary "feature" to have, but it comes in _very_ handy when creating wrapper elements for each render.

- I really need to get better at using the debugger. It is very powerful, but I am clueless about how to use it.

  - In the course, we have set up a VSCode _launch configuration_ that allows us to put breakpoints and observe the paused state of the test in the browser.

    - Looking back, I could have really used this functionality several times in the past months.

## Wrapping up

I enjoyed this workshop quite a bit. It was quick and to the point. We did not dive _that deep_ into the weeds, but the material provides you with just enough information to get started and be productive.

The debugging section was a wake-up call for me. I have almost zero experience using the debugger, and that has to change.

While the concept of _launch configuration_ is not new to me, I feel like I'm not using them effectively (or at all) in my day-to-day work.
