# Golang Serverless stuff

## Running functions locally

### Serverless Invoke

You can invoke your functions locally without any plugins. This is actually pretty amazing.
See the serverless docs for more info

### SAM plugin

With `SAM` you can even imitate `APIGW`. To do this you will need to transform your `serverless.yml` to `CloudFormation` template. There is a plugin that can do that: `serverless-sam`.

That plugin is kinda limited and works okeish at most, but it's better than nothing.
One think that you might encounter is the `permission denied` error. This is caused by the wrong paths set on your binaries.

```yaml
Properties:
  Handler: ./functions/get-restaurants
  Runtime: go1.x
  CodeUri: >-
    /Users/wn.matuszewski/Desktop/golang/fullstack-serverless/.serverless/big-mouth.zip
```

The `Handler` entry is the culprit here. You should change it to the actual path of the binary:

```yaml
Properties:
  Handler: ./.bin/get-restaurants
  Runtime: go1.x
  CodeUri: >-
    /Users/wn.matuszewski/Desktop/golang/fullstack-serverless/.serverless/big-mouth.zip
```

Now you can simulate your API fully without using `apex/gateway` or similar tools.
