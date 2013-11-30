package sss

import (
	"io"
)

// generates a random n-degree polynomial w/ a given x-intercept
func randPoly(degree int, x element, rand io.Reader) (polynomial, error) {
	result := make(polynomial, degree+1)
	result[0] = x

	buf := make([]byte, degree-1)
	_, err := io.ReadFull(rand, buf)
	if err != nil {
		return nil, err
	}

	for i := 1; i < degree; i++ {
		result[i] = element(buf[i-1])
	}

	// the Nth term can't be zero, or else it's a (N-1) degree polynomial
	for {
		buf = make([]byte, 1)
		_, err := io.ReadFull(rand, buf)
		if err != nil {
			return nil, err
		}

		if buf[0] != 0 {
			result[degree] = element(buf[0])
			return result, nil
		}
	}
}
