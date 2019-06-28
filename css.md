# CSS

## Selectors

### Specificity

Specificity is based on column-like structure (0-0-0).

- 1-0-0 for `#`
- 0-1-0 for `.`
- 0-0-1 for `element`

Whichever column has the bigger number in the left-most column wins. **If selectors are identical the bottom most wins (cascading styles)**

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
element[alt='some image description'] {
}
```

- value begins with something, **it ignores any dashes that come after it**. Mostly used on languages attributes.

```css
/*
    match all element that has alt which starts with: "some"
*/
element[alt|='some'] {
  /*
  <element lang = "some-us">
  <element lang = "some">
  */
}
```

- you can query by value which starts with (`^`) ends or (`$`) or has (`*`) some query

```css
element[alt$='some'],
element[alt*='some'],
element[alt^='some'] {
}
```

- you can also force case insensitivity using `i`
  When using `element[alt="some"]` that query inside quotes will match using case sensitivity by default. You can change that behavior

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
input[type='checkbox']:checked + label {
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

Matches element that has no content, is self closing or contains only a comment. **It cannot have any whitespace**.

#### `:blank`

Not really supported. Works like `:empty` but can contain whitespace.

### `:not`

Well supported. The gotcha is that the **selector inside parenthesis must be simple**. By simple we mean no combinators and spaces.

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

Pretty neat huh? This one is almost like `querySelectorAll`. Of course this would be too good to be true. **Support for this is quite meh**.
