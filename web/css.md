# CSS

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

### `:not`

Well supported. The gotcha is that the **selector inside parenthesis must be
simple**. By simple we mean no combinators and spaces.

```css
element: not(img); /*ok*/
div: not(.someClass); /*can be also an id*/
div: not(ul li); /*will not work, there is a space, it's not a simple selector*/
```

### `:matches`

This one is wild. Look at the syntax

```css
element: matches(#home, .someClass, [title]);
```

Pretty neat huh? This one is almost like `querySelectorAll`. Of course this
would be too good to be true. **Support for this is quite meh**.

## Pseudo elements

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
