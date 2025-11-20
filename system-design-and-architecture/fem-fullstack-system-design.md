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

- Functional requirements: what are we trying to solve? _What is the core functionality?_

  - **To fully understand the requirements, try to repeat what you've heard to the other person**. This makes sure you are both aligned on what we are building _for_.

- Non-functional requirements: how the system should perform?

  - **Use non-functional requirements to probe for engineering challenges**.

- **Designing _around_ a _core entity_ will make your architecture more robust**.

  - Focus on what matters. For example, for a banking application, you most likely want to design around a _transaction_.

- In most cases, applications perform basic CRUD operations. Designing around that functionality can help you avoid thinking about things that does not matter.

### CAP Theorem

- CAP Theorem is about how choosing to prioritize one aspect of the system, like _consistency_ influences the other, like _availability_ or _partition tolerance_.

  - _Availability_: Every request receives a response (success or error).

  - _Consistency_: Every read receives the most recent write or error.

  - _Partition Tolerance_: The system continues to operate even if messages are delayed or lost (think network issues).

  **Note**: In any distributed system, **network partitions will happen** as the transport layer is not 100% error-proof. So, in reality, you are choosing between CP and AP systems.

### System Quality

- There are **many** things you have to consider when designing a system:

  - Observability

  - Security

  - Scalability

  - Adaptability. Think making a one-way door decision that makes the system less malleable.

  - Performance

  Prioritizing one will influence the other. **It's trade-offs all the way down**.

### Non-Functional requirements
