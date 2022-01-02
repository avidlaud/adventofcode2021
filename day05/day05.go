package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func lesserGreater(a, b string) (int, int) {
	aVal, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	bVal, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	if aVal > bVal {
		return bVal, aVal
	} else {
		return aVal, bVal
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// Map coordinates as string "<x>,<y>"
	m := make(map[string]int)

	count := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Split(line, " -> ")

		startCoords := strings.Split(toks[0], ",")
		endCoords := strings.Split(toks[1], ",")

		// Vertical (share x)
		if startCoords[0] == endCoords[0] {
			start, end := lesserGreater(startCoords[1], endCoords[1])
			for i := start; i <= end; i++ {
				curr := m[startCoords[0]+","+strconv.Itoa(i)]
				if curr == 1 {
					count += 1
				}
				m[startCoords[0]+","+strconv.Itoa(i)] = curr + 1
			}
		}

		// Horizontal (share y)
		if startCoords[1] == endCoords[1] {
			start, end := lesserGreater(startCoords[0], endCoords[0])
			for i:= start; i <= end; i++ {
				curr := m[strconv.Itoa(i)+","+startCoords[1]]
				if curr == 1 {
					count += 1
				}
				m[strconv.Itoa(i)+","+startCoords[1]] = curr + 1
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)
}

