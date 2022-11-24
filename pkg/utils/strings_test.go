package utils

import (
	"fmt"
	"testing"
)

func TestStringToUint32(t *testing.T) {
	tests := []struct {
		input string
		want  uint32
		err   error
	}{
		{
			input: "4294967295",
			want:  4294967295,
			err:   nil,
		},
		{
			// INFO: This is one digit more than the max value of uint32
			input: "42949672951",
			want:  0,
			err:   ErrNumberOutOfRange,
		},
		{
			input: "42949672954294967295",
			want:  0,
			err:   ErrParsingIntegerNumber,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			result, err := StringToUint32(tc.input)

			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
			}

			if result != tc.want {
				t.Errorf("got %v, want %v", result, tc.want)
			}
		})
	}
}

func TestChunkString(t *testing.T) {
	tests := []struct {
		input string
		size  int
		want  []string
	}{
		{
			input: "",
			size:  10,
			want:  []string{""},
		},
		{
			input: "Hello world",
			size:  10,
			want:  []string{"Hello worl", "d"},
		},
		{
			input: "4294967295",
			size:  10,
			want:  []string{"4294967295"},
		},
		{
			input: "42949672954294967295",
			size:  10,
			want:  []string{"4294967295", "4294967295"},
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			result := ChunkString(tc.input, tc.size)

			if len(result) != len(tc.want) {
				t.Errorf("got %v, want %v", result, tc.want)
			}

			for i := 0; i < len(result); i++ {
				if result[i] != tc.want[i] {
					t.Errorf("got %v, want %v", result, tc.want)
				}
			}
		})
	}
}
