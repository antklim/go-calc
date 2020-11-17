package calc_test

import (
	"testing"

	"github.com/antklim/go-calc"
	"github.com/stretchr/testify/assert"
)

func TestDoUseCase(t *testing.T) {
	testCases := []struct {
		desc   string
		req    calc.Request
		assert func(t *testing.T, res *calc.Response, err error)
	}{
		{
			desc: "returns error when request validation fails",
			req:  calc.Request{},
			assert: func(t *testing.T, res *calc.Response, err error) {
				assert.Nil(t, res)
				assert.EqualError(t, err, "unsupported operation ")
			},
		},
		{
			desc: "returns response for a valid request",
			req: calc.Request{
				Operation: "add",
				Arguments: []float64{3, 4},
			},
			assert: func(t *testing.T, res *calc.Response, err error) {
				assert.NoError(t, err)
				expected := calc.Response{
					Operation: "add",
					Arguments: []float64{3, 4},
					Result:    7,
				}
				assert.Equal(t, expected, *res)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res, err := calc.DoUseCase(tC.req)
			tC.assert(t, res, err)
		})
	}
}
