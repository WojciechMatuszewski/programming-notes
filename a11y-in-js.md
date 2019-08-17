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
