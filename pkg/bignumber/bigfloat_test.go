package bignumber

import "testing"

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
			input: "1.2.3",
			want:  "",
			err:   ErrInvalidDecimalNumber,
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
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
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
