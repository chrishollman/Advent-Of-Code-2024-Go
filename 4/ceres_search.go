package dayFour

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2024-Go/utils"
)

//go:embed input.txt
var file string

type Position struct {
	Row int
	Col int
}

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func PartOne(input string) int {
	grid := utils.EmbedStringToStringSlice(input)
	rows, cols := len(grid), len(grid[0])
	target := "XMAS"
	count := 0

	isValid := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	search := func(startRow, startCol int, dir Direction) bool {
		if !isValid(startRow+dir.dx*3, startCol+dir.dy*3) {
			return false
		}

		for i := 0; i < len(target); i++ {
			row, col := startRow+dir.dx*i, startCol+dir.dy*i
			if grid[row][col] != target[i] {
				return false
			}
		}

		return true
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 'X' {
				for _, dir := range directions {
					if search(row, col, dir) {
						count++
					}
				}
			}
		}
	}

	return count
}

func PartTwo(input string) int {
	grid := utils.EmbedStringToStringSlice(input)
	rows, cols := len(grid), len(grid[0])
	count := 0

	search := func(startRow, startCol int) bool {
		if ((grid[startRow-1][startCol+1] == 'M' && grid[startRow+1][startCol-1] == 'S') ||
			(grid[startRow-1][startCol+1] == 'S' && grid[startRow+1][startCol-1] == 'M')) &&
			((grid[startRow-1][startCol-1] == 'M' && grid[startRow+1][startCol+1] == 'S') ||
				(grid[startRow-1][startCol-1] == 'S' && grid[startRow+1][startCol+1] == 'M')) {
			return true
		}

		return false
	}

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if grid[row][col] == 'A' {
				if search(row, col) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Println("Day Four - Ceres Search")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
