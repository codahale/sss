package gf256

import (
	"testing"
)

func TestElementAddition(t *testing.T) {
	expected := Element(112)
	actual := Element(100).Add(Element(20))
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementLog(t *testing.T) {
	expected := Element(226)
	actual := Element(90).Log()
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementMultiplication(t *testing.T) {
	expected := Element(254)
	actual := Element(90).Mul(21)
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementDivision(t *testing.T) {
	expected := Element(189)
	actual := Element(90).Div(Element(21))
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}
