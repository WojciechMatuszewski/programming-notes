# FEM Enterprise Accessibility

Notes from [this course](https://frontendmasters.com/workshops/enterprise-accessibility/).

## MVPs and Accessibility

- Ask yourself **what is the cheapest and fastest way we can make something usable and accessible**.

  - This does not mean using so-called "accessibility overlays". They do not work.

- Architecture plays a key role here.

  - We can make shortcuts here and there, **but we cannot compromise on architecture**.

    - As long as the architecture piece is solid, you will be able to address any accessibility/tech debt issues along the way.

- **Shift accessibility concerns left**. The earlier you voice your concerns, the better.

  - It is just like with bugs. The later in the development cycle we discover a given bug, the more costly it is to fix it!

## Accessible UIs

- It does not matter what tooling you use. **The basics still apply**.

  - Using valid HTML elements. Instead of using `div` tags everywhere, use `article`, `section` tags and others.

- **Do not create buttons with `div` elements**.

  - The `button` has a LOT of functionality. There is no excuse of trying to re-create it.

    - You are wasting cycles doing so.

- The the `sr-only` or similar classes.

  - Very useful for adding more semantic information for the screen reader, but where such information is not visible in designs (might be implied from the visual structure of the page).

  - **Be mindful that you will still be able to tab into focusable elements within the wrapper with `.sr-only`**. This might or might not be what you want (great for skip-links, not so good for other features).

  - Declarations like `display: none` or `hidden` attribute will hide the element from the accessibility tree.

- Consider using the `inert` attribute.

- **Bake accessibility testing into your CI**. Make it automatic.

  - There are tools like `Axe` which you could use as a browser extension or standalone module you invoke while testing.

## Accessible Naming & Screen Reader Concerns

- Tables, form controls, links and buttons all have names. If you use semantic HTML you get a lot for free.

  - Having good names is also crucial for testing. The more accessible your page is, the easier it will be to test!

  - You can also associate the name of the element with the content of other element, which could even be hidden via `display: none`.

    - To do this, you would use the `aria-labelledby` attribute.

- Elements can also have _descriptions_. Those are read with a delay and might be skipped by the screen reader.

  - **To provide a _description_ for a given element, you have the `title` and `aria-describedby` attributes at your disposal**.

- Use the information provided inside the **_Accessibility Tree_** to understand what title, description the element has.

  - The a11y tree lives alongside the DOM tree.

- If you are serious about accessibility and how screen reader works with your website **you need to test on Windows as well!**

  - You are most likely on MacOS, but some of your customers might not be.

- TIL that `role="application"` exists. [You can read about it here](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Roles/application_role).

  - It seem to completely disable any kind of focus-management that could be done by browser.

  - Allows you to customize all the interactions, but there is a danger that you will miss something and make the experience worse.

- **Consider wrapping images with `figure`**.

  - In addition to being more semantic, you can use `figcaption`. In addition, using the `figure` might positively impact SEO.

## Accessibility in JS Apps

- First, understand that it is okay to use JS.

  - In fact, some of the patterns recommended by many a11y-focused organizations require JS to implement.

  - In most cases, it boils down to having inaccessible HTML, not the fact that we website uses a given framework.

  - There is the [`navigation` API](https://developer.mozilla.org/en-US/docs/Web/API/Navigation_API) but it is still considered experimental. Many people hope that using this API will solve a11y problems with client-side routing.

- **Routing gets quite tricky with client-side apps**.

  - Make sure the router you are using thinks about a11y and implements features like focus switches when you navigate between pages.

  - To be honest, **the best solution would be to use MPA frameworks**.

- **Consider using `:focus-visible` instead of `:focus`**.

  - It will show the outline only when user is navigating via keyboard or similar device.

  - It WILL NOT show the outline whenever you click on stuff.

- A neat thing shown in the workshop is the `ReactAriaLiveAnnouncer` from `@react-aria`.

  - It will automatically inject a `div` with `aria-live` and `aria-relevant` attributes to the DOM and append the messages you invoke the module with.

    - Super handy.

## Test Automation for Accessibility

- **Realistically** you can **automate about 50% of a11y related testing**.

  - Things like focus order, alt text quality and so on is rather hard to test.

- There are multiple tools you could use to test accessibility.

  - There are linter plugins and plugins for your favorite testing library.

- Consider making the a11y tests fail CI.

## Organizational Skill-Building

- In the end, you cannot do it alone. **It is imperative to foster a culture where everyone cares about accessibility**. This is where you come in.

  - Speak about the downsides of having inaccessible website: lower margins, potential lawsuits and so on.

  - Make everyone responsible for it.

  - **Share your knowledge**.

- **Make the accessibility easy to test**.

- **Make accessibility in your _definition of done_**.

  - It is much harder to fix issues later on.

## Wrapping up

The workshop was more theoretical than practical. I personally would prefer more hands-on exercises, but one cannot overlook the _soft_ part of building a11y accountability within the organization.
