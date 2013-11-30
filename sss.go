// Package sss implements Shamir's Secret Sharing algorithm over GF(2^8).
//
// Shamir's Secret Sharing algorithm allows you to securely share a secret with
// N people, allowing the recovery of that secret if K of those people combine
// their shares.
//
// It begins by encoding a secret as a number (e.g., 42), and generating N random
// polynomial equations of degree K-1 which have an X-intercept equal to the
// secret. Given K=3, the following equations might be generated:
//
//     f1(x) =  78x^2 +  19x + 42
//     f2(x) = 128x^2 + 171x + 42
//     etc.
//
// These polynomials are then evaluated for values of X > 0:
//
//     f1(1) = 139
//     f2(2) = 896
//     etc.
//
// These (x, y) pairs are the shares given to the parties. In order to combine
// shares to recover the secret, these (x, y) pairs  are used as the input points
// for Lagrange interpolation, which produces a polynomial which matches the
// given points. This polynomial can be evaluated for f(0), producing the secret
// value--the common x-intercept for all the generated polynomials.
//
// If fewer than K shares are combined, the interpolated polynomial will be
// wrong, and the result of f(0) will not be the secret.
//
// This package constructs polynomials over the field GF(2^8) for each byte of
// the secret, allowing for much faster splitting and combining.
//
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
