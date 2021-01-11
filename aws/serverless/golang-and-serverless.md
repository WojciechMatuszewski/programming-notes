# Golang Serverless stuff

## Structure within `serverless.yml`

Please keep in mind that you can extract pieces of you `serverless.yml` into different files using `{file()}` syntax.

This can be done **not only on the `resources` level** but also eg. for **`functions` block level**.

```yml
functions:
  - ${file(./functions/randomGenerator/serverless.yml)}
  - ${file(./functions/uuidGenerator/serverless.yml)}
```

### Referencing stuff from other files

The `{file}` syntax is quite powerful. You can even reference keys which are defined in given file

```yml
vpc: ...
# in another file
vpc: ${file(file):vpc}
```

## APIGW

### Mapping templates

The dreaded mapping templates, **who the fuck likes VTL?!**.

It can be useful though. One scenario I've encounter is to assign a static property to request body. This can be done using `mapping templates` at `apigw` level, no need for code.

```vtl
{
  #set($inputRoot = $util.parseJson($input.json('$')))
  "body" : {
    #foreach($entry in $inputRoot.entrySet())
    "$entry.key": "$entry.value"
    #end
    "foo":"bar"
  }
}
```

Pretty straight forward right? As much as I hate using `vtl`, it's not that bad. The `foo:bar` key-pair is the static one, will always be there, no matter the body.

You can take this a step forward and only include it IF there there are no such property within the body already.

```vtl
"body" : {
    #foreach($entry in $inputRoot.entrySet())
    "$entry.key": "$entry.value"
    #end
    #if(!$inputRoot.containsKey("foo"))
    "foo": "bar"
    #end
  }
```

This can help you when you have differences between given APIs (running 2 lambdas versions).

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

### Stale resource policy on APIGW

You might have a problem where you defined a resource policy on `API Gateway` and it seem to be not working.

There are a couple of steps you can take to make sure that policy is really applied:

- first thing first, **make sure the newest version of your api is deployed**. Sadly `APIGW REST API` does not have `auto-deploy` feature (unlike `HTTP API`).

* this **might be a cashing issue**. To make sure **check using incognito**.

- the resource propagation takes a while.
  > It does take a 30-60 seconds for the change to be reflected after deploying the API. Make sure you’re using incognito mode or curl to avoid any browser caching.

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

## X-Ray madness

This is **pure BS**. Only after my second try getting X-Ray to work I finally got to the bottom why it did not work for the first time.

See, there is `github.com/aws/aws-lambda-go` library. Then you use `go get` you will , at the time of writing this, get `v1.15.0` version. Now lets switch to `X-ray` sdk : `0.9x` when you use `go get`.

Now, this would be completely ok BUT **these 2 sdks are tied together!**. **You have to have specific `X-ray` sdk version to make it work with the lambda sdk**. And here comes the BS part:

- then you update the `X-Ray` sdk to `v1.0.0-rc.1` the call will not panic, but guess what? **Your subsegments will not be there**. Frankly I think there last number of sdk versions has to be in-sync.

- **only, and only when you have the same number at the end they will work!**. So the **lambda `sdk` has to have `v1.15.0` version and the `X-Ray sdk` has to have `v1.0.0-rc.15` version**. Well that would be ok IF it was specified within the documentation. But NOPE!.

**GO FIGURE LOL!**

## Language

### `bytes.Buffer` vs `Bufio`

`bytes.Buffer` is just a simple `in-memory` buffer. It will expose `Read` and `Write` methods.
`Bufio` is used for **wrapping** underlying `writers/readers`. This is mainly used for performance. Wrapping with `Bufio` will reduce the amount of calls to `write/read`. This can be useful for files for example, when you DO NOT want every `read/write` call to hit the disc.

### JSON patterns

#### `DisallowUnknownFields`

While working with `json` data, it may happen that you do not want to handlee any unknown (properties not declared within your `struct`) key:value pairs and return an error. This is pretty simple to do using `DisallowUnknownFields` and `json.Decoder`.

```go
type Student struct {
  Name string `json:"name,omitempty"`
}

var jsonText = `{"name":"Wojtek", "foo":"bar"}`

func main() {
  dec := json.NewDecoder(bytes.NewReader([]byte(jsonText)))
  dec.DisallowUnknownFields()

  var s Student
  err := dec.Decode(&s);
  // this will result in an error
  if err != nil {
    fmt.Println(err.Error())
  }
}
```

#### Default values for JSON key value pairs

There are 2 approaches here. How you define your underlying `struct` will dictate which is the way to go.

1. Overriding default values when using `json.Unmarshal`

This one is the most straightforward way.

```go
type Student struct {
  Name string `json:"name,omitempty"`
  // I want `grade` to be 3 by default
  Grade int `json:"grade,omitempty"`
}

var jsonText = `{"name":"Wojtek"}`

func (s *Student) UnmarshalJSON(b []byte) error {
  type student Student
  st := student{
    Grade: 3
  }

  err := json.Unmarshal(b, &st)
  if err != nil {
    return err
  }

  *s = Student(st);
  return nil
}
```

Pretty straightforward right?

2. Using pointers when declaring `struct` fields to be able to perform `nil` check.

You will often see this using AWS golang package. There, the underlying `struct` fields are mostly defined as pointers.

```go
type Student struct {
  Name string `json:"name,omitempty"`
  Grade *int `json:"grade,omitempty"`
}

// when umarshaling

if st.Grade == nil {
  var g int = 3;
  st.Grade = &g
}
```

One concert here would be performance, but always remember to measure before you employ any kind of optimizations.

### Custom `marshal` for dealing with APIs

While working with APIs you might need to provide different representations of values what your `marhal/unmarshal` function is providing.
This is the ideal case for implementing those methods on structures you are maintaining within the codebase.

#### `time.Time` example

I think a good example of this would be `time.Time` type. Most APIs expect you to pass `unix` time and not `RFC` whatever.
Problem is that when you `marshal` `time.Time` you get `RFC-xxx` representation of that value.
Let's create our own `time` type which will override that behavior.

```go
package main
import (

"encoding/json"
"fmt"
"time"
)

type Time struct{
    time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
    return json.Marshal(t.Unix())
}

func (t *Time) UnmarshalJSON(data []byte) error {
    var i int64
    err := json.Unmarshal(data, &i)
    if err != nil {
        return err
   }

    t.Time = time.Unix(i, 0)
    return nil
}

func main() {
    t := Time{time.Now()}
    b, _ := json.Marshal(t)

    // timestamp
    fmt.Println(string(b))

    var t2 Time
    json.Unmarshal(b, &t2)

    // RFC-xxx time
    fmt.Println(t2)
}
```

Notice the **type embedding**. With this technique we even get the other time-related methods on our type (inherited from `time.Time`)

#### Generic structures

What happens when your API returns 2 different things for the same endpoint? This is quite common for webhooks.
I personally had such case but at that time I did not know much about golang to begin with so I could not deal with it.

So, let's imagine that your API is returning 2 different JSON responses.

```json
{
  "data": {
    "object": "bank_account",
    "id": "ba_123",
    "routing_number": "110000000"
  }
}
```

And this one

```json
{
  "data": {
    "object": "card",
    "id": "card_123",
    "last4": "4242"
  }
}
```

These structures represent completely different things. How do we deal with this?

Start with defining `struct` for each of these structures

```go
package main
type BankAccount struct {
    Object string `json:"object"`
    ID string `json:"id"`
    RoutingNumber string `json:"routing_number"`
}

type Card struct {
    Object string `json:"card"`
    ID string `json:"id"`
    Last4 string `json:"last4"`
}
```

Now we are going to embed these into 1 data structure. This data structure will be used for `marshalling` and `unmarshalling`.

```go
package main
type Data struct{
	*BankAccount
    *Card
}
```

Notice the pointers. These will be used for checking with which structure we are dealing with (eg. `Data.BankAccount == nil ?`)
Now for the parsing methods.

```go
package main
import (

"encoding/json"
"errors"
)
type Data struct {
// previous stuff
}

func (d *Data) UnmarshalJSON(data []byte) error {
	temp := struct {
		Object string `json:"object"`
	}{}

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	if temp.Object == "card" {
		var c Card
		err = json.Unmarshal(data, &c)
		if err != nil {
			return err
		}

		d.Card = &c
		d.BankAccount = nil
		return nil
	}

	if temp.Object == "bank_account" {
		var b BankAccount
		err = json.Unmarshal(data, &b)
		if err != nil {
			return err
		}

		d.Card = nil
		d.BankAccount = &b
		return nil
	}

	return errors.New("unknown type")
}
```

Notice the `temp` struct. This one is there to determine which underlying struct should be used for the `unmarshal` method.
With only 2 cases this is pretty straight forward. My hope is that with generics such cases will be much easier to handle :).
