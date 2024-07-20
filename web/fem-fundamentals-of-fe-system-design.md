# FEM Fundamentals of FE System Design

> Notes based on [this workshop](https://frontendmasters.com/workshops/systems-design/)

## Core Fundamentals

- Every website is built from boxes.

  - This is where the name _box model_ comes from.

- Every box has its own _anatomy_

  - The _context box_.

  - The _padding box_.

  - The _border box_.

  - The _margin box_.

- **Even if you do not wrap content within an HTML tag, the browser will create a "box" for that content**.

  - The content will be contained within an **_anonymous box_**.

- Some attributes or tags create **_formatting context_** in the browser.

  - The _formatting context_ dictates how the element will behave in a given "box".

  - You can also explicitly change the _formatting context_ via CSS. A good example would be setting `display: flex` on the element.

- By default, elements are rendered in so-called **_normal_ or _regular_ flow**.

  - This means elements are rendered right after each other.

  - You can change the flow in which element is rendered via CSS. For example using the `position: absolute`.

- The `position: absolute` and `position: relative` elements create a new **_rendering layer_**.

  - The GPU will NOT attempt to draw every _render layer_. There is a "selection process" where is promotes the elements to the **_graphic layer_**.

  - While you want some things to run on the GPU, you should not push everything into the GPU.

    - The more elements in the _graphic layer_ the more GPU memory-hungry your website is.

Finished part 1, part 2 next
