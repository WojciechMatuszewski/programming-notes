# Epic Web Dev notes

## Fullstack Foundations

- The **`button` element has the `type="reset"` attribute**.

  - This is super useful if you want to reset the form.

  - Please keep in mind that you can hide it and invoke it programmatically whenever you see fit.

    - This seem to be a good alternative to the `key` prop on the form! Sometimes it is hard to come up with a valid `key`.

- Kent showcases a great example where using plain form-based mutations saves us a lot of trouble.

  - Picture having to perform a network request if user clicks on a button.

    ```jsx
    <button
      onClick={() => {
        fetch("...");
      }}
    >
      Click me
    </button>
    ```

    There is a lot of things to consider.

    1. What if the user clicks the button twice?

    2. What if the request fails?

    3. What if the user clicks the button and then navigates away?

    ...

    There are a lot of edge cases! **Here is where wrapping this `button` in a form comes in handy**. If you do that, you do not have to worry about those things as they will be taken care of by the browser!

    ```jsx
    <Form method="POST">
      <button type="submit">click me</button>
    </Form>
    ```

- While I really enjoy working with the _Actions_ API, I'm not a huge fan of how Remix solved the problems of having multiple actions in the same file.

  - **Using the `name` property on a button is very neat**.

  - My main gripe is that we cannot export different functions, but rather have to perform a switch statement based on the `intent` (or other name that we set).

- The way scroll restoration works in `remix` and `react-router` is quite fascinating.

  - The `react-router` uses the _page hide_ event to save the last known scroll position for a given history entry. Each position is assigned an ID. This ID is derived from a given entry in the history (think a hash function).

  - `react-router` will then search through the saved entries, and if one matches, restore the scroll position for that given entry.

  - `remix` augments this code by injecting inline scripts to ensure the behavior works as expected with SSR.

- Kent demonstrated a very useful technique for sharing ENV variables. Instead of relying on a build tool which, in some cases makes all the values from `process.env` available, he used a _loader_ to inject those to `window`.

  ```jsx
  export async function loader() {
    return json({ env: getEnv() });
  }

  export function App() {
    const data = useLoaderData();

    return (
      <script
        dangerouslySetInnerHTML={{
          __html: `window.ENV=${JSON.stringify(data.ENV)}`,
        }}
      />
    );
  }
  ```

  While this snippet uses remix-specific APIs (the `loader` function), I think it translates well to any other framework.

- Interestingly `remix` will not prefetch all the links on the page by default. This is in contrast to Next.js.

  - What I like the most about `remix` implementation is the control I have. There are different options for `prefetch` like `intent` or `render`. Next.js does not have these options – it only exposes a boolean prop which might be problematic in some cases. For example when you want to prefetch all links on the page but only on hover.

- While it requires a bit more work, I find the `remix` way of exposing the status of the form a bit better.

  - The problem with the `useFormStatus` is that it **needs to be called in a component living inside a `form`**. In some cases, we do not have the luxury of putting the submit button inside the `form` tag. Instead we use the `id` and `formId` props to associate the button with the form.

- The way to handle metadata in `remix` is pretty similar to APIs exposed by Next.js.

  - One thing that I worries me is the crazy route names Kent had to create. He is using some kind of library for `remix` to make the routes "flat". I guess it is just a matter of getting used to those.

- Kent mentions a term I was not familiar with – the **_Splat route_**. The _splat route_ is a _wildcard route_. It sounds kind of cool!

## Web Forms

- Kent uses the `noValidate` attribute on the form and relies on the server-side validation.

  - **We still keep the HTML validation attributes for screen reader support**.

  - Specifying the `noValidate` **does not disable the support for the pseudo-classes like `:invalid`** which is pretty nice.

  - Note that the `noValidate` does not turn off every validation attribute.

  - Kent also uses the `useHydrated` hook to add the `noValidate` dynamically. We would not want to add this attribute when the JavaScript has not loaded yet. To me, the `useHydrated` hook leaks internal implementation details of the framework :C

    ```js
    function useHydrated() {
      const [hydrated, setHydrated] = useState(false);
      useEffect(() => setHydrated(true), []);
      return hydrated;
    }
    ```

    This works **because `useEffect` never runs on the server**.

- According to the workshop material, the **support for `aria-errormessage` is quite poor**. You should be using `aria-describedby` and `aria-invalid` instead.

- Kent mentions that one should either set the `aria-invalid` to `true` or `undefined` so that it is not rendered in the HTML.

  - I could not find any reference as to why setting the `aria-invalid` to `undefined` is better than setting it to either `true` or `false`.

    - While searching, I found [this great resource](https://russmaxdesign.github.io/accessible-forms/index.html) on different attributes and how they work with screen readers

- **The _accessability_ tab for a given element in DevTools is great!**. When it doubt what kind of `role` the element has, look it up there!

- The `tabIndex` of `-1` means that **users cannot focus the element via keyboard, but you can focus it programmatically**.

- The way Kent handled focus management is quite elegant. Instead of checking for each field status, we focus either the whole form, or the first invalid element.

  ```jsx
  useEffect(() => {
    const formEl = formRef.current;
    if (!formEl) {
      return;
    }

    if (actionData?.status !== "error") {
      return;
    }

    if (formEl.matches('[aria-invalid="true"]')) {
      formEl.focus();
      return;
    }

    const firstInvalidEl = formEl.querySelector('[aria-invalid="true"]');
    if (firstInvalidEl instanceof HTMLElement) {
      firstInvalidEl.focus();
    }
  }, [actionData?.status]);
  ```

  I was unaware of the `matches` method. Quite useful!

- Encoding the "global" error on the `''` field seem to be a common practice. I've done this several times, but I always felt like I'm doing something wrong.

- Kent uses the `conform` library to manage the id of the fields, validation and errors. The library seems nice, but is it really worth it pulling _yet another library_?

  - Okay, after working with nested fields, I can say that the library is pretty good. First time seeing a form library actually use the `fieldset` for something. This is great! It means the author is aware of how awesome `fieldset` is.

- First time seeing the `refine` from `zod` in action. The API seems very useful.

  ```js
  {
    file: z.instanceof(File).refine(
      (file) => {
        return file.size <= MAX_UPLOAD_SIZE;
      },
      { message: "File is too large" }
    );
  }
  ```

- While still marked as unstable, `remix` has nice APIs for handling the `multipart/form-data` requests.

- While working on adding and removing form items, Kent mentioned a very interesting quirk – **if you hit "enter" on any input, the browser will find the first "submit" button and "click" it**. Usually this is not a problem, but in some cases, it might trigger a button which you do not want to trigger. In our case, by default, the browser would trigger the "delete item" button!

  - The solution was to create a hidden submit button rendered before any other buttons.

- First time hearing the **term _honeypot_ as it relates to bots and forms**.

  - It turns out (I'm not that surprised tbh) that bots submit random forms with back links to a given site.

  - **Fundamental concept here is an input that the regular user is very unlikely to fill in**. Think an input with `display:none` or similar.

    - Bots are usually not sophisticated enough to deduce this input is a "honeypot" so they will fill it. Then you can check on the backend if this particular input was filled and take action (most likely returning a vague error message).

    - **Kent recommends using a library for this**. The [remix-utils](https://github.com/sergiodxa/remix-utils/blob/76fcb4bc706976a485e32a3e26b93404d49b3dc4/src/react/honeypot.tsx) implementation is pretty legit, but it is coupled to remix. A good source of reference implementation though.

      - The encrypted "time" field is quite interesting. It is an additional layer of protection. When form submits, we calculate the delta between the value in that field and current time. If the delta is less than X (for example 1 second), we deem the submission to come from a bot.

        - One has to have it set up so that tests do not trigger this behavior.

        - Note that the value of this "time" field is encrypted. This is yet another layer of protection against tampering with that field.

- The **protection against CSRF boils down to generating a special token you send to the frontend and include in the forms**.

  - The idea is that this token is unique on **per user-session basis**. An attacker cannot possibly craft a request with this token as it is always unique.

    - **If you do not specify the `expires` when creating the cookie, the browser will make this cookie a _session cookie_**. This means it will expire as soon as user session ends.

  - You might think that CORS would be enough to protect against CSRF. That is not the case, as some requests are so-called _simple request_ that do not require the preflight requests. As such, even if you have very narrow CORS headers, an attacker might still be able to send a request to your backend from different domain.
