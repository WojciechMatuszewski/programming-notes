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

Finished 1 17:53
