package geometry

import "math"

// Line - representation of line in 2D
type Line struct {
	StartX float64
	StartY float64
	EndX   float64
	EndY   float64
}

// LineIntersection - check if two lines intersects and where
func LineIntersection(line1 Line, line2 Line) (x float64, y float64, ok bool) {
	s1X := line1.EndX - line1.StartX
	s1Y := line1.EndY - line1.StartY
	s2X := line2.EndX - line2.StartX
	s2Y := line2.EndY - line2.StartY

	s, t := 0.0, 0.0
	if (-s2X*s1Y + s1X*s2Y) != 0 {
		s = (-s1Y*(line1.StartX-line2.StartX) + s1X*(line1.StartY-line2.StartY)) / (-s2X*s1Y + s1X*s2Y)
		t = (s2X*(line1.StartY-line2.StartY) - s2Y*(line1.StartX-line2.StartX)) / (-s2X*s1Y + s1X*s2Y)
	} else {
		return 0, 0, false
	}

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		// Collision detected
		x = math.Round(line1.StartX + (t * s1X))
		y = math.Round(line1.StartY + (t * s1Y))
		return x, y, true
	}

	return 0, 0, false // No collision
}
