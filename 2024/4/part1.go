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

func findKeyword(grid []string, x, y int, direction Direction, keyword string) (KeywordMatch, bool) {
	// Check if travelling in this direction to the end of the keyword would be within the bounds of the grid
	x_end := x + (len(keyword) - 1) * direction.x
	y_end := y + (len(keyword) - 1) * direction.y
	x_in_bound := x_end > -1 && x_end < len(grid[y])
	y_in_bound := y_end > -1 && y_end < len(grid)
	fmt.Println(len(keyword), x, x_end, x_in_bound, y, y_end, y_in_bound)
	if x_in_bound && y_in_bound {
		fmt.Println("Checking direction", direction.desc)
		matching := true
		for i := 1; i < len(keyword) && matching; i++ {
			row := grid[y + i * direction.y]
			char := row[x + i * direction.x]
			if char == keyword[i] {
				if i == len(keyword) - 1 {
					fmt.Println("Match found", direction.desc)
					return KeywordMatch{[2]int{x,y}, direction.desc}, true
				}
			} else {
				matching = false
			}
		}
	}
	return KeywordMatch{}, false
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Error! Expected two positional arguments. arg[1] filename, arg[2] keyword")
	}
	filename := os.Args[1]
	keyword := os.Args[2]

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
		{-1, 0, "left"},
		{-1,-1, "up_left"},
		{ 0,-1, "up"},
		{ 1,-1, "up_right"},
		{ 1, 0, "right"},  
		{ 1, 1, "down_right"},
		{ 0, 1, "down"},
		{-1, 1, "down_left"},
	}

	sum := 0
	matches := []KeywordMatch{}
	for y, row := range grid {
		for x, _ := range row {
			// Only start searches from the first char in our keyword
			if row[x] == keyword[0] {
				fmt.Println("Found", keyword[0], "at", x, y)
				for _, direction := range directions {
					match, found := findKeyword(grid, x, y, direction, keyword)
					if found {
						matches = append(matches, match)
						sum++
					}
				}
			}
		}
	}

	fmt.Println("Answer: ", sum)
}