package bigint

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"teladoc/pkg/utils"
)

type BigInt struct {
	magnitude []uint32
}

var ErrInvalidIntegerNumber = errors.New("invalid integer number")

// IntegerNumberMatch is a regex that matches an integer number (without decimal places)
// Ex: 123, 123456789012345678901234567890
//
// [101 reference](https://regex101.com/r/3hoFC3/1)
const IntegerNumberMatch = "^[0-9]+$"

func NewBigInt(value string) (*BigInt, error) {
	// Validate input value
	if match, err := regexp.MatchString(IntegerNumberMatch, value); !match || err != nil {
		return nil, ErrInvalidIntegerNumber
	}

	// Break the string into chunks of 9 digits
	chunks := utils.ChunkString(value, 9)
	magnitude := make([]uint32, len(chunks))

	// Convert each chunk to uint32
	for idx, chunk := range chunks {
		integer, err := utils.StringToUint32(chunk)
		if err != nil {
			return nil, err
		}

		magnitude[idx] = integer
	}

	bigInt := &BigInt{
		magnitude: magnitude,
	}

	return bigInt, nil
}

// String returns the string representation of the BigInt
func (b *BigInt) String() string {
	var sb strings.Builder

	for _, chunk := range b.magnitude {
		value := strconv.FormatUint(uint64(chunk), 10)
		sb.WriteString(value)
	}

	return sb.String()
}
