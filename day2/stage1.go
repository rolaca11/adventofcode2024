package day2

func stage1(input []Record) int {
	var result = 0

	for _, record := range input {
		if checkCorrectness(record) < 0 {
			result++
		}
	}

	return result
}
