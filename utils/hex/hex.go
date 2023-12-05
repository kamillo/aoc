package hex

import (
	"fmt"
	"math"
)

type Direction int

const (
	DirectionSE = iota
	DirectionE
	DirectionNE
	DirectionNW
	DirectionW
	DirectionSW
)

var Directions = []Hex{
	NewHex(0, 1),
	NewHex(1, 0),
	NewHex(1, -1),
	NewHex(0, -1),
	NewHex(-1, 0),
	NewHex(-1, 1),
}

type Hex struct {
	q int // x axis
	r int // y axis
	s int // z axis
}

func NewHex(q, r int) Hex {
	h := Hex{q: q, r: r, s: -q - r}
	return h
}

func (h Hex) String() string {
	return fmt.Sprintf("(%d,%d)", h.q, h.r)
}

// Adds two hexagons
func Add(a, b Hex) Hex {
	return NewHex(a.q+b.q, a.r+b.r)
}

// Subtracts two hexagons
func Subtract(a, b Hex) Hex {
	return NewHex(a.q-b.q, a.r-b.r)
}

func Length(hex Hex) int {
	return int((math.Abs(float64(hex.q)) + math.Abs(float64(hex.r)) + math.Abs(float64(hex.q)+float64(hex.r))) / 2.)
}

func Distance(a, b Hex) int {
	sub := Subtract(a, b)
	return Length(sub)
}

// Returns the neighbor hexagon at a certain direction
func Neighbor(h Hex, direction Direction) Hex {
	directionOffset := Directions[direction]
	return NewHex(h.q+directionOffset.q, h.r+directionOffset.r)
}

// Returns all (6) neighbors
func Neighbors(h Hex) (neighbors [6]Hex) {
	for i, direction := range Directions {
		neighbors[i] = NewHex(h.q+direction.q, h.r+direction.r)
	}

	return neighbors
}

// Returns the set of hexagons around a certain center for a given radius
//func HexRange(center Hex, radius int) []Hex {
//
//	var results = make([]Hex, 0)
//
//	if radius >= 0 {
//		for dx := -radius; dx <= radius; dx++ {
//
//			for dy := math.Max(float64(-radius), float64(-dx-radius)); dy <= math.Min(float64(radius), float64(-dx+radius)); dy++ {
//				results = append(results, Add(center, NewHex(int(dx), int(dy))))
//			}
//		}
//	}
//
//	return results
//}

// Returns the set of hexagons that form a rectangle with the specified width and height
//func HexRectangleGrid(width, height int) []Hex {
//
//	results := make([]Hex, 0)
//
//	for q := 0; q < width; q++ {
//		qOffset := int(math.Floor(float64(q) / 2.))
//
//		for r := -qOffset; r < height-qOffset; r++ {
//
//			results = append(results, NewHex(q, r))
//		}
//	}
//
//	return results
//}
