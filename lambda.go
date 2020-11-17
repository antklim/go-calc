package calc

import "math"

type Lambda func(arg float64) float64
type LambdaHandler func(arg float64) Lambda

var LambdaHandlers = map[string]LambdaHandler{
	"add":  Add,
	"sub":  Sub,
	"mul":  Mul,
	"div":  Div,
	"sqrt": Sqrt,
}

// Add implements addition.
func Add(a float64) Lambda {
	return func(b float64) float64 {
		return a + b
	}
}

// Mul implements multiplication.
func Mul(a float64) Lambda {
	return func(b float64) float64 {
		return a * b
	}
}

// Sub implements subtraction.
func Sub(a float64) Lambda {
	return func(b float64) float64 {
		return a - b
	}
}

// Div implements division.
func Div(a float64) Lambda {
	return func(b float64) float64 {
		return a / b
	}
}

// Sqrt implements square root.
func Sqrt(a float64) Lambda {
	return func(_ float64) float64 {
		return math.Sqrt(a)
	}
}
