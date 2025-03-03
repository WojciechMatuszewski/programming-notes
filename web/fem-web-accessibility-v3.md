# Frontend Masters Web Accessibility V3

## Semantic HTML

- Some elements do not "do anything extra" in terms of functionality, but they "read" differently for screen readers.

  - Think `footer`, `header` , `aside` and so on.

- Some elements **provide a ton of built-in functionality**.

  - Think `button`, `textarea`, `dialog` and so on.

  - Make sure you associate `label` with various inputs.

    - You can use _implicit_ or _explicit_ labels.

    ```html
    <!-- Explicit label -->
    <label for="name">My input</label>

    <input name="name" id="name" type="text" />
    ```

    ```html
    <!-- Implicit label -->
    <label>
      My input
      <input name="name" id="name" type="text" />
    </label>
    ```

    - The `label` tag only works for certain elements. **Consider using `aria-label` for others, but keep in mind that labels provided with `aria-label` won't be translated using "Google Translate" extension**.

- It _is possible_ to re-create the functionality and behavior of _semantic HTML elements_ via ARIA attributes and custom JavaScript, but it is a ton of work.

- **By default, not every element is "tabbable"**. That is a good thing. Consider the amount of tags on a single website!

  - "Tabbable elements": `a`, `button`, `input`, `select`, `textarea`, `iframe`.

  - You can use `tabindex` to make a given element "tabbable".

    - Negative `tabindex` means "allow me to programmatically focus this element".

    - A zero `tabindex` means "this element is focusable via sequential keyboard navigation". The browser will honour the order of elements in the HTML.

    - A positive `tabindex` is the same as `tabindex=0` BUT **the browser will pick the focus order relative to the number you provide for `tabindex`**.

## Keyboard Shortcuts

- Great for accessibility and keyboard-only navigation. Nowadays, more and more websites have them built-in.

## Live regions

- **Very useful to announce updates to the website markup**.

  - Think notifications or changes in the stock price.

  - For some updates, you want to interrupt whatever the screen reader is reading and ensure that the change in the markup is read immediately.

    - For this situation, you would use `aria-live="assertive"`.

  - For others, you do not care that much. It is fine for the screen reader to "catch up" with the markup.

    - For this situation, you would use `aria-live="polite"`.

- **Note that some roles create implicit live regions**. [Here is the list on MDN](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/ARIA_Live_Regions#roles_with_implicit_live_region_attributes).
