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
			// INFO: This is one digit more than the max value of uint32.
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

func TestChunkStringFromRight(t *testing.T) {
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
			want:  []string{"H", "ello world"},
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
		{
			input: "12347612074612984761239",
			size:  10,
			want:  []string{"123", "4761207461", "2984761239"},
		},
		{
			input: "897712341234",
			size:  10,
			want:  []string{"89", "7712341234"},
		},
		{
			input: "10000000000",
			size:  9,
			want:  []string{"10", "000000000"},
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			result := ChunkStringFromRight(tc.input, tc.size)

			if len(result) != len(tc.want) {
				t.Errorf("got %v, want %v", result, tc.want)
				return
			}

			for i := 0; i < len(result); i++ {
				if result[i] != tc.want[i] {
					t.Errorf("got %v, want %v", result[i], tc.want[i])
				}
			}
		})
	}
}

func TestRemoveLeadingZeros(t *testing.T) {
	tests := []struct {
		input string
		count int64
		want  string
	}{
		{
			input: "001234",
			count: 2,
			want:  "1234",
		},
		{
			input: "000",
			count: 3,
			want:  "",
		},
		{
			input: "1234",
			count: 0,
			want:  "1234",
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			got, count := RemoveLeadingZeros(tc.input)

			if got != tc.want || count != tc.count {
				t.Errorf("got %v, %v, want %v, %v", got, count, tc.want, tc.count)
			}
		})
	}
}
