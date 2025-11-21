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

- Just as with _functional requirements_, to create a list of _non-functional requirements_, you **must ask questions** about the system you are building.

  - For _non-functional requirements_, focus more on questions about "implementation details" rather than functionality.

  > How many MAUs?

  > How many transactions per second?

  > What is the desired P95 latency?

  > Are there any certifications, such as HIPAA, GDPR, or SOC2, that we need to consider?

  > How long do we need to store backups?

- Once you have a set of questions, you can start formulating _non-functional requirements_.

  > The system should support 1,000,000 MAUs.

  > The system should support 100 transactions per second.

  > The target P95 latency is 100ms.

## Chapter 4 – high-level design

- Jem talked about **starting with the entities and branching out from there**.

  - We used a similar strategy when discussing _functional_ and _non-functional_ requirements, focusing on identifying the "core" entity of the application.

  - I agree with this approach. You should narrow your focus as much as possible, especially in high-stress situations like job interviews.

- We briefly touched on various methods of communication between the "server" and the "client":

  - gRPC for service-to-service communication.

  - Server-Sent Events for real-time updates that do not need bi-directionality.

  - WebSockets for real-time updates that need bi-directionality.

  - GraphQL for stitching multiple data sources together and exposing them as one cohesive API.

  - REST for anything else. A safe default.

  **Those are only suggestions.** In the real world, it all depends on various factors.

- Next, we touched on _vertical_ vs. _horizontal_ scaling.

  - You scale _vertically_ when you add more resources to a single machine. This is quite "easy" because you do not have to change your code in any way to take advantage of it.

    - Of course, the drawback is that you have a single point of failure, and you will eventually hit a limit on how high you can scale.

  - You scale _horizontally_ when you add more machines. This is more involved, as you will most likely have to add some kind of load balancing between those machines. This change _might_ require you to modify the code of your applications.

    - This solution allows you to scale almost infinitely, but it comes with other drawbacks, such as introducing potential eventual consistency, and so on.

  By default, people will scale _vertically_ since it's easier to do. When they hit the limit, they will start scaling _horizontally_.

Start Day 2 Part 1
