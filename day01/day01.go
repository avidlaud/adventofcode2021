package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var prev int
	count, i := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting to number")
		}
		if (curr - prev) > 0 && i > 0 {
			count++
		}
		prev = curr
		i++
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}
