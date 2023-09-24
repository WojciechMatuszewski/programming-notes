# CSS

## Spacing

Instead of trying to memorize the order of the sides in the `padding` or `margin` properties, use **`padding-inline` (left-right) and `padding-block` (top-down)**. The same applies to margins. One has to swap the `padding` for margin.

There is an additional benefit to using these properties instead of the "regular" `margin` and `padding` properties. **They take the `writing-mode`, `direction`, and `text-orientation` property into the account**, which makes your code more reusable.

## Selectors

### Specificity

Specificity is based on column-like structure (0-0-0).

- 1-0-0 for `#`
- 0-1-0 for `.`
- 0-0-1 for `element`

Whichever column has the bigger number in the left-most column wins. **If
selectors are identical the bottom most wins (cascading styles)**

### Relational Selectors

There are quite a few of them. Let's consider following structure

```html
<ol>
  <li class="some_class">item1</li>
  <li>item2</li>
  <li>
    item3
    <ul>
      <li>item3.1</li>
      <li>item3.2</li>
      <li>item3.3</li>
    </ul>
  </li>
</ol>
```

#### `>` selector

_Child selector_, only children of a given element will be selected

```css
/*
    matches <li> in <ol> but not the ones nested inside <ul>
*/
ol > li {
}
```

#### `+` selector

_Adjacent sibling_. Pretty self explanatory.

```css
/*
    matches ONLY item2
*/
li.some_class + li {
}
```

#### `~` selector

_General sibling selector_. Matches all later siblings but not nested!

```css
/*
    matches item2, item3
*/
li.some_class ~ li {
}
```

### Attribute Selectors

- you can query by attribute presence, **value does not matter**

```css
/*
    match all element that has alt on them (value does not matter)
*/
element[alt] {
}
```

- you can query by exact value

```css
element[alt="some image description"] {
}
```

- value begins with something, **it ignores any dashes that come after it**.
  Mostly used on languages attributes.

```css
/*
    match all element that has alt which starts with: "some"
*/
element[alt|="some"] {
  /*
  <element lang = "some-us">
  <element lang = "some">
  */
}
```

- you can query by value which starts with (`^`) ends or (`$`) or has (`*`) some
  query

```css
element[alt$="some"],
element[alt*="some"],
element[alt^="some"] {
}
```

- you can also force case insensitivity using `i` When using
  `element[alt="some"]` that query inside quotes will match using case
  sensitivity by default. You can change that behavior

```css
element[alt="some" i]
```

### User Interface Selectors

These are:

- `:enabled`
- `:disabled`
- `:checked`
- `indeterminate`

And many more, just read the docs...

For example:

```css
input[type="checkbox"]:checked + label {
  color: red;
}
```

### Structural Selectors

These include:

- `:only-child`
- `:nth-of-type`
  - even
  - odd
  - an + b (offset can be negative)

And more... But what peaked my interest are those:

- `:root` (Angular2+ similarity?)
- `:empty`
- `:blank`

#### `:root`

Matches the root element. In HTML its the `<html>` tag.

#### `:empty`

Matches element that has no content, is self closing or contains only a comment.
**It cannot have any whitespace**.

#### `:blank`

Not really supported. Works like `:empty` but can contain whitespace.

### `:matches`

This one is wild. Look at the syntax

```css
element: matches(#home, .someClass, [title]);
```

Pretty neat huh? This one is almost like `querySelectorAll`. Of course this
would be too good to be true. **Support for this is quite meh**.

## Pseudo-classes

### The `:is` and `:where`

The `:is` and `:where` selectors are used to eliminating repetition that sometimes you have to deal with when selecting multiple children of a given element.

  ```css
  :is(article) h2,h3,h4,h5,h6 {
    color: var(--fire-red)
  }

  // the same as \/
  article h2, article h3, article h4, article h5, article h6 {
    color:var(--fire-red)
  }
  ```

The only difference between the `:is` and `:where` selectors is the specificity. The **`:where` selector has a specificity of 0, whereas the `:is` adheres to the regular specificity rules of selectors**.

### The `:has`

The `:has` selector enables you to **style the parent based on it's the children**. This is a groundbreaking change in how we think about CSS. Since always, the CSS rules had to obey the _cascading_ semantics – you could not "go back" to the parent.

The following is an example of the `:has` selector where we specify the number of columns based on the list items.

  ```css
  ul:has(li:nth-child(6)) {
    columns: 2;
  }

  ul:has(li:nth-child(11)) {
    columns: 3;
  }
  ```

Pretty amazing stuff.

### `:focus-visible`

The `:focus-visible` uses browser UA heuristics to determine when to display the focus outline. **This is not the case with the "regular" `:focus` pseudo-class** which is "dumb" in that regard. You **might have needed this to disable the focus outline on link clicks**.

In fact, most (if not all) of the browsers, migrated from `:focus` to `:focus-visible` in their UA styles.
Did you notice that, if you click a link, the focus outline is not there? **But if you use the keyboard, the focus outline is visible?**. This is the `:focus-visible` pseudo-class in action.

### `:not`

Well supported. The gotcha is that the **selector inside parenthesis must be simple**. By simple we mean no combinators and spaces.

```css
element:not(img); /*ok*/
div:not(.someClass); /*can be also an id*/
div:not(ul li); /*will not work, there is a space, it's not a simple selector*/
```

#### Using the `:not` with elements that DO NOT contain certain properties

What if you wanted to select all `p` tags that DO NOT have any class? Or maybe an image that does NOT have the `alt` tag?

```css
img:not([alt]) {
  outline: 5px solid red;
}
```

I find it pretty amazing that we can do so much logic in CSS. Of course I was not aware that using `[]` syntax checks for the _presence_ of a given attribute.

## Pseudo-elements

### `::first-letter`

You can target the first letter of any element that has text inside. This way
you can create _book-like_ text

```css
p::first-letter {
  font-size: 40px;
  color: red;
}
```

### `::selection`

You can actually style the behavior of selected (as in mouse selected) stuff.
This is pretty neat!.

```css
/*
  all selected paragraphs will have red color
*/
p::selection {
  color: red;
}
```

### `::before` and `::after`

These **has to have `content` property**. That content is actually not part of
the DOM, you cannot highlight it.

```css
p::before,
p::after {
  /*required!!!*/
  content: "";
  display: block;
  width: 50px;
  height: 50px;
}
```

One of the less known features is that if you do not want to show any content at
all you can be a _pro leet hackorz_ and use `none`

```css
p::before {
  content: none;
}
```

You can also do citations with quotes, kinda nice

```css
p::after {
  content: close-quote;
}
p::before {
  content: open-quote;
}
```

### `::placeholder`

Self explanatory. You can style placeholders 🤷‍

## Progressive enhancement

Lets say you want to use CSS grid where possible and for older browsers some
simple, other, layout.

How would we go about detecting if we can use grid ?

Well it turns out there is are so called **CSS Feature Queries**.

These can be use to detect if a CSS feature is available on given browser.

```css
@supports (display: grid) {
  // you code here
}
```

Pretty cool stuff!.

## Logical properties

Instead of using the `left` and `right` keyword for paddings, margins and so on, one **should consider using the `inline` and `block` keywords**.

```css
.foo {
  margin-inline-start: 10px; /* AKA margin-right */
  margin-inline-end: 10px; /* AKA margin-left */
  block-size: 30px; /* AKA the height */
  inline-size: 40px; /* AKA the width */
}

.bar {
  inset-inline-start: 10px; /* AKA left: 10px */
  inset-inline-end: 20px; /* AKA right: 20px */
}
```

Why should you bother? **There are at least two big reasons**.

1. **Using logical properties are named after the box model**. This makes the code more consistent.

2. **Using logical properties gives you the support for different writing modes by default**. Some languages read left to right or in other direction.
  Luckily for us, the browser will handle other writing modes automatically for us if we use logical properties.

Even if you do not have a requirement for page translation now, consider how many people use the translate plugin available in the browsers. I know you had used it at least once!

You can [learn more about the logical properties here](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_logical_properties_and_values).

## The Many Kinds of Viewport Units

Gone are the days where we only had `vh` and `vw`. **Now we have multiple ways to define the viewport units. Most of them are driven by the mobile use-case where they are dynamic due to browser address bar appearing/disappearing**. There are **too many to list here, but the most interesting unit for me is the _dynamic_ one**. This one scales based on the state of the address bar (if it is hidden or not).

You can [learn more about the different kinds of viewport units here](https://www.terluinwebdesign.nl/en/css/incoming-20-new-css-viewport-units-svh-lvh-dvh-svw-lvw-dvw/).

## CSS Variables

CSS Variables allow you to declare different values, which you can use later in your CSS definitions. **Keep in mind that you can scope the visibility of variables, like in any other programming language**.

To declare a _global variable_, use the `:root` scope.

```html
<style>
  :root {
    --my-color: black;
  }

  button {
    background: var(--my-color);
  }
</style>
```

To declare a _scoped variable_, declare the variable inside a given scope.

```html
<style>
  :root {
    --my-color: black;
  }

  button {
    background: var(--my-color);
  }

  .page {
    --my-color: blue;
  }

  .page button {
    background: var(--my-color);
  }
</style>
```

It is possible to **declare a fallback for a given variable**. Think using`??` syntax in JavaScript.

```css
.page button {
  background: var(--i-do-not-exist, black);
}
```

### CSS Properties and type-safety

Experimental, at the time of writing this, syntax that allows you to **explicitly state the type and the inheritance semantics of a given property**. It is pretty magical. Check this out.

```css
@property --my-color {
  syntax: "<color>";
  inherits: false;
  initial-value: black;
}

button {
  background: var(--my-color);
}
```

Notice the **`syntax` property**. This example says that the `--my-color` variable can only contain values valid for _color-related_ properties. If that is not the case, **and you have used JavaScript to register the property**, the browser will throw an error.

In my opinion, the most compelling use-case for these is that **you can animate values of the properties in pure CSS!**. That is not the case with "regular" CSS properties.

> [Here is the MDN documentation regarding this topic](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Types)

### Computed CSS variables gotcha

Let us say you have the following CSS.

```css
/* This is equivalent to the HTML document, but it also applies to SVG elements "root" */
:root {
  --font-size: 1rem;
  --font-size-large: calc(2 * var(--font-size));
}

h1 {
  --font-size: 2rem;
  font-size: var(--font-size-large);
}
```

Would you expect the `font-size` to now be `4rem`? It certainly would make sense would not it? **The problem is that this will not be the case**.
The **calculation happens as soon as the browser processes the definition**. This means that the **computed values, in this case, in the `:root` are immutable and only inheritable**. This behavior is not specific to `:root`.

> [You can read more about this particular gotcha here](https://moderncss.dev/how-custom-property-values-are-computed/?ck_subscriber_id=1352906140#inheritable-values-become-immutable).

To make this code work, we would have to change where the computation happens.

```css
:root {
  --font-size: 1rem;
}

.font-resize {
  font-size: calc(var(--font-size-adjust, 1) * var(--font-size));
}

.font-large {
  --font-size-adjust: 2.5
}
```

Then, you **would use the `font-resize` and `font-large` classes on the same element**. This adherers to the rules of the cascade – we are not trying to update the parent, we are updating the "current element" styles.

## Cascade layers

Have you ever had problems with CSS selectors' specificity? In the end, most of us gave up and added the `!important` to the rule (or if it is evil, declare the property as a transition which will override the `!important`). You are not alone, and the web community has your back! – enter _cascade layers_.

```css
/* If I were to remove this line, all the buttons would be black. */
@layer declared_second, declared_first;

/* Due to the explicit ordering definition, the buttons are red. */
@layer declared_first {
  button {
    background: red;
  }
}

@layer declared_second {
  button {
    background: black;
  }
}
```

Another nice thing about cascade layers is that they are **declared only once** and **if declared again, they merge the inside content**.

```css
@layer declared_second, declared_first;
@layer declared_first {
  button {
    background: red;
  }
}

@layer declared_second {
  button {
    background: black;
  }
}

/* Now the buttons have a red background with a blue text color. */
@layer declared_first {
  button {
    color: blue;
  }
}
```

## Animating the height property

Since the beginning of CSS (?), we have been unable to animate the height of the element from "auto" to 0 (or vice-versa). That lead us to use various "hacks" to create toggles. With time, these "hacks" became the go-to solutions for creating a "hide/show" animation.

In an ideal world, we could animate the height using CSS, but that is not possible. If you want to animate the height, you shall use JS, and the best way to do this is to use the so-called FLIP technique.

### Using FLIP

With the FLIP technique, you first measure the element, apply the animations, and reverse them – all in a single frame. This allows you to create an effect of smooth animation. Do not get fooled by the "simplicity" of the steps listed here – to **correctly animate the height using FLIP, you will most likely need to apply reverse transforms**. You are better off using a library that will perform the math for you.

To learn more about the FLIP technique, checkout [_Method 4_ in this blog post](https://carlanderson.xyz/how-to-animate-on-height-auto)

### Using the `max-height` property

Instead of animating the `height`, we could animate the `max-height` property. This avoids the need to apply the scale transformations – the content will shrink (remember about the `overflow` property here!) as expected.

One note about the `max-height`: **you most likely want to measure the element you animate before you animate it**. The reason is that if you apply a very high `max-height` and then animate it to 0, the animation will look weird – the timing function applies to the whole `max-height` range!

Apart from the issues with timing functions, the `max-height` can sometimes mess with your layout. Be mindful of what you put the `max-height` on – we would not want the content to create unnecessary scroll bars!

### Using the CSS Grid layout

This one is entirely new for me. Instead of animating height-related properties, we could animate the grid tracks. Combine it with the `overflow: hidden` property, and you have a beautiful "collapse" animation.

```js
document.getElementById("app").innerHTML = `
<button type = "button" id = "toggle">Toggle</button>
<div class = "box">
  <div class = "box-inner">
    <p>content</p>
    <p>content</p>
    <p>content</p>
    <p>content</p>
  </div>
</div>
`;

const button = document.getElementById("toggle");
const box = document.querySelector(".box");

button.addEventListener("click", function onButtonClick() {
  const isHidden = box.classList.contains("hidden");
  if (isHidden) {
    box.classList.remove("hidden");
    return;
  }

  box.classList.add("hidden");
});
```

```css
.box {
  display: grid;
  grid-template-rows: 1fr;
  background: red;
  transition: grid-template-rows 0.5s ease-in-out;
}

.box-inner {
  overflow: hidden;
}

.hidden {
  grid-template-rows: 0fr;
}
```

## CSS-in-JS

CSS-in-JS became de-facto day of styling our apps. Let us explore how it works on the high level and learn about it's potential drawbacks as well as the benefits it brings to the table.

### Syntax

- You write your CSS, either via some kind of `css` function or via `styled.TAG_NAME`.

  - These were popularized by _emotion.js_ and _styled-components_ libraries.

- The big advantage here is that you can use React-declared variables to style the elements. This **makes the styles dynamic**.

  - Keep in mind that **you can pretty much do the same thing with CSS variables**.

### How does CSS-in-JS work?

> Based on [this article](https://www.lauchness.com/blog/emotion-under-the-hood)

- The styles you wrote are **serialized into CSS**.

  - For **static styles**, this could happen at build time or at runtime.

  - For **dynamic styles**, this happen at runtime, **when your component runs**.

  - **Serialization is costly**. It is the major performance hot-spot in many libraries.

- Then these styles are **injected into the HTML**. This also takes a bit of time.

  - CSS-in-JS libraries usually leverage catching so not to include the same definitions multiple times. The more granular the serialized CSS is, the less duplication.

### The benefits

- The ability to co-locate CSS and JSX in the same file

- Speed of development and DX. It is easy to pick up and learn.

### The drawbacks

> Learn more [by reading this article](https://dev.to/srmagura/why-were-breaking-up-wiht-css-in-js-4g9b)

- Performance issues due to serialization and runtime dependency.

- Increased JS bundle size.

- **Using CSS-in-JS library can clutter your React dev-tools**. Most libraries inject special components responsible for handling context (theme) and other stuff.

### The bottom line

CSS-in-JS is a great way to style your apps, but it is not without its shortcomings. With the advent of _server components_, **there might be a shift away from CSS-in-JS in favour of more native solutions like _CSS modules_ or _SaSS modules_**.

## CSS Container Queries

Remember setting the _media queries_ for the whole width of the page only to style some container that lived inside another container and it's width did not necessarily depended on the width of the viewport? You got the job done, but it was not something you enjoyed.

**Now you can base the size of the child out of the size of the parent**. If the parent shrinks, the child can also shrink. **The container queries are like media queries but instead of looking at viewport, they look at the parent**. Very useful stuff.

Why is this useful? **In the age of components, it is very hard to style them according to the viewport**. Components are composable, they might "appear" in different contexts. Ideally they should "just work" no matter who the parent is.

```css
.parent {
  container: NameForTheContainer / inline-size;
}

.child {
  color: red;

  @container NameForTheContainer (max-width: 400px) {
    color: blue;
  }
}
```

The syntax is very familiar. One addition is that you have to "name" your container to use the `@container` query.

## `display: contents` blows my mind

The `display: contents` will make it so that the **element will not generate any box, as such the children will be treated as if it did not exist in terms of layout**.

```html

<div style = "display: flex">
  <!-- As if this div \/ did not exist from the layout perspective -->
  <div style = "display: contents;">
  <!-- The children are subject to the flex algorithm -->
    <span>foo</span>
    <span>bar</span>
  </div>
</div>

```

- The colors and fonts inherit from the `display: contents` parent.

- The **padding, width and all box-related properties are ignored**.

TODO: <https://rachelandrew.co.uk/archives/2017/07/20/why-display-contents-is-not-css-grid-layout-subgrid/>

## The _lobotomized owl_ selector

- Its name from from how the selector "looks" when written: `* + *`.

- It **used to be more widely used since we did not have `gap` property at our disposal**.

  - Without the `gap` property, adding spacing between children was a bit tricky. If you were not careful, you could introduce the "leftover" spacing.

    - If you did, in most cases you also had to apply negative margins.

    ```html
    <style>
      p {
        margin-block-end: 1rem;
      }
    </style>
    <section>
      <p>foo</p>
      <p>foo</p>
      <p>I will have a "leftover" spacing at the bottom</p>
    </section>
    ```

    Contrast this with the following. The `margin-block-end` is only applies to the _second_ paragraph (in this particular case, using `margin-block-start` would be a better option).

    ```html
    <style>
      section p + p {
        margin-block-end: 1rem;
      }
    </style>
    <section>
      <p>foo</p>
      <p>foo</p>
      <p>I will NOT have a "leftover" spacing at the bottom</p>
    </section>
    ```

## The `color-mix` function

- There are multiple **color models or formats or spaces** (I've seen different wording used in different articles) now in CSS.

  - These define how "rich" the color is on different displays, and also how the colors mix using the `color-mix` function.

- The syntax looks like this: `color-mox(in oklab, white 30%, black)`

- The **`color-mix` will allow you to implement the "darken(X)" semantics for a given color**.

  - I used to have this `darken` function in my SCSS code back in the day.
