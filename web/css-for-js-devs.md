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

### Wrapping

- Use the `flex-wrap` property to control the wrapping of the `flex` children. Wrapping will occur if the element cannot shrink anymore.

- The `justify-content` applies the transformation **on the content as a whole** while the `align-items` applies the transformation **on individual item**.
  This difference is subtle but noticeable on some occasions.

### Grouping and gaps

- Use the `gap` property instead of grouping elements. There is no need to artificially pollute the markup so that you can apply margins.
  The `gap` property is pretty neat. The **space is applied only between the children**.

### Ordering

- Be mindful about the `flex-direction` property. It might seem like using this property could save you some logic regarding reversing elements.
  This is not the case. **The `flex-direction` is purely visual. It does not affect the DOM layout**.

- `flex-direction` can be very helpful in some situations. **One might use `-reverse` to control focus order without using `tabindex`**.
  Utilizing the DOM placement and changing the visual ordering is much better than having artificial `tabindex` properties.

### Flexbox Interactions

- In the context of CSS, there are multiple _layout modes_. Given _layout mode_ **can, but does not have to, interact with other _layout mode_**.
  For example, the `position:absolute` layout mode is not really compatible with `flex` layout mode.

- Sometimes **mixing _layout modes_ can be very helpful**. You can do a lot with `position:sticky` declared on a `flex` child.

## CSS Grid

- In contrast to `flex` layout, the `grid` layout takes both axis into consideration.

- There are some restrictions on what kind of layouts could be created with CSS Grid, but if your layout is a _grid_ you should be good.

### Layout mode

- **If you specify `display:grid` the children of that element will be rendered using _Grid_ layout**.
  It works exactly the same as `display:flex` is.

### Grid construction

- Unlike in _flexbox_ the **sizes** for the columns and rows you specify **are not "suggestions"**.
  In the _flexbox_ land we have talked about `flex-basis` and its relation to `width` and `height` properties.
  We have also touched on the fact that these properties are a mere suggestions for the CSS engine. This is not the case here.

- Even if you explicitly set number of rows and/or columns the browser might add additional rows and/or columns depending on the number of items within the parent.
  This is where the **notion of _implicit_ and _explicit_ individual columns or rows** comes in.

### Alignment

- You can use the `justify-content` just like in _flexbox_ but the behavior might differ.

- There is additional `justify-items` property that **works on the individual item level instead of the content as a whole**.
  One might draw comparison between `justify-items` and `align-items` property from the _flexbox_ world.

- When to use `xx-content` and `xx-items` properties?
  - Use the `xx-content` properties when you want to **affect the grid structure as a whole**.
  - Use the `xx-items` properties when you want to **affect the grid items**.

### Grid Areas

- The `grid-template-areas` property is super powerful. It allows you to explicitly create "slots" for each element.
  Imagine drawing borders between countries. Neat stuff!

### Tracks and Lines

- TIL that you could have **negative values for grid columns and rows**. The `grid-row: 1/-1` is a complete valid declaration.
  The negative values for columns and rows are very useful as **they always point to the last column / row**.
  The analogy I see here is getting the last element from the array.

  ```js
  const arr = [1, 2, 3];

  // First way of getting the last element
  const lastElement_1 = arr[arr.length - 1];

  // Second way of getting the last element. Just like negative grid column / row values
  const lastElement_2 = arr.slice(-1)[0];
  ```

### Fluid Grids

- The `minmax` function is very powerful. It allows you to specify `max-width` and `min-width` at the same time.
  The benefit over the `max-width` and `min-width` is that **it composes with other functions**. One such function might be `repeat` CSS Grid function.

- The `minmax` can be used in a _fluid context_. The `minmax(min(400px, 100%), 1fr)` is an example of such behavior.
  Instead of using media queries to scale the `min` value of `minmax` I'm using the `min` function.

### Recipes

- **Centering** works on similar basis to _Flexbox_.

  - You could use the `align-content` and `justify-content` properties or set them to `center`.
  - You could use **`place-content` shorthand**. This is a **shorthand for `align-content: center` and `justify-content:center`**.
  - Please note that **`-place-content` only works in CSS Grid**.

- **`sticky` positioning plays well with `grid`**.

  - Depending on the layout you are working on **you might want to create a _sticky wrapper_ within a `grid` element.**.
    This is because the `position:sticky` does not care about _CSS Grid lines_.

- _CSS Grid_ offers a **neat way of creating _full bleed layouts_**. A _full bleed layout_ is the one where an image spans the whole page while the text is centered and constrained by some `ch` unit limit. There are some steps to creating one.
  1. Have a parent declare flexible columns: `grid-template-columns: (1fr, min(30ch, 100%), 1fr)`.
  2. Have all children be placed inside the second column: `.wrapper > * {grid-column:2}`
  3. Have the _full bleed_ child take all columns: `.full-bleed {grid-column: 1 / -1}`. It uses the trick we have learned about earlier.
