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

func TestClamp(t *testing.T) {
	for _, c := range []struct{ v, lo, hi, want int }{{5, 0, 10, 5}, {-1, 0, 10, 0}, {11, 0, 10, 10}, {3, 3, 3, 3}} {
		if got := Clamp(c.v, c.lo, c.hi); got != c.want {
			t.Fatalf("Clamp(%d, %d, %d) = %d", c.v, c.lo, c.hi, got)
		}
	}
}
