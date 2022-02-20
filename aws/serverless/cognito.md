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

## Resources

- https://bobbyhadz.com/blog/aws-cognito-amplify-bad-bugged?utm_source=newsletter&utm_medium=email&utm_content=offbynone&utm_campaign=Off-by-none%3A%20Issue%20%23133
