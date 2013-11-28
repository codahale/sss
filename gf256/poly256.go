package gf256

import (
	"fmt"
	"strings"
)

const (
	zero = Element(0)
)

// Polynomial is a polynomial whose coefficients are elements in GF(256).
// Element i is the coefficient for x^i.
type Polynomial []Element

// Degree returns the number of terms in the polynomial.
func (p Polynomial) Degree() int {
	return len(p) - 1
}

func (p Polynomial) String() string {
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

// Eval returns f(x) given x.
func (p Polynomial) Eval(x Element) (result Element) {
	// Horner's scheme
	for i := 1; i <= len(p); i++ {
		result = result.Mul(x).Add(p[len(p)-i])
	}
	return
}
