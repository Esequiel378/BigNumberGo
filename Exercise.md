# Take-Home Exercise

Implement a function that adds any two strings that each represent N arbitrarily large non-
negative numbers.

Please write and document some unit test cases for your addNumbers function. We are
looking for test cases that cover different aspects of your algorithm logic, e.g. cases you
would have in production to cover regressions.

**Requirements:**

  1. The function should be called addNumbers.
  1. The function should accept 2 string params.
  1. Each string param contains M numbers, separated by spaces.
  1. The function must add the numbers in pairs:
      1. Add the 1st number from the first string with the 1st number from the second string.
      1. Add the 2nd number from the first string with the 2nd number from the second string.
      1. etc.
  1. The output should be the sums of the pairs, separated by spaces.
  1. The strings have the same count of numbers.
  1. The numbers may include decimal places.
  1. The numbers can be arbitrarily long, e.g. 1000+ digits.

**Note: This exercise is intended to implement an algorithm, so please refrain from using
built-in high precision types such as Java's BigInt or python's long integer.**

Examples:

```go
>>> addNumbers("123 456 789", "11 22 33")
    "134 478 822"

>>> addNumbers("123456789012345678901 23456789", "12345678 234567890123456789012")
    "123456789012358024579 234567890123480245801"

>>> addNumbers("1234567.8901 2.345", "12.34 2345678901.2")
    "1234580.2301 2345678903.545"
```
