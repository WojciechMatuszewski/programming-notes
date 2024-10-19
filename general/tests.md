# About tests

## Writing good tests

> Based on [this blog post](https://mtlynch.io/good-developers-bad-tests/).

- Test code should minimize the number of layers of abstraction.

  - Keep the readers in your test function. **DO NOT make me scroll up and down to understand the logic**.

    - This means moving _all_ the data test needs into the test function.

    - You might also want to consider moving _all_ the "setup" code there as well.

      - **If this is painful, your test is doing too much!**

- **Favour simplicity over DRY**. Redundancy is okay.

  - If there are too many redundant parts, your code has structural issues.

    - Ask yourself: "why is my system so hard to test?"

## Inverse assertion

> Based on [this great blog post](https://www.epicweb.dev/inverse-assertions)

Have you ever found yourself in a situation where you wanted to assert that something _did not happen_?

Imagine a scenario where you want to assert that notification _did not_ fire after user clicks a button.

```jsx
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { Dashboard } from "~/components/dashboard.jsx";

test("does not display notification when saving a post", async () => {
  render(<Dashboard />);
  await userEvent.click(saveButton);
  expect(notification).not.toBeInTheDocument();
});
```

**This test will not work as expected**. The `not.toBeInTheDocument` resolves instantly and **will give you a false-positive when notification fires only _after_ certain time**.

You might be tempted to add a `sleep`-like function to the test and call it a day.

```jsx
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { Dashboard } from "~/components/dashboard.jsx";

test("does not display notification when saving a post", async () => {
  render(<Dashboard />);

  await userEvent.click(saveButton);

  await sleep(1_000);

  expect(notification).not.toBeInTheDocument();
});
```

But **adding the `sleep` function makes this test brittle**. Any random timeout, will cause the test to eventually be unstable and break.

Another solution might be to mock the timers and then attempt to assert the condition.

```jsx
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { Dashboard } from "~/components/dashboard.jsx";

test("does not display notification when saving a post", async () => {
  mockTimers();

  render(<Dashboard />);

  await userEvent.click(saveButton);

  timers.advance(1_000);

  expect(notification).not.toBeInTheDocument();
});
```

This is better, **but we still rely on random timeout, now hidden behind the `advance` API**.

**The ultimate solution is to "reverse" the condition**.

```jsx
import { render } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { Dashboard } from "~/components/dashboard.jsx";

test("does not display notification when saving a post", async () => {
  mockTimers();

  render(<Dashboard />);

  await userEvent.click(saveButton);

  const notificationVisiblePromise = waitFor(() => {
    expect(notification).toBeVisible();
  });

  await expect(notificationVisiblePromise).rejects.toThrow();
});
```

The `waitFor` will re-evaluate the state of the UI every so often. If it fails, it means that we could not find the notification in a given time-frame. **Exactly what we want**.

Note that you can always change the settings of `waitFor` (or any other similar API you might be using).
