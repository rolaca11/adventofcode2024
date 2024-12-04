package day2

import (
	"math"
	"strconv"
	"strings"
)

type Record []Level

type Level int

func NewRecord(input string) Record {
	var levels []Level

	for _, slice := range strings.Split(input, " ") {
		number, _ := strconv.Atoi(slice)
		levels = append(levels, Level(number))
	}

	return levels
}

func checkCorrectness(record Record) int {
	var monotonicity Level = 0
	var previous = record[0]

	for i := 1; i < len(record); i++ {
		var level = record[i]
		var diff = level - previous

		if monotonicity*diff < 0 || diff == 0 {
			return i
		}

		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return i
		}

		monotonicity = Level(int(float64(diff) / math.Abs(float64(diff))))
		previous = level
	}

	return -1
}
