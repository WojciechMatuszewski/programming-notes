# AWS Event Driven Architecture

## What does event driven mean?

Normally, aps are _coupled_ when it comes to _domain knowledge_. This usually means that the reliability is reduced.

### Command vs Event vs Query

- **event**s does not require responses. Usually contains _metadata_. Usually propagates only the data that changed.

* **command** is something specific that has to with a given application

## Event Notification

- you can think of **event stores as topics**

* this architecture is more decoupled

- there might be problems with large amount if subscribers

* there maybe problems with state tracking, events flying everywhere

- schema changes can cause unknown problems (you do not really know the amount of subscribes)

## Event Carried State Transfer

Instead of publishing changes only, now events contain all fields that are in a given _domain_. This is not a perfect solution though. Sometimes your services will require to keep event data in internal data store.

- very reliable

* very decoupled

- **redundancy**. You have to think about eventual consistency.

## Event Sourcing

You save events , all of them, so that you can _travel back in time_ if needed. This is useful for analytics.

- GDPR can be an issue

* storing events forever (potentially) can result in higher costs.

- this solution is traffic heavy.

**Snapshoting** can help with replayability.

## CQRS

This one consist of having multiple sources of truth for reading and writing. Data is organized in such a way to be efficient for a given operation (read / write). You still have event stores but these sources of truth listen to them.

- this solution is quite complex

* allows you to scale reads and writes independently

## Bounded Context

Events are published within a given context. This allows you to make assumptions that everyone in a given context understands that event.

If you want to make sure events can be understood between contexts you introduce anti corruption layer. It's job is to translate the event data between contexts

- include event version, your subscribes should probably work with multiple versions

* **Event Bridge** can help you with this. You can generate event schema there just by providing events emitted by your system. **Event Bridge is built on top of Cloud Watch Events architecture**, but they are adding new features and capabilities like transformation of your events.

## Duplicates

- systems have to be idempotent

* remember that using **aws sdk will retry by default**.

## Observability

- implement central logging service

* emit custom metrics to track delivery / reception and latency

### Summary

Event Notification pattern is implemented the most because it's the easiest. But remember that it all boils down to requirements.
Do not implement CQRS just because it's fancy. Think about **benefit vs effort**.

You can use XRay with SQS SNS and such, nice!
