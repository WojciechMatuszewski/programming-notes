# FEM Fullstack System Design

## Chapter 1 – everything is a system

- Every system has _boundaries_.

  - Systems interact with each other across those _boundaries_.

- Systems are rarely simple. They are usually _distributed_ and quite complex.

### Common System Components

1. The Client. _Someone_ or _something_ has to send the requests.
2. The Database. That local variable you have? That might count as a database as well.
3. The Server. Everything runs on some kind of server.
4. The Load Balancer. Even if you use Serverless, you are using Load Balancers, albeit indirectly.
5. The Cache. In-memory cache is still a cache!

- Consider a traditional diagram for components described above, and how it compares to a diagram where the _client_ is a _mobile_ application.

  - You do not really need a web server, because all the assets are already on the device.

  - You _still_ need some kind of server for serving HTTP requests.

- **Your diagram might look totally different than Jem's diagram because we did not specify what we are building**.

  - Any system needs a well-defined _scope_. Without it, you are just guessing.

## Chapter 2 – how to build anything

- **To be effective, you need to have the ability to _understand_ and _translate_ business requirements to a given architecture**.

  - Building, for the sake of building, is NOT the way to go.

- When you have understanding of the problem, you can start on the API and architecture design.

- At one point, we were tasked with designing a TODO application. You might think that it's pretty easy, given that, in essence, you need to support: creating, deleting, updating and listing todos. **But do you, really?**

  - Asking questions is the most powerful thing you can do!

  > What core features do we need to support?

- Do not make assumptions. **Assumptions are dangerous as they are very biased towards your current knowledge**.

  - Always ask questions. Question the intent, the need, the use case. Of course, do it in a way that's not annoying or offensive.

## Chapter 3 – understanding the problem

Day 1 part 2 41:17
