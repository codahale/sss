package sss

// Lagrange interpolation
func interpolate(points [][2]element, x element) (value element) {
	for i, a := range points {
		weight := element(1)
		for j, b := range points {
			if i != j {
				top := x.sub(b[0])
				bottom := a[0].sub(b[0])
				factor := top.div(bottom)
				weight = weight.mul(factor)
			}
		}
		value = value.add(weight.mul(a[1]))
	}
	return
}
