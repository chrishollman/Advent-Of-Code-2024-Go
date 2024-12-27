package utils

import "strings"

func EmbedStringToStringSlice(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n")
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GenerateCombinations[T any](input []T) [][]T {
	n := len(input)
	if n <= 1 {
		return nil
	}

	result := make([][]T, n)

	for i := 0; i < n; i++ {
		combo := make([]T, 0, n-1)
		for j := 0; j < n; j++ {
			if j != i {
				combo = append(combo, input[j])
			}
		}
		result[i] = combo
	}

	return result
}
