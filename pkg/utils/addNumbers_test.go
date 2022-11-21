package utils

import (
	"errors"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		lhs, rhs string
		want     string
		err      error
	}{
		{
			lhs:  "123 456 789",
			rhs:  "11 22 33",
			want: "134 478 822",
			err:  nil,
		},
		{
			lhs:  "123456789012345678901 23456789",
			rhs:  "12345678 234567890123456789012",
			want: "123456789012358024579 234567890123480245801",
			err:  nil,
		},
		{
			lhs:  "99",
			rhs:  "9",
			want: "108",
			err:  nil,
		},
		{
			lhs:  "1234567.8901 2.345",
			rhs:  "12.34 2345678901.2",
			want: "1234580.2301 2345678903.545",
			err:  nil,
		},
	}

	for idx, tc := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			got, err := AddNumbers(tc.lhs, tc.rhs)
			if !errors.Is(err, tc.err) {
				t.Errorf("got error `%v`, want `%v`", err, tc.err)
			}

			if got != tc.want {
				t.Errorf("got `%s`, want `%s`", got, tc.want)
			}
		})
	}
}
