package sss

import (
	"testing"
)

func TestInterpolation(t *testing.T) {
	expected := element(0)
	actual := interpolate(
		[]pair{
			pair{x: 1, y: 1},
			pair{x: 2, y: 2},
			pair{x: 3, y: 3},
		},
		0)

	if expected != actual {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}
