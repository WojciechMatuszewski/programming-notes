# CloudWatch Insights

So you are on-call? You better know how to use this tool.
Otherwise, you will spend too much time trying to find relevant logs. The clock is ticking while you are doing this.

In this document, I will share some of my findings regarding the _Query Syntax_ with sample queries examples.

## Use the `@logStream` field

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

## Searching through JSON logs

I really hope that you are logging in a way that produces structural logs. For your sanity sake, I really hope that the format is JSON.
Either way, the JSON logs can be easily filtered. To do that, just reference the fields directly, for example

```shell
filter error.statusCode = 404
| display error.statusCode
```

One tip here, would be to run your query without any filters applied first. By doing it, _CloudWatch Insights_ can gather
information about the data that is present within your logs. With that, the query console will have code-completion capabilities at your disposal.

## Searching for lambda timeouts

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
