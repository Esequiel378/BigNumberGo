package bigint

import (
	"fmt"
	"testing"
)

func TestNewBigInt(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   error
	}{
		{
			input: "4294967295",
			want:  "4294967295",
			err:   nil,
		},
		{
			input: "42949672954294967295",
			want:  "42949672954294967295",
			err:   nil,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			bg, err := NewBigInt(tc.input)
			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
			}

			if bg.String() != tc.want {
				t.Errorf("got %v, want %v", bg.String(), tc.want)
			}
		})
	}
}
