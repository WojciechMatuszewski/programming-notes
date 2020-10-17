# a11y in Javascript

Notes from FrontendMasters workshop.

## Visibility vs Opacity vs Visually-Hidden vs Display

### Visibility

- element still occupies given space (width and height)
- is invisible as in like `opacity:0`
- **takes accessibility information away**

### Opacity

- like `visibility` but **does NOT take accessibility information away**

### Visually-Hidden

- preserves accessibility information
- rips element from the DOM flow

### Display

- pulls element from the DOM flow
- **takes accessibility information away**

## Accessibility Tree

- parallel structure to the DOM
- uses platform a11y APIs (Mac is different from Windows)

## TabIndex

- **you should very strongly prefer `tabIndex:0`**, this will ensure that the
  page has the _natural flow_.
- `tabIndex:-1` is mainly used for focusing by javascript. You will not be able
  to focus it using _normal flow_ (eg. tab key).

- anything above `tabindex:0` will **fuck up your document flow**. By setting
  `tabindex` that way you are now responsible for managing focus on the whole
  page. GG 👋

## Native elements vs generic ones

- you should always prefer using semantic elements. They come with many features
  baked in like :
  - proper focus management
  - proper event handling
  - and many more..

## Links vs Buttons

[Great article](https://marcysutton.com/links-vs-buttons-in-modern-web-applications)

- Buttons for actions
- Links navigate

## Outline

- **DO NOT SET `outline:0`**
- use css to customize behavior (like `:focus-visible`)

## Live Regions

- used to announce something (like combobox filtering result)
- can be `polite` (non-interrupting) and `assertive` (interrupts previous
  announcement)
- live regions can be useful when dealing with forms and alerts about validity
- also about informing use that an error occurred or that something was saved

## `prefers-reduced-motion`

- you can use it with media queries (media query reacts on hardware level, eg.
  user preferences inside system settings)
- should be used to soften or turn off given animation

## Attributes

### `aria-labelledby`

This attribute is useful when you need to group elements.

It takes groups of `ids` to group those elements together.

As an example, lets suppose you have such structure

```jsx
<p>
  Your movies list is empty<Link to={to}>Add some</Link>
</p>
```

This would result in screen readers reading it as:

- "Your movies list is empty"
- "Add some"

Which feels disconnected does not it?

We can fix it by using `aria-labelledby` and a simple `div` tag.

```jsx
<div aria-labelledby="text link">
  <span id="text">Your movies list is empty</span>
  <Link to={to} id="link">
    Add some
  </Link>
</div>
```

Now both of these elements would be read after each other, much better :).

### `aria-describedby`

I would use this attribute mainly with forms and helper text / error text.

Lets say your structure looks as follows:

```jsx
<input id="id" type="text" />;
{
  helperText && <span>Helper text</span>;
}
{
  errorText && <span>Error Text</span>;
}
```

The problem with this structure is that `helperText` and `errorText` is not associated with the `input`, thus not being read by screen readers.

To fix the issue you should use `aria-describedby` attribute.

```jsx
// the ids for helper and error text should be created dinamically
const helperTextID = helperText ? `${name}-helper` : ""
const errorTextID = errorText && isInvalid ? `${name}-error}`: ""
<input id = "id" type = "text" aria-describedby= {`${helperTextID} ${errorTextID}`}>

{
  helperText && <span id = {helperTextID}>Helper text</span>;
}
{
  errorText && <span id = {errorTextId}>Error Text</span>;
}
```

Much better!
