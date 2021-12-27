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

	h, v, aim := 0, 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(a[1])
		if err != nil {
			fmt.Println(err)
		}
		switch a[0] {
		case "up":
			aim -= val
		case "down":
			aim += val
		case "forward":
			h += val
			v += aim * val
		default:
			panic("Unkown direction")
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h * v)
}

