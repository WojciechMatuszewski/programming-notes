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

Finished part 3, 32:25;
