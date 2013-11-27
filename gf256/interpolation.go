package gf256

// Interpolate uses Lagrange interpolation to return a polynomial which fits the
// given points.
func Interpolate(points [][2]Element) Polynomial {
	p := Polynomial{0}
	for i, a := range points {
		p1 := Polynomial{1}
		x := Polynomial{0, 1}
		for j, b := range points {
			if j == i {
				continue
			}
			y := Element(1).Div(points[i][0].Add(b[0]))
			p1 = p1.Mul(x.AddElement(b[0]).MulElement(y))
		}
		p = p.Add(p1.MulElement(a[1]))
	}
	return p
}
