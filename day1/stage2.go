package day1

func stage2(list1 []int, list2 []int) int {
	occurrences := countOccurrences(list2)

	result := 0

	for _, num := range list1 {
		result += occurrences[num] * num
	}

	return result
}

func countOccurrences(list []int) map[int]int {
	var result = make(map[int]int)

	for _, num := range list {
		result[num]++
	}

	return result
}
