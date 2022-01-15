package main

import (
	"bufio"
	"fmt"
	"os"
)

type BracketStack []rune

func (bs *BracketStack) push(r rune) {
	*bs = append(*bs, r)
}

func (bs *BracketStack) pop() rune {
	if len(*bs) > 0 {
		ptr := len(*bs) - 1
		top := (*bs)[ptr]
		*bs = (*bs)[:ptr]
		return top
	} else {
		return rune('a')
	}
}

func (bs *BracketStack) check(r rune) int {
	if r == rune('(') || r == rune('[') || r == rune('{') || r == rune('<') {
		bs.push(r)
	} else {
		top := bs.pop()
		if r == rune(')') && top != rune('(') {
			return 3
		} else if r == rune(']') && top != rune('[') {
			return 57
		} else if r == rune('}') && top != rune('{') {
			return 1197
		} else if r == rune('>') && top != rune('<') {
			return 25137
		}
	}
	return 0
}

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
		var stack BracketStack
		for _, e := range data {
			count += stack.check(e)
		}
	}
	fmt.Println(count)
}
