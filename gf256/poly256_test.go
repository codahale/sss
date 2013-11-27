package gf256

import (
	"testing"
)

var (
	p  = Polynomial{1, 0, 2, 3}
	p2 = Polynomial{70, 32, 6}
)

func TestPolyString(t *testing.T) {
	expected := "6x^2+32x+70"
	actual := p2.String()
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func TestPolyDegree(t *testing.T) {
	if p.Degree() != 3 {
		t.Errorf("Expected %v to be a 3rd degree polynomial but was %d", p, p.Degree())
	}
}

func TestPolyEval(t *testing.T) {
	expected := Element(17)
	actual := p.Eval(2)
	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyAddition(t *testing.T) {
	expected := Polynomial{71, 32, 2}
	actual := p.Add(p2)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyElementAddition(t *testing.T) {
	expected := Polynomial{2, 0, 2, 3}
	actual := p.AddElement(3)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolySubtraction(t *testing.T) {
	expected := Polynomial{71, 32, 4, 3}
	actual := p.Sub(p2)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyElementSubtraction(t *testing.T) {
	expected := Polynomial{41, 0, 2, 3}
	actual := p.SubElement(40)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyMultiplication(t *testing.T) {
	expected := Polynomial{70, 32, 138, 138, 108, 10}
	actual := p.Mul(p2)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyElementMultiplication(t *testing.T) {
	expected := Polynomial{30, 0, 60, 34}
	actual := p.MulElement(30)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestPolyEquals(t *testing.T) {
	p3 := Polynomial{70, 32, 5}

	if !Equal(p, p) {
		t.Errorf("%v should equal %v", p, p)
	}

	if Equal(p, p2) {
		t.Errorf("%v should not equal %v", p, p2)
	}

	if Equal(p, p3) {
		t.Errorf("%v should not equal %v", p, p3)
	}
}
