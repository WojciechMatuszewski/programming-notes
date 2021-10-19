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
  - The `baseline` is pretty magical. You can make the elements aligned on a shared `baseline` even though they do not share the same container.

### Growing and shrinking

- The `width` and `heigh` properties in the context of `flexbox` are more like suggestions. Items can shrink to their minimum `width`/`height` depending on the parent.

  - **In the `flexbox` world there is the `flex-basis` property. This property acts as the `height`/`width` but overrides the `height`/`width` properties**.

  - `flex-basis` is `flex-direction` agnostic. The `width` and `height` is not.

- Think of the `flex-shrink` property as: "How quickly should the element shrink if there is no enough space for its normal width/height".
  - `flex-shrink: 10` means that a given element will be _shrinking_ ten times faster than other elements.
  - Have you ever been in a situation where a circle turns into an oval shape after the browser window gets small enough? Yup, we have all been there.
    Instead of fiddling with CSS properties, set the `flex-shrink:0` on the element that you do not wan to get squished. Problem solved!

### The "flex" shorthand

- The `flex:1`. I've used it but I never truly understood it until now. The shorthand sets three properties.
  - The `flex-grow:1`. This means "take all available space".
  - The `flex-shrink:1`. This means "shrink at the rate of 1".
  - The `flex-basis:0%`. This means "your width is 0%". As I eluded earlier, this property overrides `width`.

### Constraints

- Neat trick Josh is touching on is the usage of `flex-shrink` with a very high value. This will produce an effect where one element seem to be shrinking and the other is not.

### Pro tip

- In the context of `flex` **use `flex-basis` instead of `width`/`height`. Remember that the `flex-basis` overrides the `width`/`height` even if it comes BEFORE the `width`/`height` property**.

## Wrapping

- Use the `flex-wrap` property to control the wrapping of the `flex` children. Wrapping will occur if the element cannot shrink anymore.

- The `justify-content` applies the transformation **on the content as a whole** while the `align-items` applies the transformation **on individual item**.
  This difference is subtle but noticeable on some occasions.

## Grouping and gaps

- Use the `gap` property instead of grouping elements. There is no need to artificially pollute the markup so that you can apply margins.
  The `gap` property is pretty neat. The **space is applied only between the children**.

## Ordering

- Be mindful about the `flex-direction` property. It might seem like using this property could save you some logic regarding reversing elements.
  This is not the case. **The `flex-direction` is purely visual. It does not affect the DOM layout**.

- `flex-direction` can be very helpful in some situations. **One might use `-reverse` to control focus order without using `tabindex`**.
  Utilizing the DOM placement and changing the visual ordering is much better than having artificial `tabindex` properties.

## Flexbox Interactions

- In the context of CSS, there are multiple _layout modes_. Given _layout mode_ **can, but does not have to, interact with other _layout mode_**.
  For example, the `position:absolute` layout mode is not really compatible with `flex` layout mode.

- Sometimes **mixing _layout modes_ can be very helpful**. You can do a lot with `position:sticky` declared on a `flex` child.
