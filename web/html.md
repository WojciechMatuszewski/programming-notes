# HTML

## The `inert` attribute

Have you ever tried to make a whole section of the UI _non-focusable_? If so, maybe your first instinct was to use `tabIndex` of `-1` on the parent, only to notice that you can focus on children Maybe you have tried to use the `pointer-events` property, only to notice that you still can focus stuff.

There are a lot of edge cases!

Here is where the `inert` attribute comes in. **When set on the parent, it makes all children (and the parent) non intractable**. No focus events, no pointer events, nothing. It's like using `fieldset` to disable every input within the form. You can read more about the [`inert` attribute here](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/inert).

Where is this useful? Especially modals. If you render the modal outside of the "main" of your page, you can then set the `main` to be `inert` when the modal displays. You **get the focus trap for free**.

> Note that, when using the `dialog` element, you DO NOT have to use the `inert` on the backdrop. It's already implemented by the browser!

Another use case are **tooltips, where you do not want the browser to ever focus on the tooltip itself**. This is much better than adding `tabIndex` of `-1` as using the `tabIndex` directly interferes with the browser way of handling focus (it might change the order of the focus).

## The `valueAsNumber` property on the input element

How many times have you had to parse the number from the `type="number"` input field? Probably a lot.
Since the value is of type `string`, you can introduce a bug while doing so. Would not it be better to leave the browser do it for you? Most likely.

It turns out, it is possible – the `type="number"` input field has the `valueAsNumber` property. You can use it in your event handlers.

```jsx
  <input type="number" onChange={e => e.currentTarget.valueAsNumber}>
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

## The magic of `content-visibility`

> Based on [this](https://web.dev/dom-size-and-interactivity/?ck_subscriber_id=1352906140) and [this](https://web.dev/content-visibility/) blog post.

At some point you might encounter a website where the number of the DOM nodes is huge. This might be a blog, this might be some other interactive site. But the problem is the same – the amount of the DOM nodes causes the browser to freeze when rendering the initial content.

**If you are dealing with a list consider virtualizing the content**. But what if that is not possible? What if the content is structured in a way that makes it impossible to collect into a list? Luckily, the browser vendors come with some help. **Enter the `content-visibility` property**.

The **`content-visibility` is a NATIVE way to tell the browser to defer rendering some parts of the webpage to when the content enters the viewport**. It is like a native virtualization. Of course it does not handle all the cases that super well (but it is a built-in API that requires 0 KiB of JS to implement).

You have three values to choose from.

- The `auto`.
- The `visible`.
- The `hidden`

Me being me, I always lean towards the simplest, the most "out-of-the-box" solution possible, so the `auto` property is very appealing. The browser will do most of the work for me, and I do not have to manage the state myself.

Overall, this is a great API to be familiar with. If you are using a framework, and not using `Fragments`, it is quite easy to case the DOM to be quite big causing rendering issues.

### Troubles with the element height

To make the `content-visibility` magic possible, the browser will have to know how _big_ the element is. This does not have to be a precise number. You are in luck, if the element already has defined height. But if you are rendering content, it is not feasible to know how long the element is. If that is the case, consider using `contain-intrinsic-size`.

```css
content-visibility: auto;
/* approximate guess \/ */
contain-intrinsic-size: 1000px;
```

**If you do not specify the `intrinsic-size` your scrollbar will jump around as the browser is rendering and removing elements from the page**.

But do not fret! Browser is here to help you.

You can also specify the `content-intrinsic-size: auto SOME_PX` which works in the following way

1. When page renders, the browser will use the `SOME_PX` value.
2. When user scrolls to the element, the browser will remember the elements _actual_ size.
3. When user scrolls up/down, the browser will use its "memory" and not the value you provided for the element height.

### Toggling content

You most likely had to toggle some UI from "visible" to "hidden" state multiple times in your career. When doing so, you might have noticed, that the component we "toggle" losses state.

```tsx
function ComponentWithState() {
  const [number, setNumber] = useState(0);

  return (
    <div>
      <p>You clicked the button {number} times </p>
      <button
        type="button"
        onClick={() => {
          setNumber((prev) => prev + 1);
        }}
      >
        Click me
      </button>
    </div>
  );
}

// In some other component
visible ? <ComponentWithState /> : null;
```

If you toggle the component this way, every time it starts being visible, the `number` will be zero.

**This example is quite trivial, but think about examples where rendering the component takes a lot of work**. In such examples, there might be visible "lag" when you toggle the component to be visible again.

#### What about `display: none`?

When working with React (I'm unsure about other frameworks) you might notice that using `display: none` seem to be working just like the `content-visibility: hidden`. **This is a feature of the framework, not the web**.

Keep in mind that React uses VDOM for keeping the state. As such you won't see this "destruction of state" while using React.

### A11Y benefit over virtualization

> Learn more [here](https://web.dev/articles/content-visibility#a_note_on_accessibility).

After adding virtualization to your website, you might notice, that the "search" browser functionality does not work that well. In fact, it pretty much does not work at all. When virtualizing a list, we render only _a subset_ of nodes at a given time. What if the user wants to find a note that is not rendered? Though luck.

**The `content-visibility` works a bit differently** – while the node is not rendered, it is still present in the DOM (memory), so **it is searchable**.

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

From what I could gather, the _description lists_ are for grouping `key:value` pairs of content. Think categories and items that belong to a given category or FAQ sections.

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
<!doctype html>
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

### Searching within the `details` tag

Imagine wanting so search something on the page. That content could be "hidden" inside the `details` tag.

But do not fret! **It turns out the browser already handles this problem** and will "open" the `details` tag if it has a matching search term.

### Keeping only one open

You can use the `name` attribute to "group" the `detail` elements. A group will have only one `detail` element open!

```html
<details name="a">
  <summary>First item</summary>
  <p>I'm content</p>
</details>
<details name="a">
  <summary>Second item</summary>
  <p>I'm content</p>
</details>
```

If I open the first `detail` element, and then the second, the first one will close.

### Animating the open/closed states

Similar to how you can animate the `dialog` enter/exit states, you can animate the `details` elements.

In the past, I've used the _View Transitions_ API to trigger a subtle fade when the `detail` element opens, but we can have much more control over the animation via CSS.

[Read more about animating it here](https://nerdy.dev/6-css-snippets-every-front-end-developer-should-know-in-2025).

**Mind the `interpolate-size: allow-keywords**. As I understand it, this property allows you to animate from "hidden" to "auto".

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

> You can learn how to animate the enter/exit states [here](https://nerdy.dev/6-css-snippets-every-front-end-developer-should-know-in-2025#transition-a-dialog).

Some time ago, browsers started to introduce the `dialog` element! This is a huge win as modals/dialogs are notoriously hard to get right (mainly the aspect of focus management and accessibility).

Another neat thing about this element is that **you do not have to use `z-index` to position it on top of the content**. The contents of the `dialog` are displayed in a special layer called **_top layer_**. No more `z-index` wars!

```html
<dialog id="myDialog">
  <form method="dialog">
    <p>Some text</p>
    <button type="submit">Close</button>
  </form>
</dialog>

<button onclick="myDialog.showModal()">Open</button>
```

You will need some JavaScript to show the dialog/modal and most likely get the return value when it closes, but apart from that, the implementation requires no JavaScript at all!

### Closing the `dialog` without form tag

> Based on [this blog post](https://nerdy.dev/closedby-any)

Instead of having an explicit `form` with a `button` inside the dialog to close it, consider using the `closedBy="any"`

```html
<dialog closedBy="any">
  <p>Hi there</p>
</dialog>
```

Tapping or clicking outside the dialog will close it.

**I still believe that having an explicit "close" button is the way to go**, since it might not be obvious to users how to close it otherwise.

### Different flavours of the `dialog` element

To show the dialog, you can use the `showModal` or `show` APIs. What is the difference?

- The `showModal` **traps the focus** and **displays the backdrop**. It behaves as a regular "modal" you are used to.

  - Good for modals.

- The `show` does not trap the focus and there is no backdrop.

  - Good for notifications and other popups. You could also implement tooltips via this API.

## The `popover` attribute

> This API is similar to the `dialog` element. You can [read about the differences here](https://developer.chrome.com/blog/introducing-popover-api#the_difference_between_a_popover_and_a_dialog).

Another great addition to the web. Just like the `dialog` element, it will display in the _top layer_.

```html
<button popovertarget="popover">Toggle the popover</button>
<div id="popover" popover>Content</div>
```

**The `popover` attribute has a feature of "soft dismiss"** which means _clicking outside_ will close the popover.

## The `radio` element is much more powerful than you think

Have you ever had to implement a gallery of some sorts, where users can _row_ through different images and then loop to the beginning when they are at the end?
You most likely used JavaScript for that right? Well, you did not have to!

Check out [this video about "a looper"](https://www.youtube.com/watch?v=vwgihljM2e4)

The "looper" is a collection of `radio` elements with the same name.

```html
<fieldset id="looperGroup">
  <legend>The Looper</legend>

  <label>
    Option 1
    <input type="radio" name="the-looper" checked value="option-1" />
  </label>
  <label>
    Option 2
    <input type="radio" name="the-looper" value="option-2" />
  </label>
  <label>
    Option 3
    <input type="radio" name="the-looper" value="option-3" />
  </label>
  <label>
    Option 4
    <input type="radio" name="the-looper" value="option-4" />
  </label>
</fieldset>
```

If you focus into the `fieldset`, you will be able to use the arrow keys to bo back and forth between different radios.
**Combine this with the `has(input:checked)` state and you have a mini-state management done in HTML and CSS**.

Okay, so how does this relate to a gallery?

You can put whatever inside those labels. Think images, or something else. Hide the radio buttons with CSS and you have a carousel!

Pretty neat stuff

## Labelable elements

You most likely use `for` and `id` pairs to associate the `label` tag with an `input` tag.

**But did you know you can associate the `label` with other elements, not just the `input` tag?**.
In fact, you can do this for all _labelable_ elements. [Link to the spec](https://html.spec.whatwg.org/multipage/forms.html#category-label).

> button input (if the type attribute is not in the Hidden state) meter output progress select textarea form-associated custom elements

Very interesting!

## The element(s) with `id` is globally accessible

If you specify the `id` attribute on the element, you can access that element by referring to the id. No need for `querySelector`.

```html
<div id="foo">
  text
  <div></div>
</div>
<script>
  // Logs the element
  console.log(foo);
</script>
```

**This will also work when you have multiple elements with the same `id`**. If that is the case, **the global variable will hold a `HTMLCollection` array**.

## The `section` tag

> Based on [this](https://www.stefanjudis.com/today-i-learned/section-accessible-name/) and [this](https://www.smashingmagazine.com/2020/01/html5-article-section/) blog post

- The `section` tag was invented for browsers to implement _HTML5 outlining_.

  - This feature would allow you to use `h1`s everywhere, and depending on how "deep" they are in `section` elements, they would become the "right" levels of headings like `h2` and `h3`.

    - **At the moment of writing this, no browser implements this spec**.

- There are no free lunches – you **have to make sure that you have proper hierarchy of headings on your website**.

  - Frameworks can help with that. I've seen component libraries implement generic "heading" component that renders the correct tag based on the parents of the tag.

- **To "label" the `section` tag, consider using `aria-label` or `aria-labelledby`**.

  - Nesting a heading tag inside a section does NOT associate the text of that heading with the "label" of the section.

    - **In fact, the "default" role of the `section` element is "generic". The `section` tag gains the "region" role only when you use `aria-label` or `aria-labelledby` attribute on the section**.

To sum up:

1. If you want the `section` tag to have a _landmark role_, use the `aria-label` or `aria-labelledby` attributes on the section. Be mindful that there might be translation issues when using `aria-label`.

2. No browser implements the _HTML outlining_ spec, so you have to be mindful of the headings hierarchy on your webpage.

3. The `article` tag is a great way to group things together. Those things should "stand on their own".

## The `table` tag

To build an accessible table, you do not have to employ lots of effort. The [example from MDN will do](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table).

**Notice the `caption` element**. I was recently building a table and failed to include the `caption` element in the markup. If I did, I bet writing tests for the table would be easier (because I could use semantic selectors).

Also, keep in mind that the **`caption` element only makes sense in the context of a table**.

## The `hr` tag

I was recently reading about the support for the `hr` element in the `select` tag. [You can read more about it here](https://developer.chrome.com/blog/hr-in-select).

```html
<select name="majors" id="major-select">
  <option value="">Select a major</option>
  <!-- This is pretty new. Might not work in your browser. Proceed with caution of production. -->
  <hr />
  <option value="arth">Art History</option>
</select>
```

That got me thinking: [if `hr` has semantic meaning within a `section` tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hr), and we should not use it to separate "random" tags, how does it work in the `select` tag?

I could not find any definitive answer. For now, I will stick to visually styling a `div` tag as a horizontal line and use `hr` inside a `section` with multiple paragraphs if necessary.

## The Invoker Commands API

> Based on [this MDN article](https://developer.mozilla.org/en-US/docs/Web/API/Invoker_Commands_API?utm_source=stefanjudis&utm_medium=email&utm_campaign=web-weekly-167-css-anchor-positioning-and-modern)

I find this API fascinating. It adds more capabilities to the HTML and, in some cases, allows developers to skip writing JS to implement a given feature.

Consider the following HTML:

```html
<button commandfor="mypopover" command="toggle-popover">Toggle the popover</button>

<div id="mypopover" popover>
  <button commandfor="mypopover" command="hide-popover">Close</button>
  Popover content
</div>
```

Notice that I can open and close the popover by the "power" of markup only. Is it not amazing? At the time of writing, the API is still work in progress. Firefox and Safari are yet to support it.

### Adding popovers on hover

> Based on [this section in the blog post](https://modernwebweekly.substack.com/i/177501059/better-tooltips-with-the-interestfor-attribute)

The `commandfor` attribute is great when you want to display something upon a user _clicking_ a button, but what about the _hover_ interaction?

**You can use `interestfor` to trigger a popover on hover.**

```html
<button interestfor="popover">Hover over me</button>
<div id="popover" popover="hint">Popover</div>
```

That's it. It could not get easier than this!

There are more styling options available for this feature. You can control the delay and so on. See [this link](https://modernwebweekly.substack.com/i/177501059/mouse-and-keyboard-delays) to learn more.

## Quotes and intl

> Based on [this article](https://www.stefanjudis.com/today-i-learned/how-to-use-language-dependent-quotes-in-css/)

First, I learned that different languages can have different-looking quotation marks. It makes sense for Asian languages, but it's interesting that French quotation marks are also different from English ones.

Either way, **you have two options**:

1. Use the `q` element if you want to add _inline quotes_.

```html
<p>They say they <q>were colorful and smiling</q></p>
```

2. Use `content: open-quote` and `content: close-quote`.

```css
blockquote::before {
  content: open-quote;
}

blockquote::after {
  content: close-quote;
}
```

The **`blockquote` element does not render quotation marks**. It's only there [to _indicate_ that the containing text is a quotation](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote).
