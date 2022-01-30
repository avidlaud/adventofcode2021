package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type OctopusGrid [][]int

type FlashStack [][2]int

func (fs *FlashStack) push(x [2]int) {
	*fs = append(*fs, x)
}

func (fs *FlashStack) pop() ([2]int, error) {
	if fs.hasItems() {
		ptr := len(*fs) - 1
		top := (*fs)[ptr]
		*fs = (*fs)[:ptr]
		return top, nil
	} else {
		return [2]int{0, 0}, errors.New("Nothing in FlashStack")
	}
}

func (fs *FlashStack) hasItems() bool {
	return len(*fs) > 0
}

func (og *OctopusGrid) increment(x, y int) bool {
	if y >= 0 && y < len(*og) && x >= 0 && x < len((*og)[0]) {
		(*og)[y][x]++
		curr := (*og)[y][x]
		if curr == 10 {
			return true
		}
	}
	return false
}

func (og *OctopusGrid) resetFlashed() (allFlashed bool) {
	allFlashed = true
	for i := range *og {
		for j := range (*og)[0] {
			if (*og)[i][j] > 9 {
				(*og)[i][j] = 0
			} else {
				allFlashed = false
			}
		}
	}
	return
}

func (fs *FlashStack) addAllToStack(og *OctopusGrid) {
	for i := range *og {
		for j := range (*og)[0] {
			fs.push([2]int{j, i})
		}
	}
}

func (fs *FlashStack) addAdjacent(pos [2]int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				fs.push([2]int{pos[0] + i, pos[1] + j})
			}
		}
	}
}

func (og *OctopusGrid) simulateTurn() bool {
	var stack FlashStack
	stack.addAllToStack(og)
	for stack.hasItems() {
		target, err := stack.pop()
		if err != nil {
			fmt.Println(err)
			return false
		}
		if og.increment(target[0], target[1]) {
			stack.addAdjacent(target)
		}
	}
	return og.resetFlashed()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var grid OctopusGrid

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		line := make([]int, len(data))
		for i, c := range data {
			line[i] = int(c - '0')
		}
		grid = append(grid, line)
	}
	count := 0
	allFlashed := false
	for !allFlashed {
		count++
		allFlashed = grid.simulateTurn()
	}
	fmt.Println(count)
}
