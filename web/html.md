# HTML

## The `fieldset` and the `input` and `form` element

Have you ever tried to disable multiple form inputs at one? It could get repetitive and difficult, especially if those inputs live in a different components. It turns out, there is a way to **disable ALL form inputs and buttons** with `fieldset` and the `disabled` attribute.

```html
<form>
  <fieldset disabled>
    <label htmlFor="name">Name</label>
    <input type="text" id="name" name="name" />
    <button type="submit">Submit</button>
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
    <input type="text" name="name" />
  </label>
</form>

<!-- Somewhere else in the HTML -->
<modal>
  Are you sure you want to submit the form?
  <button type="button">Cancel</button>
  <button type="submit">Yes</button>
</modal>
```

In such cases, especially while using JSX, you might be **tempted to use `ref` on the form and fire the submit event manually – do NOT do that!**. There is a better way, a way which does not introduce that much overhead.

**To associate the `form` with a button, use the `form` attribute on the button**. Give the `form` an id, then use that id as the value for the `form` attribute. Here is an example.

```html
<form id="form-with-confirmation">
  <label>
    Name
    <input type="text" name="name" />
  </label>
</form>

<!-- Somewhere else in the HTML -->
<modal>
  Are you sure you want to submit the form?
  <button type="button">Cancel</button>
  <button type="submit" form="form-with-confirmation">Yes</button>
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
<input type="email" multiple="true" />
```

While you can do this, **you might want to think twice before allowing the user to submit multiple email addresses**. Here are few things to consider.

- Keep in mind that you will have to parse this list to extract individual addresses (most likely).
- Keep in mind that **the validation error messages vary from browser to browser**.
- Keep in mind that the **iOS keyboard does not include the comma by default**. This might result in a cumbersome experience for users on mobile.

## Huge amounts of DOM nodes and the `content-visibility`

> Based on [this](https://web.dev/dom-size-and-interactivity/?ck_subscriber_id=1352906140) and [this](https://web.dev/content-visibility/) blog post.

At some point you might encounter a website where the number of the DOM nodes is huge. This might be a blog, this might be some other interactive site. But the problem is the same – the amount of the DOM nodes causes the browser to freeze when rendering the initial content.

**If you are dealing with a list consider virtualizing the content**. But what if that is not possible? What if the content is structured in a way that makes it impossible to collect into a list? Luckily, the browser vendors come with some help. \*\*Enter the `content-visibility` property.

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

## The `time` and `dateTime`

For dates, consider using the `time` tag. As [described by the MDN](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time) it represents specific period in time. A common use-case when displaying timestamps, for example under posts of a given user.

```html
<article>
  <header>
    <a href="profile">Wojciech Matuszewski</a>
  </header>
  <footer>
    <time datetime="2018-07-07">July 7</time>
  </footer>
</article>
```

Notice that **the `datetime` does not have to have the same format as the date you display**. The `datetime` is for SEO robots and other machines.

## The `dl` tag

From what I could gather, the _description lists_ are for grouping `key:value` pairs of content. Think categories and items that belong to a given category.

To style the _description lists_, the `grid` type of layout should most likely be the "go-to" as it nicely works with the `key:value` pair model of the elements within the _description list_.

```html
<dl>
  <dt>Job</dt>
  <dd>President</dd>
  <dt>Email</dt>
  <dd>kobayashi.aoi@acme.co</dd>
</dl>
```

## The `template` element

The `template` HTML element creates the so-called _inert_ DOM tree. **You would use is to "prepare" a tree and then clone its contents to the real DOM**. The HTML parser will not evaluate the `template` element, though it will "process" it to ensure the contents of the `template` element are syntactically valid. [You can read more about the `template` element on MDN](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template).

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>HTML + CSS</title>
    <link rel="stylesheet" href="styles.css" />
  </head>
  <body>
    <div id="main"></div>
    <script>
      const main = document.querySelector("#main");

      const template = document.createElement("template");

      // Much faster than directly writing to `main`
      template.innerHTML = `
        <p>first</p>
        <p>second</p>
      `;

      main.appendChild(template.content);
    </script>
  </body>
</html>
```

[According to this blog post](https://nolanlawson.com/2023/12/02/lets-learn-how-modern-javascript-frameworks-work-by-building-one) the `template` was originally designed for _web components_, but now is at the core of many web frameworks (used to update the DOM).

## Native Accordion (collapsible)

Yea you could use the `button` & `div` and look at a11y spec but... **there is a native way of doing collapsible boxes**.
You can use `summary` and `detail` HTML tags.

```html
<details>
  <summary>title</summary>
  <p>content</p>
</details>
```

Sadly at the time of writing this, there is **no built-in way of animating the collapse state**.

## Native Combobox

Again, the same with as with `Accordion`. There is a native way of doing this by using `input` and `datalist`.

```html
<input list="languages" placeholder="Choose language" />

<datalist id="languages">
  <option>Python</option>
  <option>Javascript</option>
  <option>Java</option>
</datalist>
```

The main benefit of this approach is also it's main drawback. Since it is a native implementation, it is not that flexible. **You will have a hard time styling it**. I could not find a way to style it even a tiny bit. **The demos I've encountered seem to work on CodePen but does not work when I run them locally via plain `index.html` and my browser**.

## Native Dialog

Some time ago, browsers started to introduce the `dialog` element! This is a huge win as modals/dialogs are notoriously hard to get right (mainly the aspect of focus management and accessibility).

Another neat thing about this element is that **you do not have to use `z-index` to position it on top of the content**. The contents of the `dialog` are displayed in a special layer called **_top layer_**. No more `z-index` wars!

```html
<dialog>
  <form method="dialog">
    <p>Some text</p>
    <button type="submit">Close</button>
  </form>
</dialog>
```

You will need some JavaScript to show the dialog/modal and most likely get the return value when it closes, but apart from that, the implementation requires no JavaScript at all!
