# CS Primer – Distributed Systems

> Learning [this material](https://csprimer.com/courses/distributed-systems/)

## KV store introduction

- Did you know you can [create _indexes_](https://redis.io/docs/latest/develop/clients/patterns/indexes/) in Redis?

  Consider a scenario where you save user payments. You _might_ need a query for "give me a payment by ID", but you might also need a query for "give me all payment for a given user".

  If we step back a bit, and ignore the fact that this pattern might be better served by a relational database, to make those two access pattern work, you might need to add a secondary index to Redis!

  I've also seen people use `SADD` command to append multiple elements to the same set.

### Milestones

- Server which echos over UDP.

- Server which sets on SET and GET.

- Test with two clients.

Finished part 1 37:48
