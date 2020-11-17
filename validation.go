package calc

import "errors"

// Validation describes operation validation rules.
type Validation func(args ...float64) error

func commonValidation(args ...float64) error {
	if args == nil {
		return errors.New("arguments list cannot be nil")
	}

	return nil
}

func minOneOperandValidation(args ...float64) error {
	if err := commonValidation(args...); err != nil {
		return err
	}

	if len(args) < 1 {
		return errors.New("arguments list should have at least one element")
	}

	return nil
}

func minTwoOperandsValidation(args ...float64) error {
	if err := commonValidation(args...); err != nil {
		return err
	}

	if len(args) < 2 {
		return errors.New("arguments list should have at least two elements")
	}

	return nil
}

// AddValidation addition arguments validation rules.
var AddValidation Validation = minTwoOperandsValidation

// SubValidation subtraction arguments validation rules.
var SubValidation Validation = minTwoOperandsValidation

// MulValidation multiplication arguments validation rules.
var MulValidation Validation = minTwoOperandsValidation

// DivValidation division arguments validation rules.
func DivValidation(args ...float64) error {
	if err := minTwoOperandsValidation(args...); err != nil {
		return err
	}

	if args[1] == 0 {
		return errors.New("divisor cannot be zero")
	}

	return nil
}

// SqrtValidation square root arguments validation rules.
func SqrtValidation(args ...float64) error {
	if err := minOneOperandValidation(args...); err != nil {
		return err
	}

	if args[0] < 0 {
		return errors.New("square root operand cannot be negative")
	}

	return nil
}
