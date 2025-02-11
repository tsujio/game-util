package mathutil

func ConvertCoordinateWorldToScreen(pos *Vector3D, cameraY float64, screenZ float64, screenWidth, screenHeight int) *Vector2D {
	x := pos.X * screenZ / pos.Z
	y := (pos.Y - cameraY) * screenZ / pos.Z
	return &Vector2D{X: x + float64(screenWidth)/2, Y: y + float64(screenHeight)/2}
}

func ConvertCoordinateScreenToWorld(pos *Vector2D, z float64, cameraY float64, screenZ float64, screenWidth, screenHeight int) *Vector3D {
	x := (pos.X - float64(screenWidth)/2) * z / screenZ
	y := (pos.Y-float64(screenHeight)/2)*z/screenZ + cameraY
	return &Vector3D{X: x, Y: y, Z: z}
}
