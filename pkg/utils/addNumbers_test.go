package utils

import (
	"errors"
	"fmt"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		lhs, rhs string
		result   string
		err      error
	}{
		{
			lhs:    "123 456 789",
			rhs:    "11 22 33",
			result: "134 478 822",
			err:    nil,
		},
		{
			lhs:    "123456789012345678901 23456789",
			rhs:    "12345678 234567890123456789012",
			result: "123456789012358024579 234567890123480245801",
			err:    nil,
		},
		{
			lhs:    "99",
			rhs:    "9",
			result: "108",
			err:    nil,
		},
		{
			lhs:    "1234567.8901 2.345",
			rhs:    "12.34 2345678901.2",
			result: "1234580.2301 2345678903.545",
			err:    nil,
		},
		{
			lhs:    "19.9 10",
			rhs:    "10 29.9",
			result: "29.9 39.9",
			err:    nil,
		},
		{
			lhs:    "123..3 123. .123",
			rhs:    "1 2 3",
			result: "",
			err:    ErrInvalidInputNumber,
		},
		{
			lhs:    "1 2 3",
			rhs:    "123..3 123. .123",
			result: "",
			err:    ErrInvalidInputNumber,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			got, err := AddNumbers(tc.lhs, tc.rhs)
			if !errors.Is(err, tc.err) {
				t.Errorf("got error `%v`, want `%v`", err, tc.err)
			}

			if got != tc.result {
				t.Errorf("got `%s`, want `%s`", got, tc.result)
			}
		})
	}
}

func TestAddTwoStringNumbers(t *testing.T) {
	tests := []struct {
		lhs, rhs string
		result   string
		carry    int
		opts     []AddTwoStringNumbersOptions
		err      error
	}{
		{
			lhs:    "123",
			rhs:    "456",
			result: "579",
			opts:   []AddTwoStringNumbersOptions{},
			carry:  0,
			err:    nil,
		},
		{
			lhs:    "99",
			rhs:    "9",
			result: "08",
			opts:   []AddTwoStringNumbersOptions{WithCarry(0)},
			carry:  1,
			err:    nil,
		},
		{
			lhs:    "100",
			rhs:    "9",
			result: "110",
			opts:   []AddTwoStringNumbersOptions{WithCarry(1)},
			carry:  0,
			err:    nil,
		},
		{
			lhs:    "1",
			rhs:    "1",
			result: "",
			opts:   []AddTwoStringNumbersOptions{WithCarry(3)},
			carry:  3,
			err:    ErrInvalidCarryValue,
		},
		{
			lhs:    "11.2",
			rhs:    "10.2",
			result: "",
			opts:   []AddTwoStringNumbersOptions{},
			carry:  0,
			err:    ErrInvalidIntegerNumber,
		},
	}

	for idx, tc := range tests {
		testname := fmt.Sprintf("test#%d", idx)

		t.Run(testname, func(t *testing.T) {
			result, carry, err := addTwoStringNumbers(tc.lhs, tc.rhs, tc.opts...)
			if !errors.Is(err, tc.err) {
				t.Errorf("got error `%v`, want `%v`", err, tc.err)
			}

			if result != tc.result {
				t.Errorf("got `%s`, want `%s`", result, tc.result)
			}

			if carry != tc.carry {
				t.Errorf("got `%d`, want `%d`", carry, tc.carry)
			}
		})
	}
}
