# ProtoBuffers

- think of `json_schema` but better, as in lighter, with stricter types.

- documentation can be embedded within a `schema`.

- you can **generate code** based on the `schema`.

- data cannot be really opened with code editor. You will need some tooling to read it.

## Field tags

The schema itself is pretty simple:

```protobuf
message MyMessage {
  int32 id = 1;
}
```

the `1` is the so called `field tag`. The `tag` is very important.

- use **1 to 15** tags for **for frequently populated fields**. This is because the **`tag` itself weights less**.

- there is a maximum number you can _count_ to. It's quite huge, 2^29 - 1.

- **the tags are used to MATCH FIELDS when serializing and deserializing data**.

## Repeated fields

- these are used to describe **lists**.

```protobuf
message Person {
  repeated string phone_numbers = 1;
}
```

## Constraints

- your code has to place constraints on the fields eg. `int32` cannot be 13 for `month` field.

## Custom types

- you can refer to your custom `messsage` by just using that name as field type.

  ```protobuf
  message Person {
    Date birth = 1;
  }

  message Date {}
  ```

## Nesting types

- you can nest `message` definition within a `message` definition. Allows you to create nested structures.

```protobuf
message Address {
    string address_line_1 = 1;
    string address_line_2 = 2;
    string zip_code = 3;
    string city = 4;
    string country = 5;
  }

  repeated Address addresses = 1;
  }
```

## Imports

- you can import `message`s from other files. Use `import "path"` syntax.

## Packages

- `package NAME` just creates a namespace. Important while importing stuff.
