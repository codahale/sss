package gf256

import (
	"testing"
)

func TestInterpolation(t *testing.T) {
	expected := Element(0)
	actual := Interpolate(
		[][2]Element{
			[2]Element{1, 1},
			[2]Element{2, 2},
			[2]Element{3, 3},
		},
		0)

	if expected != actual {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}
