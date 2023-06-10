# DynamoDB

Random stuff about DynamoDB

## Basics

- table must have _partition key_ and **optional** _sort key_ (or range key)

- The **WCU is measured in 1KB increments (rounded up) FOR THE WHOLE ITEM you are writing to**.

  - This means that, **if the item you are writing to is large, but you update only a small portion of it, you will pay a lot for the write**.

    - This is why **you should consider splitting the "large" attributes apart from the small ones that you frequently update**.

- The **RCU is measured in 4KB increments (rounded up)**.

## Contribution Insights

- you can view graphs on access patterns for your database

- there is an underlying cost for activating it.

## Storage layer

- there are two options when it comes to storage layer (where your data resides). The `DynamoDB Standard` and the `DynamoDB Standard-IA`

- you can **switch between them at will**

- the **`DynamoDB Standard` has lower throughput costs** than the **`DynamoDB Standard-IA`**

- the **`DynamoDB Standard-IA` has lower storage costs**

## Data types

### Scalar Types

- exactly one value (string, number, binary, boolean and null)

- **keys can only be string or number**

### Set Types

- multiple scalar values (string set, number set, binary set)

For example(**Each element of a given set must be the same type**):

```js
["Apples", "Oranges", "Grapes"], [1, 2, 3, 4, 5, 6];
```

### Document Types

- complex structures with nested attributes (list or map)
- no restriction on data types stored
- you can mix multiple types together

  For example:

```js
{
    name: "John",
    age:22,
    address: {
        city: "Stamford",
        state: "Alabama"
    }
}
```

Example with List:

```js
["John", 1234, "Apples"];
```

## Capacity and performance

You can use either _provisioned capacity_ or _on demand_ mode to control how many operations your table can handle.

### On-demand

This mode might sound like a serverless dream come true. You pay only for what you use and you do not have to worry about scaling your throughput up or down.

And in 99% of cases, this is exactly what is happening. But in that 1% of cases, you might want to "warm" or "pre-provision" internal DDB resources to handle given load.

#### Scaling

The **`on-demand` mode does not scale that well for small workloads**. The [aws documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.InitialThroughput) says that **the table will provision up to double the previous peak traffic on a table**. This means that if you have a huge spike, you will most likely experience throttling on the table.

> However, throttling can occur if you exceed double your previous peak within 30 minutes.

Like everything in software world, there are tradeoffs. Usually you should be good with on-demand, **but for super spiky workloads (for which I often see this mode advertised)**, you might be better off with **auto scaling with provisioned capacity**.

### Provisioned capacity

Can you forecast the amount of read and write operations your application makes?

If so, you might want to look into _provisioned capacity_ mode for cost optimization reasons. Before you do so, is the engineering time to make those calculations (and make sure that they are up to date) worth the effort?

### Adaptive capacity

You might have been thought to avoid hot partitions while designing your access patterns. This makes complete sense – the more throughput one partition gets, the less the other would get, right?

To some extent, this is true, and it would be a very much valid concern if it were not for **adaptive capacity**. So what is _adaptive capacity_?

The mode itself is very well described in [this AWS blog piece](https://aws.amazon.com/blogs/database/how-amazon-dynamodb-adaptive-capacity-accommodates-uneven-data-access-patterns-or-why-what-you-know-about-dynamodb-might-be-outdated/)

> In practice, it is difficult to achieve perfectly uniform access. To accommodate uneven data access patterns, DynamoDB adaptive capacity lets your application continue reading and writing to hot partitions without request failures (as long as you don’t exceed your overall table-level throughput, of course).

You might be wondering if there is any delay between the throttling happening and the _adaptive capacity_ kicking in. **It used to be the case that we would have to wait a bit for the _adaptive capacity_ to kick in, but it is no longer the case**. [Here is the announcement blog entry](https://aws.amazon.com/about-aws/whats-new/2019/05/amazon-dynamodb-adaptive-capacity-is-now-instant/) that talks about _adaptive capacity_ being instant.

### Auto scaling

While the _adaptive capacity_ concern was to ensure enough throughput is allocated for a given partition, the **role of the _auto scaling_ is to ensure your table has enough WCU/RCU to handle the load**.

The _auto scaling_ feature **uses _AWS CloudWatch Alarms_ under the hood** to scale the WCU/RCU. This means that **there will be a delay (minutes) between the _auto scaling_ kicking in and the traffic increase**. Ideally your traffic would rise gradually, but this might not always be the case.

### Throughput Capacity

- **1 capacity unit = 1 request/sec**
- used to control read/write throughput

## Table Design

- while building your data model rely on user stories
- try to use single table design. This will allow you to avoid N+1 problem. The N+1 problem is where you get some data and loops through the results of that data (reaching to the database again).

### Indexes

- Mandatory Primary Key - Either simple or composite
- Simple Primary Key - Only Partition or Hash key
- Composite Primary Key - Partition Key + Sort or Range Key
- Partition or Hash Key decides the target partition

#### Indexes basics

1. If the table has only partition key (**also called hash key**) then that key has to be unique.
2. If the table has partition key and sort key (**also called range key**) **their combination must be unique**

### Secondary Indexes

#### Local Secondary Indexes

They are called local because they are tied with partition key (hash key).

Partition key is responsible for putting things in the same buckets and secondary indexes allow to do querying operations inside those buckets. Since we are only doing operations inside _buckets_ it's pretty fast.

- They **support _strongly consistent_ reads** but they are **limited by the partition size (they cannot be split, like in the case of GSI)**.

- You **cannot create LSIs after the table is already created**.

- You should favour GSIs unless you need strong consistency on the index. Do you really?

- The **LSI is also the copy of the base table**. This behavior is the same as GSI.

  > A local secondary index also contains a copy of some or all of the attributes from its base table.

  - But the **LSIs are synchronous and SHARE throughput with the base table**. In the end **you will also pay for the writes to LSI as well as the writes to the base table**.

#### Global Secondary Indexes

These do not have to be tied with partition key, but can, you can have GSI HASH and Partition key. **GSI is like creating a separate table with items that only contain that GSI**. Global secondary indexes are **stored on their own partitions** (separate from the table).

The **GSI entry (pk + sk, or pk) does not have to be unique**. This is different than the LSI / primary indexes. You can also change the GSI value of the GSI pk and / or GSI sk without any restrictions. This is quite logical since otherwise it would not be possible to create _sparse indexes_.

> In a DynamoDB table, each key value must be unique. However, the key values in a global secondary index do not need to be unique. <https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html>

Another thing to note is that **the set of data-types you can create the GSI from is limited**. For example, you **cannot create a GSI on a boolean value**. Before you plan your data model, ensure that the data-type is compatible with the GSI/index scheme.

##### GSI asynchronicity

The GSI replication is asynchronous. The leader node acknowledges the write and then sends the write to a stream to be picked up by other nodes (that stream is most likely a kinesis data stream). **That is why you cannot use `consistentRead: true` while using GSI**.

The **GSIs have much lower throughput than the base table**. This is because the **GSIs have much lower cardinality, as there is only a single PK and SK for a given GSI table**. This means that the **replication can lag quite a bit behind the main table**. See [this tweet for more information](https://twitter.com/ksshams/status/1651196657831403520?s=20).

If you want to learn more about GSIs, [check out this video](https://youtu.be/ifSckJlatWE?t=2114).

##### WCU and GSIs

While the GSI enable you to create elaborate and powerful querying patterns, they also can cause issues in regarding to throttling and WCU / RCU consumption.

With the RCU case, **please remember that querying given GSI consumes a sum RCU in terms of all projected attribute sizes across all of the items returned**. This usually is not a problem, but I've noticed that engineers often overlook this property of the GSI.

The RCU case is similar to the RCU one. The total provisioned throughput cost of a write consists of the sum of write capacity units consumed by writing to the base table nad those consumed by updating the global secondary index.

Imagine having 10 GSIs and writing an item that touches only half of them. You will be paying a lot more for a single write that you would have if the GSI were not there. This argument alone should makes us question each and every new GSI we plan to add, especially in the single table environment.

**Note that GSIs have their own RCU/WCU limits that are separate from the base table**.

###### Implications of GSI throttling

[Here is a great article](https://medium.com/shelf-io-engineering/five-ways-to-deal-with-aws-dynamodb-gsi-throttling-1a489803a981) on this subject. There is also [this Twitter thread](https://twitter.com/ksshams/status/1651196703004053505?s=20).

The main takeaways I have after reading it are:

- The base table is **not directly** affected when **GSI RCU** goes over the limit.

- The base table **is affected** when **GSI WCU** goes over limit. By "affected" we mean to say that **write operations will NOT succeed on the base table as well as the GSI** – [source](https://repost.aws/knowledge-center/dynamodb-gsi-throttling-table).

- There are various techniques that could help here. The most notable for me are splitting the table or sharding the GSI.

##### Considerations for not creating GSIs

You might decide to skip on creating a GSI and instead choose to use `Scan` API to perform ad-hoc querying. I would say this is a good pattern if you are certain that the cost of having those GSIs, and their influence on RCU / WCU consumption, would incur a non-trivial increase to your overall cost.

A blog post on this topic <https://roger20federer.medium.com/dynamodb-when-to-not-use-query-and-use-scan-61e4ab90c1df>

##### Provisioning capacity for an index

If you have your table set to **provisioned capacity mode** you will be able to set capacity for a given GSI. This might come in handy in cases where you want to optimize costs or are very well informed about the access patterns.

#### Sort Key

Sort key enables _rich query capabilities_. **If you provided sort key (also called range key) your partition key (hash key) does not have to be unique**.
You can think about it like putting things that have the same partition key in the bucket and sorting (_querying_) them by sort key.

### Sparse Indexes

- this an **alternative for filtering**. It works on an idea that when you scan / query you only pull the data from indexes

- this is where you **create an LSI or a GSI** (most likely GSI) on **attribute that is not present on some items**.

Picture orders within a restaurant. Some of the orders might be _open_ - indicating that the order is not yet fulfilled. Instead of using _Scan_ or creating a _Filter_ expression you could create a _sparse index_ on the _open_ attribute.

Some of the orders were already fulfilled so they do not have that attribute, but some of them have. With this setup you could literally _Scan_ or use _Query_ to get all orders that are open.

Carefully picking HASH key is very important with this approach.

### GSI Overloading

- you should minimize the amount of GSIs you have.

- sometimes called **partition overloading**

- whats more important is that **attributes can be sort keys for GSI**

### Time-related attributes

- You **most likely want to add some kind of time-related metadata to each item**.

  - Think of `createdAt` or `updatedAt`.

- **Consider granular attributes in addition to more generic `createdAt` and `updatedAt` ones**.

  - The reason is that you might want to perform an operation that gets items for a certain day, year and so on.

    - As an alternative, one could use the `>` operator on the timestamp.

      ```text
      #created_at > TIMESTAMP
      ```

## Projections

While creating **GSI (HASH / HASH + RANGE)** you can **project other attributes on those keys**. This is an important concept because **keys store data, they have some `weight`**. By default, when you have GSI, you only have access to attributes that are your keys. To use other attributes, you should use projections.

You could also project everything, but that is kind of inefficient. **There is an underlying storage cost that comes with every index!**

## Modeling

### One-To-Many (simple HASH + SK)

- remember that hash + sk has to be unique. You can leverage the fact that sk can be different for the same hash, **this is called overloading**.

- remember that Query has a powerful _Beings with_ option.

- remember that when using _Query_ you do not have to provide SK.

### Adjacency List and Many-to-Many

- uses overloaded GSI partition key.

- this is where you have HASH and SK and the GSI is usually the SK. This allows you to make queries both ways.

- while filtering by GSI (basically an local SK) you might want to use filter expressions. This is due to the fact that _Query_ will also return the item with SK equal to given value.

### Hierarchical Data

- picture stores which can be queried by _postal code_, _region_ , _city_ , _street_.

- this is where the notion of **pre-computed sort key comes in**. It's usually a concatenation of data separated by some symbol (maybe _#_). Use it along with _Beings with_ query expression.

- this will require you to create **GSI** (quite important) on the table which houses the _pre-computed sort key_.

- do not be afraid to make _Queries_ like _Begins with: ATTR#ATTR2#ATTR3_. This is how you ensure the hierarchical nature of your data.

## Write Sharding

Imagine voting situation with 2 candidates. **Even with DynamoDB adaptive scaling (PAY PER REQUEST) you will face problems if a lot of votes come in**.
One solution to this problem would be to **expand the key space**. Because if you think about it, instead of having only 2 ids for 2 candidates, you can have multiple ids for the same candidate, thus the name sharding.

```
candidateA -> used for aggregation
candidateA_1
candidateA_2
...
candidateA_N (N is the shard number)
candidateB -> used for aggregation
candidateB_1
candidateB_2
...
candidateB_N (N is the shard number)
```

### Time Series Data

This is where you have a row describing some time-based entry, maybe a measurement of a thermometer or smth. You want to get all the measurements from a given day so, most likely, your PK will be the date. You are probably not going to be quering measurements from X days before so why have them in the same table with the same WCU / RCU as your _main_ data set.

In this pattern you have **separate tables** for **current, previous and much older sets**. This way you can allocate appropriate RCU / WCU values to given tables.

<https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-time-series.html>

### Aggregation

What about aggregation in regards to time series data? Because data has to be aggregated somehow.

It is a best practice to run aggregation outside of the DynamoDB computation. In our example this means that we schedule a lambda function triggered by CloudWatch Events. We scan the sharded table and sum up every minute or so.

As an **alternative** you **could use DynamoDB streams**. Remember that _Dynamodb streams_ are timely ordered. To avoid throttling you can setup _batch window_ to have your function wait X seconds before being invoked. This should invoke your function whenever _batch window_ expires or the batch is _6 MB_ in size.

### Selective write sharding

_Write sharding_ is great at _expanding the key space_, but it does come with a drawback. How in the world are you suppose to know in which sharded partition given item resides?
If you did not put any though into how the shard number is computed, this could be a problem. Imagine a table of customers and mobile phone providers.

```txt
PK                        SK     PHONE
P#${PROVIDER_ID}#1        C#1    5222222
P#${PROVIDER_ID}#2        C#2    6666666
P#${OTHER_PROVIDER_ID}#3  C#3    7777777
```

To get a given customer knowing the ${PROVIDER_ID} and the customer number, you would have to issue parallel queries. This might be OK depending on the situation, but
it would be neat to make sure this access is a simple `key:value` lookup.

This is where the idea of _selective write sharding_ comes in. Instead of mindlessly sharding given partition, we could ensure that the ID of the shard can be computed given some information. In our case, we should be able to compute the shard ID if we know the customer phone number and the provider ID.

```txt
PK                                          SK     PHONE
P#${PROVIDER_ID}#${HASH(PHONE) % 10}        C#1    5222222
P#${PROVIDER_ID}#${HASH(PHONE) % 10}        C#2    6666666
P#${OTHER_PROVIDER_ID}#${HASH(PHONE) % 10}  C#3    7777777
```

Now the shard ID is deterministic and derived from the customers phone number. Given the ${PROVIDER_ID} and the customers PHONE we are able to get that customer by performing a `key:value` lookup

## API

- use `ConditionExpression` to fail specific operations. This can be used together with `Transactions (read / write)` to create powerful operations.

- you can use logical operators like `AND` with `ConditionExpression` to create powerful conditions.

### UpdateExpression

- you **can use SET, DELETE, REMOVE, ADD in one _UpdateExpression_**.

  - if you want to **perform multiple operations of the same _kind_** (multiple ADDs, multiple DELETEs) just separate those using `,`.

    ```text
    ADD #count :count, ADD #somethingElse :value
    ```

  - if you want to **perform multiple operations of different _kind_** (multiple ADDs with multiple DELETEs) you only need commas separating operations of the same _kind_

    ```text
    ADD #count :count, ADD #somethingElse :value DELETE #ids :ids
    ```

- there is **no `AND` keyword**. This keyword is present in `KeyConditionExpression`

### Transactions

The ability to perform transactional operations makes DDB really powerful.
There is one caveat you might not be aware of first, that will definitely come into play if you heavily really on transactions.

#### Optimistic Concurrency Control (OCC)

The DDB transactions work on the premise that multiple transactions can be performed without interfering each other.
Whenever you do a transaction, a check is performed if another transaction is already "working" on a given entity. If so, an error will be thrown.

You can retry the transaction, to be super safe you could pass the `ClientRequestToken` to ensure idempotency.

If you design your tables correctly, you should not be having much issues with the way DDB handles concurrent transactions.
Usually you can just retry the request, ensuring that you have valid _Condition Expressions_ in place.

#### Transactions and other operations

Imagine yourself performing a transaction that involves changing items A and B. At the same time, you kick-off a `GetItem` request for the item A and B. It turns out that the read operations may return different results, all is based on timing.

- Both `GetItem` requests are run before the TransactWriteItems request.

- Both `GetItem` requests are run after the TransactWriteItems request.

- `GetItem` request for item A is run before the TransactWriteItems request. For item B the `GetItem` is run after TransactWriteItems.

- `GetItem` request for item B is run before the TransactWriteItems request. For item A the `GetItem` is run after TransactWriteItems.

Definitely interesting. [Rick Houlihan claims that he actually never used the transaction API](https://twitter.com/houlihan_rick/status/1430240266095669249)

#### Transactions and WCU/RCU

According to [this blog post](https://aws.amazon.com/blogs/database/optimize-amazon-dynamodb-transaction-resilience/) for any given operation within the transaction _DynamoDB_ will perform two read or write operations (depending on the underlying operation)

> As DynamoDB transactional APIs perform two underlying reads or writes for every item in the transaction

This makes sense since _DynamoDB_ uses _two-phase commit protocol_ where the data is first written/read to/from the storage nodes, then committed.
All of these operations are, of course, not free. **For every operation that is within your transaction you have to count double the WCU/RCU**.

I find this very fascinating as I never once though about WCU/RCU and the throttling with the _on-demand_ pricing mode. I guess I have not yet built an application that scaled to traffic that would be high enough to have to worry about it.

#### Transaction conflicts

The bigger your transactions are, the more likely you will run into transaction conflicts. This is where there are multiple operations competing for the same "resource" within the _DynamoDB_ table.

Imagine `PutItem` that operates on the item with `pk` of `1` and, at the same time, firing of a transaction that has the same operation.

Note that _DynamoDB_ uses the notion of _optimistic concurrency control_ (OCC). This means that there might be multiple in-flight transactions at any given time. There is no possibility of deadlocks as discussed earlier, before deadlock can occur, you will get the transaction conflict error.

#### Transactions on high contention items

The _DynamoDB_ transact API is great for items frequently read in parallel. If they are, you might face many "transaction conflict" errors.

One solution here would be to retry, but as we all know, this type of solution can only carry us so far. A much better solution would be to use _DynamoDB_ streams. You update one part of the application, then the stream reader will update the second.

What happens in the case of errors in the stream reader? Well, you would need to implement some kind of rollback mechanism. It might seem like it's a lot more work to do, but the chance of the stream reader failing is much lower than having your transaction fail due to high contention on the item itself.

This is related to the _DynamoDB saga pattern_ that I write about in the next section.

#### Transactions spanning multiple services

Imagine a scenario where the user checks out their cart. To carry out that operation successfully, we might need to perform multiple operations in a transaction-like manner:

1. Checking the inventory
2. Updating the global product count
3. Other operations ...

Like in the _Step Functions_ world, one might use the _Saga pattern_ with _DynamoDB Streams_ to achieve the desired result.

The other service subscribes to the stream updates, then writes to the origin. The front end would poll (or get notified via WebSocket) the transaction status. The transaction itself is _asynchronous_ because it spans multiple services.

[Check out this great video](https://youtu.be/IgFvWaSQaeg?t=1496) for a deep dive regarding this pattern.

## Locking

To prevent the situation where two requests change the same piece of data (the double-booking problem), it is vital to implement some kind of locking strategy for a given piece of data. There are two locking strategies I'm aware of.

### Optimistic locking

It's where you set a _version_ attribute on an item. A given operation can only update this item if the _version_ attribute matches the one specified in the request. You **would use the `ConditionExpression` here** and ensure that the _version_ attribute is the same as it was when you fetched the item.

The strategy is called _optimistic_ because it **allows for multiple non-clashing writes**. As long as the version is the same, the operation will go through.

### Pessimistic locking

Here, you are dealing with some sort of lock. If a given operation acquires a clock for an item, any other operation cannot change that item. **A good example would be the `TransactWrite` operation**. Here, if any other operation tries to write to that item, that operation will fail.

You can also implement physical _lock items_ in the Database.

## Consistency

Usually, most of the `read` operations will be `eventually consistent`. This is due to the fact how DynamoDB is built.

When you have the same `PK`, items with that `PK` will be on the same partition. Each **partition** has **leader and follower nodes**. Data is replicated from the leader to the follower nodes asynchronously. Usually DynamoDB reads from a random node - either leader or a follower node. When that happens there might be a point of time where that data is not yet there on given node.

To make sure your reads are always targeting the `leader` node (where data is placed first) - you have to manually tell dynamo, on operation basis, to perform a `consistent read / write`.

### Strongly consistent reads are NOT the silver bullet

After reading about the semantics of _strongly consistent reads_ you might be tempted to use them by default. I would argue that this is not the right way to go about the DynamoDB API usage.

Here is the kicker – the **strongly consistent reads are as fresh as the data you put into the database**. Keep in mind that **a lot can happen during the time it takes you to retrieve that data**. What if some other thread modified the value when you were fetching it? Then we are back to the semantics of eventually consistent read.

Apart from the possible in-transit mutations, there is one situation where **using the _strongly consistent reads is an overkill – when you use a condition expression**. What is the point of using SC if DynamoDB checks data integrity before it commits it to the storage for you? **Keep in mind that SC operations have an additional cost**.

So **when should you use SC reads?** In [this Twitter thread](https://twitter.com/ksshams/status/1620138322495680512), Khwaja talks about **monotonic reads**, which in this context mean **reads that never "travel back in time" – the data they read is as stale as the latest put in a given thread**. A good example would be a counter – you will never go from 33 to 32, you will always "read forward".

### Global Tables consistency

As good as the _Global Tables_ feature of _DynamoDB_ might seem, We should only use it for "append-only" use-cases. In the end, the tables are regional, and the "global" part is the **asynchronous** replication between those tables. The **collisions are resolved in _last-writer-wins_ manner**.

The biggest no-no when it comes to **atomic updates** is that the **transactions are not distributed**, meaning that you might have conflicts when someone performs operations on two tables in different regions at the same time.

> Take for instance a simple counter that you want to increment by 1 and need strictly linear counts. If you use global tables, you basically need to say "x-region" is now my primary. However, if there is a network partition, you can't really failover to a different region and guarantee you didn't already issue a number on a count. So you can't safely failover and must halt operations

For a in-depth explanation of how _DynamoDB Global Tables_ handle consistency, checkout this video: <https://www.youtube.com/watch?v=fqxL3WQ53GM&t=645s>

## Sorting

### Lexical sorting

There are some rules on how _string_ type is sorted, the order is basically a dictionary order with

- **uppercase letters come before lowercase**

- numbers and symbols are relevant

The most important part here is that **uppercase letters come before lowercase**. This can trip you out. This is why you see so many people use all uppercase or lowercase values in their tables.

### Timestamps

As long as the format is sortable, it's ok. You can use _epoch time_ or _ISO-8601_. Does not really matter, just make sure they are sortable.

### UUIDs

Regular _UUID_ will not do - it's not sortable. What you need is something that contains the timestamp and enough randomness to prevent collisions. You might look into

- KSUID

- ULID: this one even has a spec.

### ScanIndexForward

So you have all the necessary information about sorting related things to use this attribute. Remember, _DynamoDB_, by default, always scans forward, that means in ascending manner.

### Pre-pending arbitrary symbols

- sometimes you want to get the enmity and all the other entities that relates to the entity, think GitHub repo and issues for this repo

- to ensure that you get the repo first and then the issues within the same query, you might look into prepending the `issues` SK with some arbitrary character, this would _push_ the `repo` entity up top

- that character usually is `#` or `0`, depending on the use case and the structure of the data.

## PartiQL

When using DDB you have 2 choices when it comes to what kind of DSL you are going to use to perform your operations.

1. Use the native DSL.

2. Use the PartiQL SQL like syntax.

If I were you, I would always lean towards the native DSL. The native DSL guards you from doing silly things like scanning your whole table.
It's much easier to write `SELECT * FROM ...` than to write `db.Scan(...)`. The latter version forces you to consciously use the `scan` API.

### Conditions support and `batchWrite` API

Single DDB `UpdateItem` and `PutItem` and `DeleteItem` operations fully support the conditions.
The conditions are often used to, for example, create a new item if it does not yet exist.

The `batchWrite` API is different though. While the names of the operations are the same, **the native `batchWrite` item API does not support conditions on the operations**. Here is an excerpt from the [AWS documentation regarding the topic](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItem.html)

> For example, you cannot specify conditions on individual put and delete requests, and BatchWriteItem does not return deleted items in the response.

Here is the deal with `PartiQL` though, **the `batchExecuteStatement` API used within the context of `PartiQL` DOES support conditions on the operations - with limitations!**.

The _with limitations_ part is very important. According to my initial research, **the supported conditions have to contain an equality check on all key attributes**. Might be useful sometimes like removing / updating an item and making sure it exists in a process. Here is a very simple code in Go doing just that

```go
 out, err := client.BatchExecuteStatement(context.Background(), &dynamodb.BatchExecuteStatementInput{
  Statements: []types.BatchStatementRequest{
   {
    Statement: aws.String("UPDATE testTable SET foo=1 WHERE id='4'"),
   },
  },
 })
```

### The tweet

This section was inspired by [this tweet](https://twitter.com/jeremy_daly/status/1281628318895415299).
According to the Rick Houlihan answer the operation Jeremy tried to perform should be possible with `PartiQL`.

I believe though, that it impossible to perform that action by using the `batchExecuteStatement` API - the main reason being the restriction regarding the _WHERE_ clause.

As an alternative, the transaction API might be used, but it does not return the new data that you might have just updated.

### The Update / Insert dilemma

With the "regular" DynamoDB, one does not need to pick between an "update" and "insert" behavior. There is one operation that allows you to perform those operations - the `PutItem` API.

That is not the case with the `PartiQL`, though. Here, we either update an existing item or insert a new item that does not exist.
How would we know which one to choose? We most likely would need to perform a `SELECT` statement on a given item before performing the operation.

All of this could be done with one statement in DDB by using `SET attribute = if_not_exists(attribute, :value)` update expressions.

### WORM models with PartiQL

[Inspiration – this Twitter thread](https://twitter.com/NoSQLKnowHow/status/1458139547527647239).

It turns out that due to the distinct nature between the `Update` and `Insert` statements, one might use that to create a WORM data structure on top of _DynamoDB_.

Let us consider the "regular" _DynamoDB_ API for a moment. The `putItem` API call can create a new item or override an existing item.

This is not precisely the case with PartiQL. Here, the `Insert` statement will throw an error if the item we are trying to insert already exists.
Let us look into an example to understand how it works.

```go
db := dynamodb.NewFromConfig(cfg)

_, err = db.PutItem(
  ctx,
  &dynamodb.PutItemInput{
    Item: map[string]types.AttributeValue{
      "pk": &types.AttributeValueMemberS{
        Value: "pk",
      },
      "property": &types.AttributeValueMemberS{
        Value: "value",
      },
    },
    TableName: aws.String("test-table"),
  },
)
if err != nil {
  panic(err)
}

_, err = db.ExecuteStatement(
  ctx,
  &dynamodb.ExecuteStatementInput{
    Statement: aws.String(
      // Watch out for the dashes and single quotes!
      "INSERT INTO \"test-table\" value {'pk': 'pk', 'property': 'overwritten-value'}",
    ),
  },
)
// operation error DynamoDB: ExecuteStatement, https response error StatusCode: 400, RequestID: 2QNMCR7HRCVUVNK4VOJGGFQOIVVV4KQNSO5AEMVJF66Q9ASUAAJG, DuplicateItemException: Duplicate primary key exists in table
if err != nil {
  panic(err)
}
```

As we can see, the `Insert` PartiQL statement failed due to the "duplicate primary key" error. If we swap the `Insert` to the `insertItem` API call, this will not be the case.

```go
db := dynamodb.NewFromConfig(cfg)

_, err = db.PutItem(
  ctx,
  &dynamodb.PutItemInput{
    Item: map[string]types.AttributeValue{
      "pk": &types.AttributeValueMemberS{
        Value: "pk",
      },
      "property": &types.AttributeValueMemberS{
        Value: "value",
      },
    },
    TableName: aws.String("test-table"),
  },
)
if err != nil {
  panic(err)
}

_, err = db.PutItem(
  ctx,
  &dynamodb.PutItemInput{
    Item: map[string]types.AttributeValue{
      "pk": &types.AttributeValueMemberS{
        Value: "pk",
      },
      "property": &types.AttributeValueMemberS{
        Value: "overwritten-value",
      },
    },
    TableName: aws.String("test-table"),
  },
)
// No error occurs. Multiple `PutItem` API calls overwrite the same item.
if err != nil {
  panic(err)
}
```

All of this information is vital for building WORM - _write once read many_ models on top of _DynamoDB_.

How might one build such a model?

- Disallow all _DynamoDB_ API calls. [This IAM policy might come in handy](https://github.com/aws-samples/aws-dynamodb-examples/blob/master/DynamoDBIAMPolicies/AmazonDynamoDBAppendOnlyAccess.json).
- Disallow all PartiQL `Update` calls.
- Use the `Insert` statements for the first write.
- Use the `Select` statements for the writes.

While this setup might not be a "true" WORM datastore, it might be sufficient for your needs.
If you find this solution lacking, I would suggest looking into the [S3 Object Lock features](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock.html).

## Analytics

You might think that DynamoDB, a NoSQL database, might be unusable for any analytics use-cases. While there is SOME truth to that, it is not all doom and gloom.

### Utilizing periodic `Scan` operations

Depending on the size of your data, performing the `Scan` operation on your table might be a completely valid strategy. It all depends on the frequency and the size of your data.

If you need to perform the `Scan` operation infrequently and the size of your data is relatively small, I would say that you should go for it. In most cases, you pay more for GSI replication than performing the `Scan` operation once in a while.

#### `Scan` operation within an index (sparse indexes)

**Keep in mind that you can use the `Scan` operation on a GSI or LSI**. This means that that you will fetch only the items that have a given index.
This is **great use-case for migrations**, as you get a partial benefit of the `Query` API (you cannot use the `KeyConditionExpression`), but you do not scan the entire table.

### Export to S3

You can also export your data in DynamoDB to S3 and use BI-specific tools on the exported data. The DynamoDB does not consume any RCU/WCU during the export operation because **the export is based on the PITR backup instead of the live data**.

Remember that the export operation might take a while, even for tables with barely any data. One of my colleagues, Graham, raised [the question about the wait time on Twitter](https://twitter.com/Grundlefleck/status/1511359776478879752).

### DynamoDB streams to Firehose to S3

Ah, the "classic" data pipeline. Firehose is an excellent tool for putting stuff into S3. As long as your data is on S3, you can use Athena for BI things.

You can find a blogpost regarding this architecture [here](https://aws.amazon.com/blogs/database/how-to-perform-advanced-analytics-and-build-visualizations-of-your-amazon-dynamodb-data-by-using-amazon-athena/).

## Cost considerations

### Measure before you optimize

_DynamoDB_ is great. It allows you to pull the WCU/RCU information right from the operation you have just performed. There is an option within the SDK to get the `ReturnConsumedCapacity`. You can then send that information somewhere, maybe to your analytics pipeline,
where you would chart the cost of each operation.

### Use Reserved Capacity

AWS has a lot of services that enable you to pay an upfront fee in exchange for better prices on a given resource.
The _AWS DynamoDB_ is no different. **If you are running _provisioned capacity mode_ consider using _reserved capacity_ along with _auto scaling_**.
Such a combination is the most cost-effective way of using the service.

### Switch to Standard-IA

Usually, the data we store is rarely accessed after some period of time. In that case, one way to save on storage would be to use the `DynamoDB Standard-IA` storage class. The `DynamoDB Standard-IA` has higher throughput costs but lower storage costs than the `DynamoDB Standard` class.

Remember that you can switch between the two at will (at least that's what I've read in the docs. As of writing this, I have yet to test the functionality myself).

### Avoid keeping big blobs of data along small, frequently accessed ones in the same item

Imagine a scenario where you have a table that keeps users profiles. The APP that your DDB is for allows the users to upload their photo. For historical reasons, the photos were stored as base64 encoded strings within the DDB, on the user profile item. Since the base64 string can grow up to 400kb, we have a problem.

Apart from the obvious problem of having a 400kb limit on the item, **we are wasting money**.
See, every time the **user profile is updated, DDB has to read the whole item in memory and THEN perform the update - you pay from the read and write**.

> Even if you update just a subset of the item's attributes, UpdateItem will still consume the full amount of provisioned throughput (the larger of the "before" and "after" item sizes). <https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html>

With a big attribute of user photo present on the item, we will be paying way to much for a simple `UpdateItem` operation.
There are **two ways of dealing with such situations**:

1. Upload the user photo to s3 and keep the pointer to that object in DDB.
2. Split the user profile item into two items, one that contains **only the image** and one that contains **all other attributes**.

Both of them perform the same optimization - having the user profile item "weight" less, thus making each operation on that item cheaper.

#### Splitting items and global tables

Be aware of the eventual consistency when using the "split-item" pattern and global tables. Since you have to issue multiple writes to update an item associated with a single identity, there might be a point in time when some parts of the item are in inconsistent state (The `name` item is up to date, but the `surname` item write is still in-flight).

I've written about this topic [on the Stedi blog](https://www.stedi.com/blog/how-to-ensure-cross-region-data-integrity-with-amazon-dynamodb-global-tables).

### Use TTL feature to purge unused data

DDB exposes a feature where, given an attribute marked as "TTL" (the name of the attribute does not matter, you have to say which attribute is the "TTL attribute") items will be deleted when the TTL expires.

It's not instantaneous though. There might be up to 48 hours of delay between "TTL" expiring and the item being deleted. This is due to the fact
that the sweeper that runs the deletion is spun up on spare capacity of DDB (source: <https://youtu.be/S02CRffcoX8?t=1368>)

### Using filtering instead of a GSI

Depending on how your GSI is set up, you might be paying too much.

Imagine a scenario where the GSI is set up to replicate 100% of the item attributes. Every time a new attribute is added, you will be paying for that
GSI replication. There is a sweet spot where such replication makes sense, but only when you actually use that GSI frequently. **In other scenarios you might be better of doing scans!**.

Developers new to _DynamoDB_ often hear "avoid scans, use GSI instead". This advice, while most likely given in a good intentions in mind, might lead to actually paying more for _DynamoDB_. Of course, the caveat is WHEN that statement holds true.

I would firstly advise anyone to perform pricing calculations on their own and see the cost difference between performing a scan operation and the cost of "maintaining" a GSI (mind the GSI replication).

**As a rule of thumb, the less the given access pattern is used, the more likely scan operation will be cheaper for you, but only if the GSI replication % is high**. Otherwise, **if the GSI replication % is very low (sparse index or just replicating keys), the GSI will most likely be cheaper**.

### Consider _Provisioned Capacity_

I get it. I really do. The _On Demand_ model is really tempting. It's the serverless dream right?
Do not get me wrong, the _On Demand_ mode has it's place and time and you definitely **should use it**.

But maybe, during the lifecycle of your application, you started noticing patterns. And I'm not talking about having increased traffic during christmas or similar. I'm talking a day-to-day traffic patterns. If that's the case, you might want to look into _Provisioned Capacity_ most likely coupled with _Auto Scaling_.

#### Ramping up to avoid throttling

The _Auto Scaling_ takes time to catch up, it's not perfect. When you have some scripts that perform a lot of operations,
add logic so that the writes are spread and increase gradually.

#### Update the _Provisioned Capacity_ settings before a big spike

Usually done before big events that are very profitable for your business. In such critical moments, it's better to burn a bit more on a database, than to crash because of the load.

## Data integrity

As they often say, the data is the new gold. I support this statement as, in most cases, data drives any product forward (be it a customer, order, or any other data entity).

How can we ensure we do not violate data integrity if data is so important?

### Data types

When writing to DynamoDB, you have to explicitly tell the service what the data type of a given attribute is. This alone does not guarantee protection against writing the wrong type of data for a given attribute, but it forces developers to at least think about it.

I would say that the solution here would be to validate the data you will write against a JSONSchema. That should give you enough confidence.

### Required fields

Very similar to the previous point. Depending on the operation, you can use a JSONSchema (for create operations) or the DynamoDB conditions syntax (for update operations).

### Use data-access layer

You have to consider many things while working with data and databases. Would not it be nice to centralize all that logic into one module? Definitely.

So, as they preach in most programming books – use the data-access layer and work on domain objects within your application and not the "raw" database objects.

This approach will save you time and money in debugging time.

### Resources

- <https://serverlessfirst.com/data-integrity-considerations-writing-to-dynamodb>
