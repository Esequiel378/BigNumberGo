package bignumber

import "errors"

var (
	// ErrInvalidIntegerNumber is returned when the input string is not a valid integer number
	ErrInvalidIntegerNumber = errors.New("invalid integer number")
	// ErrInvalidDecimalNumber is returned when the input string is not a valid decimal number
	ErrInvalidDecimalNumber = errors.New("invalid decimal number")
	// ErrConvertingChunkToInteger is returned when a chunk cannot be converted to integer
	ErrConvertingChunkToInteger = errors.New("error converting chunk to integer")
)

// BigNumber is a number with arbitrary precision.
type BigNumber interface {
	String() string
	Add(BigNumber) BigNumber
}
