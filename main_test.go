package main

import (
	"fmt"
	"testing"
)

func TestExper(t *testing.T) {
	parameters := []struct {
		input    string
		expected int
	}{
		{
			"12+34+56",
			102,
		},
		{
			"100+12-5",
			107,
		},
		{
			"0-1",
			-1,
		},
		{
			"-1-1",
			-2,
		},
		{
			"2*3-1",
			5,
		},
		{
			"2+3*5",
			17,
		},
		{
			"2+6/2",
			5,
		},
		{
			"(2+3)*4",
			20,
		},
	}

	for _, tt := range parameters {
		t.Run(fmt.Sprintf("calc %s", tt.input), func(t *testing.T) {
			s := &Source{Str: tt.input}
			if a := s.Expr(); a != tt.expected {
				t.Logf("expected: %d, actual: %d", tt.expected, a)
				t.Fail()
			}
		})
	}
}
