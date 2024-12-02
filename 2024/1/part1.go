package main

import (
	"fmt"
	"slices"
	"strings"
	"strconv"
	"bufio"
	"os"
	"log"
)

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
		fmt.Println(x, y)
	}

	slices.Sort(list_a)
	slices.Sort(list_b)

	sum := 0

    for index, _ := range list_a {
		pair := []int{list_a[index], list_b[index]}
		slices.Sort(pair);
		distance := pair[1] - pair[0]
		sum += distance
		//fmt.Println(pair[1], pair[0], ":", distance)
	}

	fmt.Println("Answer: ", sum)
}