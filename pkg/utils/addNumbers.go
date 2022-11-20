package utils

import (
	"errors"
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

func addStringNumbers(lhs, rhs string) (string, error) {
	// TODO:
	// 1. Make sure that len(lhs) >= len(rhs)
	// 2. Split lhs and rhs into two parts: integer and decimal
	// 3. Iterate over len(lhs) from right to left
	// 4. Get index values for lhs and rhs
	// 6. If the sum is greater than 10, then set the carry to 1
	// 7. Add the sum to the result
	return "", nil
}
