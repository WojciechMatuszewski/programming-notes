# CSS for JS devs

## Fundamentals

- I was not sure about the difference between _pseudo-elements_ and _pseudo-classes_. The difference is that _pseudo-elements_ target "sub-elements" within a given element. On the flipside, the _pseudo-classes_ target the elements' specific state.
  The syntax is also a bit different. For the _pseudo-elements_ we use `::`, for the _pseudo-classes_ we use `:`

- The difference between `rem` and `em` values. Here is the difference:
  - The `rem` value is **relative to HTML tag font size**.
  - The `em` value is **relative to the current tag font size**.

## Rendering logic 1

- **Some** CSS properties are inherited from the parent. A good example would be the `color` property. Please note that **the notion of property inheritance only applies to some properties**.

- The `inherit` property allows you to **manually force inheritance for a given property**.

- Josh presents a neat JavaScript analogue to how the CSS specificity works

  ```js
  const appliedStyles = {
    ...inheritedStyles,
    ...tagStyles,
    ...classStyles,
    ...idStyles,
    ...inlineStyles,
    ...importantStyles,
  };
  ```

  Later he argues that one does not even need to be aware of the _specificity cascade_. I second his opinion. With modern tooling and auto-generated selectors, it is seldom an issue.

- I've been using `display:block` and `display:inline` without putting any though to it so far. Josh cements my knowledge about those.
  - `block` property means _stacked on top of each other_.
  - `inline` property means _in-line_. Just like people standing in a queue.

### The Box Model

Not understanding the _box model_ is like not understanding what a _closure_ is.
You might get away with it, but you might get stuck once you encounter a problem that requires its understanding.

- The `box-sizing` property changes how the overall `width` and `height` of the element is calculated.

  - The `content-box` property tells the CSS engine to **add** the padding/border calculations to the width/height of the element.
  - The `border-box` property tells the CSS engine to **subtract** the padding/border calculations to the width/height of the element.

  The rule of thumb is to always have `border-box` on each element. It makes the calculations so much simpler as this is how we usually think of width and height.

- The `padding` property is like the stuffing that your jacked comes with.

  - I always forget what number corresponds to which side in the `padding: 0 0 0 0` rule. Think of a clock: `top right bottom left`.

- In the `border` section I was surprised to learn that creating element with _double border_ does not require 2 HTML tags. There is the `border: double` property.

- Think of `margin` property like **your personal space**. During the pandemic health authorities required us to stay 2 meters apart from each other.

  - Margin can have a negative value.

### Flow Layout

- You **cannot change** the **width** and the **heigh** of an an `inline` positioned element.

  - When thinking about `inline` elements, think about "going with the flow".

- The `block` positioned elements are greedy. They fill all available space unless told not to do so.

- The `inline` positioned elements **can wrap**. Neat!

- The `inline-block` positioning is interesting. For the **parent of the element, the element is considered to be `inline`**, but for the **element itself can be styled like a `block` positioned element**.
  - This gives you the ability to **add `width` and `height` to the element**.
  - One notable gotcha is that **the `inline-block` element does not wrap**.

#### Width

- TIL that you can specify `max-content`, `min-content` and `fit-content` as `width` properties. I though that these are available only in the context of CSS grid.

## Flexbox

- `flexbox` is all about thinking "how am I controlling the layout on THIS(single) axis.

- The `align-items: baseline` is pretty interesting. It works similarly to the `flex-end` but in the relation to typography.
  - Could be used in the context of inputs where the `label` text is huge and you want the input to be positioned at `baseline` of that.
