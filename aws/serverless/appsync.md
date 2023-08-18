# AppSync

## JavaScript / TypeScript resolvers

VTL is no longer the single way to write resolvers. At first AWS announced the support for [JavaScript pipeline resolvers](https://aws.amazon.com/blogs/aws/aws-appsync-graphql-apis-supports-javascript-resolvers/), and then enchanted the feature with the support for [JavaScript unit resolvers](https://aws.amazon.com/blogs/mobile/aws-appsync-now-supports-javascript-for-all-resolvers-in-graphql-apis/).

This is great news, as writing the resolvers is much easier now than it was before. Especially given the ability to write them in TypeScript and lint the code using `eslint`.

### Gotchas

- The **pagination token returned from DynamoDB is encrypted** and I **could not find a way to decrypt it**. Usually this would not be a problem, but it might be [since DynamoDB returns the `NextToken` even when there might not be "next results"](https://stackoverflow.com/questions/64422854/appsync-pagination-issue-where-no-items-are-returned-but-a-nexttoken-is).

  - This is the price we pay of abstracting the DynamoDB calls. You can [see the solution to this problem here](https://github.com/WojciechMatuszewski/dynamodb-pagination-gotcha). The TLDR is: **use AWS Lambda AppSync integration instead of VTL resolver for this particular use-case**.

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
    $util.qr($result.put("__typename", "OtherProfile"))
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

**I would only use this if my API will never, ever be public**.

### Pipeline resolvers

This is the approach I would recommend anyone uses.

With pipeline resolvers, you do not have to introduce intermediate types, you instead introduce intermediate resolver which will hydrate the results.With _pipeline resolvers_ we would be able to get back to our previous schema definition

```graphql
type Query {
  getLikes(userId: ID!): TweetsPage! # This Query will be resolved using pipeline resolvers
  getTweets(userId: ID!): TweetsPage!
}
```

There is a bit more configuration to this approach, but I think the benefits heavily outweigh the negatives.

## Depth limiting

If you are not careful and your schema is recursive, an attacker might run costly queries against your endpoint.

Imagine picking friends of a friend of a friend ... and so on. Such requests might fire expensive back-end computations, which will cause issues and high costs if unbounded.

There are many ways of guarding against such behavior. Some resort to computing _query cost_. A user has a "currency" assigned and spends that "currency" making queries. Other solutions employ other heuristics.

We can implement one of such heuristics in _AppSync_. I personally feel like this way of checking query complexity might not be sufficient for all but will most likely get the job done in 80% of scenarios.

### The `selectionSetList` context attribute

One might use the `selectionSetList` VTL context attribute to determine the cost of a query. The first axis to look at would be the **length of the list**. The second would be **the depth of the queries contained in the list**.

Here is an example `$context.info` variable that contains the `selectionSetList`

```json
{
  "fieldName": "getPost",
  "parentTypeName": "Query",
  "variables": {
    "postId": "123",
    "authorId": "456"
  },
  "selectionSetList": [
    "postId",
    "title",
    "secondTitle"
    "content",
    "author",
    "author/authorId",
    "author/name",
    "secondAuthor",
    "secondAuthor/authorId",
    "inlineFragComments",
    "inlineFragComments/id",
    "postFragComments",
    "postFragComments/id"
  ],
  "selectionSetGraphQL": "{\n  getPost(id: $postId) {\n    postId\n    title\n    secondTitle: title\n    content\n    author(id: $authorId) {\n      authorId\n      name\n    }\n    secondAuthor(id: \"789\") {\n      authorId\n    }\n    ... on Post {\n      inlineFrag: comments {\n        id\n      }\n    }\n    ... postFrag\n  }\n}"
}
```

Looking at the `selectionSetList`, we can see that the maximum depth of the query is three (three levels of nesting). How can we programmatically probe for that information?

[This blog post shows you how to do just that](https://robertbulmer.medium.com/aws-appsync-rate-and-max-depth-limiting-c536e5ba43d6).

Here is the template.

```vtl
#set($selectionSetList = $ctx.info.selectionSetList)

#foreach ($item in $selectionSetList)
    #if($item.matches(".*/.*/./."))
        $util.error("Error: Queries with more than 3 levels found. At level - $item")
    #end
#end#return($ctx.prev.result)
```

Not that scary, is not it?
