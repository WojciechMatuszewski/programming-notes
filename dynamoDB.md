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
['Apples', 'Oranges', 'Grapes'], [1, 2, 3, 4, 5, 6];
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
['John', 1234, 'Apples'];
```

### Capacity Units

> This allows us to control performance at the table level.

#### Throughput Capacity

- **1 capacity unit = 1 request/sec**
- used to control read/write throughput

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

They are called local because they are tied with partition key (hash key). Partition key is responsible for putting things in the same buckets and secondary indexes allow to do quering operations inside those buckets.

Since we are only doing operations inside _buckets_ it's pretty fast.

#### Global Secondary Indexes

These do not have to be tied with partition key.
They work _outside the buckets_. Global secondary indexes are **stored on their own partitions** (separate from the table).

#### Sort Key

Sort key enables _rich query capabilities_. **If you provided sort key (also called range key) your partition key (hash key) does not have to be unique**.

You can think about it like putting things that have the same partition key in the bucket and sorting (_quering_) them by sort key.
