package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		return rune('@')
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

func (bs *BracketStack) insert(r rune) {
	if r == rune('(') || r == rune('[') || r == rune('{') || r == rune('<') {
		bs.push(r)
	} else {
		bs.pop()
	}
}

func (bs *BracketStack) stackVal() int {
	top := rune('a')
	val := 0
	for top != rune('@') {
		if top == rune('(') {
			val *= 5
			val += 1
		} else if top == rune('[') {
			val *= 5
			val += 2
		} else if top == rune('{') {
			val *= 5
			val += 3
		} else if top == rune('<') {
			val *= 5
			val += 4
		}
		top = bs.pop()
	}
	return val
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	counts := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Text()
		var checkStack BracketStack
		curr := 0
		for _, e := range data {
			curr += checkStack.check(e)
		}
		if curr != 0 {
			continue
		}
		var stack BracketStack
		for _, e := range data {
			stack.insert(e)
		}
		counts = append(counts, stack.stackVal())
	}
	sort.Ints(counts)
	fmt.Println(counts[(len(counts) / 2)])
}
