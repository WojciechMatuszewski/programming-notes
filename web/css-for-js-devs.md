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
@media (max-width: 300px) {} /* valid */
@media (font-size: 32px) {} /* invalid */
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

2. The `rem` unit which **is relative to the ROOT font size**. By default, the root HTML font size is `16px`.

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
      ...aStyles
    }
    ```

  - When determining the end-result, one has to also take into the account the **specificity of a given CSS selector**.
    While important when you write vanilla JS, if you use modern tooling, you do not really have to know the specificity of a given selector.

### Directions

- There is the **`block` direction (think lego blocks stacked on top of each other)**, and there is the **`inline` direction (think people standing in a line)**.

  - The above holds true for English and vast majority of other languages, but it's not exactly true for the arabic and some of the Asian languages.

  - You might want to consider **using the `margin-block-start`, `margin-block-end`, `margin-inline-start` and `margin-inline-end`** properties to style margins. These are universal and will adjust accordingly based on the direction of the document.

### The Box Model

- The box model describes **how big a given element will be**.

  - The box model **describes how the content, `padding`, `border` and `margin`** interact with each other.

- By default, the browsers specify the `box-sizing` to have a value of `content-box`.

  - This means that, the `width` and `height` of the child does not take into the account the padding and the margin of that element.

  - Since the behavior described above might be confusing, one can use the `box-sizing: border-box` declaration.

    - The `border-box` means that `width` and the `height` properties should account for padding, margin and border properties of a given element.

#### Padding

- The **_inner space_ of a given element**.

  - You have your usual suspects of `padding-left` and such, but **you should also consider using the logical properties like `padding-block` or `padding-block-start` or `padding-inline-start` and so on**.

- For padding, **use `px` rather than other units**. You most likely do not want the padding to change alongside the text size.

- To **remember the shortcut notation, imagine a clock. Start from the `12:00` and go clock-wise**.

  - So the definition like `padding: 10px 20px` means `10px` for up/down and `20px` for right and left.

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

- The **`inline-block` elements behave as _inline_ for the parent, but they accept _block_ element properties**.

  - This allows you to provide a margins and other properties to the element, but still have it flow _inline_.

  - Keep in mind that **`inline-block` elements do not wrap**.

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
    <p style = {{margin-bottom: 20px}}>
      foo
    </p>
    <br/>
    <p style = {{margin-top: 20px}}>
      bar
    </p>
    ```

  But, the following WILL collapse

    ```html
    <div>
    <p style = {{margin-bottom: 20px}}>
      foo
    </p>
    </div>
    <p style = {{margin-top: 20px}}>
      bar
    </p>
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

### Stacking Contexts

- There are many factors which influence how elements stack on top of each other. Most notable are the **layout mode and the DOM order** and the **`z-index`** value.

  - As a rule of thumb, the **_positioned_ elements will always render on top of non-positioned ones**.

- The famous **`z-index` property only works with _positioned_ elements and grid/flex children**.

  - The **value of the `z-index` is compared relative to all the elements in a given stacking context**. A very important nuance to understand as this is the reason you sometimes find yourself bumping the `z-index` to a very high value to no avail.

    ```html
    <header zIndex = "2" position="relative">
      My header
    </header>

    <main zIndex = "1" position="relative">
      <div zIndex = "999999">VeryHigh</div>
    </main>
    ```

  The `VeryHigh` **will NEVER overlay on top of the `My header`**. The `999999` is very high, but it only is comparable to other `zIndex` values inside the `main`.

- There are many ways one could go about creating the stacking context. The **most common way I've seen is to use the `position: relative` and `z-index: SOME_NUMBER` declarations**. You can find the full [list here](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Positioning/Understanding_z_index/The_stacking_context).

#### Managing z-index

- To avoid "`z-index` wars" you **should create stacking contexts when necessary**.

  - As noted previously, there are multiple ways one could do that, but **the best way to avoid the "`z-index` wars" is to not use the `z-index` at all!**.

  - To **create a stacking context without using the `z-index` property, use the `isolation: isolate` declaration**.

- Another very **useful thing you could do is to set the `isolation: isolate` on the `#root` element of your React application**. This guarantees that any modals (injected via portals) which render at the level of the `#root` element do not overlap with any elements from within the application itself.

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
