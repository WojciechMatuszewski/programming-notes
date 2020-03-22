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

## Gotchas

### Nested functions with defer

`defer` is used to, well _defer_ the call. There is one gotcha though (probably there are many ;p). **Only the outer function is deferred**. So lets look at some examples:

```go
f, err := os.Open("file.txt")
if err != nil {
  // code
}

defer log.Printf("closed file: %v", f.Close()) // the gotcha is here
```

So the above code wont work as expected. The `f.Close()` will be called immediately and the result would be passed to the `v` formatting arg. `defer` only works directly on `log.Printf` here (in this case).

## Patterns

### Functional options

While working with `go` you might write and API that looks similar to this one:

```go
type Server struct {
  addr string
  timeout time.Duration
  // .. some other config prop
}

func NewServer(addr string, timeout time.Duration, /*... more props */) *Server {
  return &Server{
    addr: addr,
    timeout: timeout,
  }
}
```

We passing of arguments might get out of hand, and what if we want to make some properties optional ?
Well you might try `config structs` or different methods like `NewServerWithTimeout` or `NewServerWithSomething` but this approach can be hard to maintain.

Enter `functional options`.

```go
// exported type mainly used for documentation
type Option func(s *Server)


func Timeout(timeout time.Duration) Option {
  return func(s *Server) {
    s.timeout = timeout
  }
}

type Server struct {
  addr string
  // default: no timeout
  timeout time.Duration
}

func NewServer(addr string, ...opts Option) *Server {
  server := Server{addr: addr}

  for _, opt := range opts {
    opt(&server)
  }

  return &server
}
```

Now, `NewServer` is a `variadic` function which takes multiple config options. We can specify defaults when creating a server then mutate the server with config applied through `Option` type.

## DynamoDB

### `dynamodbav`

It turns out `DynamoDB` has it's own custom `JSON` tag. This tag is used to control how the value will be `unmarshalled`.

One notable tag is `unixtime`, but there is a catch, the underlying type has to be `time.Time`. So

```go
type MyStruct struct {
  createdAt time.Time `dynamodbav:"unixtime"`
}
```

Notice that **I did not specify the `json` tag**. This is explained within `dynamodbattribute/encode.go`. The documentation states that **`dynamodbav` will be favoured whenever possible**. This means that you do not have to specify the `json` tag.

Is this a good idea? You will probably end up specifying both `json` and `dynamodbav` for clarity sake, but I really like the idea. `unixtime` brings a lot of convenience :)

## Language

### `bytes.Buffer` vs `Bufio`

`bytes.Buffer` is just a simple `in-memory` buffer. It will expose `Read` and `Write` methods.
`Bufio` is used for **wrapping** underlying `writers/readers`. This is mainly used for performance. Wrapping with `Bufio` will reduce the amount of calls to `write/read`. This can be useful for files for example, when you DO NOT want every `read/write` call to hit the disc.

```

```
