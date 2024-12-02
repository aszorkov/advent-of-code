package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"log"
)

func checkVal(val int, hash map[int][2]int, col int) {
	entry := hash[val]
	if len(entry) == 0 {
		entry = [2]int{0, 0}
	}
	entry[col]++
	hash[val] = entry
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

	list_a := []int{}
	list_b := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineStr := scanner.Text()
		slicedStr := strings.Fields(lineStr)
		x, _ := strconv.Atoi(slicedStr[0])
		y, _ := strconv.Atoi(slicedStr[1])
		list_a = append(list_a, x)
		list_b = append(list_b, y)
	}

	sum := 0
	hashmap := make(map[int][2]int)

	// Builds a hashmap like this:
	// 1234: [x, y]
	// Where x is num occurrences in left col and y is num occurrences in right col
    for index, _ := range list_a {
		checkVal(list_a[index], hashmap, 0)
		checkVal(list_b[index], hashmap, 1)
	}

	// For each entry in map, multiply key value by x by y
	for key, value := range hashmap {
		fmt.Println("Key: ", key, ", Left: ", value[0], ", Right: ", value[1])
		sum += key * value[0] * value[1]
	}

	fmt.Println("Answer: ", sum)
}