# HTML Forms – attributes, tips and tricks

When working on any kind of application, you will need to write HTML forms. It is very rare that your application does
not have any kind of inputs.

Since you will be writing forms (probably a lot of them), it is essential to learn how `form` HTML tag works, and tags
related to it.

## Properly labelling the form

Consider the following code

```html
<form>
  <fieldset>
    <label>
      <span>Some label</span>
      <input type="text" name="text" />
    </label>
  </fieldset>
</form>
```

There are a couple of issues

1. The `form` will not have an accessible name. See [this link](https://www.stefanjudis.com/today-i-learned/forms-require-an-accessible-name/?utm_source=stefanjudis&utm_medium=email&utm_campaign=web-weekly-167-css-anchor-positioning-and-modern) to learn more.

2. The `group` which `fieldset` creates will not have an accessible name.

How we can fix those issues?

```html
<form aria-labelledby="legend-id">
  <fieldset>
    <legend id="legend-id">Your form</legend>
    <label>
      <span>Some label</span>
      <input type="text" name="text" />
    </label>
  </fieldset>
</form>
```

1. To fix the `form`s accessible name, we are explicitly adding `aria-labelled` by to it. This way, the `form` will render as `form` landmark. See [this guide](https://www.w3.org/WAI/ARIA/apg/patterns/landmarks/examples/form.html) to learn more.

2. To fix the `group`s accessible name, we render a `legend` inside the `fieldset`.

**This is how forms should look like**.

## When does the form "submit"?

> Based on [this short blog post entry](https://www.stefanjudis.com/today-i-learned/implicit-form-submission-doesnt-work-always)

For the longest time, I thought that the form will fire the `submit` event whenever you hit `enter` on one of its inputs. This heuristics appeared to be consistent, because 99% of the forms I wrote **had the "submit" button**.

```html
<form>
  <!-- some inputs -->
  <button type="submit">Submit</button>
</form>
```

**It turns out, the "implicit" submit "on enter" does not always happen**. It will happen when the following conditions are met.

1. The `form` has a `submit` button.
2. The `form` **only has a single `input` element**.

So, I could have a form without the `submit` button, with a single `input` element that would also trigger this "implicit" submit event.

```html
<form>
  <!-- Hitting enter while editing this field will "submit" the form -->
  <input type="text" />
</form>
```

But, **the "submit" event will not fire you have multiple inputs WITHOUT the "submit" button**.

```html
<form>
  <!-- Hitting enter while editing this field will NOT "submit" the form -->
  <input type="text" />
  <input type="text" />
</form>
```

If you **wish to preserve the "submit" behavior in this case, use a `hidden` submit button**.

```html
<form>
  <!-- Hitting enter while editing this field will NOT "submit" the form -->
  <input type="text" />
  <input type="text" />
  <button type="submit" hidden>Hidden Submit</button>
</form>
```

## Getting all elements of form from `event`

Use `event.target.elements` to get all form elements. **You have to add name prop to the form elements**

## Resetting the form values after submission (uncontrolled)

You can use the `target.reset()` to reset the form values.

## Using the `disabled` attribute on a `button`

> [Based on this great article](https://css-tricks.com/making-disabled-buttons-more-inclusive/).

One thing that I came to appreciate is seeing a good use of the `disabled` attribute on form-related elements, like
inputs or buttons.

One rule of thumb I follow is **to avoid putting the `disabled` attribute on the "submit" button of the form**.
The `disabled` attribute will make it impossible to focus `button`.

**If there was a focus on the button prior to clicking it, now the focus will be put on the document itself!**
This is not ideal, as the user might be confused what just happened (his focus is seemingly "gone").

**Instead of the `disabled` attribute, consider the `aria-disabled`**. It will NOT prevent button clicks for you,
but it will provide the same _semantics_ as the `disabled` attribute. The "click prevention" part should be handled
within the application itself.

### What about the `pointer-events`?

I oftentimes see developers use the `pointer-events: none` in hopes of disabling "interactions" for certain elements.
**The issue with `pointer-events: none` is that it does not prevent keyboard events**.

Usually, a **better alternative is the `insert` attribute**. This one is a bit nuclear, as it will prevent any clicks,
keyboard and other interactions, but in most cases, it is what should be used instead.

Note that the `insert` is _a relatively_ new attribute. Please check the browser support before using it.

## The `search` element

The `search` element creates a _search landmark_, so you don't also need to add `role` of search to the form.

```html
<search>
  <form>
    <label>
      Search
      <input type="search" name="search" />
    </label>
  </form>
</search>
```

- The **`search` element is not for displaying search results**.

- You can have multiple `search` elements on the page.

- **Consider using the `output` element to tie the form with the results**.

- You still need to annotate the input with `type="search"`.

### Semantic search with pre-defined queries

Combine the `search` with `datalist` and you have created a very good search experience.

You can [play around with the code here](https://codesandbox.io/p/sandbox/mwwww4).

```html
<div>
  <search>
    <form>
      <fieldset>
        <legend>Search for user</legend>
        <!-- Notice the `list` attribute here -->
        <input type="search" name="search" list="predefined-queries" />
        <button type="submit">Search</button>
      </fieldset>
    </form>
  </search>
  <!--  Predefined queries -->
  <datalist id="predefined-queries">
    <option value="Ben"></option>
    <option value="Tom"></option>
    <option value="Amanda"></option>
    <option value="Thomas"></option>
  </datalist>
</div>
```

## The `type=search` input

There are several differences between the `type="search"` and `type="text"` inputs.

1. The **browser might display the "X" icon to clear the search value**.

2. The **browser might suggest previously entered values when typing**.

## The `type=file` input

### Styling the native button

If you use the `type=file` input, by default, the browser will render a button for the user to click on to bring up the
file picker.

This button will have native browser styles, which you most likely would want to change. How to go about changing it?

1. Use the `::file-selector-button` pseudo-element

This option is good when you do not need any _droppable_ areas for your files (you most likely need them). Read more
about [this pseudo-element on MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/::file-selector-button). You can
also find
more [information here](https://www.stefanjudis.com/today-i-learned/how-to-style-the-select-button-of-file-inputs).

### Resetting the input after file upload

It is imperative to reset the input value after the file was uploaded. If we fail to do this, the user will NOT be able
to upload the same file again. People often forget to do this, and this leads to bugs.

```jsx
function FormStuff() {
  return (
    <form>
      <input
        type="file"
        onChange={(event) => {
          const files = event.currentTarget.files;
          // process the files

          // Reset the input
          event.currentTarget.value = "";
        }}
      />
    </form>
  );
}
```

## Form errors & validation

### Attribute-based validation

While for more robust forms, you might need to reach out for a JS API, in some cases, you might get away with an
HTML-only validation of inputs. There are many attributes on the native HTML form elements that could aid you.

- The `minLength`
- The `required`
- The `pattern`

Also, let us not forget about the `type` property of the input. The most common are `type="email"` or `type="tel"`.

#### Visual cues

The input elements have different "states" represented by pseudo-classes. You can use them to style the inputs based on
their validity.

- The `:invalid` pseudo-class **allows you to target invalid inputs, but the `:invalid` is also applied when the first
  loads!**.

- The `:user-invalid` is **very similar to `:invalid` but IS NOT applied when the page first loads**.

  - Sadly, the **browser support for `:user-invalid` is lacking**, so we cannot really rely on it.

The one technique to "fix" the surprising (at least to me) behavior of the `:invalid` pseudo-class is to
the `:not(:placeholder-shown)` to only set the invalid styles when the input is filled.

```css
input:not(:placeholder-shown):invalid {
  border: 1px solid red;
}
```

This **requires your inputs to have placeholders**. If they do not have placeholders, you might want to try to style the
inputs `:invalid` state, only when the form was submitted.

```css
form.submitted input:invalid {
  border: 1px solid red;
}
```

The `submitted` class would be added in JavaScript. I could not find any other way to make this work.

#### How JavaScript fits the picture

You know how, when you use the native HTML validation, the browser will display validation errors in the popovers? You
can **control the contents of those popovers by using JS and the `Constraint Validation API`**.

```js
function FormStuff() {
  const inputRef = useRef();

  useEffect(() => {
    inputRef.current.addEventListener("input", (element) => {
      if (element.currentTarget.validity.typeMismatch) {
        element.currentTarget.setCustomValidity("I am expecting an email address!");
      }
    });
  }, []);

  return (
    <form>
      <input ref={inputRef} type="email" required={true} minLength="2" />
    </form>
  );
}
```

Now, upon submitting the form, the error the user will see is the one that you specified (of course, given the email is
invalid). There is
a [very good article on this topic on MDN](https://developer.mozilla.org/en-US/docs/Learn/Forms/Form_validation#validating_forms_using_javascript).

You can push this technique quite far. You could, instead of using `setCustomValidity`, append the error message to some
kind of `span` (ensure it is accessible). Then, you got yourself a quite capable form validation story.

### Form errors

Based on [this great video](https://www.accessibilityoz.com/resources/videos/error-messages-in-forms/).

#### Inline errors

There are three things you have to do to correctly annotate the input as invalid with an error message.

1. Set the `aria-invalid` property to `true`.

2. Add the `aria-describedby` attribute to the input.

3. Render the inline error message with an `id` that equals to the `aria-describedby` on the input.

Here is an example.

```jsx
<label htmlFor="name">Your name</label>
<input
    type="text"
    id="name"
    aria-invalid={nameError ? true : undefined}
    aria-describedby={nameError ? "nameError" : undefined}
    {...register("name", { required: "This field is required" })}
/>
<br />
{
    nameError ? <span id="nameError">{nameError.message}</span> : null
}
```

Doing so, has a couple of benefits.

1. Makes your input accessible.

2. Enables you to use semantic queries when testing using `@testing-library/react`.

3. Enables you to style the input based on invalid state via CSS only.

```css
input[aria-invalid="true"] {
  border: 2px solid red;
}
```

We are using the `aria-invalid` here instead of `:invalid` because we are not using the native HTML form validation.
Instead, I've opted to use the `react-hook-form` library.

#### Global form errors

Imagine a situation where the form fields are valid, but upon form submission, an error is returned from the API. Since
the error is not associated with any fields, this error pertains to the form as a whole. In such case, we have to use
another method of notifying the user about the error.

1. Create a container for the "form error." Keep it empty.

```jsx
<form>
  <div role="alert" tabIndex={-1}></div>
</form>
```

2. When an error occurs, **focus this container** and add the necessary attributes.

```jsx
<form>
  <div role="alert" tabIndex={0} aria-labelledby="formErrorHeading">
    <h2 id="formErrorHeading" className="sr-only">
      Failed to submit the form
    </h2>
    <span>Your error message</span>
  </div>
</form>
```

It is vital to add correct attributes to the container and **ensure you switch the `tabIndex` attribute**.

Here is the full snippet. Tested on a macOS screen reader, and it reads the errors as I would expect.

```jsx
import { useMutation } from "@tanstack/react-query";
import { useRef } from "react";
import { useForm } from "react-hook-form";
import "./styles.css";

function App() {
  const { mutate, isError } = useMutation({
    mutationFn: async () => {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          reject();
        }, 1000);
      });
    },
    onError: () => {
      requestAnimationFrame(() => {
        formErrorRef.current?.focus();
      });
    },
  });
  const formErrorRef = useRef < HTMLDivElement > null;

  const { handleSubmit, register, formState } = useForm({
    defaultValues: {
      name: "",
    },
  });

  const nameError = formState.errors.name;

  return (
    <form
      onSubmit={handleSubmit((values) => {
        mutate();
      })}
    >
      <label htmlFor="name">Your name</label>
      <input
        type="text"
        id="name"
        aria-invalid={nameError ? true : undefined}
        aria-describedby={nameError ? "nameError" : undefined}
        {...register("name", { required: "This field is required" })}
      />
      <br />
      {nameError ? <span id="nameError">{nameError.message}</span> : null}
      <br />
      <button type="submit">Submit</button>
      <div
        // You cannot use `visibility: none` here as that will make this container inaccessible
        ref={formErrorRef}
        role="alert"
        tabIndex={isError ? 0 : -1}
        aria-labelledby="formErrorHeading"
      >
        <h2 id="formErrorHeading" className="sr-only">
          Failed to submit the form
        </h2>
        <span>Error message</span>
      </div>
    </form>
  );
}

export default App;
```

Notice a couple of things.

1. The `requestAnimationFrame` inside the `onError` is here to ensure we get the browser the chance to change styles on
   our "form error" container. Otherwise, we might try to focus an element tha has `tabIndex=-1`.

2. I hid the heading "Failed to submit the form" behind the `sr-only` class. I do not think it's valuable to display it
   for all the users.

## Focusing the inputs

> [You can read more about `autofocus` here](https://htmhell.dev/adventcalendar/2024/2/)

There is the `autofocus` property one might use. This should help you with establishing the focus on a given input when
it is inserted into HTML.

**This means that adding this attribute dynamically to an already rendered element WILL NOT WORK**.

```html
<input name="some name" autofocus />
```

Of course, this is not a silver bullet. There are a lot of things to consider when using this attribute. If the form is below the fold, the page might unexpectedly scroll, leaving the user confused.

**Consider using the `autofocus` attribute for single-purpose pages with forms**. For example, a "login page" consisting only of inputs and the submit button.

## Helpful input attributes

> Based on [this great blog post](https://garrettdimon.com/journal/posts/fine-tuning-text-inputs)

I love when forms specify the necessary attributes for the browser and my password manager to be as helpful as possible.
You too, can create such experiences – you _just_ need to provide the right attributes to the inputs in the form!

- Use the `autocomplete` attribute to provide hints to the browser regarding a given field.

  - **Pay attention to the `name` of the field. The value you put there matters in the context of `autocomplete` attribute**.

  - [Read more about the `autocomplete` attribute here](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete)

    - The concept of _sections_ is especially interesting to me.

- The `spellcheck` attribute might be helpful for `textarea` tags.

- The `autocapitalize` might be helpful for "first name" and "last name" inputs.

## The accessible custom checkbox

> [Based on this great talk](https://www.youtube.com/watch?v=ob_M_qXeDVE&list=PLdMmkhK3RIbWMNgDNTfahg7ICBYxNhWZB&index=30)

We still have a limited ability to style checkboxes. In most cases, you will be asked a custom checkbox styles that are quite impossible to replicate with the native `input type = "checkbox"`.

How we can implement a custom checkbox with accessibility in mind?

1. Set the proper `width` and `height` sizes for the input. The sizes should pretty much match what you are replacing it with.

2. Hide it, but make sure user can still click it. **This means using `opacity: 0` rather than `sr-only` class**.

```css
input[type="checkbox"] {
  position: absolute;
  opacity: 0;
  block-size: 1rem;
  inline-size: 1rem;
}
```

**You can use this technique for any other element that you wish to replace with an image, like `svg`**.

## The `.submit` vs `.requestSubmit` method

> Based on [this blog post](https://www.stefanjudis.com/today-i-learned/requestsubmit-offers-a-way-to-validate-a-form-before-submitting-it/)

Have you ever had to submit the form via JavaScript manually? I bet you do.

And I bet you tried to use the `.submit` method, like so:

```js
const form = document.querySelector("form");

form.addEventListener("submit", () => {
  // some code
});

form.submit();
```

The problem is that, **when using the `.submit` method, the `submit` event is never fired**. The form is submitted "as is". The browser **will not run any validation checks on the form inputs**.

To _submit the form as if the user did it_, you **should consider using the `.requestSubmit` method instead**.

```js
const form = document.querySelector("form");

form.addEventListener("submit", () => {
  // some code
});

form.requestSubmit();
```

Now, the browser will check the validity of the form inputs before firing the `submit` event.

## The _Post / Redirect / GET_ pattern and native form submissions

Have you ever heard about the _PRG_ pattern and what it solves? I have not, and this is mainly because I did not have to deal with multi-page apps that did not leverage any frameworks.

What is the problem exactly? Consider submitting a form on a page. The server will respond with HTML. If you then refresh the page, **the browser will repeat the LAST request that generated the page**. So, if the last request was a POST request submitting a form, the browser will issue that request.

This could be problematic right? We would not want this to happen! **And this is the "issue" the _PRG_ pattern attempts to solve**.

Instead of responding with HTML to that POST request, you redirect the user to a given page that they are currently at. This ensures the LAST request that generated the page was a GET request, so the danger for the browser to repeat that POST request upon page refresh is zero!

## The `fieldset` and the `input` and `form` element

> You can read more about this [here](https://tetralogical.com/blog/2025/01/31/foundations-fieldset-and-legend/).

Have you ever tried to disable multiple form inputs at one? It could get repetitive and difficult, especially if those inputs live in a different components. It turns out, there is a way to **disable ALL form inputs and buttons** with `fieldset` and the `disabled` attribute.

```html
<form>
  <fieldset disabled>
    <legend>The name for the fieldset</legend>
    <label htmlFor="name">Name</label>
    <input type="text" id="name" name="name" />
    <button type="submit">Submit</button>
  </fieldset>
</form>
```

### The importance of the `legend` element

Without the `legend` (or some other accessible name), the `fieldset` won't properly read to screen readers. The `fieldset` creates a group, but that group won't be named without the `legend` element.

I usually use `sr-only` (or similar) class to "hide" the `legend` element, but still make the `fieldset` accessible to screen-readers.

```html
<fieldset>
  <legend className="sr-only">Contact form</legend>
</fieldset>
```

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

## The "auto-grow" `input`, `textarea` and `select` elements

> Based on the contents of [this blog post](https://olliewilliams.xyz/blog/html-renaissance/)

> Read more about it [here](https://frontendmasters.com/blog/what-you-need-to-know-about-modern-css-2025-edition/#field-sizing)

How many times have you had to implement a `textarea` tag that grows with the content?

How many times have you reached for a library like `react-textarea-autosize`?

**There is no need to add _yet another library_ do your dependencies for this**.

Use `field-sizing` CSS property with `content` value. See [this MDN article](https://developer.mozilla.org/en-US/docs/Web/CSS/field-sizing).

## The `output` HTML tag

> Based on [this article](https://denodell.com/blog/html-best-kept-secret-output-tag).

> [More examples of usage](https://rud.is/drop/output.html).

If, based on input(s) values, you want to announce something to the user, look no further but to the `output` tag.

I really like the "Password strength" example.

```html
<label for="password">Password</label>
<input type="password" id="password" name="password" />
<output for="password"> Password strength: ${calculated strength} </output>
```

Yes, you could use `role="status"` to an element instead, but I believe we ought to use semantic HTML tags!

## Simplest possible OTP form

> Based on [this article](https://cloudfour.com/thinks/simple-one-time-passcode-inputs/)

Before you reach for _yet another_ dependency to implement the OTP input, stop and consider using plain HTML.

```html
<form>
  <fieldset>
    <legend>Code</legend>
    <input type="text" inputmode="numeric" autocomplete="one-time-code" maxlength="6" pattern="\d{4}" required />
  </fieldset>
</form>
```

1. We are using `type="text"` because the value of the input is not really a "number". `0004` would still be valid.

2. The `inputmode="numeric"` enables virtual numeric keyboard on touch devices, like mobile phones.

3. The `autocomplete="one-time-code"` is a hint for password managers and browsers.

4. The `maxlength` and `pattern` deal with validation. **You should also validate on the backend as well!**.

Yes, it is not this fancy input with multiple "sub-inputs", but do you really need it to be?
