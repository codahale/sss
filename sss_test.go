package sss

import (
	"testing"
)

func TestRoundtrip(t *testing.T) {
	n := 30
	k := 2

	expected := "well hello there!"
	shares, err := Split(n, k, []byte(expected))
	if err != nil {
		t.Error(err)
	}

	subset := make(map[int][]byte, k)
	for x, y := range shares {
		subset[x] = y
		if len(subset) == k {
			break
		}
	}

	actual := string(Combine(subset))
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
