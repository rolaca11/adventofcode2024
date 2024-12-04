package day2

func stage2(input []Record) int {
	var result = 0

	for _, record := range input {
		for index, _ := range record {
			r := remove(record, index)
			if checkCorrectness(r) < 0 {
				result++
				break
			}
		}
	}

	return result
}

func remove(record Record, i int) Record {
	var s = make(Record, 0)
	return append(append(s, record[:i]...), record[i+1:]...)
}
