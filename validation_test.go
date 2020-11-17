package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestValidations(t *testing.T) {
	testCases := []struct {
		desc       string
		args       []float64
		validation calc.Validation
		assert     func(t *testing.T, err error)
	}{
		{
			desc:       "addition validation returns error when arguments list is nil",
			args:       nil,
			validation: calc.AddValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list cannot be nil")
			},
		},
		{
			desc:       "addition validation returns error when arguments list is empty",
			args:       []float64{},
			validation: calc.AddValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list should have at least two elements")
			},
		},
		{
			desc:       "addition validation returns nil when arguments list is valid",
			args:       []float64{1, 2},
			validation: calc.AddValidation,
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			desc:       "subtraction validation returns error when arguments list is nil",
			args:       nil,
			validation: calc.SubValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list cannot be nil")
			},
		},
		{
			desc:       "subtraction validation returns error when arguments list is empty",
			args:       []float64{},
			validation: calc.SubValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list should have at least two elements")
			},
		},
		{
			desc:       "subtraction validation returns nil when arguments list is valid",
			args:       []float64{1, 2},
			validation: calc.SubValidation,
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			desc:       "multiplication validation returns error when arguments list is nil",
			args:       nil,
			validation: calc.MulValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list cannot be nil")
			},
		},
		{
			desc:       "multiplication validation returns error when arguments list is empty",
			args:       []float64{},
			validation: calc.MulValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list should have at least two elements")
			},
		},
		{
			desc:       "multiplication validation returns nil when arguments list is valid",
			args:       []float64{1, 2},
			validation: calc.MulValidation,
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			desc:       "division validation returns error when arguments list is nil",
			args:       nil,
			validation: calc.DivValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list cannot be nil")
			},
		},
		{
			desc:       "division validation returns error when arguments list is empty",
			args:       []float64{},
			validation: calc.DivValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list should have at least two elements")
			},
		},
		{
			desc:       "division validation returns error when divisor is zero",
			args:       []float64{1, 0},
			validation: calc.DivValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "divisor cannot be zero")
			},
		},
		{
			desc:       "division validation returns nil when arguments list is valid",
			args:       []float64{1, 2},
			validation: calc.DivValidation,
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			desc:       "square root validation returns error when arguments list is nil",
			args:       nil,
			validation: calc.SqrtValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list cannot be nil")
			},
		},
		{
			desc:       "square root validation returns error when arguments list is empty",
			args:       []float64{},
			validation: calc.SqrtValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list should have at least one element")
			},
		},
		{
			desc:       "square root validation returns error when argument is negative",
			args:       []float64{-1},
			validation: calc.SqrtValidation,
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "square root argument cannot be negative")
			},
		},
		{
			desc:       "square root validation returns nil when arguments list is valid",
			args:       []float64{1},
			validation: calc.SqrtValidation,
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.validation(tC.args)
			tC.assert(t, err)
		})
	}
}
