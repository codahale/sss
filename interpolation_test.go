package sss

import (
	"testing"
)

func TestInterpolation(t *testing.T) {
	expected := element(0)
	actual := interpolate(
		[][2]element{
			[2]element{1, 1},
			[2]element{2, 2},
			[2]element{3, 3},
		},
		0)

	if expected != actual {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}
