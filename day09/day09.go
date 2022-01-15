package main

import (
	"bufio"
	"fmt"
	"os"
)

func inspectPoint(col, row int, heightmap [][]int) int {
	// NESW
	adj := [4]int{10, 10, 10, 10}
	if row-1 >= 0 {
		adj[0] = heightmap[row-1][col]
	}
	if row+1 < len(heightmap) {
		adj[2] = heightmap[row+1][col]
	}
	if col-1 >= 0 {
		adj[1] = heightmap[row][col-1]
	}
	if col+1 < len(heightmap[0]) {
		adj[3] = heightmap[row][col+1]
	}
	curr := heightmap[row][col]
	for _, e := range adj {
		if e <= curr {
			return 0
		}
	}
	return curr + 1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	heightmap := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		line := make([]int, len(data))
		for i, c := range data {
			line[i] = int(c - '0')
		}
		heightmap = append(heightmap, line)
	}

	count := 0
	for i := range heightmap {
		for j := range heightmap[0] {
			count += inspectPoint(j, i, heightmap)
		}
	}
	fmt.Println(count)
}
