package gf256

import (
	"testing"
)

func TestInterpolation(t *testing.T) {
	expected := Polynomial{0, 143}
	actual := Interpolate([][2]Element{
		[2]Element{1, 1},
		[2]Element{2, 2},
		[2]Element{3, 3},
	})

	if !Equal(actual, expected) {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}
