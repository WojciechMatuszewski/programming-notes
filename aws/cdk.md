# AWS CDK things

## The CDK tree

Whenever you initialize a new _construct_, by convention, you will most likely need to pass the _scope_ as the first parameter to the given _construct_ initialization call.

```ts
const bucket = new s3.Bucket(this, {...})
```

The act of passing down the scope for a given constructs creates a tree-like structure.

This tree like structure is pretty powerful as it allows layering.
These are usually called _construct levels_ also known as L1, L2 and L3 _constructs_.

### Performing operations in the context of the CDK tree

The CDK exposes basic operations that allow you to traverse the tree.
These are tied to the exposed "escape hatches" that allow you to fine-tune the _CloudFormation_ output (more on those later).

#### Searching for a child - traverse tree down

You can use `tryFindChild` or if you feel adventurous (want to catch errors if the child is not found) the `findChild` functions.

Suppose you have a stack that looks as follows

```ts
export class CdkLearningStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const bucket = new s3.Bucket(this, "BucketIdentifier", {});
  }
}
```

To find the L1 child of the `bucket` _construct_ you would write

```ts
bucket.node.tryFindChild("Resource");

// or the implicit version

bucket.node.defaultChild;
```

Very important here, **the `Resource` name is a convention. This is the identifier, the L2 `s3.Bucket` _construct_ decided to give the L1 resource, thus it appears as the `defaultChild`**

The traversal also works for your own custom _constructs_

```ts
export class CdkLearningStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const myConstruct = new MyConstruct(this, "ConstructIdentifier");
  }
}

class MyConstruct extends cdk.Construct {
  constructor(scope: cdk.Construct, id: string) {
    super(scope, id);

    const bucket = new s3.Bucket(this, "BucketIdentifier");
  }
}
```

To get the `bucket` within the `MyConstruct` you would write

```ts
myConstruct.node.tryFindChild("BucketIdentifier");
```

Please note that doing

```ts
myConstruct.node.defaultChild; // returns undefined
```

would return `undefined`. This is because the `MyConstruct` does not contain a child with the id of `Resource`.

#### The `Resource` identifier

As written earlier, the `Resource` identifier is a convention.
If I were to modify my previous example so that the `MyConstruct` has a child with a `Resource` identifier, like so

```ts
class MyConstruct extends cdk.Construct {
  constructor(scope: cdk.Construct, id: string) {
    super(scope, id);

    const bucket = new s3.Bucket(this, "Resource");
  }
}
```

the `myConstruct.node.defaultChild` would yield the `s3.Bucket` _construct_.

Of course, nothing is stopping me from chaining the properties to go even deeper in my traversal

```ts
myConstruct.node.defaultChild.node.defaultChild;
```

The above would yield the L1 bucket _construct_, a child of the `s3.Bucket` _construct_.

#### Searching for a child - traverse tree up

The AWS CDK also allows us to traverse the tree up from a given _construct_ point of view.

This is very useful whenever you want to check if a given resource already exist
so that you do not create it again.

The `cdk.Stack.of(SCOPE)` allows us to get to the parent. From that point, all the capabilities that were mentioned before are available to us.

So let's say I have this kind of situation

```ts
export class CdkLearningStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const singletonLambda = new lambda.Function(this, "SingletonLambda", {});

    const myConstruct = new MyConstruct(this, "ConstructIdentifier");
  }
}

class MyConstruct extends cdk.Construct {
  constructor(scope: cdk.Construct, id: string) {
    super(scope, id);

    const lambda = // ?
  }
}
```

I would like to reference the `SingletonLambda` within the `MyConstruct` without passing it as a property.

With our new found knowledge, I could write

```ts
const lambda = cdk.Stack.of(this).node.tryFindChild("SingletonLambda");
```

Now, this is a very contrived example. In a real-world scenario I would definitely pass
the `SingletonLambda` as a prop to `MyConstruct`.

We do not have to look far for the real use-case scenario through. The L3 `s3-deployment` construct is using this technique of
searching for a given child of it's parent _construct_ to ensure that the lambda that uploads the assets to s3 is only defined once.
The consumer of the `BucketDeployment` _construct_ does not even know that the lambda is defined in the first place!

```ts
export class CdkLearningStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);


    const myConstruct = new MyConstruct(this, "ConstructIdentifier");
  }
}

class MyConstruct extends cdk.Construct {
  constructor(scope: cdk.Construct, id: string) {
    super(scope, id);

    // The `s3Deployment.BucketDeployment` will create a lambda that uploads the assets to s3.
    // This lambda is a singleton, it's only created once in the parent scope.
    const deployment = new s3Deployment.BucketDeployment(this, "bucketDeployment", {..})
  }
}
```

## The escape hatches

Be wary, this section will teach you things you should only be doing as a last resort. Like really.

### Modifying the _CloudFormation_ LogicalID

This could be useful for various reasons.
Maybe yo do not like how your _CloudFormation_ outputs look? maybe you need your _CloudFormation LogicalID_ be more deterministic (I would advice to re-think your design decisions if that's the case).

You can override the _LogicalID_ of a given node by using `overrideLogicalId` function

```ts
const bucket = new s3.Bucket(this, "MyBucket");
const cfnBucket = bucket.node.defaultChild as s3.CfnBucket;
cfnBucket.overrideLogicalId("MyLogicalID");
```

### Modifying the _CloudFormation_ parameters and other properties

This works in a very similar fashion as overriding the _LogicalID_, we just use a different method

To add things to the template you would write

```ts
const bucket = new s3.Bucket(this, "MyBucket");
const cfnBucket = bucket.node.defaultChild as s3.CfnBucket;
cfnBucket.addOverride(
  "Properties.MyImaginaryPropertyKey",
  "MyImaginaryPropertyValue"
);
```

And to remove a given property, all you need to do is to override it with `undefined`

```ts
const bucket = new s3.Bucket(this, "MyBucket", {
  bucketName: "MyBucketName",
});
const cfnBucket = bucket.node.defaultChild as s3.CfnBucket;

// Removing the `BucketName` property from the underlying CloudFormation template
cfnBucket.addOverride("Properties.BucketName", undefined);
```

## Additional resources

- A great example on how to build a custom CDK resource https://www.youtube.com/watch?v=tDXE7S6J_AY
