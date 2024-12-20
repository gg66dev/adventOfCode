package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	fmt.Printf("safe reports part I: %v\n", safeReports)
	return safeReports
}

// --- Part Two ---
// The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

// The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

// Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

// More of the above example's reports are now safe:

// 7 6 4 2 1: Safe without removing any level.
// 1 2 7 8 9: Unsafe regardless of which level is removed.
// 9 7 6 2 1: Unsafe regardless of which level is removed.
// 1 3 2 4 5: Safe by removing the second level, 3.
// 8 6 4 4 1: Safe by removing the third level, 4.
// 1 3 6 7 9: Safe without removing any level.
// Thanks to the Problem Dampener, 4 reports are actually safe!

// Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?

// reports no error in all levels

func GetNumOfSafeReportsDampener(filename string) int {
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

	//safeReports := 0
	var okDirect [][]string
	var twoDouble [][]string
	var badIncrease [][]string
	var badDecrease [][]string
	var badDirection [][]string
	for /*i*/ _, s1 := range reportData {
		processReports(
			s1,
			&okDirect,
			&twoDouble,
			&badIncrease,
			&badDecrease,
			&badDirection,
		)
	}
	var goodIncreaseIgnoring1Lvl [][]string
	var badIncreaseIgnoring1Lvl [][]string
	increaseIgnoring1Lvl(
		badIncrease,
		&goodIncreaseIgnoring1Lvl,
		&badIncreaseIgnoring1Lvl,
	)
	var goodDecreaseIgnoring1Lvl [][]string
	var badDecreaseIgnoring1Lvl [][]string
	decreaseIgnoring1Lvl(
		badDecrease,
		&goodDecreaseIgnoring1Lvl,
		&badDecreaseIgnoring1Lvl,
	)
	//fmt.Println("+++++++++ Print Ok Direct +++++++++++++")
	//for _, s1 := range okDirect {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Two Double +++++++++++++")
	//for _, s1 := range twoDouble {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Bad Increase +++++++++++++")
	//for _, s1 := range badIncrease {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Bad Decrease +++++++++++++")
	//for _, s1 := range badDecrease {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Bad Direction +++++++++++++")
	//for _, s1 := range badDirection {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Bad Increase ignoring 1 lvl +++++++++++++")
	//for _, s1 := range badIncreaseIgnoring1Lvl {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Good Increase ignoring 1 lvl +++++++++++++")
	//for _, s1 := range goodIncreaseIgnoring1Lvl {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Bad Decrease ignoring 1 lvl +++++++++++++")
	//for _, s1 := range badDecreaseIgnoring1Lvl {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")
	//fmt.Println("+++++++++ Print Good Decrease ignoring 1 lvl +++++++++++++")
	//for _, s1 := range goodDecreaseIgnoring1Lvl {
	//	fmt.Printf("%v\n", s1)
	//}
	//fmt.Println("+++++++++ +++++++++++++++ +++++++++++++")

	//fmt.Printf("total input: %v\n", len(reportData))
	//fmt.Printf("total process: %v\n", (len(okDirect) + len(twoDouble) + len(badIncrease) + len(badDecrease) + len(badDirection)))
	//fmt.Println("------------------------------------")
	//fmt.Printf("total directOK: %v\n", len(okDirect))
	//fmt.Printf("total twoDouble: %v\n", len(twoDouble))
	//fmt.Printf("total badDirection: %v\n", len(badDirection))
	//fmt.Println("------------------------------------")
	//fmt.Printf("total badIncrease: %v\n", len(badIncrease))
	//fmt.Printf("> total badIncreaseIgnoring1Lvl: %v\n", len(badIncreaseIgnoring1Lvl))
	//fmt.Printf("> total goodIncreaseIgnoring1Lvl: %v\n", len(goodIncreaseIgnoring1Lvl))
	//fmt.Println("------------------------------------")
	//fmt.Printf("total badDecrease: %v\n", len(badDecrease))
	//fmt.Printf("> total badDecreaseIgnoring1Lvl: %v\n", len(badDecreaseIgnoring1Lvl))
	//fmt.Printf("> total goodDecreaseIgnoring1Lvl: %v\n", len(goodDecreaseIgnoring1Lvl))
	//fmt.Println("------------------------------------")
	safeReports := len(okDirect) + len(goodIncreaseIgnoring1Lvl) + len(goodDecreaseIgnoring1Lvl)
	fmt.Printf("safe reports part II: %v\n", safeReports)
	return safeReports
}

func processReports(
	s1 []string,
	okDirect *[][]string,
	twoDouble *[][]string,
	badIncrease *[][]string,
	badDecrease *[][]string,
	badDirection *[][]string,
) {
	var vectorUnic []string
	var vectorAsc []bool
	var vectorDes []bool
	var vectorDiff []int
	for i := range s1 {
		if !slices.Contains(vectorUnic, s1[i]) {
			vectorUnic = append(vectorUnic, s1[i])
		}
		if i == len(s1)-1 {
			continue
		}
		left, _ := strconv.Atoi(s1[i])
		right, _ := strconv.Atoi(s1[i+1])
		vectorAsc = append(vectorAsc, (right > left && left+3 >= right))
		vectorDes = append(vectorDes, (left > right && left-3 <= right))
		vectorDiff = append(vectorDiff, (left - right))

	}
	if len(s1)-len(vectorUnic) > 1 {
		*twoDouble = append(*twoDouble, s1)
		return
	}
	isDsc := false
	numOfNeg := 0
	numOfPos := 0
	numOfZeros := 0
	for _, e := range vectorDiff {
		if e > 0 {
			numOfPos++
		} else if e < 0 {
			numOfNeg++
		} else {
			numOfZeros++
		}
	}
	if numOfPos == numOfNeg {
		*badDirection = append(*badDirection, s1)
		return
	}

	if numOfPos > 2 {
		isDsc = true
	}
	var errIdx []int
	if isDsc {
		for idx, b := range vectorDes {
			if !b {
				errIdx = append(errIdx, idx)
			}
		}
		if len(errIdx) == 0 {
			*okDirect = append(*okDirect, s1)
			return
		} else {
			*badDecrease = append(*badDecrease, s1)
		}
	} else {
		for idx, b := range vectorAsc {
			if !b {
				errIdx = append(errIdx, idx)
			}
		}
		if len(errIdx) == 0 {
			*okDirect = append(*okDirect, s1)
			return
		} else {
			*badIncrease = append(*badIncrease, s1)
		}

	}

}

func increaseIgnoring1Lvl(
	badIncrease [][]string,
	goodIncreaseIgnoring1Lvl *[][]string,
	badIncreaseIgnoring1Lvl *[][]string,
) {
	for _, s2 := range badIncrease {
		numBads := 0
		for i := 0; i < len(s2); i++ {
			if i == len(s2)-1 {
				continue
			}
			left, _ := strconv.Atoi(s2[i])
			right, _ := strconv.Atoi(s2[i+1])
			if right == left || right < left || left+3 < right {
				numBads++
				if i == len(s2)-2 {
					break
				}
				nextRight, _ := strconv.Atoi(s2[i+2])
				if nextRight == left || nextRight < left || left+3 < nextRight {
					if i != 0 {
						numBads++
					}
				} else {
					i++
				}

			}
		}
		if numBads == 1 {
			*goodIncreaseIgnoring1Lvl = append(*goodIncreaseIgnoring1Lvl, s2)
		} else {
			*badIncreaseIgnoring1Lvl = append(*badIncreaseIgnoring1Lvl, s2)
		}
	}

}

func decreaseIgnoring1Lvl(
	badDecrease [][]string,
	goodDecreaseIgnoring1Lvl *[][]string,
	badDecreaseIgnoring1Lvl *[][]string,
) {
	for _, s2 := range badDecrease {
		numBads := 0
		for i := 0; i < len(s2); i++ {
			if i == len(s2)-1 {
				continue
			}
			left, _ := strconv.Atoi(s2[i])
			right, _ := strconv.Atoi(s2[i+1])
			if right == left || right > left || left-3 > right {
				numBads++
				if i == len(s2)-2 {
					break
				}
				nextRight, _ := strconv.Atoi(s2[i+2])
				if nextRight == left || nextRight > left || left-3 > nextRight {
					if i != 0 {
						numBads++
					}
				} else {
					i++
				}
			}
		}
		if numBads == 1 {
			*goodDecreaseIgnoring1Lvl = append(*goodDecreaseIgnoring1Lvl, s2)
		} else {
			*badDecreaseIgnoring1Lvl = append(*badDecreaseIgnoring1Lvl, s2)
		}
	}

}
