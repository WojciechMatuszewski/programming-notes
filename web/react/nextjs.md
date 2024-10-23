# Next.js

## Why you can't set a cookie in RSC

It all as to do with streaming. The RSC model fully embraces streaming the chunks of HTML over the wire. **This concept is tied very closely with `Suspense`**.

**Next.js wont let you use the `cookies().set` in RSC because at that time the request headers were already sent to the client**. Since the server implements streaming, the first response from the server contains all the necessary headers related to streaming for the browser to understand next payloads.

You should consider using _server actions_, _route handlers_ or _middleware_ as the place where you set cookies as there the response is NOT streamed from the server.

> [Here is a great video explaining this "issue" in more detail](https://www.youtube.com/watch?v=ejO8V5vt-7I)

## `Suspense` in Pages router

- **According to my testing**, the `Suspense` will be triggered on the server, and **Next.js will WAIT for all `Suspense` boundaries to finish fetching before returning the HTML**.

  - If you are not careful, this could mean that your initial page load is slower than it needs to be.

  - It seems like the alternative might be fetching _only_ the necessary data via `getServerSideProps` and then rendering any components that use "Suspense" on the client.

    - You can render a component on the client via `dynamic` API.

## `Suspense` in App router

- The **`App` router supports streaming**. This means that **Next.js will NOT wait for all `Suspense` boundaries to finish fetching before returning the HTML**.

  - This is great, as it allows us to use `Suspense` without worrying about slowing down the initial page load.
