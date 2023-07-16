# HTML Forms – attributes, tips and tricks

When working on any kind of application, you will need to write HTML forms. It is very rare that your application does not have any kind of inputs.
Since you will be writing forms (probably a lot of them), it is essential to learn how `form` HTML tag works, and tags related to it.

## The `type=file` input

### Styling the native button

If you use the `type=file` input, by default, the browser will render a button for the user to click on to bring up the file picker.
This button will have native browser styles, which you most likely would want to change. How to go about changing it?

1. Use the `::file-selector-button` pseudo-element

  This option is good when you do not need any _droppable_ areas for your files (you most likely need them). Read more about [this pseudo-element on MDN](https://developer.mozilla.org/en-US/docs/Web/CSS/::file-selector-button). You can also find more [information here](https://www.stefanjudis.com/today-i-learned/how-to-style-the-select-button-of-file-inputs).

## Accessible form errors

Based on [this great video](https://www.accessibilityoz.com/resources/videos/error-messages-in-forms/).

### Inline errors

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

### Global form errors

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
        style={{ visibility: isError ? "visible" : "hidden" }}
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
