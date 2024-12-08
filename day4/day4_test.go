package day4

import (
	"os"
	"testing"
)

func TestStage1Example(t *testing.T) {
	file, _ := os.ReadFile("./stage1_example")

	result := stage1(NewPlane(string(file)))
	if result != 18 {
		t.Error(result)
	}
}

func TestStage1(t *testing.T) {
	file, _ := os.ReadFile("./stage1")

	result := stage1(NewPlane(string(file)))
	if result != 2524 {
		t.Error(result)
	}
}

func TestStage2Example(t *testing.T) {
	file, _ := os.ReadFile("./stage1_example")

	result := stage2(NewPlane(string(file)))
	if result != 9 {
		t.Error(result)
	}
}

func TestStage2(t *testing.T) {
	file, _ := os.ReadFile("./stage1")

	result := stage2(NewPlane(string(file)))
	if result != 1873 {
		t.Error(result)
	}
}
