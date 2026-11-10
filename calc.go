// Package calc is a tiny library the famulus review is tried on.
package calc

import "errors"

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

// Gcd returns the greatest common divisor of a and b, never negative; Gcd(0, 0) is 0.
// For math.MinInt the divisor does not fit in int and the result is not defined.
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}
