# Lambda extensions

## What is it

Lambda extensions are used to augment your Lambda functions. They allow you to plug into the lambda life-cycle
and do things based on these events.

Main use-cases are

- monitoring
- caching (only on a given container basis)

They only run when there is something to do. They can run **when the given lambda execution environment starts** and can run after your handler has returned the data.
They do not block your lambda, as soon as your lambda returns the data, you are not paying for the lambda.

## How to deploy an extension

Lambda extension is just a layer with a special directory structure.
Lambda extensions also work with container deployments.

## Types of extensions

There are two types of lambda extensions.

- **internal extensions** run as separate thread, but **in-process with your main code**.
  These are mainly scripts used to configure your runtime of choice environment. Think configuring the `max-memory` Node.Js setting.
  Since most common flags can be set through environment variables, you would most **use this type of extension for things that you CANNOT set through environment variables, like runtime flags**

- **external extensions** run as a **separate process**. These are very powerful, but share the CPU, memory and IAM stuff.
  You can basically run any code your want here. The AWS itself already uses _external lambda extensions_ to gather additional telemetry information whenever you select the "Advanced CloudWatch Insights" option for a given lambda.

## Pluming with external extensions

With the release of lambda extensions, a new APIS were added so that the extension can talk to Lambda service.
These are the _Logs API_ and the _Extensions API_. Of course, the _Runtime API_ was not changed, the external extension does not have the ability to talk to the runtime API.

### Collecting logs

- the extension can submit a _subscribe request_ so that the logs that the lambda service will push to cloudwatch are also pushed to a given endpoint

- you are responsible for spinning the endpoint the lambda service will push the logs to

## Pricing

The pricing model is the same as AWS Lambda.

## Gotchas

- The **extension initialization phase might influence the duration of the cold start**. This is because the _init_ phase of the execution environment is considered done only when the all three components (runtime, your lambda and the extension) are ready.
