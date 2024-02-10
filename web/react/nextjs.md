# Next.js

## Why you can't set a cookie in RSC

It all as to do with streaming. The RSC model fully embraces streaming the chunks of HTML over the wire. **This concept is tied very closely with `Suspense`**.

**Next.js wont let you use the `cookies().set` in RSC because at that time the request headers were already sent to the client**. Since the server implements streaming, the first response from the server contains all the necessary headers related to streaming for the browser to understand next payloads.

You should consider using _server actions_, _route handlers_ or _middleware_ as the place where you set cookies as there the response is NOT streamed from the server.

> [Here is a great video explaining this "issue" in more detail](https://www.youtube.com/watch?v=ejO8V5vt-7I)
