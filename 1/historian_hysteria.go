package dayOne

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2024-Go/utils"
)

//go:embed input.txt
var file string

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for _, line := range lines {
		split := strings.Fields(line)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	res := 0
	for i := 0; i < len(left); i++ {
		res += utils.AbsInt(left[i] - right[i])
	}

	return res
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)
	left := make([]int, 0, len(lines))
	right := make(map[int]int, len(lines))

	for _, line := range lines {
		split := strings.Fields(line)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])

		left = append(left, l)
		right[r]++
	}

	res := 0
	for _, v := range left {
		res += v * right[v]
	}

	return res
}

func main() {
	fmt.Println("Day One - Historian Hysteria")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
