package dayThree

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:embed input.txt
var file string

func PartOne(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	res := 0

	matches := regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		lNum, _ := strconv.Atoi(match[1])
		rNum, _ := strconv.Atoi(match[2])
		res += lNum * rNum
	}

	return res
}

func PartTwo(input string) int {
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	intervalRegex := regexp.MustCompile(`do\(\)|don't\(\)`)

	multiplyStrings := func(text string) int {
		sum := 0
		for _, match := range mulRegex.FindAllStringSubmatch(text, -1) {
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])
			sum += l * r
		}
		return sum
	}

	result := 0
	startIdx := 0
	openInterval := true
	for _, interval := range intervalRegex.FindAllStringIndex(input, -1) {
		shouldCloseInterval := false
		if interval[1]-interval[0] == 7 {
			shouldCloseInterval = true
		}

		if openInterval && shouldCloseInterval {
			result += multiplyStrings(input[startIdx:interval[0]])
			openInterval = false
		} else if !openInterval && !shouldCloseInterval {
			startIdx = interval[1]
			openInterval = true
		}
	}

	if openInterval {
		result += multiplyStrings(input[startIdx:])
	}

	return result
}

func main() {
	fmt.Println("Day Three - Mull It Over")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
