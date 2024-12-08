package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

type KeywordMatch struct {
	origin [2]int
	direction string
}

type Direction struct {
	x int
	y int
	desc string
}

func stringReverse(input string) (output string) {
	for _, c := range input {
		output = string(c) + output
	}
	return
}

func findKeyword(grid []string, x, y int, direction Direction, keyword string) (KeywordMatch, bool) {
	// Make sure we've been given valid x and y values
	if(x > -1 && y > -1 && y < len(grid) && x < len(grid[y])) {
		x_end := x + (len(keyword) - 1) * direction.x
		y_end := y + (len(keyword) - 1) * direction.y
		x_in_bound := x_end > -1 && x_end < len(grid[y])
		y_in_bound := y_end > -1 && y_end < len(grid)
		fmt.Println(len(keyword), x, x_end, x_in_bound, y, y_end, y_in_bound)
		if x_in_bound && y_in_bound {
			fmt.Println("Checking direction", direction.desc)
			s := ""
			for i := 0; i < len(keyword); i++ {
				row := grid[y + i * direction.y]
				char := row[x + i * direction.x]
				s += string(char)
			}
			fmt.Println("Got: ", s)
			keyword_reversed := stringReverse(keyword)
			if s == keyword || s == keyword_reversed {
				return KeywordMatch{[2]int{x,y}, direction.desc}, true
			}
		}
	}
	return KeywordMatch{}, false
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error! Expected one positional argument. arg[1]: filename")
	}
	filename := os.Args[1]

	fmt.Println("Reading file: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	grid := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	directions := []Direction{
		{ 1,-1, "up_right"},
		{ 1, 1, "down_right"},
	}

	sum := 0
	matches := []KeywordMatch{}
	for y, row := range grid {
		for x, _ := range row {
			if string(row[x]) == "A" {
				fmt.Println("Found 'A' at", x, y)
				match1, found1 := findKeyword(grid, x-1, y+1, directions[0], "MAS")
				match2, found2 := findKeyword(grid, x-1, y-1, directions[1], "MAS")
				if found1 && found2 {
					matches = append(matches, match1)
					matches = append(matches, match2)
					sum++
				}
			}
		}
	}

	fmt.Println("Answer: ", sum)
}