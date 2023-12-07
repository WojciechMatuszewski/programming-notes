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

### Wrapping up

- It looks to me that using `noValidate` but still providing those HTML-based validation attributes is the way to go. It makes sense as it allows assistive technologies to read out validation spec for the end user, while also allowing us to have our custom "look and feel" for error messages.

  - You can even go further and add it only when React hydrated the page.

- To properly associate the HTML input element with an error message, use the `aria-describedby="some_id"` and `aria-invalid={true}`.

- First time seeing the concept of a _honeypot_ in action. Very useful pattern.

- The CSRF token makes sense. It prevents an attacker _forging_ a request to our site from a 3rd party domain.

  - Keep in mind that CORS is not enough here. Some request might classify as "simple request". For those, the browser will not send the OPTIONS (preflight) request.

  - Encrypting the CSRF token is a very good idea.

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

Now: Change email (106)
Before: https://nolanlawson.com/2023/12/02/lets-learn-how-modern-javascript-frameworks-work-by-building-one/?utm_source=stefanjudis
