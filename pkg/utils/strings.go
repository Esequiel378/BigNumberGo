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

// ChunkString breaks a string into chunks of a given size.
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

		chunks = append(chunks, value[start:end])
	}

	return chunks
}
