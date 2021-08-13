# DynamoDB

## Basics

- table must have _partition key_ and **optional** _sort key_ (or range key)

## Contribution Insights

- you can view graphs on access patterns for your database

* there is an underlying cost for activating it.

## PartiQL support

- you can use the _PartiSQL_ to read from your table

* **can result in full table scans if you are not careful**. While you can do the famous `select * from WHERE = ...` expression, it will be costly

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

You can use either _provisioned capacity_ or _on demand_ mode to control
how many operations your table can handle.

### On demand

This mode might sound like a serverless dream come true. You pay only for what
you use and you do not have to worry about scaling your throughput up or down.

And in 99% of cases, this is exactly what is happening. But in that 1% of cases, you might want to "warm" or "pre-provision" internal DDB resources to handle given load.

### Provisioned capacity

Can you forecast the amount of read and write operations your application makes?
If so, you might want to look into _provisioned capacity_ mode for cost optimization reasons. Before you do so, is the engineering time to make those calculations (and make sure that they are up to date) worth the effort?

### Throughput Capacity

- **1 capacity unit = 1 request/sec**
- used to control read/write throughput

## Table Design

- while building your data model rely on user stories
- try to use single table design. This will allow you to avoid N+1 problem. The N+1 problem is where you get some data and loops through the results of that data (reaching to the database again).

## Indexes

- Mandatory Primary Key - Either simple or composite
- Simple Primary Key - Only Partition or Hash key
- Composite Primary Key - Partition Key + Sort or Range Key
- Partition or Hash Key decides the target partition

### Indexes basics

1. If the table has only partition key (**also called hash key**) then that key has to be unique.
2. If the table has partition key and sort key (**also called range key**) **their combination must be unique**

### Secondary Indexes

#### Local Secondary Indexes

They are called local because they are tied with partition key (hash key).
Partition key is responsible for putting things in the same buckets and secondary indexes allow to do querying operations inside those buckets.
Since we are only doing operations inside _buckets_ it's pretty fast.

#### Global Secondary Indexes

These do not have to be tied with partition key, but can, you can have GSI HASH and Partition key.
They work _outside the buckets_. Global secondary indexes are **stored on their own partitions** (separate from the table).

The **GSI entry (pk + sk, or pk) does not have to be unique**. This is different than the LSI / primary indexes.
You can also change the GSI value of the GSI pk and / or GSI sk without any restrictions. This is quite logical since otherwise it would not be possible to create _sparse indexes_.

> In a DynamoDB table, each key value must be unique. However, the key values in a global secondary index do not need to be unique.

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html

##### WCU and GSIs

While the GSI enable you to create elaborate and powerful querying patterns, they also can cause issues in regarding to throttling and WCU / RCU consumption.

With the RCU case, please remember that querying given GSI consumes a sum RCU in terms of all projected attribute sizes across all of the items returned. This usually is not a problem, but I've noticed that engineers often overlook this property of the GSI.

The RCU case is similar to the RCU one. The total provisioned throughput cost of a write consists of the sum of write capacity units consumed by writing to the base table nad those consumed by updating the global secondary index.
Imagine having 10 GSIs and writing an item that touches only half of them. You will be paying a lot more for a single write that you would have if the GSI were not there. This argument alone should makes us question each and every new GSI we plan to add, especially in the single table environment.

##### Considerations for not creating GSIs

You might decide to skip on creating a GSI and instead choose to use `Scan` API to perform ad-hoc querying.
I would say this is a good pattern if you are certain that the cost of having those GSIs, and their influence on RCU / WCU consumption, would
incur a non-trivial increase to your overall cost.

A blog post on this topic https://roger20federer.medium.com/dynamodb-when-to-not-use-query-and-use-scan-61e4ab90c1df

#### Sort Key

Sort key enables _rich query capabilities_. **If you provided sort key (also called range key) your partition key (hash key) does not have to be unique**.
You can think about it like putting things that have the same partition key in the bucket and sorting (_quering_) them by sort key.

### Spare Indexes

- this an **alternative for filtering**. It works on an idea that when you scan / query you only pull the data from indexes

- this is where you **create an LSI or a GSI** (most likely GSI) on **attribute that is not present on some items**.

Picture orders within a restaurant. Some of the orders might be _open_ - indicating that the order is not yet fulfilled. Instead of using _Scan_ or creating a _Filter_ expression you could create a _sparse index_ on the _open_ attribute.

Some of the orders were already fulfilled so they do not have that attribute, but some of them have. With this setup you could literally _Scan_ or use _Query_ to get all orders that are open.

Carefully picking HASH key is very important with this approach.

### GSI Overloading

- you should minimize the amount of GSIs you have.

- sometimes called **partition overloading**

- whats more important is that **attributes can be sort keys for GSI**

## Projections

While creating **GSI (HASH / HASH + RANGE)** you can **project other attributes on those keys**. This is an important concept because **keys store data, they have some `weight`**.
By default, when you have GSI, you only have access to attributes that are your keys. To use other attributes, you should use projections.

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

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-time-series.html

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

* you can use logical operators like `AND` with `ConditionExpression` to create powerful conditions.

### UpdateExpression

- you **can use SET, DELETE, REMOVE, ADD in one _UpdateExpression_**.

  - if you want to **perform multiple operations of the same _kind_** (multiple ADDs, multiple DELETEs) just separate those using `,`.

    ```
    ADD #count :count, ADD #somethingElse :value
    ```

  - if you want to **perform multiple operations of different _kind_** (multiple ADDs with multiple DELETEs) you only need commas separating operations of the same _kind_

    ```
    ADD #count :count, ADD #somethingElse :value DELETE #ids :ids
    ```

- there is **no `AND` keyword**. This keyword is present in `KeyConditionExpression`

### Transactions

The ability to perform transactional operations makes DDB really powerful.
There is one caveat you might not be aware of first, that will definitely come into play if you heavily really on transactions.

### Optimistic Concurrency Control (OCC)

The DDB transactions works on the premise that multiple transactions can be performed without interfering each other.
Whenever you do a transaction, a check is performed if another transaction is already "working" on a given entity. If so, an error will be thrown.

You can retry the transaction, to be super safe you could pass the `ClientRequestToken` to ensure idempotency.

If you design your tables correctly, you should not be having much issues with the way DDB handles concurrent transactions.
Usually you can just retry the request, ensuring that you have valid _Condition Expressions_ in place.

## Consistency

Usually, most of the `read` operations will be `eventually consistent`. This is due to the fact how DynamoDB is built.

When you have the same `PK`, items with that `PK` will be on the same partition. Each **partition** has **leader and follower nodes**. Data is replicated from the leader to the follower nodes asynchronously. Usually DynamoDB reads from a random node - either leader or a follower node. When that happens there might be a point of time where that data is not yet there on given node.

To make sure your reads are always targeting the `leader` node (where data is placed first) - you have to manually tell dynamo, on operation basis, to perform a `consistent read / write`.

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

### Prepending arbitrary symbols

- sometimes you want to get the enmity and all the other entities that relates to the entity, think GitHub repo and issues for this repo

* to ensure that you get the repo first and then the issues within the same query, you might look into prepending the `issues` SK with some arbitrary character, this would _push_ the `repo` entity up top

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

The _with limitations_ part is very important. From my initial research, **the supported conditions have to contain an equality check on all key attributes**. Might be useful sometimes like removing / updating an item and making sure it exists in a process. Here is a very simple code in Go doing just that

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

## Cost considerations

### Avoid keeping big blobs of data along small, frequently accessed ones in the same item

Imagine a scenario where you have a table that keeps users profiles. The APP that your DDB is for allows the users to upload their photo.
For historical reasons, the photos were stored as base64 encoded strings within the DDB, on the user profile item. Since the base64 string can grow up to 400kb, we have a problem.

Apart from the obvious problem of having a 400kb limit on the item, **we are wasting money**.
See, every time the **user profile is updated, DDB has to read the whole item in memory and THEN perform the update - you pay from the read and write**.

> Even if you update just a subset of the item's attributes, UpdateItem will still consume the full amount of provisioned throughput (the larger of the "before" and "after" item sizes). https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html

With a big attribute of user photo present on the item, we will be paying way to much for a simple `UpdateItem` operation.
There are **two ways of dealing with such situations**:

1. Upload the user photo to s3 and keep the pointer to that object in DDB.
2. Split the user profile item into two items, one that contains **only the image** and one that contains **all other attributes**.

Both of them perform the same optimization - having the user profile item "weight" less, thus making each operation on that item cheaper.

### Use TTL feature to purge unused data

DDB exposes a feature where, given an attribute marked as "TTL" (the name of the attribute does not matter, you have to say which attribute is the "TTL attribute") items will be deleted when the TTL expires.

It's not instantaneous though. There might be up to 48 hours of delay between "TTL" expiring and the item being deleted. This is due to the fact
that the sweeper that runs the deletion is spun up on spare capacity of DDB (source: https://youtu.be/S02CRffcoX8?t=1368)

### Mind the GSIs

Depending on how your GSI is set up, you might be paying too much.

Imagine a scenario where the GSI is set up to replicate 100% of the item attributes. Every time a new attribute is added, you will be paying for that
GSI replication. There is a sweet spot where such replication makes sense, but only when you actually use that GSI frequently. **In other scenarios you might be better of doing scans!**.

Here a workshop you can use to learn more https://github.com/robm26/cost
