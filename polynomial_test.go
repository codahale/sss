package sss

import (
	"testing"
)

var (
	p  = polynomial{1, 0, 2, 3}
	p2 = polynomial{70, 32, 6}
)

func TestPolyDegree(t *testing.T) {
	expected := 3
	actual := p.degree()
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyEval(t *testing.T) {
	expected := element(17)
	actual := p.eval(2)
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
