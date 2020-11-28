# Testing Serverless

## Introduction

Testing. One of my favorite things to do actually. The more software I write, the more I love writing tests. Features are the things that put food on our tables, but the tests... the tests are the things that keep everything in check, in order.

Many of you are probably familiar with the notion of _unit tests_. These are small, **fast**, tests that check if a given _unit_, be it a function or a _Component_, works correctly.

When it comes to serverless, we still have _unit tests_, but they play lesser role in overall testing story. With us using managed services more and more, it's the _integration_ and _end to end_ tests that matter the most. 100% of coverage on your lambda handler is almost worth nothing if it cannot get the data from _DynamoDB_ because of missing IAM permissions.

This file contains all the knowledge I have about testing serverless apps. We are going to start at the most basic level, I'm talking a simple API with _API Gateway_ and _DynamoDB_, and end at _SQS_ and other batch processing services.

## Testing simple APIS

This is where everything makes sense. We have our simple architecture with _API Gateway_ fronting one lambda function (which could be skipped) that talks to _DynamoDB_

### Unit tests

**This section assumes that you have nothing (PR scope) deployed to the cloud**.

Even though we are writing _unit tests_ here, I would strongly advice you **against mocking anything**. If you want to test the _integration_ with the _DynamoDB_ you can do that in the _integration test_.

Let's say your handler looks like this

```ts
const handler: APIGatewayProxyHandlerV2 = async (event) => {
  // some logic

  await docClient.put().promise();

  return; // some return stuff
};
```

This is a very simple handler. To test it, you might be tempted to just invoke the handler with the event in your test suite.
While this is a strategy, we would be venturing to _integration testing_ territory.

Instead of doing that, extract the logic you have in your handler, to a separate function.

```ts
export function performMyLogic() {
  // some logic
}

const handler: APIGatewayProxyHandlerV2 = async (event) => {
  const logicResult = performMyLogic();
  // do something with the `logicResult`

  await docClient.put().promise();

  return; // some return stuff
};
```

Now all you have to do in your tests, is to test (hopefully) a pure function of `performMyLogic`.

```ts
import { performMyLogic } from "./handler";

describe("performMyLogic", () => {
  test("works as expected", () => {
      expect(performMyLogic()).toBe(...)
  });
});
```

That is it. Nothing to deploy, test is fast and clean.

### Integration tests

**This section assumes that you have all the PR infrastructure that you need deployed to the cloud**.

This is where you are going to actually invoke your handler. Again, **no mocking**. We will be talking to a real, existing, deployed, _DynamoDB_ table.

First, we need the event. There is an easy way of generating one. Please refer to the [sample GET event](https://github.com/awsdocs/aws-lambda-developer-guide/blob/master/sample-apps/nodejs-apig/event-v2.json). You will need it (change it to POST event, let the _TypeScript_ guide you).

Since I'm writing a simple POST endpoint, the failure modes I have to consider are:

1. Is the `body` preset on the request?
2. Is the `body` malformed?
3. Is there data I need inside the `body`?
4. Is my lambda creating _DynamoDB_ entities correctly?

That is a lot to cover, but having those tests will really, really make you certain that the system you are building actually works.

A sample test which tests one of the failure modes would look something like this:

```ts
import { createHandler } from "../../functions/simple-api/handler";
import invokeEvent from "../../events/simple-api-event.json";

const tableName =
  "testingServerlessRefresher-simpleAPItable78ABB1E1-ORU7YEL1MKO7";

describe("returns bad request if", () => {
  test("body does not exist", async () => {
    const handler = createHandler(tableName);
    const event = { ...invokeEvent, body: undefined };

    const result = await handler(event, {} as any, () => null);
    expect(result).toEqual({
      statusCode: 400,
      body: JSON.stringify({ message: "Body not found" }),
    });
  });
});
```

Notice that I'm creating my `handler` by passing the `tableName`. I think having the handler as a _higher order function_ is a best practice. It allows you to inject dependencies as you see fit (in this case I'm hard coding the table name, normally the name would live in environment variables).

Let's talk about the fourth failure mode. Here we have 2 choices. We can either use the native `aws-sdk`, or a 3rd party library like [aws-testing-library](https://www.npmjs.com/package/aws-testing-library). The choice is yours, I will use the aforementioned 3rd party library.

The test would look something like this

```ts
jest.mock("ulid");

const mockUlid = mocked(ulid);

test("successfully writes to the DynamoDB table", async () => {
  const itemId = jest.requireActual("ulid").ulid();
  mockUlid.mockReturnValue(itemId);

  const event = { ...invokeEvent, body: '{"firstName": "Wojtek"}' };

  const result = await handler(event, {} as any, () => null);

  const expectedItem = {
    id: itemId,
    firstName: "Wojtek",
    lastName: "Matuszewski",
  };
  await expect({
    region: "eu-central-1",
    table: tableName,
    timeout: 5000,
  }).toHaveItem({ id: expectedItem.id }, expectedItem);

  expect(result).toEqual({
    statusCode: 200,
    body: JSON.stringify(expectedItem),
  });
});
```

Okay, so I used `jest.mock` for one thing, and that is the library I'm using to generate ids. Otherwise I would have to perform a scan on the _DynamoDB table_ which is not something you want to do.

The main takeaway here is that we are actually talking to a **real, deployed _DynamoDB table_**. This is the most important thing.

You might be asking asking yourself

> But what about permissions?

The _IAM_ is something that we will test in the _E2E_ tests. Lets do just that.

### E2E tests

Since integration tests take a long time to run, it's vital to think when do you run those.
For some projects, it might be completely ok, to run these after merging to master (after _staging_ environment is deployed). For others, you might want to run these in your per PR ci/cd pipeline. Up to you.

In this category of tests, you will be testing the whole flow. That means making a POST request to the endpoint exposed by an _API Gateway_.

The test for this would look something like this:

```ts
import phin from "phin";

test("saves and returns the item", async () => {
  const response = await phin({
    url: "https://87wa3d2y4c.execute-api.eu-central-1.amazonaws.com/",
    method: "POST",
    data: {
      firstName: "Karol",
    } as any,
    parse: "json",
  });

  expect(response.statusCode).toEqual(200);
  expect(response.body).toEqual({
    firstName: "Karol",
    id: expect.any(String),
    lastName: "Unknown",
  });
});
```

This test might seem trivial, but **there is so much value that you gain by having this test**. If this test works, you know for certain that:

1. Your lambda function has correct IAM permissions to call the _DynamoDB_ table
2. Your lambda function has correct _environment variables_ for _DynamoDB_ related things specified
3. _API Gateway_ is actually fronting a correct handler

The power of e2e tests in a serverless app is immense. You them!

## Testing asynchronous flows

**This section assumes most, if not all, your PR infrastructure is deployed to the cloud**.

Tests for asynchronous flows will take a long time to run. This is by nature, we are talking _asynchronocity_ here. I also think there line between _integration_ vs _E2E_ test is a bit blurred here. I'm not going to be dividing this section as before.

While there are a lot of AWS services that allow you to build _asynchronous_ flows (note that I'm not talking about services that have batching capabilities like SQS), let's focus on _EventBridge_ and _S3_.

### Testing asynchronous _EventBridge_ flows

#### Receiving the event

Here, you would like to know, that a given event will trigger your customer. For simplicity sake, let's say that consumer is a lambda function. It might sound crazy, but we can listen and match on _CloudWatch logs_ events.

Normally, when you are developing, you should have a logging system in place, which allows you to set sampling based on the environment variable you pass to a lambda function. This approach works ideally for such scenarios. You can set the sampling to 100% for the staging / developer environment and adjust it to a reasonable levels on production.

So add the log line in your handler

```ts
const handler: EventBridgeHandler<any, any, any> = async (event) => {
  // logic

  // use some kind of logger with sampling
  console.log(event);

  // logic
};
```

And then the test itself

```ts
import EventBridge from "aws-sdk/clients/eventbridge";
import { ulid } from "ulid";

const eb = new EventBridge();
const busName = "testingServerlessRefresherasyncFlowbusC977BBFB";

const functionName =
  "testingServerlessRefreshe-asyncFlowreceiver06D8229-WWVZV9IH3RHC";

const testTimeout = 30 * 1000;

test(
  "event is passed to the handler",
  async () => {
    const userId = ulid();

    const detail = {
      id: userId,
      firstName: "Wojciech",
      lastName: "Matuszewski",
    };

    await eb
      .putEvents({
        Entries: [
          {
            EventBusName: busName,
            Detail: JSON.stringify(detail),
            DetailType: "user",
            Source: "async-flow",
          },
        ],
      })
      .promise();

    await expect({
      region: "eu-central-1",
      function: functionName,
      timeout: testTimeout,
    }).toHaveLog(userId);
  },
  testTimeout
);
```

This might seem weird, I totally get you. It might event be something you run on your PR environment only (due to logging sampling and costs of _CloudWatch_), but **please consider how much confidence you gain by having this test**.

1. You know that your pattern is correct
2. You know that correct handler is invoked

This is huge! No more manual testing and "seeing if things work". By introducing manual procedures you are introducing toil. We do not want that.

## Testing batch processing and streaming
