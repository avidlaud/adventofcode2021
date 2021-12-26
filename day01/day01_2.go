package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type circularQueue struct{
	size	int
	values	[]int
	ptr		int
}

func newCircularQueue(size int) *circularQueue {
	values := make([]int, size)
	for i := 0; i < size; i++ {
		values[i] = math.MaxInt
	}
	return &circularQueue{
		size:	size,
		values:	values,
	}
}

func (cq *circularQueue) insert(x int) (out int) {
	out = cq.values[cq.ptr]
	cq.values[cq.ptr] = x
	cq.ptr = (cq.ptr + 1) % cq.size
	return out
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	cq := newCircularQueue(3)
	count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting to number")
		}
		if curr > cq.insert(curr) {
			count++
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}


	
