package fi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
)

type testCase struct {
	name string
	in   fi.Fi
	out  fi.Fi
}

func TestNext(t *testing.T) {
	tests := []testCase{
		{
			name: "zero val",
			in:   fi.Fi{},
			out: fi.Fi{
				Previous: 0,
				Current:  1,
			},
		},
		{
			name: "1 current",
			in: fi.Fi{
				Previous: 0,
				Current:  1,
			},
			out: fi.Fi{
				Previous: 1,
				Current:  1,
			},
		},
	}
	for _, tc := range tests {
		result := fi.Next(tc.in)
		assert.Equal(t, tc.out, result)
	}
}
func TestPrev(t *testing.T) {
	tests := []testCase{
		{
			name: "zero val",
			in:   fi.Fi{},
			out: fi.Fi{
				Previous: 0,
				Current:  0,
			},
		},
		{
			name: "1 val",
			in: fi.Fi{
				Previous: 0,
				Current:  1,
			},
			out: fi.Fi{
				Previous: 0,
				Current:  0,
			},
		},
		{
			name: "1 next val",
			in: fi.Fi{
				Previous: 1,
				Current:  1,
			},
			out: fi.Fi{
				Previous: 0,
				Current:  1,
			},
		},
		{
			name: "1 next^2 val",
			in: fi.Fi{
				Previous: 1,
				Current:  2,
			},
			out: fi.Fi{
				Previous: 1,
				Current:  1,
			},
		},
		{
			name: "1 next^3 val",
			in: fi.Fi{
				Previous: 2,
				Current:  3,
			},
			out: fi.Fi{
				Previous: 1,
				Current:  2,
			},
		},
	}
	for _, tc := range tests {
		result := fi.Previous(tc.in)
		assert.Equal(t, tc.out, result)
	}
}
