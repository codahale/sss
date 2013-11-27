package gf256

import (
	"io"
)

// RandPoly returns a random polynomial of the given degree with the given
// x-intercept. This should only be used with crypto/rand.Reader outside of
// testing.
func RandPoly(degree int, x Element, rand io.Reader) (Polynomial, error) {
	result := make(Polynomial, degree+1)
	result[0] = x

	buf := make([]byte, degree-1)
	_, err := io.ReadFull(rand, buf)
	if err != nil {
		return nil, err
	}

	for i := 1; i < degree; i++ {
		result[i] = Element(buf[i-1])
	}

	for {
		buf = make([]byte, 1)
		_, err := io.ReadFull(rand, buf)
		if err != nil {
			return nil, err
		}

		if buf[0] != 0 {
			result[degree] = Element(buf[0])
			return result, nil
		}
	}
}
