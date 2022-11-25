package utils

import (
	"errors"
	"strconv"
)

var (
	// ErrParsingIntegerNumber is returned when a string cannot be parsed to an integer.
	ErrParsingIntegerNumber = errors.New("error parsing integer number")
	// ErrNumberOutOfRange is returned when a number is out of range.
	ErrNumberOutOfRange = errors.New("number out of range")
)

// maxUint32 is the maximum value of a uint32.
const maxUint32 = int64(^uint32(0))

// RemoveLeadingZeros removes leading zeros from a string returning
// the string without leading zeros and the count of leading zeros.
func RemoveLeadingZeros(value string) (string, int64) {
	var count int64

	for _, char := range value {
		if char == '0' {
			count++
			continue
		}

		break
	}

	// Remove leading zeros from the decimal part to avoid zeros leading chunks
	value = value[count:]

	return value, count
}

// StringToUint32 converts a string to uint32.
func StringToUint32(value string) (uint32, error) {
	integer, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, ErrParsingIntegerNumber
	}

	if int64(integer) > maxUint32 {
		return 0, ErrNumberOutOfRange
	}

	result := uint32(integer)

	return result, nil
}

// ChunkStringFromRight breaks a string into chunks of a given size from the right.
func ChunkStringFromRight(value string, chunkSize int) []string {
	if len(value) <= chunkSize {
		return []string{value}
	}

	var chunks []string

	for end := len(value); end >= 0; end -= chunkSize {
		start := end - chunkSize

		if start < 0 {
			start = 0
		}

		if start == end {
			continue
		}

		chunks = append([]string{value[start:end]}, chunks...)
	}

	return chunks
}
