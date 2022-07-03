# FrontendMasters – HTML Semantics & CSS Selectors

Notes from [this workshop](https://frontendmasters.com/workshops/semantics-selectors/).

> Slides & website [here](https://semantics-selectors.css.education/).

## Why do HTML CSS matter?

- HTML is not responsible for the display of the page!

  - HTML is NOT used for page styling.

  - HTML comes with many hidden features for SEO and such.

- We use very few number elements that are available to us. This impacts SEO and accessibility.

## Lists and Selector Review

- The `box-sizing: border-box` makes the most sense but is not the default, most likely, due to legacy considerations. Imagine all those websites breaking if the `border-box` is the new default (as opposed to the `content-box`).

- **Some CSS properties**, like `font-family` or `color`, when applied to parent, **will propagate to children**.

- **To capture the hierarchical nature of CSS**, one might want to **read CSS selectors from right to left**.

  ```cs
  ul a {
  color: red
  }
  ```

  Means _"select all elements that are descendants of the ul element"_.

<!-- Finished part 1 -->
