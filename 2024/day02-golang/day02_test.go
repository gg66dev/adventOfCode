package main

import (
	"testing"
)

/*
	func TestGetNumSafeReportExample(t *testing.T) {
		result := GetNumOfSafeReports("example.txt")
		if result != 2 {
			t.Errorf("GetNumOfSafeReports(\"example.txt\") = %d; expected 2", result)
		}
	}

	func TestGetNumSafeReportDampenerExample(t *testing.T) {
		result := GetNumOfSafeReportsDampener("example.txt")
		if result != 4 {
			t.Errorf("GetNumOfSafeReportsDampener(\"example.txt\") = %d; expected 4", result)
		}
	}
*/
func TestGetNumSafeReportInput(t *testing.T) {
	result := GetNumOfSafeReports("input.txt")
	if result != 524 {
		t.Errorf("GetNumOfSafeReports(\"input.txt\") = %d; expected 524", result)
	}
}

func TestGetNumSafeReportDampenerInput(t *testing.T) {
	result := GetNumOfSafeReportsDampener("input.txt")
	if result != 569 {
		t.Errorf("GetNumOfSafeReportsDampener(\"input.txt\") = %d; expected 569", result)
	}
}

// bad answer 656 (too high)
// bad answer 627 (too high)
// bad answer 593
// bad answer 597
// bad answer 552
// bad answer 544
// bad answer 540
// bad answer 539
// bad answer 542
// bad answer 543
// bad answer 574
// bad answer 568
// bad answer 560
// bad answer 564
// good answer 569

/*
func TestGetNumSafeReportDampenerInput3(t *testing.T) {
	result := GetNumOfSafeReportsDampener("input3.txt")
	if result != 5 {
		t.Errorf("GetNumOfSafeReportsDampener(\"input3.txt\") = %d; expected 5", result)
	}
}

func TestGetNumSafeReportDampenerInput4(t *testing.T) {
	result := GetNumOfSafeReportsDampener("input4.txt")
	if result != 5 {
		t.Errorf("GetNumOfSafeReportsDampener(\"input4.txt\") = %d; expected 5", result)
	}
}
*/
