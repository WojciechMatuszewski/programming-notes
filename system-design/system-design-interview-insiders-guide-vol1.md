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

- My first thought was about rate limiting on the backend, but in fact we can also do it on the front end.
  A good example of ambiguity that I should ask a question about.

- There are different rate algorithms used for rate limiting. The one I'm the most familiar with is the _token bucket_ one. Not from the implementation perspective, but from usage. This rate limiting algorithm is the most common in AWS (I think).

  - One interesting thing here is that one might need to have multiple buckets for multiple API endpoints.
    This one is also implemented within Amazon API Gateway!

- Apart from the _token bucket_ algorithm there is also _leaking bucket_, _fixed window counter_, _sliding window log_, _sliding window counter_ and others.

- For the implementation, the author points to Redis and a request middleware. These are good choices, but I would have chosen DynamoDB as the data-layer solution.

- To communicate back to the client about the rate limit rules (the amount, wait time and such) the author chose to use `X-Ratelimit` prefixed headers. Interesting.

- Author points to two pitfalls that are really dangerous – the race condition and synchronization issues.

  - My mind was blown when the book introduced running Lua script inside Redis as a potential solution. Sorted sets in Redis seem more maintainable to me.
  - The synchronization issue is wild. Especially with distributed rate limiters architecture. _Eventual consistency_ seems like a solution but it feels "off" to me to make rate limiting logic _eventually consistent_. I guess it is a matter of tradeoffs if you have such problems.

- What we have considered so far is rate limiting on the 7th layer: the application one. One might implement rate limiting on different layers.
