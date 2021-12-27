package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	bits := make([]int, 12)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i, b := range(scanner.Text()) {
			if b == 48 {
				bits[i] -= 1
			} else {
				bits[i] += 1
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}

	gamma := make([]string, 12)
	epsilon := make([]string, 12)
	for i, e := range(bits) {
		if e >= 0 {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
	}
	
	g, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	e, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g * e)
}

