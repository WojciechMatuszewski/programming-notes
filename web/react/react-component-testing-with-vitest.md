# React Component testing with Vitest

Learnings based on [this course](https://github.com/epicweb-dev/react-component-testing-with-vitest)

## Learnings

- Historically, we used libraries like JSDOM to _emulate_ browser DOM in Node.js.

  - This worked to some extent, but there are problems with that approach.

    - The main issue is _incompatibility_. JSDOM will never be a "real" DOM for Node.js. It is a library maintained by the open-source community. A lot of things behave differently than in a "real" DOM.

      - A good example is the `.text` method on the `File` instance. It is there in Node.js, it is there in DOM, but it's nowhere to be found in JSDOM. **You could write correct JavaScript code, but your test might still fail** due to this difference.

- Instead of using JSDOM and other layers related to it, we **want to run our tests in a real browser – just like your users would run your application**.

  - This has huge benefits. If you run your tests that way, you have much greater confidence that the code will also work for users!

    - In addition, you can run tests across multiple browsers.

- In the course, we use the Vite browser mode. This allows us to test on a real browser and leverage all the nice things about vite.

  - Why would use use Vite browser mode instead of Playwright Component tests?

    - I'm a bit torn on this one. On the one hand, you could say that this allows you to power your unit and component/integration tests via the same framework. But then if you use Playwright Component tests, you power your end-to-end and component/integration tests via the same framework.

      - But, if you use vitest for unit/integration tests for FE and BE already, using the Vitest browser mode for component tests might be the best choice.

- Vitest allows you to **create "test workspaces"** which can target different environments.

  - [Read this documentation page](https://vitest.dev/guide/workspace.html#defining-a-workspace) to learn more.

  - **When doing this consider using different `tsconfig.json` files for each environment**.

    - This allows you to control what global types are available in which files. You probably do not want to allow others to use DOM APIs in the context of a Node.js specific test, right?

Start 18: Best practices
