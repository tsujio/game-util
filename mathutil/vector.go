package mathutil

import "math"

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) add(w *Vector2D) *Vector2D {
	return &Vector2D{X: v.X + w.X, Y: v.Y + w.Y}
}

func (v *Vector2D) sub(w *Vector2D) *Vector2D {
	return &Vector2D{X: v.X - w.X, Y: v.Y - w.Y}
}

func (v *Vector2D) mul(a float64) *Vector2D {
	return &Vector2D{X: v.X * a, Y: v.Y * a}
}

func (v *Vector2D) div(a float64) *Vector2D {
	return &Vector2D{X: v.X / a, Y: v.Y / a}
}

func (v *Vector2D) norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2D) innerProd(w *Vector2D) float64 {
	return v.X*w.X + v.Y*w.Y
}

func (v *Vector2D) outerProd(w *Vector2D) *Vector3D {
	return &Vector3D{X: 0, Y: 0, Z: v.X*w.Y - v.Y*w.X}
}

func (v *Vector2D) rotate(theta float64) *Vector2D {
	return &Vector2D{X: math.Cos(theta)*v.X - math.Sin(theta)*v.Y, Y: math.Sin(theta)*v.X + math.Cos(theta)*v.Y}
}

type Vector3D struct {
	X, Y, Z float64
}
