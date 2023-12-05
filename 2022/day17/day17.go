package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

const (
	boardWidth = 7
)

var shapes = [][][]int{
	{{1, 1, 1, 1}},
	{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
	{{0, 0, 1}, {0, 0, 1}, {1, 1, 1}},
	{{1}, {1}, {1}, {1}},
	{{1, 1}, {1, 1}},
}

type tetromino struct {
	x, y  int
	shape [][]int
}

func (t *tetromino) moveLeft(board *Board) {
	if t.x > 0 && canMove(tetromino{shape: t.shape, x: t.x - 1, y: t.y}, board) {
		t.x--
	}
}

func (t *tetromino) moveRight(board *Board) {
	if t.x < boardWidth-len(t.shape[0]) && canMove(tetromino{shape: t.shape, x: t.x + 1, y: t.y}, board) {
		t.x++
	}
}

func canMove(t tetromino, board *Board) bool {
	for i, row := range t.shape {
		for j, cell := range row {
			if cell == 1 && board[t.y+i][t.x+j] == 1 {
				return false
			}
		}
	}
	return true
}

func (t *tetromino) lock(board *Board) {
	for i, row := range t.shape {
		for j, cell := range row {
			if cell == 1 {
				board[t.y+i][t.x+j] = cell
			}
		}
	}

	if t.y < highest {
		highest = t.y
	}

	// for i := highest; i < len(board); i++ {
	// 	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(utils.SliceToString(board[i][:], ""), "1", "#"), "0", "."))
	// }
	// fmt.Println()
}

type Board [5 * 2022][7]int

type game struct {
	board            Board
	currentTetromino *tetromino
}

var highest = 0

func (g *game) update(i int) bool {
	moved := g.currentTetromino.moveDown(g)
	if !moved {
		g.currentTetromino.lock(&g.board)
		// g.checkCompletedLines()
		g.currentTetromino = generateTetromino(i + 1)
	}

	return moved
}

func generateTetromino(i int) *tetromino {
	shape := shapes[i%len(shapes)]

	return &tetromino{
		shape: shape,
		x:     2,
		y:     highest - len(shape) - 3,
	}
}

// func (g *game) checkCompletedLines() {
// 	// Check for completed lines and remove them.
// 	for i := range g.board {
// 		lineComplete := true
// 		for _, cell := range g.board[i] {
// 			if cell == 0 {
// 				lineComplete = false
// 				break
// 			}
// 		}
// 		if lineComplete {
// 			// Remove the completed line and move all lines above it down.
// 			g.board = append(g.board[:i], g.board[i+1:]...)
// 			g.board = append([][]int{{0, 0, 0, 0, 0, 0, 0}}, g.board...)
// 		}
// 	}
// }

func (t *tetromino) collides(board Board) bool {
	if t.y+len(t.shape) >= len(board) {
		return true
	}
	for i, row := range t.shape {
		for j, cell := range row {
			if cell == 1 && board[t.y+1+i][t.x+j] == 1 {
				return true
			}
		}
	}
	return false
}

func (t *tetromino) moveDown(g *game) bool {
	if !t.collides(g.board) {
		t.y++
		return true
	}
	return false
}

type gameState struct {
	lastLine string
	shape    int
	jet      int
}

func main() {
	lines := utils.GetLines("input.txt")
	jets := strings.Split(lines[0], "")

	g := game{}
	highest = len(g.board)
	g.currentTetromino = generateTetromino(0)
	j := 0
	//state := map[gameState]bool{}

	for i := 0; i < 2022; i += 1 {
		for ; ; j = (j + 1) % len(jets) {
			jet := jets[j]

			switch jet {
			case "<":
				g.currentTetromino.moveLeft(&g.board)
			case ">":
				g.currentTetromino.moveRight(&g.board)
			}

			if !g.update(i) {
				j = (j + 1) % len(jets)
				break
			}
		}
	}
	for i := highest; i < len(g.board); i++ {
		fmt.Println(strings.ReplaceAll(strings.ReplaceAll(utils.SliceToString(g.board[i][:], ""), "1", "#"), "0", "."))
	}
	fmt.Println(len(g.board) - highest)
}
