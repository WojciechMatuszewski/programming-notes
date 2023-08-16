# Going through the "Every Layout" course

[You can find the course here](https://every-layout.dev/rudiments/boxes/).

## Fundamentals

### Boxes

- The **_box model_ is the formula upon which layout boxes are based, and comprises _content_, _padding_, _border_ and margin**.

  - Each element on the page has a "box" wrapped around it. Even if you cannot see it, the "box" is there.

- At the core, there is `display: block` or `display: inline`.

  - The `block` elements will span the entire width of the page (or the height, depending on the writing mode).

  - The `inline` elements will only take the space of their inner contents. This means that there might be a lot of `inline` elements after each other in a single "row".

    - The `block` elements are like paragraphs, and the `inline` elements are like words.

- There is also the `inline-block` which, in addition to behaving like `inline` element, allows you to set `margin` and `padding` on all it's dimensions (for `inline` elements one can only set the horizontal margin and padding)

- You should **prefer the _logical properties_ whenever possible**. They will work across different writing modes.

  - _Logical properties_ eschew terminology like "left" and "right" and instead use the _box model_ compatible `inline` or `block` properties.

- If you set a `width` on a `content-box` (default) element, the `width` will not take the `padding` or `border` into the account. If you are not cautious, you might case the element to overflow the parent!

  - That is why, we mostly use `border-box`. Nowadays almost all (if not all) CSS overrides include the `box-sizing:border-box` rule.

### Composition

<https://every-layout.dev/rudiments/composition/>
