# System Design Misc

## About real-time updates

Applications can communicate via HTTP or GRPC or any other means. How can we make sure that the client can receive data
while being at "standby" to do so?

This is relevant in many cases: messaging apps, stock trading platforms, notification services. Anywhere where you
expect the system to provide updates to the customers over a period of time.

### Polling (also known as short polling)

This is where the application _continuously_ makes requests to the other application to update its state and the data it
has. The requests are usually made on some interval, for example, every 5 seconds.

**The main benefit of this approach** is that it is **very straightforward** to implement. You do not have to deploy any
new architecture or tools. All you need to do is to write some kind of loop in the code, and, within that loop, make a
request to the other application.

**The main drawback of this approach** is that, **the more clients perform the pooling, the higher the chance for the
server to get overloaded**. Imagine a chat application with millions of users. If each connected client
periodically makes requests to the backend systems, the server could get overloaded.

**Another drawback** is the fact that **a lot of requests might return the same data indicating that nothing has changed
on the server**. This means that we should not have made this request in the first place, but there is no way for us to
know that!

### Long Polling

Similar approach to the polling, **but instead of responding immediately, the server will "hold" onto the request and
either return after certain timeout or return whenever there was an update to the data**.
This approach is more performant than the _short polling_.

The **main drawback** of this approach that **creating and holding onto connections is resource intensive**.
Like in the case of _short polling_, we might encounter scaling issues.

The **main benefit** of this approach is that **it is widely supported by browsers and servers**.
The implementation is not that involved and relatively easy to pull off.
This **could be a good start to implement given requirement, but we definitely can do better**.

### WebSockets

This approach **will establish a persistent bidirectional connection between server and the client**. Both the server
and the client can send events to each other (not applicable in all cases, but very valuable in some).

The **main drawback** of this approach is **resource waste if there is no traffic over the connection**. Imagine having
millions of clients connected to a given server and waiting there for updates. Those updates might never come, and
during that time, the server had to manage and keep those connections in memory!

**Another drawback (but in some cases it is actually the behavior you want)** is the fact that **when the connection is
interrupted, you will need to re-connect yourself**. Depending on the situation, this actually might be very beneficial.
Imagine a situation where multiple clients just lost the connection and they automatically re-connect. They will most
likely reconnect at roughly the same time **causing the thundering herd problem**.

The **main benefit** is the **bidirectional nature** of this architecture. This characteristic is quite powerful.
Think of all the multiplayer games or software that enables collaboration!

**Another benefit** is the fact that **you send the headers only once – when establishing the connection**. This means
that additional exchange of data does not involve sending any request-specific metadata. This means that **the requests
are pretty fast, faster than sending them through HTTP**.

### Server Side Events (SSE)

This approach **will establish a persistent, uni-directional connection between the server and the client**. The client
will not be able to send any data to the server.

The **main drawback**, just like in the case of WebSockets, is the **fact that maintaining those connections takes a lot
of resources**. If the connection is not used frequently, we should not keep it.

**Another drawback (or a positive, depending on the situation)** is the fact that **SSE will automatically re-connect
clients**. As I mentioned before, this might or might not be what you want.

The **main benefit** is **speed as you do not have to send headers with each request**.
In addition, you most likely do not need a library to handle all the complexity SSE transport. Unlike WebSockets,
the setup to make this work is not that involved.

## Mutex

Mutex is a construct that makes sure **only a single "consumer" can access a given piece of state/data/resource (or even a combination of operations)**.

The caller must acquire a so-called "lock". **Only one caller can hold the lock at a given time**. All other callers have to wait for the caller to "free" the lock to act upon the data (again, only one of them will acquire the lock in the end).

## RWLock vs. Mutex

Mutex is quite restrictive in how it works. If one caller acquires the lock, all others have to wait, no matter what operation they want to perform (read or write).

**RWLock allows for concurrent reads as long as the write lock is not acquired by any caller**.

As soon as one caller acquires the _write lock_, all callers have to wait for that lock to be "freed". **This also includes callers who "just" want to read data**.

## R.A.D.I.O.S framework

- Requirements

- High-Level Architecture

- Data Model

- Interfaces and API Design

- Optimizations, Observability

- Security and A11Y

**Whatever you are doing, make sure you think about the "sad paths" as well!**

### Requirements

Functional and non-functional

- **Functional** describe what the system **must do**. For example displaying a list of items.

- **Non-functional** describe how well the system must fulfil functional requirements or under what constraints. For example, the list should appear under 200ms for the user.

### High-Level Architecture and Data Model

Here is where you can lay-out components and how they communicate with each other.

Is component A communicating with component B? Or perhaps they should communicate via some kind of "shared" bridge?

Things like that.

### Data Model

This is where you design _how_ the data would look like.

Is the data normalized? Perhaps you use a tree-like structure?

### Interfaces and API Design

Frontend applications communicate with backends. What kind of API should we use?

Should we use GraphQL? Why?

Should we use REST? Why?

How should we structure endpoints?

Perhaps a BFF pattern?

Long-pooling? WebSockets? SSEs?

For example, we could choose a Polling / Long Polling approach.

- Easy to set-up. Most changes on the frontend, you only need an endpoint on the backend.

- Might be wasteful since the data might not change when you request it.

- It will definitely take longer for changes to appear in the UI (depending on the polling interval).

WebSockets are hard to scale properly, but allow for bi-directional communication with FE. Do you need it? Perhaps SSE would be easier, since it's "less stateful" and uses HTTP instead of TCP?

SSE also require state since you need to have a "handle" for the connected client, but since you can't receive messages from the client, the code is usually less involved.

### Performance

There is a lot to it. You have a lot of knobs to turn.

- Using `Accept-Encoding` headers from FE. This will help if the server supports compression and responds with compressed data. If that's the case, the browser will automatically de-compress responses for you!

- CND's. Some CDNs compress the data on the edge, so you do not even have to worry about that on the server side. Having said that, you will most likely incur a storage cost here.

- Batching requests.

- Client-side cache.

- Using HTTP2 for multiplexing or different protocols like WebSockets.

- Bundling.

- Lazy loading resources.

- Virtualization.

- Deferring non-critical resources.

- Talking about metrics like _time-to-interactive_ or _first-contentful-paint_.

- Tree shaking.

- Rendering techniques, like SSR, Server Components.

- Using different formats for images (usually taken care of by the CDN).

Things like that...

### Observability

You need to know if you application is working or not.

- Error Monitoring tools like Sentry or Datadog.

- User tracking tools.

- Performance Monitoring tools like Sentry or Datadog.

- Tools that allow you to see application logs.

Things like that...

### Security

- Proper CORS (Cross-Origin Resource Sharing) headers.

  - CORS is all about telling the browser which cross-origin resources it can "expose" to the client JavaScript.

    - **Backend responds with CORS headers, the browser enforces them**.

- Rate limiting in place.

- **Content-Security-Policy** headers which **tell the browser which sources of content are allowed and which kind of actions are permitted**.

  - For example, you can allow loading scripts only from certain domains.

    **This is a good protective measure against cross-side-scripting (XSS)**.

- DDOS prevention. Usually handled by Cloudflare or other vendors.

### A11Y

Browser renders HTML, so the HTML must be semantically sound.

It's hard to read websites that use tiny fonts, so those muse be adjusted accordingly.

It is much faster to use a website using keyboard, so we need to make sure that's possible.

Some users might be color blind, so using specific contrasts between colors might be required.

Things like that...
