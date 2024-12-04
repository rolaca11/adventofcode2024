package day3

import (
	"regexp"
	"strconv"
)

func stage1(input []byte) int {
	pattern, _ := regexp.Compile("mul\\((\\d+),(\\d+)\\)")

	matches := pattern.FindAllStringSubmatch(string(input), -1)
	accumulator := 0

	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		accumulator += first * second
	}

	return accumulator
}
