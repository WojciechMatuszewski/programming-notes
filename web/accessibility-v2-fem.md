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
