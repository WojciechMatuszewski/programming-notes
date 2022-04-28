# System Design Interview - An Insider's Guide: Volume 1

My notes while reading the book. Random thoughts and such.
I decided to give the book a read since I like the "bigger picture" when it comes to architecting solutions. I often question how people on my team are aware of all the edge cases and such. By reading this book, I hope to gain insights into how "generic" systems are built.

## Chapter 1: Scale from zero to millions of users

Great introductory chapter. It touched on many "core" mechanics that govern most (I hope) applications that we use daily.

The author talked about:

- Scaling methods (vertical / horizontal). Both for the web tier and the data tier.

- Using load balancers for scaling and disaster-recovery purposes.

- Caching, be it locally via some cache instance or CDN.

- Deploying to multi-region environments and all the challenges that come with that.

- Data sharding.

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

There are so many different ways to handle rate-limiting I was not aware (unsurprisingly) that there were so many different algorithms. I'm most familiar with the _token bucket_ algorithm that AWS uses for most, if not all, of its services.

Some notes from the architecture part itself:

- Rate limiting implementation as a middleware between the client and the server.

- Redis acts as a data layer for the rate-limiting rules and keeps the state of the logic. Redis has been a prevalent choice in this book so far.

- Rate limiting services might run in a distributed environment. If that is the case, the architecture must take **rate conditions** and **synchronization** issues into account.

- For resiliency, you might need to make your rate limiter _eventually consistent_ which feels a bit weird to me.

## Chapter 5: Design consistent hashing

This chapter is hefty on the theory side but important.

Consistent hashing algorithms power data/load distribution across many different servers — something we take for granted these days.

The chapter touches on the theory behind **_ring buffers_** and their role in load distribution. Interestingly, the first, quite sophisticated implementation was insufficient at providing good load distribution.

The key takeaway is this: consistent hashing allows you to overcome the "celebrity" load problem. The consistent hashing mechanisms are used in various systems, not only databases or load balancers (Discord uses consistent hashing too).
