package day3

import (
	"os"
	"testing"
)

func TestStage1Example(t *testing.T) {
	file, _ := os.ReadFile("./stage1_example")

	result := stage1(string(file))
	if result != 161 {
		t.Error(result)
	}
}

func TestStage1(t *testing.T) {
	file, _ := os.ReadFile("./stage1")

	result := stage1(string(file))
	if result != 170068701 {
		t.Error(result)
	}
}

func TestStage2Example(t *testing.T) {
	file, _ := os.ReadFile("./stage2_example")

	result := stage2(string(file))
	if result != 48 {
		t.Error(result)
	}
}

func TestStage2(t *testing.T) {
	file, _ := os.ReadFile("./stage1")

	result := stage2(string(file))
	if result != 78683433 {
		t.Error(result)
	}
}
