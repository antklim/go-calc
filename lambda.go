package calc

import "math"

type Lambda func(arg float64) float64
type Handler func(arg float64) Lambda

var OperationHandlers = map[Operation]Handler{
	Add:  add,
	Sub:  sub,
	Mul:  mul,
	Div:  div,
	Sqrt: sqrt,
}

func add(a float64) Lambda {
	return func(b float64) float64 {
		return a + b
	}
}

func mul(a float64) Lambda {
	return func(b float64) float64 {
		return a * b
	}
}

func sub(a float64) Lambda {
	return func(b float64) float64 {
		return a - b
	}
}

func div(a float64) Lambda {
	return func(b float64) float64 {
		return a / b
	}
}

func sqrt(a float64) Lambda {
	return func(_ float64) float64 {
		return math.Sqrt(a)
	}
}
