# Day One

## Basics

- use one of the method on the `context` to see how much time is left before timeout

- use **lambda destinations** whenerver you can.

- use **lambda layers**.

- split lambda traffic for canary deployments.

- with lambda, if your runtime is not supported, you might look at the community-created runtimes

- **DO NOT default aws-sdk which is included within lambda**. This sdk is only for the console and **SHOULD NOT BE USED FOR PRODUCTION**. Included SDK version can change without notice.

## Testing

- use `Pact` for consumer driven testing.

## CI/CD

- use multiple accounts to avoid aws limits

- do not easly discoveribility roles for deployment accounts

- look into **ABAC**.

- if you need load testing, look into **serverless-artillery**.

- when using `npm` look into `npm ci`, if you are using `yarn`, use `--frozen-lockfile` flag to make your builds reproducible.

## Logging

- remember about expiring logs, there is a cost for keeping log storage.

- `XRay` is kinda useful but somewhat limited. Does not support every service.

- use `correlation ID`. This is the `ID` that helps you string different logs together. Without any kind of system this might be hard for distributed services.

## Metrics

- use `Embedded Metric Format`. This allows you to push logs to metrics without calling AWS service directly. There is no performance overhead.

## Retries

-
