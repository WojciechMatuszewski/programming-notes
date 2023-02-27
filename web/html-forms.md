# HTML Forms – attributes, tips and tricks

When working on any kind of application, you will need to write HTML forms. It is very rare that your application does not have any kind of inputs.
Since you will be writing forms (probably a lot of them), it is essential to learn how `form` HTML tag works, and tags related to it.

## The `type=file` input

### Styling the native button

If you use the `type=file` input, by default, the browser will render a button for the user to click on to bring up the file picker.
This button will have native browser styles, which you most likely would want to change. How to go about changing it?

1. Use the `::file-selector-button` pseudo-element

  This option is good when you do not need any _droppable_ areas for your files (you most likely need them). Read more about [this pseudo-element on MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/::file-selector-button). You can also find more [information here](https://www.stefanjudis.com/today-i-learned/how-to-style-the-select-button-of-file-inputs).
