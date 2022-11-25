package bignumber

import (
	"errors"
	"strings"
)

var (
	// ErrInvalidIntegerNumber is returned when the input string is not a valid integer number.
	ErrInvalidIntegerNumber = errors.New("invalid integer number")
	// ErrInvalidDecimalNumber is returned when the input string is not a valid decimal number.
	ErrInvalidDecimalNumber = errors.New("invalid decimal number")
	// ErrConvertingChunkToInteger is returned when a chunk cannot be converted to integer.
	ErrConvertingChunkToInteger = errors.New("error converting chunk to integer")
	// ErrInputWithDifferentNumbersCount is returned when the input strings have different numbers count.
	ErrInputWithDifferentNumbersCount = errors.New("input with different numbers count")
	// ErrTrimmingDecimalPart is returned when the decimal part cannot be trimmed.
	ErrTrimmingDecimalPart = errors.New("error trimming decimal part")
)

// AddNumbers takse two string params containing M numbers
// separated by spaces and returns sum of the pairs.
//
// + The input parameters should have the same count of numbers.
// + The numbers may include decimal places.
// + The numbers can be arbitrarily long, e.g. 1000+ digits.
//
// Examples:
//
//	>> AddNumbers("123 456 789", "11 22 33")
//	"134 478 822"
//
//	>> AddNumbers("123456789012345678901 23456789", "12345678 234567890123456789012")
//	"123456789012358024579 234567890123480245801"
//
//	>> AddNumbers("1234567.8901 2.345", "12.34 2345678901.2")
//	"1234582.2301 2345678903.545"
func AddNumbers(lhs, rhs string) (string, error) {
	var zero string

	lhsElements, rhsElements := strings.Split(lhs, " "), strings.Split(rhs, " ")

	if len(lhsElements) != len(rhsElements) {
		return zero, ErrInputWithDifferentNumbersCount
	}

	results := make([]string, len(lhsElements))

	for idx := 0; idx < len(lhsElements); idx++ {
		// TODO: chekIsInBounds
		lhsElm, rhsElm := lhsElements[idx], rhsElements[idx]

		isLHSDecimal := strings.Contains(lhsElm, ".")
		isRHSDecimal := strings.Contains(rhsElm, ".")

		isDecimalAddition := isLHSDecimal || isRHSDecimal

		// Convert both numbers to decimal numbers
		// if one of the numbers already is.
		if !isLHSDecimal {
			lhsElm += ".0"
		}

		if !isRHSDecimal {
			rhsElm += ".0"
		}

		// Create BigFloat intances
		lhs, err := NewBigFloat(lhsElm)
		if err != nil {
			return zero, err
		}

		rhs, err := NewBigFloat(rhsElm)
		if err != nil {
			return zero, err
		}

		// Perform aritmetic addition
		sum, err := lhs.Add(rhs)
		if err != nil {
			return zero, err
		}

		result := sum.String()

		// Remove the decimal part if the result should be an integer.
		if !isDecimalAddition {
			var found bool

			result, _, found = strings.Cut(result, ".")
			if !found {
				return zero, ErrTrimmingDecimalPart
			}
		}

		// Store the result
		results[idx] = result
	}

	// Jois the results to a single string
	resultString := strings.Join(results, " ")

	return resultString, nil
}
