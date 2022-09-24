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
