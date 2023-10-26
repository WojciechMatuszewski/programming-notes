# HTML Forms – attributes, tips and tricks

When working on any kind of application, you will need to write HTML forms. It is very rare that your application does not have any kind of inputs.
Since you will be writing forms (probably a lot of them), it is essential to learn how `form` HTML tag works, and tags related to it.

## Getting all elements of form from `event`

Use `event.target.elements` to get all form elements. **You have to add name prop to the form elements**

## Resetting the form values after submission (uncontrolled)

You can use the `target.reset()` to reset the form values.

## The `type=file` input

### Styling the native button

If you use the `type=file` input, by default, the browser will render a button for the user to click on to bring up the file picker.
This button will have native browser styles, which you most likely would want to change. How to go about changing it?

1. Use the `::file-selector-button` pseudo-element

  This option is good when you do not need any _droppable_ areas for your files (you most likely need them). Read more about [this pseudo-element on MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/::file-selector-button). You can also find more [information here](https://www.stefanjudis.com/today-i-learned/how-to-style-the-select-button-of-file-inputs).

### Resetting the input after file upload

It is imperative to reset the input value after the file was uploaded. If we fail to do this, the user will NOT be able to upload the same file again. People often forget to do this and this leads to bugs.

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

While for more robust forms, you might need to reach out for a JS API, in some cases, you might get away with an HTML-only validation of inputs. There are many attributes on the native HTML form elements that could aid you.

- The `minLength`
- The `required`
- The `pattern`

Also, let us not forget about the `type` property of the input. The most common are `type="email"` or `type="tel"`.

#### Visual cues

The input elements have different "states" represented by pseudo-classes. You can use them to style the inputs based on their validity.

- The `:invalid` pseudo-class **allows you to target invalid inputs, but the `:invalid` is also applied when the first first loads!**.

- The `:user-invalid` is **very similar to `:invalid` but IS NOT applied when the page first loads**.

  - Sadly, the **browser support for `:user-invalid` is lacking**, so we cannot really rely on it.

The one technique to "fix" the surprising (at least to me) behavior of the `:invalid` pseudo-class is to the `:not(:placeholder-shown)` to only set the invalid styles when the input is filled.

```css
input:not(:placeholder-shown):invalid {
  border: 1px solid red;
}
```

This **requires your inputs to have placeholders**. If they do not have placeholders, you might want to try to style the inputs `:invalid` state, only when the form was submitted.

```css
form.submitted input:invalid {
  border: 1px solid red;
}
```

The `submitted` class would be added in JavaScript. I could not find any other way to make this work.

#### How JavaScript fits the picture

You know how, when you use the native HTML validation, the browser will display validation errors in the popovers? You can **control the contents of those popovers by using JS and the `Constraint Validation API`**.

```js
function FormStuff() {
  const inputRef = useRef();

  useEffect(() => {
    inputRef.current.addEventListener("input", (element) => {
      if (element.currentTarget.validity.typeMismatch) {
        element.currentTarget.setCustomValidity(
          "I am expecting an email address!"
        );
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

Now, upon submitting the form, the error the user will see is the one that you specified (of course, given the email is invalid). There is a [very good article on this topic on MDN](https://developer.mozilla.org/en-US/docs/Learn/Forms/Form_validation#validating_forms_using_javascript).

You can push this technique quite far. You could, instead of using `setCustomValidity`, append the error message to some kind of `span` (just ensure it is accessible). Then, you got yourself a quite capable form validation story.

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
  aria-describedby="nameError"
  {...register("name", { required: "This field is required" })}
/>
<br />
{nameError ? <span id="nameError">{nameError.message}</span> : null}
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

We are using the `aria-invalid` here instead of `:invalid` because we are not using the native HTML form validation. Instead I've opted to use the `react-hook-form` library.

#### Global form errors

Imagine a situation where the form fields are valid, but upon form submission, an error is returned from the API. Since the error is not associated with any fields, this error pertains to the form as a whole. In such case, we have to use other method of notifying the user about the error.

1. Create a container for the "form error". Keep it empty.

```jsx
<form>
  <div role = "alert" tabIndex = {-1}>
  </div>
</form>
```

2. When an error occurs, **focus this container** and add the necessary attributes.

```jsx
<form>
  <div role = "alert" tabIndex={0} aria-labelledby="formErrorHeading">
    <h2 id = "formErrorHeading" className = "sr-only">Failed to submit the form</h2>
    <span>Your error message</span>
  </div>
</form>
```

It is vital to add correct attributes to the container and **ensure you switch the `tabIndex` attribute**.

Here is the full snippet. Tested on a MacOS screen reader and it reads the errors as I would expect.

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
    }
  });
  const formErrorRef = useRef<HTMLDivElement>(null);

  const { handleSubmit, register, formState } = useForm({
    defaultValues: {
      name: ""
    }
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
        aria-describedby="nameError"
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

1. The `requestAnimationFrame` inside the `onError` is here to ensure we get the browser the chance to change styles on our "form error" container. Otherwise, we might try to focus an element tha has `tabIndex=-1`.

2. I hid the heading "Failed to submit the form" behind the `sr-only` class. I do not think it's valuable to display it for all the users.

## Focusing the inputs

There is the `autofocus` property one might use. This should help you with establishing the focus on a given input when the page first loads. **Sadly, this approach does not work with client-side rendered applications**. The DOM node has to be there already when the page loads. This attribute applies to **all elements, not just form-controls**.

```html
<input name = "some name" autofocus/>
```

Of course, this is not a silver bullet. There are a lot of things to consider when using this attribute. If the form is below the fold, the page might unexpectedly scroll, leaving the user confused.
