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

func TestPow(t *testing.T) {
	for _, c := range []struct{ base, exp, want int }{{2, 10, 1024}, {7, 0, 1}, {-3, 3, -27}, {0, 5, 0}} {
		if got, err := Pow(c.base, c.exp); err != nil || got != c.want {
			t.Fatalf("Pow(%d, %d) = %d, %v", c.base, c.exp, got, err)
		}
	}
	if _, err := Pow(2, -1); !errors.Is(err, ErrNegExp) {
		t.Fatalf("Pow(2, -1) err = %v", err)
	}
}
