# AWS & Serverless

## Serverless Lambdas

`serverless framework` is a tool to create lambdas (you can even use webpack
:o), **configure AWS services** (Cloud Formation).

### Sync and Async Lambda mode

#### Sync mode

> You call a lambda and you are waiting for the response to come back.

-   client waiting for the response.

#### Async mode

> You fire it and forget about it.

-   send me an email in the background.
-   update dynamoDB (like a scheduled job).

### Cold State

So you have your lambda in the cloud. AWS **spins your lambdas** in containers.
These containers gets discarded (it's like 5-20 mins of inactivity). If AWS
decided to discard container your lambda was in it's in a _cold state_. When
called **a new container must be created** resulting in slow response speed.

### Lambda@Edge

> Instead of going to the origin server request will go to the nearest node
> available.

What's more exciting is that you can do computations there now. Previously it
was not possible.

## AWS Cognito

This service is used for authentication but there are 2 variations:

-   user pools
-   federated identities

What's the difference between those two?

### AWS Cognito UserPools

**This service is used for authentication (adding one to your application)**.

-   sign-in
-   sign-up
-   password policies
-   social-login with Facebook or Google.

Using this service enables you to grant users access to your application
(previously mentioned auth flow) **and also grant access to AWS Cognito
Federated Identities**.

### AWS Cognito Federated Identities (aka Identity Pools)

This service allows users to **access AWS services under your AWS account**.
It's like lending someone your account without giving them password and login.

## AWS IAM roles

Think about having multiple developers on team. You have one intern coming in
and you do not want him to be able to poke auth stuff, but you still want him to
contribute to the project. You create specific IAM role for that user on your
AWS IAM console.

You can do the same with group of people.

So using **AWS Cognito Federated Identities** creates IAM roles.

## Cloud Formation

So AWS has a lot of features. A shit ton of them. `Cloud Formation` is an API
that ties them together. Instead you going to their website and clicking stuff
you could write one config file and send that to amazon and given services will
be configured. `serverless framework` is an abstraction over that since even
writing that file _is really hard_.

## Api Gateway

Well, a gateway. You can think about it as **a proxy**. That proxy could, for
example, be used to **route request to our lambdas**.

> Api Gateway sits between your client and your lambda. It's a long lived
> process (like a normal server).

## VPC (Virtual Private Cloud)

Whenever you create aws account you get one automatically. **It's a collection
of tools that basically allows your lambda to be exposed to the internet**.

## EC2

So long the era of costly private servers. With EC2 you can rent a virtual
machine (of different kind, optimized for different things or general use) for
cheap and play with that. You can control the exposition to the internet by
defining ports and such.

## S3 (Simple Storage Service)

-   safe place to store your files
-   data is spread between multiple devices and facilities
-   unlimited storage
-   files are store in something called **buckets**

### Glacier and Glacier Deep Archive

These are used for storing archive data. Something you do not need to get right
away. It is very cost effective.

## EBS (Elastic Block Store)

> persistent block storage volumes for use with with EC2.

So this is basically virtual hard disk in the cloud. Aws provides a variety of
disks, from SSDs to normal Magnetic drives.

## CloudWatch

CloudWatch is used for monitoring (mostly performance). You can check your
lambdas logs there and other stuff.

## CloudTrail

Do not confuse with `CloudWatch`. Think of this as a camera looking at you and
records what you are doing.

> Every time you create S3 bucket or make an API call to AWS it records that.

You can later identify which users did what.

## Databases

Amazon has a plethora of different databases available. From traditional, SQL
(relational) databases to Key-Vale ones (DynamoDB)

### RDS

RDS is the name for relational data bases. They have 2 key features:

-   Multi-AZ for disaster recovery, this is basically means that when your
    database fails, Amazon will dynamically update EC2 DNS to point to the other
    database.
-   Read-Replicas for performance. One database is connected to the other one.
    That second one holds copies of your writes to the first one. In case of
    failure Amazon **will not** handle it automatically here.

### NoSql/JSON

Mainly `dynamoDB`.

### Aurora

-   AWS solution for Relational-Database
-   Much faster than eg. `MySQL`
-   Costs much less than traditional relational database
-   Compatible with existing relational databases

### Misc

#### RedShift

Amazon's answer for data warehousing, doing online analytics processing with very complex queries.

#### Elastic Cache

Sits between your website and database. Basically cache as a service. Basically is a Redis (or Memcached) managed by AWS.
