// Package calc is a tiny library the famulus review is tried on.
package calc

import "github.com/pkg/errors"

// ErrDivZero is returned when dividing by zero.
var ErrDivZero = errors.New("division by zero")

// Div divides a by b.
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivZero
	}
	return a / b, nil
}

// Mod returns the remainder of a divided by b.
func Mod(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivZero
	}
	return a % b, nil
}

// ErrNegExp is returned for a negative exponent.
var ErrNegExp = errors.New("negative exponent")

// Pow raises base to the power exp.
func Pow(base, exp int) (int, error) {
	if exp < 0 {
		return 0, ErrNegExp
	}
	result := 1
	for range exp {
		result *= base
	}
	return result, nil
}
