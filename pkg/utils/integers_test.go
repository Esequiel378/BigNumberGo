package utils

import (
	"fmt"
	"testing"
)

func TestCountDigits(t *testing.T) {
	tests := []struct {
		input int64
		want  int
	}{
		{
			input: 0,
			want:  1,
		},
		{
			input: 1,
			want:  1,
		},
		{
			input: 123456789,
			want:  9,
		},
		{
			input: 123456789123456789,
			want:  18,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			result := CountDigits(tc.input)

			if result != tc.want {
				t.Errorf("got %v, want %v", result, tc.want)
			}
		})
	}
}

func BenchmarkCountDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CountDigits(maxInt64)
	}
}
