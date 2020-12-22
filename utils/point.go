package utils

type Point2D struct {
	X, Y int
}

func NewPoint2D(x, y int) Point2D {
	return Point2D{x, y}
}

func (p1 Point2D) Add(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y}
}

type Point3D struct {
	X, Y, Z float64
}

func NewPoint3D(x, y, z float64) Point3D {
	return Point3D{x, y, z}
}
