package main

import (
	"testing"
)

func TestGetNumSafeReportExample(t *testing.T) {
	result := GetNumOfSafeReports("example.txt")
	if result != 2 {
		t.Errorf("GetNumOfSafeReports(\"example.txt\") = %d; expected 2", result)
	}
}

func TestGetNumSafeReportInput(t *testing.T) {
	result := GetNumOfSafeReports("input.txt")
	if result != 524 {
		t.Errorf("GetNumOfSafeReports(\"input.txt\") = %d; expected 1000", result)
	}
}
