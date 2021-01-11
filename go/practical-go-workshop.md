# Practical GO workshop

## Identifiers

- **optimize for the reader**

* a good name for a variable does not describe the variable contents but the purpose

- you should consider short variable names for variables that live short. Find balance and not strive blindly to use 1 letter variables.

* do not mix and match long and short variable names

```go
func (s *SNMP) Fetch(oid []int, index int) (int, error)

    // vs (clearly suffers when it comes to readability \/)
    func (s *SNMP) Fetch(o []int, i int) (int, error)
```

- do not let package names steal good variable names.
  a good example of this problem is `context.Context`. This is why you will most use `ctx` here.

* use **names who have established name through programming languages**.
  these include `i`, `j`, `db` all of the usual stuff.

## Initializers

1. **when declaring, but not initializing `use` var`**

```go
var players int

var things []Thing

var thing Thing
```

2. **when `declaring and initializing` use `:=` syntax**

```go
players  := 0
```

If you initialize with `var` syntax you are performing work that is **redundant!**.

```go
var players int = 0
// is the same as
var players = 0

// thus you could just write it as
players := 0
```

Of course, do not follow this rule blindly. There are exceptions, mostly where you have 2 related variables.

```go
// weird right?
var min int
max := 1000

// much better
min, max := 0, 1000
```

- **make tricky things stand out**.
  This is also a case where you might consider breaking this rule (but think, deeply, before you do that).

## Comments

- describe **behavior** not **contents**.
  A good example of this are http `status codes`.

  ```go
  const (
    StatusOk = 200 // RFC ..
  )
  ```

* always document public symbols, this means putting meaningful comments above the things you export.

- avoid comments that do not add any value. You will most likely fall into that trap when using `interfaces`

  ```go
  // Read implements the io.Reader interface
  func (r *FileReader) Read(buf []byte) (int, error)
  ```

  That comment literally does not bring any value. You might as well delete it.

## Package design

- try not to leak implementation details via `types` or other means.

- do **not** use **`utility` packages**. You will regret it, really. Some duplication will not kill you.

- make zero value useful. A good example of this would be `sync.Mutex`

  ```go
  type MyInt struct {
    mu sync.Mutex
    val int
  }

  func main() {
    var i MyInt

    // usable without initialization
    i.mu.Lock()
    i.val++
    i.mu.Unlock()
  }
  ```

  or your good friend `bytes.Buffer`

  ```go
  func main() {
    var b bytes.Buffer
    b.WriteString("Hello")
    io.Copy(os.Stdout, &b)
  }
  ```

- try to **avoid pkg directory**. This is something that the go library moved away from long time ago.
  Using the _pkg_ directory is like putting _I_ before interfaces - needles boilerplate.

## API Design

- **avoid** functions with **multiple arguments of the same type**.
  Imagine having a function:

  ```go
  func copy(to, from string) error
  ```

  Would be hard to distingush between

  ```go
  copy("/tmp/backup", "presentation")
  copy("presentation", "/tmp/backup")
  ```

  Instead, **use aliases, and receivers**

  ```go
  type Source string
  func (src Source) CopyTo(dest string) error{}

  func main() {
    var from Source = "presentantion.md"
    from.CopyTo("...")
  }
  ```

  Much better!

https://youtu.be/gi7t6Pl9rxE?list=WL&t=7704
