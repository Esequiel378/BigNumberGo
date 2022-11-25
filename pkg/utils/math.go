package utils

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Abs returns the absolute value of x.
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}

	return x
}
