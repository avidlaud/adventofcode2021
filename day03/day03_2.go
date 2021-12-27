package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mostCommonBit(readings []uint64, mask uint64) int8 {
	counter := 0
	for _, reading := range readings {
		if reading & mask != 0 {
			counter += 1
		} else {
			counter -= 1
		}
	}
	return int8(counter)
}

func filterBits(readings []uint64, mask, expected uint64) []uint64 {
	pass := []uint64{}
	for _, reading := range readings {
		if reading & mask == expected {
			pass = append(pass, reading)
		}
	}
	return pass
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	readings := []uint64{}

	for scanner.Scan() {
		data := scanner.Text()
		x, err := strconv.ParseUint(data, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		readings = append(readings, x)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
	o := make([]uint64, len(readings))
	co := make([]uint64, len(readings))
	copy(o, readings)
	copy(co, readings)

	var oFinal, coFinal, expectedM, expectedL uint64

	// Iterate through bits
	for mask := uint64(1) << 11; mask > 0; mask >>= 1 {
		if len(o) == 1 {
			oFinal = o[0]
		} else {
			if mostCommonBit(o, mask) >= 0 {
				expectedM = 1
			} else {
				expectedM = 0
			}
			o = filterBits(o, mask, expectedM * mask)
		}

		if len(co) == 1 {
			coFinal = co[0]
		} else {
			if mostCommonBit(co, mask) >= 0 {
				expectedL = 0
			} else {
				expectedL = 1
			}
			co = filterBits(co, mask, expectedL * mask)
		}
	}
	fmt.Println(oFinal * coFinal)
}

