package sss

import (
	"fmt"
	"strings"
)

const (
	zero = element(0)
)

// polynomial is a polynomial whose coefficients are elements in GF(256).
// Element i is the coefficient for x^i.
type polynomial []element

// degree returns the number of terms in the polynomial.
func (p polynomial) degree() int {
	return len(p) - 1
}

func (p polynomial) String() string {
	coeffs := make([]string, 0)
	for i := len(p) - 1; i >= 0; i-- {
		e := p[i]
		if int(e) != 0 {
			switch i {
			case 0:
				coeffs = append(coeffs, fmt.Sprintf("%v", e))
			case 1:
				coeffs = append(coeffs, fmt.Sprintf("%vx", e))
			default:
				coeffs = append(coeffs, fmt.Sprintf("%vx^%d", e, i))
			}
		}
	}
	return strings.Join(coeffs, "+")
}

// eval returns f(x) given x.
func (p polynomial) eval(x element) (result element) {
	// Horner's scheme
	for i := 1; i <= len(p); i++ {
		result = result.mul(x).add(p[len(p)-i])
	}
	return
}
