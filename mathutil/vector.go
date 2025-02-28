package mathutil

import "math"

type Vector2D struct {
	X, Y float64
}

func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{X: x, Y: y}
}

func (v *Vector2D) Clone() *Vector2D {
	return &Vector2D{X: v.X, Y: v.Y}
}

func (v *Vector2D) Add(w *Vector2D) *Vector2D {
	return &Vector2D{X: v.X + w.X, Y: v.Y + w.Y}
}

func (v *Vector2D) Sub(w *Vector2D) *Vector2D {
	return &Vector2D{X: v.X - w.X, Y: v.Y - w.Y}
}

func (v *Vector2D) Mul(a float64) *Vector2D {
	return &Vector2D{X: v.X * a, Y: v.Y * a}
}

func (v *Vector2D) Div(a float64) *Vector2D {
	return &Vector2D{X: v.X / a, Y: v.Y / a}
}

func (v *Vector2D) NormSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector2D) Norm() float64 {
	return math.Sqrt(v.NormSq())
}

func (v *Vector2D) Normalize() *Vector2D {
	return v.Div(v.Norm())
}

func (v *Vector2D) Dot(w *Vector2D) float64 {
	return v.X*w.X + v.Y*w.Y
}

func (v *Vector2D) Cross(w *Vector2D) *Vector3D {
	return &Vector3D{X: 0, Y: 0, Z: v.X*w.Y - v.Y*w.X}
}

func (v *Vector2D) Rotate(theta float64) *Vector2D {
	return &Vector2D{X: math.Cos(theta)*v.X - math.Sin(theta)*v.Y, Y: math.Sin(theta)*v.X + math.Cos(theta)*v.Y}
}

func (v *Vector2D) Unpack() (float64, float64) {
	return v.X, v.Y
}

type Vector3D struct {
	X, Y, Z float64
}

func NewVector3D(x, y, z float64) *Vector3D {
	return &Vector3D{X: x, Y: y, Z: z}
}

func (v *Vector3D) Clone() *Vector3D {
	return &Vector3D{X: v.X, Y: v.Y, Z: v.Z}
}

func (v *Vector3D) Add(w *Vector3D) *Vector3D {
	return &Vector3D{X: v.X + w.X, Y: v.Y + w.Y, Z: v.Z + w.Z}
}

func (v *Vector3D) Sub(w *Vector3D) *Vector3D {
	return &Vector3D{X: v.X - w.X, Y: v.Y - w.Y, Z: v.Z - w.Z}
}

func (v *Vector3D) Mul(a float64) *Vector3D {
	return &Vector3D{X: v.X * a, Y: v.Y * a, Z: v.Z * a}
}

func (v *Vector3D) Div(a float64) *Vector3D {
	return &Vector3D{X: v.X / a, Y: v.Y / a, Z: v.Z / a}
}

func (v *Vector3D) NormSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector3D) Norm() float64 {
	return math.Sqrt(v.NormSq())
}

func (v *Vector3D) Normalize() *Vector3D {
	return v.Div(v.Norm())
}

func (v *Vector3D) Dot(w *Vector3D) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v *Vector3D) Cross(w *Vector3D) *Vector3D {
	return &Vector3D{X: v.Y*w.Z - v.Z*w.Y, Y: v.Z*w.X - v.X*w.Z, Z: v.X*w.Y - v.Y*w.X}
}

func (v *Vector3D) Unpack() (float64, float64, float64) {
	return v.X, v.Y, v.Z
}
