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
