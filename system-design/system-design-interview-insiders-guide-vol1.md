# System Design Interview - An Insider's Guide: Volume 1

My notes while reading the book. Random thoughts and such.

I decided to give the book a read since I like the "bigger picture" when architecting solutions. I often question how people on my team are aware of all the edge cases. By reading this book, I hope to understand how "generic" systems are built.

## Chapter 1: Scale from zero to millions of users

Great introductory chapter. It touched on many "core" mechanics that govern most (I hope) applications that we use daily.

The author talked about:

- **Scaling methods (vertical / horizontal)**. Both for the web tier and the data tier.

- Using load balancers for scaling and disaster-recovery purposes.

- Caching, be it locally via some cache instance or CDN.

- Deploying to multi-region environments and all the challenges that come with that.

- **Data sharding**.

## Chapter 2: Back-of-the-envelope estimation

This chapter introduces units of scale and how to interpret them.
In my humble opinion, it is practical because every engineer should possess this knowledge.

I liked the brief Twitter example and estimating QPS (Query per second) numbers.

Also, it does put into perspective what 99.99% availability SLA means – scary!

## Chapter 3: A framework for system design interviews

This chapter is a guide on how to behave during system design interviews. Not relevant for me at this point in my career, but a good read nevertheless.

As is often the case in the interview process, it boils down to **asking questions** and **communicating early and often**. The worst thing you can do is to stay silent within your head.

As long as you communicate and focus on the higher-level problems, you should be good!

## Chapter 4: Design a rate limiter

This chapter contains a lot of helpful information. I love it!

First off, the rate-limiting might not only be implemented on the client or the server – we might want to introduce rate-limiting on different layers of the networking stack! There is no way I would have thought about it.

There are so many different ways to handle rate-limiting I was not aware (unsurprisingly) that there were so many different algorithms. I'm most familiar with the **_token bucket_** algorithm that AWS uses for most, if not all, of its services.

Some notes from the architecture part itself:

- Rate limiting implementation as a middleware between the client and the server.

- Redis acts as a data layer for the rate-limiting rules and keeps the state of the logic. Redis has been a prevalent choice in this book so far.

- Rate limiting services might run in a distributed environment. If that is the case, the architecture must take **rate conditions** and **synchronization** issues into account.

- For resiliency, you might need to **make your rate limiter _eventually consistent_** which feels a bit weird to me.

## Chapter 5: Design consistent hashing

This chapter is hefty on the theory side but important.

**Consistent hashing algorithms power data/load distribution across many different servers** — something we take for granted these days.

The chapter touches on the theory behind **_ring buffers_** and their role in load distribution. Interestingly, the first, quite sophisticated implementation was insufficient at providing good load distribution.

The key takeaway is this: **consistent hashing allows you to overcome the "celebrity" load problem**. The consistent hashing mechanisms are used in various systems, not only databases or load balancers (Discord uses consistent hashing too).

## Chapter 6: Design a key-value store

The chapter starts with the single-server approach – suitable for meager traffic and storage requirements. The implementation might be an in-memory hash-table.

Since the characteristics of the single-server architecture do not fit most production systems, one has to consider building a distributed one. This is where most, if not all, complexity comes from.

The author begins with the **CAP theorem – a classical thesis about the _consistency_, _availability_, and _partition tolerance_**. The takeaway is that you cannot have all of them. You can only pick two.

- CA systems are not practical since network failures are common in the real world.
- CP systems sacrifice availability.
- AP systems sacrifice consistency. I think these are the most common.

Since there is no one-fits-all solution, you as an engineer are responsible for understanding which model fits your needs.

Next, the author looks at data partitioning and describes the previously mentioned **_hash ring_** in more detail. The _hash ring_ is a prevalent technique to ensure uniform data spread.

Because of data replication needs, the system might or might not be strongly consistent. It all depends on the implementation. **The _coordinator_ component (in AWS land, would that be _the router_?)** is responsible for reading/writing to nodes.

In a distributed word, one also must think about race conditions and data inconsistencies. The author suggests using **_versioning_ and \_vector clock_s** – a fascinating technique that keeps the changes in the form of vectors. Diffing the vectors allows the system to know that there is a conflict and resolve the conflict accordingly.

Handling outages and failures differ depending on the severity of the failure. The first step is to know that the failure occurred – usually done by using the **_gossip protocol_**.

Other nodes in the system handle temporary failures.

- Permanent failures require data syncing and diffing – Merkle trees come to the rescue in this situation.
- Data center outages are usually taken care of via DNS and switching between different data centers.

Zooming in on the node itself, it is vital that it is independent of the other nodes, meaning it can carry out all the necessary functions – just like our single-server approach did.

## Chapter 7: Design a unique ID generator in distributed systems

As was the case with the previous chapters, the author starts with a single-server solution and deems it unfit. I get it as using auto-incrementing ID in a single database in a distributed compute context sounds like a bad idea.

I find it fascinating that there are many approaches to tackling this issue.

- One could **use _multi-master replication_ with the database auto-increment feature**. The increment "tick" would be different in each database. This approach has elasticity issues. Imagine adding or removing a node. You would have to recalculate the IDs.

- Another approach is to **use UUIDs**. Nodes generating those are entirely independent as the probability of the collision is very low. This approach is not without its problems – the UUIDs are usually 128 bits long, are not numeric, and do not "go up" with time.

- Another approach is to use the **_ticket server_**. A technique introduced by Flicker. Funny enough, this approach uses the auto-increment feature of a given database. The ticket server works in a multi-node configuration with data synchronization. Pulling this architecture off is not a small feat!

- A good approach for numeric IDs is the **Twitter snowflake approach** (the name is somewhat relevant to our culture). We split the ID into five parts, each with a different meaning – just like the IP packet?

What follows is a deep dive into the **Twitter snowflake approach**. As with any ID generation algorithm, there is some period after which we expect the collision to happen. In this case, it's 69 years.

Since part of the **Twitter snowflake approach** is computed based on a timestamp, one has to ensure that **all the machines involved in the ID generation have their clocks synchronized**. I'm not even surprised that [AWS has a service for this – the Time Sync Service](https://aws.amazon.com/about-aws/whats-new/2017/11/introducing-the-amazon-time-sync-service/)

## Chapter 8: Design a URL shortener

The author starts with an API design, first tacking the POST and GET routes.

What follows is an elaboration on the differences between the 301 and 302 status codes.

- 301 status code is a **permanent redirect** meaning that **the browser caches the response** and will NOT invoke the API in subsequent requests. I was not aware of that. Good to know. This status code should be **used when reducing the server load is essential**.

- 302 status code is a **temporary redirect**, meaning that **the browser will NOT cache the response**. This is **useful when tracking is essential**.

We would not be able to deploy the service and operate it efficiently without a good hashing function. The hashing function is responsible for "truncating" the long URL into a shorter one. As for the data storage layer – the author recommends a hash table **for small systems**. I'm not sure why the author did not pick a key:value store for real-life scenarios but instead decided to use a relational database.

What follows is a description of different hashing techniques.

- The **hash + collision resolution** solution uses a well-known hashing algorithm and detects the collisions by directly reaching the database. Not very efficient.

- Another approach is the **base 62 conversion**, not based on the link but a unique ID. Quite fascinating. There is no way I would have come up with this idea alone. I would most likely try to generate the short URL based on the long URL.

The author decided to go with the classic data layer and the caching layer frontend with an API for the architecture piece. The database and caching layers are widespread in this book.

## Chapter 9: Design a web crawler

Building a web crawler can be an excellent exercise in using concurrent features of your language of choice. One might think that creating a web crawler is relatively straightforward – you visit the page, collect URLs and move to the next page. As we will see, that is not the case.
Do you want only to support HTML or maybe additional content types like `.pdf` or `.jpeg` files?

- How do you handle malicious links?

- How would you ensure that you do not make too many requests to a given website?

- How do you parse invalid HTML? Remember – browsers are permissive when it comes to HTML syntax.

- How do you handle duplicate content? There is a lot of duplicate content out there on the web.

When traversing web pages and saving the links, you must choose between **DFS and BFS** algorithms (because the natural structure of the web is a graph). Since the **depth of the tree might be very large, you should favor the BFS algorithm**.

- Naive BFS implementation might lead to problems, like overflowing a given page with requests. The author suggests adding webpage ranking and the **_URL frontier_** to the mix to ensure our web crawler behaves like a good bot.

The author suggests splitting the work between multiple workers. Each worker has its own FIFO queue containing **URLs in prioritized order**. One might implement a separate set of queues to ensure politeness (ensuring that the web crawler is not DDOSing websites).

Another vital thing to keep in mind is the system's robustness.

- Use **checksums or hashes to find duplicate web pages**.

- Watch out for **spider traps** – websites that sole purpose is to trap your crawler in an infinite loop. One solution to this problem is to ignore URLs X characters long or more.

- The system has to be extendable. You might need to add support for different file extensions later on!
