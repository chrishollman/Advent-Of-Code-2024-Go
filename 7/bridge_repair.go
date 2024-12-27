package daySeven

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
	return solve(input, false)
}

func PartTwo(input string) int {
	return solve(input, true)
}

func parseInput(line string) (int, []int) {
	lineSplit := strings.Split(line, ":")
	target, _ := strconv.Atoi(lineSplit[0])

	var nums []int
	for _, strnum := range strings.Fields(strings.TrimSpace(lineSplit[1])) {
		num, _ := strconv.Atoi(strnum)
		nums = append(nums, num)
	}
	return target, nums
}

func solve(input string, isPart2 bool) int {
	lines := utils.EmbedStringToStringSlice(input)
	total := 0

	for _, line := range lines {
		target, nums := parseInput(line)
		if canReach(nums, target, isPart2) {
			total += target
		}
	}

	return total
}

func canReach(numbers []int, target int, isPart2 bool) bool {
	if len(numbers) == 0 {
		return target == 0
	}

	dp := make([]map[int]bool, len(numbers))
	dp[0] = make(map[int]bool)
	dp[0][numbers[0]] = true

	for i := 1; i < len(numbers); i++ {
		dp[i] = make(map[int]bool)
		for val := range dp[i-1] {
			dp[i][val+numbers[i]] = true
			dp[i][val*numbers[i]] = true

			if isPart2 {
				left := strconv.Itoa(val)
				right := strconv.Itoa(numbers[i])
				concat, _ := strconv.Atoi(left + right)
				dp[i][concat] = true
			}
		}
	}

	return dp[len(numbers)-1][target]
}

func main() {
	fmt.Println("Day Seven - Bridge Repair")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
