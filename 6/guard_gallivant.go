package daySix

import (
	_ "embed"
	"fmt"
	"slices"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2024-Go/utils"
)

//go:embed input.txt
var file string

type Direction struct {
	dx, dy int
}

type Position struct {
	row, col int
}

type Grid [][]string

var directions = []Direction{
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
	{0, -1}, // Up
}

const (
	right = '>'
	down  = 'v'
	left  = '<'
	up    = '^'
)

func PartOne(input string) int {
	grid, start, startDir := parseGrid(input)
	visited := make(map[Position]struct{})

	traverseGrid(grid, start, startDir, visited)

	return len(visited)
}

func PartTwo(input string) int {
	grid, start, startDir := parseGrid(input)
	visited := make(map[Position]struct{})

	traverseGrid(grid, start, startDir, visited)

	loopCount := 0
	for pos := range visited {
		if isLoop(grid, start, startDir, pos) {
			loopCount++
		}
	}
	return loopCount
}

func parseGrid(input string) (Grid, Position, Direction) {
	lines := utils.EmbedStringToStringSlice(input)
	grid := make(Grid, len(lines))
	var start Position
	var startDir Direction

	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, ch := range line {
			grid[i][j] = string(ch)
			if slices.Contains([]rune{right, down, left, up}, ch) {
				start = Position{i, j}
				startDir = directionFromRune(ch)
			}
		}
	}
	return grid, start, startDir
}

func traverseGrid(grid Grid, pos Position, dir Direction, visited map[Position]struct{}) {
	maxRow, maxCol := len(grid), len(grid[0])

	for {
		next := Position{pos.row + dir.dy, pos.col + dir.dx}
		if !grid.moveIsValid(next, maxRow, maxCol) {
			break
		}

		for grid[next.row][next.col] == "#" {
			dir = rotateDirection(dir)
			next = Position{pos.row + dir.dy, pos.col + dir.dx}
			if !grid.moveIsValid(next, maxRow, maxCol) {
				return
			}
		}

		visited[next] = struct{}{}
		pos = next
	}
}

func isLoop(grid Grid, start Position, startDir Direction, blockPos Position) bool {
	maxRow, maxCol := len(grid), len(grid[0])
	pos := start
	dir := startDir
	seenTurns := make(map[string]struct{})

	origValue := grid[blockPos.row][blockPos.col]
	grid[blockPos.row][blockPos.col] = "#"
	defer func() { grid[blockPos.row][blockPos.col] = origValue }()

	for {
		next := Position{pos.row + dir.dy, pos.col + dir.dx}
		if !grid.moveIsValid(next, maxRow, maxCol) {
			return false
		}

		for grid[next.row][next.col] == "#" {
			turnKey := fmt.Sprintf("%d,%d-%d,%d", pos.row, pos.col, next.row, next.col)
			if _, exists := seenTurns[turnKey]; exists {
				return true
			}
			seenTurns[turnKey] = struct{}{}
			dir = rotateDirection(dir)
			next = Position{pos.row + dir.dy, pos.col + dir.dx}
			if !grid.moveIsValid(next, maxRow, maxCol) {
				return false
			}
		}
		pos = next
	}
}

func (g Grid) moveIsValid(pos Position, maxRow, maxCol int) bool {
	return pos.row >= 0 && pos.row < maxRow && pos.col >= 0 && pos.col < maxCol
}

func directionFromRune(r rune) Direction {
	switch r {
	case right:
		return directions[0]
	case down:
		return directions[1]
	case left:
		return directions[2]
	case up:
		return directions[3]
	}
	panic("invalid direction")
}

func rotateDirection(d Direction) Direction {
	for i, dir := range directions {
		if dir == d {
			return directions[(i+1)%4]
		}
	}
	panic("invalid direction")
}

func main() {
	fmt.Println("Day Six - Guard Gallivant")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
