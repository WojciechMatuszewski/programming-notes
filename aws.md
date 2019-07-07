# AWS stuff

## AWS Cognito

This service is used for authentication but there are 2 variations:

- user pools
- federated identities

What's the difference between those two?

### AWS Cognito UserPools

**This service is used for authentication (adding one to your application)**.

- sign-in
- sign-up
- password policies
- social-login with Facebook or Google.

Using this service enables you to grant users access to your application (previously mentioned auth flow) **and also grant access to AWS Cognito Federated Identities**.

### AWS Cognito Federated Identities (aka Identity Pools)

This service allows users to **access AWS services under your AWS account**. It's like lending someone your account without giving them password and login.

## AWS IAM roles

Think about having multiple developers on team. You have one intern comming in and you do not want him to be able to poke auth stuff, but you still want him to contribute to the project. You create specific IAM role for that user on your AWS IAM console.

You can do the same with group of people.

So using **AWS Cognito Federated Identities** creates IAM roles.
