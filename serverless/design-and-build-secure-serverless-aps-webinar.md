# Webinar

## Security Background

### Shared Responsibility Model

- AWS is protecting the infrastructure that runs the services

- customers are responsible for building secure applications

## IAM for Lambda

- apply always the least privilege iam permissions

- use **SAM parameterized policies** to develop faster

### Execution Policy

- your lambda is able to do stuff

> Lambda function A can read from DynamoDB table users

### Function

- some services are allowed to trigger your lambda

> Actions on bucket X can invoke lambda function Z

## APP Configuration

### Secrets

- do not hard code secrets within the code

### Lambda Env variables

- key value pairs passed to a function, **can be optionally encrypted using KMS**.

### SSM Parameter store

- can be useful for sharing the same variable for multiple lambda functions

- is actually **free service**

- parameters can be encrypted

- parameters can be tagged

- granular permissions for parameters

- can be used for feature flags

- you might hold secrets here (simple, **without rotation**)

### Secrets Manager

- **manage, retrieve, version and rotate credentials**

- granular permissions

- supports RDS by default

- secrets manager is exposes much more complex controls over passwords that SSM Parameter Store

- **not a free service**

### Using Secrets Manager inside Parameter Store

- use special prefix for parameter inside Parameter Store `/secretsmanager/YOUR_STUFF`

### Using Parameter store

- remember to **have permissions to decrypt the parameter if your parameter is encrypted**

## AWS RDS Proxy

> Pools and shares connections to make applications more scalable, resilient and secure

- instead of talking to the DB you talk to the proxy

- it solves the _pooling problem_. This is where db has a limited number of connections that it can keep open. This can be problematic since serverless workflows have tendency to have concurrent workers

## Closing notes

- workshop which you can take [link](amzn.to/serverless-security)

## QA

- do not be afraid of friction with IAM. It takes time to know all the permissions, **do not use \***
