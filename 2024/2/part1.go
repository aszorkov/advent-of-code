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
	direction string
	status string
}

func checkReport(input []string) Report {
	report := Report{[]int{}, "", "safe"}
	for index, val := range input {
		i, _ := strconv.Atoi(val)
		report.data = append(report.data, i)
		if(index > 0 && report.direction != "unsafe") {
			diff := i - report.data[index - 1]
			if diff == 0 || diff < -3 || diff > 3 {
				report.status = "unsafe"
			} else {
				var direction = "descending"
				if diff > 0 {
					direction = "ascending"
				}
				if report.direction == "" {
					report.direction = direction
				} else if report.direction != direction {
					report.status = "unsafe"
				}
			}
		}
	}
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
		report := checkReport(slicedStr)
		if report.status == "safe" {
			safe++
		} else if report.status == "unsafe" {
			unsafe++
		}
		reports = append(reports, report)
		fmt.Println(report)
	}

	fmt.Println("Answer: ", safe)
}