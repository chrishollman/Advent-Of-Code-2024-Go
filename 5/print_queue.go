package dayFive

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var file string

func PartOne(input string) int {
	rawRules, rawUpdates := reader(input)
	rules := map[int][]int{}
	total := 0

	for _, rawRule := range rawRules {
		target, condition := rawRule[0], rawRule[1]
		if _, exists := rules[target]; exists {
			rules[target] = append(rules[target], condition)
		} else {
			rules[target] = []int{condition}
		}
	}

	for _, rawUpdate := range rawUpdates {
		sorted := santaSort(rules, rawUpdate)
		if slices.Equal(sorted, rawUpdate) {
			total += sorted[len(sorted)/2]
		}
	}

	return total
}

func PartTwo(input string) int {
	rawRules, rawUpdates := reader(input)
	rules := map[int][]int{}
	total := 0

	for _, rawRule := range rawRules {
		target, condition := rawRule[0], rawRule[1]
		if _, exists := rules[target]; exists {
			rules[target] = append(rules[target], condition)
		} else {
			rules[target] = []int{condition}
		}
	}

	for _, rawUpdate := range rawUpdates {
		sorted := santaSort(rules, rawUpdate)
		fmt.Println(rawUpdate, sorted)
		if !slices.Equal(sorted, rawUpdate) {
			fmt.Printf("Adding %d\n", sorted[len(sorted)/2])
			total += sorted[len(sorted)/2]
		}
	}

	return total
}

func reader(input string) ([][]int, [][]int) {
	lines := strings.Split(input, "\n")
	before := make([][]int, 0)
	after := make([][]int, 0)

	blank := false
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			blank = true
			continue
		}

		if !blank {
			split := strings.Split(line, "|")
			n1, _ := strconv.Atoi(split[0])
			n2, _ := strconv.Atoi(split[1])
			before = append(before, []int{n1, n2})
		} else {
			split := strings.Split(line, ",")
			tmp := make([]int, 0)
			for _, numString := range split {
				n, _ := strconv.Atoi(numString)
				tmp = append(tmp, n)
			}
			after = append(after, tmp)
		}
	}

	return before, after
}

func santaSort(rules map[int][]int, target []int) []int {
	result := make([]int, len(target))
	copy(result, target)

	n := len(result)
	for i := 0; i < n-1; i++ {
		values, _ := rules[result[i]]
		for j := i + 1; j < n; j++ {
			if !slices.Contains(values, result[j]) {
				result[i], result[j] = result[j], result[i]
				i = -1
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println("Day Five - Print Queue")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(file))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(file))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
