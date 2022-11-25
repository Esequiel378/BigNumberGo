package bignumber

import (
	"errors"
	"fmt"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		lhs  string
		rhs  string
		want string
		err  error
	}{
		{
			lhs:  "1 2 3",
			rhs:  "1 2 3",
			want: "2 4 6",
			err:  nil,
		},
		{
			lhs:  "1 2 3 4",
			rhs:  "1 2 3 4",
			want: "2 4 6 8",
			err:  nil,
		},
		{
			lhs:  "1.3 2 3",
			rhs:  "1 2 3",
			want: "2.3 4 6",
			err:  nil,
		},
		{
			lhs:  "1 2 3",
			rhs:  "1 2 3 4",
			want: "",
			err:  ErrInputWithDifferentNumbersCount,
		},
		{
			lhs:  "1 2 3 4",
			rhs:  "1 2 3",
			want: "",
			err:  ErrInputWithDifferentNumbersCount,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			got, err := AddNumbers(tc.lhs, tc.rhs)
			if !errors.Is(err, tc.err) {
				t.Errorf("got %v, want %v", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
