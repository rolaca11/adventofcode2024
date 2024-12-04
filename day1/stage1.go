package day1

import (
	"math"
	"sort"
)

func stage1(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	result := 0

	for i := 0; i < len(list1); i++ {
		result += int(math.Abs(float64(list2[i] - list1[i])))
	}

	return result
}
