package calc

import "math"

// LambdaFunc ...
type LambdaFunc func(arg float64) float64

// Lambda ...
type Lambda func(arg float64) LambdaFunc

// Do applies lambda to arguments.
func Do(lambda Lambda, args []float64) float64 {
	a, b := args[0], float64(0)
	if len(args) > 1 {
		b = args[1]
	}

	return lambda(a)(b)
}

// Add implements addition.
func Add(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a + b
	}
}

// Mul implements multiplication.
func Mul(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a * b
	}
}

// Sub implements subtraction.
func Sub(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a - b
	}
}

// Div implements division.
func Div(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a / b
	}
}

// Sqrt implements square root.
func Sqrt(a float64) LambdaFunc {
	return func(_ float64) float64 {
		return math.Sqrt(a)
	}
}
