package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestLambdaOperations(t *testing.T) {
	testCases := []struct {
		desc     string
		oper     string
		a        float64
		b        float64
		expected float64
	}{
		{
			desc:     "adds operands when oper is 'add'",
			oper:     "add",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			desc:     "subtracts operands when oper is 'sub'",
			oper:     "sub",
			a:        1,
			b:        2,
			expected: -1,
		},
		{
			desc:     "multiplies operands when oper is 'mul'",
			oper:     "mul",
			a:        1,
			b:        2,
			expected: 2,
		},
		{
			desc:     "divides operands when oper is 'div'",
			oper:     "div",
			a:        1,
			b:        2,
			expected: 0.5,
		},
		{
			desc:     "returns square root of operand when oper is 'sqrt'",
			oper:     "sqrt",
			a:        1,
			expected: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := calc.LambdaHandlers[tC.oper](tC.a)(tC.b)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
