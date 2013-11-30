package sss

type polynomial []element

func (p polynomial) degree() int {
	return len(p) - 1
}

func (p polynomial) eval(x element) (result element) {
	// Horner's scheme
	for i := 1; i <= len(p); i++ {
		result = result.mul(x).add(p[len(p)-i])
	}
	return
}
