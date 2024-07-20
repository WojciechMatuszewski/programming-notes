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

## DOM API

- The `document` is equivalent to the `html` tag.

  ```js
  window.document === document;
  ```

- Many of the APIs you could use, like the `getElementById` are powered by browser-created hash maps.

  - The browser creates many structures when reading the HTML. These structures are then used to power many of the APIs available to you.

  - The `getElementById` is much faster than more "loose" methods like `getElementByClassName` or `querySelector`.

  - **Browser caches your queries, so performing the same query multiple times will only be costly for the first time**.

    - Of course, it **depends wether the DOM was changed in-between the queries**.

- Interestingly, **some APIs return so-called _live collection_ of elements, and others return "just" a _collection_ of elements**.

  - The _live collection_ changes as the DOM changes.

- No matter which way of inserting elements to the DOM you choose, all of them are not that great for performance.

  - The **"worst offenders" are `innerHTML` and `insertAdjacentHTML`**.

  - Consider using `[createDocumentFragment](https://developer.mozilla.org/en-US/docs/Web/API/Document/createDocumentFragment)` to perform all your operations in memory before inserting elements into the DOM.

    - **The `DocumentFragment` does not support the `innerHTML`**.

      - Consider using the `[template](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template)` element if you need `innerHTML`.

- **You can have multiple selectors in the `querySelectorAll` function**.

  ```js
  document.querySelectorAll("first, second, third");
  ```

  That is nice to know!

Finished Part 2 -41:13
