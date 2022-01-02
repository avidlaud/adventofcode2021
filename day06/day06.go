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
	newFishes := 0
	// Simulate a day
	for i := 0; i < len(fishes); i++ {
		if fishes[i] == 0 {
			newFishes += 1
			fishes[i] = 6
		} else {
			fishes[i] = fishes[i] - 1
		}
	}
	spawned := make([]int, newFishes)
	for i := 0; i < len(spawned); i++ {
		spawned[i] = 8
	}
	return spawn(append(fishes, spawned...), days-1)
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

	vals := make([]int, 9)
	for i := 0; i < 9; i++ {
		vals[i] = len(spawn([]int{i}, 80))
	}

	count := 0

	for _, fish := range initialFishes {
		x, err := strconv.Atoi(fish)
		if err != nil {
			fmt.Println(err)
			return
		}
		count += vals[x]
	}

	fmt.Println(count)
}
