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

- Consider how a traditional diagram for the components described above compares to a diagram where the _client_ is a _mobile_ application.

  - You do not really need a web server because all the assets are already on the device.

  - You _still_ need some kind of server to handle HTTP requests.

- **Your diagram might look totally different than Jem's diagram because we did not specify what we are building**.

  - Any system needs a well-defined _scope_. Without it, you are just guessing.

## Chapter 2 – how to build anything

- **To be effective, you need to have the ability to _understand_ and _translate_ business requirements to a given architecture**.

  - Building, for the sake of building, is NOT the way to go.

- When you have an understanding of the problem, you can start on the API and architecture design.

- At one point, we were tasked with designing a TODO application. You might think that it's pretty easy, given that, in essence, you need to support creating, deleting, updating, and listing todos. **But do you, really?**

  - Asking questions is the most powerful thing you can do!

  > What core features do we need to support?

- Do not make assumptions. **Assumptions are dangerous because they are heavily biased toward your current knowledge**.

  - Always ask questions. Question the intent, the need, the use case. Of course, do it in a way that's not annoying or offensive.

## Chapter 3 – understanding the problem

- Functional requirements: what are we trying to solve? _What is the core functionality?_

  - **To fully understand the requirements, try to repeat what you've heard to the other person**. This ensures you are both aligned on what you are building.

- Non-functional requirements: How should the system perform?

  - **Use non-functional requirements to probe for engineering challenges**.

- **Designing _around_ a _core entity_ will make your architecture more robust**.

  - Focus on what matters. For example, for a banking application, you most likely want to design around a _transaction_.

- In most cases, applications perform basic CRUD operations. Designing around that functionality can help you avoid thinking about things that do not matter.

### CAP Theorem

- CAP Theorem explains how choosing to prioritize one aspect of the system, such as _consistency_, influences others, such as _availability_ or _partition tolerance_.

  - _Availability_: Every request receives a response (success or error).

  - _Consistency_: Every read receives the most recent write or error.

  - _Partition Tolerance_: The system continues to operate even if messages are delayed or lost (such as network issues).

  **Note**: In any distributed system, **network partitions will happen** as the transport layer is not 100% error-proof. So, in reality, you are choosing between CP and AP systems.

### System Quality

- There are **many** things you have to consider when designing a system:

  - Observability

  - Security

  - Scalability

  - Adaptability. Think about making a one-way-door decision that makes the system less malleable.

  - Performance

  Prioritizing one will influence the other. **It's trade-offs all the way down**.

### Non-Functional Requirements

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

  - You scale _horizontally_ when you add more machines. This is more involved, as you will most likely have to add some kind of load balancing between those machines. This change _might_ require you to modify your application code.

    - This solution allows you to scale almost infinitely, but it comes with other drawbacks, such as introducing potential eventual consistency.

  By default, teams will scale _vertically_ since it is easier to do. Once they hit the limit, they will start scaling _horizontally_.

## Chapter 5 – data storage

> At the end of the day, most of what you do is reading from and writing to a database.

### Dimensions

- Structured vs. unstructured data

  - _Structured data_ has relationships.

  - _Unstructured data_ is like the `JSON` column in PostgreSQL.

- Persisted vs. ephemeral

- Read-optimized vs. write-optimized

  - Most applications are read-heavy rather than write-heavy.

- Consistency vs. availability

  - If you shard your data, you sacrifice consistency but gain availability.

### Data Storage

- Relational

  - SQL.

  - Structured data with relationships.

  - Enforced schema.

  - ACID transactions.

  - Traditional databases such as PostgreSQL were not built to scale horizontally.

    - This means that scaling those to meet the demand might be more challenging.

- Non-relational

  - NoSQL.

  - Semi-structured or unstructured data.

  - Flexible. There are usually no schemas.

  - These types of databases usually scale horizontally quite well.

    - Most of the non-relational databases have sharding built-in.

    - It is much easier for those databases to scale _horizontally_ since they do not have to worry about ACID transactions.

### Database Scaling

- Two approaches: **partitioning** and **sharding**.

  - _Partitioning_ is when you create, **within the same database**, multiple tables that contain the same data but are _partitioned_ by some kind of heuristic.

    - For example, you might have **multiple `user` tables that are partitioned based on the `user_id` column**.

    - **_Partitioning_ is easy because it only requires changes in your code**. You can scale your database _vertically_ to meet the demand for new tables.

  - _Sharding_ is when you create multiple _databases_ that contain the data partitioned based on a "shard key".

    - For example, you might have **multiple databases holding the `user` data sharded on the `MD5(first_name)` value**.

    - **_Sharding_ is quite hard because it requires you to deploy new infrastructure**.

- _Sharding_ is quite hard because you can **over-index on a particular shard**, making the data "imbalanced".

  - If you have most of your data on one database, you eliminate the benefit of sharding.

  - Understanding your data is key.

### Availability

> What happens when things go wrong?

- Backups are key component of every architecture.

  - Primary/replica architecture.

    - You write to primary, then the write is _replicated_ to the replicas. You can read from replicas.

      - This introduces eventual consistency since replicas might not have the most up-to-date data yet.

  - Primary/primary architecture.

    - You have multiple primary instances that you can read or write to.

      - How do you reconcile conflicts between different primary databases?

  - Peer to Peer

- There are **various strategies to consider**:

  - _Transactional_: The primary won't yield back until all the replicas acknowledge the write.

    - This will increase the latency.

    - What if replicas fail?

  - _Snapshot_: You can take period snapshots of the primary and reconcile replicas.

    - Definitely a good approach for making sure the system is resilient.

    - Snapshots can be _iterative_ or "whole". The _iterative_ store only diffs between A and B points in time.

  - _Merge_: Each database can have its own separate backup. Then, all the separate backups are combined into one "main" backup.

### Caching

- **Client-based cache is the fastest possible cache you can have**.

  - Yes, Redis is fast, but consider that you have to make a network request to get the data.

- At pretty much every layer of the stack, including the transport protocol, there are ways to cache the data:

  - The HTTP cache on the transport level.

  - The client cache on the browser level.

  - The in-memory/disc cache on the server level.

  - The query cache on the database level.

  - The CDN cache on the assets/transport level.

#### Various approaches to caching

- The **cache aside (lazy loading)** pattern (your application code does a lot of work):

  When reading:

  1. Check the cache.
  2. If you have the data, return the result.
  3. If not, read from the DB, update the cache, return the result.

  When writing:

  1. Write to the cache AND to the database at the same time.

- The **write-through cache** pattern (cache does a lot of work):

  When reading:

  1. Read from the cache.
  2. On miss, read from DB.
  3. Write to cache.

  When writing:

  1. Write to the cache.
  2. Write to the database.

- The **write behind** pattern (the cache does a lot of work):

  When reading:

  1. Read from cache.

  When writing:

  1. Write to cache.
  2. Return.
  3. Async write to database by the cache.

#### Caching Tradeoffs

- The biggest tradeoff here is **performance vs. freshness**.

  - The data in cache might be stale. Is that okay? It depends on your needs.

#### Cache Invalidation

- Time-based expiration (TTL).

- Event-based.

- Version tagging.

  - Think invalidating the cache based on the _version_ of the data. CDNs often do this.

- Refresh ahead.

  - This one is quite similar to "stale-while-revalidate" pattern.

  - The cache will request new resources just before the old ones expire.

    - Useful for mobile applications!

### Estimations

The main point in estimating RPS or storage requirements is to **validate your architecture**.

If you can agree, that you will need to save X amount of GBs a month in terms of data, you might choose one database or another.

**Probe for scale required, not exact numbers**.

Start Day 2 Part 3 31:46
