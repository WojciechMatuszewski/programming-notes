# Testing Serverless

## Introduction

Testing. One of my favorite things to do. The more software I write, the more I love writing tests. Features are the things that put food on our tables, but the tests are the things that keep everything in check and in order.

Many of you are probably familiar with the notion of _unit tests_. These are small, **fast**, tests that check if a given _unit_, be it a function or a _Component_, works correctly.

When it comes to serverless, we still have _unit tests_, but they play a lesser role in the overall testing story. With us using managed services more and more, it's the _integration_ and _end to end_ tests that matter the most. 100% of coverage on your lambda handler is almost worth nothing if it cannot get the data from _DynamoDB_ because of missing IAM permissions.

This file contains all the knowledge I have about testing serverless apps. We will start at the most basic level. I'm talking a simple API with _API Gateway_ and _DynamoDB_, and end at _SQS_ and other batch processing services.

## Testing simple APIS

This is where everything makes sense. We have our simple architecture with _API Gateway_ fronting one lambda function (which could be skipped) that talks to _DynamoDB_

### Unit tests

**This section assumes that you have nothing (PR scope) deployed to the cloud**.

Even though we are writing _unit tests_ here, I strongly advise you **against** mocking anything\*\*. If you want to test the _integration_ with _DynamoDB_, you can do that in the _integration test_.

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
import invokeEvent from "../../events/simple-api-event.json";
import { createHandler } from "../../functions/simple-api/handler";

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

#### Sending the event

Let's say we have an _API Gateway_ fronting a simple lambda function which will send events to the _EventBridge_. We want to test that the handler is sending that event correctly. In such case, we will introduce _test only_ resource, mainly _SQS queue_.
The queue will be used as a target of the _EventBridge_. Then we could assert that after I invoked the API, the event is in the queue.

The test would look something like this

```ts
import phin from "phin";

const apiUrl = "https://ex59t535b2.execute-api.eu-central-1.amazonaws.com/";
const queueUrl =
  "https://sqs.eu-central-1.amazonaws.com/286420114124/testingServerlessRefresher-asyncFlowqueueB14A1593-14PFRLCUOZCTA";

const testTimeout = 30 * 1000;

test(
  "event is send",
  async () => {
    const result = await phin({
      method: "POST",
      parse: "json",
      url: apiUrl,
      data: {
        firstName: "Karol",
        lastName: "Matuszewski",
      } as any,
    });

    expect(result.statusCode).toEqual(200);

    expect({
      region: "eu-central-1",
      queueUrl,
      timeout: testTimeout,
    }).toHaveMessage(
      (msg) =>
        msg.detail.firstName === "Karol" &&
        msg.detail.lastName === "Matuszewski"
    );
  },
  testTimeout
);
```

It's pretty self explanatory right? Again, the confidence we gain here is huge

1. Is my handler sending the correct event?
2. Is the execution role of my handler correctly set up?

One thing I want to stress here. The **_SQS_ queue is only for testing purposes**. You would not want to deploy this resource to your production environment. Due to this, I've set up a relatively short _visibility timeout_ on the queue itself - 300 second.

### Testing `S3` flows

We could have couple of scenarios here. For simplicity sake, let's say I want to get a _presigned url_ back, which allows me to upload the image. After the upload is complete, the image should be in a bucket under a given key.

```ts
test(
  "creates a presigned url which enables the user to upload the image",
  async () => {
    const result = await phin<{ url: string; fields: Record<string, string> }>({
      method: "GET",
      parse: "json",
      url: `${apiUrl}/get-upload-url?name=image.png`,
    });

    expect(result.statusCode).toEqual(200);
    expect(result.body).toEqual({
      url: expect.any(String),
      fields: expect.any(Object),
    });

    const { url, fields } = result.body;

    const fileToUpload = fs.readFileSync(join(__dirname, "image.png"));

    const formData = new FormData();
    Object.entries(fields).forEach(([key, value]) =>
      formData.append(key, value)
    );
    formData.append("file", fileToUpload);

    const uploadResult = await phin({
      url: url,
      method: "POST",
      data: formData.getBuffer(),
      headers: formData.getHeaders(),
    });

    expect(uploadResult.statusCode).toEqual(204);

    expect({
      region: "eu-central-1",
      bucket: bucketName,
      timeout: testTimeout,
    }).toHaveObject(fields.key);
  },
  testTimeout
);
```

One thing to notice here is that the `/get-upload-url` path returns **presigned POST data**. There are numerous benefits to using _presigned POST_ instead of _presigned GET/PUT_ urls with S3, mainly the _Conditions API_.

Just like every _E2E_ test, the confidence gain from this test is insane

1. Is my handler configured correctly when it comes to IAM?
2. Are the conditions specified in the handler correct?
3. Can the user upload images to S3 using the data returned in the response?

### Testing Step Functions flows

I'm aware of three ways to test _Step Functions_ flows.

1. Use the [local mocking capabilities](https://aws.amazon.com/blogs/compute/mocking-service-integrations-with-aws-step-functions-local/).
1. Decompose a given task into multiple parameters that you could import to your IaC code and into your test code. Graham Allan talks about this technique [here].
1. Test the whole thing end-to-end and assert on the [_Step Function_ execution history](https://grundlefleck.github.io/2022/01/12/how-using-the-same-programming-language-for-iac-made-a-step-function-testable.html).

I would argue that you can use each of those methods in some scenarios. It all depends on the needs we have. What follows is me writing about each testing technique in more detail while giving guidance on using it.

TODO: https://github.com/WojciechMatuszewski/testing-aws-step-functions

#### Local mocking capabilities

I have a somewhat strong opinion regarding mocking AWS services – I do not think that is a good idea (albeit it might be in the context of unit tests).

So what would be, **in my opinion**, a good use-case for using the local mocking capabilities (_aws-stepfunctions-local_ Docker image)? In my mind, such an environment is **great for testing the steps that transform the data** – the `Pass` states.

How many times have you got stuck on some kind of transformation? Maybe the [_intrinsic function_](https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-intrinsic-functions.html) syntax was wrong? Or maybe the `resultPath` or other parameters?

##### An example

I want to test a `Pass` state that concatenates two strings together using the `States.Format` function. The following is an _AWS CDK_ definition of the infrastructure we will be working with.

```ts
export class LocalTesting extends Construct {
  public transformIncomingDataStep: aws_stepfunctions.Pass;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.transformIncomingDataStep = new aws_stepfunctions.Pass(
      this,
      "TransformIncomingDataStep",
      {
        parameters: {
          payload: aws_stepfunctions.JsonPath.stringAt(
            "States.Format('{} {}', $.firstName, $.lastName)"
          ),
        },
      }
    );

    const machine = new aws_stepfunctions.StateMachine(this, "StateMachine", {
      definition: this.transformIncomingDataStep,
    });
  }
}
```

I made the `transformIncomingDataStep` public on purpose. As I eluded earlier, I'm only concerned with steps that transform data and do not directly interact with other AWS services.

Writing the test for the `transformIncomingDataStep` will be a bit involved. We have to:

1. Spin up the Docker container with the Step Functions local image.
1. Create a Step Function that only contains the `transformIncomingDataStep` ASL.
1. Run the Step Function.
1. Assert on the result.

Starting with the first point. I like to use the `testcontainers` package for Docker image orchestration in tests.

```ts
let container: StartedTestContainer | undefined;

beforeAll(async () => {
  container = await new GenericContainer("amazon/aws-stepfunctions-local")
    .withExposedPorts(8083)
    .start();
}, 15_000);

afterAll(async () => {
  await container?.stop();
}, 15_000);
```

And here is the test itself.

```ts
test("The `TransformIncomingDataStep` works as expected", async () => {
  const testIdentifier = Buffer.from(
    `${__filename}_${expect.getState().currentTestName}`
  )
    .toString("hex")
    .replace(/\d/g, "");

  const stack = new cdk.Stack();
  const construct = new LocalTesting(stack, testIdentifier);

  const stepFunctionsASL = construct.transformIncomingDataStep.toStateJson();

  const sfnClient = new SFNClient({
    endpoint: `http://localhost:${container?.getMappedPort(8083)}`,
  });

  const createSFNResult = await sfnClient.send(
    new CreateStateMachineCommand({
      definition: JSON.stringify({
        StartAt: "TestingStep",
        States: {
          TestingStep: {
            ...stepFunctionsASL,
            End: true,
          },
        },
      }),
      name: testIdentifier,
      roleArn: "arn:aws:iam::012345678901:role/DummyRole",
    })
  );

  const machineARN = createSFNResult.stateMachineArn;

  const sendPayloadResult = await sfnClient.send(
    new StartExecutionCommand({
      stateMachineArn: machineARN,
      input: JSON.stringify({
        firstName: "Wojtek",
        lastName: "Matuszewski",
      }),
      name: testIdentifier,
    })
  );

  await waitFor(async () => {
    const getExecutionHistoryResult = await sfnClient.send(
      new GetExecutionHistoryCommand({
        executionArn: sendPayloadResult.executionArn,
      })
    );

    const exitedState = getExecutionHistoryResult.events?.find(
      (event) => event.type == "PassStateExited"
    );

    expect(exitedState).toBeDefined();
    expect(exitedState?.stateExitedEventDetails?.output).toEqual(
      JSON.stringify({ payload: "Wojtek Matuszewski" })
    );
  });
}, 20_000);
```

To somewhat guarantee the uniqueness of different identifiers used within the test (Step Function `name` and the Execution `name`), I created the `testIdentifier` variable.

While the test setup might feel bloated, remember that you only have to code for it once. In the long run, it will save you time as deploying the changes, running the state machine, looking at the AWS console is a time-consuming process.

#### Decomposing the `Task` state

When it comes to decomposing the `Task` states, the technique is similar to what we have done in the [Local mocking capabilities](#local-mocking-capabilities) section.

Instead of the `Pass` state, we make the service API call parameters publicly accessible. [Graham Allan](https://twitter.com/Grundlefleck) already wrote an [excellent article](https://grundlefleck.github.io/2022/01/12/how-using-the-same-programming-language-for-iac-made-a-step-function-testable.html) on this topic. I recommend you check out and try the approach he is proposing!

#### End-to-end testing

Depending on the complexity of your Step Function, the end-to-end testing can either give you a lot of confidence in the code you are shipping or be a slow process that tests only a particular branch of the Step Functions logic.

Speaking from experience, the more complex the unit of deployment is (in our case, a Step Function), the harder it is to test all possible scenarios and branches of the logic embedded within that unit.

To test Step Functions flows end to end, we have to:

- Assume a role that allows us to invoke and probe the given Step Function. This might be the role you use for deploying our development stack.
- Start the Step Functions execution.
- Assert on the execution history.

The difficulty arises whenever a given Task depends on a response from some other service that we do not control. Suppose we cannot force a given response from that particular service. In that case, we might find ourselves in a situation where testing the given Step Function branch is impossible!

## Testing batch processing and streaming

There is a lot to think about when it comes to batch processing and streams. I think the main gotchas are idempotency and ensuring that one _poison pill_ is will not cause our service to halt completely.

We can tackle the second concern in our unit tests.

### Batch processing with SQS

#### Unit tests

Let's say you are pulling from _SQS_ queue. You are knowledgeable in AWS so you know that the _lambda service_ will automatically delete the messages that you process. But you might not be aware is that **whenever you throw an error from your lambda function, THE WHOLE BATCH will be retried**. This is usually not a problem when the batch size is 1, but that is often not the case.

What you have to do in such situation is to delete messages manually, while processing the batch in parallel. The logic for that would look something like this

```ts
import { SQSHandler } from "aws-lambda";
import SQS from "aws-sdk/clients/sqs";

const sqs = new SQS();

async function performWork() {}

const handler: SQSHandler = async (event) => {
  const promises = event.Records.map(async (record) => {
    await performWork();
    await sqs
      .deleteMessage({
        ReceiptHandle: record.receiptHandle,
        QueueUrl: "QUEUE_URL",
      })
      .promise();
  });

  const results = await Promise.allSettled(promises);
  const hasErrors = results.find((result) => result.status === "rejected");

  if (hasErrors) {
    throw new Error("Errors occurred");
  }
};

export { handler };
```

Notice that I'm throwing an error if there were some errors. That is completely OK, because I'm manually deleting messages. This way, the next event will only contain messages that were problematic. Combine this approach with bisecting on error and you have something that is resilient to failures.

So to the test itself. Previously I was advocating for extracting the logic to a separate function which could be exported. Here I will be using _dependency injection_ for the _SQS_ service and the _performWork_ function.

```ts
import { SQSHandler } from "aws-lambda";
import SQS from "aws-sdk/clients/sqs";

const sqs = new SQS();

async function performWork() {}

function isRejected(result: PromiseSettledResult<void>) {
  return result.status === "rejected";
}

function createHandler(worker: () => void, sqsService: SQS) {
  const handler: SQSHandler = async (event) => {
    // code
  };

  return handler;
}

const handler = createHandler(performWork, sqs);

export { handler };
```

This way I can easily pass the dependencies in the test itself. The test is rather straightforward, I'm not going to be pasting it here.

#### E2E tests

Let's face it, while working with any kind of batching oriented workloads, you are going to have to deal with _Dead letter queues_. I would go as far as to say that _DLQs_ are must haves.

How we can ensure that we are adhering to best practices (deleting messages manually and actually having _DLQs_ set up)? We can write tests for that!

These tests will be really slow. I'm talking more than a minute. This is due to native retry mechanism you get when your lambda is hooked to _SQS_.

So let's imagine a scenario where we want to send emails through _SES_. Our lambda is fronted by _SQS_ so that we can perform work in batches. We are going to write a tests which ensures that

1. Only the poisoned message lands in _DQL_ (the other one is deleted)
2. We are actually have _DQL_ setup up
3. Our lambda function has permissions to send an email through _SES_ service

The test for this would look something like this

```ts
test(
  "bad requests land in dlq",
  async () => {
    const badMessageId = ulid();

    // Can be executed once every 60 seconds!
    await sqs.purgeQueue({ QueueUrl: dqlUrl }).promise();

    await sqs
      .sendMessageBatch({
        Entries: [
          {
            Id: ulid(),
            MessageBody: JSON.stringify({
              // This is an email exposed by SES that will always succeed
              destination: "success@simulator.amazonses.com",
              source: "MY_OWN_EMAIL_CONFIRMED_IN_SES",
            }),
          },
          {
            Id: ulid(),
            MessageBody: JSON.stringify({
              destination: "success@simulator.amazonses.com",
              source: "dupa@dupa.pl",
              id: badMessageId,
            }),
          },
        ],
        QueueUrl: queueUrl,
      })
      .promise();

    await expect({
      region: "eu-central-1",
      queueUrl: dqlUrl,
      timeout: testTimeout,
      poolEvery: 2000,
    }).toHaveMessage((msg) => {
      return msg.id === badMessageId;
    });

    await new Promise((resolve) => setTimeout(resolve, 6000));

    const { Attributes } = await sqs
      .getQueueAttributes({ QueueUrl: dqlUrl, AttributeNames: ["All"] })
      .promise();

    expect(Attributes!.ApproximateNumberOfMessages).toEqual(1);
    expect(Attributes!.ApproximateNumberOfMessagesNotVisible).toEqual(0);
    // Depending on how your test is set up, you might need to remove the message you are making the assertions on.
    // The message will not be automatically removed since we are pooling manually and with help of EventSourceMapping.
  },
  testTimeout
);
```

Again, so much confidence! This test is not without it's drawbacks though. First of all it takes a lot of time to run, but I think the most important drawback is that due to us calling `purgeQueue` we can run it once per 60 seconds.

#### Alternative using CDK provider framework

All the previous tests that I've shown you were running locally. One might argue that is not the best from the security perspective. We are using our local credentials after all.

We can go an alternative route and have the assertions, be it not so powerful, run in a lambda instead!
Enter the [_AWS CDK provider framework_](https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html#aws-cdk-custom-resources).

Usually, we think about the _provider framework_ in the context of creating resources that are not supported by _CloudFormation_.
But there is a hidden gem amongst the _provider framework_ API that will allow us to run the test assertions every X seconds – the [_asynchronous handler_](https://docs.aws.amazon.com/cdk/api/latest/docs/custom-resources-readme.html#aws-cdk-custom-resources).

##### Benefits

The main benefit of having your tests live very close to the infrastructure they test is that **your stack will not deploy if your test fail. Running the tests is NOT a separate command**.

The setup I'm referring to is not complex. Here is the pseudo-code.

```ts
const appStack = new AppStack();

const testStack = new TestStack(appStack);

// Explicit over implicit
testStack.addDependency(appStack);
```

##### Drawbacks

The main drawback of this approach is that **your deployments will take longer**. Please note that I'm not referring your CI/CD pipeline.
Since the tests are part of the infrastructure, your deploy step will take additional time.

Another issue one might face is the issue with messed up _CloudFormation_ stacks. I personally never encountered this situation but I imagine it would be possible to end up in situation where _CloudFormation_ rollback is not possible?

##### How

[This repository](https://github.com/elthrasher/cdk-async-testing-example) contains all you need.

### Stream with _DynamoDB_ streams

This is basically the same strategy as for checking if something is triggered by _EventBridge_ or other similar services.

Let's say that I'm going to be pushing to _EventBridge_ from _DynamoDB_ stream. One neat thing I like to do is to make a _CloudWatch_ log group a target of _EventBridge_ rule. This way I do not have to log anything inside the stream handler.

The test would look something like this:

```ts
test(
  "events are sent do event bridge",
  async () => {
    const itemId = ulid();
    await docClient
      .put({ Item: { id: itemId }, TableName: tableName })
      .promise();

    const searchLog = async () => {
      const result = await logs
        .filterLogEvents({
          logGroupName,
          filterPattern: itemId,
        })
        .promise();

      if (!result.events) return [];

      return result.events.map((event) => JSON.parse(event.message!).detail.id);
    };

    await waitForExpect(
      async () => {
        const events = await searchLog();
        return expect(events).toContain(itemId);
      },
      testTimeout,
      2000
    );
  },
  testTimeout
);
```

Sadly, the `aws-testing-library` does not have a matcher for a specific log group, so I had to write my code for that.
Otherwise, it's almost the same as previous tests.

## Using the CDK provider framework

You can incorporate your E2E tests directly into the infrastructure and let CloudFormation handle the case when the tests fail – roll back the changes.
As good as it might sound, the approach is not without drawbacks. Instead of having the test logs visible in your CI environment, we will have to use the CloudWatch console to view them – it might be a significant pain point since the CloudWatch console is not exactly lovely to work with.

If you want to learn more about utilizing the provider framework for testing, consider [giving this talk a watch](https://youtu.be/s8tO-ymVQPg?t=8515) then, consult [this repository](https://github.com/WojciechMatuszewski/cdk-async-testing-example).
