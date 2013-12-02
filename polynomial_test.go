package sss

import (
	"bytes"
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

func TestRandPoly(t *testing.T) {
	b := []byte{1, 2, 3}

	expected := polynomial{10, 1, 2, 3}
	actual, err := randPoly(3, 10, bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}

	if !equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestRandPolyEOF(t *testing.T) {
	b := []byte{1}

	p, err := randPoly(3, 10, bytes.NewReader(b))
	if p != nil {
		t.Errorf("Expected an error but got %v", p)
	}

	if err == nil {
		t.Error("No error returned")
	}
}

func TestRandPolyEOFFullSize(t *testing.T) {
	b := []byte{1, 2, 0, 0, 0, 0}

	p, err := randPoly(3, 10, bytes.NewReader(b))
	if p != nil {
		t.Errorf("Expected an error but got %v", p)
	}

	if err == nil {
		t.Error("No error returned")
	}
}

func TestRandPolyFullSize(t *testing.T) {
	b := []byte{1, 2, 0, 4}

	expected := polynomial{10, 1, 2, 4}
	actual, err := randPoly(3, 10, bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}

	if !equal(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

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
