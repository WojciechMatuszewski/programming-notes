# Frontend Masters Next.js Fundamentals v4

- The _route groups_ allows you to use a separate layouts for different pages **without having to create a different URL for those pages**.

  - Think: "I need these routes to share a common layout, but I do not want to change the URL structure".

  - Be mindful of "route collisions" when using route groups.

    - Since the URL does not change, you might, by accident, override a `page.tsx` for that url by creating `page.tsx` inside a "route group" directory.

- CSS in JS does not work that well in the context of SRCs.

  - Oftentimes there is a lot of setup involved to get all the generated styles from a given page and put them in the "head" during SSR.

  - Using CSS classes (perhaps via TailwindCSS) allows you to write your components without having to worry about all of that.

- In dev mode, **static pages are not compiled unless you visit them**.

  - The compilation step might take some time, so your website might feel sluggish whenever you work on it locally (at least at first).

- _Server actions_ **implicitly create an HTTP endpoint**. You have to be careful here – _think about how this influences your security posture_. There is also a question about rate-limiting.

  - Next.js has [dedicated section talking about security](https://nextjs.org/docs/app/building-your-application/data-fetching/server-actions-and-mutations#security) in the context of server actions.

- Scott mentions that Next.js will _split_ the "server code" and the "client code" into separate bundles.

  - This aligns with what I've read about – that there are different "environments" React runs in.

    - Library authors have to configure their `package.json` `exports` field correctly to hint to the webpack plugin which bundle to use for which environment.

      - A good example of this is the [`package.json` file in `React`](https://github.com/facebook/react/blob/2980f27779cf37a9656b25418a3c5cfca989e244/packages/react-dom/package.json#L51). Notice different "default" and the "react-server" values.

      - Another good example is the [`server-only` package](https://www.npmjs.com/package/server-only).

        - Go into the "code" tab and check the `package.json` file!

- **DAL** stands for _data access layer_. Not to be confused with _server actions_ or _server components_.

  - You might be tempted to use `server-only` import in the file you have these defined. **This will not work, as the `server-only` needs "react-server" environment to function**.

    - I personally do not put any pragma directives in those files.

- Next.js definitely has come a long way when it comes to caching.

  - In the course, we are using a setting called "dynamicIO". You can read more about it [https://nextjs.org/docs/app/api-reference/config/next-config-js/dynamicIO](here).

    - The `use cache` directive allows you to cache a page that fetches data on the server. To be honest, I find this API quite confusing...

      - You can also use `Suspense` without the `use cache` to signal to Next.js that the page is _dynamic_ and should NOT be cached.

  - One can't also forget about the `cache` function `React` exports.

    - **The `use cache` is for caching across requests**.

      - It allows you to cache JSX or non-serializable data. But that data WILL not be part of the cache key.

    - **The `cache` function is per-request cache**. You might think of this as memoization of results for a function that gets wiped as soon as request finishes.

      - By "request" I mean the request for a given page.

    - **If you want to check if the cache is working properly, especially the `cache` function, consider building your application**.

      - When running in dev-mode, React might invoke your function multiple times.

  In the course, we cached the `getIssues` function.

  ```tsx
  async function getIssues() {
    "use cache";
    cacheTag("issues");
  }
  ```

  If you want to revalidate the cache, you would have to call `revalidateTag` in _server action_.

- Getting granular with `Suspense` has its benefits.

  - Push the suspense boundary as close to the data it fetches. This way, the rest of the JSX could be rendered statically, and you stream only the part that needs to wait for a network response.

- The auto-generated endpoints for _server actions_ are quite awkward to work with if the request fails.

  - Usually, you would expect the endpoint to return 4xx or 5xx with a proper error message. Instead, **the docs point you to return error information as the result of the _server action_**.

- The API routes are pretty awesome. They work with "regular" `Request` and `Response` objects.

  - You can use `NextRequest` and `NextResponse` objects to simplify your code a bit.

- **You can use _edge runtime_ for your "page" or "layout" components**. [Here are the docs](https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#runtime).

  - Due to limited memory and CPU limitations, setting `runtime = "edge"` for a given page might not work that well.
