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

// Finished at 30:00
