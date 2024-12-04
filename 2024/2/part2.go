package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"log"
)

type Report struct {
	data []int 
	safe bool
	direction string
	removed []int
}

type ValidationResult struct {
	safe bool
	direction string
}

func validateData(data []int) ValidationResult {
	safe := true
	reportDirection := ""
	for i := 1; safe && i < len(data); i++ {
		diff := data[i] - data[i - 1]
		if diff == 0 || diff < -3 || diff > 3 {
			safe = false
		}
		var pairDirection = ""
		if diff > 0 {
			pairDirection = "ascending"
		}
		if diff < 0 {
			pairDirection = "descending"
		}
		if reportDirection == "" {
			reportDirection = pairDirection
		} else if reportDirection != pairDirection {
			safe = false
			reportDirection = ""
		}
	}
	return ValidationResult{safe, reportDirection}
}

// --- Expected args ---
// data (mandatory): An array of integers representing the report to validate
// tolerance (optional): An integer representing the number of levels we can ignore for a given report
// level (optional): An integer representing the current recursion level
func buildReport(data []int, args ...int) Report {
	report := Report{data, true, "", []int{}}
	if len(args) == 0 {
		args = append(args, make([]int, 2)...)
	}
	if len(args) == 1 {
		args = append(args, 0)
	}
	fmt.Println("Checking data: ", data)
	printBuffer := []int{}
	for i, complete := 0, false; !complete && i < len(data); i++ {
		printBuffer = append(printBuffer, i)
		result := validateData(data)
		report.safe = result.safe
		report.direction = result.direction
		if !result.safe && args[1] < args[0] {
			dataMinusIndex := make([]int, 0)
			dataMinusIndex = append(dataMinusIndex, data[:i]...)
			dataMinusIndex = append(dataMinusIndex, data[i + 1:]...)
			fmt.Println("Retrying on: ", dataMinusIndex, "after removing indice [", i, "]")
			reportMinusIndex := buildReport(dataMinusIndex, args[0], args[1] + 1)
			if reportMinusIndex.safe {
				report.removed = append(report.removed, i)
				report.direction = reportMinusIndex.direction
				report.safe = reportMinusIndex.safe
				complete = true
			}
		} else {
			complete = true
		}
	}
	fmt.Println("Checked on indices: ", printBuffer)
	return report
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error! Expected only one positional argument")
	}
	filename := os.Args[1]
	fmt.Println("Reading file: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safe := 0
	unsafe := 0
	reports := []Report{}

	for scanner.Scan() {
		line := scanner.Text()
		slicedStr := strings.Fields(line)
		data := make([]int, len(slicedStr))
		for index, val := range slicedStr {
			i, _ := strconv.Atoi(val)
			data[index] = i
		}
		report := buildReport(data, 1)
		if report.safe {
			safe++
		} else {
			unsafe++
		}
		reports = append(reports, report)
		fmt.Println("Report: ", report)
	}

	fmt.Println("Answer: ", safe)
}