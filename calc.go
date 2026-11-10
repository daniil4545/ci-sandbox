// Package calc is a tiny library the famulus review is tried on.
package calc

import "errors"

// ErrDivZero is returned when dividing by zero.
var ErrDivZero = errors.New("division by zero")

// Div returns the quotient of dividend and divisor.
func Div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, ErrDivZero
	}
	return dividend / divisor, nil
}

// Mod returns the remainder of a divided by b.
func Mod(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivZero
	}
	return a % b, nil
}

// Abs returns the absolute value of a.
func Abs(a int) int {
	if a < 0 {
		return Neg(a)
	}
	return a
}

// Neg returns a with the opposite sign.
func Neg(a int) int {
	return -a
}
