package bignumber

import (
	"fmt"
	"testing"
)

func TestNewBigFloat(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   error
	}{
		{
			input: "1.2",
			want:  "1.2",
			err:   nil,
		},
		{
			input: "1.01",
			want:  "1.01",
			err:   nil,
		},
		{
			input: "asdf.01",
			want:  "",
			err:   ErrConvertingChunkToInteger,
		},
		{
			input: "01.asdf",
			want:  "",
			err:   ErrConvertingChunkToInteger,
		},
		{
			// Will be splited into 1 1.2.3
			input: "1.2.3",
			want:  "",
			err:   ErrConvertingChunkToInteger,
		},
		{
			input: "123",
			want:  "",
			err:   ErrInvalidDecimalNumber,
		},
		{
			input: "",
			want:  "",
			err:   ErrInvalidDecimalNumber,
		},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			bg, err := NewBigFloat(tt.input)
			if err != tt.err {
				t.Errorf("got %v, want %v", err, tt.err)
				return
			}

			if bg != nil && bg.String() != tt.want {
				t.Errorf("got %v, want %v", bg, tt.want)
			}
		})
	}
}

func TestNewBigFloatAdd(t *testing.T) {
	tests := []struct {
		lhs  string
		rhs  string
		want string
		err  error
	}{
		{
			lhs:  "1.2",
			rhs:  "1.2",
			want: "2.4",
			err:  nil,
		},
		{
			lhs:  "12347612074612984761239.21127612734691273469127461293748612340",
			rhs:  "897712341234.340282366920938463463374607431768211455",
			want: "12347612075510697102473.551558494267851198154649220369254334855",
			err:  nil,
		},
		{
			lhs:  "1239.21127612734691273469",
			rhs:  "1234.1234712349812",
			want: "2473.33474736232811273469",
			err:  nil,
		},
		{
			lhs:  "1234567.8901",
			rhs:  "12.34",
			want: "1234580.2301",
			err:  nil,
		},
		{
			lhs:  "2.345",
			rhs:  "2345678901.2",
			want: "2345678903.545",
			err:  nil,
		},
		{
			lhs:  "1239.91127612734",
			rhs:  "1234.1",
			want: "2474.01127612734",
			err:  nil,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			// INFO: Errors are ignored since `TestNewBigFloat`
			// already tests the constructor
			bg1, _ := NewBigFloat(tc.lhs)
			bg2, _ := NewBigFloat(tc.rhs)

			bg, err := bg1.Add(bg2)
			if err != tc.err {
				t.Errorf("got %v, want %v", err, tc.err)
				return
			}

			if bg.String() != tc.want {
				t.Errorf("got %v, want %v", bg, tc.want)
			}
		})
	}
}
