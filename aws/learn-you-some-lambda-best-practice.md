# Lambda best practices

There is a difference between `exports =` and `module.exports =`. **Always preffer `module.exports` because it's the `module.exports` that gets returned**.
With `exports` you are mutating an `module.exports` and if there is an `exports` and `module.exports` within the same file, it's the `module.exports` that will be honoured.

## Tuning function memory

- the more memory your lambda has and the longer is runs it will cost you more

* memory allocation also controls CPU allocation

- use **aws lambda power tuning tool** to see memory vs cost for your function

* use **lumigo-cli** to make your life easier. This tool is a collection of utility tools for AWS and also includes the aforementioned _power tuning tool_

- within **lumigo-cli** you can set different profiles and visualize results, for example:

  > lumigo-cli powertune-lambda -n io-bound-example -r eu-central-1 -s speed

## Cold Start

- there are 2 factors that pay a role when it comes to cold starts: **duration** and **frequency**. You should mainly **focus on duration** since **frequency is usually out of your control**.

* look for **init duration** within **cloud watch logs** or an **initialization segment** within **XRay**.

- previous methods were only for a single function on a single invocation resolution. **Use lumigo-cli to see cold starts for all your functions on all your invocations**.
  > lumigo-cli analyze-cold-starts

### Initialization time

- **adding more memory (thus CPU) WILL NOT AFFECT THE SPEED OF THE INITIALIZATION**. This is because lambda is run at full power, always, at initialization.

- it's **always faster** to load **dependencies from a layer**

* **unused dependencies** are **irreverent when it comes to initialization**

- only require what explicitly need

### Measuring cold start performance

- you can simply invoke your function

* to **force cold start** you can just **change / add env variable**

- use **lumigo-cli** to **measure particular lambda cold starts**
  > lumigo-cli measure-lambda-cold-starts -r eu-central-1 -n io-bound-example -i 100

### Improving cold starts with Node.js

- put stuff into `dev-dependencies`. This is quite dangerous since the version can change. As an alternative **publish lambda-sdk as lambda layer**.

* only require what you need

- use `webpack` to bundle your dependencies

## Layers

- use `Lambda layers` to abstract common logic.

* these can be anything that you want to get into execution environment, literally, for all lambda cares it could be a binary file.

- layers are versioned automatically, you can reference up to 5 layers.

* when working with `node.js` you will have to set `NODE_PATH` environment variable:
  ```yml
  NODE_PATH: "./:/opt/node_modules"
  ```
  this is so that you can use both the root `node_modules` and the ones defined within your layers. A good explanation of `NODE_PATH`:
  > NODE_PATH is like the windows path environment variable. Whenever node can't find a file, it looks through the paths in the paths stored in the NODE_PATH variable.

## Provisioned Concurrency

- you cannot use it on `$Latest`

* makes it so there are X containers within your function, warm, ready for execution

- this means that you can have functions which are not handled by provisioned concurrency.

* there is one **gotcha that might throw you off**. When you invoke your function **it will still report old `Init Duration`**. This is due to the fact that **the reported `Init Duration` happens at the time the concurrency is provisioned**.

- you can have **alias pointing to version which has provisioned concurrency enabled**, but always remember that **provisioned concurrency is always assigned to version**

### When to use Provisioned Concurrency

- for really strict latency requirements

* for spikes of traffic

- when cold starts stack up (multiple sync lambda calls)

### Tips

- use **weighted distribution with alias** to make sure you do not loose provisioned concurrency executions when updating versions

* use **scheduled auto scaling** for spikes of traffic you know ahead of time.

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

* `SNS / EventBridge` can be a bit problematic. There are retry mechanism in place, but what if they fail?. Remember that `SNS` wait for the acknowledge message for 15s, otherwise will consider given message as a failure in delivery. You can always rely on `dead-letter queue`

There might a problem with `Kinesis` where one _poison message_ can really screw you up, being delivered over and over to your lambda function. You can control this behavior though using `on-failure destination` and `maximum age of record`.

## Service Proxy

This is where you bypass compute layer (your lambda) and go directly to a service. The example here would be going to eg. DynamoDB directly.

You **should consider** using **service proxies** when you are **concerned about cold start overhead or a burst limit of lambda**.

This solution is not without downsites though:

- you loose exponential backoff and retries (SDK is doing that for you)

* you loose logging

- you loose the ability to use 3rd party tracing tools

* you loose the ability for having custom error handling logic

- **VTL is hard to test / write**

## Load testing

- model your load test to resemble the actual traffic. Do not create artificially steep concurrency curve.

* use user stories for testing. Do not hammer a specific endpoint. By doing that you are testing the AWS and not your service.

- artillery.io can be helpful. As an alternative you should look into `serverless-artillery`

## Handling RDS connections

- default config for RDS is just bad for lambda.

* use **RDS Proxy**.

## Lambda Destinations

- works only for **async / stream based invocations**. So SQS, SNS, Kinesis, DynamoDB streams ...

* gives you much more information than DLQ, even provides you the stack trace.

- prefer Lambda Destinations when you can.

* remember that Lambda Destinations are not only about failure. **There is also `onSuccess` event you can configure**. You would use it when you have `one hop` situation. Other than that **prefer step functions**.

## Async with Lambda

- Lambda async invoke: buffer is managed for you

* SQS event source to Lambda: you manage the buffer and DQL

- Stream event source to Lambda: you manage the bufffer, each shard is like a FIFO queue. With Stream data there is a concept of **poison pill**. This is where you cannot, for some reason, process batch of requests and due to this, you cannot make any progress.

## Secret Management

### Storage

- there are 2 services which you should look at: `SSM Parameters Store` and `Secrets Manager`.

* `SSM Parameter Store` is mostly free, unless you want to use `advanced parameters` (larger payload).

- remember that `Secrets Manager` can **rotate credentials**. That is **applicable even for RDS credentials**.

### Distribution

- remember to **never store secrets as env variables within your function**.

* you should **fetch secrets at runtime**, cache and invalidate the cache every few minutes.

- if you are writting JS lambda, consider using `middy`. It is a great tool - easy to use, gets the job done!.

## APIGW Throttling gotcha

If you create your API with a default settings, each stage has the default `10k` req/s throttling limit. This is an **region wide limit!**.
This means that I as an attacker only need to **attack your 1 endpoint to take down all your APIS within a given region**.

Pretty bad huh?

- always use custom throttling limits

* consider using `serverless-api-gateway-throttling` plugin if you are working with `serverless-framework`.

- if you are using `AWS SAM` you can set the desired throttling on the `AWS::Serverless::Api` type.

Read more [here](https://theburningmonk.com/2019/10/the-api-gateway-security-flaw-you-need-to-pay-attention-to/)

## Multi-region

- you can create SNS topics in multiple regions and have 1 SQS subscribe to these topics. This can cause duplication though, watch out!

* there is a term `static stability`. This is where a system would continue to work even if a region failed.

![image](./assets/sns-multi-region.png)

## Partial Failures

- by default `Kinesis` will **retry until success**. This is no good. Luckily AWS introduced payload splitting and DLQ.

* SQS **fails the whole batch**. This is probably not what you want. You could:
  - set `batchSize:1`. This is bad for throughput.
  - have idempotent workflow so that you can process multiple messages multiple times
  - call `deleteMessage` yourself for the successful messages. This way they will not get retried. After that, you can safely return an error.

## Alarms

- one alaram should monitor different failure modes

* use alarams to tell if something is wrong, not what.

Notable alarms are:

- `concurrent exectuions`: about 80% of the regional limit
- `IteratorAge`: for `Kinesis`.
- `DeadLetterErrors`: for async lambda invocations. **You should have deadletter queue setup, even better use lambda destinations**.
- `Throttles`: self explanatory

## Logging

- use simple `os.Stdout`. Logs are written to `CloudWatch` logs asynchronously.

* there is a **cost of ingestion** when using `CloudWatch` logs.

- use **sampling** to **log only a part (given %) of given log level**.

* if you really need it, you can stream logs to a 3rd party service like logz.io

## Distributed tracing

### `X-Ray`

- you should strongly consider using `xray-core` only.

* `X-Ray` is not that good. Your traces will be cut after publishing to `SQS` or `EventBridge`. You will have to then use different filter to find those. Weird.

- `service map` view suffers from the same problem. You no longer have to jump in between filters, but still the dots are _not connected_ if you will.

* some _flow arrows_ are even missing, WOW! (apigw => lambda => back to apigw).

- **`ServiceLens`** gives you different view (look inside CloudWatch section). Still not great, **it does not event support most of the services**.

* one huge benefit of `X-Ray` is that the traces are saved _asynchronously_.

### `Lumigo`

- specialized service, seems better than `X-Ray`.

* no code changes required, you need to install serverless-plugin

### `Epsagon`

- specialized service, UI not that nice as `Lumigo`.

* a bit better capturing than `Lumigo`.

### `Thundra`

- you need to wrap lambda handlers with the `thundra` client.

* not that good reporting as `Epsagon` and `Lumigo`.

## Correlation IDS

- used to make sense of our logs.

* are used to **keep track of a origin request, when there is a lot of services involved in the chain**.

- this is very **useful with `Logs Insights`**. You can create simple filter:

  ```
   fields @timestamp, `x-correlation-id`
  | filter `x-correlaction-id` = 'REQUEST_ID'
  ```

![correlaction-ids](./assets/correlation-ids.png)

## Lambda powertools (node.js)

- multiple _middlewares_ for your lambda functions

* pretty nice packages, help you with `correlaction ids`.

- sampling is done at _transaction_ level, not at single invocation level.

## Cost

- use **AWS Billing** and `tags`!

* use **Cost Explorer** for very detailed graphs.

- **rightsize your lambda invocations!**.

* use `Step Functions` only for **core business workflows**. They are expensive!. You could also use `Express workflows`.

- CloudWatch charges per dimmension!. **DO NOT USE REQUEST_ID AS DIMMENSION!**.

* `NAT Gateway` can be **very expensive**.

- use **sampling of logs**.

* set _retention period_ on `CloudWatch logs`.

- `Lambda` is likely the cheapest part of your infrastructure. use `HTTP APIs` whenever you can.

* in **high thruput scenarios ALB is much cheaper than APIGW REST and HTTP APIs**.
