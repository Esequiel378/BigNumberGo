package utils

// maxUint32 is the maximum value of a uint32
const maxUint32 = int64(^uint32(0))

// maxInt64 is the maximum value of a int64
const maxInt64 = int64(^uint64(0) >> 1)

// CountDigits returns the number of digits in a 64 bits number
func CountDigits(integer int64) int {
	if integer == 0 {
		return 1
	}

	var count int

	for integer != 0 {
		integer /= 10
		count++
	}

	return count
}
