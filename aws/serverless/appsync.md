# AppSync

## DataSources and IAM

The AWS AppSync service has a notion of DataSources. These are services (like Lambda, DynamoDB, RDS, etc...) which will be used to provide data given a mapping template.

In the context of IAM, it's the DataSources that need permissions to perform operations on a given service. When you are using only a single DataSource within a context of a mapping template, this is usually taken care of by your framework (like aws-cdk or sam or serverless framework).

But there are situations when you need to do some work inside the mapping template that touches more than 1 table (think TransactWriteItem across two or more tables). This is where you have to customize the IAM permissions of a given DataSource.

The act of doing so is rather straightforward, the most important thing to remember here is to add permissions to the DataSource, not the resolver.

An example for aws-cdk

```ts
const likeATweetRequest = appsync.MappingTemplate.fromFile(
  getMappingTemplatePath("Mutation.like.request.vtl")
)
  .renderTemplate()
  .replace("LikesTable", props.likesTable.tableName)
  .replace("TweetsTable", props.tweetsTable.tableName)
  .replace("UsersTable", props.usersTable.tableName);

likesTableDataSource.createResolver({
  typeName: "Mutation",
  fieldName: "like",
  requestMappingTemplate: appsync.MappingTemplate.fromString(likeATweetRequest),
  responseMappingTemplate: appsync.MappingTemplate.fromFile(
    getMappingTemplatePath("Mutation.like.response.vtl")
  ),
});

// Add missing permissions
props.usersTable.grantWriteData(likesTableDataSource.grantPrincipal);
props.tweetsTable.grantWriteData(likesTableDataSource.grantPrincipal);
```

A quick refresher, the Principal in the context of IAM is a person or application that can make a request for an action or operation on an AWS resource. In our case, it's the DataSource making a request for DynamoDB related operation.

## Working with interfaces and unions

GraphQL unions and interfaces are powerful tools you have at your disposal when modeling a schema. The most important thing to remember is that **if you have extended the interface implementation, you have to have a way to distinguish between different interface or union implementations**. Please note that GraphQL unions do not have to share any fields.

The act of doing so is not that hard, all you have to do is to specify the `__typename` field on the data that you are storing. If you are using DynamoDB, do not mistake this with the type attribute that you often see defined on a given entity. The `__typename` is purely GraphQL related and will be for getting the data.

For the implementation itself, you can just add `__typename` as an attribute on a given DynamoDB entry

```ts
const newTweet = {
  __typename: TweetTypes.TWEET,
  id,
  text,
  creator: username,
  createdAt: timestamp,
  replies: 0,
  likes: 0,
  retweets: 0,
} as const;

await docClient.put({ TableName: "foo", Item: newTweet });
```

Or you can resolve the `__typename` dynamically within the .vtl template itself.

```vtl
#set($result = $ctx.result)
#set($id = $ctx.result.id)
#set($username = $ctx.identity.username)


#if($id == $username)
    $util.qr($result.put("__typename", "MyProfile"))
#else
    $util.qr($result.put("__typenam", "OtherProfile"))
#end

$util.toJson($context.result)
```

It really all depends on the use case here. Either way, when you are done with the implementation, you can now use the GraphQL fragments to get the data.

```graphql
query getMyTimeline($limit: Int!, $nextToken: String) {
  getMyTimeline(limit: $limit, nextToken: $nextToken) {
    nextToken
    tweets {
      id
      createdAt
      profile {
        id
        name
        screenName
      }
      ... on Tweet {
        id
        createdAt
        text
        likes
        retweets
        replies
      }
    }
  }
}
```

## Ways of hydrating results

This is a common problem in any kind of GraphQL api. First you fetch set of ids, then you want to hydrate them, as in fetch the data that correspond to a given id. With _AppSync_ you can do it using two techniques.

### Creating a special type

This is something I usually would avoid as it leaks implementation details to your schema. Either way, it is a solution to the problem so let's take a look.

Let's say you are working on two _Queries_ that return _TweetsPage_

```graphql
type Query {
  getLikes(userId: ID!): TweetsPage!
  getTweets(userId: ID!): TweetsPage!
}
```

The second query, `getLikes` is much more complex(assumes that you are not using single table design) than the `getTweets` query. This is because the `getLikes` involves first checking which tweets are liked, then hydrating those tweets. With the `getTweets` we can perform a single _Query_ operation and get the results.

Since we cannot fulfill the `getLikes` query in 1 step (like we can with `getTweets`), we have to introduce intermediate type which will hold the _unhydrated_ results.

```graphql
type Query {
  getLikes(userId: ID!): UnhydratedTweetsPage!
  getTweets(userId: ID!): TweetsPage!
}
```

Now we can directly attach a resolver to the `UnhydratedTweetsPage` and not interfere with `TweetsPage`. As you can see, we leaked the implementation details to our schema. From the clients perspective the results are _hydrated_.

### Pipeline resolvers

This is the approach I would recommend anyone uses.

With pipeline resolvers, you do not have to introduce intermediate types, you instead introduce intermediate resolver which will hydrate the results.

With _pipeline resolvers_ we would be able to get back to our previous schema definition

```graphql
type Query {
  getLikes(userId: ID!): TweetsPage! # This Query will be resolved using pipeline resolvers
  getTweets(userId: ID!): TweetsPage!
}
```
