package utils

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInputWithDifferentNumbersCount = errors.New("input with different numbers count")

// AddNumbers takse two string params containing M numbers
// separated by spaces and returns sum of the pairs.
//
// + The input parameters should have the same count of numbers.
// + The numbers may include decimal places.
// + The numbers can be arbitrarily long, e.g. 1000+ digits.
//
// Examples:
//  >> AddNumbers("123 456 789", "11 22 33")
//  "134 478 822"
//
//  >> AddNumbers("123456789012345678901 23456789", "12345678 234567890123456789012")
//  "123456789012358024579 234567890123480245801"
//
//  >> AddNumbers("1234567.8901 2.345", "12.34 2345678901.2")
//  "1234582.2301 2345678903.545"
func AddNumbers(lhs, rhs string) (zero string, err error) {
	lhsElements, rhsElements := strings.Split(lhs, " "), strings.Split(rhs, " ")

	if len(lhsElements) != len(rhsElements) {
		return zero, ErrInputWithDifferentNumbersCount
	}

	result := make([]string, len(lhsElements))

	for idx := 0; idx < len(lhsElements); idx++ {
		// TODO: chekIsInBounds
		lhsElm, rhsElm := lhsElements[idx], rhsElements[idx]

		var r string
		r, err = addStringNumbers(lhsElm, rhsElm)
		if err != nil {
			return zero, err
		}

		result[idx] = r
	}

	resultString := strings.Join(result, " ")

	return resultString, nil
}

func addStringNumbers(lhs, rhs string) (zero string, err error) {
	// TODO: Implement decimal addition
	if strings.Contains(lhs, ".") || strings.Contains(rhs, ".") {
		return "", errors.New("not implemented")
	}

	// Make sure that lhs is the largest number
	if len(lhs) < len(rhs) {
		lhs, rhs = rhs, lhs
	}

	result := make([]string, len(lhs))
	carry := 0

	// Iterate from the end of the string
	for idx := 1; idx <= len(lhs); idx++ {
		// String numbers to compute sum
		var lhsStr, rhsStr string

		// Get the current digit from the lhs string
		lhsStr = string(lhs[len(lhs)-idx])
		// rhsStr may not exist if the number is shorter than lhs
		// N + 0 = N
		rhsStr = "0"

		// If exists, get the current digit from the rhs string
		if len(rhs)-idx >= 0 {
			rhsStr = string(rhs[len(rhs)-idx])
		}

		// Int numbers to compute sum
		var lhsValue, rhsValue int

		// Convert lhsValue to int
		lhsValue, err = strconv.Atoi(lhsStr)
		if err != nil {
			return zero, err
		}

		// Convert rhsValue to int
		rhsValue, err = strconv.Atoi(rhsStr)
		if err != nil {
			return zero, err
		}

		// Compute sum with carry
		sum := carry + lhsValue + rhsValue
		// Compute carry for the next iteration
		carry = sum / 10

		// Store the result
		result[len(lhs)-idx] = strconv.Itoa(sum % 10)
	}

	// If there is a carry left, add it to the result
	if carry > 0 {
		result = append([]string{strconv.Itoa(carry)}, result...)
	}

	// Join the result to a string
	resultStr := strings.Join(result, "")

	return resultStr, nil
}
