# Random Things I've picked up while studying

## Rolling back vs. rolling forward

> Taken from [this AWS Builder's Library article](https://aws.amazon.com/builders-library/cicd-pipeline/).

You might have heard the term _roll back_ before in the context of CI/CD pipelines. It is when you want to bring the system to the old state before a change. **This is the most common way of restoring the system in the case of an emergency**, where a change introduces a bug.

But what happens if the **change latent change – a change introduced a while back that now sits sandwiched between other changes?**
In such cases, the engineer must **pragmatically decide whether it is worth rolling back all the changes, including the latent change, OR is it worth rolling forward**. _Rolling forward_ means creating a change to fix the underlying issue and letting it advance through the CI/CD pipeline. This ensures that we do not roll back valuable releases, like bug fixes.

## Different types of Deployments

### Canary deployment

The idea behind the _canary deployment_ is to **route some percentage of traffic to the new deployment and monitor how the system behaves**.
If everything is ok, you can route more and more traffic to that deployment, until all of the traffic is switched to it.

**The main drawback is the user experience IF things go wrong on the canary**. The users will feel it. If the system suddenly breaks, those users will not have a good time using your app until you route them back to the "old" deployment.

This type of deployment **is great at getting some "real-life" testing done to the new version of the app**.

## Dark Read deployment

Instead of routing the users to new deployment, **you "clone" the request to the new version of the system, but serve the old system to all users**.
This allows you to **monitor how the new deployment behaves without impacting users if things go wrong**.

This deployment type **quite complex and could be costly, but is definitely worth it for critical systems**. If your new deployment explodes, there is no need to rollback – no users are using it!

**Having said that, it is a great way to load test your new deployment**. You can't do that with canary deployments as it would require you to shift _all_ traffic to the new version, rendering the canary deployment an "all at once" deployment.

## TLS Termination

When making a request to a server, you most likely use HTTPS protocol. This ensures the data you sent is encrypted and secure.

To handle the HTTPS request, the server would have to decrypt the data, handle it, encrypt it back and send it to you – **this is a lot of work for a single server to do**.

The **term _TLS Termination_ refers to a point where the data is decrypted and forwarded to the destination**. **This is mostly implemented via _load balancers_, _api gateways_ or _CDNs_**.

When the destination responds, the "TLS Terminator" will re-encrypt the data.

## Different types of Proxies

Keep in mind that the definition change depending on from which _side_ of the request you are looking at.

From the origins perspective, the proxy would be a reverse proxy and so on.

### The _forward proxy_ (or _proxy_)

The **_forward proxy_ sits between you AND the destination**.

A good **metaphor is having an assistant call a restaurant and arrange diner for you**. You have "the assistant in front of you" and the restaurant is the "destination".

From technical perspective, **a good example of a proxy would be any CDN or some tool that filters network traffic**.

### The _reverse proxy_

The **_reverse proxy_ sits between the incoming traffic and the server**.

A good **metaphor is visiting a restaurant and asking for table. The waiter will assign you a seat. From the restaurants perspective, the waiter is a _reverse proxy_**.

From technical perspective, **a good example of a _reverse proxy_ would be a load balancer**.

## The `.env` files

I was watching [this video](https://www.youtube.com/watch?v=j2JRBZaMDSg) where the author makes a compelling case _against_ `.env` files.

- If your `.env` files are not checked into the repository, you might have lots of trouble running the app as it existed at some point in the past.

  - Why would that be useful? Well, if you want to fix an issue, you most likely want to go back to how the app was functioning previously and see what the difference is between the working version and the current one.

- [There is also this follow-up video](https://www.youtube.com/watch?v=5lb3T3R_z2k) where the author talks about the security implications of `.env` files on a server.

  - There have been numerous cases where companies _hosted_ their `.env` files as static files accessible by anyone. In fact, in my work, I see requests to various `.env` files happen all the time! People are actively trying to "probe" for those files in hopes of stealing application secrets.

The author makes a case for using vendor-specific services to host secrets, like _AWS Secrets Manager_. I agree with this approach, but I also think that this solution is not for everyone.

Keep in mind that, depending on the infrastructure, reading secrets might add latency to your application, especially in serverless environments, even when caching those secrets in-memory. In addition, one has to be mindful of the complexity such a solution brings. You usually will need to add some kind of library, or write custom code, to fetch those secrets.

## Caching

Caching is quite hard. There is a tradeoff between complexity and reliability of the solution. One also has to remember about keeping the data relatively fresh (depends on the application). Otherwise the users could potentially see outdated data.

### CPU

- Your CPU has multiple layers of cache!

  - These are called L1, L2 and L3 caches.

  - The **higher the number, the bigger the latency accessing the cache, but the cache itself is bigger**.

### In-memory cache

- As long as the language you are using does not garbage collect piece of data, you can use it as a cache.

- This technique of _in-memory_ caching is quite popular with serverless functions (likes _AWS Lambda_) and frontend applications.

  - The frontend application will cache responses from the backend and re-use them whenever possible.

  - In the case of AWS Lambda, we initialize some variable _outside_ of the handler. This variable will live until the execution environment is torn down.

### Memcached

- Rather "bare bones". There is not much functionality.

- **Supports partitioning through the _content hash ring_**.

- Supports LRU eviction.

  - Implemented through **_doubly linked list with a hash map_**.

  - You evict from the tail.

  - When you read the node, the hashmap allows you to know where it is in the linked list. Then you move the head to be that node that you have just read.

- If you need something custom, you can always implement it on-top of this cache.

## Redis

- A much more sophisticated than Memcached. It has a lot of additional features.

- It is almost like a database. You have transactions, _write ahead logs_ and other similar features.

- Redis forces you to use single-leader replication. Memcached is easier to replicate.

## System resiliency

> Based on [this youtube video](https://www.youtube.com/watch?v=rvHd4Y76-fs)

### When adding retries can do more harm then good

Your client adheres to all the best practices you learned about when it comes to making an HTTP calls.

- You have retries in place.
- You have request timeouts in place.
- You add jitter in-between the retries.
- You use exponential backoff timeouts to wait before making another retry call.

**But what happens if the API you are calling is overloaded?**

In that situation, you are _adding even more work to the API_ by performing those retries!

Marc mentions it is like your boss telling you that you have X more amount of work to do, even though you are asking him for a vacation.

**If the system is overloaded, performing retries can make the outage even worse, especially if multiple clients are doing it**.

### Good usage of retires

There is no silver bullet in engineering, and the _way_ you handle errors depends on what is going on in the API you are calling.

- **Skipping retires** is a good strategy when the system is overloaded (systemic failures).

- **X retries** is a good strategy for transient errors. This **is the most common approach because transient errors are the most common errors**.

  - So, you are in a situation where you _think_ you are engineering for robustness, but you might be shooting yourself in the foot!

- **Adaptive amount of retries** is great for both systemic and transient failures.

  - You can read [about it here](https://docs.aws.amazon.com/sdkref/latest/guide/feature-retry-behavior.html#standardvsadaptiveimplementation).

    - In essence, you use _token buckets_ to deduce weather the request ought to be retried or not.

### What about backoff and jitter?

To understand when to use which one, we have to understand what does a _closed_ and _open_ system is.

- **A open system** is a system where requests come at random, and you have to way of controlling the amount of requests. This a web server powering a website.

- **A closed system** is a system where you have a fixed number of clients you are aware of. Think five (exactly five) machines polling some API for results.

So, when to use _backoff_ and when to use _jitter_?

- **Backoff is _very effective_ for closed systems and _mostly ineffective_ in an open system**.

  - For _closed systems_, we _know_ how many requests to expect, so telling the clients to slow down is quite effective.

  - For _open systems_, we _do not know_ when someone will attempt to make a request (a new client might attempt to connect), so telling everyone to "backoff" is not that effective.

- **According to Marc, using jitter is _always_ a good idea**.

  - By _jitter_ we mean not retrying every X on a fixed interval, but adding some "randomness factor" to X.

### Circuit breakers

The _circuit breaker_ pattern is where you have some logic deciding, based on the responses from the service, weather to send a request to it, or reject the attempt. **While a good solution, if implemented "naively" could lead to unnecessary downtimes, especially in sharded systems**.

Consider a sharded system with four shards A,B,C,D where shard D is overloaded and rejects requests.

If your circuit breaker implementation is not robust enough, **you might reject ALL requests because you see errors on shard `D`**.
So, all clients are blocked from making the requests to other shards until shard `D` becomes stable again. This is not a good situation to be in.

[Marc mentions](https://youtu.be/rvHd4Y76-fs?t=1572) some solutions to this problem, like implementing more sophisticated algorithms to decide if the circuit breaker should be in "on" or "off" state.

### Tail latency and techniques to improve it

First, let us define a the _tail latency_.

From what I understand, the **_tail latency_ represent the system latency at the very end of the P99.xxx spectrum**. For example, you might be looking at `p99.99` or `p99.999` where the response times are quite long.

I would say that, in most systems, you do not have to worry about tail latency. Making sure `p95` or `p99` is good enough should do it. But some systems, for example those AWS is building, need to optimize further, thus they focus on the _tail latency_.

So, how can we optimize the _tail latency_?

- **The "first response of two** technique where you send two requests at the same time, and use the first resolved one as the result.

  - This **increases the amount of work the server has to do**, but it makes the `p99.xxx` scenarios much less frequent since it is very unlikely that two requests will suffer form the "tail latency problem".

- **The hedging** technique where you send the first request, and if it does not come back after certain amount of time, you send another.

  - This **increases the work and might lead to issues we talked about related to retries and system overload**. If the system is overloaded, and takes a while to respond, you will be adding _even more_ work to the system!

- **The erasure coding** technique where you send X amount of requests, and assemble response from Y (X > Y).

  - I'm unsure if one can use this calling a single endpoint, but it is great technique when API supports partial responses.
