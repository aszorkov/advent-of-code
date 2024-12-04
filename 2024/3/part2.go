package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error! Expected only one positional argument")
	}
	filename := os.Args[1]
	fmt.Println("Reading file: ", filename)
	b, err := os.ReadFile(filename) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

	s := string(b)

	re_valid_calls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re_valid_calls.FindAllString(s, -1)

	sum := 0
	re_digits := regexp.MustCompile(`\d+`)
	for _, val := range matches {
		digits := re_digits.FindAllString(val, 2)
		x, _ := strconv.Atoi(digits[0])
		y, _ := strconv.Atoi(digits[1])
		result := x * y
		fmt.Println("Pair: ", digits, ", Result:", result)
		sum += result
	}

	fmt.Println("Answer: ", sum)
}