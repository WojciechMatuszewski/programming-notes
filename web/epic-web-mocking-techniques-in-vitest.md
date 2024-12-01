# Epic Web Dev – Mocking Techniques in Vitest

## Establishing test boundaries with mocks

- **What you mock** influences how _much_ code you will cover with tests.

- You can think of mocking a module as setting a boundary.

  - The unmocked part is what you wish to test.

  - The mocked part is what you want to control to influence the test.

    > Nothing beyond this boundary matters for my test

- It is imperative to **establish the test boundary at the lowest possible level**.

  - If you establish the boundary "too high" in the module tree, you might end up testing against your mock, which has zero benefit and only introduces more code.

## Functions

- You have your regular **mocks**, that you can "inspect" to check what arguments they were called with and how many times they were called with.

  - **If you mock a function, you will "throw away" all of its implementation**.

    - This might be what you want! But in some cases, this is undesirable.

- Another option would be to use **spies**.

  - Those **do not override the implementation, but allow you to "inspect" the underlying function**.

- **Bottom line is that, while you CAN use mocks and spies**, you **should consider dependency injection first**.

  - DI is not only about _interfaces_ and _abstract classes_.

    - **Sometimes, all you need, is to pass the things you "import" as parameters**.

## Date and Time

### Dates

- **To make sure your test is robust, you have to "freeze" time**.

  - This, most likely, means mocking the "system time". While doing so, **ensure you pass the timezone information**.

    - If you do not, the test might fail for other people from your team living in different time-zones.

- Again, **perhaps passing the `now` date as a parameter** would be a better idea?

### Timers

- Here, instead of "freezing" the system date, you mock the timers.

  - Any testing framework has methods to advance timers by some time. Use that to help you with assertions.

### Ticks and Tasks

- In addition to controlling _time_, you can **control things related to the Node.js event loop**.

  Consider the following scenario (this code also works in the browser! [MDN documentation](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget))

  ```ts
  class Controller extends EventTarget {
    constructor() {
      super();
      queueMicrotask(() => {
        this.dispatchEvent(new Event("connection"));
      });
    }
  }
  ```

  We need to call the `dispatchEvent` within the `queueMicrotask`. Otherwise, the consumers would never get the `connection` event since it would run _before_ they had a chance to register a listener.

## Globals

- Consider using `globalThis` when spying on globals.

  ```ts
  vi.spyOn(globalThis.console, "log");
  ```

  The `globalThis` will work in any environment!

- For global variables, we can use `vi.stubGlobals`

- For environment variables, we cna use `vi.stubEnv`

- I **really like the DX of `vitest`**. Those functions to mock globals are great.

  - Having said that, I fear that their convince promotes bad practices that, instead of passing those as parameters, you use them globally throughout the application.
