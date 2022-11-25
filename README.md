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

_**The `AddNumbers` function can is located in the [bignumber.go](https://github.com/Esequiel378/teladoc-take-home-exercise/blob/d2d50a75f41d21b6458f5031774f71b5e6ed0cfe/pkg/bignumber/bignumber.go#L36) file, within the `pkg` directory**_

#### Solution 1: Naive approach

> Located in the [naive-solution](https://github.com/Esequiel378/teladoc-take-home-exercise/tree/solutions/naive-solution) branch

This approach will solve the problem by doing a basic right-to-left sum of each individual digit, handling the
decimal part of the number as a separated sum, add the carry from the last addition the integer part
original number

This approach will solve the addition of N + M by doing a on digit at a time sum from right-to-left, handling the
fractional part of N and M as a separated addition and adding the carry value to the integer addition of N and M

This approach will solve the problem of N + M by doing and aritmetic addition of each digit from right-to-left carring
the remainding of the sum to the next addition

**Integer addition `489 + 14`:**

<pre>
<strong><span>&#8226;</span></strong>  9  + 4     = 13   carry = 1   result = 3
<strong><span>&#8226;</span></strong> (1) + 8 + 1 = 10   carry = 1   result = 0
<strong><span>&#8226;</span></strong> (1) + 4 + 0 = 5    carry = 0   result = 5
</pre>

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

<pre>
<strong><span>&#8226;</span></strong>  9  + 3     = 12   carry = 1   result = 2
<strong><span>&#8226;</span></strong> (1) + 9 + 0 = 10   carry = 1   result = 0
<strong><span>&#8226;</span></strong> (1) + 1 + 1 = 3    carry = 0   result = 3
</pre>

```go
19.9
10.3
---
30.2
```

#### Solution 2: 32 bits chunks

> Located in the [master](https://github.com/Esequiel378/teladoc-take-home-exercise/tree/master) branch

This approach will solve the addition of N + M by doing a right-to-left aritmetic addition on
chunks of 32 bits numbers. To performe this chunk addition we need to pre-process the input string

1. Split the string in two (integer and decimal)
1. Split the string in chunks of 9 digits (~32 bits) <sup>**1**</sup> numbers
1. Each chunk of strings are parsed and stored as `uint32` numbers
1. The aritmetic addition is perform like in the [Solution 1: Naive approach](#solution-1-naive-approach)
   but using the 32 bits numbers instead of one digit at a time, which is more performant

> **1**: A 32 bits integer number (unsigned) has 10 digits and max number of `4294967295`.
> The reason I decided to use 9 digits, is to be able to perform the addition of two 32 bits
> numbers like this `999999999 + 999999999 = 1999999998` witout overflowing the 32 bits range
> `1999999998 < 4294967295`.
