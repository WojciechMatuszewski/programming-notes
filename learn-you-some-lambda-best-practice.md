# Lambda best practices

## Tuning function memory

- the more memory your lambda has and the longer is runs it will cost you more

- memory allocation also controls CPU allocation

- use **aws lambda power tuning tool** to see memory vs cost for your function

- use **lumigo-cli** to make your life easier. This tool is a collection of utility tools for AWS and also includes the aforementioned _power tuning tool_

- within **lumigo-cli** you can set different profiles and visualize results, for example:

  > lumigo-cli powertune-lambda -n io-bound-example -r eu-central-1 -s speed

## Cold Start

- there are 2 factors that pay a role when it comes to cold starts: **duration** and **frequency**. You should mainly **focus on duration** since **frequency is usually out of your control**.

- look for **init duration** within **cloud watch logs** or an **initialization segment** within **XRay**.

- previous methods were only for a single function on a single invocation resolution. **Use lumigo-cli to see cold starts for all your functions on all your invocations**.
  > lumigo-cli analyze-cold-starts

### Initialization time

- **adding more memory (thus CPU) WILL NOT AFFECT THE SPEED OF THE INITIALIZATION**. This is because lambda is run at full power, always, at initialization.

- it's **always faster** to load **dependencies from a layer**

- **unused dependencies** are **irreverent when it comes to initialization**

- only require what explicitly need

### Measuring cold start performance

- you can simply invoke your function

- to **force cold start** you can just **change / add env variable**

- use **lumigo-cli** to **measure particular lambda cold starts**
  > lumigo-cli measure-lambda-cold-starts -r eu-central-1 -n io-bound-example -i 100

### Improving cold starts with Node.js

- put stuff into `dev-dependencies`. This is quite dangerous since the version can change. As an alternative **publish lambda-sdk as lambda layer**.
- only require what you need
- use `webpack` to bundle your dependencies

## Provisioned Concurrency

- you cannot use it on `$Latest`
- makes it so there are X containers within your function, warm, ready for execution
- this means that you can have functions which are not handled by provisioned concurrency.

- there is one **gotcha that might throw you off**. When you invoke your function **it will still report old `Init Duration`**. This is due to the fact that **the reported `Init Duration` happens at the time the concurrency is provisioned**.

- you can have **alias pointing to version which has provisioned concurrency enabled**, but always remember that **provisioned concurrency is always assigned to version**

### When to use Provisioned Concurrency

- for really strict latency requirements
- for spikes of traffic
- when cold starts stack up (multiple sync lambda calls)

### Tips

- use **weighted distribution with alias** to make sure you do not loose provisioned concurrency executions when updating versions

- use **scheduled auto scaling** for spikes of traffic you know ahead of time.

## Enabling HTTP keep-alive

This used to be the case where sdks for lambdas were not using `keep-alive` by default. This would result in having to establish a handshake with every request.

**That is no longer the case** and **both js and golang sdk have keep alive enabled by default**.

## Fan-out

You have `ventilator` that spits out work to multiple workers. Since lambda scales horizontally , it's an ideal candidate for a worker pool.

- for **low traffic** use **SQS / SNS / EventBridge**

- for **high traffic** use **Kinesis**

It's also important to understand how concurrency scales in regard to messages per second.

- `SNS` and `EventBridge` scale concurrency linearly, that means there is **no batching**

- `SQS` **has batching** which means that the scaling for more gradual

- `Kinesis` scales more aggressively than `SQS` since your can have batching on `ShardID` level, but still not as aggressively as `SNS` or `EventBridge`

## Controlling Concurrency

This concept is quite important since having no control over concurrency can be a problem to your downstream system.

- with `Kinesis` the flood of messages can be amortized by the ability to hold messages for up to 7 days.

- `SNS / EventBridge` can be a bit problematic. There are retry mechanism in place, but what if they fail?. Remember that `SNS` wait for the acknowledge message for 15s, otherwise will consider given message as a failure in delivery. You can always rely on `dead-letter queue`

There might a problem with `Kinesis` where one _poison message_ can really screw you up, being delivered over and over to your lambda function. You can control this behavior though using `on-failure destination` and `maximum age of record`.
