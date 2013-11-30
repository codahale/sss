package sss

import (
	"fmt"
	"strings"
)

type polynomial []element

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

func (p polynomial) eval(x element) (result element) {
	// Horner's scheme
	for i := 1; i <= len(p); i++ {
		result = result.mul(x).add(p[len(p)-i])
	}
	return
}
