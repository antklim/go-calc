package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestRequestValidation(t *testing.T) {
	testCases := []struct {
		desc   string
		req    calc.Request
		assert func(t *testing.T, err error)
	}{
		{
			desc: "returns error when unsupported operation provided",
			req:  calc.Request{Operation: "foo"},
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "unsupported operation foo")
			},
		},
		{
			desc: "propagates operation validation error",
			req:  calc.Request{Operation: "add"},
			assert: func(t *testing.T, err error) {
				assert.EqualError(t, err, "arguments list cannot be nil")
			},
		},
		{
			desc: "returns nil when request is valid",
			req:  calc.Request{Operation: "add", Arguments: []float64{1, 2}},
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.req.Validate()
			tC.assert(t, err)
		})
	}
}
