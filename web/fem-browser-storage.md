# Browser Data Storage FEM course

## Introduction

- The concept of an **_origin_** is important when talking about browser storage.

  - The _origin_ is the _protocol_ + _host_ + _port_. So you can have different origins for the same domain, for example an http and https origins for a domain example.com.

  - The **storage access is _origin_ based**. Different origins cannot access each other locally stored data.

- The **storage quota is _origin_ based in MOST browsers**.

  - There is a difference in Firefox. Firefox applies the quotas for a given domain plus all the subdomains for that domain.

  - Safari uses so-called partitions to help with privacy. It boils down to separating the storage access for websites embedding your website as an `iframe` and your website as a standalone thing.

- There are a lot of **different _web clients_ that can access web storage**.

  - There are PWAs, Web Views, the browser and others

- Overall, the landscape is rather complex. There are a lot of exceptions related to different browsers, web clients and origins. A lot of combinations.

- You **might want to prefix storage names per some kind of heuristics** to avoid collisions.

## State of APIs

- There are several choices.

  - _Cookies_ are **NOT a good way to store data** since they are sent to the server. The more data you save in cookies, the bigger the request payload is. This is suboptimal.
s
  - _Local storage_ is most widely used, but **it should not be**. There are better ways

    - The main reason for avoiding it is the performance. The API is synchronous. **For small amounts of data it is okay, but for larger payloads consider other solutions**.

    - The limits for the data size are quite low. **12MB for _session storage_ and 5MB for _local storage_**.

    - It is **not available in _Workers_ and _Service Workers_**.

  - _WebSQL_ is deprecated.

  - _Application Cache_ is deprecated.

  - _IndexedDB_ is **pretty good solution for data storage**.

    - The API is not the best, but we can create an abstraction layer over it.

    - You would **store objects here**.

    - **Available in _Workers_ and _Service Workers_**.

  - _File and Directories_ is deprecated.

  - _Cache Storage_ is **pretty good solution for data storage**.

    - You would **store HTTP responses here**

  - _FileSystem Access_ might be pretty good in the future. Lacks browser support.

  - That is the only storage technology to require users permissions.

- The **bottom line** is to use either **IndexedDB, CacheStorage and Web Storage**.

## Debugging tools

- Most, if not all, browsers have a panel in developer tools dedicated to storage.

- Keep in mind that the data you store is public. The user might "break" the data you stored.

## Quotas and persistence

- The **storage quota is shared amongst all the storage APIs**.

  - Of course, some of them have an upfront limit, the like _browser storage_ (aka _local storage_), but despite this, the size of the browser storage still contributes to the overall quota.

- There is more than enough space available to you.

- You can **request for the data to be persisted through different sessions via the `navigator.storage.persist()`**.

  - Take in mind that, some browsers will ask the user for confirmation (Firefox), but some will not (Chrome).

## Coding

- Use the `navigator` object.

  - To check the quota, use the `navigator.storage.estimate`

  - To check / ask for persistent storage, use the `navigator.storage.persist`

- The **_local storage_ requires a lot of error handling and sanity checks**.

  - Ensure that the data is JSON – you have to wrap the `parse` with a `try/catch`.

  - Ensure that the parsed data is the data you expect – maybe employ some kind of schema validation?

    - This **problem of validation is not only related to _local storage_. User can fiddle with IndexedDB and other storage APIs as well**!

- **Keep in mind that, the _local storage_ does not count towards your quota** so it is hard to know how much more data you can store.

- There is a **_IndexedDB v2.0_ which is an improvement in API**.

  - The API is event based, but there is a thin wrapper to convert it to promises. [The _idb_ package](https://github.com/jakearchibald/idb).

    - There are a lot of libraries that sit on top of IndexedDB.

- _IndexedDB_ supports transactions. This is quite neat!

  - _IndexedDB_ **supports versioning**. Each "database" is versioned.

    - **We do not have to prefix the key in IndexedDB because the separation happens at the database level**.

- With the IndexedDB or Local Storage one could implement the **storage-first pattern** (not the serverless one :p) where you fetch the data from the storage layer and not the API.

  - If you do this, you need to invest in a mechanism to invalidate the cache. Since we do not operate on HTTP layer here, adding an HTTP response will not do.

  - In most cases, using the application manifest (some kind of file denoting application version or it's options) will do. If the timestamp of the version of the DB is different than the manifest, the application knows that it should revalidate the local data.

- Another technique is the **server-first pattern**. The local data is a fallback for a failed network request.

Finished Part 4 50:22
