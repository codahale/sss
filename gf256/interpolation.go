package gf256

// Interpolate uses Lagrange interpolation to calculate the x-intercept
// of the polynomial which best fits the given points.
func Interpolate(points [][2]Element, x Element) (value Element) {
	for i, a := range points {
		weight := Element(1)
		for j, b := range points {
			if i != j {
				top := x.Sub(b[0])
				bottom := a[0].Sub(b[0])
				factor := top.Div(bottom)
				weight = weight.Mul(factor)
			}
		}
		value = value.Add(weight.Mul(a[1]))
	}
	return
}
