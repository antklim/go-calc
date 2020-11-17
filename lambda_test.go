package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestLambdaHandlers(t *testing.T) {
	testCases := []struct {
		desc     string
		lambda   calc.Lambda
		a        float64
		b        float64
		expected float64
	}{
		{
			desc:     "adds operands when handler is Add",
			lambda:   calc.Add,
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			desc:     "subtracts operands when handler is Sub",
			lambda:   calc.Sub,
			a:        1,
			b:        2,
			expected: -1,
		},
		{
			desc:     "multiplies operands when handler is Mul",
			lambda:   calc.Mul,
			a:        1,
			b:        2,
			expected: 2,
		},
		{
			desc:     "divides operands when handler is Div",
			lambda:   calc.Div,
			a:        1,
			b:        2,
			expected: 0.5,
		},
		{
			desc:     "returns square root of operand when handler is Sqrt",
			lambda:   calc.Sqrt,
			a:        1,
			expected: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.lambda(tC.a)(tC.b)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
