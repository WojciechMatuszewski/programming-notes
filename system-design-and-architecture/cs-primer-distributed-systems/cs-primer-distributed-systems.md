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

Start lesson 8
