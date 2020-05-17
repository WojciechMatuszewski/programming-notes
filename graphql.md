# GraphqlStuff

- https://egghead.io/courses/graphql-query-language

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

See how I'm colocating the data-needs within the vicinity of the `Child` using `Fragment` ?. This allows me to easily tweak it and make changes to it without touching the `Parent` query. **This is not an ideal solution** but I think it's worth considering.

## Approaches for creating a schema

As with almost everything within JS ecosystem, there are multiple ways you can go about when it comes to creating a schema.

// TODO:

### Schema first approach (SDL)

### Code first approach

### Using Annotations
