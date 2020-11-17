package calc

import "math"

// LambdaFunc ...
type LambdaFunc func(arg float64) float64

// Lambda ...
type Lambda func(arg float64) LambdaFunc

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
