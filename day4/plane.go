package day4

import (
	"errors"
	"strings"
)

const left = 1
const right = 5
const top = 2
const bottom = 6
const topLeft = 3
const bottomRight = 7
const topRight = 4
const bottomLeft = 8

type Node struct {
	char     int32
	left     *Node
	top      *Node
	topLeft  *Node
	topRight *Node
}

type Plane []Line

type Line []Node

func NewPlane(input string) Plane {
	result := make(Plane, 0)

	for y, row := range strings.Split(input, "\n") {
		resultRow := make(Line, 0)
		for x, char := range row {
			var left, top, topLeft, topRight *Node = nil, nil, nil, nil

			if x > 0 {
				left = &resultRow[x-1]
				if y > 0 {
					topLeft = &result[y-1][x-1]
				}
			}

			if y > 0 {
				top = &result[y-1][x]
				if x < len(result[y-1])-1 {
					topRight = &result[y-1][x+1]
				}
			}

			resultRow = append(resultRow, Node{
				char:     char,
				left:     left,
				top:      top,
				topLeft:  topLeft,
				topRight: topRight,
			})
		}
		result = append(result, resultRow)
	}

	return result
}

func (p Plane) Line(x int, y int, direction int) (Line, error) {
	result := make(Line, 0)

	if x >= 0 && y >= 0 && len(p) > y && len(p[y]) > x {
		result = append(result, p[y][x])

		nextX := x
		nextY := y

		if direction == top {
			nextY--
		} else if direction == topLeft {
			nextX--
			nextY--
		} else if direction == topRight {
			nextX++
			nextY--
		} else if direction == left {
			nextX--
		} else if direction == bottom {
			nextY++
		} else if direction == right {
			nextX++
		} else if direction == bottomLeft {
			nextX--
			nextY++
		} else if direction == bottomRight {
			nextX++
			nextY++
		}

		next, err := p.Line(nextX, nextY, direction)
		if err == nil {
			result = append(result, next...)
		}
	} else {
		return nil, errors.New("out of bounds")
	}

	return result, nil
}

func (l Line) String() string {
	s := ""

	for _, node := range l {
		s = s + string(node.char)
	}

	return s
}
