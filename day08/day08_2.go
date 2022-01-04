package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func determineNumber(dict [][]bool, num []bool) int {
	for i, e := range dict {
		for a, b := range e {
			if b != num[a] {
				break
			}
			return i
		}
	}
	fmt.Println("Could not determine number!")
	return -1
}

func countOccurrence(nums []string) ([]int, map[int][]int) {
	occurence := make([]int, 7)
	for _, num := range nums {
		for _, rune := range num {
			occurence[rune-'a'] += 1
		}
	}
	occurMap := make(map[int][]int)
	for i, e := range occurence {
		x := occurMap[e]
		x = append(x, i)
		occurMap[e] = x
	}
	return occurence, occurMap
}

func countLength(nums []string) map[int][]string {
	lengthMap := make(map[int][]string)
	for _, num := range nums {
		x := lengthMap[len(num)]
		x = append(x, num)
		lengthMap[len(num)] = x
	}
	return lengthMap
}

func generateDict(occurMap map[int][]int, lengthMap map[int][]string) []int {
	translate := make([]int, 7)
	// Segment that occurs six times is B
	B := occurMap[6][0]
	translate[B] = 1
	// Segment that occurs four times is E
	E := occurMap[4][0]
	translate[E] = 4
	// Segment that occurs nine times is F
	F := occurMap[9][0]
	translate[F] = 5
	// Segment that occurs in length 3 (7) but not 2 (1) is A
	A := findDiff(lengthMap[3][0], lengthMap[2][0])
	translate[A] = 0
	// Segment that occurs in length 2 but is not F is C
	C := removeSegments(lengthMap[2][0], []int{F})
	translate[C] = 2
	// Segment that occurs in length 4 that is not B, C, or F is D
	D := removeSegments(lengthMap[4][0], []int{B, C, F})
	translate[D] = 3
	// Last one remaining is G
	G := 21 - A - B - C - D - E - F
	translate[G] = 6
	return translate
}

func findDiff(a, b string) int {
	for _, x := range a {
		if !strings.ContainsRune(b, x) {
			return int(x - rune('a'))
		}
	}
	fmt.Println("Error finding rune difference!")
	return -1
}

func removeSegments(orig string, remove []int) int {
CHECKNUM:
	for _, e := range orig {
		x := int(e - 'a')
		for _, c := range remove {
			if x == c {
				continue CHECKNUM
			}
		}
		return x
	}
	fmt.Println("Error removing rune")
	return -1
}

func translate(dict []int, num string) int {
	seg := make([]bool, 7)
	for _, rune := range num {
		seg[dict[rune-'a']] = true
	}
	return getSegment(seg)
}

func getSegment(seg []bool) int {
	key := ""
	for i, e := range seg {
		if e {
			key += strconv.Itoa(i)
		}
	}
	switch key {
	case "012456":
		return 0
	case "25":
		return 1
	case "02346":
		return 2
	case "02356":
		return 3
	case "1235":
		return 4
	case "01356":
		return 5
	case "013456":
		return 6
	case "025":
		return 7
	case "0123456":
		return 8
	case "012356":
		return 9
	default:
		fmt.Println("Error getting segment")
		return -1
	}

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
		tokens := strings.Split(data, " | ")
		nums := strings.Fields(tokens[0])
		res := strings.Fields(tokens[1])
		_, occurMap := countOccurrence(nums)
		dict := generateDict(occurMap, countLength(nums))
		n := ""
		for _, r := range res {
			n += strconv.Itoa(translate(dict, r))
		}
		x, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
			return
		}
		count += x
	}

	fmt.Println(count)
}
