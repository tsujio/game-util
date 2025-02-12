package mathutil

import (
	"math"
)

func ConvertCoordinateWorldToScreen(pos *Vector3D, cameraY float64, screenZ float64, screenWidth, screenHeight int) *Vector2D {
	x := pos.X * screenZ / pos.Z
	y := (pos.Y - cameraY) * screenZ / pos.Z
	return &Vector2D{X: x + float64(screenWidth)/2, Y: y + float64(screenHeight)/2}
}

func ConvertCoordinateScreenToWorld(pos *Vector2D, dst *Vector3D, cameraY float64, screenZ float64, screenWidth, screenHeight int) *Vector3D {
	count := 0
	if !math.IsNaN(dst.X) {
		count++
	}
	if !math.IsNaN(dst.Y) {
		count++
	}
	if !math.IsNaN(dst.Z) {
		count++
	}
	if count != 1 {
		panic("dst must have just one non-NaN value")
	}

	if !math.IsNaN(dst.X) {
		srcX := pos.X - float64(screenWidth)/2
		y := (pos.Y-float64(screenHeight)/2)*dst.X/srcX + cameraY
		z := screenZ * dst.X / srcX
		return &Vector3D{X: dst.X, Y: y, Z: z}
	} else if !math.IsNaN(dst.Y) {
		ratio := (dst.Y - cameraY) / (pos.Y - float64(screenHeight)/2)
		x := (pos.X - float64(screenWidth)/2) * ratio
		z := screenZ * ratio
		return &Vector3D{X: x, Y: dst.Y, Z: z}
	} else if !math.IsNaN(dst.Z) {
		x := (pos.X - float64(screenWidth)/2) * dst.Z / screenZ
		y := (pos.Y-float64(screenHeight)/2)*dst.Z/screenZ + cameraY
		return &Vector3D{X: x, Y: y, Z: dst.Z}
	} else {
		panic("error")
	}
}
