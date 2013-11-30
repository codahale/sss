package sss

import (
	"testing"
)

func equal(a, b polynomial) bool {
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
	return false
}

func TestPolyEquals(t *testing.T) {
	p3 := polynomial{70, 32, 5}

	if !equal(p, p) {
		t.Errorf("%v should equal %v", p, p)
	}

	if equal(p, p2) {
		t.Errorf("%v should not equal %v", p, p2)
	}

	if equal(p, p3) {
		t.Errorf("%v should not equal %v", p, p3)
	}
}
