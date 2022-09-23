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
