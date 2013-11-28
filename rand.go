package sss

import (
	"io"
)

// randPoly returns a random polynomial of the given degree with the given
// x-intercept. This should only be used with crypto/rand.Reader outside of
// testing.
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
