package utils

// Line - representation of line in 2D
type Line[T ~int | ~float64] struct {
	StartX T
	StartY T
	EndX   T
	EndY   T
}

type Line3D struct {
	StartX int
	StartY int
	StartZ int
	EndX   int
	EndY   int
	EndZ   int
}

// LineIntersection - check if two lines intersects and where
func (line1 *Line[T]) LineIntersection(line2 Line[T]) (x float64, y float64, ok bool) {
	s1X := line1.EndX - line1.StartX
	s1Y := line1.EndY - line1.StartY
	s2X := line2.EndX - line2.StartX
	s2Y := line2.EndY - line2.StartY

	s, t := 0.0, 0.0
	if (-s2X*s1Y + s1X*s2Y) != 0 {
		s = float64(-s1Y*(line1.StartX-line2.StartX)+s1X*(line1.StartY-line2.StartY)) / float64(-s2X*s1Y+s1X*s2Y)
		t = float64(s2X*(line1.StartY-line2.StartY)-s2Y*(line1.StartX-line2.StartX)) / float64(-s2X*s1Y+s1X*s2Y)
	} else {
		return 0, 0, false
	}

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		// Collision detected
		x = float64(line1.StartX) + (t * float64(s1X))
		y = float64(line1.StartY) + (t * float64(s1Y))
		return x, y, true
	}

	return 0, 0, false // No collision
}

func LineIntersection(slope1, intercept1, slope2, intercept2 float64) (bool, float64, float64) {
	if slope1 == slope2 {
		if intercept1 == intercept2 {
			return true, 0, intercept1 // The lines are the same
		}
		return false, 0, 0 // The lines are parallel
	}
	x := (intercept2 - intercept1) / (slope1 - slope2)
	y := slope1*x + intercept1
	return true, x, y
}
