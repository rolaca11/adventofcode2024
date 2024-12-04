package day2

import (
	"os"
	"strings"
	"testing"
)

func TestStage1Example(t *testing.T) {
	var file, _ = os.ReadFile("./stage1_example")

	var records []Record
	for _, line := range strings.Split(string(file), "\n") {
		records = append(records, NewRecord(line))
	}

	result := stage1(records)
	if result != 2 {
		t.Error(result)
	}
}

func TestStage1(t *testing.T) {
	var file, _ = os.ReadFile("./stage1")

	var records []Record
	for _, line := range strings.Split(string(file), "\n") {
		records = append(records, NewRecord(line))
	}

	result := stage1(records)
	if result != 479 {
		t.Error(result)
	}
}

func TestStage2Example(t *testing.T) {
	var file, _ = os.ReadFile("./stage2_example")

	var records []Record
	for _, line := range strings.Split(string(file), "\n") {
		records = append(records, NewRecord(line))
	}

	result := stage2(records)
	if result != 4 {
		t.Error(result)
	}
}

func TestStage2(t *testing.T) {
	var file, _ = os.ReadFile("./stage2")

	var records []Record
	for _, line := range strings.Split(string(file), "\n") {
		records = append(records, NewRecord(line))
	}

	result := stage2(records)
	if result != 479 {
		t.Error(result)
	}
}

func TestRemove(t *testing.T) {
	record := NewRecord("10 11 12")
	remove(record, 1)
	t.Error(record)

}
