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
		{
			lhs:  "12347612074612984761239.21127612734691273469127461293748612340 2 3",
			rhs:  "1 2 3",
			want: "12347612074612984761240.21127612734691273469127461293748612340 4 6",
			err:  nil,
		},
		{
			lhs:  "12347612074612984761239.2112761270 2 3",
			rhs:  "76127346912734.461291234761 2 3",
			want: "12347612150740331673973.672567361761 4 6",
			err:  nil,
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
