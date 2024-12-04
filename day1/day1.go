package day1

import (
	"strconv"
	"strings"
)

func readInput(input string) ([]int, []int) {
	var list1, list2 []int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}
