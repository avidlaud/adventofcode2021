package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateFuel(positions []int, target int) int {
	fuel := 0
	for _, e := range positions {
		fuel += abs(e - target)
	}
	return fuel
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

	tokens := strings.Split(input, ",")
	positions := make([]int, len(tokens))

	minPosition := math.MaxInt
	maxPosition := 0

	for i, tok := range tokens {
		pos, err := strconv.Atoi(tok)
		if err != nil {
			fmt.Println(err)
			return
		}
		positions[i] = pos
		if pos < minPosition {
			minPosition = pos
		}
		if pos > maxPosition {
			maxPosition = pos
		}
	}

	minFuel := math.MaxInt

	for target := minPosition; target < maxPosition; target++ {
		fuelUsed := calculateFuel(positions, target)
		if fuelUsed < minFuel {
			minFuel = fuelUsed
		}
	}

	fmt.Println(minFuel)
}
