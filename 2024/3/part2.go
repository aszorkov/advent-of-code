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

	re_valid_blocks := regexp.MustCompile(`(?sm)(?:\A|do\(\)).*?(?:don't\(\)|\z)`)
	re_valid_calls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	re_digits := regexp.MustCompile(`\d+`)

	valid_blocks := re_valid_blocks.FindAllString(s, -1)
	sum := 0
	
	for i, block := range valid_blocks {
		fmt.Println("Block", i + 1, ":", block)
		valid_calls := re_valid_calls.FindAllString(block, -1)
		for _, call := range valid_calls {
			digits := re_digits.FindAllString(call, 2)
			x, _ := strconv.Atoi(digits[0])
			y, _ := strconv.Atoi(digits[1])
			result := x * y
			fmt.Println("Capture group:", call, "Pair: ", digits, ", Result:", result)
			sum += result
		}
	}

	fmt.Println("Answer: ", sum)
}