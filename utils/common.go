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
