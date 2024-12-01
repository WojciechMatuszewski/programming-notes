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
