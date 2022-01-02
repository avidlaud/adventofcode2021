package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spawn(fishes []int, days int) []int {
	if days == 0 {
		return fishes
	}
	newFishes := make([]int, 9)
	// Simulate a day
	for i := 0; i < len(fishes); i++ {
		if i == 0 {
			newFishes[8] = fishes[i]
			newFishes[6] = fishes[i]
		} else {
			newFishes[i-1] += fishes[i]
		}
	}
	return spawn(newFishes, days-1)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	input := scanner.Text()

	initialFishes := strings.Split(input, ",")

	fishes := make([]int, 9)

	for _, fish := range initialFishes {
		x, err := strconv.Atoi(fish)
		if err != nil {
			fmt.Println(err)
			return
		}
		fishes[x] += 1
	}

	finalCounts := spawn(fishes, 256)

	count := 0
	for _, s := range finalCounts {
		count += s
	}

	fmt.Println(count)
}
