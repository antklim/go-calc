package calc

import "errors"

// Validation describes operation validation rules.
type Validation func(args []float64) error

func commonValidation(args []float64) error {
	if args == nil {
		return errors.New("arguments list cannot be nil")
	}

	return nil
}

func minOneArgumentValidation(args []float64) error {
	if err := commonValidation(args); err != nil {
		return err
	}

	if len(args) < 1 {
		return errors.New("arguments list should have at least one element")
	}

	return nil
}

func minTwoArgumentsValidation(args []float64) error {
	if err := commonValidation(args); err != nil {
		return err
	}

	if len(args) < 2 {
		return errors.New("arguments list should have at least two elements")
	}

	return nil
}

// AddValidation addition arguments validation rules.
var AddValidation Validation = minTwoArgumentsValidation

// SubValidation subtraction arguments validation rules.
var SubValidation Validation = minTwoArgumentsValidation

// MulValidation multiplication arguments validation rules.
var MulValidation Validation = minTwoArgumentsValidation

// DivValidation division arguments validation rules.
func DivValidation(args []float64) error {
	if err := minTwoArgumentsValidation(args); err != nil {
		return err
	}

	if args[1] == 0 {
		return errors.New("divisor cannot be zero")
	}

	return nil
}

// SqrtValidation square root arguments validation rules.
func SqrtValidation(args []float64) error {
	if err := minOneArgumentValidation(args); err != nil {
		return err
	}

	if args[0] < 0 {
		return errors.New("square root argument cannot be negative")
	}

	return nil
}
