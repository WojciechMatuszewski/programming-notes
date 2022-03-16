# Cogito

## Identity Pools vs User Pools

Ever since I learned about AWS Cognito, I have wondered what the difference is between the User Pools and the Identity Pools. Today, I have finally decided to capture my knowledge regarding this topic in this section. I hope you will find it helpful, dear reader.

**Think of the User Pools as a quasi database for your AWS Cognito users**. It can contain metadata, called _attributes_, to group users into _groups_. Apart from keeping the users' data, you can also specify password policies.

On the other hand, there are the Identity Pools. **Think of the Identity Pools as the _credential vending machine_ for Amazon Cognito or unauthenticated users**. After you have the credentials, you can **directly call other AWS services the credentials give you access to**.

I would argue that if you are aware of this distinction, you know enough to get you 99% there in most cases.

## Authorizer

The default, natively implemented Amazon Cognito Authorizer will do for most cases. But for some cases, it might not be enough, especially if we want to pass additional context to the resource we are "guarding".

You can implement your custom authorizer.

1. Check the validity of the JWT.
1. Pass the necessary context.
1. Ensure that the policy is scoped correctly.

One use case might be adding additional context, like [in this article](https://aws.amazon.com/blogs/compute/capturing-client-events-using-amazon-api-gateway-and-amazon-eventbridge/).

## The _ID Token_ and the _Access Token_

After you successfully go through the _Amazon Cognito_ authentication flow, you get two tokens from the service – the _ID Token_ and the _Access Token_.

The significant difference between them is that **the _Access Token_ does not contain any information about the identity, while the _ID Token_ does**.

While you could, in theory, use either of them for authorization purposes, [the best practice](https://auth0.com/docs/secure/tokens?_ga=2.253547273.1898510496.1593591557-1741611737.1593591372#id-tokens) is to use **_Access Token_ for making requests to protected resources**. Leave the _ID Token_ for getting the user-specific information.

- The **_ID Token_** is meant to be used by the application only (most likely the front end of your application).

  You **should not use the _ID Token_ as a way of gaining access to an API**. The _ID Token_ contains info about the user. We would not want to expose that unnecessarily.

- The **_ID Token_** contains list of **claims**. Claims are _assertions_ on a particular subject. Having an _email_ claim means that Amazon Cognito asserted that the _ID Token_ receiver has a given email (the value is the value of the _email_ claim).

- You **cannot revoke the _ID Token_**. You will have to wait for it to expire (the expiration TTL is 60 minutes)

- The **_Access Token_** is meant to be used for API **authorization**. It **does not contain any sensitive user info** but instead contains _scopes_ and other attributes.

- The **_Access Token_** contains **scopes**. Scopes **define a resource that the user has access to**.

- There is a problem with revoking _Access_ and _ID_ tokens. Sadly it's not supported by Cognito.

  > Amazon Cognito now supports token revocation, and Amplify (from version 4.1.0) will revoke Amazon Cognito tokens if the application is online. This means Cognito refresh token cannot be used anymore to generate new Access and Id Tokens.

  > Access and Id Tokens are short-lived (60 minutes by default but can be set from 5 minutes to 1 day). After a revocation, these tokens cannot be used with Cognito User Pools anymore, however they are still valid when used with other services like AppSync or API Gateway.

  > For limiting subsequent calls to these other services after invalidating tokens, we recommend lowering token expiration time for your app client in the Cognito User Pools console. If you are using the Amplify CLI this can be accessed by running amplify console auth.

### The Amazon Cognito authorizer

By default, the Amazon Cognito APIGW Rest authorizer will honor the _ID Token_ and the _Access Token_ (see [this GitHub repository](https://github.com/WojciechMatuszewski/apigw-mockintegration-cognito)).

Depending on your environment, this might or might not be a problem. Remember that we should not be using the _ID Token_ as a means of authorization.

To make sure the authorizer works well with the _Access Token_, one has to amend the configuration of the Amazon Cognito User Pool. Again, see [this GitHub repository](https://github.com/WojciechMatuszewski/apigw-mockintegration-cognito) for more information.

## Resources

- https://bobbyhadz.com/blog/aws-cognito-amplify-bad-bugged?utm_source=newsletter&utm_medium=email&utm_content=offbynone&utm_campaign=Off-by-none%3A%20Issue%20%23133
