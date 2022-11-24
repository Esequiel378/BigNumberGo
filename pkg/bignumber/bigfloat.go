package bignumber

import (
	"regexp"
	"strings"
)

// BigFloat is a decimal number with arbitrary precision.
type BigFloat struct {
	integer *BigInt
	decimal *BigInt
}

// InputNumberMatch is a regex that matches an integer or decimal number
// Ex: 123, 123.456, 123.456789
//
// [101 reference](https://regex101.com/r/8w9NWG/1)
const InputNumberMatch = "^[0-9]*[.][0-9]+?$"

// NewBigFloat creates a new BigFloat from a string
// The string must be a valid integer number
// and must contain decimal places
//
// Ex: 123.3, 123456789012345678901234567890.1123123, etc.
func NewBigFloat(value string) (*BigFloat, error) {
	// Validate input value
	match, err := regexp.MatchString(InputNumberMatch, value)
	if !match || err != nil {
		return nil, ErrInvalidDecimalNumber
	}

	integer, decimal, found := strings.Cut(value, ".")
	if !found {
		return nil, ErrInvalidDecimalNumber
	}

	integerBigInt, err := NewBigInt(integer)
	if err != nil {
		return nil, err
	}

	decimalBigInt, err := NewBigInt(decimal)
	if err != nil {
		return nil, err
	}

	bigFloat := &BigFloat{
		integer: integerBigInt,
		decimal: decimalBigInt,
	}

	return bigFloat, nil
}

// String returns the string representation of the BigFloat.
func (b BigFloat) String() string {
	integer := b.integer.String()
	decimal := b.decimal.String()

	return integer + "." + decimal
}
