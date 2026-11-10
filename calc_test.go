package calc

import (
	"errors"
	"testing"
)

func TestDiv(t *testing.T) {
	if got, err := Div(7, 2); err != nil || got != 3 {
		t.Fatalf("Div(7, 2) = %d, %v", got, err)
	}
	if _, err := Div(1, 0); !errors.Is(err, ErrDivZero) {
		t.Fatalf("Div(1, 0) err = %v", err)
	}
}

func TestMod(t *testing.T) {
	if got, err := Mod(7, 2); err != nil || got != 1 {
		t.Fatalf("Mod(7, 2) = %d, %v", got, err)
	}
	if _, err := Mod(1, 0); !errors.Is(err, ErrDivZero) {
		t.Fatalf("Mod(1, 0) err = %v", err)
	}
}

func TestAvg(t *testing.T) {
	for _, c := range []struct {
		xs   []int
		want int
	}{{[]int{0, 4, 8}, 4}, {[]int{0, -3, -4}, -2}, {[]int{0}, 0}} {
		if got, err := Avg(c.xs...); err != nil || got != c.want {
			t.Fatalf("Avg(%v) = %d, %v", c.xs, got, err)
		}
	}
	if _, err := Avg(); !errors.Is(err, ErrEmpty) {
		t.Fatalf("Avg() err = %v", err)
	}
}
