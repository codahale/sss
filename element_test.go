package sss

import (
	"testing"
)

func TestElementAddition(t *testing.T) {
	expected := element(112)
	actual := element(100).add(element(20))
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementSubtraction(t *testing.T) {
	expected := element(112)
	actual := element(100).sub(element(20))
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementLog(t *testing.T) {
	expected := element(226)
	actual := element(90).log()
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementMultiplication(t *testing.T) {
	expected := element(254)
	actual := element(90).mul(21)
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementDivision(t *testing.T) {
	expected := element(189)
	actual := element(90).div(element(21))
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementZeroDivision(t *testing.T) {
	expected := element(0)
	actual := element(0).div(element(2))
	if actual != expected {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}

func TestElementDivideByZero(t *testing.T) {
	defer func() {
		m := recover()
		if m != "div by zero" {
			t.Error(m)
		}
	}()

	element(2).div(element(0))
	t.Error("Shouldn't have been able to divide those")
}
