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

// Clamp limits v to the range from lo to hi.
func Clamp(v, lo, hi int) int {
	return min(max(v, lo), hi)
}
