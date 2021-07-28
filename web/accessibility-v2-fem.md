# Accessibility v2 FEM

## What is accessibility

> People with disabilities can use the web

- When the website is designed with accessability in mind, everyone can use it, not only people without any disabilities.

- Accessability means creating websites with empathy in mind. Just like we talk about _mechanical empathy_ we also have to talk about
  empathy to other human beings.

- Problems with accessability arise whenever we end up developing complex solutions that have little to do with a plain HTML.

- A lot more people have some kind of disabilities than you think. Some estimates say that one in fourth in USA.

- Whenever you build your site with accessability in mind, you end up making it better for everyone.
  Thinking about accessibility forces you to be methodological about the structure of your app.

## Screen readers

- There is no governing body when it comes to screen readers.

- Multiple versions of multiple screen readers exists.

- Usually what works for one screen reader will work for another.
  From time to time you might face a bug pertaining to a specific version of a specific screen reader.

## Accessible HTML

- A lot of accessability problems can usually be solved by writing semantic HTML.

- Please do not use `div` for everything.
  For example, use `label` tag with `input` tags. You literally get accessibility for free. **Please note that only some elements can be used with label tag**.

- TIL: There is `<input type = "submit">` that acts like a submit button.

- You might want to use `visuallyHidden` css class for hiding elements for sighted users but making sure the content is read by screen readers.

### Exercise takeaways

- You can use `p` and `span` tags within `label`.

- You might want to add `aria-label` to links. Usually the text within the link is not really descriptive.

```html
<a href="google.com" aria-label="descriptive label"
  >Click here (not really descriptive)</a
>
```

## ARIA

- Special attributes for your HTML attributes.

- Jon talks about multiple ARIA attributes.

- Jon talks about multiple ARIA roles states and properties.

- ARIA can be used within CSS as selectors.

### Live regions

- Used to control how the screen reader will react whenever a given piece of content changes.

- Think of Uber and the "click to search a driver" flow. If the driver is en-route to your place, the content of the tag changes and given the correct `aria-live` setting, will be read to the user.

- Please note there is no way to implement fine-grained controls via pure HTML. The `aria-live` has only three settings. If you want to customize it more, like implementing a solution that reads to a screen reader every X seconds you would need to use JavaScript and hide/show some DOM elements.

## Focus Management

- Many popular sites like Twitter and Facebook have keyboard shortcuts defined.

- _Skip links_ are great way to skip the cumbersome navigation through the navigation (or any repeated elements on the page).

- The `tabindex` attribute is important for some interactions.

- Tab trapping for modals and popups is a must.

- Use 3rd party libraries for interactive and accessible elements. One such example would be the [`react aria` library](https://react-spectrum.adobe.com/react-aria/useDialog.html).

### Skip link exercise takeaways

- You can use `href` attribute to focus specific attributes on action using the `a` tag
  For example `<a href = "#action">` upon clicking would focus an element with id of `action` (in the exercise that was a button).

## Visual considerations

- Color contrast ratios matter for accessability compliance.

- Most sites fail with forms where the only information about the form state (is it valid or not) is depicted by input border color.
  Usually, this kind of visual feedback tells nothing people with disabilities. Add some kind of labels to your forms!

## HTML Markup considerations

- Validate your HTML from time to time.

- Consider the language attribute for html tag.

## `PrefersXX` APIs

- Use the `prefers-reduced-motion` for animations.

- Use the `prefers-color-scheme` for dark / light mode settings.
