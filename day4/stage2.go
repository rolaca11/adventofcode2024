package day4

import "strings"

func stage2(input Plane) int {
	result := 0

	for y, line := range input {
		for x, node := range line {
			if node.char == 'A' {
				backSlash, err := input.Line(x-1, y-1, bottomRight)
				if err != nil {
					continue
				}
				slash, err := input.Line(x+1, y-1, bottomLeft)
				if err != nil {
					continue
				}

				if (strings.HasPrefix(backSlash.String(), "MAS") || strings.HasPrefix(backSlash.String(), "SAM")) &&
					(strings.HasPrefix(slash.String(), "MAS") || strings.HasPrefix(slash.String(), "SAM")) {
					result++
				}
			}
		}
	}

	return result
}
