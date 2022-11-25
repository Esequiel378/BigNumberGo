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

		if !isLHSDecimal {
			lhsElm += ".0"
		}

		if !isRHSDecimal {
			rhsElm += ".0"
		}

		lhs, err := NewBigFloat(lhsElm)
		if err != nil {
			return zero, err
		}

		rhs, err := NewBigFloat(rhsElm)
		if err != nil {
			return zero, err
		}

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

		results[idx] = result
	}

	resultString := strings.Join(results, " ")

	return resultString, nil
}
