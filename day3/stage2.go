package day3

import (
	"strings"
)

func stage2(input string) int {
	processedInput := strings.ReplaceAll(input, "\n", "")
	processedInput = strings.ReplaceAll(processedInput, "don't()", "\ndon't()")
	processedInput = strings.ReplaceAll(processedInput, "do()", "\ndo()")

	filteredLines := ""
	for _, line := range strings.Split(processedInput, "\n") {
		if !strings.HasPrefix(line, "don't()") {
			filteredLines += line + "\n"
		}
	}

	return stage1(filteredLines)
}
