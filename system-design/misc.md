# System Design Misc

## About real-time updates

Applications can communicate via HTTP or GRPC or any other means. How can we make sure that the client can receive data while being at "standby" to do so?

This is relevant in many cases: messaging apps, stock trading platforms, notification services. Anywhere where you expect the system to provide updates to the customers over period of time.

### Polling (also known as short polling)

This is where the application _continuously_ makes requests to the other application to update its state and the data it has. The requests are usually made on some interval, for example every 5 seconds.

**The main benefit of this approach** is that it is **very straightforward** to implement. You do not have to deploy any new architecture or tools. All you need to do is to write some kind of loop in the code, and, within that loop, make a request to the other application.

**The main drawback of this approach** is that, **the more clients perform the pooling, the higher the chance for the server to get overloaded**. Imagine a chat application with millions of users. If each connected client would periodically made requests to the backend systems, the server could get overloaded.

**Another drawback** is the fact that **a lot of requests might return the same data indicating that nothing has changed on the server**. This means that we should not have made this request in the first place, but there is no way for us to know that!

### Long Polling

Similar approach to the polling, **but instead of responding immediately, the server will "hold" onto the request and either return after certain timeout or return whenever there was an update to the data**. This approach is more performant than the _short polling_.

The **main drawback** of this approach that **creating and holding onto connections is resource intensive**. Like in the case of _short polling_, we might encounter scaling issues.

The **main benefit** of this approach is that **it is widely supported by browsers and servers**. The implementation is not that involved and relatively easy to pull off. This **could be a good start to implement given requirement, but we definitely can do better**.

### WebSockets

This approach **will establish a persistent bi-directional connection between server and the client**. Both the server and the client can send events to each other (not applicable in all cases, but very valuable in some).

The **main drawback** of this approach is **resource waste if there is no traffic over the connection**. Imagine having millions of clients connected to a given server and waiting there for updates. Those updates might never come, and during that time, the server had to manage and keep those connections in memory!

**Another drawback (but in some cases it is actually the behavior you want)** is the fact that **when the connection is interrupted, you will need to re-connect yourself**. Depending on the situation, this actually might be very beneficial. Imagine a situation where multiple clients just lost the connection and they automatically re-connect. They will most likely reconnect at roughly the same time **causing the thundering herd problem**.

The **main benefit** is the **bi-directional nature** of this architecture. This characteristics is quite powerful. Think of all the multiplayer games or software that enables collaboration!

**Another benefit** is the fact that **you send the headers only once – when establishing the connection**. This means that subsequent exchange of data does not involve sending any request-specific metadata. This means that **the requests are pretty fast, faster than sending them through HTTP**.

### Server Side Events

This approach **will establish a persistent, uni-directional connection between server and the client**.

The **main drawback**, just like in the case of WebSockets, is the **fact that maintaining those connections takes a lot of resources**. If the connection is not utilized frequently, we should not keep it.

**Another drawback (or a positive, depending on the situation)** is the fact that **SSE will automatically re-connect clients**. As I mentioned before, this might or might not be what you want.

The **main benefit** is **speed as you do not have to send headers with each request**.
