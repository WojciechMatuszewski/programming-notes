# Neetcode.io System Design Interview

> Learning and building things from [this course](https://neetcode.io/courses/system-design-interview/3).

## Design a Rate Limiter

- The easiest way to implement rate limiting is to **implement the logic inside the application**.

  - That is why you **might want to clarify if the rate limiter is for a single microservice or the whole backend**.

    - If it is for single micro-service, why not go with the _in-code_ solution?

    - If not, then you will have to implement it as a shared service, especially since one API might be distributed across different machines.

      > I would immediately reach for APIGW here, or at least ask if adding a gateway in front of the API is possible. Having said that, one has to remember that rate-limiting per user is not that great in APIGW.
      > Keep in mind that the **APIGW usage plans require API keys to work**, and the **number of API keys you can create is limited!**.

### Non-functional requirements

- The **latency is very important**. You need something that will be quick. We do not want to add unnecessary latency to each request.

- The **scalability is also vital**. You want the rate limiter to scale horizontally.

- The **storage is worth considering**. Every application might have different rate limiting rules.

- A **crucial aspect is the availability and how the system behaves if the rate limiter goes down**.

  - I would argue that, for most businesses, it makes sense to **fail open** – allow the app to function as if never happened. Yes, some people might perform too many requests, but your application should function as expected.

  - There is also **failing closed** – returning an error to the user. I'm not a fan of such behavior.

### Implementation

- The rate limiter might act as a reverse proxy.

  - If the user is rate limited, the request should (?) never arrive into a given service.

  - It does overload the concept of a rate limiter to me a bit. Should not the routing be a separate service?

- The notion of _sticky sessions_ is **crucial** if we scale the **persistence layer of the rate limiter**.

  - Imagine accessing different storage nodes with different information for the same user for each request. A nightmare!

### Thinking in AWS

- The easiest solution, most likely, would be to **leverage the APIGW rate limiting capabilities**. But there are some caveats to that.

  1. Keep in mind that APIGW "only" supports 10k requests/s. This is a soft limit.

  2. Keep in mind that to **have the rate limiting applied on the tenant-level** one would have to **use usage plans and api keys**. There is a **limit on how many api-keys we can create**.

- As for my own implementation. I see two ways we can go about this.

  1. Use the **APIGW authorizer as a rate limiter**. Retrofitting the authorizer to act as a proxy could work. Instead of checking the IAM and the token (though that could also be possible), we could deny the access to the API based on the amount of requests a given user made. For the persistance layer, I would use DynamoDB with DAX. Most likely two tables, one for rules (DAX) and one for the request count (without DAX).

  2. Use CloudFront as the proxy? A wildcard (no idea if that would work at all), but we could use the CloudFront as the proxy which would talk to the persistance layer.

  3. Use **WAF rate limiting capabilities**. It has the ability to rate limit based on the IP address, but **there is no way to configure the algorithm is uses to compute the count of requests**. An advantage here is that one can use WAF with both APIGW and ALB.

## Design a Link shortener

- The idea is simple, map a string into another string, but shorter.

- You most likely want to expire those links.

### Non-functional requirements

- High availability (well, duh!)

- Low latency

  - Optimize for reads

  - There might really be a LOT of reads. **While thinking about the architecture, also take the cost into the consideration**.

### Implementation

- Since we do not need atomic operations, a NoSQL database could be the right choice here.

- Due to the large amount of reads, one **could consider adding caching in-between the client and the storage layer**.

  - With the cache in place, the eviction algorithm is quite important. The teacher suggest going with LRU – [_Least Recently Used_](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)

    > Keeping and computing what is the _least recently used_ item is quite expensive.

- Since the implementation relies on redirects, **you must know what is the difference between `301` and `302` status codes**.

  - The `301` status code is for **permanent moves, it causes the browser to CACHE the end location**.

  - The `302` status code is for **temporary moves, the browser will not cache the request**.

    - Imagine the case where you want to apply analytics on the backend. Then you most likely want to know about all the redirects happening. In such case, the `302` status code could be a better choice. Otherwise, `301` is the way to go.

### Thinking in AWS

- **S3 can act as a "redirect" engine**.

  - This, I would call it, a niche feature, is pretty much all we want.

    1. It handles deletion of old objects via _lifecycle rules_. There is **a limit of how many _expiration rules_ one can set inside a given bucket / day**. The limit is **very low – 100/day**, but there are workarounds. [See this stackoverflow thread for more ideas](https://stackoverflow.com/questions/12185879/s3-per-object-expiry).

    2. Enables you to redirect the user to a given destination.

    3. Handles a lot of traffic.

    4. Eliminates the need for the database.

    **Keep in mind that, you will pay for each S3 GET request**, but that cost would also be there for any other solution.

- Use **APIGW + AWS Lambda + DynamoDB**.

  - Before the S3 redirecting was possible, people used to compute the shortened URL using AWS Lambda, and save those on DynamoDB.

  - Here, the scale is also quite big, but since there are more "pieces of the puzzle", the maintenance cost will be higher.

  - **Keep in mind that DynamoDB TTL could have a delay up to 48 hours**.

    - At least this one is documented. I could not find the information about the S3 expiry delay.

  - The latency will be much higher than the S3 version, since there are a lot of services involved.

## Design Twitter

- Read heavy system. This points to a usage of Redis or any other caches.

### Requirements

- Following others.

- Creating tweets. With images and videos.

- Viewing "the feed". This one is particularly interesting due to the sheer amount of users on the platform.

### Non-functional requirements

- Most people will be viewing tweets, not creating them. This means a lot of reads.

- Since the tweets could be "media rich", the storage capacity of the system needs to be huge.

  - This usually points to the fact that we do NOT need strong consistency.

    - As a sidenote, I think you do not need strong consistency in almost all cases.

### Implementation

- Consider using NoSQL for the tweets, and some kind of Graph Database for the "follow" functionality.

  - This will enable you to scale better. Keeping everything in one DB (even sharded) might not be a good solution due to the amount of data.

    - When talking about the relational DB here, consider sharding based on userId. You can have a hash function compute the correct shard for a given user.

- Object storage for the media. **Remember about the CDN**.

  - Here the teacher talks about the different ways to populate data on the CDN.

    - You can have the "pull-based" CDN (sometimes called _reverse proxy_), where the asset is cached only after requested.

    - You can have the "push-based" CDN, where, as soon as the asset is uploaded, it is also uploaded to the CDN.

    - In this particular case, **consider the "pull-based" CDN approach**. Pushing everything into the CDN might create a lot of overhead.

- The way to distribute tweets is to use a queue with a worker pool that creates a feed for a given user and saves it to the cache.

  - There might be a problem with this approach for people with huge amount of followers. You would need to re-compute the feed for all people who are following that one person.

    - **Insert the celebrity tweet at runtime**. Do NOT fan-out when the celebrity tweets. Instead fetch that tweet when the user requests their feed.

    - Do NOT compute the feed for inactive users.

### Thinking in AWS

- The CDN would be CloudFront. I think, by default, it uses the "reverse-proxy/pull" mechanism.

- The Servers would most likely be some kind of containers. Using AWS Lambdas at that scale does not make sense.

  - I would pick Fargate or ECS here.

- The Cache would be ElastiCache. Keep in mind that you can use Memcached or Redis under the hood.

- For Object Storage one cannot go wrong with S3.

- To compute the feed, I would use SQS with a combination of Fargate of ECS. Please note that you have to manually pull messages, you have to set that in your code!

  - If you need a fan-out, I propose SNS, as it supports millions of subscribers.

## Design Discord

### Functional Requirements

- Servers

- Channels

  - The most interesting functionality is the "show me the first unread message" when navigating to a given channel.

### Implementation

- Receiving messages via WebSockets or Polling. Polling is not a great solution as it creates a lot of traffic.

  - You could also try using _Server Sent Events_ here, but please note that these are uni-directional.

    - You [can read more about SSE here](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events).

- Shard the database based on the `channelId`. There will be more channels than servers.

- Ordering based on the date the message was sent is quite important. The `sent_at` should be the index.

- To derive the "mentions" and "you have x unread messages" use a second table depicting the user activity.

  - Each time the message mentions the user, update the user activity table with that information.

  - You also need to keep track of the `last_read_at` for a given user + channel combination.

### Thinking in AWS

- For the WebSockets, I would definitely use the IoT Core MQTT.

  - It supports huge scale, is serverless and has powerful filtering capabilities.

- For the database, I would reach for DynamoDB.

  - We can group the user activity collections nicely.

- To asynchronously update the user activity for a given message, one could use DynamoDB streams with AWS Lambda integration.

  - This will work for small scale, but for a high traffic, I would push the DynamoDB messages to Kinesis.

    - **Kinesis with AWS Lambda has a much higher throughput than DynamoDB Stream + AWS Lambda subscriber**.

- Should we even cache the messages for a given channel?

  - The cache would be updated very frequently since there could be a lot of messages posted every second.

    - I do not have a great answer to this question. "It depends" :p

## Design Youtube

### Non-functional requirements

- Reliability of the data.

  - A single video could be watched by a massive amount of people.

- Favour availability instead of consistency.

  - Given the amount of data, using strong-consistency is not feasible.

### Implementation

- Videos stored in the Object Storage. To achieve high availability, those should be replicated across AZs.

- Video metadata in a NoSQL database. There are no relations so it makes sense.

  - Even if you have to "join" some data, you could de-normalize it and allow some duplication.

- Encoding should be async. Use a message queue to push the "this video needs to be encoded" to the "encoding" service.

  - We will have multiple video artifacts. One is the "raw" video, one is the "encoded" video.

    - To distribute the videos, we will use a CDN.

- To improve latency, load small chunks of the video. There is no need to send the user the whole video.

  - **Use TCP for the streaming. The UDP is a good choice for live-streaming, but for a video that is already encoded, TCP is a better choice**.

    - We should favour reliability. You do not want to have any "gaps" in the video.

    - For live-streams, you the video feed to always display the "freshest" chunks, so it makes sense to sacrifice a couple of seconds of missing video to stay "up-to-date" with the stream.

### Thinking in AWS

- The CDN is CloudFront

- The Object Store is S3

- The mechanism to encode videos could use the _Elemental_ suite of offerings from AWS.

  - The `MediaConvert` service integrates with S3. Perfect fit! We do not have to have any queues in the architecture.

- For the metadata store, one could use DynamoDB

- To upload the video, one could use AWS S3 presigned URLs.

  - **To implement resumability, consider using the _multipart upload_**.

    - This is not so easy to do, as you might need to create a presigned URL for each part.

## Design Google Maps

### Functional requirements

- Navigate from source to destination in the shortest amount of time.

  - This does not necessarily mean that the route is the shortest.

- Track user location.

- Get the ETA.

### Non-functional requirements

- Accuracy.

- It is okay to have some latency.

- Reliability is important. You do not want the app crash / data to be missing when user is navigating.

### Spatial Indexing

Before we dive into the implementation, one has to understand the concept of **_spatial indexing_**.

**_Spatial indexing_ is the process of diving the area in squares**. With area divided into N squares, you can **attach a number, or some kind of hash to that square**. If you have that, you can address a given square pretty easy, as each "area" starts with the same hash, then you can recursively traverse the hash to pinpoint a given square. **Attaching hashes to the squares on the map is called _Geohashing_**.

It seems like, in the world of Amazon, DynamoDB is pushed as the database to hold the geo-related data.

- [The dynamodb-geo library](https://github.com/amazon-archives/dynamodb-geo).

- [Series of articles on this subject](https://aws.amazon.com/blogs/mobile/geo-library-for-amazon-dynamodb-part-1-table-structure/).

### Implementation

- Two services – the _Route Service_ and the _Location Service_.

  - The _Route Service_ reads from the _Location Service_. This allows to compute the best route given the traffic.

- The _Location Service_ is a write-heavy system. You will need some kind of time-series based database for the entires.

- To view the map, we have to serve images to the user. Since returning an image for the entire world is not feasible, the service would have to return multiple images for the frontend to stich them together.

  - Serve images via the CDN.

### Thinking in AWS

- For the geo-data one could use DynamoDB.

  - If you need a graph-based database, consider [Amazon Neptune](https://aws.amazon.com/neptune/).

- The CDN is obviously CloudFront.

- For the _Location Service_, I would use [Amazon Timestream](https://aws.amazon.com/timestream/).

  - For ingestion, consider Kinesis.

## Design a Key-Value Store

- The _isolation_ is the norm in SQL databases.

  - **_Isolation_ means that if two transactions were dispatched at the same time, they will appear as if they happened right after each other**.

    - This means that some transactions might wait for others to finish.

    - You can achieve the same thing in DynamoDB with `ConditionExpression` to check if the value you are trying to read/update is what you think it is.

      - This will require a read of the item before writing, so there might be some inconsistencies.

      - DynamoDB uses _last writer-wins_ strategy when replicating data.

- For indexing, most databases use some kind of tree-structure, like B-Tree or other.

  - There is also the LSM Tree used in databases like Cassandra. This one is optimized for writes.

- When talking about replication, we have to be mindful of consistency.

  - Here, the famous CAP theorem comes into play – it's two out of three things _Consistency_, _Availability_ and _Partition_.

    - Most databases sacrifice on consistency. I would argue that availability is the most important thing.

  - There **is also the concept of the Quorum**. It dictates how many nodes have to be in-sync for the write/read to be accepted.

- When talking about partitioning, the hashing algorithm is imperative.

  - Ideally, you would be able to spread the data evenly.

  - Even better, if we **virtualize the nodes into vnodes. This should help if one of the nodes goes down**.

    - If we do not virtualize the nodes, and one node goes down, a single node might be suddenly taking on much more writes.

      - With virtual nodes, even though one node is down, the granularity of vnodes allow us to spread the writes more evenly across available nodes.

## Design Message Queue

- _Subscriber_ and _consumer_

- _Publisher_ and _producers_

- _At least once delivery_ is the way to go. It is very hard to guarantee only single delivery due to possible network issues.

- The video describes an architecture with a _metadata store_ where we keep track of how many topics received a given message (one message could go into multiple topics).

  - With this knowledge, we could delete the message from the database if all the topics received a given message.

- One also has to consider whether the system is a _push_ or _pull_ based.

  - In the _pull_ based system, the subscriber might over-poll the messages. Also, _pull_ based systems are usually slower as the polling happens only so often. Having said that, the polling allows for batching which is a great optimization.

  - In the _pull_ based system, the subscriber will never over-poll as it processed the messages as they come in. **This system is not suitable for batching, but the batching could be done on the publisher side (sending multiple messages in one single payload)**.

    - In AWS, **SQS is a _pull_ based and SNS is a _push_ based system**.
