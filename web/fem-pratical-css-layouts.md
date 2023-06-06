# FEM Practical CSS Layouts

Watching [this workshop](https://frontendmasters.com/workshops/practical-css/). [Website for the course](https://practical.css.education/).

## Establishing base styling

- Consider **defining global CSS variables on the `:root` level rather than the `html` level**.

  - The `:root` has a bigger specificity than the `html` element selector.

  - The **`:root` is the "true" root of the document**. Keep in mind that the CSS can be used in other formats, where there might not be the `html` element – **for example `svg`**.

- The **`rem` is relative to the "root"** while **`em` is relative to the font size of element**.

  - Since `em` is relative to the font size of a given element, it will not be consistent between different tags, like `h1` and `span`.

  - Using `rem` makes the spacing consistent as it is the "root" that is the single source of truth.

- For setting typography scale, you might find [typescale](https://typescale.com/) useful.

- The **order of pseudo-selectors, at least for the `a` tag, matters!**.

  - This is new to me. I've always though that it does not matter in which order we put the pseudo-selectors. [That is NOT the case according to MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/:link).

## The "easy design"

- Jen recommends _mobile-first_ approach.

  - I second this opinion. It is much easier to add elements into the page than to remove them. Removing elements can lead to removing useful information from the page. If your design includes all the information the page needs on the smaller screens, there is no need to worry about missing elements.

- The `main` means the "main focus of the page".

  - The most important section should most likely contain `h1`.

- The `aside` acts as a "support" for the `main` content.

- Jen wraps the foundational styles with a `@layer`.

  - The `@layer` also known as **cascade layers** are a tool to **avoid specificity problems on selectors**.

    - If you **opt-in into using layers, it's the order in which layers are declared not the CSS selectors specificity that matters**.

      - This means you can have a selector with a very high specificity in one layer, but then the styles can be overwritten by a selector with a very low specificity inside a layer defined after the first one.

      - This helps as now the only think that matters is which order your declare the layers.

Finished day 1 part 2, 22:37 -> <https://codepen.io/WojciechMatuszewski/pen/yLQBrWV?editors=1100>
