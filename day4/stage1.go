package day4

import "strings"

func stage1(input Plane) int {
	result := 0

	for y, line := range input {
		for x, node := range line {
			if node.char == 'X' {
				if input.checkDirection(x, y, left) {
					result++
				}
				if input.checkDirection(x, y, right) {
					result++
				}
				if input.checkDirection(x, y, top) {
					result++
				}
				if input.checkDirection(x, y, bottom) {
					result++
				}
				if input.checkDirection(x, y, topLeft) {
					result++
				}
				if input.checkDirection(x, y, bottomRight) {
					result++
				}
				if input.checkDirection(x, y, topRight) {
					result++
				}
				if input.checkDirection(x, y, bottomLeft) {
					result++
				}
			}
		}
	}

	return result
}

func (p Plane) checkDirection(x int, y int, direction int) bool {
	line, err := p.Line(x, y, direction)
	if err != nil {
		return false
	}

	return strings.HasPrefix(line.String(), "XMAS")
}
