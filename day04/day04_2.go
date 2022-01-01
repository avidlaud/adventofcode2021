package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type BingoCard struct {
	size	int
	nums	[]uint64
}

func newBingoCard(size int) *BingoCard {
	return &BingoCard {
		size: size,
		nums: make([]uint64, size * size),
	}
}

func (bc *BingoCard) fillBingoCard(numbers []string) {
	for row, nums := range numbers {
		entries := strings.Fields(nums)
		for i, entry := range entries {
			x, err := strconv.ParseUint(entry, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			bc.nums[(row * bc.size) + i] = x
		}
	}
}

func (bc *BingoCard) checkBingoCard(drawn []bool) bool {
	CHECKROWS:
	for row := 0; row < bc.size; row++ {
		for col := 0; col < bc.size; col++ {
			if drawn[bc.nums[(row * bc.size) + col]] == false {
				continue CHECKROWS
			}
		}
		return true
	}

	CHECKCOLS:
	for col := 0; col < bc.size; col++ {
		for row := 0; row < bc.size; row++ {
			if drawn[bc.nums[(row * bc.size) + col]] == false {
				continue CHECKCOLS
			}
		}
		return true
	}
	
	return false
}

func (bc *BingoCard) scoreBingoCard(drawn []bool) uint64 {
	score := uint64(0)
	for i := 0; i < bc.size * bc.size; i++ {
		if drawn[bc.nums[i]] == false {
			score += bc.nums[i]
		}
	}
	return score
}

func lastOne(drawn []bool) int {
	foundOne := -1
	for i, e := range drawn {
		if e == false {
			if foundOne != -1 {
				return -1
			}
			foundOne = i
		}
	}
	return foundOne
}

func allDone(done []bool) bool {
	for _, e := range done {
		if e == false {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Get the list of bingo draws
	scanner.Scan()
	drawsRaw := scanner.Text()
	draws := strings.Split(drawsRaw, ",")
	
	drawNumbers := make([]uint64, len(draws))

	maxDraw := uint64(0)
	for i, e := range draws {
		x, err := strconv.ParseUint(e, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		drawNumbers[i] = x
		if x > maxDraw {
			maxDraw = x
		}
	}

	boardLines := make([]string, 0)

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		boardLines = append(boardLines, line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}

	BOARD_SIZE := 5
	boards := make([]*BingoCard, 0)

	for i := 0; i < len(boardLines); i += 6 {
		board := newBingoCard(BOARD_SIZE)
		board.fillBingoCard(boardLines[i:i+5])
		boards = append(boards, board)
	}

	drawn := make([]bool, maxDraw + 1)
	done := make([]bool, len(boards))
	prev := uint64(0)

	last := -1

	for _, e := range drawNumbers {
		x := lastOne(done)
		if x != -1 {
			last = x
		}
		if allDone(done) {
			fmt.Println((uint64(prev) * boards[last].scoreBingoCard(drawn)))
			return
		}
		drawn[e] = true
		for i, board := range boards {
			if done[i] == false && board.checkBingoCard(drawn) {
				done[i] = true
			}
		}
		prev = e
	}

	fmt.Println("No board found")
}

