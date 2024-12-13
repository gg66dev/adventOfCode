package main

import (
	"bufio"
	//"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// --- Day 2: Red-Nosed Reports ---
// Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.

// While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.

// They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.

// The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9
// This example data contains six reports each containing five levels.

// The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
// In the example above, the reports can be found safe or unsafe by checking those rules:

// 7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
// 1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
// 9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
// 1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
// 8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
// 1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
// So, in this example, 2 reports are safe.

// Analyze the unusual data from the engineers. How many reports are safe?

// To begin, get your puzzle input.

func eqSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func GetNumOfSafeReports(filename string) int {
	// structure to store input
	var reportData [][]string

	// read file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	var line string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		result := strings.Split(line, " ")
		//		fmt.Println("Result:", result)
		reportData = append(reportData, result)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	safeReports := 0
	for /*i*/ _, s1 := range reportData {
		//		fmt.Printf("Start: i=%v s=%v\n", i, s1)
		s2 := make([]string, len(s1))
		_ = copy(s2, s1)
		s3 := make([]string, len(s1))
		_ = copy(s3, s1)
		sort.Slice(s2, func(i, j int) bool {
			ii, _ := strconv.Atoi(s2[i])
			ij, _ := strconv.Atoi(s2[j])
			return ii < ij
		})
		sort.Slice(s3, func(i, j int) bool {
			ii, _ := strconv.Atoi(s3[i])
			ij, _ := strconv.Atoi(s3[j])
			return ii > ij
		})
		if !eqSlice(s1, s2) {
			if !eqSlice(s1, s3) {
				continue
			}
		}
		ok := true
		for idx, _ := range s1 {
			if idx == len(s1)-1 {
				continue
			}
			left, _ := strconv.ParseFloat(s1[idx], 64)
			right, _ := strconv.ParseFloat(s1[idx+1], 64)
			//log.Printf("%v %v", left, right)
			if math.Abs(left-right) > 3 || math.Abs(left-right) == 0 {
				ok = false
				break
			}

		}
		if ok {
			safeReports++
		}
	}
	return safeReports
}
