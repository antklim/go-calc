package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	testCases := []struct {
		desc     string
		lambda   calc.Lambda
		args     []float64
		expected float64
	}{
		{
			desc:     "applies Add to arguments",
			lambda:   calc.Add,
			args:     []float64{1, 2},
			expected: 3,
		},
		{
			desc:     "applies Sub to arguments",
			lambda:   calc.Sub,
			args:     []float64{1, 2},
			expected: -1,
		},
		{
			desc:     "applies Mul to arguments",
			lambda:   calc.Mul,
			args:     []float64{1, 2},
			expected: 2,
		},
		{
			desc:     "applies Div to arguments",
			lambda:   calc.Div,
			args:     []float64{1, 2},
			expected: 0.5,
		},
		{
			desc:     "applies Sqrt to arguments",
			lambda:   calc.Sqrt,
			args:     []float64{1},
			expected: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := calc.Do(tC.lambda, tC.args)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestLambdas(t *testing.T) {
	testCases := []struct {
		desc     string
		lambda   calc.Lambda
		a        float64
		b        float64
		expected float64
	}{
		{
			desc:     "sums arguments when handler is Add",
			lambda:   calc.Add,
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			desc:     "subtracts arguments when handler is Sub",
			lambda:   calc.Sub,
			a:        1,
			b:        2,
			expected: -1,
		},
		{
			desc:     "multiplies arguments when handler is Mul",
			lambda:   calc.Mul,
			a:        1,
			b:        2,
			expected: 2,
		},
		{
			desc:     "divides arguments when handler is Div",
			lambda:   calc.Div,
			a:        1,
			b:        2,
			expected: 0.5,
		},
		{
			desc:     "returns square root of argument when handler is Sqrt",
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
