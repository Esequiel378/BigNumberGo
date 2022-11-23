package utils

import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInputWithDifferentNumbersCount = errors.New("input with different numbers count")
	ErrInvalidCarryValue              = errors.New("invalid carry value")
	ErrInvalidInputNumber             = errors.New("invalid input number")
	ErrInvalidIntegerNumber           = errors.New("invalid integer number")
	ErrInvalidDecimalNumber           = errors.New("invalid decimal number")
)

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

	results := make([]string, len(lhsElements))

	for idx := 0; idx < len(lhsElements); idx++ {
		// TODO: chekIsInBounds
		lhsElm, rhsElm := lhsElements[idx], rhsElements[idx]

		var result string
		result, err = addStringNumbers(lhsElm, rhsElm)
		if err != nil {
			return zero, err
		}

		results[idx] = result
	}

	resultString := strings.Join(results, " ")

	return resultString, nil
}

// InputNumberMatch is a regex that matches an integer or decimal number
// Ex: 123, 123.456, 123.456789
//
// [101 reference](https://regex101.com/r/SuRnhV/1)
const InputNumberMatch = "^[0-9]*[.]?[0-9]+?$"

// addStringNumbers adds two string numbers and returns the result
// The numbers may include decimal places.
//
// Examples
//  >> addStringNumbers("123", "11") "134"
//
//  >> addStringNumbers("123456789012345678901", "12345678")
//  "123456789012358024579"
//
//  >> addStringNumbers("1234567.8901", "12.34")
//  "1234582.2301"
func addStringNumbers(lhs, rhs string) (string, error) {
	var (
		match bool
		zero string
		err  error
	)

	// Validate lhs input value
	if match, err = regexp.MatchString(InputNumberMatch, lhs); !match || err != nil {
		return zero, ErrInvalidInputNumber
	}

	// Validate rhs input value
	if match, err = regexp.MatchString(InputNumberMatch, rhs); !match || err != nil {
		return zero, ErrInvalidInputNumber
	}

	// Make sure that lhs is the largest number
	if len(lhs) < len(rhs) {
		lhs, rhs = rhs, lhs
	}

	// Get the integer part of the numbers
	lhs, lhsDecimal, lhsFound := strings.Cut(lhs, ".")
	rhs, rhsDecimal, rhsFound := strings.Cut(rhs, ".")

	// Make sure that lhs is the largest number
	if len(lhsDecimal) < len(rhsDecimal) {
		lhsDecimal, rhsDecimal = rhsDecimal, lhsDecimal
	}

	// Fill the missing decimal part with zeros
	rhsDecimal += strings.Repeat("0", len(lhsDecimal)-len(rhsDecimal))

	var (
		carry 	int
		decimal string
	)

	// Compute the sum of the decimal part
	if lhsFound && rhsFound {
		decimal, carry, err = addTwoStringNumbers(lhsDecimal, rhsDecimal)
		if err != nil {
			return zero, err
		}
	}

	// compute addition
	result, carry, err := addTwoStringNumbers(lhs, rhs, WithCarry(carry))
	if err != nil {
		return zero, err
	}

	// If there is a carry left, add it to the result
	if carry > 0 {
		// Carry can be only 1
		result = "1" + result
	}

	if len(decimal) > 0 {
		result += "." + decimal
	}

	return result, nil
}

// AddTwoStringNumbersOptions is a function that can be used to modify the behaviour of addTwoStringNumbers
//
// This is based on the [Option Pattern  in Golang](https://blog.damavis.com/en/option-pattern-in-golang)
type AddTwoStringNumbersOptions = func(lhs, rhs string, carry int) (string, string, int)

// WithCarry is an option that can be used to add a initial carry to the addition
func WithCarry(carry int) AddTwoStringNumbersOptions {
	return func(lhs, rhs string, _ int) (string, string, int) {
		return lhs, rhs, carry
	}
}

// IntegerNumberMatch is a regex that matches an integer number (without decimal places)
// Ex: 123, 123456789012345678901234567890
//
// [101 reference](https://regex101.com/r/3hoFC3/1)
const IntegerNumberMatch = "^[0-9]+$"

// addTwoStringNumbers adds two integers numbers from a string
// and returns the result with the carry of the last digits
// The numbers must not include decimal places.
//
// Examples:
//  >> addTwoStringNumbers("123", "456")
//  "579", 0, nil
//
//  >> addTwoStringNumbers("99", "9")
//  "08", 1, nil
func addTwoStringNumbers(
	lhs, rhs string,
	opts ...AddTwoStringNumbersOptions,
) (string, int, error) {
	var (
		match	bool
		zero  	string
		carry 	int
		err  	error
	)

	// Validate lhs input value
	if match, err = regexp.MatchString(IntegerNumberMatch, lhs); !match || err != nil {
		return zero, carry, ErrInvalidIntegerNumber
	}

	// Validate rhs input value
	if match, err = regexp.MatchString(IntegerNumberMatch, rhs); !match || err != nil {
		return zero, carry, ErrInvalidIntegerNumber
	}

	for _, opt := range opts {
		lhs, rhs, carry = opt(lhs, rhs, carry)
	}

	if carry < 0 || carry > 1 {
		return zero, carry, ErrInvalidCarryValue
	}

	// Make sure that lhs is the largest number
	if len(lhs) < len(rhs) {
		lhs, rhs = rhs, lhs
	}

	// result store the digits of the sum
	result := make([]string, len(lhs))

	// Iterate from the end of the string
	for idx := 1; idx <= len(lhs); idx++ {
		// String numbers to compute sum
		var lhsStr, rhsStr string

		// Get the current digit from the lhs string
		lhsStr = string(lhs[len(lhs)-idx])
		// rhsStr may not exist if the number is shorter than lhs
		// X + 0 = X
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
			err = fmt.Errorf("failed to covert string `%s` into int: %v", lhsStr, err)
			return zero, carry, err
		}

		// Convert rhsValue to int
		rhsValue, err = strconv.Atoi(rhsStr)
		if err != nil {
			err = fmt.Errorf("failed to covert string `%s` into int: %v", rhsStr, err)
			return zero, carry, err
		}

		// Compute sum with carry
		sum := carry + lhsValue + rhsValue
		// Compute carry for the next iteration
		carry = sum / 10

		// Store the result
		result[len(lhs)-idx] = strconv.Itoa(sum % 10)
	}

	resultStr := strings.Join(result, "")

	return resultStr, carry, nil
}
