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

- _Composition over inheritance_ rule also applies to how we write CSS.

- **You should favour layout primitives over elaborate CSS selectors**.

  - Take a look at TailwindCSS. It uses primitives for all the styles and it is really great to work with!

- If the styles compose well, you will not need to depend on media-queries (in almost all cases) to make the view look good on different screen sizes.

### Units

- Nowadays screen employ _sub-pixel_ rendering, so the `width: 1px` might not actually be 1 pixel.

  - Because of that, you **should consider using alternative units like `em`, `rem` and `ch`**. All of them are relative to the user's default font size.

- The combination of `calc` and _viewport units_ allows you to skip the `@media` queries all together!

- The `rem` is relative to the user's default font size. The `em` unit is relative to the immediate parent.

  ```css
    h2 {
      font-size: 2.5rem;
    }

    h2 strong {
      <!-- Relative to the 2.5rem -->
      font-size: 1.125em;
    }
  ```

- The `ch` unit pertains to the _approximate_ width of one character.

  - **There is also `ex` which is the _approximate_ height of one character**. TIL.

### Global and local styling

- CSS classes are portable. This portability promotes the "atom-based" programming where we compose classes.

- There are other ways to style elements. One could use `id` selectors, _inline styles_ or leverage the _Shadow DOM_.

  - The `id` selectors are high in specificity. This might be a problem when you want to override some styles.

  - The _inline styles_ are usually hard to maintain.

  - The `Shadow DOM`, while it works great for components, prevents you from taking the advantage of _global styles_. Most styles are "unable to get into" the `Shadow DOM`.

### Modular scale

- The notion of creating multiple scaled values to denote an increase in a property. Like `font-size` or `line-height`.

  ```css
    :root {
      --ratio: 1.5;
      --s-5: calc(var(--s-4) / var(--ratio));
      --s-4: calc(var(--s-3) / var(--ratio));
      --s-3: calc(var(--s-2) / var(--ratio));
      --s-2: calc(var(--s-1) / var(--ratio));
      --s-1: calc(var(--s0) / var(--ratio));
      --s0: 1rem;
      --s1: calc(var(--s0) * var(--ratio));
      --s2: calc(var(--s1) * var(--ratio));
      --s3: calc(var(--s2) * var(--ratio));
      --s4: calc(var(--s3) * var(--ratio));
      --s5: calc(var(--s4) * var(--ratio));
    }
  ```

### Axioms

- Instead of fixed width / height values, you **should deal in tolerances, like `max-inline-size`**. While this leaves you at the mercy of the layout algorithms, it will make your page fluid.

- **Consider using _global styles_ for axioms first. Then leverage the "exception-based styling" technique**.

  ```css
      <!-- The general rule -->
      * {
        max-inline-size: 60ch;
      }

      <!-- The exceptions -->
      html,
      body,
      div,
      header,
      nav,
      main,
      footer {
        max-inline-size: none;
      }
  ```

## Layouts

### The Stack

- The stack component is used to provide vertical spacing between boxes.

- The stack takes advantage of the **`owl selector` which looks like `.someClass > * + *`**.

  - You can [read more about the _owl selector_ here](https://alistapart.com/article/axiomatic-css-and-lobotomized-owls/).

  - Using the _owl_ selector (it looks like an OWLs face), ensures that the elements have consistent vertical spacing and that there are no margin "left overs". Some people call them "margin glue".

```css
  .stack > * + * {
    margin-block-end: 1.5rem;
  }

  /* Or you can use the recursive version. */

  .stack-recursive * + * {
    margin-block-end: 1.5rem;
  }
```

- Using the "recursive variant" ensures that you can interleave the _stacks_ with other elements and the spacing will be preserved.

### The Box

- The box as a container to **style any intrinsic properties, like `padding` or `border`**.

  - We have the _Stack_ for margins (and it should only be responsible for that).

- The book depicts the usage of the _owl_ selector for borders inside the `box`.
