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

// ErrEmpty is returned when there are no values.
var ErrEmpty = errors.New("no values")

// Avg returns the mean of xs rounded toward zero.
func Avg(xs ...int) (int, error) {
	if len(xs) == 0 {
		return 0, ErrEmpty
	}
	sum := 0
	for i := 1; i < len(xs); i++ {
		sum += xs[i]
	}
	return sum / len(xs), nil
}
