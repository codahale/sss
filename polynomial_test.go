package sss

import (
	"testing"
)

var (
	p  = polynomial{1, 0, 2, 3}
	p2 = polynomial{70, 32, 6}
)

func TestPolyDegree(t *testing.T) {
	if p.degree() != 3 {
		t.Errorf("Expected %v to be a 3rd degree polynomial but was %d", p, p.degree())
	}
}

func TestPolyEval(t *testing.T) {
	expected := element(17)
	actual := p.eval(2)
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
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
