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

# Feedback for our solution

## Message Processing vs Stream Processing

### Message Processing

- one message contains everything we need

* built-in DLQ

### Stream Processing

- think click stream: you need A LOT of messages to get some value out of it

* no DLQ

## Feedback

- would start with SNS and SQS and maybe Event Bridge, this would allow scalability

* for ordering maybe we could try **SNS FIFO**. Currently it's on preview. With **SNS FIFO** you can push to **SQS FIFO** which would solve the problem with ordering and would allow you to drop DynamoDB.

- **SQS has MessageGroupID attribute**. This is very similar to Kinesis PartitionID.

* SNS and SQS solution is scalable and you pay only for what you use.

- with SNS and SQS you could either send messages to another SQS queue and provide retention there (14 days) or stream the messages into S3. There is an example within serverless repo.

* we can utilize batching to save on SQS costs (pay per request)

- with SQS and SNS the latency would be better. Considering the Kinesis shard-subscribed limitations which would require you to create pipeline of streams (lambda to kinesis pushing to other kinesis streams) the SQS and SNS seems better.

* Job Ads Stream can be replaced with SNS FIFO, that would probably be the most cost efficient solution. Again, no manual scaling.

- the whole Atlas Topic and Ads Queue can be replaced with Event Bridge. But, as of now there is no ordering (planned for the future). The cost is roughly the same as SQS + SNS.

* **MessageGroupID integration works different than lambda PartitionID**. There is a hidden pooling service between SQS and your lambda. This pooling service invokes your lambda function. At start there are **5 poolers** which scale by number of 60.

- binlog: not just academic, not enough info, have not seen it yet. Enrichment has to be done manually, requires a lot of effort.

* FIFO SQS and retries: when message is dropped to DLQ you probably have to process that message out of order

- **Event Fork Pipeline** for replayability.

* **Athena** can be used for pulling data from S3. Using timestamps and local jobAdID we can express the order.

- when it comes to GDPR you could use Athena for extracting.

## Misc

- scaling control for Kinesis (concurrent invocations for a given shard?)

* **serverless application lens**

- there is a free **architecting serverless course**

# Day 2

## Setup

E-commerce, not really fancy, had to operate in several countries. On prem, monolithic. DB was Oracle. That db was a central place for a lot of domains.
No TDD , dedicated test teams. Overall pretty bad :c

## Challenges

- coupling

* testing

- basically monolith which is not really HA

## Goals

Replacing the monolith, going for microservices (they used Event Carried State Transfer).

### Loading Data

Using VPN connection and so called _exporters_. This is basically ACL for the single source of truth.
_Exporter_ was fetching on the basis of version for a given product (there was a versioning table). If the version changed the real query was dispatched. Then that data was exported into kinesis (multiple subscribers).

There were some problems with Stock data (static and dynamic from a CRM). Decision was to create 2 topics for that, one from the DB one from the CRM (to the same subscriber).

### Why Kinesis?

- they did not want to go with Kafka because the infrastructure was not really event-sourcing like (no need to reply).

* SQS and SNS was not considered due to ordering constrains (fifo did not exist back then).

### Kinesis Insights

- number of subscribers could be an issue (remember up to 5 consumers and 5 reads/s). There is enchanted fan-out for Kinesis, might be interesting to look into.

### Deletion

Since the events are hold in the separate dbs the decision to delete the data has to be on per consumer basis.
