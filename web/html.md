# HTML

## The `fieldset` and the `input` and `form` element

Have you ever tried to disable multiple form inputs at one? It could get repetitive and difficult, especially if those inputs live in a different components. It turns out, there is a way to **disable ALL form inputs and buttons** with `fieldset` and the `disabled` attribute.

```html
<form>
    <fieldset disabled>
        <label htmlFor = "name">Name</label>
        <input type = "text" id = "name" name = "name"/>
        <button type = "submit">Submit</button>
    </fieldset>
</form>
```

Pretty neat right?

## The `form` id attribute and the submit button

In some situations, the submit button of a `form` element lives in a completely different place than the form itself – think confirming the form submission with a modal or a dropdown.

```html
<form>
    <label>
        Name
        <input type = "text" name = "name"/>
    </label>
</form>

<!-- Somewhere else in the HTML -->
<modal>
    Are you sure you want to submit the form?
    <button type = "button">Cancel</button>
    <button type = "submit">Yes</button>
</modal>
```

In such cases, especially while using JSX, you might be **tempted to use `ref` on the form and fire the submit event manually – do NOT do that!**. There is a better way, a way which does not introduce that much overhead.

**To associate the `form` with a button, use the `form` attribute on the button**. Give the `form` an id, then use that id as the value for the `form` attribute. Here is an example.

```html
<form id = "form-with-confirmation">
    <label>
        Name
        <input type = "text" name = "name"/>
    </label>
</form>

<!-- Somewhere else in the HTML -->
<modal>
    Are you sure you want to submit the form?
    <button type = "button">Cancel</button>
    <button type = "submit" form = "form-with-confirmation">Yes</button>
</modal>
```

By using the platform, you drastically reduce the overhead of your code. Sometimes it is worth looking into MDN first, before writing complex JS/React code.

## The `inert` attribute

Have you ever tried to make a whole section of the UI _non-focusable_? If so, maybe your first instinct was to use `tabIndex` of `-1` on the parent, only to notice that you can focus on children Maybe you have tried to use the `pointer-events` property, only to notice that you still can focus stuff.

There are a lot of edge cases!

Here is where the `inert` attribute comes in. **When set on the parent, it makes all children (and the parent) non intractable**. No focus events, no pointer events, nothing. It's like using `fieldset` to disable every input within the form. You can read more about the [`inert` attribute here](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/inert).

Where is this useful? Especially modals. If you render the modal outside of the "main" of your page, you can then set the `main` to be `inert` when the modal displays. You **get the focus trap for free**.

Another use case are **tooltips, where you do not want the browser to ever focus on the tooltip itself**. This is much better than adding `tabIndex` of `-1` as using the `tabIndex` directly interferes with the browser way of handling focus (it might change the order of the focus).