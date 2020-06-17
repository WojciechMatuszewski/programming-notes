# Day Two

## Secrets

- use SSM, secrets manager if you need rotation

- you can either fetch and decrypt the secrets **at deploy time, or during runtime**.

- avoid having your secrets in plain sight within `.env` of your function. This is a mistake you can make easly.

## VPC

- when deploying to multiple regions and your lambda should be within VPC, use `serverless-vpc-discovery` plugin.

- remember that `NAT Gateway` and `Interface Endpoints` costs per hour and data transferred.

## Security

- apply per function policy

## Leading pratices

- use CF caching. When you deploy edge optimized API, APIGW uses CF under the hood but not the cashing mechanism from CF.

- watch out for `Iterator Age` while looking at `Kinesis` metrics. High `Iterator Age` can indicate problems.
  > Age is the difference between the current time and when the last record of the GetRecords call was written to the stream.

## Design Patterns

- with **dynamoDB streams** using lambda as target is free.

- `Kinesis` is much cheaper for high throuput workloads than SNS or DynamoDB.

- `step functions` are great, but **expensive**. Carefuly evaluate your options. You can also **look into `express workflows`**.

- try out the `saga pattern`.
