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

func TestGcd(t *testing.T) {
	for _, c := range []struct{ a, b, want int }{{12, 18, 6}, {18, 12, 6}, {-12, 18, 6}, {12, -18, 6}, {7, 0, 7}, {0, -7, 7}, {0, 0, 0}, {17, 5, 1}} {
		if got := Gcd(c.a, c.b); got != c.want {
			t.Fatalf("Gcd(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
