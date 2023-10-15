# Caching

Caching is quite hard. There is a tradeoff between complexity and reliability of the solution. One also has to remember about keeping the data relatively fresh (depends on the application). Otherwise the users could potentially see outdated data.

## CPU

- Your CPU has multiple layers of cache!

  - These are called L1, L2 and L3 caches.

  - The **higher the number, the bigger the latency accessing the cache, but the cache itself is bigger**.

## In-memory cache

- As long as the language you are using does not garbage collect piece of data, you can use it as a cache.

- This technique of _in-memory_ caching is quite popular with serverless functions (likes _AWS Lambda_) and frontend applications.

  - The frontend application will cache responses from the backend and re-use them whenever possible.

  - In the case of AWS Lambda, we initialize some variable _outside_ of the handler. This variable will live until the execution environment is torn down.

## Memcached

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
