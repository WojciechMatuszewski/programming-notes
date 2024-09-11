# CSS for JS developers

## Module 0 - Fundamentals Recap

### Anatomy of a _Style Rule_

Let us dissect the following rule.

```css
p {
  margin: 32px;
}
```

- We call the **whole thing** a **rule**.
- We call the `p` a **the selector**.
- We call the `margin: 32px` a **declaration**.
- We call the `margin` a **property**.
- We call the `px` an **unit**.

### Media Queries

It just so happens that the most popular _media feature_ overlaps with a quite popular css _property_ (the `max-width` or `min-width`). You **cannot** use css properties with `@media` syntax.

```css
@media (max-width: 300px) {
}

/* valid */
@media (font-size: 32px) {
}

/* invalid */
```

### Selectors

#### Pseudo-classes

They offer you a way to **style the element based on its internal (browser) state**. There is a bunch of them for a given element (see MDN). The most popular ones are `:hover`, `:focus` and `:checked`.

#### Pseudo elements

Instead of targeting the state of a given element, **they target sub-elements of a given HTML element**. You have probably used them already. Rules such as `input::placeholder` or `input::before` are examples of rules for _pseudo-elements_.

We call them **_pseudo-elements_ because they target things in the DOM we did NOT explicitly define with HTML tags**.

#### Combinators

These lets you target elements in a hierarchical fashion. One could **target direct children via the `>` combinator**. Or maybe you fancy **to target all descendants of a given element by utilizing a whitespace between tags (`ul li`)**. There are tons and tons of other combinators which are extremely useful.

### Color

There are many ways to represent colors in CSS, some notable ones are:

1. Using the direct color name, like `color: red`.
2. Using the HEX codes, like `color: #FF0000`;
3. Using the RGB scheme, like `color: rgb(255,0, 0)`.
4. Using the HSL scheme, like `color: hsl(0deg 100% 50%)`. Notice that there are no commas in-between the values. Weird.

### Units

There are many units one can use. Here are some notable ones:

1. The `em` unit which **is relative to the parent font size**.

2. The `rem` unit which **is relative to the ROOT font size**. **By default**, the root HTML font size is `16px` (this can be changed in the settings).

You should not be setting the `html` tag font size explicitly. Doing so **will override the default font size of the user settings**. Instead, use percentages!

```css
html {
  font-size: 120%;
}
```

3. The `px` unit which **is NOT relative to anything**.

4. The `%` unit. This unit does not work for font sizes.

[This article](https://www.joshwcomeau.com/css/surprising-truth-about-pixels-and-accessibility/#introduction) tackles the question when should one use `rems` and when one should reach for `px` units. The **bottom line is to use `rem` when you want your element to scale with users default font size, otherwise consider using `px`**.

### Typography

Fonts are very important – they can make or break a given website. In web, we have so called _font families_. The name stems from the fact that each font has multiple variants we could use.

For the the font-related properties, the most common are `font-weight`, `font-size` and `line-height`.

1. **If your font family does not have the specified `font-weight`, the browser will try to "bold" the characters automatically**. Usually, such attempt ends with a weird-looking font. Before using given `font-weight` value, ensure that the font you are using supports it!

2. The `line-height` accepts an _unitless_ value (as well as a value with an unit). **You should prefer the _unitless_ variant** as it scales with user zoom settings.

## Rendering Logic 1

### Built-in Declarations and Inheritance

- Certain CSS properties inherit values from the parent HTML tag.

  - **Most of the properties that inherit are typography-related**. Think the `color`, `font-size`, `text-shadow` properties.

  - You can think of inheritance in CSS as similar to the prototypal-inheritance in JavaScript.

  - You can **force inheritance** by using the `inherit` property value.

### The cascade

- The name CSS expands to _Cascading Style Sheets_. The _cascading_ part means that, the order in which you declare the rules matter.

  - Think of the _cascade_ as analogy to merging different objects in JavaScript.

    ```js
    const result = {
      ...pStyles,
      ...aStyles,
    };
    ```

  - When determining the end-result, one has to also take into the account the **specificity of a given CSS selector**.
    While important when you write vanilla JS, if you use modern tooling, you do not really have to know the specificity of a given selector.

### Directions

- There is the **`block` direction (think lego blocks stacked on top of each other)**, and there is the **`inline` direction (think people standing in a line)**.

  - The above holds true for English and vast majority of other languages, but it's not exactly true for the arabic and some of the Asian languages.

  - You might want to consider **using the `margin-block-start`, `margin-block-end`, `margin-inline-start` and `margin-inline-end`** properties to style margins. These are universal and will adjust accordingly based on the direction of the document.

### The Box Model

- First, understand that **every element on the page is effectively a box**. You can think **of the term _box model_ as the implicit box that wraps every element**.

- The box model describes **how big a given element will be**.

  - The box model **describes how the content, `padding`, `border` and `margin`** interact with each other.

- By default, the browsers specify the **`box-sizing` to have a value of `content-box`**.

  - This means that, **the `width` and `height` of the child does not take into the account the padding and the margin of that element**.

  - Since the behavior described above might be confusing, one can use the `box-sizing: border-box` declaration.

    - The **`border-box` means that `width` and the `height` properties should account for padding, margin and border properties of a given element**.

#### Padding

- The **_inner space_ of a given element**.

  - You have your usual suspects of `padding-left` and such, but **you should also consider using the logical properties like `padding-block` or `padding-block-start` or `padding-inline-start` and so on**.

- For padding, **use `px` rather than other units**. You most likely do not want the padding to change alongside the text size.

- To **remember the shortcut notation, imagine a clock. Start from the `12:00` and go clock-wise**.

  - So the definition like `padding: 10px 20px` means `10px` for up/down and `20px` for right and left.

- TIL that **percentages in padding always refer to the width of the element**, not it's height. Makes sense if you think about it, but I did not know that.

#### Border

- Use **`currentColor` to "synchronize" the border color with the color of the text**.

  - The neat thing about the `currentColor` is that it can be used in any property which accepts a color.

- It is **common to not understand the difference between the `outline` and the `border`**.

  - The **`border` affects the layout, the `outline` does not!**.

- You should never disable the `outline` property. It breaks the navigation with the keyboard (there is no visual feedback when navigating).

#### Margin

- Think of the margin as the "personal space" of a given element. Margin does not account for the element width and height in the `border-box` model.

- The syntax of `margin` property is the same as the `padding` one. **Consider using the logical properties of `margin-block-start` and so on**.

- A single most important difference between the `margin` and the `padding` properties is the fact that, **with margin you can use negative values like `-32px`**.

  - This allows you to "escape" the parent box (think stretched images inside an article).

  - On the subject of stretched content, there are two possible solutions for the layout described above.

    1. Use the `calc` on the `img` tag.

       ```css
       img {
       	 margin: 0 -32px;
       	 <!-- using with 100% will NOT work as the width is calculated based on the parent -->
       	 width: calc(100% + 64px);
       }
       ```

    2. Wrap the `img` with a `div` and let the `div` do the "escaping".

       ```css
       img {
         width: 100%;
       }

       .img-container {
         margin: 0 -32px;
       }
       ```

### Flow layout

- There are 7 layout modes, the most notable ones are the _flow layout_, _flexbox_ and the _grid_.

  - **Most (the exception being for `img`, `video`, `canvas` and `button` tags) _inline_ elements ignore properties that could effect the layout of the page**.

    - Try using a `width: 300px` on a `span`. If you did not change the default browser styles, the `width` will be ignored!

- The _block_ elements do not spare space, they will fill all the available space.

  - There is a way to "contain" a block element with `width: fit-content` declaration.

- The **browser treats _inline_ elements as typography. This means that they have a "magic" bottom space" attached to them**.

  - The reason kind of makes sense for text – this ensures that the text is not crammed. But it does not make sense for the `img` tag.

  - In addition, for the **_inline_ elements, one can only specify the horizontal paddings and margins**.

- The **`inline-block` elements behave as _inline_ for the parent, but they accept _block_ element properties**.

  - This allows you to provide a margins and other properties to the element, but still have it flow _inline_.

  - Keep in mind that **`inline-block` elements do not wrap (apart from the text within the element)**. **The ability to wrap without any additional CSS is unique to `inline` elements**.

    - Try creating an `inline` and `inline-block` box with text. Add a border to that box and see how it wraps. The `inline` box will wrap "around" different sections of the text. [You can learn more about this behavior here](https://youtu.be/kj7WGnUDaI4?t=1001).

      - **This wrapping behavior explains why setting `height` on `inline` elements do not work**. If the element is wrapped across multiple lines, what would the `height` apply to?

#### Width Algorithms

- You have the _percentage-based_, _unit-based_ and _logical-based_ attributes at your disposal.

  - The **`fit-content`, `max-content` and `min-content`** are really useful. Keep in mind that `min-content` means the _smallest possible given the longest child_.

#### Height Algorithms

- As opposed to width, the default behavior for `height` on block elements is to be as small as possible, while also containing the whole content.

- Keep in mind that percentage-based unit is relative to the parent.

  - This plays a key role when setting the "page wrapper" height to `100%` – it might not work out of the box. Usually, this is caused by not setting the explicit height on the `html` and `body` (can also be percentage-based).

### Margin collapse

#### Rules of margin collapse

- **Only vertical margins collapse (in vertical-writing mode)**. I did not know that.

- **Margins only collapse in the _flow_ layout**. If you specify the `display: flex` on the parent, the margins will NOT collapse. TIL

- The bigger margin wins.

- Margins **must be "touching" in order for them to collapse (padding, border and other in-between elements will make elements NOT touch)**. This is very interesting. It means that, the following WILL NOT collapse

  ```html
  <p style="{{margin-bottom:" 20px}}>foo</p>
  <br />
  <p style="{{margin-top:" 20px}}>bar</p>
  ```

  But, the following WILL collapse

  ```html
  <div>
    <p style="{{margin-bottom:" 20px}}>foo</p>
  </div>
  <p style="{{margin-top:" 20px}}>bar</p>
  ```

  The biggest takeaway here is that **margin is used to create a space between sibling elements, even if it means "transferring" margin to the parent**.

- The margin of `0px` is going to collapse. An element with no `margin` property is considered to a `0px` margin.

#### Using Margin Effectively

A lot of people in JavaScript ecosystem advocate for completely ditching the `margin` property. This sounds reasonable and doable in the era of _component driven_ frameworks like React or Vue.

### Workshop

There is a lot of thing you can do with only margin, padding and some colors.

## Rendering logic II

- When **you set the `position` to `static` you tell the browser to NOT use the _positioned_ layout**.

  - You could achieve the same effect by using the `initial` value.

### Relative positioning

- Kicks in when you set the `position` property to `relative`.

- Gives you **the ability to use additional CSS properties**. These are the **`top`, `left`, `right` and `bottom`**.

  - They can take both positive and negative values (akin to the margin), but **using them DOES NOT impact the layout like setting a margin does**.

- You can **apply the relative positioning to both block and inline elements**. This makes it possible to move inline elements around. You could not do that with margin.

### Absolute positioning

- Kicks in when you set the `position` property to `absolute`.

  - **Using the `position: absolute` declaration WILL affect the layout**.

  - Using the `position: absolute` will **pull the element out of the flow layout**.

- Gives you **the ability to use additional CSS properties**. These are the **`top`, `left`, `right` and `bottom`**.

  - They can take both positive and negative values.

- The values for `top` and so on, are **calculated based on the containing block (more on this below), not the flow position as in the case of `relative` positioning**.

#### Centering Trick

- Remember how in the flow layout, the `margin: auto` only worked horizontally? It turns out, that, for **`absolute` positioned elements, the `margin: auto` works horizontally as well as vertically**.

  - This enables you to center the `absolute` element inside the containing block.

    ```css
    .box {
      width: 100px;
      height: 100px;
      margin: auto;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
    }
    ```

  - Keep in mind that, for this trick to work, **the element has to have a well-defined size**.

### Containing blocks

- The `absolute` positioned elements **are contained by elements with positioned layout OR the initial containing block – the viewport**.

  - In most cases, **people wrap the `absolute` positioned elements with `relative` positioned parents, but using `fixed`, `static` or even `absolute` positioned elements will also do**.

  - The `absolute` positioned elements **ignores all parents until it find either the one with positioned layout or the initial containing block**.

- The **containing block is also called `offset parent`**.

  - This makes more sense if you think about the names of the logical properties used to style position of the `absolute` element, like `offset-inline-start` and similar.

### Stacking Contexts

> [More information here](https://youtu.be/kj7WGnUDaI4?t=1991).

- There are many factors which influence how elements stack on top of each other. Most notable are the **layout mode and the DOM order** and the **`z-index`** value.

  - As a rule of thumb, the **_positioned_ elements will always render on top of non-positioned ones**.

- The famous **`z-index` property only works with _positioned_ elements and grid/flex children**. This means that, **the default _flow_ layout does not work with `z-index`**.

  - The **value of the `z-index` is compared relative to all the elements in a given stacking context**. A very important nuance to understand as this is the reason you sometimes find yourself bumping the `z-index` to a very high value to no avail.

    ```html
    <header zIndex="2" position="relative">My header</header>

    <main zIndex="1" position="relative">
      <div zIndex="999999">VeryHigh</div>
    </main>
    ```

  The `VeryHigh` **will NEVER overlay on top of the `My header`**. The `999999` is very high, but it only is comparable to other `zIndex` values inside the `main`.

- There are many ways one could go about creating the stacking context. The **most common way I've seen is to use the `position: relative` and `z-index: SOME_NUMBER` declarations**. You can find the full [list here](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Positioning/Understanding_z_index/The_stacking_context).

#### Managing z-index

- To avoid "`z-index` wars" you **should create stacking contexts when necessary**.

  - As noted previously, there are multiple ways one could do that, but **the best way to avoid the "`z-index` wars" is to not use the `z-index` at all!**.

  - To **create a stacking context without using the `z-index` property, use the `isolation: isolate` declaration**.

- Another very **useful thing you could do is to set the `isolation: isolate` on the `#root` element of your React application**. This guarantees that any modals (injected via portals) which render at the level of the `#root` element do not overlap with any elements from within the application itself.

- Note that **other CSS properties also implicitly create stacking context**. [View the list on MDN here](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_positioned_layout/Understanding_z-index/Stacking_context#description).

- **Think of the stacking context and `z-index` as versioning scheme**. Check this out

```html
<style>
  header,
  main,
  .tooltip {
    /* Needed for the layout to switch from the "flow" to "positioned". "flow" layout does not implement `z-index` */
    position: relative;
  }

  header {
    z-index: 2;
  }

  main {
    z-index: 1;
  }

  .tooltip {
    z-index: 9999;
  }
</style>

<!-- Version 2.0 -->
<header></header>
<!-- Version 1.0 -->
<main>
  <!-- Version 1.9999 It cannot possibly go higher than 2 -->
  <div class="tooltip">tooltip</div>
</main>
```

> Check out [this talk for more information](https://youtu.be/Xt1Cw4qM3Ec?t=1696).

### Fixed positioning

- The `position: fixed` declaration makes the element **behave as if `absolute` but it will also follow you when you scroll**.

  - As opposed to the `absolute` positioned elements which are bounded by their _positioned parents_, the **`fixed` elements are contained by the viewport**.

- You can **"break" the `fixed` position by applying transforms on the parent**.

### Overflow

- Make sure you configure your Mac to always show scroll bars. This way you will avoid a situation where a scroll bar is hidden for you, but visible for someone on a Windows machine.

- The `overflow: auto` is very smart, use it.

#### Horizontal Overflow

- For images, **which behave as `inline` elements by default** use the `white-space: nowrap` declaration.

  - Now, that I understand that `inline` elements are really treated like typography, I understand why `white-space: nowrap` works here.

#### Positioned Layout

- Keep the concept of _containing blocks_ in mind. If you grasp that concept, you will know, for example, why a given element peeks outside of its parent, even though the parent has `overflow: hidden` declaration! (it is because the parent is NOT using positioned layout).

- The **`fixed` positioned** elements will **always escape overflow**. This is because the `fixed` positioned elements are contained by the "initial containing block" and not by the parent with some kind of CSS property.

### Sticky Positioning

- The **most important** thing to keep in mind regarding `sticky` position is the fact that **`sticky` positioned elements CANNOT leave the bounds of their parent**.

  - If your parent is the height of the `sticky` positioned elements, no matter how hard you try, the element will not "stick" – it has no space!

- The `sticky` position "hijacks" the `top`, `left`, `right` and `bottom` properties to make them work differently than in the case of other `position: XX` declarations.

  - In `sticky` case these positions denote when a given element should "stick" (how much space from the corners of the PARENT ELEMENT!).

### Hiding content

- There are many ways to hide content using CSS, the most notable ones are `display: none`, `visibility: hidden` and `opacity` declarations.

  - The `display: none` will **make it so that the element is effectively removed from the DOM (even though they are physically still there)**. You cannot see it, you cannot click it, you cannot interact with it.

  - The `visibility: hidden` is less strict than the `display: none` counterpart. The **element will NOT be visible, but it will still take the space in the layout**, but **you will not be able to interact with that element, just like in the case of `display: none`**.

  - And finally, the `opacity`. Here, **you still can interact with elements**. You should not use this property for hiding content but rather for animations.

#### Accessibility considerations

- To **hide the content from the screen, but make it available to screen readers, use the `visuallyHidden` trick**.

  ```js
  const visuallyHidden = {
    position: "absolute",
    overflow: "hidden",
    clip: "rect(0 0 0 0)",
    height: "1px",
    width: "1px",
    margin: "-1px",
    padding: 0,
    border: 0,
  };
  ```

  - Using `aria-label` is okay as well.

- If you want to **hide the content from screen readers, but keep it visible on the screen** use the `aria-hidden` attribute.

## Modern Component Architecture

### Component Libraries

#### Breadcrumbs

- It appears that screen readers will read out the contents of the `&:before` and `&:after` pseudo-classes.

  - That is why, in the example, the `/` is created via the CSS shape, and not via the `content` property.

  - Very interesting.

#### Button

- Changing how the `outline` looks and feels is **very powerful mechanism for creating "multiple borders" effect**.

  - Check out the `outline-offset` property!

  - You can also leverage the `box-shadow` property to create three borders inside a single element.

#### Dynamic tags

- AKA polymorphic components, AKA the `as` prop.

- If you want to learn more about _polymorphic components_, check out my notes about them in the `react-typescript` file.

  - It is hard to create a good, fully typed, polymorphic `Box` component.

### Single Source of Styles

- The ability to refer to the components within the emotion styles rocks! Check this out.

  ```jsx
  const TextLink = styled.a`
  	color: blue;
  	text-decoration: none;
  `;

  const Figure = styled.figure`
  	${TextLink} {
  		color: black;
  		text-decoration: underline;
  	}
  ```

### Workshop

#### Progress bar

- A handy trick of using the `overflow: hidden` declaration to ensure border radius is consistent, even if the element overflows.

#### Select

- Josh used a very interesting technique – he hid the native select (using the `opacity` property) and then put a "presentational" bit on top of it.

  - All of this would not be necessary if we could style the caret of the `select`. Damn you browser styles!

  - You can **remove the caret with the `appearance` property**.

## Flexbox

- Another layout mode, **very relevant, even in the context of CSS Grid**.

- Deals with **a single axis, be it the Y or X axis**.

- You apply the `flexbox` value to the `display` property.

  - If you do so, **this will affect the children of a given element and not the element with the `display: flexbox` declaration**.

### Directions and Alignment

- You can set the **primary axis by using the `flex-direction` property**. By default, it has a value of `row`.

  - To control how elements align in relation to **primary axis**, use the `justify-content` property.

  - To control how elements align in relation to **secondary axis**, use the `align-items` property.

#### Alignment tricks

- Use the **`baseline` value** to align text in relation to the biggest text in a given row.

  - Do not be afraid of nesting `flexbox` declarations. For example, if you wish to use `baseline` and then center the elements, you have to nest the `flexbox` declarations.

- Use the `align-self` (secondary axis) to manipulate a given child.

  - Keep in mind that **`justify-self` (primary axis) DOES NOT exist**.

### Growing and shrinking

- In **the `flexbox` model, the `width` and the `height` properties are more like suggestions (by default)**. The element can shrink to a size which is less than the `width`.

  - Since, depending on the `flex-direction` you operate on different axis, **use `flex-basis` instead of `width` or `height` as it takes into the account the `flex-direction` property**.

    - Keep in mind that **the `flex-basis` is also a suggestion to the layout algorithm**. Nothing stops the item to shrink below the specified `flex-basis`.

- The `flex-shrink` controls how much a given element will shrink in relation to other elements. Keep in mind that **the shrinking only happens up to a `min-content` threshold!**. After that, the box will overflow.

  - You can achieve a neat effect where it seems like some elements start to shrink first, by using a very high `flex-shrink` value. It is an illusion. Remember that the setting is a ratio, other elements are shrinking as well, it is just that they shrink very slowly.

- Using the `flex-grow` will cause the element to **consume a ratio of available space**. If no other elements have that property, it will consume all available space.

#### The "flex" shorthand

- The `flex: 1` declaration **is NOT a shorthand for `flex-grow: 1`**.

  - It is a combination of three properties, and in the nutshell, it **allows you to ensure that children share the same amount of space** (the `flex-grow: 1` also takes the content into the account, so despite setting the `flex-grow:1` on all children, you might end up with different widths).

    - The `flex: 1` expands to the following.

      ```css
      flex-grow: 1;
      flex-shrink: 1;
      flex-basis: 0%;
      ```

      The `flex-basis` is crucial here. If the width is 0, then we can evenly distribute space, even if the content of the children is different.

#### Constraints

- Keep in mind that **you can still use the `max/min-width` and `max/min-height` to your advantage**.

  - The `flex-basis` operates akin to `width/height` but it will respect the `min/max` values.

    - This means that **`flex-basis` where the value is less than the content will be ignored**.

- Using the `flex-shrink` trick (with a very high value) comes in handy in many layouts.

#### Shorthand Gotchas

- It is vital to understand what the `flex: 1` declaration really means. If you do not, you might try to write the following.

  ```css
  .item {
    flex: 1;
    width: 200px;
  }
  ```

  In this case, **the `width` is ignored as `flex-basis` (0px) will ALWAYS win**. This is a rare case where the order of the declarations does not matter.

  What you ought to do, in such situation, is the following.

  ```css
  .item {
    flex: 1 1 200px;
  }
  ```

  Now, the `flex-basis` is set to `200px`. As it should have been.

### Wrapping

- With `flex-wrap: wrap` you can re-create the behavior of _inline positioned_ elements.

- **By default, the `width` of a flex element is at lest its content. This can cause issues when the element wont wrap even if it overflows its parent**.

  - You can **"fix" that by adding `min-width: 0`**. I bet you encountered similar situations where this was a problem.

### Groups and gaps

- Use **the `gap` property to create gaps between items**. This property is supported in all major browsers

- Use `margin-left/right: auto` to separate groups of items from each other. As an alternative, you could wrap items with `div` tags, but that pollutes the markup.

### Ordering

- You can **manipulate the visual order of the elements by using either the `-reverse` properties on `flex-direction` or `order` property**.

  - Keep in mind that **this is visual order ONLY!. If you wish to create different "tabbing" order, you have to modify the DOM**.

    - Modifying how elements sit in the DOM and using the `order` or similar property is very powerful. It allows you to create experiences, where visually, the elements are aligned left to right, but the tabbing starts from the right (think aside and a list of articles).

### Flexbox interactions

- When there is a **conflict between the layout modes, the _positioned_ layout always wins**.

  - This means that, a `fixed` positioned child will not participate in a flex layout.

  - There is an **exception for `relative` and `sticky` _positioned_ layout**.

- The **margin collapse does not work for flex layout**.

  - It only works for _flow_ layout so `block`, `inline` or `inline-block`.

- The `z-index` property works in flexbox just as it is working for _positioned_ layout.

### Recipes

- Using `flex: 1` or `flex: 2` and manipulating proportions can come in quite handy.

- Using `align-self` to "pull" the content in a `flex-direction: row` layout is crucial when using sticky positioning.

- Using `align-items` or `justify-content` **works with children that overflow the parent container** which is quite amazing.

### Workshop

- Using the additional `Side` element to center the nav menu, even though there is a logo in the nav as well was very clever.

- If you want to **use _baseline_ alignment and also have the children at the center** you can either manually center the elements via the padding, or add another flex wrapper.

- The combination of `min-width` and `flex:1` is powerful. Use it to create "grid-like" layouts especially if you cannot use CSS grid for some reason (of course, the best solution would be to use CSS grid).

## Flexbox encore

- Controls distribution of elements on the single primary axis. The axis changes depending on the `flex-direction` property value.

- When you apply the `display: flex` on an element, **you do not change that element layout mode, you change its children layout mode**.

  - This is crucial to understand. You can have an element with `display: flex` that behaves like a block element.

  - The **`flex` layout mode is NOT recursive**. If you set it on the parent, only the immediate children will be affected.

### Directions and Alignment

- To control which axis is the primary one, use the `flex-direction` property.

  - It really is helpful to get this mental model right. Instead of trying out random properties, knowing which axis is the primary one and using the correct declaration makes a difference. It also feels quite good :)

- The word **_justify_ is for primary axis, the word _align_ is for the secondary axis**.

#### Alignment Tricks

- In some cases, **instead of centering on the secondary axis, you might want to use the `baseline` alignment instead**.

  - Combine the `baseline` on the children with the `center` on the parent and you have a great looking layout!

- To align specific single child, use the `align-self` property on the child.

### Growing and Shrinking

- The **`width` and `height` are more like SUGGESTIONS rather than hard rules**.

  - The elements can shrink below their `width` and `height`.

- The `flex-basis` is **like `height` or `width` depending on which axis is the primary one**.

  - If you have `flex-basis` set, the `flex-basis` will always win with either `width` or `height` depending on which axis is the primary one.

    - The `flex-basis` **cannot scale the element below it's minimum content size**. If you set it to 0, the contents of the element will NOT overflow. This is not the case with width, where the contents of the element WOULD overflow.

  - The **`flex-grow` will override the `flex-basis` property**.

- The `flex-shrink` is for setting the **speed with which a given element shrinks relative to other elements**.

  - You can **disable shrinking** by setting the **`flex-shrink` to 0**.

  - A neat trick is to **apply a very large `flex-shrink` to one element, and a low number to another element**.

#### The `flex` Shorthand

- The `flex` shorthand is setting the **`flex-grow`, `flex-basis` and `flex-shrink`**.

  - You usually use it for **making the elements have equal size**.

  - It works mainly due to **setting the `flex-basis` to 0** – if the element width/height is 0, all space is "available space" that can be filled when an element has `flex-grow` set to 1.

#### Constraints

- While the `flex-basis` controls the _hypothetical width or height_ of an element, it **will respect the `max-width` or `max-height` properties**.

  - Using both `flex-basis` and `max-XX` allows you to put constraints on the element width/height.

#### Shorthand Gotchas

- The most common gotcha is \*\*setting the `width/height` and using the `flex` shorthand together.

  - Keep in mind that **the `flex-basis` will ALWAYS override the `width`/`height`**. The **order of CSS declarations does not matter here!**.

### Wrapping

- Like in the case of `inline` elements, **`flex` positioned children can wrap**.

  - Control wrapping via the `flex-wrap` property.

  - Keep in mind that **when an element wraps, it's size will most likely be different than the other children on a different row/column**.

    - This is because the sizing in flexbox is per row/column.

- When using wrapping, you are effectively creating _two-dimensional_ layouts.

  - There is an additional property for aligning the whole content – the `align-content` (so far, we have only used the `align-items` for secondary axis alignment).

### Groups and Gaps

- Use the `gap` property, it behaves as you would expect it to behave.

### Ordering

- Allows you to **change the VISUAL order of the flex children**.

  - This only applies to visuals, **the DOM structure is untouched!**.

- You have a couple of properties at your disposal.

  1.  The `order` property.
  2.  The `flex-direction` property.
  3.  The `flex-wrap: wrap-reverse`.

- A **common technique here is to reverse the order of the DOM and apply the `flex-direction: row-reverse**.

  - Imagine an `aside` that you want on the right, but you want it to focus first, before the content on the left.

  - Imagine a modal where the first button is "Cancel" but you want to focus the second button first (most libraries focus first focusable element).

### Flexbox Interactions

- An element can **only interact with a single layout mode when also participating in a `flex` layout**.

  - This means that **adding `position: fixed` to a `flex: 1` child makes all flex-related properties be ignored**.

  - As always, there is an exception to this rule – the _relative positioning_ and _sticky positioning_. You can use both `flex: 1`, and `position: relative`.

- The **_margin collapse_ is NOT a thing in flexbox**.

  - We are talking about the margin collapsing in _Flow_ layout mode.

- The **`z-index` works with _flex layout_ without the need to set the `position: relative`**.

### Recipes

- When using **`position: sticky` within a flex container, pay attention to how tall a given element is**.

  - If the **`position: sticky` does not seem to work, maybe the element has the height of the parent? If so use `align-self`**.

### Workshop

- Creating artificial margins via the `flex: 1` component is a great technique to ensure that elements stay in the center, even if there is an element on the left/right of that given element.

  ```tsx
  	<FlexSpacer>
  		</Content>
  	</FlexSpacer>
  	</MoreContent>
  	</FlexSpacer>
  ```

- To create _grid-like_ layouts, use the `min-width` and `flex: 1`.

  - The `min-width` is a firm statement, the elements will not go below that width.

  - The `flex: 1` tells the elements that they can grow and shrink. Keep in mind that **the elements will wrap if there is no available space on a given axis**. If you do not set the `min-width`, the elements will shrink to the _min-content_ size.

## Responsive and Behavioral CSS

### Working with Mobile Devices

- Nowadays, the phones most people use, have _high DPI_ screens. This means that `1px` in CSS does not necessarily correspond to `1px` on a screen – it could be `3px` or even more depending on the DPI of the screen.

- There is also this special `meta` tag which relates to how should mobile devices render your webpage.

  ```html
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  ```

  This will tell the mobile browser to match the viewport width with the device width and set the "zoom" to 1 (so the default).

#### Mobile testing

- Browser emulation can only do so much. To gain confidence that your webpage works as expected on mobile devices, you need to test it on... mobile devices.

- There are a lot of services which offer access to different mobile phones. The most popular one is BrowserStack.

### Media Queries

- Both _mobile-first_ and _desktop-first_ approaches makes sense. It all depends on the circumstances.

#### Other Queries

- Using the `:hover` **might not be what you want for mobile devices**. Depending on the browser, **tapping on an element might trigger the hover state and the second tap will trigger the click**.

  - There are tools to **only apply hover on devices with mouse/trackpads**. To do so, use the following media query.

    ```css
    @media (hover: hover) and (pointer: fine) {
      button:hover {
        text-decoration: underline;
      }
    }
    ```

    This reads: _when there is a hover state, and the pointer device is mouse or trackpad_.

- There is another group of media queries: the **preference-based media queries**. These gained a lot of traction in recent years.

  ```css
  @media (prefers-reduced-motion: no-preference) {
  }

  @media (prefers-color-scheme: dark) {
  }
  ```

### Breakpoints

- Use breakpoints to provide consistent experience across different mobile / tablet and desktop devices.

- When used with CSS in JS, you can interpolate the media queries into your styles.

- **Consider using `rem`s as your media query values instead of `px`**. The main advantage is that, when user changes the default font size, your website will behave accordingly.

### CSS Variables

- Note that **CSS Variables inherit from the parent (where they were declared)**.

  - You can have "global" or "local" CSS variables, it all depends on where you declare the variable.

  - In articles you will often see CSS variables declared on **`:root` which is an alias for `html`**.

  - You **can disable the inheritance by using the `@property` query, but it is NOT widely supported in all browsers yet**.

    ```css
    @property --text-color {
      syntax: "<color>";
      inherits: false;
      initial-value: black;
    }
    ```

- The `var` function **takes in the second parameter which acts as a _default value_**. This is pretty neat if you are **setting the CSS variables dynamically**.

  ```css
  .btn {
    padding: var(--inner-spacing, 16px);
  }
  ```

  - Keep in mind that **CSS variables are accessible via JavaScript**.

  - You can also **change the value of the CSS variable using a media query** which is pretty neat!

    ```css
    :root {
      --spacing: 16px;
    }

    @media (min-width: 600px) {
      :root {
        --spacing: 32px;
      }
    }
    ```

#### Variable fragments

- Since CSS variables **are composable**, they allow you to get rid of a lot of repetition.

  - All of this is possible because the CSS variables are **evaluated when used, and NOT when defined**.

### The Magic of Calc

- You can do math with CSS, using the `calc` property.

  - TIL that **you can mix units with `calc`**.

    ```css
    body {
      width: calc(50% + 32px);
      height: calc(24 / 16 * 1rem);
    }
    ```

### Viewport units

- From the previous section (the one about CSS variables and the `@property` rule) you might think that **CSS is a strongly-typed language**. That is **indeed the fact!**.

  - There are many units you can register – `<color>` or `<length>` and so on.

  - The _viewport units_ are of `<length>` type.

- The most frequently used are the `vh` and `vw` units.

  - They are not without their problems – **they do not work well on mobile where the viewport resizes when the address bar disappears**.

  - Luckily there are **other similar _viewport units_ which are dynamic in nature**. Read more about them [here](https://blog.webdevsimplified.com/2022-08/css-viewport-units/).

    - Sadly, [the browser support is not that great](https://caniuse.com/viewport-unit-variants).

### Clamping values

- The `clamp` value allows you to **create fluid layouts** where the `max/min-width/height` properties are not used.

  - This is cool because you **can combine the `clamp` with `min/max-width/height` for the ultimate responsive layout**.

    ```css
    .wrapper {
      width: clamp(/* minimum */ 500px, /* ideal */ 65%, /* maximum */ 800px);
      max-width: 100%;
    }
    ```

    The `.wrapper` will never be longer than either `800px` or `100%` which is very handy for "full width" layouts.

### Scrollburglars

- Use a **dev tools snippet to detect which element overflows the parent**, it will save you A LOT of time.

  - Read how to add a snippet to your dev tools [here](https://developer.chrome.com/docs/devtools/javascript/snippets/#create).

  - Use the following snippet:

    ```javascript
    const findOverflows = () => {
      const documentWidth = document.documentElement.offsetWidth;
      document.querySelectorAll("*").forEach((element) => {
        const box = element.getBoundingClientRect();
        if (box.left < 0 || box.right > documentWidth) {
          console.log(element);
          element.style.border = "1px solid red";
        }
      });
    };

    findOverflows();
    ```

### Responsive Typography

- As a rule of thumb, **you should not change the root font size**. Let the user do that.

- Depending on the content and its place (whether the text sits in a footer, or is a heading), first and foremost, we want to ensure that text is readable.

- **By default, iOS will zoom in forms if their font size is less than `16px`**.

  - You probably want to avoid this effect.

#### Fluid Typography

- With `clamp` you **can scale the size of the text depending on the viewport size** which is quite nice as it allows you to make heading scale with the viewport.

  - Depending solely on `vw` and `vh` does not cover **the case where the user uses the browser zoom feature**. To make your text scale accordingly, mix an unit which scales with zoom.

    ```css
    h1 {
      font-size: clamp(1.5rem, 4vw + 1rem /* 1rem scales with browser zoom */, 3rem);
    }
    ```

    We do not have to use `calc` here as `clamp` will automatically perform the conversion between units for us.

- **Use** these techniques **for headings rather than the body text**. The body text is already at a good size at its default.

### Fluid Design

- With fluid approach **you cannot apply styles conditionally** which might be a deal breaker in some circumstances.

  - Think a situation where you have to amend the `border-radius` property on mobile devices.

- The responsive approach only reacts to viewport changes, the **fluid approach reacts to the container changes**.

  - This drawback will become less relevant once [container queries are live](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Container_Queries).

- In the end it boils down to what is suitable for each situation. Keep in mind that it is better to be more verbose than to be clever.

### Workshop

- It seems like the `calc` support for `@media` **is spotty or non-existent**. I could not find a definitive answer, but it appears that it is not widely supported.

  - Instead of using `calc`, use interpolation to perform operations.

  - Of course there is a limitation that, without `calc`, you cannot perform operations with different units.

## Typography and Images

### Text Rendering

- Different browsers might display text a bit differently due to how they implement the _kerning algorithm_.

- _WebKit_ is a browser rendering engine created by Apple, and used in Safari and all mobile iOS browsers.

#### Text Overflow

- The default algorithm looks for _soft wrap opportunities_ (whitespace and `-`) characters. If found, it will break the word at that character.

  - You **could use HTML entities (`&nbsp;` and similar) characters to produce "non-breaking whitespace"**.

    - These can come in very handy when must ensure words stay together, like `10 USD`.

      ```html
      <p>That sandwich costs $10&nbsp;USD</p>
      ```

- TIL that there is an alternative text layout algorithm created by Adobe. It uses JavaScript to modify the HTML in a way that, arguably, is more appealing to the eye.

- It turns out **there is an alternative for `word-break` called `overflow-wrap`**.

  - I've always used the `word-break` and the `break-word` value which is **deprecated**.

  - Instead of using `word-break: break-word` one should use `word-break: normal` or `overflow-wrap: anywhere`.

- One can **use `hyphens: auto` to ensure that the browser adds the hyphens where a word break occurred**.

- Remember that you have to use `overflow: hidden` declaration for **single-line ellipsis**.

  - **Multi-line ellipsis is a bit more complicated, but still doable**.

    ```css
    p {
      display: -webkit-box;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 3;
      overflow: hidden;
    }
    ```

  - **Multi-line ellipsis should ONLY be used with _flow layout_**. If you use it with flex/grid layout modes, the result might not be what you expect.

### Print-style Layouts

- Yet another layout mode – the **Multi-Column Layout** allows you to split content into even columns.

  ```css
  .wrapper {
    columns: 3;
    column-gap: 16px;
    break-inside: avoid; /* Ensures that a given child is not broken in-between columns */
  }
  ```

- While `float` property is pretty niche, **it is the `float` property that allows you to wrap the text around an image or other blocks**.

  - You **cannot do that with any other layout mode / property**.

  - The **combination of `float:left` and `p::first-letter` allows you to create a book-like first sentence** where the first letter is bigger and "tucked" around the rest of the sentence.

### Masonry Grid with Columns

- The _masonry_ layout was popularized by sites like Unsplash and Pinterest

- It is **not possible to create a robust _masonry_ layout without using JavaScript at this very moment**. We can get close, but not quite there.

  - To get to "maybe good enough" state, you should use the `columns` property.

### Text Styling

- The **`1ch` is equal to the width of the `0` character, at the current font size**

  - For the longest time, I though that the `1ch` corresponds to a single character rather than a width of a certain character.

  - This means that **if you set `max-width: 50ch` that does not mean the element will wrap at 50 characters**. It all depends on the `0` character width. You might end up with more characters than 50.

  - In the end it does not matter than much. As long as your paragraphs are between 50-75 characters range, you should be good.

### Font Stacks

- We call it a "stack" because the **`font-family` property takes in multiple comma separated font names**. If a given font is not available, the browser will pick the next one from the list.

  ```css
  .title {
    font-family: "Lato", Futura, Helvetica, Arial, sans-serif;
  }
  ```

  - The last "item" in the list is the _category_ of the font.

- An alternative approach to using a "custom" fonts is to use the **system font stack**.

  ```css
  p {
    font-family:
      -apple-system,
      BlinkMacSystemFont,
      avenir next,
      avenir,
      segoe ui,
      helvetica neue,
      helvetica,
      Ubuntu,
      roboto,
      noto,
      arial,
      sans-serif;
  }
  ```

  Here, we more or less have a guarantee, that the website will look the same way for people on Windows and the same way for people on MacOS (though there might be differences between the operating systems).

### Web Fonts

- You can either **host a custom font yourself or use a 3rd party service**.

  - **Hosting the font yourself** has a number of **advantages**.

    1. Performance – there are no HTTPs handshakes involved since you are not reaching out to a 3rd party domain.

    2. Versioning – you can lock a version of your font. Since you are hosting it yourself, you are in full control over what's available to the application.

    3. Privacy – big brother Google (or other service) cannot track you.

    4. Performance – **You can embed the font file right into your CSS**.

- TIL that **if your font does not have a "bold" variant and you try to use one, the browser will try to emulate the "bold" variant**.

  - This usually does not end well. The "generated" font does not look as good as the original one.

  - The **same will happen for italic variant of the font**.

#### Font Loading UX

- There are **many ways to load a font files**. All depend on how crucial the font is, and what kind of drawbacks you deem okay.

  1.  The `font-display: swap` could cause **FOUT (_Flash of Unstyled Text_)** since the browser will not wait that long for the custom font to load. It will most likely swap it when the text is already visible.

  2.  The `font-display: block` could cause **FOIT (_Flash of Invisible Text_)** since the browser will wait for the custom font to load before displaying any text.

  3.  The `font-display: fallback` is a **happy medium** between the `swap` and `block`. The browser will wait some time, if during that time the font did not load, the default font is used. If the font loaded, the browser will swap the font when the text is already visible.

  4.  The `font-display: optional` is **for fonts which are not critical and displaying them is a minor improvement rather than a "feature"**. Here, the browser will wait for a short period of time for the font to load. If it does not load, it will use the default font and will NEVER swap to the custom font.

#### Font Optimization

- There a couple of ways you could optimize your font.

  1.  Use **a variable font** – more on that later.

  2.  Use **unicode-ranges to only download a subset of the font character set**.

  3.  Host the font files yourself.

- One neat trick is to download the font files from Google fonts and then host them yourself.

  - Font files coming from Google fonts are already well optimized.

#### Variable Fonts

- A _variable font_ is a font which is configurable and lives in a single files.

  - Historically you had to download multiple files for different `font-weight` values, with _variable fonts_ that is not the case.

- _variable fonts_ are great since you only have to download a single file. That file is also, most likely, smaller than the "legacy" version of the font.

- With _variable font_ you can **specify custom `font-weight` values like `777`**. The browser will NOT round up to the nearest available weight, the font itself will display as if the `font-weight` was `777` which is neat.

### Icons

- Two ways to have custom icons: **a custom font with icons or SVGs**.

  - **SVGs are much better** way of providing icons to the webpage. They are usually better looking, more accessible and can be tweaked. The font icons are kind of a black box.

- There are SVGs icon packages tailor-made for React.

  - Keep in mind that **SVG elements behave like `inline` elements (typography)**. This means that **they will have additional space added to them, even if you remove the `padding` and `margin` properties**.

    - To get rid of that "magic space", make them a `block` or `inline-block` element.

      ```css
      svg {
        display: block;
      }
      ```

### Images

- Remember about the **`alt` attribute**. It is important to specify it (your linter probably already complains if you do not).

- One neat fact about SVGs is that **SVGs can grow to infinite size without loosing fidelity**.

#### Fit and Position

- Images are considered to be **_replaced elements_**. The browser will replace the underlying `img` tag with a foreign entity.

- The `img` tag **is "special" in a sense that it is a `inline` element, but, by default, we are still able to use `width` and `height` properties on it**.

- The **browser tries to preserve the aspect ratio**, but if you provide both the `width` and the `height` properties, it gives up.

- You can **control how image displays within a container by leveraging the `object-fit` and `object-position` properties**.

#### Images and Flexbox

- **Images behave very "weirdly" when used as direct children of a `flex` parent**.

  - The best course of action is to **wrap the `img` tag with a `div`**. This way you will not loose your sanity trying to make them behave like block-level elements.

#### Aspect Ratio

- As noted earlier, if you set the `height` or the `width` property, the browser will try to preserve the aspect ratio of the image.

- If you set **both the `height` and the `width`, the browser will NOT preserve the aspect ratio** (something has to give).

  - To preserve the aspect ratio, **use the `aspect-ratio` property**. It is widely supported by all of the browsers.

  - Historically, developers utilized the `padding` property along with the _absolute positioning_ to lock in the aspect ratio. Luckily this technique is a thing of the past.

#### Responsive images

- With the `srcset` attribute (on the `img` tag) you can provide multiple resolutions of the same image. The browser will pick the best one and display it.

- As an **alternative to `srcset` attribute** one might use **the `picture` tag with multiple `source` tags inside it**. It looks a bit complicated, but it is most likely worth it.

  ```html
  <picture>
    <source
      type="image/avif"
      srcset="
        /cfj-mats/responsive-diamond.avif    1x,
        /cfj-mats/responsive-diamond@2x.avif 2x,
        /cfj-mats/responsive-diamond@3x.avif 3x
      "
    />
    <source
      type="image/webp"
      srcset="
        /cfj-mats/responsive-diamond.webp    1x,
        /cfj-mats/responsive-diamond@2x.webp 2x,
        /cfj-mats/responsive-diamond@3x.webp 3x
      "
    />
    <img alt="" src="/cfj-mats/responsive-diamond.png" />
  </picture>
  ```

  Here, you can **provide multiple formats as well as multiple resolutions for a given format**. If the browser does not recognize the image format, it will skip it. If everything fails, the browser will render a plain `img` tag.

  - As for the styling, the **`picture` element behaves like an _inline_ `span` element** so you might want to add `display: block` to it.

- You can generate multiple resolutions for a given image in your build pipeline. Look into `next/image` package.

#### Background images

- We use the `background-image` mainly for the ability to repeat the image with `background-repeat` property.

### Workshop

- To optimize a given font, see if it's not available in Google fonts. If it is, do the following.

  1.  Pick the right font weights and other properties.
  2.  Put the Google fonts link onto your website.
  3.  **Download the font from the Google servers**. You can do that by using "Open in new tab" while looking at the font request in the network tab.
  4.  Put the font file in your project.

  Keep in mind that, by doing all of these steps, you will shave a lot of time out of the request for the font. Since the font is hosted on your domain, there are no handshakes and other request that the browser needs to make.

- The combination of `picture` and `source` is really powerful. Use it!

## CSS Grid

### Mental Model

- **The rows/columns do not have to be the same size, but they do have to be consistent**.

- The rows/columns are _invisible_ from the DOM perspective. The layout purely driven by CSS.

### Grid Flow and Layout Modes

- Similarly to `flexbox` model, you enable the _Grid Flow_ by using the `display` property with a value of `grid`.

- CSS defaults to a grid of `1xN` where N is the number of rows.

  - The height of the grid container will distribute evenly between children.

  - You can **change how the grid children are placed by using `grid-auto-flow` property**.

    - This might feel similar to flexbox `flex-direction`, but **in CSS grid there is no "primary" or "secondary" axis**, so you do not have to worry about `align-items` switching the reference axis.

- As is the case with flexbox, **`display: grid` on the parent affects CHILDREN AND NOT THE PARENT!**.

- The CSS Grid interacts with `position: absolute`.

  - **If you set the `position:relative` on the grid, the children with `position: absolute` will be contained to a given grid cell**. Pretty neat!

    - You can **use this to move the element outside of its designed grid-area**.

### Grid construction

- **Unlike `flexbox`** the **size you give each column/row are HARD limits, not suggestions** (the difference between `flex-basis` and `width`, remember that `flex-basis` always wins!).

- **You can think of the `fr` unit as `flex-grow` proportions** (this means they will grow/shrink according to the size and also what they contain), but keep in mind that, instead of working on children, it scales the columns/rows instead.

- To create rows in `grid` layout, use `grid-template-rows`. To create columns, use `grid-template-columns`.

  - There are multiple nifty functions you can use while creating rows/columns.

- Keep in mind that you can **mix `grid` with `flexbox` (on children)**.

- **Be mindful of the tab order when placing items in the grid**. The tab order will, by default, honour the DOM order. If you put items all over the place, the users might have problems traversing the page with keyboard.

#### Alignment

- The `justify-content` **controls how COLUMNS are distributed INSIDE the grid parent**.

  - The `justify-items` (**ignored in `flexbox`) applies to the child elements (but you declare that property on the parent)**.

- The `align-content` **controls how ROWS are distributed INSIDE the grid parent**.

  - The `align-items` **applies to the child elements, and controls how a child is distributed inside a given row**.

#### Grid Areas

- One of the coolest things about the Grid layout is that **you can explicitly define the names of the rows/columns and assign elements to those areas**.

  ```css
  .wrapper {
    display: grid;
    grid-template-areas:
      "sidebar header"
      "sidebar main";
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 80px 2fr;
  }

  <!-- And so on... -- > .sidebar {
    grid-area: sidebar;
  }
  ```

  - This makes it so **you can shuffle the elements in the DOM, but the layout will stay the same**. Each element is assigned to a designed "area" in the grid, so the order of DOM nodes does not matter. Having said that, **keep the keyboard and focus navigation in mind while changing the DOM order**.

- It seems like I'm **unable to use `max` or `min` functions to denote the width/height of a column/row**. The only similar supported function is the `minmax` one and **the combination of `minmax(min/max(VALUE), VALUE)**.

#### Track and Lines

- The `grid-template-areas` is a syntactic sugar over the `grid-column` and `grid-gap` properties.

- The CSS Grid composes of tracks (rows and columns). Each track has a number assigned to it (starting from 1, or -1 if you are looking at bottom right).

  - The **negative track numbers allow you to ensure a given element spans the whole grid area**, like so.

    ```css
    .grid-child {
      grid-row: 1 / -1; /* Spans all the rows, no matter how many of them there are */
    }
    ```

    Note that **the `/` symbol does not indicate division, but rather a separator**.

### Fluid Grids

- There are few ingredients you need to create a _fluid grid_.

  1.  The `repeat` function.
  2.  The `auto-fit` or `auto-fill` property.
  3.  The `minmax` function.

  The resulting snippet looks the following:

  ```css
  .wrapper {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  }
  ```

  - **The `minmax` is a combination of `min` and `max` functions and works akin to the `clamp`**. Keep in mind that **the first value cannot be "flexible" like 1fr**. Using `px`, `rem` or `%` works fine.

    - A good mental model of `minmax(X, Y)` is thinking about it in the following way:

      > I want my column to be `Y`, but never be less than `X`

- The **single most crucial difference between the `auto-fill` and `auto-fit`** keywords is the **how they treat the available space**.

  - The `auto-fill` **will create empty columns** where there is available space to do so.

    > `auto-fill` the container with as much columns as you can.

  - The `auto-fit` **will stretch existing columns to fill the available space**.

    > `auto-fit` the **existing columns** to the available space.

- To create a **truly fluid grid layout** one has to use the `min`/`max` function alongside the `minmax` function.

  ```css
  .wrapper {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(min(250px, 100%), 1fr));
  }
  ```

  This makes it so that **the minimum column size is computed and dynamic**. Say goodbye to the overflow on mobile devices!

- You **cannot use the "intrinsic" sizes with the `repeat` function**. This means that `min-content`, `max-content`, `auto` will NOT work.

### Grid Dividers

- It is **not possible to style grid dividers (tracks)**.

- If you want to create an effect where there tracks of CSS grid have different color, use the **background on the grid elements** and the **background on the grid container**. This way you can create an illusion of "borders" that correspond to grid dividers.

### Grid Recipes

#### Two line center

Instead of using `flexbox`, like so

```css
.wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}
```

You could use `grid`, like so:

```css
.wrapper {
  display: grid;
  place-content: center;
}
```

The `place-content: center` is a shorthand for `justify-content: center` and `align-content: center`. Personally, I see it more of a party trick than anything else. We are saving one line of CSS, but the probability of someone having an understanding of this snippet is, I would argue, low.

#### Sticky grids

- **Sticky position does NOT care about grid rows and columns**. As far as it is concerned, the children are constrained by the parent, not the grid tracks.

- If the element does not "stick" **it might span the whole container, despite a small amount of content**. In such situations, you most likely want to use `align-self: start` or similar.

#### Full Bleed Layouts

- Instead of fighting with negative margins (although they are required for padding), one could use CSS grid to create the so-called _full bleed layout_.

  ```css
  .wrapper {
    display: grid;
    grid-template-columns: 1fr min(30ch, 100%) 1fr;
    padding: 0 16px;
  }

  .wrapper > section {
    grid-column: 2;
  }

  .full-bleed {
    grid-column: 1 / -1;
    margin: 0 -16px;
  }
  ```

  Keep in mind that **the `1ch` represents the WIDTH of the `0` character and NOT the number of characters**.

  - To get rid of the implicit dependency between the parent `padding` and the children `margin`, you can use CSS variables.

### Managing Overflow

- The **`fr` unit will grow based on the children**.

  - To allow the `fr` unit to be as small as it needs to be, use `minmax`

    ```css
    .wrapper {
      grid-template-columns: minmax(0, 1fr);
    }
    ```

### Grid Quirks

- Some browsers might limit the number of grid rows. If you bump on this limit, you are rendering way too many things at once. Look into [DOM virtualization](https://github.com/bvaughn/react-window).

- **Similarly to `flexbox`, grid children will NOT collapse margins**.

- You can use the `z-index` on a grid child without having to use the `position` property.

### Workshop

- There are **two ways to center something in relation to the WHOLE page (not a row)**. Think centering a logo which is then accompanied by menu on the left and some other actions on the right.

  1.  Using the _spacer_ components with `flex: 1`. This is what we have seen in one of the previous workshops.

  2.  Using the `1fr` unit in the context of CSS grid and the `auto` on the child you want to center relatively.

- Wrapping elements with borders and creating vertical space between the borders is surprisingly challenging.

  1.  One could use the `:after` or `:before` elements with a given `height` to simulate the border. Then you can also set margins on the `:after` or `:before`.
  2.  You can create a wrapping elements which only do the border. This approach seems simpler and easier to understand.

- Use the `white-space: prewrap` to honour the whitespace in the text. This is especially important when dealing with text from a 3rd party, like CMS.

- To **create "borders" between grid tracks, you can either use `margins` and `border` or one of the children or you can use special "divider" elements that fit nicely inside the grid**.

  - Personally I much prefer the "dividers" approach. This approach decouples the borders from the child styling.

- To **force the grid to overflow in horizontal direction** you can use the `grid-auto-flow: column` and `grid-auto-columns` property. This forces the children NOT to wrap and using the `grid-auto-columns` you can set the minimum width.

- The "world famous" CSS grid snippet requires you to have at least one "concrete" value in the `minmax` function. If you were to use `fr` or `min-max` or similar sizes, the declaration will not be correct.

## Animations

### Transforms

- Using the **`translate` property will NOT change the element in-flow position**. This means that it works differently than margin. If you use `translate` on element X, the elements Y layout will not be affected.

  - If you use **percentages as the value for `translate`, these are relative to the ELEMENT size and NOT the container size**. That is in stark difference to the `left/right/top/bottom` properties.

  - Keep in mind that you **can mix units using the `calc` function**.

    ```css
    .element {
      transform: translateX(calc(100% + 50px));
    }
    ```

- The **`scale` will transform the children of the element as well as the element itself**. If not countered, you will see the children stretch.

  - You might want to use `overflow: hidden` on the parent element to manage "spill-over" of the element you apply the `scale` on.

- The **order of operations for the `transform` property is important**. The operations are applied **sequentially one after each other – RIGHT TO LEFT!**.

> Side note on the order of properties. It seems like the browsers are also matching the selectors from right to left.

- The **`transform` property does not work with `inline` elements**.

### CSS Transitions

- You can **animate multiple properties using a comma separated list**.

  ```css
  .wrapper {
    transition:
      transform 250ms,
      opacity 400ms;
  }
  ```

  Notice that each property has a different timing.

  - There is also the `transition:all`, but in reality it is a footgun. You should not use it, it is overly generic.

- Consider using the `ease` and other _easing functions_ to make your animations feel more authentic. The **`linear` is almost never (if not ever) a good option**.

  - You can create your own easing functions with the `cubic-bezier` function.

- When **translating, consider applying animations on the children, not the element itself** when using **hover as the animations trigger**.

  - If you do not do this, you might observe an effect where hovering at the very bottom of the element causes the element to "jump" depending on the animations direction.

    - This is caused by the fact that the hover applies on the element and that element moves (the `transform`). If you do not move your mouse, the element will loose hover and transition back to its origin. But you have a mouse waiting out there to trigger the animation again, and so it repeats.

### Keyframe Animations

- You can either use the `from` and `to` keyword or percentages.

  - The browser does a good job of interpolating between different states, but [as this video shows](https://www.youtube.com/watch?v=azoIMhKOucQ), you might come across some interesting edge cases.

- You can apply _easing functions_ via the `animation-timing-function`.

- The **timing function applies to EACH STEP**. We do NOT get a single ease for the whole duration.

  - This means that **the more `keyframes` you add, the longer your animation will be**. Keep that in mind.

- There is a shorthand, but I find the explicit declarations much more readable.

- The **styles in `@keyframes` will take precedence over your regular styles**.

  - On some browsers, you might even be able to override the `!important` declarations.

#### Fill modes

- The `animation-fill-mode` allows us **to persist the animation end-state after the animation is finished**.

  - How many times you wrote an animation, only for the element to snap back to it's original state?

  - You **usually want to use `forwards`, but the `backwards` can be used when you want to reveal an element**.

  - Or you can use the `both` value for `animation-fill-mode` property.

#### Dynamic Updates

- You can apply the animations _dynamically via JS_, similar how you can change CSS variables.

  - If you remove the animation styles altogether, you might see an abrupt end to an animation.
    **Consider using `animation-play-state: paused` for disabling animations**.

#### Exercises

- Interestingly, you can **declare variables inside the `keyframes` keyword**.

  - Those variables are scoped to the element where the `keyframes` is used.

- It would be really great if we could have a tool which shows where the `transform-origin` is.

  - I'm very surprised that is not a thing, or at least appears not to be a thing, in the Chrome dev tools

### Animation performance

- You should **favour `transform` and `opacity` animations heavily**. Animating other properties is compute expensive and might cause the animation to be sluggish.

  - Some properties are not that bad™️, the `background-color` will not cause a layout shift so animating it is much better than animating the `width` or `height`.

  - **The styles of a given element also plays a huge role in how expensive animating it would be**. For example, **tweaking the `height` on `fixed` positioned element will be very cheap** as it does not "interact" with other elements, layout-wise.

- The `will-change` property **instructs the browser to hand-off the rendering of an element to GPU on page load**.

  - By default, if you apply the `transform` animations, the CPU hands off that element to GPU. There are subtle rendering differences between the two, so you might see text rendered differently.

  - If the browser hands off the element to the GPU when the page loads, there will be no change in how the element is rendered when you animate it.

### Designing animations

#### Orchestration

- Sometimes, to create a great looking animations, you might need to stagger multiple elements – animate them "in turn".

  - If you want to separate the _enter_ and _exit_ animations and stagger multiple elements, you would be better off using a library. There is a lot to code to write (JavaScript code since it requires state).

- Do not forget about the `onTransitionEnd` callback. **It could be useful if you want to animate multiple elements one after the other**.

### Accessibility

- Use **the `@media (prefers-reduced-motion: reduce) {}` query to disable one-off animations**.

  - You can **access this, and other queries, using the `matchMedia` function in JS**.

  - A **better way** would be to **start with no animations, and then add them inside `@media (prefers-reduced-motion: no-preference)`** query.

    - Why do we start with no animations and not disable them globally? **You most likely do NOT want to disable all animations altogether**. Some animations are pretty much required so that the user is aware that something changed. For those cases, you should reduce the animation "umph" rather than disabling it.

### Ecosystem World Tour

- The _Web Animations API_ is a JavaScript alternative to CSS `@keyframes`.

  - The **main difference is that the _Web Animations API_ enables you to provide a single timing function for the WHOLE animation**.

- The _Framer Motion_ library is really good if your application is written using React.

### Workshop

- It is **important to add _hover_ styles only when it makes sense to do so**.

  - Keep in mind that, on mobile devices, the _hover_ state might trigger when user clicks on an element.

  - To enable _hover_ state **only on mouse-powered devices** use the `hover` and `pointer` **media features**.

    ```css
    .wrapper {
      @media (hover: hover) and (pointer: fine) {
        &:hover {
        }
      }
    }
    ```

- Sometimes, to achieve the animation you want, you will need to duplicate the HTML.

  - If you do that, consider using `aria-hidden` so that the screen reader does not read the repeated markup.

## Little Big Details

### CSS Filters

#### Color Manipulation

- You can animate the `filter` property as the **alternative to animating the `background-color`**. The `filter` is hardware-accelerated in SOME browsers.

- There are multiple _filter-related_ CSS functions, like `brightness`, `contrast` and `grayscale`.

  - You can apply multiple of them onto a single element.

    ```css
    .wrapper {
      filter: brightness(120%) contrast(110%) grayscale(50%);
    }
    ```

#### Blur Filter

- Using **`blur` is really expensive. Make sure you have a good reason for using it**.

- Keep in mind that **`blur` as means of hiding content is not really effective**. People who use screen readers, will be able to read the content. Use `aria-hidden` or similar.

#### Backdrop Filters

- Sometimes, when you use `blur` you really only want to `blur` the background of an element, not the element contents. This is where the `backdrop-filter` comes in handy.

- The `backdrop-filter` can also be used with other properties, not only the `blur`.

  ```css
  .header {
    backdrop-filter: brightness(150%) hue-rotate(30deg) blur(5px);
  }
  ```

### Border Radius

- TIL that the **more specific variations of the property, like `border-top-left-radius` take in two values**.

  - It turns out, even the `border-radius` property itself takes in multiple values.

    ```css
    .wrapper {
      border-radius: 10% 20% 30% 40% / 50% 60% 70% 80%;
    }
    ```

#### Nested Radiuses

- If you apply the `border-radius` on the element and its child, the rounding might look off.

  - The edges will not be of the same size, even if you set the `border-radius` to have the same size.

  - To make sure whey optically are the same, you will have to increase the element `border-radius`.

#### Circular Radius

- The **`border-radius` has an implicit limit which is relative to the element height**.

  - If you provide a value bigger than the limit, the maximum value of the `border-radius` will be applied.

  - This allows you to create really funky looking shapes.

### Shadows

- There are a couple of ways to create a shadow in css.

  1.  The `box-shadow`. Arguably the most commonly used.
  2.  The `filter: drop-shadow`. It **produces slightly different shadow effect than the `box-shadow` one**.
  3.  The `text-shadow` which I had no idea existed.

#### Contoured Shadows

- Here the difference between the `box-shadow` and `filter: drop-shadow` is much more pronounced.

  - The **`filter:drop-shadow` applies the shadow to non-transparent parts of the image**, while the **`box-shadow` will apply the shadow effect to the whole image**.

    - As you can imagine, in most cases, you want to use the `filter: drop-shadow`.

- If you want **the shadow to "follow" the contour of an element, use the `filter: drop-shadow`**.

#### Single-Sided Shadows

- This is **where the `box-shadow` shines** as the **`filter: drop-shadow` does not allow you to "direct" the shadow in a given direction**.

#### Inset Shadows

- Only available via the `box-shadow` property.

#### Designing Shadows

- Consider **designing shadows while looking at your website holistically**.

  - To take things even further, make sure that **all shadows follow a ratio** which you agree on.

  - This way your website will feel more consistent.

### Colors

#### Accessibility

- It is vital to ensure the text on the page has enough contrast in relation to the background. Otherwise some people will not be able to read your website!

  - There are many accessibility-related tools to help you here. Chrome has a built-in contrast checker.

  - This does not only apply to text, you should also mind the images.

#### Selection Colors

- You can change the _selection color_ (usually blue by default) using the `::selection` pseudo-class.

#### Accent Colors

- Have you ever tried to style the default `<input type = "checkbox">` realizing that you are very limited in the styling options?

  - **This is where the `accent-color` comes in**. It **allows you to specify the colors for native elements like radios, checkboxes, sliders and progress elements**.

  - It might be just enough for the designer to be happy and you not having to rebuild some of the native elements.

### Mobile UX Improvements

- We should not be making buttons tiny in general, but this rule is very important for mobile devices.

  - **You can increase the "surface area" of a given element by using the `:after` or `:before` pseudo elements!**.

    - This technique is a good **alternative to increasing the element width or height**.

      ```css
      button {
        position: relative;
        height: 32px;
      }

      button::after {
        --tap-increment: -8px;
        content: "";
        position: absolute;
        top: var(--tap-increment);
        left: var(--tap-increment);
        right: var(--tap-increment);
        bottom: var(--tap-increment);
      }
      ```

- Consider using **`user-select: none` for buttons**. On mobile it can **be cumbersome to tap a button only for the OS to select the text inside of it**.

### Pointer Events

- With `pointer-events` property you can... disable the pointer events on a given element!

  - Keep in mind that focus and text selection will still work.

  - This property **applies to all children, unless you set the `pointer-events: all` on a given child**.

### Clipping With clip-path

- Do you remember creating different shapes with CSS borders? Well, these times are over.

  - The `clip-path` is well supported and is much more idiomatic than the invisible borders.

#### Animations

- You can even animate the transition going from different `clip-path` values! Check this out.

  ```css
  .triangle {
    clip-path: polygon(0% 0%, 100% 0%, 100% 100%, 0% 100%);
    transition: clip-path 250ms;
    will-change: transform;
  }

  .triangle-wrapper:hover .triangle,
  .triangle-wrapper:focus .triangle {
    clip-path: polygon(0% 0%, 100% 50%, 100% 50%, 0% 100%);
  }
  ```

  Pretty amazing if you ask me.

#### Rounded shapes

- It turns out the `clip-path` is usable with different _shape-defining_ functions, like `ellipse`, `circle` or `polygon` (the one we have seen already in the previous example).

  - All of this makes it so that **the `clip-path` can be used like a super-charged `border-radius`**.

- Keep in mind that you can animate between different shapes. Such ability unlocks a new way creating visual experiences.

#### With shadows

- The `drop-shadow` function applies \*\*before the `clip-path`.

  - This means that you will most likely not see the shadow as it will be clipped.

  - You **have to apply the shadow on the parent**.

#### Exercises

- Keep in mind that **you can reveal elements using `clip-path`**. This allows you to create reveals that seem to "blend-in" with existing elements.

  - This is possible because you can **declare the same coordinates for multiple `clip-path` values**.

### Optical alignment

- Keep in mind that **text takes in more space than in Figma, and this is without the "magic space"** of an _inline_ element.

- Sometimes, you have to tweak the space manually. It **is better to make the spacing FEEL even than to be mathematically even**.

### Scrolling

#### Smooth Scrolling

- The _smooth scrolling_ **should ONLY be applied on interaction, and NOT globally**.

  - If you apply it globally, it is known as **_scroll-jacking_** and it is never a good idea.

- If you do apply the `scroll-behavior: smooth`, **make sure to guard it behind the `prefers-reduced-motion: no-preference` query**.

  - Keep in mind that you cannot control the speed and the easing of the scroll animation. It is controlled by the OS.

- The `scroll-behavior: smooth` has a JS equivalent, the `scrollIntoView` or `scrollTo` with the `behavior` parameter.

#### Scroll snapping

- While implementing a slider menu, you probably wrestled a lot with making sure the scroll stops at a given image. All this work that historically been done in JS can be thrown away and replaced with two lines of CSS.

- Use the `scroll-snap-type` and `scroll-snap-align`. These properties are well supported in all browsers.

  - These properties really make our job easier.

#### Scrollbar Colors

- It is possible to style the scrollbar, but you should not overdo it.

  - If you completely replace the native one, users might be confused and do not recognize the scrollbar!

#### Scroll Optimization

- To avoid **clipping content when scrolling to anchors with sticky headers, use the `scroll-margin-top` property**.

  - So good that it exists. I vividly remember having to deal with it without this property.

- To **improve CLS** on sites with async data, **consider using `overflow-y: scroll` to force the vertical scroll by default**.

  - The scrollbar track shifts the page by a couple of pixels. On some websites, this experience might be jarring as the content loads. To prevent this shift, we make sure the scrollbar is always displayed.

### Focus Improvements

#### Focus Visible

- Works on the same basis as `:focus` but **triggers only when the element is focused and the user is using a non-pointer device like the keyboard**.

  - This is a huge win, as previously, it was very annoying to have the focus rings on clicks.

  - What's more, **nowadays most of the browser styles use `:focus-visible` instead of `:focus`**.

- The bottom line is that **there is NO REASON TO REMOVE THE FOCUS OUTLINE**.

#### Focus Within

- It enables you to **style the parent based on the focus state of the child**.

  ```css
  .parent {
    &:focus-within {
      background: red;
    }
  }
  ```

  **But it matches the `focus` state and not the `focus-visible` state**. There is a difference between these two states.

#### Focus Outlines

- The focus outline **works differently for interactive and non-interactive elements**.

- In general, we should not apply styles to the outline. The styles are a convention and everyone is used to them.

  - You can experiment with `outline-color` but not all browsers support it.

### Floats

- You might consider `float` property obsolete, but that is not the case. They have a place in the modern CSS ecosystem!

- They are mainly used for **wrapping the text around images / or other shapes**.

  - Yes, **shapes**. The element **does not have to be a rectangle. It can be an circle or an image itself**.

  - To specify the exact shape, **look into the `shape-outside` property**.
