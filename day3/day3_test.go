package day3

import (
	"os"
	"testing"
)

func TestStage1Example(t *testing.T) {
	file, _ := os.ReadFile("./stage1_example")

	result := stage1(file)
	if result != 161 {
		t.Error(result)
	}
}

func TestStage1(t *testing.T) {
	file, _ := os.ReadFile("./stage1")

	result := stage1(file)
	if result != 161 {
		t.Error(result)
	}
}
