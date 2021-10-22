# CloudWatch

## Latency percentile

While using the _CloudWatch_ service you might encounter graphs with `p99` or `p50` axis.
These are so called _percentiles_ and are very useful in gauging the performance of your system.

In the **context of latency**, the **`p99` means that 99% requests are processed in less than the `p99` latency**.
Here is an excerpt from the [_Google Cloud_ docs](https://cloud.google.com/spanner/docs/latency) that explains this topic very well.

> 99th percentile latency: The maximum latency, in seconds, for the fastest 99% of requests. For example, if the 99th percentile latency is 2 seconds, then Cloud Spanner processed 99% of requests in less than 2 seconds.

### Why should you care?

Should not we be looking at the average? Why should I be concerned with 1% of requests that are slower than X?
**In my opinion** the answer is: **it depends**.

If you are just starting out and in the process of forming your business, you should not really care.
In the contrary, **if the traffic to your service is high you should care**. This is because of the _economy of scale_.
**The more users you have, the more customers will experience the worst case scenario latency**.

## CloudWatch Insights

So you are on-call? You better know how to use this tool.
Otherwise, you will spend too much time trying to find relevant logs. The clock is ticking while you are doing this.

In this document, I will share some of my findings regarding the _Query Syntax_ with sample queries examples.

### Use the `@logStream` field

So an error occurred and you got paged. You fire off the _CloudWatch Insights_ and filter for errors.
You see the error logs, but where are logs that are _related_ to that error log?

You cannot really query _related_ logs by using _CloudWatch Insights_ natively.
But there is a way, to get to them easily - display `@logGroup` field within your query.

Given this simple query

```shell
filter level = "error"
| display @timestamp, @logStream
```

You will have a link to the log stream presented to you. When you click that link, you will be redirected right to that log.
You might want to copy the `@timestamp` value as well, to orient yourself on the new page.

**The link will not be clickable if you selected multiple log groups for a given query**.

### Searching through JSON logs

I really hope that you are logging in a way that produces structural logs. For your sanity sake, I really hope that the format is JSON.
Either way, the JSON logs can be easily filtered. To do that, just reference the fields directly, for example

```shell
filter error.statusCode = 404
| display error.statusCode
```

One tip here, would be to run your query without any filters applied first. By doing it, _CloudWatch Insights_ can gather
information about the data that is present within your logs. With that, the query console will have code-completion capabilities at your disposal.

### Searching for lambda timeouts

These things will happen, so it's worth knowing how to find relevant logs for timeouts.

One strategy would be to log something just before the lambda is about to timeout.
Here, the strategy is simple, run a timer with a deadline of `remainingTimeInMilis - DELTA).

A simple JS implementation

```js
setTimeout(() => {
  log.info("TIMEOUT!");
}, deadline);
```

You can pull the `remainingTimeInMilis` from the `context`.

If you are not doing that, you can still get the relevant log. When lambda times out, it will produce a log containing `Task Timed Out ...`.
Then all you have to do is get the `@logGroup` field and go to _CloudWatch Logs_ console.

## Contributor Insights

You probably are aware that you should not use high cardinality data as a dimension value within your metrics (if you do not, now you know). If you do so - think using `requestId` as a dimension for a metric, your AWS bill will explode.

> CloudWatch treats each unique combination of dimensions as a separate metric

And one of the things you pay for is the amount of metrics you have :)

So what is the alternative then? For tracking high cardinally data I would recommend using _CloudWatch Contributor Insights_.

The _Contributor Insights_ allow you select log groups from which the data will be gathered.
You can pick and choose properties from those log groups and then create aggregates based on their values.

So now, the `requestId` is not so scary, you just pick it as a field upon which the _Contributor Insights_ will create an aggregate.

Of course, this way of presenting information is mostly used for _count-based_ metrics. But either way it can be useful (you can achieve the same results by using _CloudWatch Insights_ but you would have to create a relatively complex query yourself).
