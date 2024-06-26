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

### Wrapping up

- Use the `type="reset"` on the button to implement "reset" functionality. This blew my mind. So simple, yet I think not so many people know this.

- Single-element forms with buttons are a good idea.

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
      { message: "File is too large" },
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

### Wrapping up

- It looks to me that using `noValidate` but still providing those HTML-based validation attributes is the way to go. It makes sense as it allows assistive technologies to read out validation spec for the end user, while also allowing us to have our custom "look and feel" for error messages.

  - You can even go further and add it only when React hydrated the page.

- To properly associate the HTML input element with an error message, use the `aria-describedby="some_id"` and `aria-invalid={true}`.

- First time seeing the concept of a _honeypot_ in action. Very useful pattern.

- The CSRF token makes sense. It prevents an attacker _forging_ a request to our site from a 3rd party domain.

  - Keep in mind that CORS is not enough here. Some request might classify as "simple request". For those, the browser will not send the OPTIONS (preflight) request.

  - Encrypting the CSRF token is a very good idea.

- The **browser is only capable of using `POST` or `GET` method on forms**. The framework might, and in this case, will handle `DELETE` form methods, but when JavaScript has not loaded yet, using `DELETE` method on a form will trigger a `POST` request – most likely not something you want. [Kent talks about this limitation in this video](https://www.epicweb.dev/tips/only-use-get-and-post).

As a rule of thumb, consider only using `POST` and `GET` methods on the form.

## Data Modeling Deep Dive

- The prisma client has the _prisma studio_

  ```bash
  npx prisma studio
  ```

  Pretty neat tool to see what you have in your database.

- Kent shows a neat trick related to sqlite and exporting the data to another db.

  - You can run the dump on the sqlite and then use that file to re-create what you have in sqlite in another database.

  - Of course, note that there are differences between databases. Some fields are unique to PostgreSQL and so on.

- **Kent recommends avoiding polymorphisms in schema design**. I do agree.

  - At the start, it might seem like having a model like _file_ and using that to create _image_ which is shared between an user and a note is a good idea. Inventively you hit the case where you start to embed model-specific ids on that _image_ model. The _image_ model is now anemic – it has a lot of optional properties and those point to other different models.

    - To avoid this problem altogether, create separate models that do one thing and are used only in one context.

  - Storing images in the database makes sense for the small scale we are dealing here with.

- When talking about migrations, **Kent mentions the _widen then narrow_ or _expand and contract_ pattern for performing zero downtime migrations**.

  - The basic idea is to expand the application and database schema to allow for all the possible cases, then gradually narrow the application and database schema.

  - This process could pose a lot of challenges, especially if the migrations are not done fast enough. In the worst case scenario, your application stays in the "wide" state because you did not have the time to migrate all the data.

    - This is not bad from the data perspective (apart from having duplicate data, but storage is cheap nowadays). It is bad from maintainability and operations perspective.

- Prisma has this neat feature of _nested writes_ where you can create multiple entities that rely on each other with a single API call (of course, underneath the library is performing a transaction).

  ```js
  await prisma.note.create({
    data: {
      id: "d27a197e",
      title: "Basic Koala Facts",
      content:
        "Koalas are found in the eucalyptus forests of eastern Australia. They have grey fur with a cream-coloured chest, and strong, clawed feet, perfect for living in the branches of trees!",
      ownerId: kody.id,
      images: {
        create: [...],
      },
    },
  });
  ```

  **Many times I wanted similar feature in DynamoDB, where I could reference ID of one item in another operation**. While it is not surfaced to the API, this is what is happening underneath – the `note` gets created and then the `nodeId` is propagated to the `images`. **This propagation of IDs is specific to _nested writes_**. There is the `$transaction` API, but it does not surface this functionality.

  Sadly I could not find any relevant information on how this feature is really implemented.

- The _seed_ script should be idempotent. I've been in situations where that is not the case, and it was a bit of a pain.

  - In addition, Kent ensures that the data is really unique across the whole script lifetime (by creating _unique value enforcer_). Pretty good practice!

- Prisma will, by default, use _lazy initialization_ and connect to the database upon first request to the database. This might be something you do not want to do. If you want to connect eagerly, call the `$connect` method yourself.

- Kent uses a `singleton` function to make the Prisma client a singleton. Not strictly necessary on production, **but very handy during development where HMR will re-require the files. If the Prisma client is not a singleton, you will end up with multiple instances of the client!**

- Nested queries (with the `select` keyword) are pretty awesome.

  ```js
  prisma.user.findUnique({
    where: {
      username: params.username,
    },
    select: {
      name: true,
      notes: { select: { id: true } }, // nested query
    },
  });
  ```

  Having said that, just like in GraphQL, there is a danger that nesting too many levels will cause the query performance to drop. Internally Prisma is doing joins and joins consume CPU and memory.

- Prisma also has a "nested updates" API. The API is just as, if not more, powerful as the nested queries one.

  ```js
  await prisma.note.update({
    select: { id: true },
    where: { id: params.noteId },
    data: {
      title,
      content,
      images: {
        deleteMany: {
          id: { notIn: imageUpdates.map((i) => i.id) },
        },
        create: newImages.map((newImage) => {
          return { ...newImage };
        }),
        updateMany: imageUpdates.map((imageUpdate) => {
          return {
            where: {
              id: imageUpdate.id,
            },
            data: {
              ...imageUpdate,
              id: imageUpdate.blob ? cuid() : imageUpdate.id,
            },
          };
        }),
      },
    },
  });
  ```

  This will create necessary transactions and updates. Of course, it would be wise to see what kind of queries prisma executes.

- Some databases will automatically add an index for non-unique foreign keys. Some will not.

  - If adding such index would be beneficial to your application performance, use the `@@index` keyword in the Prisma DSL.

  - **To know if adding an index is a good idea, use the `EXPLAIN QUERY PLAN` command to see if the database is scanning without an index**. If that is the case, adding index on the column you are filtering/ordering by would make a lot of sense.

- You might want to consider adding **multi-column indexes to speed up your queries as well**.

  - There is an **important thing to consider: the order in which the columns are defined in the index matters!**

    - Consider the following query

      ```sql
        SELECT Note.updatedAt
        FROM Note
        WHERE Note.ownerId = user.id
        ORDER BY Note.updatedAt DESC
        LIMIT 1
      ```

      Without indexes, the database would have to create a temporary B-TREE structure for the `ORDER BY` clause. **If you add an index for (`ownerId`,`updatedAt`) the database will use the index for the `ORDER BY` clause**. If I were to reverse the order of columns, the database would not be able to use the index for the `ORDER BY` clause, as **the leading column is not used in the `WHERE` clause**.

      A good metaphor here is a set of folders. First, you want to go into a folder for a given `ownerId`, then see the first folder in that folder (folders are already sorted since we have an index on `updatedAt`).

### Wrapping up

- While using the `$transaction` API in Prisma is a valid choice, consider nesting the operations. This way prisma will take care of the transactions for you.

- Do not be afraid of writing manual SQL. Sometimes it is necessary.

- Monitoring the queries speed is important. Without it, you would not be able to know if a certain query is slow.

- Consider adding indexes on non-unique foreign keys.

## Web Auth

- **For authentication purposes, strongly consider using cookies instead of `localStorage`**. There are several problems with `localStorage`.

  1. It is not accessible on the server. This means that the initial render of the page has to be either a loading spinner, or a very generic shell which then updates to reflect the signed-in user. Both options are not that great.

  2. It is accessible by extensions and scripts. Image a malicious extension/script which steals the user data. Not fun.

  There are probably more, but those came to my mind first. As you can see, the **security aspect of `localStorage` is not that great**.

- I wonder what is the story behind authorization in loaders/actions.

  - Since each loader/action is like a separate API endpoint, having authorization piece before request makes it to those functions seems crucial to me.

- Kent implements the _theme switch_ by sending a request to a backend and returning a response with a specific `set-cookie` header.

  - Very good idea. Holding the theme in the cookie means that the page already has the necessary classes when it is rendered as the theme is available during SSR.

    - **In Next.js accessing the cookie in RSC will make the whole page dynamic. This might not be what you want**.

  - To add the _system preferences default_ one has to inject a little script into the `head` of the document that checks the preferences and sets the necessary attributes on the HTML. This is pretty much a standard way of doing the dark-mode toggle without the flash of white/dark theme.

- Kent shows how to implement **_cookie flash pattern_** (it has nothing to do with Adobe Flash) for displaying a notification for the user when they delete a note. **First time hearing about this pattern**.

  - When user deletes a note, we want to redirect him to the "notes list" page and display a toast that the note was deleted. **We keep the state of the toast in a cookie scoped to user session**.

    - Alternatives would be a bit harder to implement. It boils down to moving the state from the server to the client that works across different pages and, in the `remix` case, loaders. Cookies are an ideal mechanism for this!

  - **When we read the cookie in the loader, we also unset it, so that the toast only shows once**.

  - The name "flash" comes from the fact that we are using this cookie once to, supposedly, display some kind of information. After that's done, the cookie is no longer of use for us.

- A few recommendations regarding passwords that Kent made (I'm still not convinced about implementing the auth myself, but I think the concepts in this sections are pretty interesting and worth knowing, especially the knowledge around cookies).

  - Do not encrypt passwords. The encryption key might get lost. Instead hash the passwords with a strong hashing algorithm. You should also add salt for each hashing call you perform.

  - Do not store passwords on the same table you store other user data. Moving the passwords to a separate database makes it less likely for someone to accidentally send that data to the frontend when they should not have.

  - Of course, this is a tip of the iceberg. **I still think that creating a custom auth implementation is not a good idea**.

- The `zod` `refine` or `superRefine` transformations are pretty handy when dealing with authentication stuff.

  - In our case, we first check if the form data matches the schema, then we use `superRefine` to verify that the `username` is not taken, and then use the `transform` function to create the user and return the data.

    ```js
    SignupFormSchema.superRefine(async (data, ctx) => {
      // validate that the user does not exist
    }).transform(async (data) => {
      // create the user and return
    });
    ```

    **The `transform` API allows you to validate and return data in a single pass. You cannot do that with `superRefine`**.

- To unset a given cookie, `remix` sets the `expires` attribute to time in the past.

- Creating "authenticated" and "non authenticated" routes requires us to add more code into the _loaders_ and _actions_.

  - While this works, I wonder if there is a better way to do it. I can encapsulate that logic within a function, but it would be nice to have it in one single place. Maybe doing that work in root loader would be a better idea?

    - Sadly [this does not seem to be possible](https://remix.run/docs/en/main/guides/faq#how-can-i-have-a-parent-route-loader-validate-the-user-and-protect-all-child-routes).

- Kent mentions very important detail when implementing the "redirectTo" search param functionality. **Main, you should not trust this param – you should check if it starts with `/`**. Otherwise you risk a malicious actor crafting a specific URL that redirects the user to a 3rd party domain. This would enable the attacker to steal the users credentials.

- While preparing the codebase for the roles and permissions related features, I had to manually update the migration script to "pre seed" the database so that the roles and permissions are already there when application starts up.

  - This could also have been done with some kind of script running on CI before we deploy the application to production.

- Something that I've noticed is that the `loader` and `action` functions code tend to grow quite fast, even if you abstract the blocks of inner logic to separate functions.

- In one of the modules, we have switched from purely cookie-managed sessions, to a combination of a database and cookie-managed sessions.

  - The main benefit of this approach is that **each "sign in" has its own session record in the database. If we delete that record, we will effectively log the user out**.

  - This was not possible with purely cookie-managed solution as we did not have a way to indicate that a given cookie should be considered "expired" before it `expires` property.

    - Instead of holding the `userId` in the cookie, we now hold the `sessionId` instead.

  - The **main drawback** is that some operations, like logging in or logging out now take more time. I think this is a fair price to pay, as users do not necessarily expect the login process to be super quick.

- One pattern that I really like, is the **usage of `zod` for checking if `.env` variables exists and inferring the type of the `.env` schema**.

  - I usually write those types by hand, but I think I will start using `zod` or any kind of validation library for this, just like Kent does it.

  ```ts
  const schema = z.object({
    NODE_ENV: z.enum(["production", "development", "test"] as const),
    SESSION_SECRET: z.string(),
    HONEYPOT_SECRET: z.string(),
    RESEND_API_KEY: z.string(),
  });

  declare global {
    namespace NodeJS {
      interface ProcessEnv extends z.infer<typeof schema> {}
    }
  }
  ```

- The **`fetch` API will throw when there is some kind of network issue, and will return successfully with `ok=false` if the status code is 4xx or 5xx**

  - Always worth keeping in mind. This got me so many times!

- Kent uses the [`close-with-grace` package](https://www.npmjs.com/package/close-with-grace). Pretty neat library.

- Kent uses cookies to communicate between different pages.

  - The flow is as follows: first the user provides the email, then they are redirected to the onboarding flow. Doing this with state might be tricky as we want to keep the user on that onboarding route even if they refresh the page!

  - Seeing how Kent uses cookies, I have to say I've been underusing that API.

- In the TOTP (_time-based one time password_) flow, Kent decided to automatically verify the user if they land on the `/verify` page with certain set of query params. **This this approach sounds nice, it has drawbacks which Kent also mentioned**.

  - For one, **some email clients might "click" the links to check for viruses and spam**. If that is the case, it would not be the user who verifies the email, but a bot. If user were to try to click that link, they would land on a page with an error because the verification was already deleted from the DB.

  - This "automatic" verification might be confusing for users. If I get an email regarding verification, I would expect to land on some kind of "verification page" with the form pre-filled. If we redirect the users automatically, they might be confused about what just happened (at first I was confused as well!)

- When working with images or blobs of data, sometimes you see them encoded as _Data URLs_. For example, a _data url_ for an image could look like this

  ```text
    data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOQAAADkCAYAAACIV4iNAAAAAklEQVR4AewaftIAAAxnSURBVO3BQY4cy5LAQDLR978yR0tfBZCoaineHzezP1hrXeFhrXWNh7XWNR7WWtd4WGtd42GtdY2HtdY1HtZa13hYa13jYa11jYe11jUe1lrXeFhrXeNhrXWNh7XWNR7WWtf44UMqf1PFicpUcaJyUvGbVKaKm6icVEwqU8Wk8kbFicpUMan8TRWfeFhrXeNhrXWNh7XWNX74sopvUvkmlZOKSWWq+ITKicpJxYnKVDGpTBWTyknFScVJxaRyovJNFd+k8k0Pa61rPKy1rvGw1rrGD79M5Y2KN1SmikllqphUJpWpYlKZKiaVNyreUPkmlZOKSeWkYlJ5o+JE5ZtU3qj4TQ9rrWs8rLWu8bDWusYP60jljYoTlROVqWKqmFSmit9UMalMKlPFicqkMlWcVPwveVhrXeNhrXWNh7XWNX74H6PyRsWJylQxqUwVJxUnFZPKGyrfpPJGxW+q+F/2sNa6xsNa6xoPa61r/PDLKv5LVKaKk4pJ5Y2KSeWk4g2VqWJSOamYVKaKSeWkYqqYVCaVqeKbKm7ysNa6xsNa6xoPa61r/PBlKv9SxaRyojJVTCpTxaQyVUwqU8WkMlVMKicqU8UnKiaVqWJSmSomlROVqWJSOVGZKk5Ubvaw1rrGw1rrGg9rrWv88KGKm6icqEwVk8qJyonKGxWfqPhExRsqn1CZKk4qPlHxX/Kw1rrGw1rrGg9rrWv88CGVqWJS+aaKqeJE5TdVTCpTxaRyUnGi8ptUTireqDhR+UTFico3Vfymh7XWNR7WWtd4WGtd44d/rGJSmSpOVKaKqeJE5RMqJypTxYnKVDGpnFR8omJSOVE5qfhExaQyqUwVb1Tc5GGtdY2HtdY1HtZa17A/+CKVqWJSmSomlTcq3lB5o+INlaniRGWqmFSmiknlmypOVKaKSeUTFZPKJypOVL6p4hMPa61rPKy1rvGw1rqG/cEvUvlExYnKScWkclIxqUwVk8pUMalMFZPKScWk8kbFicpU8YbKGxXfpPKbKk5UpopPPKy1rvGw1rrGw1rrGj98mcpU8YbKpDJVTBWTyhsVb6hMFScVJxUnKlPFJ1SmiknljYpJZaqYVE4qJpWpYqo4UZkq3lCZKn7Tw1rrGg9rrWs8rLWu8cOHVKaKSWWqmFSmir9J5aTiEypvVLyhclIxVZxUTCpTxTdVvKEyVUwqb6i8oTJVfNPDWusaD2utazysta7xwz9WcaIyVUwqb1R8QuWkYqo4UZlU3qiYVCaVk4pJZao4UZkqPqEyVUwqk8pUcaIyVZyonKhMFZ94WGtd42GtdY2HtdY17A/+IZWTikllqphUTipOVKaKE5WpYlJ5o2JSOak4UTmpmFSmik [... MORE STUFF]
  ```

  In our case, we encode the QRCode image to a _data url_. Then we send that _data url_ to the frontend.
  You [can read more about data urls here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs).

- **Every time you _commit_ the session, you have to specify when it expires**. If you do not, the default will apply. The default is to have the cookie live only till the end of user session.

  - To "fix" this issue, in one of the exercises, we had to monkey-patch the `commitSession` API. We did this via the `Object.defineProperty`

    ```js
    const originalCommitSession = sessionStorage.commitSession;
    Object.defineProperty(sessionStorage, "commitSession", {
      value: async (...args: Parameters<typeof originalCommitSession>) => {
        const [session, options] = args;

        if (options?.expires) {
          session.set("expires", options.expires);
        }

        if (options?.maxAge) {
          const expiresAt = new Date(Date.now() + options.maxAge * 1_000);
          session.set("expires", expiresAt);
        }

        // We either set it prior, or the cookie already had the `expires` or it is a "session" cookie.
        const expires = session.get("expires");
        return await originalCommitSession(session, {
          ...options,
          expires,
        });
      },
    });
    ```

    **This works because the `sessionStorage` is a singleton**. One could also use the `new Proxy` API here.

- The implementation of the OAuth flow for Github was pretty straightforward.

  - To ensure our application could work offline, we did mock some of the internals of the library we are using when the environment variables start with a certain string. In our case, they had to start with `MOCK_`

    - While the written mocks rely on the internal logic of the library (as only HTTP calls are mocked), I think it is a good idea.

      - Keep in mind that we only mock the APIs when the environment variables start with `MOCK_`. This gives us freedom to isolate set of e2e tests to test the real thing and not bother in others.

      - One could argue that it is not worth mocking. I would also agree. It all depends on the team you are working on and how much "offline-first" development minded others are.

- There are a lot of edge cases to think about when creating a "login with..." feature.

  - What if the user clicks "login with..." and the email the provider account is associated with already exists in our database?

    - We should create a "connection" for the user and let them in.

  - What if the user tries to add 3rd party connection and the email the provider account is associated with another user?

    - We should display a rather vague error.

### Wrapping up

- **Cookies are not only for auth-related state**. We heavily utilized cookies for `redirectTo` parameters.

  - In fact, cookies were an ideal mechanism for this, as they enabled us to share state in-between routes.

  - **When implementing the `redirectTo` feature you should not trust the value of that param. Always sanitize/validate it**. You would not want an attacker crafting URL with `redirectTo` to some kind of 3rd party domain.

- The **_cookie flash_ pattern** is very useful.

  - First, at the root of your application, you set code that reads the cookie that contains "toast payload" (if it exists of course).

    - After getting that information from the cookie, you want to destroy that cookie so that the user does not see the toast multiple times when they refresh the page.

  - Then you pass that information to React, displaying the toast if the payload is defined.

  - Now, in whatever route you want, you set the "toast payload" and redirect the user.

    - When user lands on a given page, a toast is displayed. **This is, IMHO, much better than using URL for the toast state. Using the URL would pollute the URL and make it less readable**.

- Implementing the QRcode and one-time-password is not that scary.

  - Sure, it involves a lot of work, but it is not that super complex.

- The OAuth spec is quite complex. You would be better off using a library for the redirect dance that has to occur.

- You can monkey-patch objects via the `Object.defineProperty`.

- Creating separate models for passwords, connections and verifications (one-time password prompt) is a good idea.

## Web Application Testing

- **I cannot stress enough how handy is the _accessibility_ view for a given element in the dev tools**.

  - How many times have you written the `getByRole` query only to realize it does not select the element you want? In most cases, this was because I provided wrong _role_ parameter.

    - The _accessibility_ tab displays the _role_ and much more information about a given DOM node.

- Fixtures allow you to create temporary resources to be used in the test. I like them better than using the `beforeEach` and `afterEach` hooks for creating the resources as you do not have to create "temporary" variables.

  ```ts
  const test = base.extend<{
    insertNewUser(): Promise<{
      id: string;
      name: string | null;
      username: string;
    }>;
  }>({
    insertNewUser: async ({}, runTheTest) => {
      const userData = createUser();

      await runTheTest(async () => {
        const newUser = await prisma.user.create({
          data: userData,
          select: { id: true, name: true, username: true },
        });
        return newUser;
      });

      await prisma.user.deleteMany({ where: { username: userData.username } });
    },
  });

  const expect = test.expect;
  ```

  Contrast this with the `beforeEach` and `afterEach`

  ```ts
    let user;

    beforeEach(async () => {
      user = ...
    })

    afterEach(async () => {
      if (user) {
        await deleteUser(user)
      }
    })
  ```

- To communicate between different processes, Kent uses the file system.

  - The process A writes to the file system.

  - The process B (_Playwright_) reads from the file system.

- It is **imperative to keep the stdout/stderr clean while running tests**.

  - Kent mentions how hard it was to develop tests while working in a codebase where running them polluted the logs. I can definitely relate.

    - Apart from the difficult of adding new tests, having **unnecessary** logs in the test output is pretty demoralizing. I find it works on the same basis as the famous "broken window" – since one is broken, why not break another and another. With that mindset you end up in a codebase where nobody cares about having clean test logs.

  - For logs in particular, you might want to use spies for the console. **Make sure to reset the spy after you assert on it**.

    - The resetting part is most likely best done in various before/after hooks the testing library you use provides.

    ```ts
    // You probably want all this code in some kind of setup file.
    const originalConsoleError = console.error;
    let consoleError: SpyInstance<Parameters<typeof console.error>>;

    beforeEach(() => {
      consoleError = vi.spyOn(console, "error").mockImplementation((...args) => {
        originalConsoleError(...args);
        // This will cause the test to fail if one does not mock the error.
        throw new Error("`console.error` called. If you expect this to happen, mock the console.error");
      });
    });

    // The vitest and jest can automatically restore mocks so we do not have to call `mockRestore` here.
    afterEach(() => {});
    ```

- In the section about _component testing_ Kent mentions three kind of users

  1. The end user
  2. The developer user
  3. The "test user"

  **You want to avoid testing like "test user" would use your app**. Who is the "test user"? The "test user" is knowledgeable about the internals of the components of your app. **If you test like "test user" you start to test implementation details end users do not care about**. This is a bad place to be since **it is the end users who are brining the money, not the "test user"**.

  Note on the _developer user_. Those users are also very important. Here, we ensure that the components are easy to use. If they are, other developers are more likely to ship faster.

- Testing react hooks in isolation has many benefits, especially if those hooks are critical to your app functionality. But there is also a drawback – **testing react hooks in isolation facilities testing implementation details**.

  - As discussed earlier, testing implementation details, in most cases, is not ideal. Of course, sometimes that is the right thing to do.

  - Kent proposes an alternative – a **concept called _a test component_**. So instead of testing the hook itself, you create a "stub" component which uses the hook API. Think of it as writing a "story", but not for a component, but rather for the hook.

  ```tsx
  function TestComponent() {
    const [defaultPrevented, setDefaultPrevented] = useState<"idle" | "no" | "yes">("idle");

    const dc = useDoubleCheck();

    return (
      <div>
        <output>Default Prevented: {defaultPrevented}</output>
        <button
          {...dc.getButtonProps({
            onClick: (e) => setDefaultPrevented(e.defaultPrevented ? "yes" : "no"),
          })}
        >
          {dc.doubleCheck ? "You sure?" : "Click me"}
        </button>
      </div>
    );
  }
  ```

  In the test, you assert on the output HTML rather than on `result.current` as it is the case with `renderHook` function. This solution is also far from ideal. I personally thing there is a room for both approaches, but if I were to choose, I would choose the _test component_ approach.

- I was very pleasantly surprised that `remix` exposes the `createRemixStub` function. This function will setup all the necessary context providers so that you can test the component/route that utilizes _loader_ or any other API.

  ```tsx
  const user = createFakeUser();

  // from @remix-run/testing. Pretty sweet!
  const App = createRemixStub([
    {
      path: "/users/:username",
      Component: UsernameRoute,
      loader: async () => {
        return json({
          user,
          userJoinedDisplay: user.createdAt.toLocaleDateString(),
        });
      },
    },
  ]);

  await render(<App initialEntries={[`/users/${user.username}`]} />, {
    wrapper: ({ children }) => (
      <AuthenticityTokenProvider token="test-csrf-token">{children}</AuthenticityTokenProvider>
    ),
  });
  ```

- When testing authenticated requests via `vitest` I could not shake the feeling that the amount of code we have to write is quite big.

  - I think it would be easier to test those loaders/components in `playwright` where we can authenticate users much more easily.

- The `set-cookie` header is for creating a cookie. This is where you specify all the properties that the cookie should have, like `httpOnly` and other stuff.

  - The `cookie` header is for passing the cookie value. Here, we simulate the browser behavior where the browser includes the `cookie` in each request we make to the backend.

- Creating separate database for the tests is the way to go.

  - Luckily for us, we are using SQlite as a database. This means we can use a `.db` file. This means we can create the file when test are starting.

    - Another way to ensure isolation is to introduce some kind of segregation of data in your application. Most applications are _tenancy-based_ meaning the data is separated between tenants. If that is the case, you just put the test data for one test into randomly generated tenant (like `userId`).

      - If that is NOT the case, it might be worth introducing tenancy if only for the sake of testing. Believe me, it is very hard to write tests on singleton resources.

  - We can take things further and create a `.db` file per worker the test framework is using. Usually, testing frameworks are parallelizing tests in separate processes. To avoid collisions and data contention, you might want to create an _instance_ of the database per the worker.

    - In `vitest` we can use the `VITEST_POOL_ID` environment variable.

      ```js
      const databaseFile = `./tests/prisma/data.${process.env.VITEST_POOL_ID ?? 0}.db`;

      const databasePath = path.join(process.cwd(), databaseFile);
      ```

  - In addition, one could use `globalSetup` where they create a "golden database" which is copied for each worker. No need for seeding each time worker runs!

### Wrapping up

- Testing in Playwright or Cypress is awesome.

  - It is a shame that Kent did not showcase the _component testing_ capabilities of those tools.

    - I feel like we had to write a lot of setup code (albeit that could be abstracted) for integration tests of the components. Having said that, **it is awesome that remix exposes the `createRemixStub` function**. It would be so nice for Next.js to do this as well!

- With `vitest` being 1.0 it seems to be a defacto testing tool now.

- Custom assertions are worth your time.

- Creating some kind of segregation, be it at database, or data-level is essential in testing. Without that, you will not be able to have parallel tests going.

## Overall

- Very good workshop. I've learned a ton.

  - Especially around cookies.

- The workshops are centered around `remix`. As such, sometimes I felt like I'm learning more about the framework than the web itself.

- The **forms workshop is very worth your time**.

  - I'm glad that Kent spent some time on accessibility. Tying the error messages with the form fields, and doing it in a way that is accessible is very important skill to have.

    - As a reminder, you specify the `aria-describedby` and `aria-invalid` on the input, and then `id` which points to `aria-describedby` on the wrapped that displays the error.

      - Consider using the `useId` for the ids.

  - The "trick" with specifying the validation-related HTML attributes and setting `noValidate={true}` was new to me.

    - The deal is that, we want to validate in JS to have custom error popups, but we want the screen-readers to read out to users the constraint for a given field. Of course, sometimes it is not necessary to define them in plain HTML, but in most cases that is possible.

  - Resetting the form via `type="reset"` button was new to me as well. TIL!
