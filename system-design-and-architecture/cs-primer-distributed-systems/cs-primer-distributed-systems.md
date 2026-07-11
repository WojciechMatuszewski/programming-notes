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

Start Part 2
