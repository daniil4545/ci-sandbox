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

func TestAbs(t *testing.T) {
	for _, c := range []struct{ a, want int }{{5, 5}, {-5, 5}, {0, 0}} {
		if got := Abs(c.a); got != c.want {
			t.Fatalf("Abs(%d) = %d", c.a, got)
		}
	}
}

func TestNeg(t *testing.T) {
	if Neg(3) != -3 || Neg(-3) != 3 || Neg(0) != 0 {
		t.Fatal("Neg changes the sign")
	}
}
