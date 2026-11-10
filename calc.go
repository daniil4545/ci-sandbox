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

// IsEven reports whether a divides by 2.
func IsEven(a int) bool {
	return remainder(a, 2) == 0
}

// remainder returns what is left of a divided by b, 0 for a zero b.
func remainder(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
