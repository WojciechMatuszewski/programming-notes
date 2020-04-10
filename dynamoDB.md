# DynamoDB

## Basics

- table must have _partition key_ and **optional** _sort key_ (or range key)

### Data types

#### Scalar Types

- exactly one value (string, number, binary, boolean and null)

- **keys can only be string or number**

#### Set Types

- multiple scalar values (string set, number set, binary set)

For example(**Each element of a given set must be the same type**):

```js
["Apples", "Oranges", "Grapes"], [1, 2, 3, 4, 5, 6];
```

#### Document Types

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

### Capacity Units

> This allows us to control performance at the table level.

#### Throughput Capacity

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

They are called local because they are tied with partition key (hash key). Partition key is responsible for putting things in the same buckets and secondary indexes allow to do querying operations inside those buckets.

Since we are only doing operations inside _buckets_ it's pretty fast.

#### Global Secondary Indexes

These do not have to be tied with partition key.
They work _outside the buckets_. Global secondary indexes are **stored on their own partitions** (separate from the table).

#### Sort Key

Sort key enables _rich query capabilities_. **If you provided sort key (also called range key) your partition key (hash key) does not have to be unique**.

You can think about it like putting things that have the same partition key in the bucket and sorting (_quering_) them by sort key.

### Spare Indexes

- this an **alternative for filtering**. It works on an idea that when you scan / query you only pull the data from indexes

- this is where you **create an LSI** on **attribute that sometimes is not present on some records**.

Picture orders within a restaurant. Some of the orders might be _open_ - indicating that the order is not yet fulfilled. Instead of using _Scan_ or creating a _Filter_ expression you could create a _sparse index_ on the _open_ attribute.

Some of the orders were already fulfilled so they do not have that attribute, but some of them have. With this setup you could literally _Scan_ or use _Query_ to get all orders that are open.

Carefully picking HASH key is very important with this approach.

### GSI Overloading

- you should minimize the amount of GSIs you have.

- sometimes called **partition overloading**

- whats more important is that **attributes can be sort keys for GSI**

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

### Aggregation

In this example, what about aggregation? Because data has to be aggregated somehow.

It is a best practice to run aggregation outside of the DynamoDB computation. In our example this means that we schedule a lambda function triggered by CloudWatch Events. We scan the table sharded table and sum up every minute or so.

As an **alternative** you **could use DynamoDB streams**. Remember that _Dynamodb streams_ are timely ordered. To avoid throttling you can setup _batch window_ to have your function wait X seconds before being invoked. This should invoke your function whenever _batch window_ expires or the batch is _6 MB_ in size.

## API

- use `ConditionExpression` to fail specific operations. This can be used together with `Transactions (read / write)` to create powerful operations.
