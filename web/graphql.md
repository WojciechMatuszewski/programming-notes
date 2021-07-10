# GraphqlStuff

## Enums

Just like in typescript you can define enums (you probably only need constant enums in TypeScript).

Here is how we would define custom enum in GraphlQL:

```graphql
enum PetCategory {
  CAT
  DOG
}
```

## Aliases

This one is quite useful, especially when doing pagination (I'm looking at you Prisma). Let's imagine you are fetching paginated data and want to know how many results (in total) there are. Usually you have to do separate query for that. Aliases makes it soo easy to do in one query string.

This is especially handy when `aggregate` prop lives on the `query` itself.

**Aliases can be added to any field**.

```graphql
query {
    resultsCount: {
        getItems {
            aggregate {

            }
        }
    }
    getItems(/* variables */) {

    }
}
```

## Interfaces

> Interface is an abstract type that includes set of fields. These fields must be used when creating new instances of that interface

There is high resemblance between GraphQL interfaces and typescripts.

Interfaces have different implementations of them.

```graphql
interface Human {
  age: Int!
  name: String!
}
# It's a pity that you have to re-include the types
type Myself implements Human {
  age: Int!
  name: String!
  workPlace: String!
}
```

## Unions

Unions allow you to return lists of multiple types. Nothing more special about it, its pretty useful though.

```graphql
union Pet = Cat | Dog
```

While fetching data you can specify fragments on each

```graphql
query {
        getAllPets {
            ... on Cat {

            }
            ... on Dog {

            }
        }
    }
```

## Persisted Queries

Query strings can get big. You might never though about this before but sometimes your queries can weight a lot, and I mean a lot.

With persisted Queries, your server will have a hash for a given query. You only send that hash and variables, no Query string needed.

This approach allows for more security and limits potential malicious attacks. Sending very large Query in hops of overloading the server wont work since it does not know anything about that query.

This approach works for **static Queries**.

## Static Queries / Mutations

Static Queries / Mutation are queries / mutations that has the same string, no matter the variables you pass. This means that they are created using variables, not computed at run-time using interpolation.

## Anemic Types

These are types that you should avoid. The term `anemic` means something without strong reason for being defined. A good example would be an class where all (or most) of it's methods are simple `getters` and `setters`. There is no value in that.

These types can be used within `mutations` and `queries`. An example:

```graphql
updateCheckout(input: UpdateCheckoutInput!): UpdateCheckoutPayload

  input UpdateCheckoutInput {
    email: Email
    address: Address
    creditCard: CreditCard
  }
```

Notice that all of the fields are `optional`. This design goes against what true strength of `GraphQL` really is - **it's type system and enforcing required fields**.

The `UpdateCheckoutInput` is very `data-driven` instead of focusing on **behaviors**. Prefer types that have less nullable fields and do one thing and one thing well.

## Fragments and data-driven components

You probably heard of [Relay](https://relay.dev/docs/en/introduction-to-relay). It's a well known `GraphQL` client framework. While working with `Relay` you will be treating your component and the data it needs as one.

While you cannot practice this methodology using `apollo-client` that much, you still can reap the benefits of at least trying to imitate it.

An example:

```js
const PARENT_QUERY = gql`
  ${CHIlD_FRAGMENT}
  query ParentQuery {
    # stuff
  }
`;
function Parent() {}

const CHILD_FRAGMENT = gql`
  fragment SomeFragment on SomeType {
    # stuff
  }
`;
function Child() {}
```

See how I'm collocating the data-needs within the vicinity of the `Child` using `Fragment` ?. This allows me to easily tweak it and make changes to it without touching the `Parent` query. **This is not an ideal solution** but I think it's worth considering.

### Fragment matchers

When your schema is using _GraphQL interfaces_ or _unions_ you might have encountered a warning about _Heuristic fragment matcher_.

The warning itself is not critical but should not be taken lightly. Since your _GraphQL client_ usually (unless you fixed the warning :) ) does not know nothing about your schema, it has no way of validating if the fragments you are using on a type that implements an _interface_ are indeed correct.

The default behavior is to use something called _heuristic fragment matcher_. The _heuristic_ means _good but not ideal_. Since this fragment matcher cannot handle fragments on unions or interfaces, you will need to implement your own.

You can use `graphql-codegen` to do that.

## Approaches for creating a schema

As with almost everything within JS ecosystem, there are multiple ways you can go about when it comes to creating a schema.

### Schema first approach (SDL)

This is where you create schema using the SDL (usually the `gql` tag).

I've been using SDL approach through this document.

### Code first approach

I have not used this personally, but there are a lot of benefits when it comes to creating schema this way.

- you can easily create custom types, SDL is dumb and you cannot do that.

- code-sharing should be easier.

The **ultimate way** of creating a schema would probably be a **mix of SDL and code first approach**. This is where you design your schema using SDL, but you implement it using code-first approach.

### Using Annotations

Mainly used with languages such as `Java`. There is a danger to couple `GraphQL` definitions with your implementation details since it's so close to your domain code.

## Egghead Nik Graf

- start with writing the `queries` not the types. It's a good way to start

- naming fields is very important, it will have an impact later on. **Always be as explicit as you can whenever naming fields**.
  Remember not to expose implementation details here. As per golang idioms, you **should not repeat the underling type within the field name**.

- there should be a balance between `nullable` vs `non-nullable` fields. Remember that `nullable` fields give you the opportunity to return partial results if something goes wrong.

- you **cannot create union of `Scalar` types**.

- prefer the `Relay` pagination spec. It's considered based practice, it will save you a lot of headaches. Remember that you can _extend_ the spec, it's not rigid.

- what out for impossible states when it comes to query parameters. When you have multiple ways of quering the same type, use `typeBySomething` convention.

  ```graphql
  productBySlug(slug: String!): Product!
  # optional just `product`. It's a team decision
  productById(id: ID!): Product!
  ```

- the `input` notation actually is a _Relay_ spec. Pretty sweet!

- try to **return entities which were affected by the mutation** in the **result of the mutation**. This will probably be a type which contains `payload` within it's name.

- always design with domain in mind. Again, be as verbose as possible.

- _update_ type mutations require tradeoffs. This is where you create an input with all fields optional, but some of them might be actually required on the entity are trying to update.
  This is where you might need to implement some validation on the backend to make sure user does not set required values as `null`s.

## GraphQL query complexity

Usually, most of the APIs you will deploy as part of your work (or work with) will have some rate-limiting applied to them.
The reason is understandable - we would not want someone to DDOS our servers (there are servers in serverless).

While this approach works quite well for REST APIs, it lacks flexibility when it comes to the GraphQL API.
In GraphQL world, the query can be simple and very complex (even recursive). So if we were to apply the REST-style rate-limiting to those APIs, we could be throttling our users on operations that do not take that much time to compute on our time while also allowing them to DDOS us with very complex queries.

All these problems could be solved by using a GraphQL query complexity analyzer. Each user is given a finite amount of "points". These points are spent doing operations and refilled on some basis. While this approach might sound like we are implementing a bucketing strategy to rate limiting, it's much more than that.

The queries are evaluated based on their complexity, not how often they are made. Thus, simple queries cost less than mutations which cost less than very complex queries that involve pagination.

Few open-source libraries allow you to get the complexity of a given query before it's executed and go from there. But they are not perfect.
[IBM researched the topic](https://www.youtube.com/watch?v=5Xtw5XDIyFw) where they found that the open-source solutions do not work that well. Sadly, AFAIK, they did not open-source their solution.
