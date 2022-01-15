package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func search(grid *[][]int, row, col int) int {
	if col < 0 || row < 0 || row >= len(*grid) || col >= len((*grid)[0]) || (*grid)[row][col] == 9 {
		return 0
	}
	(*grid)[row][col] = 9
	n := search(grid, row-1, col)
	s := search(grid, row+1, col)
	e := search(grid, row, col+1)
	w := search(grid, row, col-1)
	return 1 + n + s + e + w
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

	lowpoints := [][2]int{}
	for i := range heightmap {
		for j := range heightmap[0] {
			if inspectPoint(j, i, heightmap) > 0 {
				lowpoints = append(lowpoints, [2]int{i, j})
			}
		}
	}

	sizes := make([]int, len(lowpoints))
	for i, e := range lowpoints {
		sizes[i] = search(&heightmap, e[0], e[1])
	}

	sort.Ints(sizes)
	fmt.Println(sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3])
}
