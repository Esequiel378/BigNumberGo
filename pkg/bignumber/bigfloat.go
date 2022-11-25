package bignumber

import (
	"regexp"
	"strings"

	"teladoc/pkg/utils"
)

// BigFloat is a decimal number with arbitrary precision.
type BigFloat struct {
	// integer is the integer part of the number.
	integer *BigInt
	// decimal is the decimal part of the number.
	decimal *BigInt
	// leadingZeros is the number of leading zeros in the decimal part
	// Should be used to calculate the magnitude of the decimal part.
	leadingZeros string
	// precision is the number of decimal places.
	precision int64
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

	// Split integer and decimal parts
	integer, decimal, found := strings.Cut(value, ".")
	if !found {
		return nil, ErrInvalidDecimalNumber
	}

	// INFO: this should be calculated before removing the leading zeros
	precision := len(decimal)

	// Compute the number of leading zeros in the decimal part and
	// remove them to avoid zeros leading chunks
	decimal, leadingZerosCount := utils.RemoveLeadingZeros(decimal)

	// Create BigInts from integer and decimal parts
	integerBigInt, err := NewBigInt(integer)
	if err != nil {
		return nil, err
	}

	if decimal == "" {
		decimal = "0"
	}

	decimalBigInt, err := NewBigInt(decimal)
	if err != nil {
		return nil, err
	}

	leadingZeros := strings.Repeat("0", int(leadingZerosCount))

	bigFloat := &BigFloat{
		integer:      integerBigInt,
		decimal:      decimalBigInt,
		leadingZeros: leadingZeros,
		precision:    int64(precision),
	}

	return bigFloat, nil
}

// String returns the string representation of the BigFloat.
func (b BigFloat) String() string {
	integer := b.integer.String()
	decimal := b.decimal.String()

	return integer + "." + b.leadingZeros + decimal
}

// Precision returns the number of decimal places.
func (b BigFloat) Precision() int64 {
	return b.precision
}

// Add adds two BigFloats and returns the result.
func (b BigFloat) Add(other *BigFloat) (*BigFloat, error) {
	lhsDecimal := b.decimal
	rhsDecimal := other.decimal
	deltaZeros := utils.Abs(b.Precision() - other.Precision())

	// Make sure that lhs has more decimal places than rhs
	if b.Precision() < other.Precision() {
		lhsDecimal, rhsDecimal = rhsDecimal, lhsDecimal
	}

	if deltaZeros > 0 {
		zeros := strings.Repeat("0", int(deltaZeros))

		newValue := rhsDecimal.String() + zeros

		var err error

		rhsDecimal, err = NewBigInt(newValue)
		if err != nil {
			return nil, err
		}
	}

	decimal := lhsDecimal.Add(rhsDecimal)
	integer := b.integer.Add(other.integer)

	var (
		integerValue = integer.String()
		decimalValue = decimal.String()
	)

	// Check if decimal should carry over to integer
	if decimal.Length() > lhsDecimal.Length() {
		// Add carry to integer
		oneBigInt, err := NewBigInt("1")
		if err != nil {
			return nil, err
		}

		integer = integer.Add(oneBigInt)

		// Remove the carry value form the decimal and
		// update the float values
		decimalValue = decimal.String()[1:]
		integerValue = integer.String()
	}

	newFloatValue := integerValue + "." + decimalValue

	bigFloat, err := NewBigFloat(newFloatValue)
	if err != nil {
		return nil, err
	}

	return bigFloat, nil
}
