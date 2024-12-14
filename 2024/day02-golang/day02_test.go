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
		t.Errorf("GetNumOfSafeReports(\"input.txt\") = %d; expected 524", result)
	}
}

func TestGetNumSafeReportDampenerExample(t *testing.T) {
	result := GetNumOfSafeReportsDampener("example.txt")
	if result != 4 {
		t.Errorf("GetNumOfSafeReportsDampener(\"example.txt\") = %d; expected 4", result)
	}
}

func TestGetNumSafeReportDampenerSlice0(t *testing.T) {
	var s1 []string
	var result int
	s1 = []string{"7", "6", "4", "2", "1"}
	result = GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 0 - 1 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
	s1 = []string{"1", "2", "7", "8", "9"}
	result = GetNumOfSafeReportDampenerSlice(s1)
	if result != 0 {
		t.Errorf("test 0 - 2 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}

	s1 = []string{"9", "7", "6", "2", "1"}
	result = GetNumOfSafeReportDampenerSlice(s1)

	if result != 0 {
		t.Errorf("test 0 - 3 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}
	s1 = []string{"1", "3", "2", "4", "5"}
	result = GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 0 - 4 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}

	s1 = []string{"8", "6", "4", "4", "1"}
	result = GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 0 - 5 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
	s1 = []string{"1", "3", "6", "7", "9"}
	result = GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 0 - 6 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}

}

func TestGetNumSafeReportDampenerInput(t *testing.T) {
	result := GetNumOfSafeReportsDampener("input.txt")
	if result != 1000 {
		t.Errorf("GetNumOfSafeReportsDampener(\"input.txt\") = %d; expected 1000", result)
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

func TestGetNumSafeReportDampenerSlice1(t *testing.T) {
	// s1: [76 75 75 73 72 69 71]

	s1 := []string{"76", "75", "75", "73", "72", "69", "71"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 0 {
		t.Errorf("test 1 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}
}

func TestGetNumSafeReportDampenerSlice2(t *testing.T) {
	// s1: [45 45 47 47 48]

	s1 := []string{"45", "45", "47", "47", "48"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 0 {
		t.Errorf("test 2 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}
}

func TestGetNumSafeReportDampenerSlice3(t *testing.T) {
	// s1: [46 42 45 42 45]
	s1 := []string{"46", "42", "45", "42", "45"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 0 {
		t.Errorf("test 3 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}
}

func TestGetNumSafeReportDampenerSlice4(t *testing.T) {
	// s1: [95 99 97 99 96]
	s1 := []string{"95", "99", "97", "99", "96"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 0 {
		t.Errorf("test 4 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}
}

func TestGetNumSafeReportDampenerSlice5(t *testing.T) {
	// s1: [10 12 13 16 17]
	s1 := []string{"10", "12", "13", "16", "17"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 5 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
}

func TestGetNumSafeReportDampenerSlice6(t *testing.T) {
	// s1: [10 12 13 16 17]
	s1 := []string{"8", "6", "4", "4", "1"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 6 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
}

func TestGetNumSafeReportDampenerSlice7(t *testing.T) {
	// s1:  [77 78 84 85 88 85]
	s1 := []string{"77", "78", "84", "85", "88", "85"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 0 {
		t.Errorf("test 7 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 0", result)
	}
}
func TestGetNumSafeReportDampenerSlice8(t *testing.T) {
	s1 := []string{"1", "2", "3", "4", "5", "9"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 8 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
}

func TestGetNumSafeReportDampenerSlice9(t *testing.T) {
	s1 := []string{"9", "5", "4", "3", "2", "1"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 9 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
}

func TestGetNumSafeReportDampenerSlice10(t *testing.T) {
	s1 := []string{"54", "59", "57", "58", "59"}

	result := GetNumOfSafeReportDampenerSlice(s1)
	if result != 1 {
		t.Errorf("test 10 - GetNumOfSafeReportDampenerSlice(s1)  = %d; expected 1", result)
	}
}
