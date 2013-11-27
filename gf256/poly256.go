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

// Add returns the sum of the receiver and the argument.
func (p Polynomial) Add(a Polynomial) Polynomial {
	minDeg, maxDeg := degrees(p, a)
	big := p
	if len(p) < len(a) {
		big = a
	}

	coeff := make(Polynomial, 0)

	for i := 0; i < minDeg; i++ {
		coeff = append(coeff, p[i].Add(a[i]))
	}

	for i := minDeg; i < maxDeg; i++ {
		coeff = append(coeff, big[i])
	}

	return coeff
}

// AddElement returns the sum of the receiver and the argument.
func (p Polynomial) AddElement(a Element) Polynomial {
	coeff := make(Polynomial, len(p))
	for i, v := range p {
		coeff[i] = v
	}
	coeff[0] = coeff[0].Add(a)
	return coeff
}

// SubElement returns the different between the receiver and the argument.
func (p Polynomial) SubElement(a Element) Polynomial {
	coeff := make(Polynomial, len(p))
	for i, v := range p {
		coeff[i] = v
	}
	coeff[0] = coeff[0].Sub(a)
	return coeff
}

// Sub returns the difference between the receiver and the argument.
func (p Polynomial) Sub(a Polynomial) Polynomial {
	minDeg, maxDeg := degrees(p, a)

	coeff := make(Polynomial, 0)

	for i := 0; i <= minDeg; i++ {
		coeff = append(coeff, p[i].Add(a[i]))
	}

	for i := minDeg + 1; i <= maxDeg; i++ {
		if p.Degree() > a.Degree() {
			coeff = append(coeff, p[i])
		} else {
			coeff = append(coeff, zero.Add(a[i]))
		}
	}

	return coeff
}

// Mul returns the product of the receiver and the argument.
func (p Polynomial) Mul(a Polynomial) Polynomial {
	coeff := make(Polynomial, p.Degree()+a.Degree()+1)
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(a); j++ {
			coeff[i+j] = coeff[i+j].Add(p[i].Mul(a[j]))
		}
	}
	return coeff
}

// MulElement returns the product of the receiver and the argument.
func (p Polynomial) MulElement(a Element) Polynomial {
	coeff := make(Polynomial, len(p))
	for i, v := range p {
		coeff[i] = v.Mul(a)
	}
	return coeff
}

// Equal returns a boolean reporting whether a == b.
func Equal(a, b Polynomial) bool {
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
	return false
}

func degrees(a, b Polynomial) (int, int) {
	minDeg := a.Degree()
	if b.Degree() < minDeg {
		minDeg = b.Degree()
	}

	maxDeg := a.Degree()
	if b.Degree() > maxDeg {
		maxDeg = b.Degree()
	}

	return minDeg, maxDeg
}
