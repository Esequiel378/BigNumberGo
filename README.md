# Teladoc - Take-Home Exercise

> You can find the problem description in the [Exercise.md](./Exercise.md) file

### 👷 CI - Continuous Integration

The project uses [github actions](https://github.com/features/actions) to run jobs for build, linting, and testing
you can inspect the results in the [actions](https://github.com/Esequiel378/teladoc-take-home-exercise/actions) tab

### ✅ Run tests

> To run test manually, run the next command

```shell
go test ./...
```

### 📈 Run tests coverage

> To inspect the test coverage, run the next command

```shell
go test -coverpkg=./... ./...
```

#### Solution 1: Naive approach

This approach will solve the problem by doing a basic right-to-left sum of each individual digit, handling the
decimal part of the number as a separated sum, add the carry from the last addition the the integer part
original number

This approach will solve the addition of N + M by doing a on digit at a time sum from right-to-left, handling the
fractional part of N and M as a separated addition and adding the carry value to the integer addition of N and M

This approach will solve the problem of N + M by doing and aritmetic addition of each digit from right-to-left carring
the remainding of the sum to the next addition

**Integer addition `489 + 14`:**

1. **Add** `9 + 4 = 13`         carry = `1`   result = `3`
1. **Add** `(1) + 8 + 1 = 10`   carry = `1`   result = `0`
1. **Add** `(1) + 4 + 0 = 5`    carry = `0`   result = `5`

```go
489
 14
---
503
```

**Fractional addition `19.9 + 10.3`:**

1. Split both numbers by the period and do the same process for the integer addition starting with the fractional part
1. The fractional addition will return a carry value, that should be the initial carry value for the integer addition
1. Join the integer and the decimal part with a dot (`.`)

1. **Add** `9` + 3 = `12`             carry = `1`   result = `2`
1. **Add** `(1)` + `9` + `0` = `10`   carry = `1`   result = `0`
1. **Add** `(1)` + `1` + `1` = `3`    carry = `0`   result = `3`

```go
19.9
10.3
---
30.2
```
