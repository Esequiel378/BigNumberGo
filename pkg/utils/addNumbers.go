package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type BigInt struct {
	magnitude []uint32
}

var (
	ErrInvalidIntegerNumber = errors.New("invalid integer number")
	ErrParsingIntegerNumber = errors.New("error parsing integer number")
	ErrNumberOutOfRange     = errors.New("number out of range")
)

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
	chunks := ChunkString(value, 9)
	magnitude := make([]uint32, len(chunks))

	// Convert each chunk to uint32
	for idx, chunk := range chunks {
		integer, err := StringToUint32(chunk)
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

const maxUint32 = uint64(^uint32(0))

// StringToUint32 converts a string to uint32
func StringToUint32(value string) (uint32, error) {
	integer, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, ErrParsingIntegerNumber
	}

	if integer > maxUint32 {
		return 0, ErrNumberOutOfRange
	}

	result := uint32(integer)

	return result, nil
}

// ChunkString breaks a string into chunks of a given size
func ChunkString(value string, chunkSize int) []string {
	if len(value) <= chunkSize {
		return []string{value}
	}

	var chunks []string

	for start := 0; start < len(value); start += chunkSize {
		end := start + chunkSize

		if end > len(value) {
			end = len(value)
		}

		chunks = append(chunks, string(value[start:end]))
	}

	return chunks
}
