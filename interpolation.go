package sss

type pair struct {
	x, y element
}

// Lagrange interpolation
func interpolate(points []pair, x element) (value element) {
	for i, a := range points {
		weight := element(1)
		for j, b := range points {
			if i != j {
				top := x.sub(b.x)
				bottom := a.x.sub(b.x)
				factor := top.div(bottom)
				weight = weight.mul(factor)
			}
		}
		value = value.add(weight.mul(a.y))
	}
	return
}
