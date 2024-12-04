package day1

import (
	"fmt"
	"os"
	"testing"
)

func TestStage1Example(t *testing.T) {
	file, err := os.ReadFile("./stage1_example")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read file: %s", err))
	}

	result := stage1(readInput(string(file)))
	if result != 11 {
		t.Error(result)
	}
}

func TestStage1(t *testing.T) {
	file, err := os.ReadFile("./stage1")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read file: %s", err))
	}

	result := stage1(readInput(string(file)))
	if result != 3569916 {
		t.Error(result)
	}
}

func TestStage2Example(t *testing.T) {
	file, err := os.ReadFile("./stage1_example")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read file: %s", err))
	}

	result := stage2(readInput(string(file)))
	if result != 31 {
		t.Error(result)
	}
}

func TestStage2(t *testing.T) {
	file, err := os.ReadFile("./stage1")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read file: %s", err))
	}

	result := stage2(readInput(string(file)))
	if result != 26407426 {
		t.Error(result)
	}
}
