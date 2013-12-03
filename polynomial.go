package sss

import (
	"io"
)

// a GF(2^8) polynomial
type polynomial []element

// the degree of the polynomial
func (p polynomial) degree() int {
	return len(p) - 1
}

// evaluate the polynomial at the given point
func (p polynomial) eval(x element) (result element) {
	// Horner's scheme
	for i := 1; i <= len(p); i++ {
		result = result.mul(x).add(p[len(p)-i])
	}
	return
}

// generates a random n-degree polynomial w/ a given x-intercept
func generate(degree int, x element, rand io.Reader) (polynomial, error) {
	result := make(polynomial, degree+1)
	result[0] = x

	buf := make([]byte, degree-1)
	if _, err := io.ReadFull(rand, buf); err != nil {
		return nil, err
	}

	for i := 1; i < degree; i++ {
		result[i] = element(buf[i-1])
	}

	// the Nth term can't be zero, or else it's a (N-1) degree polynomial
	for {
		buf = make([]byte, 1)
		if _, err := io.ReadFull(rand, buf); err != nil {
			return nil, err
		}

		if buf[0] != 0 {
			result[degree] = element(buf[0])
			return result, nil
		}
	}
}

// an input/output pair
type pair struct {
	x, y element
}

// Lagrange interpolation
func interpolate(points []pair, x element) (value element) {
	for i, a := range points {
		weight := element(1)
		for j, b := range points {
			if i != j {
				top := x.sub(b.x)
				bottom := a.x.sub(b.x)
				factor := top.div(bottom)
				weight = weight.mul(factor)
			}
		}
		value = value.add(weight.mul(a.y))
	}
	return
}
