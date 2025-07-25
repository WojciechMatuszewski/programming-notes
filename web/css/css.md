# CSS

## Centering the element

For the longest time, it was the case that, to center the element you had to change the `display` or `position` property of the element to center it.

- You would use the `position: absolute` with `translate` and `inset` (top/bottom)

- You would use `display: grid` with `place-items: center`.

- You would use `display: flex` with `place-items: center`.

At the time of writing this, **we can now center any block element using `align-items` and `justify-items`**.

This is great **because setting the `display` has implications for how margins collapse and other properties**. As such, while viable way of centering things, it always produces "side-effects" we might not want.

## The `height: 100%` does not work!

> [More information here](https://www.joshwcomeau.com/css/height-enigma/).

How many times have you written the following snippet of CSS only to realize it is not doing what you expect.

```css
.some-class {
  height: 100%;
}
```

Especially, in the case where you want to style some kind of "main" container

```html
<html>
  <head>
    <style>
      main {
        /* does not work! */
        height: 100%;
      }
    </style>
  </head>
  <body>
    <main>content</main>
  </body>
</html>
```

In retrospective, in almost all cases I can think of, it was the `height` that did not do what I was expecting it would do, and not the `width`. Why is that? **It has to do how browsers calculate the `height` and `width`**.

> [Check out this great video for a full explanation](https://youtu.be/Xt1Cw4qM3Ec?t=736)

- To **calculate the `width` browsers look at the parent of a given element**. This is recursive. The last parent is the `html` element that has the default width of the document.

- To **calculate the `height` browsers look at the children of a given element**. This "looking at the children" can create recursive conditions that render the `height: 100%` useless.

If we apply this logic to our example, we can see why the `height: 100%` is not working.

1. The `main` wants to be `height: 100%` of its parent.
2. The parent of `main` (the `body`) asks `main` how tall is it.
3. The `main` answers that it wants to be `height: 100%` of its parent.
4. And so on...

To break this recursive chain, one has to specify the height on the `html` and the `body` – in other worlds, **the parent heigh can't depend on the height of the children**.

```html
<html>
  <head>
    <style>
      html,
      body {
        /* Explicit height. No longer dependant on the height of the children */
        height: 100%;
      }

      main {
        /* works as expected */
        height: 100%;
      }
    </style>
  </head>
  <body>
    <main>content</main>
  </body>
</html>
```

Now, the `body` can ask the `html` about the `height` and answer the question from `main`. Remember that the `html` has the height of a document (implied height of the screen).

### About percentage values for height

Usually, for the `height: <NUMBER>%` to work, the parent's height has to use a concrete value, like `300px` or `10rem`. Think about it: without a "concrete" number, how would the browser be able to calculate a percentage of that number?

If that is the case, how come the `height: 100%` on the `main` works, if the `body` and `html` both have a `height: 100%` applied to them?

[See this section](https://www.joshwcomeau.com/css/height-enigma/#percentages-all-the-way-down-3). It boils down to the fact that **the `html` tag has an "implied concrete height" that is equal to the viewport!**

So, you might often see the following CSS in older projects:

```css
html,
body,
#root {
  height: 100%;
}
```

Nowadays, this is not necessary given that we have the `vh` and `svh` units available to us.

## Spacing

Instead of trying to memorize the order of the sides in the `padding` or `margin` properties, use **`padding-inline` (
left-right) and `padding-block` (top-down)**. The same applies to margins. One has to swap the `padding` for margin.

There is an additional benefit to using these properties instead of the "regular" `margin` and `padding` properties. \*
\*They take the `writing-mode`, `direction`, and `text-orientation` property into the account\*\*, which makes your code
more reusable.

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

_An adjacent sibling_. Pretty self-explanatory.

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
element: matches (#home, .someClass, [title]);
```

Pretty neat huh? This one is almost like `querySelectorAll`. Of course this
would be too good to be true. **Support for this is quite meh**.

## Pseudo-classes

### The `:is` and `:where`

The `:is` and `:where` selectors are used to eliminating repetition that sometimes you have to deal with when selecting
multiple children of a given element.

```css
:is(article) h2,
h3,
h4,
h5,
h6 {
  color: var(--fire-red);
}

/ / the same as \/ article h2,
article h3,
article h4,
article h5,
article h6 {
  color: var(--fire-red);
}
```

The only difference between the `:is` and `:where` selectors is the specificity.

The **`:where` selector has a specificity of 0, whereas the `:is` takes the specificity of the most specific argument**.

### The `:has`

The `:has` selector enables you to **style the parent based on it's the children**.

This is a groundbreaking change in how we think about CSS. Since always, the CSS rules had to obey the _cascading_ semantics – you could not "go back" to the parent.

The following is an example of the `:has` selector where we specify the number of columns based on the list items.

```css
ul:has(li:nth-child(6)) {
  columns: 2;
}

ul:has(li:nth-child(11)) {
  columns: 3;
}
```

---

In addition, **you can use the `:has` as the so-called _anywhere_ selector**. The **name stems from the usage of `:has` on the `body` tag**.

```css
body:has(input.blur-answer:checked) .answer {
  filter: blur(5px);
}
```

```css
/* Given certain element with a focus inside the card, style the card */
.card:has(:where([data-action="delete"], [data-action="open"]):focus-visible) {
  border-top: 10px solid #f7bccb;
  box-shadow: 0 0 0 2px #f7bccb;
  transform: scale(1.02);
}

/* If there is a card with a focus within, style all other cards */
body:has(.card:focus-within) .card:not(:focus-within) {
  filter: grayscale();
  background-color: #f6f7f6;
}
```

---

If you are well-versed with CSS selectors, **you can achieve behavior that seemingly defies the cascade – apply style "
backwards"**.

```css
ul li:has(+ .select-before) {
  /* styles */
}
```

This will **element that occurs BEFORE the `.select-before`**. Let that sink in... we are styling backwards! This technique is quite useful for styling labels if the input value is invalid.

```css
label:has(+ input:user-invalid) {
  /* some styles */
}
```

---

You can even produce layouts that have **different styles based on the amount of children a given container has**.

```css
.container:has(> *:nth-child(10)) {
  /* This only selects the container if it has 10 or more children */
}
```

How cool is that??

> Side note: instead of `:invalid` you most likely want to use the new `:user-invalid`. [Read more about it on MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/:user-invalid).

---

What about styling **all the siblings but not the element we are interacting with**? Hell yes!

```css
.card-list:has(.card:hover) .card:not(:hover) {
  scale: 0.9;
}
```

Wave your goodbyes to using JavaScript for this!

#### `:has` and inputs

The `:has` selector is very powerful when combined with inputs. Check out this example with checkbox.

```html
<style>
  .notice {
    display: none;
  }

  .body:has(#toggle:checked) .notice {
    display: block;
  }
</style>

<input type="checkbox" id="toggle" />
<div class="notice">Some notice</div>
```

Quite powerful if you ask me. And you are not limited to checkboxes. Keep in mind that other input types also have "
internal state" (the `text` has the value, the `file` has the file name).

### `:nth-child`

> Based on [this great blog post](https://developer.chrome.com/docs/css-ui/css-nth-child-of-s)

When learning about this pseudo-class you most likely encountered definitions like the following.

```css
:nth-child(2) /* second child */
:nth-child(2n) /* all even children */
:nth-child(6n + 1) /* start with n=0 and increment it */
```

In most cases, the number syntax inside the brackets is sufficient to get things done.

**But, did you know you can apply "filters" inside the brackets of the `:nth-child` selector?**

Let us consider the following selector.

```css
:nth-child(2 of .some-special-class);
```

This will select the **second child with `.some-special-class`**. The kicker is that **the child does not necessarily have to be the second child of the parent**. This is in stark difference with the "regular" `:nth-child` selector.

```css
:nth-child(-n + 3 of li.important); /*First three list items that have the `.important` class*/

li.important: nth-child(-n + 3); /*If they are the first three children AND match the selector `li.important`*/
```

Consider the following, perhaps invalid, HTML structure.

```html
<ul>
  <li>First</li>
  <span>some span</span>
  <li>Second</li>
  <li>Third</li>
  <span>some span</span>
  <li>Fourth</li>
  <li>Fifth</li>
</ul>
```

And the following CSS.

```css
li:nth-child(-n + 3) {
  border: 1px solid red;
}

ul :nth-child(-n + 3 of li) {
  border: 1px solid red;
}
```

The `li:nth-child(-n + 3)` will put a border on the "First" and "Second" `li`. **The "Third" `li` will not have the right styles applied**. The `li:nth-child(-n + 3)` means **style the `li` if its the first, second or third position within the parent**. Since we have a `span` in-between the first, and the second `li`, the third `li` counts as a fourth element, therefore the style is not applied.

The `ul :nth-child(-n + 3 of li)` will **apply the border to the "First", "Second" and "Third" `li` elements**. It no longer looks at the position of the elements, but rather the number of occurrences within the parent the selector matches.

### `:focus-visible`

> [Read more about this selector here](https://fullystacked.net/programmatic-focus-styles/)

The `:focus-visible` uses browser UA heuristics to determine when to display the focus outline.

**This is not the case with the "regular" `:focus` pseudo-class** which is "dumb" in that regard. You **might have needed this to disable the focus outline on link clicks**.

In fact, most (if not all) of the browsers, migrated from `:focus` to `:focus-visible` in their UA styles.

Did you notice that, if you click a link, the focus outline is not there? **But if you use the keyboard, the focus outline is visible?**. This is the `:focus-visible` pseudo-class in action.

#### Programmatic focus

TIL that the `.focus` method has options you can pass. One of the options is `focusVisible`. If you specify its value to `true`, the browser is supposed to apply the `focus-visible` styles to the element.

There is also the `preventScroll` parameter which sounds quite useful!

### `:not`

Well-supported. The gotcha is that the **selector inside parenthesis must be simple**. By simple we mean no combinators and spaces.

```css
element: not (img); /*ok*/
div: not (.someClass); /*can be also an id*/
div: not (ul li); /*will not work, there is a space, it's not a simple selector*/
```

#### Using the `:not` with elements that DO NOT contain certain properties

What if you wanted to select all `p` tags that DO NOT have any class? Or maybe an image that does NOT have the `alt`
tag?

```css
img:not([alt]) {
  outline: 5px solid red;
}
```

I find it pretty amazing that we can do so much logic in CSS. Of course I was not aware that using `[]` syntax checks
for the _presence_ of a given attribute.

#### Elements that are NOT descendants of particular element

When learning about the `:not` selector, I always thought about it in terms of applying the "not" to the "current element" I have selected.

```css
/* Select `p` tags that do not have `.foo` class */
p:not(.foo) {
}
```

Now, **it turns out, you can also provide more complex selectors to `not`**.

```css
/* Select all inputs that are NOT the descendants of a form element. */
input:not(form input) {
}
```

Notice that, in this case, the selector in the `not` does not "apply" to the "current element" directly. It is as if the `not` selector allowed you to "reach up" in a DOM tree in a way.

#### Combining `:not` with `:has`

> Based on this [this blog post](https://www.matuzo.at/blog/2022/100daysof-day50/).

Consider the following

```css
/* First selector */
.card:not(:has(img)) {
}

/* Second selector */
.card:has(:not(img)) {
}
```

The selectors are clearly different, but what would be the outcome of each selector? The same? or different?

**When debugging CSS selectors, consider unraveling them from "inside out"**. If we apply this methodology here, we get:

1. For the first selector, we are trying to select a `.card` that contains an image and "inverse" the check.

2. For the second selector, we are selecting any element that is not an image and checking if the `.card` contains that element.

See how they differ?

1. The first selector selects all `.card`s that DO NOT have an image.

```html
<!-- Would not match. There is an img element inside the card. -->
<div class = "card">
  <span>foo<span>
  <img/>
</div>


<!-- Would match. There is NOT img element inside the card. -->
<div class = "card">
  <span>foo<span>
</div>
```

2. The second selector selects all `.card`s that have elements DIFFERENT than an image.

```html
<!-- Would match. There is a span element. That element is NOT an image. -->
<div class = "card">
  <span>foo<span>
  <img/>
</div>


<!-- Would match. There is a span element. That element is NOT an image. -->
<div class = "card">
  <span>foo<span>
</div>
```

**Be mindful when combining pseudo-classes**. Changing their order will impact how your selector result.

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

Self-explanatory. You can style placeholders 🤷‍

## Progressive enhancement

Let's say you want to use CSS grid where possible and for older browsers some
simple, other, layout.

How would we go about detecting if we can use grid?

Well, it turns out there is are so-called **CSS Feature Queries**.

These can be used to detect if a CSS feature is available on given browser.

```css
@supports (display: grid) {
// you code here
}
```

You can also check if the browser supports given selectors

```css
@supports selector(p:has(a)) {
}
```

## Logical properties

Instead of using the `left` and `right` keyword for paddings, margins and so on, one **should consider using
the `inline` and `block` keywords**.

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

2. **Using logical properties gives you the support for different writing modes by default**. Some languages read left
   to right or in other direction.
   Luckily for us, the browser will handle other writing modes automatically for us if we use logical properties.

Even if you do not have a requirement for page translation now, consider how many people use the translate plugin
available in the browsers. I know you had used it at least once!

You can [learn more about the logical properties here](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_logical_properties_and_values).

### Logical properties values gotchas

So you started adapting the _logical properties_. Nice!

The other day, you wanted to create a container that spans 50% of the viewport width. Naturally, you do something like the following:

```css
.container {
  inline-size: 50vw;
}
```

Without a second thought, you push the code to production – we are done right? **Sadly, you just nullified the benefit of _logical properties_ by using `50vw`**.

The **`vw` unit is not "responsive" to the changes in writing direction**. You ought to use a _logical properties_ equivalent of `vw` (and `vh` for that matter).

Enter the `vi` and `vb` units.

- The `vi` is for the _inline direction_.
- The `vb` is for the _block direction_.

So, the previous snippet should be rewritten to the following

```css
.container {
  inline-size: 50vi;
}
```

[You can learn more about those units here](https://www.stefanjudis.com/today-i-learned/viewport-units-can-consider-the-writing-mode).

## The Many Kinds of Viewport Units

Gone are the days when we only had `vh` and `vw`. **Now we have multiple ways to define the viewport units. Most of
them are driven by the mobile use-case where they are dynamic due to browser address bar appearing/disappearing**. There
are **too many to list here, but the most interesting unit for me is the _dynamic_ one**. This one scales based on the
state of the address bar (if it is hidden or not).

You can [learn more about the different kinds of viewport units here](https://www.terluinwebdesign.nl/en/css/incoming-20-new-css-viewport-units-svh-lvh-dvh-svw-lvw-dvw/).

## The `ch` unit

> Based on [this great blog post](https://clagnut.com/blog/2432/).

1. **The `ch` unit does not correspond to the "count" of _characters_ in a given line**.

2. The **`ch` unit represents the width of the `0` character within the selected font OR 0.5 rem is no `0` is present**.

The implications of those points are quite large.

- Unless the you have font consisting of only `0`, the `width: 66ch` will not equal to 66 characters.

- The resulting width of `width: 66ch` declaration is **purely font-dependant**.

It is not "set if and forget it" kind of deal. You **might consider using rems instead**.

## Relative color syntax

> [Read this excellent blog post to learn more](https://ishadeed.com/article/css-relative-colors/).

Gone are the days where you had to spend X amount of time manually deriving colors from a "brand" color or similar. Now, we have a built-in way to do this: enter [_CSS relative colors_](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_colors/Relative_colors)

```css
:root {
  --brand: red;
  --brand-light: hsl(from var(--brand) h s calc(l + 30));
}
```

Notice how powerful this is!

1. I can mix and match colors defined via different "color functions".

2. I can use `calc` inside those functions.

### The `hsl` color space

The `rgb` or `hex` representations of a color does not tell us much. [This blog post explains the problem quite well](https://css-tricks.com/almanac/functions/h/hsl/#aa-why-hsl).

Consider the following colors:

```css
rgb(237, 70, 44)

rgb(255, 120, 100)

rgb(200, 50, 30)
```

The numbers there are not really that meaningful right? I mean, you might have an intuition for them, but I certainly do not know what color they represent.

Now, consider HSL:

```css
hsl(8, 84%, 55%)
hsl(8, 84%, 70%)
hsl(8, 84%, 40%)
```

If you remember what HSL stands for – `h` for _hue_, `s` for _saturation_, and `l` for _lightness_ – you might be able to guess that all of these represent some "reddish" colors.

I personally find the HSL representation of colors much more logical to use.

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

> [Here is the MDN documentation regarding this topic](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Types)

> [Here is a great blog post](https://moderncss.dev/providing-type-definitions-for-css-with-at-property) about this topic.

Experimental, at the time of writing this, syntax that allows you to **explicitly state the type and the inheritance semantics of a given property**.
It is pretty magical. Check this out.

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

Notice the **`syntax` property**. This example says that the `--my-color` variable can only contain values valid for _color-related_ properties.

If that is not the case, **and you have used JavaScript to register the property**, the browser will throw an error.

In my opinion, the most compelling use-case for these is that **you can animate values of the properties in pure CSS!** – that is not the case with "regular" CSS properties.

#### Improved debugging experience

If you ever found yourself debugging CSS variables, you know how hard it can get. **The nice thing about `@property`-defined variables is that you get much more information in browser devtools**.

Let us say you provided a wrong _type_ to the variable via JS. If you open the devtools console, you will see a warning near the property value with an explanation on WHY you see this warning in the first place.

#### "Literal values" syntax

Apart from using predefined values like `<color>` or `<length>` **you can actually provide literal values**!

```css
@property --my-color {
  syntax: "blue | cyan";
  inherits: false;
  initial-value: black;
}

.section {
  --my-color: cyan;
}
```

You can even go further and provide "mixed" values for the `syntax` property. The `<color> | <length>` will also work!

#### Note about the `inherits` property

**The `inherits` property is required**. If you fail to include it in the definition, the whole `@property` block will be invalid.

To illustrate how it works, consider the following example:

```css
@property --brand {
  syntax: "<color>";
  /* Notice the "false" value here! */
  inherits: false;
  initial-value: blue;
}

.button {
  --brand: red;
  background-color: var(--brand);
  color: white;

  span {
    display: block;
    width: 30px;
    height: 30px;
    border: 1px solid var(--brand);
  }
}
```

What would be the border color of the `span` inside the `button` class?

**If you set the `inherits` to `false`, the color would be `blue`** (initial value of the `--brand` variable). This is because the change of the `--brand` variable in `.button` class will NOT propagate to the `span` tag.

**If you set the `inherits` to `true`, the color would be `red`**. This is because the change of the `--brand` variable WILL propagate to the `span`.

Take a moment to appreciate how much control this gives you. At the time of writing this, I can't think of another API that would allow me to control inheritance in such a way.

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

Would you expect the `font-size` to now be `4rem`? It certainly would make sense would not it? **That will NOT be the case**.

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
  --font-size-adjust: 2.5;
}
```

Then, you **would use the `font-resize` and `font-large` classes on the same element**. This adherers to the rules of
the cascade – we are not trying to update the parent, we are updating the "current element" styles.

### Conditional CSS

Would it not be awesome to have the ability to apply some CSS styles conditionally? Yes, one could use JavaScript to do
this, but I think there is something to be said about doing all of that in CSS. Note that I'm not talking about
conditions based on the viewport. I'm talking about the conditions based on the "state" of the rules themselves!

It turns out that it is possible, and, at the time of writing this, there are at least two ways I can think of that
allow us to apply conditions in CSS.

#### The "space toggle"

> Based on resources [from this GH repo](https://github.com/propjockey/css-sweeper?tab=readme-ov-file#basics-of-space-toggle).

> Also, check [this article out](https://www.bram.us/2023/09/16/solved-by-css-scroll-driven-animations-detect-if-an-element-can-scroll-or-not).

This works, because **an empty space is a valid CSS custom-property value**. If **the space value essentially acts as `true` value in other programming languages**.

```css
.box {
  /* notice the space here! */
  --toggler: ;
  /* read this as "if toggler then .." */
  --red-if-toggler: var(--toggler) red;

  /* the background will be red */
  background: var(--red-if-toggler, blue);
  width: 50px;
  height: 50px;
}
```

Okay, so if empty space acts as a "true", how do we represent the `false` value? To **represent the `false` use the `initial` keyword**.

```css
.box {
  --toggler: initial;
  /* read this as "if toggler then .." */
  --red-if-toggler: var(--toggler) red;

  /* the background will be blue */
  background: var(--red-if-toggler, blue);
  width: 50px;
  height: 50px;
}
```

You can toggle the "toggler" using _media queries_.

```css
.box {
  --toggler: initial;
  --red-if-toggler: var(--toggler) red;
  background: var(--red-if-toggler, blue);
  width: 50px;
  height: 50px;
}

@media (max-width: 600px) {
  .box {
    --toggler: ;
  }
}
```

---

And here is the usage with the `hover` pseudo-selector;

```css
:root {
  /*Remember – empty space acts as a "true" value*/
  --hover-false: ;
  --hover-true: initial;
}

a {
  background: var(--hover-false, red) var(--hover-true, blue);
}

a:hover {
  --hover-false: initial;
  --hover-true: ;
}
```

The space toggle **the biggest benefit is that it works cross-browser**. This behavior is well-defined in the CSS spec.
The **biggest drawback is that it looks like a hack and could be hard to understand for others**.

Use it wisely!

I first learned about the "space toggle" while reading [this article](https://www.bram.us/2023/09/16/solved-by-css-scroll-driven-animations-detect-if-an-element-can-scroll-or-not).

**It showcases how to detect if the element is scrollable via CSS only**. Amazing technique!

#### The style query

At the time of writing, this only works in newest Chrome-based browsers.

```css
.container {
  container-type: inline-size;
  container-name: box-container;
  --toggle: 0;
}

@media (max-width: 600px) {
  .container {
    --toggle: 1;
  }
}

.box {
  background: var(--bg);
  width: 50px;
  height: 50px;
}

@container box-container style(--toggle: 1) {
  .box {
    --bg: red;
  }
}

@container box-container style(--toggle: 0) {
  .box {
    --bg: blue;
  }
}
```

Using **_style queries_ is, at least for me, much more readable**. There is no "magic" with the empty space, but the
syntax is much more verbose.

## Cascade layers

> [Learn more here](https://youtu.be/D79TND9w_AY?list=PLuZ_sYdawLiWenx-X315dfZNOaliVnSTY&t=539).

Have you ever had problems with CSS selectors' specificity? In the end, most of us gave up and added the `!important` to
the rule (or if it is evil, declare the property as a transition which will override the `!important`). You are not
alone, and the web community has your back! – enter _cascade layers_.

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

You **can import styles directly _into_ a layer**.

```css
@import "theme.css" layer(utilities);
```

## Inheritance and "proximity"

> Based on [this blog post](https://jwdallas.com/posts/nesteddarkmode/).

Consider the following snippet.

```html
<style>
  [data-theme="red"] a {
    color: red;
  }

  [data-theme="blue"] a {
    color: blue;
  }
</style>

<div data-theme="blue">
  <a>Should be blue</a>
  <div data-theme="red">
    <a>Should be red (but is actually blue)</a>
  </div>
</div>
```

CSS compares two selectors in isolation. Since both selectors have the same specificity, the one that was last declared
wins (due to _cascading_ nature of CSS).

How could we make these styles work? **We can achieve the desired end-result via inheritance**.

```html
<style>
  [data-theme="red"] {
    color: red;
  }

  [data-theme="blue"] {
    color: blue;
  }
</style>

<div data-theme="blue">
  <a>Is blue</a>
  <div data-theme="red">
    <a>Is red</a>
  </div>
</div>
```

Notice that we do not target any specific tag. **If two or more ancestor elements have set values, the child element will always use the value from its closest parent**.

While this might not be super helpful for the `color`, you should know that **custom properties also are subject of inheritance**.

```html
<style>
  [data-theme="red"] {
    --background-color: red;
  }

  [data-theme="blue"] {
    --background-color: blue;
  }

  a {
    background-color: var(--background-color);
  }
</style>

<div data-theme="blue">
  <a>Is blue</a>
  <div data-theme="red">
    <a>Is red</a>
  </div>
</div>
```

Now we are talking are we not? **This could be used for theming specific sections of a given page**!

## The usefulness of `overflow: clip`

> Based on [this great blog post](https://ishadeed.com/article/overflow-clip/)

**The main issue with `overflow:hidden` is that if you specify it for one axis, the other axis becomes `overflow:auto`**.

```css
.foo {
  overflow-y: hidden;
  /* The X axis is now implicitly "auto" */
}
```

This means that **the other axis will most likely show a scrollbar and actually clip both axis**.

The **benefit of `overflow: clip` is the ability to specify which axis we want to clip without affecting the other axis**.

In addition, **you can control the "boundary" of the clipped area**.

```css
.foo {
  overflow-x: clip;
  overflow-clip-margin: 40px;
}
```

## Animating the height `auto` property (or any other size-related property)

Since the beginning of CSS (?), we have been unable to animate the height of the element from "auto" to 0 (or
vice-versa). That lead us to use various "hacks" to create toggles. With time, these "hacks" became the go-to solutions
for creating a "hide/show" animation.

In an ideal world, we could animate the height using CSS, but that is not possible. If you want to animate the height,
you shall use JS, and the best way to do this is to use the so-called FLIP technique.

### Using FLIP

With the FLIP technique, you first measure the element, apply the animations, and reverse them – all in a single frame.
This allows you to create an effect of smooth animation. Do not get fooled by the "simplicity" of the steps listed
here – to **correctly animate the height using FLIP, you will most likely need to apply reverse transforms**. You are
better off using a library that will perform the math for you.

To learn more about the FLIP technique, checkout [_Method4_ in this blog post](https://carlanderson.xyz/how-to-animate-on-height-auto)

### Using the `max-height` property

Instead of animating the `height`, we could animate the `max-height` property. This avoids the need to apply the scale
transformations – the content will shrink (remember about the `overflow` property here!) as expected.

One note about the `max-height`: **you most likely want to measure the element you animate before you animate it**. The
reason is that if you apply a very high `max-height` and then animate it to 0, the animation will look weird – the
timing function applies to the whole `max-height` range!

Apart from the issues with timing functions, the `max-height` can sometimes mess with your layout. Be mindful of what
you put the `max-height` on – we would not want the content to create unnecessary scroll bars!

### Using the CSS Grid layout

This one is entirely new for me. Instead of animating height-related properties, we could animate the grid tracks.
Combine it with the `overflow: hidden` property, and you have a beautiful "collapse" animation.

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

This technique is used in [this free course](https://www.epicweb.dev/tutorials/fluid-hover-cards-with-tailwind-css).

Looking at the performance snapshot from Chrome, animating height this way will cause **layout and style recalculation** but **it seems to be much cheaper than reflow**.

### Using the `calc-size`

> Based on [this article](https://developer.chrome.com/docs/css-ui/animate-to-height-auto/)

**The API of the `calc-size` is mind-blowing**.

Relatively new addition to the CSS, it allows you to animate from `0` to `max-height` or any other other "auto" property. **Please note that this is not the same as animating from `display: none` to `display: block`** – for that we have **`@starting-style`**.

```css
.animate {
  width: 30px;
  overflow: clip;
  display: block;
  white-space: nowrap;
  transition: width 0.3s ease;

  &:hover {
    width: calc-size(max-content, size);
  }
}
```

When you hover over the element, it will animate from `30px` to `max-content`. No `grid` or FLIP needed.

**When the `calc-size` really shines, is the ability to perform operations on the `size` keyword**. Check this out.

```css
.animate {
  width: 30px;
  overflow: clip;
  display: block;
  white-space: nowrap;
  transition: width 0.3s ease;

  &:hover {
    width: calc-size(max-content, size / 2);
  }
}
```

When you hover over the element, it will animate from `30px` to **half of the `max-content`**. How amazing is that?

**Combine this with [`round` CSS function](https://developer.mozilla.org/en-US/docs/Web/CSS/round) and the sky is the limit**.

```css
p {
  width: calc-size(max-content, round(up, size, 50px));
}
```

The `p` all have the `width` of `max-content` but rounded to the nearest multiply of `50`. **Again, how amazing is that?**

### Using `interpolate-size`

> [Read more about this property here](https://developer.mozilla.org/en-US/docs/Web/CSS/interpolate-size)

This property allows you **to enable native animations between length/percentage values and _intrinsic sizes_ such as `auto` or `fit-content`**. At the time of writing, this property is not YET widely supported, but it is only a matter of time when it will be.

You _might_ consider adding `interpolate-size: allow-keywords;` to your CSS reset, but be mindful of users who prefer the "reduced motion" experience.

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <style>
      #box {
        /* From "auto" size */
        interpolate-size: allow-keywords;
        background: red;

        transition:
          block-size 1s,
          inline-size 1s;

        &:hover {
          /* To a well-defined size */
          block-size: 80px;
          inline-size: 80px;
        }
      }
    </style>
  </head>
  <body>
    <div id="box">box</div>
  </body>
</html>
```

## CSS-in-JS

CSS-in-JS became de-facto day of styling our apps. Let us explore how it works on the high level and learn about it's
potential drawbacks as well as the benefits it brings to the table.

### Syntax

- You write your CSS, either via some kind of `css` function or via `styled.TAG_NAME`.

  - These were popularized by _emotion.js_ and _styled-components_ libraries.

- The big advantage here is that you can use React-declared variables to style the elements. This **makes the styles
  dynamic**.

  - Keep in mind that **you can pretty much do the same thing with CSS variables**.

### How does CSS-in-JS work?

> Based on [this article](https://www.lauchness.com/blog/emotion-under-the-hood)

- The styles you wrote are **serialized into CSS**.

  - For **static styles**, this could happen at build time or at runtime.

  - For **dynamic styles**, this happen at runtime, **when your component runs**.

  - **Serialization is costly**. It is the major performance hot-spot in many libraries.

- Then these styles are **injected into the HTML**. This also takes a bit of time.

  - CSS-in-JS libraries usually leverage catching so not to include the same definitions multiple times. The more
    granular the serialized CSS is, the less duplication.

### The benefits

- The ability to co-locate CSS and JSX in the same file

- Speed of development and DX. It is easy to pick up and learn.

### The drawbacks

> Learn more [by reading this article](https://dev.to/srmagura/why-were-breaking-up-wiht-css-in-js-4g9b)

- Performance issues due to serialization and runtime dependency.

- Increased JS bundle size.

- **Using CSS-in-JS library can clutter your React dev-tools**. Most libraries inject special components responsible for
  handling context (theme) and other stuff.

### The bottom line

CSS-in-JS is a great way to style your apps, but it is not without its shortcomings. With the advent of _server
components_, **there might be a shift away from CSS-in-JS in favour of more native solutions like _CSS modules_
or _SaSS modules_**.

## CSS Container Queries

> More information [in this great blog post](https://www.joshwcomeau.com/css/container-queries-introduction)

Remember setting the _media queries_ for the whole width of the page only to style some container that lived inside
another container and it's width did not necessarily depended on the width of the viewport? You got the job done, but it
was not something you enjoyed.

**Now you can base the size of the child out of the size of the parent**. If the parent shrinks, the child can also
shrink. **The container queries are like media queries, but instead of looking at viewport, they look at the parent**.

Handy stuff.

Why is this useful? **In the age of components, it is hard to style them according to the viewport**. Components
are composable, they might "appear" in different contexts. Ideally, they should "just work" no matter who the parent is.

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

### Relationship with _CSS containment_

What is this `inline-size` that every code snippet has? Why do we even need to specify it?

Think about the case **where using container query changes the height of the element**.

```css
.foo {
  @container (max-width: 2rem) {
    font-size: 30rem;
  }
}
```

We have just created an infinite loop. When you shrink `.foo` you will bump its `font-size`. This will make `foo` bigger, removing the bigger `font-size`. This makes the `.foo` smaller, adding the bigger `font-size`...

The **`inline-size` or `size`** is part of [CSS containment](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_containment) spec. It **instructs the element to skip retrieving information about its size (or inline-size) from the children it renders**.

This prevents the infinite loop from happening!

#### Why `inline-size`?

Armed with the knowledge about _CSS containment_ and this _infinite loop_ problem, let us consider the `inline-size` value of `container` property.

**If you change the container width, it is very probable that the container will change its height**. Adding/removing children also will, most likely, affect the height.

If that is the case, using `container: inline-size` makes sense. **We want the container to have the width dictated by its parent (or CSS), but react to `height` changes based on the children**.

Notice that there is no infinite loop to be had here. Height information comes from the children. Width information comes from the parent (or CSS).

### Pitfalls

Every API/feature has pitfalls. This one is no different.

1. **The container has to have an explicit size or come from `flex` or `grid` layout and NOT from its children**.

```css
.nav {
  container: nav / inline-size;
  display: flex;
  gap: 1rem;
}
```

In the code above, the `width` of the container will have an effective width of zero. Use `flex: 1` to fix that problem.

2. You **cannot query the container against itself**.

```css
.nav {
  container: nav / inline-size;
  display: flex;
  gap: 1rem;
}

@container nav (min-width: 700px) {
  /*This will not work!*/
  background: red;
}
```

## `display: contents` blows my mind

The `display: contents` will make it so that the **element will not generate any box, as such the children will be treated as if it did not exist in terms of layout**.

```html
<div style="display: flex">
  <!-- As if this div \/ did not exist from the layout perspective -->
  <div style="display: contents;">
    <!-- The children are subject to the flex algorithm -->
    <span>foo</span>
    <span>bar</span>
  </div>
</div>
```

- The colors and fonts inherit from the `display: contents` parent.

- The **padding, width and all box-related properties are ignored**.

Think of the `display: contents` **as a "ungrouping" mechanism**.

### But what about `subgrid`?

While using the `display: contents` might be useful in some other situations, the first time I've learned about this
property, my mind immediately started comparing how the `display: contents` works [with the `subgrid` property](https://web.dev/css-subgrid).

**The main difference is that, with `subgrid` you preserve the dimensions of the parent**. This means that you have more control over the parents dimensions.

The `display: contents` makes the box "invisible" to the HTML, so the children will be at the mercy of the grandparent dimensions.
Also, using the `display: contents` to make `subgrid` work is a workaround. While the layout might work, the intent behind the code is kind of lost.

Where `subgrid` shines is the ability to have consistent heights of rows across different elements, like cards.
The **children can use the rows from the parent, and the columns from the grandparent, or some combination of both**.

[Here is a good video about this topic](https://www.youtube.com/watch?v=rmF_iE0lMzw).

[And here is a demo with a form, where all inputs "start" at the same width](https://codepen.io/chriscoyier/pen/YzxqJap).

## The _lobotomized owl_ selector

- Its name from how the selector "looks" when written: `* + *`.

- It **used to be more widely used since we did not have `gap` property at our disposal**.

  - Without the `gap` property, adding spacing between children was a bit tricky. If you were not careful, you could
    introduce the "leftover" spacing.

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

    Contrast this with the following. The `margin-block-end` is only applies to the _second_ paragraph (in this
    particular case, using `margin-block-start` would be a better option).

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

### Combining with `has` selector

> Based on [this great blog post](https://frontendmasters.com/blog/selecting-previous-siblings)

The _lobotomized owl_ selector allows you to select N _next_ siblings to a given elements

```css
.card {
  /* Select the second sibling to the "current" card. */
  & + * + * {
  }
}
```

But what if you want to select N _previous_ elements? **You can do this with `has` selector**!

```css
.card {
  /* Select the second sibling to the "current" card */
  & + * + * {
  }

  /* Select the second "previous" element to the current card */
  :has(* + * + &) {
  }
}
```

**Notice the usage of `&` in the `:has` selector above**. The "last" element has to point to that "current" card, so it has to be `&`.

## Customize the `select` tag with CSS

For the longest time, the `select` element was not very useful apart from implementing very basic UIs. The main problem was that you could not style it that well. If you wanted a good-looking UI for the `select` element, you had to lean on 3rd-party packages, or implement custom `select` behavior yourself.

I'm happy to say that this is no longer the case! Just recently, Chrome released [a way for us to style the native `select` element](https://developer.chrome.com/blog/a-customizable-select).

**Be mindful of what behavior you _lose_ when you opt into `appearance: base-select;`**. I wish we did not lose the built-in mobile operating-system "select" component.

## The `color-mix` function

- There are multiple **color models or formats or spaces** (I've seen different wording used in different articles) now
  in CSS.

  - These define how "rich" the color is on different displays, and also how the colors mix using the `color-mix`
    function.

- The syntax looks like this: `color-mox(in oklab, white 30%, black)`

- The **`color-mix` will allow you to implement the "darken(X)" semantics for a given color**.

  - I used to have this `darken` function in my SCSS code back in the day.

## The `tabular-nums`

- Use the `font-variant-numeric: tabular-nums;` every where you do not want the content shifting when some number
  changes. I bet you encountered a situation where changing from `10` to `12` caused a slight layout shift. If you were
  using the `tabular-nums` that would not be the case.

## CSS Grid and custom names

- At the time of writing, there are two ways to create custom names when defining grid layouts.

  1.  Naming the lines.

  2.  Naming the areas.

  **Naming the lines is different than naming the areas**. The syntax is different and, overall, the end-result is a bit
  different.

### Naming the lines

Below is an example of naming the lines and then using them.

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>HTML + CSS</title>
    <link rel="stylesheet" href="styles.css" />
    <style>
      .grid-container {
        display: grid;
        grid-template-columns: 1fr [content-start] 1fr [content-end] 1fr;
      }

      .main-content {
        /* Notice the name here */
        grid-area: content;
      }
    </style>
  </head>
  <body>
    <div class="grid-container">
      <div class="main-content">
        <p>...</p>
        <p>...</p>
        <p>...</p>
        <p>...</p>
        <p>...</p>
      </div>
    </div>
  </body>
</html>
```

**Notice that I did not specify the `-start` and `-end` within the `grid-area` definition**. [According to the MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_grid_layout/Grid_layout_using_named_grid_lines) this is a built-in feature. I was not aware of this!

> While you can choose any name, if you append -start and -end to the lines around an area, as I have in the example above, grid will create you a named area of the main name used.

One important detail here is to understand that the `content` area lives between the first and the third column. I'm naming the lines here, not the areas!

#### "Overloading" line names

You can have multiple names assigned to a given "line".

```css
.layout {
  display: grid;
  grid-template-columns: [aside-start] 200px [aside-end main-start] 1fr [main-end];
}
```

Notice the `[aside-end main-start]`. We effectively assigned two names for a single line.

### Naming the areas

Here, we are going to use `grid-template-areas`. I find this approach a bit easier to understand.

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>HTML + CSS</title>
    <link rel="stylesheet" href="styles.css" />
    <style>
      .grid-container {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-template-areas: "sider-1 content sider-2";
      }

      .main-content {
        /* Notice the name here */
        grid-area: content;
      }
    </style>
  </head>
  <body>
    <div class="grid-container">
      <div class="main-content">
        <p>...</p>
        <p>...</p>
        <p>...</p>
        <p>...</p>
        <p>...</p>
      </div>
    </div>
  </body>
</html>
```

**Note that the value of `sider content sider` would not be valid**. If you were to write `grid-area: sider` where would the browser put the element?

Split it in two and assign to each side? If you switch from naming lines to naming areas, you might try to write such definition and wonder why it does not work.

#### Empty grid cells

You can define "gaps" in the grid areas by _dots_.

You can either have one or multiple dots assigned as an "empty" cell. In the example below, I'm using three dots, but nothing stops me from using one or two or four dots or any other number of dots to denote empty space.

```css
.element {
  grid-template-areas:
    "aside main"
    "... footer";
}
```

Pretty interesting. It only adds to the power of this feature.

## The usefulness of `user-select`

You most likely wanted to copy some ID from somewhere right? Did you have a hard time selecting the whole ID string? I
know I have.

With [the `user-select: all`](https://developer.mozilla.org/en-US/docs/Web/CSS/user-select) the browser would select the
whole string if you clicked on it! Pretty neat.

## The `appearance: none` and styling inputs

If you wish to re-style the default user-agent input styles (like checkboxes and other inputs), you might want to
use `appearance: none`. This definition tells the browser that you wish to be able to override the UA styles. Here is a
sample CSS (I'm using native CSS nesting).

```css
input[type="checkbox"] {
  appearance: none;

  position: relative;
  background: lightgray;
  height: 1rem;
  width: 1rem;
  border: 1px solid black;
  vertical-align: middle;

  &::before {
    content: "";
    display: none;
    width: 10px;
    height: 10px;
    background: red;
    position: absolute;
    inset-block: 0;
    inset-inline: 0;
    transform: translateY(50%);
  }

  &:checked::before {
    display: block;
  }
}
```

To my best knowledge, there are some subtle bugs with this property in older browsers (MDN mentions it). Having said
that, I could not find any definite answer regarding which browser versions are affected.

Some UI libraries do not use this property – they hide the native checkbox, either by using `opacity: 0` or some other
technique. Some UI libraries use this property with custom styling directly applied on the `checkbox`.

## Clickable target sizes and how to make them bigger

> Based on [this great article](https://ishadeed.com/article/target-size/).

The web nowadays is full of rich interactions. Those interactions are are mostly triggered by user actions, like click
or touch. **As developers we need to make sure the icons/buttons user is to click on are big enough**. If not, the user
might not be able to get to a certain spot in the website or complete the purchase – imagine the "buy" button being a
tiny one.

While the designers usually do a great job of making sure "action buttons" are large enough for everyone to click/tap
onto, they might miss the need to make icons and other contextual buttons large enough! This is where your expertise
come in.

First, know that there are guidelines for how large a "clickable" target should be. They differ from company to company,
but you should be safe with 44x44px. Yes, you have read that right, it is 44x44 not 16x16 or some other lower number.

### Add padding around the element

**Before adding padding to anything, understand that this could affect the layout of the page**. It could make the
parent container larger, and sometimes it is not desired behavior. Think the "close" button in the modal header. If you
add a lot of padding to that button, the overall size of the header might increase.

Having said that, it is the easiest way to improve the UX of the "clickable" elements.

### Use _pseudo-classes_

This one is interesting. Instead of adding `padding` to the element, we will be increasing the size of the "clickable"
area by using _pseudo-classes_.

```css
.card {
  /* Some styles */
  position: relative;
}

a:before {
  content: "";
  position: absolute;
  inset: 0;
}
```

In the example above, we have a card with a link inside it. We would like the card click to correspond to the `a` click.
We can achieve that by using _pseudo-selectors_! **Note that this particular example might pose a11y concerns as it
makes the contents of the card non-selectable**.

There are more great examples in the article I've linked to at the very start.

### Make buttons "full width"

This practice is quite common. Instead of having buttons side-by-side, we display them in a block-order. Of course,
increasing padding will also work here.

## The `fit-content` property

Have you ever wanted to make the element size to _exactly_ match its contents? Well, you have the `width: fit-content`
at your disposal!

This property is quite useful if you want to center an element with dynamic width (you cannot predict what it will be).
In addition to `fit-content`, consider using the `max-width` property as well.

```css
.centered {
  position: fixed;
  inset: 0;
  width: fit-content;
  height: fit-content;
  margin: auto;
  max-width: 80dvw;
}
```

There are also the `min-content` and `max-content`. The `fit-content` makes it so that the element will stretch, but
will never exceed the `max-content` value.

## Animating display via `transition-behavior`

> Read more on [MDN here](https://developer.mozilla.org/en-US/docs/Web/CSS/transition-behavior).

Since the `display` is a "discrete" _animation type_, by default, it will not work with any transitions. This is how it worked for pretty much forever, and required you to use JavaScript to create "show/appear" animations when toggling display.

**That is no longer the case if you use the `transition-behavior: allow-discrete`**.

```css
dialog {
  display: none;
  opacity: 0;

  transition-duration: 1s;
  transition-property: opacity, display;

  /* The key to making this whole thing work \/ */
  transition-behavior: allow-discrete;
}

dialog[open] {
  display: block;
  opacity: 1;

  @starting-style {
    opacity: 0;
  }
}
```

The `@starting-style`, in this example, as I understand it, is the value you want to animate from when the `display` changes.

**One gotcha to keep in mind is that you can't nest within the _pseudo-classes_**. So you if wish to animate the `backdrop`, you need to move the `@starting-style` outside of the selector.

```css
dialog {
  display: none;
  opacity: 0;

  transition-duration: 1s;
  transition-property: opacity, display;

  /* The key to making this whole thing work \/ */
  transition-behavior: allow-discrete;
}

dialog[open] {
  display: block;
  opacity: 1;

  @starting-style {
    opacity: 0;
  }
}

dialog::backdrop {
  display: none;
  background: black;
  opacity: 0

  transition-duration: 1s;
  transition-property: opacity, display;

  /* The key to making this whole thing work \/ */
  transition-behavior: allow-discrete;
}

dialog[open]::backdrop {
  display: block;
  opacity: 0.5;
}

@starting-style {
  dialog[open]::backdrop {
    opacity: 0;
  }
}
```

**Notice that the definition inside the `@starting-style` can be different for different "directions"**.

### Note about `display` property

It is worth noting _how_ the browser "animates" the `display` property.

- **When animating from `display: none` to `display: block` (or any other value)**, browser will flip the element to `block` at the very start of the animation, so it is visible. **This is why you usually want to animate `opacity` alongside the `display` property**. Otherwise the element will "pop in" out of nowhere!

- **When animating from `display: block` to `none`**, browser will **flip the element to `none` at the END of the animation**. Again, this is to ensure the element is visible thought the animation.

## The Popover API

The [popover API](https://developer.mozilla.org/en-US/docs/Web/API/Popover_API) introduced a wave of very useful features.

First is the fact that it allows you to natively created popovers on a webpage **without using JS**.

The following snippet will take you very far.

```html
<button popovertarget="mypopover">click me</button>
<div id="mypopover" popover>popover content</div>
```

At the time of writing, the default is to place the popover at the middle of the viewport.

Then, you can control the look and feel via different states of the `popover` element.

```css
/* Starting styles */
[popover] {
  opacity: 0;
  transition: all 0.7s allow-discrete;
}

/* Animate to */
[popover]:popover-open {
  opacity: 1;
}

/* When toggling the button */
@starting-style {
  [popover]:popover-open {
    opacity: 0;
  }
}
```

Since, **by default the popover is NOT anchored to the target**, you will most likely want to combine this with the [CSS Anchor Positioning](https://www.oidaisdes.org/popover-api-accessibility.en/).

```html
<button popovertarget="mypopover" id="target">click me</button>
<div id="mypopover" anchor="target" popover>popover content</div>
```

Notice that I did not have to apply any additional styles to make the functionality work. Of course, If I need to customize the positioning, I can do that!

- To learn more, consult the [MDN article](https://developer.mozilla.org/en-US/docs/Web/API/Popover_API) about the Popover API.

- To learn more about positioning the elements anchored to another element, read [this blog post](https://developer.chrome.com/docs/css-ui/anchor-positioning-api#use_anchor_with_top-layer_elements_like_popover_and_dialog).

## The `display: stretch`

By default, _block positioned_ elements, will _stretch_ the whole available space.

```html
<div class="container">
  <!-- The box will take up all the horizontal space of the `container` -->
  <div class="box"></div>
</div>
```

That is not the case for `inline` elements like `button` tag. **For `inline` elements, the default width is `fit-content`**.

What if you want to make the button span `100%` of the container width? You most likely would do the following:

```css
.button {
  display: block;
  inline-size: 100%;
}
```

Okay, so far so good. Now what if you add margins?

```css
.button {
  display: block;
  inline-size: 100%;
  margin-inline: 20px;
}
```

As soon as you do that, you will **notice that the button overflows the `.container`**. The size of the `margin-inline` is not subtracted from the `inline-size` of `100%`.

**One solution** would be to use `calc`.

```css
.button {
  --inline-margin: 20px;
  margin-inline: var(--inline-margin);
  inline-size: calc(100% - 2 * var(--inline-margin));
  display: block;
}
```

**But you could also use the `inline-size: stretch` property**.

```css
.button {
  inline-size: -webkit-fill-available;
  inline-size: -moz-available;
  inline-size: stretch;
  margin-inline: 20px;
}
```

At the time of writing this, the `stretch` requires vendor prefixes – something your CSS processor does for you.

Interestingly, the `display: stretch` is yet to be documented on MDN.

## The `@scope` and _proximity_ of classes

> [Based on this blog post](https://frontendmasters.com/blog/one-thing-scope-can-do-is-reduce-concerns-about-source-order/)

Consider the following HTML. What would be the color of the `#box` element?

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <style>
      .wrapper-2 #box {
        background: red;
      }

      .wrapper-1 #box {
        background: blue;
      }
    </style>
  </head>
  <body>
    <div class="wrapper-1">
      <div class="wrapper-2">
        <div id="box">Box</div>
      </div>
    </div>
  </body>
</html>
```

The `wrapper-2` class should override the `wrapper-1` class, right? **That is not the case**.

**Remember: by default, if the selectors have the same specificity, it's the order in which they are defined that determines which overrides which**.

That is quite suboptimal. If we do not do anything about it, it forces us to keep track of the order in which we define our selectors. You know that this is not happening, especially when working within a team.

So, what can we do about it? **We can use the `@scope` property to switch from "order of definition" to "proximity" heuristic**.

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <style>
      @scope (.wrapper-2) {
        #box {
          background: red;
        }
      }

      @scope (.wrapper-1) {
        #box {
          background: blue;
        }
      }
    </style>
  </head>
  <body>
    <div class="wrapper-1">
      <div class="wrapper-2">
        <div id="box">Box</div>
      </div>
    </div>
  </body>
</html>
```

Now, the box would have red background color.

### The `@scope` vs. `@layer`

- **Use `@layer` for organizing cascade rather than "scoping" styles to a specific selector**.

- **Use `@scope` for creating scoped styles that do not leak "outside" of their boundaries**.

Be _very_ careful while defining a `@layer` inside a `@layer`. You might be tempted to do this, because adding `@layer` affects specificity, but using `@scope` is better suited for this job.

## `text-wrap: balance` and `text-wrap: pretty`

Have you ever had to deal with "orphaned" words in a paragraph? Or perhaps a heading that had that one word in a new line?

Now, in 2025, we have **new `text-wrap` properties we can use to avoid such issues**.

I [think MDN explains it best](https://developer.mozilla.org/en-US/docs/Web/CSS/text-wrap):

- Use `text-wrap: balance` for headings and text that is not overly long.

  - Uses an "expensive" algorithm (in comparison to `pretty`) to calculate the layout.

  - Has an effect only on a couple of lines of text.

- Use `text-wrap: pretty` for paragraphs.

  - Uses a less compute-intensive algorithm (in comparison to `balance`) to calculate the layout of words.

  - Can be used for long paragraphs.

## `text-box-trim` property

Have you ever noticed that text has a bit of "vertical padding", even without any `padding` properties applied to it?

The `text-box-trim` property _will_ [allow you to control which of the over and under edges of the text content to trim from a text element's block container](https://developer.mozilla.org/en-US/docs/Web/CSS/text-box-trim).

You most likely want to pair it with `text-box-edge` property – [see MDN for details](https://developer.mozilla.org/en-US/docs/Web/CSS/text-box-edge).

So, if you want to **make the text have as little vertical height as possible, use the following snippet**.

```css
p {
  text-box-edge: cap alphabetic;
  text-box-trim: trim-both;
}
```

## Custom attributes

You can select things based on their attributes in CSS, like so (we have covered this in "attribute selectors" section):

```css
img[alt="some-text"] {
  /* stuff */
}
```

But, there is also the [`attr` function](https://developer.mozilla.org/en-US/docs/Web/CSS/attr) which is about to get extended to support _any attribute_ on a DOM node!

```html
<style>
  .grid {
    display: grid;
    grid-template-columns: repeat(attr(data-columns type(<number>)), 1fr);
  }
</style>

<div class="grid" data-columns="3">
  <div></div>
  <div></div>
  <div></div>
</div>
```

**There is something to be said about using CSS variables to achieve the same thing**:

```html
<style>
  .grid {
    display: grid;
    grid-template-columns: repeat(var(--columns), 1fr);
  }
</style>

<div class="grid" style="--columns: 3">
  <div></div>
  <div></div>
  <div></div>
</div>
```

The CSS variables approach might seem easier, but consider that you are mixing _styling_ with _configuration_ (at least it seems as that is the case for me).

## Checking if the element is scrollable via CSS

> Taken from [this video](https://www.epicweb.dev/talks/less-cruft-more-power-leverage-the-power-of-the-web-platform)

With the `animation-timeline: scroll(self)` we can apply a `@keyframes` animation to the element _only_ when it is scrollable!

Since you can change CSS variables within the `@keyframes` animation, you can change the element when it has a scrollbar.

### Using _space toggles_

> [You can read more about the scape toggle technique here](https://github.com/propjockey/css-sweeper#basics-of-space-toggle)

```css
@keyframes detect-scroll {
  from,
  to {
    /* So-called space-toggle */
    --can-scroll: ;
  }
}

.box {
  /* Set the initial value. Critical to get this working! */
  --can-scroll: initial;

  animation: detect-scroll;
  animation-timeline: scroll(self);

  --some-var-when-can-scroll: var(--can-scroll) <SOME_VALUE>;
  --some-var-when-cant-scroll: <SOME_VALUE>;

  /* And now you can apply CSS styles based on those values */
  border: 2px solid var(--some-var-when-can-scroll, var(--some-var-when-cant-scroll));
}
```

**The main drawback** of this approach is the complexity the _space toggle_ brings. To me, it is a bit of a hacky solution that relies on the fact that ` ` value will be skipped by the browser when analyzing the property.

```css
--some-var-when-can-scroll: var(--can-scroll) <SOME_VALUE>;
```

If the `--can-scroll` holds an "empty value", the `--some-var-when-can-scroll` will evaluate to `<SOME_VALUE>`. Given that the `--some-var-when-can-scroll` holds a valid value, the `border` will use that value for the color. If not, it will use the fallback.

It used to be that **you would use this approach to achieve greater browser compatibility**, but nowadays, the _style queries_ browser support is just as good.

### Using _style queries_

> [Based on this blog post](https://www.bram.us/2023/09/16/solved-by-css-scroll-driven-animations-detect-if-an-element-can-scroll-or-not/)

```css
@keyframes detect-scroll {
  from,
  to {
    --can-scroll: 1;
  }
}

.box {
  --can-scroll: 0;

  border: 2px solid gray;

  overflow-y: auto;

  animation: detect-scroll;
  animation-timeline: scroll(self);
}

@container style(--can-scroll: 1) {
  span {
    border: 2px solid red;
  }
}
```

Here, it is much apparent what is going on!

## `unset` vs. `revert` vs. `initial`

> [More information here](https://youtu.be/bssZTv3BgkE?t=632).

Sometimes you want to "reset" all the styles of an element to the ones from the UA. How would you do that?

When searching for solutions to this problem, I found three values: `unset`, `revert`, and `initial`. What is the difference between these three? And which one should you use?

- The `initial` property will **revert the styles to the _CSS specification defaults_ rather than UA styles**.

  - So, `initial` most likely does not do what you thought it does. I always thought of _initial_ as being the UA styles, but that is not the case!

- The `revert` works differently for _inherited properties_ and _non-inherited properties_.

  - For _inherited properties_, like `color`, it behaves like `inherit` for that property.

  - For _non-inherited properties_, like `display`, **it will style the property according to the UA**.

- The `unset` also works differently for _inherited properties_ and _non-inherited properties_.

  - For _inherited properties_, like `color`, it behaves like `inherit` for that property.

  - For _non-inherited properties_, like `display`, it **behaves like `initial` for that property**.

### Please revert all the style changes I made to this element

If you want to revert **all** the style changes, you can use `revert` to apply back the UA styles for all _non-inheritable_ properties.

```css
.foo {
  all: revert;
}
```

## The `margin-trim` property

> Based on [this blog post](https://webkit.org/blog/16854/margin-trim/)

The `margin-trim` property is there to help you remove unwanted space between elements.

IMO, components should not set their own margins – instead those should be controlled by the _container_ or their _parent_. But, we all know that CSS can be quite messy, and sometimes we have to use other approaches like negative margins OR the `margin-trim` property!

Let us say you have paragraphs inside an article.

```css
article {
  padding: 2lh;

  p {
    margin-block: 1lh;
  }
}
```

The margin of the first and the last paragraph will be added to the padding of the article creating too much space. We can "fix" that issue by using `margin-trim` property.

```css
article {
  padding: 2lh;
  margin-trim: block;

  p {
    margin-block: 1lh;
  }
}
```

Now everything should look nice and tidy!

## Layout testing

> Based on [this article](https://piccalil.li/blog/a-revisit-of-the-every-layout-sidebar-with-has-and-selector-performance/) and the "Layout testing" section.

I doubt I will ever use this, but this "technique" peaked my interest.

```css
:root {
  --error-outline: 0.25rem solid red;
}

:has(> .sidebar) > :only-child,
:has(> .sidebar) > :nth-child(3) {
  outline: var(--error-outline);
  --error: "Sidebar layouts must include exactly two child elements.";
}
```

If the element containing a `.sidebar` has only one, or three (or more) children, that child element would get a red outline. The `--error` is only there to show you a descriptive "error message" in devtools.

Very interesting, is it not?
