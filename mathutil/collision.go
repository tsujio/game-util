package mathutil

const epsilon = -1e-6

func PointLineDistance(p, p1, v1 *Vector2D) (d float64, h *Vector2D, t float64) {
	t = 0.0
	if l := v1.NormSq(); l > 0 {
		t = v1.Dot(p.Sub(p1)) / l
	}
	h = p1.Add(v1.Mul(t))
	d = h.Sub(p).Norm()
	return
}

func PointLineSegmentDistance(p, p1, v1 *Vector2D) (d float64, h *Vector2D, t float64) {
	e1 := p1.Add(v1)

	d, h, t = PointLineDistance(p, p1, v1)

	if p.Sub(p1).Dot(e1.Sub(p1)) < 0 {
		h = p1.Clone()
		d = p1.Sub(p).Norm()
	} else if p.Sub(e1).Dot(p1.Sub(e1)) < 0 {
		h = e1.Clone()
		d = e1.Sub(p).Norm()
	}

	return
}

func LineLineDistance(p1, v1, p2, v2 *Vector2D) (d float64, h1 *Vector2D, t1 float64, h2 *Vector2D, t2 float64) {
	if cross := v1.Cross(v2).Z; -epsilon < cross || cross < epsilon {
		d, h2, t2 = PointLineDistance(p1, p2, v2)
		h1 = p1.Clone()
		t1 = 0
		return
	}

	dv1v2 := v1.Dot(v2)
	dv1v1 := v1.NormSq()
	dv2v2 := v2.NormSq()
	p21p11 := p1.Sub(p2)
	t1 = (dv1v2*v2.Dot(p21p11) - dv2v2*v1.Dot(p21p11)) / (dv1v1*dv2v2 - dv1v2*dv1v2)
	h1 = p1.Add(v1.Mul(t1))
	t2 = v2.Dot(h1.Sub(p2)) / dv2v2
	h2 = p2.Add(v2.Mul(t2))
	d = h2.Sub(h1).Norm()
	return
}

func clamp(t float64) float64 {
	if t < 0 {
		return 0
	} else if t > 1 {
		return 1
	} else {
		return t
	}
}

func LineSegmentLineSegmentDistance(p1, v1, p2, v2 *Vector2D) (d float64, h1 *Vector2D, t1 float64, h2 *Vector2D, t2 float64) {
	if l1, l2 := v1.NormSq(), v2.NormSq(); l1 < 1e-6 {
		if l2 < 1e-6 {
			d = p2.Sub(p1).Norm()
			h1 = p1.Clone()
			h2 = p2.Clone()
			t1 = 0
			t2 = 0
			return
		} else {
			d, h2, t2 = PointLineSegmentDistance(p1, p2, v2)
			h1 = p1.Clone()
			t1 = 0
			t2 = clamp(t2)
			return
		}
	} else if l2 < 1e-6 {
		d, h1, t1 = PointLineSegmentDistance(p2, p1, v1)
		h2 = p2.Clone()
		t1 = clamp(t1)
		t2 = 0
		return
	}

	if cross := v1.Cross(v2).Z; -epsilon < cross || cross < epsilon {
		h1 = p1.Clone()
		t1 = 0
		if d, h2, t2 = PointLineSegmentDistance(p1, p2, v2); 0 <= t2 && t2 < 1 {
			return
		}
	} else if d, h1, t1, h2, t2 = LineLineDistance(p1, v1, p2, v2); 0 <= t1 && t1 <= 1 && 0 <= t2 && t2 <= 1 {
		return
	}

	t1 = clamp(t1)
	h1 = p1.Add(v1.Mul(t1))
	if d, h2, t2 = PointLineSegmentDistance(h1, p2, v2); 0 <= t2 && t2 <= 1 {
		return
	}

	t2 = clamp(t2)
	h2 = p2.Add(v2.Mul(t2))
	if d, h1, t1 = PointLineSegmentDistance(h2, p1, v1); 0 <= t1 && t1 <= 1 {
		return
	}

	t1 = clamp(t1)
	h1 = p1.Add(v1.Mul(t1))
	d = h2.Sub(h1).Norm()

	return
}

func CapsulesCollide(p1, v1 *Vector2D, r1 float64, p2, v2 *Vector2D, r2 float64) bool {
	d, _, _, _, _ := LineSegmentLineSegmentDistance(p1, v1, p2, v2)
	return d <= r1+r2
}
