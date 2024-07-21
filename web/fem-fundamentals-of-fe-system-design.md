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

## Observer API

- There are **three different kinds of observers**.

  - The **_intersection_ observer**.

    - This one tells you when a given element enters a viewport.

    - **It is much more performant than any kind of vanilla code you might write that does similar thing**.

      - In addition, it can track multiple items at once!

  - The **_mutation_ observer**.

    - This one tells you about the changes to the DOM structure.

      - You can granularly pick what you want to track: `childList`, `attributes` and others.

  - The **_resize_ observer**.

    - This one tells you about changes to the viewport size.

    - There are other ways to _react_ to viewport changes.

      - Via CSS Media Query.

      - Via CSS Container Query.

      - Via "resize" event. This one is particularly inefficient.

      The **_resize observer_ is quite nice since it gives you a callback in JavaScript you can use**. The bottom line here is to avoid using the "resize" event and consider other APIs.

- **While going through the _mutation_ observer API exercise, we used a couple of interesting web APIs**.

  - The `contenteditable` property on the section displaying the title.

  - To put the cursor _after_ the newly created heading, I've used the `range` and `selection` APIs

    ```js
    const heading = getHeading(target);
    target.replaceWith(heading);

    const range = document.createRange();
    const selection = window.getSelection();

    range.setStartAfter(heading);
    range.collapse();

    selection.removeAllRanges();
    selection.addRange(range);
    ```

## Virtualization

- Implementation of the virtualization is quite hard.

  - The basic premise is that you _swap_ elements and re-use the elements that are no longer in the viewport to show the new data.

  - In addition, you **have to move the observers as well**!

    - Since the elements are `position: absolute` the observers won't be "pushed" down when you update transforms.

- During the implementation, we used `.at` quite heavily. Neat to see this API used more and more.

## Application State Design

- **Depending on how large your data is**, you might want to look into **_data normalization_**.

  - This **usually means flattening your data structures to an object**.

    - This allows you to look things up in a O(1) time.

    - Of course, **it all depends on the access patterns you have**.

- There are many APIs for storing data in the browser.

  - The `localStorage`. Has Synchronous API. Not ideal for large amounts of data.

  - The `sessionStorage`. Only persists during the session.

  - The `Indexed DB`. Quite complex, but supports almost everything.

## Network Connectivity

- Two core network protocols to be aware of:

  - The **`UDP` protocol which is _lossy_**.

    - It does not need the "I got the data confirmation".

    - Used for streaming and anywhere where potential data loss is acceptable.

  - The **`TCP` protocol which is _lossless_**.

    - It involves so-called _handshake_. This process takes some time.

    - Ensures the server received data.

- There are many ways to receive data over time.

  - You have the **_long pooling_** which is very easy to implement, but battery & latency inefficient.

  - You have the **_server sent events_** which are much more efficient than _long pooling_, but they do not supporting sending data to the server.

  - You have the **_websockets_** which are quite complex in the implementation, but are very good from efficiency and functionality perspective.

## Performance Optimization

- Optimizations does not necessarily only happen in the browser.

  - **For example, upgrading your server to use HTTP 2 or 3 could be a huge win**.

    - This is related to the amount of open connections each protocol supports.

- **Make sure you are compiling to somewhat latest ES version**.

  - Of course, it all depends on the environment your users use, but in most cases, ES2020 target (or above) should be okay.

- You can pre-fetch assets via `rel="prefetch"` directive on a given link.

  - There is also `rel="prefetch"` directive.

- As for images, it boils down to using the latest format you can in a given situation.

  - The newer image formats are usually much smaller while providing the same visual fidelity.

  - **Consider compressing your SVGs**. If you use a bundler, the bundler is probably already doing that for you.

## Summary

A great workshop that also included a system-design mock interview at the end. I wish we did a bit more coding, but then this workshop would be quite long.

It was eye-opening to see how much complexity goes into creating a virtual scrolling library. In addition, we have used so many DOM APIs thought the workshop which was a great refresher.
