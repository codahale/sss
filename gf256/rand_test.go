package gf256

import (
	"bytes"
	"testing"
)

func TestRandPoly(t *testing.T) {
	b := []byte{1, 2, 3}

	expected := Polynomial{10, 1, 2, 3}
	actual, err := RandPoly(3, 10, bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}

	if !Equal(expected, actual) {
		t.Errorf("Expected \n%v but was \n%v", expected, actual)
	}
}

func TestRandPolyEOF(t *testing.T) {
	b := []byte{1}

	p, err := RandPoly(3, 10, bytes.NewReader(b))
	if p != nil {
		t.Errorf("Expected an error but got %v", p)
	}

	if err == nil {
		t.Error("No error returned")
	}
}

func TestRandPolyEOFFullSize(t *testing.T) {
	b := []byte{1, 2, 0, 0, 0, 0}

	p, err := RandPoly(3, 10, bytes.NewReader(b))
	if p != nil {
		t.Errorf("Expected an error but got %v", p)
	}

	if err == nil {
		t.Error("No error returned")
	}
}

func TestRandPolyFullSize(t *testing.T) {
	b := []byte{1, 2, 0, 4}

	expected := Polynomial{10, 1, 2, 4}
	actual, err := RandPoly(3, 10, bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}

	if !Equal(expected, actual) {
		t.Errorf("Expected \n%v but was \n%v", expected, actual)
	}
}
