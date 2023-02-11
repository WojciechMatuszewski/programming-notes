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
