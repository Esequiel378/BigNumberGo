package bignumber

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
		{
			input: "qwer",
			want:  "",
			err:   ErrInvalidIntegerNumber,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			bg, err := NewBigInt(tc.input)
			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
			}

			if bg != nil && bg.String() != tc.want {
				t.Errorf("got %v, want %v", bg.String(), tc.want)
			}
		})
	}
}

func TestBigIntAdd(t *testing.T) {
	tests := []struct {
		lhs    string
		rhs    string
		result string
	}{
		{
			lhs:    "999999999",
			rhs:    "999999999",
			result: "1999999998",
		},
		{
			lhs:    "340282366920938463463374607431768211455",
			rhs:    "340282366920938463463374607431768211455",
			result: "680564733841876926926749214863536422910",
		},
		{
			lhs:    "21127612734691273469127461293748612340",
			rhs:    "340282366920938463463374607431768211455",
			result: "361409979655629736932502068725516823795",
		},
		{
			lhs:    "123456789012345678901",
			rhs:    "12345678",
			result: "123456789012358024579",
		},
		{
			lhs:    "23456789",
			rhs:    "234567890123456789012",
			result: "234567890123480245801",
		},
		{
			lhs:    "12347612074612984761239",
			rhs:    "897712341234",
			result: "12347612075510697102473",
		},
		{
			lhs:    "123",
			rhs:    "456",
			result: "579",
		},
		{
			lhs:    "99",
			rhs:    "9",
			result: "108",
		},
		{
			lhs:    "100",
			rhs:    "9",
			result: "109",
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			bg1, _ := NewBigInt(tc.lhs)
			bg2, _ := NewBigInt(tc.rhs)

			bg3 := bg1.Add(bg2)

			if bg3.String() != tc.result {
				t.Errorf("got %v, want %v", bg3.String(), tc.result)
			}
		})
	}
}
