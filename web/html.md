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

## The `valueAsNumber` property on the input element

How many times have you had to parse the number from the `type="number"` input field? Probably a lot.
Since the value is of type `string`, you can introduce a bug while doing so. Would not it be better to leave the browser do it for you? Most likely.

It turns out, it is possible – the `type="number"` input field has the `valueAsNumber` property. You can use it in your event handlers.

```jsx
  <input type = "number" onChange = {e => e.currentTarget.valueAsNumber}>
```

Pretty cool if you ask me!

## The input type of `email` and multiple email addresses

You might be familiar with the `multiple` attribute on the `input` element when used in conjunction with the `file` type. It turns out that **the `multiple` attribute also works for inputs of type `email`**.

```html
<input type = "email" multiple = "true">
```

While you can do this, **you might want to think twice before allowing the user to submit multiple email addresses**. Here are few things to consider.

- Keep in mind that you will have to parse this list to extract individual addresses (most likely).
- Keep in mind that **the validation error messages vary from browser to browser**.
- Keep in mind that the **iOS keyboard does not include the comma by default**. This might result in a cumbersome experience for users on mobile.

## Huge amounts of DOM nodes and the `content-visibility`

> Based on [this](https://web.dev/dom-size-and-interactivity/?ck_subscriber_id=1352906140) and [this](https://web.dev/content-visibility/) blog post.

At some point you might encounter a website where the number of the DOM nodes is huge. This might be a blog, this might be some other interactive site. But the problem is the same – the amount of the DOM nodes causes the browser to freeze when rendering the initial content.

**If you are dealing with a list consider virtualizing the content**. But what if that is not possible? What if the content is structured in a way that makes it impossible to collect into a list? Luckily, the browser vendors come with some help. **Enter the `content-visibility` property.

The **`content-visibility` is a NATIVE way to tell the browser to defer rendering some parts of the webpage to when the content enters the viewport**. It is like a native virtualization, but of course it does not handle all the cases that super well (but it is a built-in API that requires 0 KiB of JS to implement).

You have three values to choose from.

- The `auto`.
- The `visible`.
- The `hidden`

Of course, me being me, I always lean towards the simplest, the most "out-of-the-box" solution possible, so the `auto` property is very appealing to me. The browser will do most of the work for me, and I do not have to manage the state myself.

To be clear: **this is not a silver bullet**. With **the `auto` property, the scrollbar might jump around while the browser shows/hides the content**.
Some articles recommend using the `IntersectionObserver` API to handle the state manually, which is what you most likely will end up doing.

Overall, this is a great API to be familiar with. If you are using a framework, and not using `Fragments`, it is quite easy to case the DOM to be quite big causing rendering issues.

## DOM and Shadow DOM

### Shadow DOM

- This concept is **used in the context of custom web components**.

- It is **separate from the "main" DOM**.

  - This means you can have multiple IDs that repeat, and you do not have to worry about specificity rules.

- The main benefit is the **style isolation. Your CSS will NOT clash with any other CSS on the page**.

- Apart from the regular selectors, you get access to the `:host` and `:host-context` pseudo-classes.

  - These only exist in the context of the Shadow DOM.
