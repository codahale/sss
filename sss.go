// Package sss implements the Shamir Secret Sharing algorithm over GF(2^8).
// This package has not been audited by cryptography or security professionals.
package sss

import (
	"crypto/rand"
)

// Split the given secret into N shares of which K are required to recover the
// secret. Returns a map of share IDs (1-255) to shares.
func Split(n, k int, secret []byte) (map[int][]byte, error) {
	shares := make(map[int][]byte, n)

	for _, b := range secret {
		p, err := randPoly(k-1, element(b), rand.Reader)
		if err != nil {
			return nil, err
		}

		for x := 1; x <= n; x++ {
			y := p.eval(element(x))
			shares[x] = append(shares[x], byte(y))
		}
	}

	return shares, nil
}

// Combine the given shares into the original secret.
// N.B.: There is no way to know whether the returned value is, in fact, the
// original secret.
func Combine(shares map[int][]byte) []byte {
	var secret []byte
	for _, v := range shares {
		secret = make([]byte, len(v))
		break
	}

	points := make([][2]element, len(shares))
	for i := range secret {
		p := 0
		for k, v := range shares {
			points[p][0] = element(k)
			points[p][1] = element(v[i])
			p++
		}

		s := interpolate(points, 0)
		secret[i] = byte(s)
	}

	return secret
}
