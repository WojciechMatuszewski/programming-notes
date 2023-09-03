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

### The Center

- The author argues that we should never use `text-align: center` on paragraphs, but rather reserve this declaration for headings and short lines of text. I agree with this. Centered paragraph looks weird and is difficult to read.

- At it's core, the **_center_ component uses the `auto` value for the _inline_ margins**. This allows us to defer the margins calculations to the browser.

  - The **author suggests being specific here, as the `margin: 0 auto` might undo some other margin styles** which, in some cases, is not desireable.Consider using the `margin-inline: auto` here.

```css
.center {
  max-inline-size: 60ch;
  margin-inline: auto;
  /* This ensures that the max content inline size is 60ch and not 60ch - padding */
  box-sizing: content-box;
  /* Random padding value */
  padding-inline: 12px;
}
```

- In addition to the `margin-inline: auto` you might want to support **_intrinsic centering_. That is centering of elements based on their natural size**. Of course, this is optional stuff.

  - To achieve this, use the `flexbox` alongside with `align-items` and `flex-direction`.

```css
.center {
  max-inline-size: 60ch;
  margin-inline: auto;
  /* This ensures that the max content inline size is 60ch and not 60ch - padding */
  box-sizing: content-box;
  /* Random padding value */
  padding-inline: 12px;

  /* Intrinsic centering */
  display: flex;
  flex-direction: column;
  align-items: center;
}
```

- Notice that we do not use `px` or `rem` or `em` values for the `inline-size`. The `ch` are much better unit as they actually represent the text size.

### The Cluster

- _The Cluster_ component is for positioning elements in an _inline-like_ flow with consistent spacing between them.

  - For such layouts, I usually reach out for `display: flex` and `flex-wrap: wrap`. So did the author.

```css
.cluster {
  /* Keep in mind that you have the `justify-content` at your disposal! */
  display: flex;
  flex-wrap: wrap;
  gap: var(--space, 1rem);
}
```

- **There was a lot of hoops to go through prior to `gap` property**. Without it, when elements wrapped, we would most likely end-up with excessive margins on both sides.

  - To solve this issue, we had to apply negative margins on the container itself. It was not fun.

    ```css
      .cluster {
        --space: 1rem;
      }

      .cluster > *{
      display: flex;
      flex-wrap: wrap;
      /* ↓ multiply by -1 to negate the halved value */
      margin: calc(var(--space) / 2* -1);
      }

      .cluster > *>* {
        /*↓ half the value, because of the 'doubling up'*/
        margin: calc(var(--space) / 2);
      }
    ```

### The Sidebar

- By designing to _ideal element dimensions_ you can do away with `@media` breakpoints (the viewport-based ones).

  - Please note **that with introduction of `@container` queries, this is much easier to do**.

  - An example of setting the _ideal element dimensions_ is the `flex-basis` and `flex-wrap: wrap` rule.

- **Keep in mind that `flex-basis` is only a suggestion to the browser. Even if you set `flex-basis: 0` the box will NEVER be smaller than its content**.

  - The author uses this knowledge to define the "content" of the sidebar part as such.

    ```css
    .not-sidebar {
      flex-basis: 0;
      /* The `flex-grow` is denotes a proportion, not an actual value in px or rems */
      flex-grow: 999;
      min-inline-size: 50%
    }
    ```

  - Similar trick could be achieved with the `flex-shrink` ratio, where the two adjacent boxes shrink with different "speed".

Here is the full CSS snippet

```css
.with-sidebar {
  display: flex;
  flex-wrap: wrap;
  gap: var(--gutter, 1rem);
}

.sidebar {
  /* The `flex-basis` here is optional. We might want to inherit the width from children */
  flex-basis: 20rem;
  flex-grow: 1;
}

.not-sidebar {
  flex-basis: 0;
  /* When the wrapping happens, the element will fill all the width of the parent */
  flex-grow: 999;
  min-inline-size: 50%;
}
```

### The Switcher

The problem this component tries to solve is the following: when using `flex-wrap: wrap` and `flex-grow: 1`, the element that just wrapped, might span the whole parent. This **is, in most cases, undesirable as it could be perceived as "picked out" by the user**.

```text
┌──────────────┐ ┌───────────┐
│              │ │           │
└──────────────┘ └───────────┘

┌────────────────────────────┐
│ Element that sticks out    │
└────────────────────────────┘
```

To solve this, author presents the **so-called _Flexbox Holy Albatross_ technique** which leans heavily on `flex-basis` and a very clever usage of calc. [Here is an article on the topic](https://heydonworks.com/article/the-flexbox-holy-albatross-reincarnated/).

```css
.switcher {
  display: flex;
  flex-wrap: wrap;
}

.switcher > * {
  flex-grow: 1;
  /* 30rem is an arbitrary number */
  flex-basis: calc((30rem - 100%) * 999)
}
```

The **key to understanding this piece of code is to understand how the `flex-basis` works**.

- The `flex-basis` is a suggestion. The content might grow above the defined `flex-basis`, but it will NEVER have value less than the content of the box it applies to.

- **Negative `flex-basis` valued are ignored**.

Given the points above, we can deconstruct the `flex-basis: calc((30rem - 100%) * 999)`;

- If the container `inline-size` is greater than `30rem`, the `flex-basis` is a big negative number. As such the `flex-basis` is ignored. The `flex-grow` takes over. This makes all the elements grows to tak up an equal proportion of horizontal space.

- If the container `inline-size` is less than `30rem`, the `flex-basis` is a big positive number. As such the elements will wrap as the `flex-basis` is much bigger number than the container `inline-size`.

```
  -n * 999             n * 999

┌──┐ ┌──┐ ┌──┐        ┌─────┐
└──┘ └──┘ └──┘        └─────┘

                      ┌─────┐
                      └─────┘

                      ┌─────┐
                      └─────┘
```

- To support gutters use the `gap` property.

- To support different proportions of the children, use the `flex-grow: X`.

  ```css
  .switcher > :nth-child(2) {
    flex-grow: 2;
  }
  ```

#### Quantity query

What if you want to ensure that, if the `switcher` contains more than X children, you always present the "pancake" layout instead of having the elements sit in one row? **This is possible to achieve without any media queries**. Enter the world of _quantity queries_.

First, you need to be aware of the `nth-last-child` selector. This selector allows you to pick the "nth-child" **counting from the end**.

```text
     :nth-last-child(3)
┌────┐ ┌────┐ ┌────┐ ┌────┐
│    │ │ XX │ │    │ │    │
└────┘ └────┘ └────┘ └────┘
```

Then, you can combine this with the `n+X` syntax you can pass into the `:nth-last-child`. **Remember that `n` starts counting from 0, not 1** (this, of course makes sense), but the **CSS starts counting elements from 1**.

```text
    :nth-last-child(n+3)
┌────┐ ┌────┐ ┌────┐ ┌────┐
│ XX │ │ XX │ │    │ │    │
└────┘ └────┘ └────┘ └────┘
```

Then, you can apply the `~` selector to select all elements that are preceded by a given element.

```text

   :nth-last-child(n+3),
   :nth-last-child(n+3) ~ *
┌────┐ ┌────┐ ┌────┐ ┌────┐
│ XX │ │ XX │ │ XX │ │ XX │
└────┘ └────┘ └────┘ └────┘
```

This means that **you can effectively apply styles if the container has equal or more than X children**. Note that, if there were 2 boxes, the selector would not apply any styles as there is no "3rd child counting from the end".

```text
 :nth-last-child(n+3),
 :nth-last-child(n+3) ~ *
┌──────────┐  ┌──────────┐
│          │  │          │
│          │  │          │
└──────────┘  └──────────┘
```

You can **reverse the count – make the query apply when there are equal or fewer than X children** as well. This is done with the combination of `:nth-last-child` and `:first-child`. **If the `:nth-last-child` is also `:first-child` that means that the list has less than or equal to elements**.

```
           :nth-last-child(-n + 3):first-child
     ┌─────────┐   ┌─────────┐   ┌─────────┐
     │         │   │         │   │         │
     │   XX    │   │         │   │         │
     └─────────┘   └─────────┘   └─────────┘


         :nth-last-child(-n + 3):first-child
┌──────────┐  ┌────────┐ ┌────────┐ ┌───────┐
│          │  │        │ │        │ │       │
│          │  │        │ │        │ │       │
└──────────┘  └────────┘ └────────┘ └───────┘
```

To style all the "rest" elements, use the `~` selector. Like in the case of "more than or equal to" query.

### The Cover

- There are many ways to _vertically_ center the content in the CSS.

- Before the introduction of Flexbox it was quite hard.

  - One could use the `position:relative` and `position:absolute` / `transform: translateY`, but it did not guarantee that the content would not overflow.

- A go-to solution for _vertically_ centering stuff, is the `justify-content: center` in flexbox.

  - Depending on the amount of items, you might also want to look into `justify-content: space-between`.

### The Grid

- We should **create CSS based on the content, not the width**.

  - In some cases, we know the number of columns up-front, but those cases are rare.

- To achieve a fluid grid layout, we will need to use the `repeat`, `auto-fit` and `minmax` keywords when defining the columns

  ```css
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  }
  ```

  The problem here is the hardcoded value in `minmax`. With `flex-basis` we only told the browser the _ideal_ width, but here it is a hard boundary. **It might cause overflow!**

- To **"fix" the issue with hardcoded value inside `minmax`, use the `min` CSS function**.

  ```css
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(min(250px, 100%), 1fr));
  }
  ```

  Here, **there is no hardcoded value as the `min` will always yield `100%` if the column is less than `250px`**.

### The frame

- `img` along with `iframe`, `video` and `embed` tags are so-called _replaced_ elements.

  - The _replaced_ elements are elements **whos contents are NOT affected by the current document's styles**. If you think about it, they are pretty unique.

- While you could use `background-image` to embed images on a site, **it is preferable to use the `img` tag**. Due to specific user settings, the `background-image` could be removed.

- **The `display: flex` does not affect _replaced_ elements**. This means that our `frame` element could target both _replaced_ and _non-replaced_ elements.

  ```css
  .frame {
    aspect-ratio: 16 / 9;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
  }

  .frame > :is(img, video) {
    inline-size: 100%;
    block-size: 100%;
    object-fit: cover;
  }
  ```

  **Take a note of the `aspect-ratio` property**. Before it was widely supported, we had to rely on the `padding` to do the work for us (with a very specific percentage-based values).

### The Reel

- At its most basic level, the carousel could be implemented as scrollable container.

  - You can enhance it with many modern features, [like the `scroll-snap`](https://developer.mozilla.org/en-US/docs/Web/CSS/scroll-snap-type).

- **You can implement the "there is more to scroll" visual features using `background-attachment` property**.

  - Read [this blog for more information](https://lea.verou.me/blog/2012/04/background-attachment-local). The technique is pretty fascinating.

### The Imposter

- The `transform: translate` property **takes the element dimensions into the account**.

  - This means that, the `transform: translate(50%, 50%)` will transform the position of the element in respect to the element dimensions.

  ```css
    .imposter {
      position: absolute;
      inset-block-start: 50%;
      inset-inline-start: 50;
      transform: translate(-50%, -50%);
    }
  ```

  Of course, this is only part of the story, as we also have to worry about the overflow. To ensure no overflow happens, once could set the maximum dimensions for width and height.

  ```css
  .imposter {
    <!-- stuff from earlier -->
    max-inline-size: 100%;
    max-block-size: 100%;
  }
  ```

- The **`position: absolute` is, in this case, used to center the element based on the document or positioning container**. We should swap the `position: absolute` to `position: fixed` to center things in relation to the _viewport_.

### The Icon

- Use `svg` for icons. Stay away from font-based icons. [They are not good](https://cloudfour.com/thinks/seriously-dont-use-icon-fonts/).

  - You have to do a LOT of work to have a good font-based icons system in place. That is not the case for SVGs.

- `SVGs` are `inline` elements by default. **To change their vertical alignment one could use `flexbox` or `vertical-align: middle/baseline` properties**.

  - **The `vertical-align: middle` result is not necessary what you would expect**. You will most likely need to hand-roll a `px/rem` value here.

- There is a lot of gotchas with matching the icon to the font height. You will need to tweak the `height` / `width` values according to whether the copy after the font starts with an uppercase or lowercase letters.

```css
.icon {
  /* Have the same height as the font */
  height: 0.75em;
  height: 1cap;

  /* Have the same width as the font */
  width: 0.75em;
  width: 1cap;
}

.with-icon {
  /* Eliminates any "magic" space created by inline elements */
  display: inline-flex;
  align-items: baseline;
  gap: 0.5em;
}
```

- When using **icons without any text, consider using `aria-label` attribute to ensure they are accessible!**

  - Quite an important detail which will help you with testing as well (assuming you use a library with `getByRole` or `getByLabelText` selectors).

- The author also **recommends setting the `role="img` to the `svg` when it is used without any text**.

  - Seems nice and makes sense – the screen reader will announce it as labelled image.
