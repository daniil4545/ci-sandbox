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
