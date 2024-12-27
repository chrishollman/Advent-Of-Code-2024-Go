package dayTwo

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2024-Go/utils"
)

//go:embed input.txt
var file string

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)
	safe := 0
	for _, line := range lines {
		strNums := strings.Fields(line)
		increasing, decreasing, valid := true, true, true

		for i := 1; i < len(strNums); i++ {
			curr, _ := strconv.Atoi(strNums[i])
			prev, _ := strconv.Atoi(strNums[i-1])

			if curr < prev {
				decreasing = false
			} else {
				increasing = false
			}

			if increasing == decreasing {
				valid = false
				break
			}

			diff := utils.AbsInt(curr - prev)
			if diff < 1 || diff > 3 {
				valid = false
				break
			}
		}

		if valid {
			safe++
		}
	}

	return safe
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)
	safe := 0
	for _, line := range lines {
		strNums := strings.Fields(line)
		nums := make([]int, len(strNums))
		for i, strNum := range strNums {
			nums[i], _ = strconv.Atoi(strNum)
		}

		numCombinations := utils.GenerateCombinations(nums)
		for _, combination := range numCombinations {
			if isValidCombination(combination) {
				safe++
				break
			}
		}
	}

	return safe
}

func isValidCombination(combination []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(combination); i++ {
		curr, prev := combination[i], combination[i-1]
		if curr < prev {
			decreasing = false
		} else {
			increasing = false
		}
		if increasing == decreasing {
			return false
		}
		diff := utils.AbsInt(curr - prev)
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Day Two - Red-Nosed Reports")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
