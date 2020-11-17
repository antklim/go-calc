package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestLambdaOperations(t *testing.T) {
	testCases := []struct {
		desc     string
		oper     calc.Operation
		a        float64
		b        float64
		expected float64
	}{
		{
			desc:     "adds operands when oper is Add",
			oper:     calc.Add,
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			desc:     "subtracts operands when oper is Sub",
			oper:     calc.Sub,
			a:        1,
			b:        2,
			expected: -1,
		},
		{
			desc:     "multiplies operands when oper is Mul",
			oper:     calc.Mul,
			a:        1,
			b:        2,
			expected: 2,
		},
		{
			desc:     "divides operands when oper is Div",
			oper:     calc.Div,
			a:        1,
			b:        2,
			expected: 0.5,
		},
		{
			desc:     "returns square root of operand when oper is Sqrt",
			oper:     calc.Sqrt,
			a:        1,
			expected: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := calc.OperationHandlers[tC.oper](tC.a)(tC.b)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
