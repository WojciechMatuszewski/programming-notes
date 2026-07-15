# CS Primer – Distributed Systems

> Learning [this material](https://csprimer.com/courses/distributed-systems/)

## KV store introduction

- Did you know you can [create _indexes_](https://redis.io/docs/latest/develop/clients/patterns/indexes/) in Redis?

  Consider a scenario where you save user payments. You _might_ need a query for "give me a payment by ID", but you might also need a query for "give me all payment for a given user".

  If we step back a bit, and ignore the fact that this pattern might be better served by a relational database, to make those two access pattern work, you might need to add a secondary index to Redis!

  I've also seen people use `SADD` command to append multiple elements to the same set.

### Notes during implementation

- I was re-introduced to the _double dispatch_ pattern, which really helps with avoiding any kind of "type switches".

  If you return a generic interface from a function, you either have to narrow it down (by the "type switch") OR make it so that this interface can "self-dispatch".

  In Golang, this looks like the following:

  ```go
  // Without double-dispatch
    switch command := command.(type) {
      case GetCommand:
        value, err := handler.HandleGet(command)
      case SetCommand:
        err := handler.HandleSet(command)
    }

  // With double-dispatch
  command.Execute(handler)

  // Then, inside of each command
  func (gc GetCommand) Execute(handler) {
    handler.HandleGet(gc)
  }
  ```

## Introduction to distributed systems

- We want to build systems that are:

  - Scalable.

  - Reliable.

  - Maintainable.

  If you do not have one of those three, you will be in trouble.

- To build 👆, we have the following tools (non-exhaustive list):

  - **Replication**. This basically means **putting the same thing into multiple places**.

  - **Partitioning**. This basically means **putting something that was in one place, and splitting it**.

- When designing distributed systems, **you will need to make tradeoffs**. The axis on which you case base your tradeoffs on:

  - **Consistency**: if I make a write, and then read, will that read contain the most-up-to-date data?

  - **Availability**: would you rather return with "success" to the client, even when data is not consistent, or would you rather fail the request?

    This is interesting. Consider why you are adding replicas – most likely to _increase_ reliability. So if you choose to fail, if the write failed to propagate, why do you have the replicas in the first place?

    But it's not all-or-nothing. You can have _synchronous_ replicas mixed with _async_ replicas.

  - **Complexity**. The more complex the system is, the harder it is to maintain.

  All of this relates to the CAP theorem.

- During the webinar, Oz mentioned **the pitfall of using auto-incrementing IDs in the context of distributed systems**.

  If you are making writes to multiple locations independently, and both are working out of the same ID, you will be in huge trouble.

  **This means that if your system uses auto-incrementing ids now, you might never be able to partition it**. Quite problematic!

- One thing that I started to pay more attention to is **whether the migration will cause a lock on the table an how that influences the system**.

  Some tables are rarely accessed, so that is not a problem, but some might be critical, and adding a lock on them might be problematic.

## Wire Formats

- Text-based wire formats. Think `JSON`, `CSV` or `XML`.

  **If we look those through the lens of how easy it is to evolve those formats, the main problem is lack of `schema` integration**.

  Ideally, the wire format would be versioned, or have some kind of well-known schema, so you can "update" or "downgrade" it as you see fit.

  We also have to think about how _efficient_ it is to encode, decode those formats and also how _large_ those can get.

- Sometimes you might be tempted to use `base64` for encoding non-text stuff and sending it via text-based wire formats.

  While this works, it should be a signal to you, that perhaps there is a better wire format to use for this particular data.

- Using `gzip` (or any other compression mechanism) **is a very good way to reduce the size of messages you send over the wire**.

  Remember that browsers will automatically de-code `gzip` for you!

- Binary-based formats. Think `Proto Buffers`, `Avro`.

  Those are usually very efficient at encoding and decoding. Consider how easy it is to "jump" to specific field: you know the offset in bytes!

  They also **have schema integration built-in as first-class concern**. This means it's much easier to evolve the API.

## API Patterns – REST, SOAP, RPC, GraphQL

- Conceptually, the "contract" between clients is unrelated to the wire format.

- Interestingly, SOAP was built to operate over _any_ protocol, such as HTTP, SMTP, TCP or UDP.

- These approaches can be compared by the interface they expose, although SOAP is a messaging protocol and can also carry RPC:

  - REST exposes resources through a uniform interface. The "uniform interface" are the HTTP verbs. For example, the `GET` is consistent across all resources.

  - SOAP services expose operations on messages.

  - RPC exposes procedures, or "functions".

  - GraphQL exposes a typed schema whose fields clients select and traverse.

- RPC is _really_ good at creating an _exact_ interface between services. It is not built in flexibility in mind.

  Note: RPC is often associated with ProtoBuffers or other binary format. While you _can_ use RPC with this wire format, nothing stops you from implementing RPC via HTTP (think `tRPC` or _server functions_ in FE world).

## Replication

- _Replication_ at it's core means copying the same thing into multiple places.

  - To reduce latency: the closer the data is to the client, the faster the response will be. CDNs are great at this.

  - To increase availability.

  - To increase _read_ throughput. You can read from multiple sources. The load is distributed across multiple nodes. Scaling _write_ throughout requires partitioning which is much harder to do than replication.

- Be mindful of _horizontal_ vs. _vertical_ scaling.

  - _Vertical_ is usually easier to achieve, because you most likely do not need to change anything in your application, but it has a hard cap of how large the machine you use can be. You "just" update the configuration for the machine you use.

  - _Horizontal_ is usually harder to achieve, as it might require changes in your application, but it's much more maintainable.

- When deploying _replication_, you need to think about:

  - **Replication lag**: how long does it take for the writes to replicate?

    You can make it so that all writes have to synchronously replicate to _all_ readers before you return with successful write. **This will greatly reduce your write throughput**.

    You can accept that read-after-write might be stale. That all writes will _eventually_ propagate to all nodes.

    Or you can make it so that _some_ replicas need to acknowledge the write before returning to the client upon writing.

  - **How you replicate**. This can go pretty deep into Database mechanics, but one way might be _statement-based replication_.

    But what if the outcome of this statement is non-deterministic? For example IDs that should be the same across multiple replicas, but are created internally? **For some setups, the statement-based replication is the answer, for others, not so much**.

    **You could look at WAL, and replicate based on that**. The problem here is that the structure of entries are fixed, and it might be hard to keep the service running while you upgrade your database.

  - **Replication topology**: Do you have multiple writers trying to replicate to multiple readers?

Start replication.
