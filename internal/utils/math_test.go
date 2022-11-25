package utils

import (
	"fmt"
	"testing"
)

func TestAbsWithInt64(t *testing.T) {
	tests := []struct {
		input int64
		want  int64
	}{
		{
			input: -1,
			want:  1,
		},
		{
			input: 1,
			want:  1,
		},
		{
			input: 0,
			want:  0,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			got := Abs(tc.input)

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAbsWithFloat64(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{
			input: -1.0,
			want:  1.0,
		},
		{
			input: 1.0,
			want:  1.0,
		},
		{
			input: 0.0,
			want:  0.0,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			got := Abs(tc.input)

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
