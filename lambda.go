package calc

import "math"

type LambdaFunc func(arg float64) float64
type Lambda func(arg float64) LambdaFunc

func Add(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a + b
	}
}

func Mul(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a * b
	}
}

func Sub(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a - b
	}
}

func Div(a float64) LambdaFunc {
	return func(b float64) float64 {
		return a / b
	}
}

func Sqrt(a float64) LambdaFunc {
	return func(_ float64) float64 {
		return math.Sqrt(a)
	}
}
