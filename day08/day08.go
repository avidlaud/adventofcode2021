package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		tokens := strings.Split(data, " | ")
		output := strings.Fields(tokens[1])
		for _, x := range output {
			if !(len(x) == 5 || len(x) == 6) {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
