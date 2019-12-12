package main

import "strconv"

import "fmt"

func main() {
	tally := validRange(172851, 675869)
	fmt.Printf("Possibilities: %d", tally)
	fmt.Println()
}

func validRange(start int, end int) int {
	if start > end {
		return 0
	}

	result := 0
	for v := start; v < end; v++ {
		if valid(v) {
			result++
		}
	}
	return result
}

func valid(number int) bool {
	value := strconv.Itoa(number)
	double := false

	// check length
	if len(value) != 6 {
		return false
	}

	var prev rune
	for i, next := range value {
		if i > 0 {
			// check double
			if prev == next {
				double = true
			}
			// check never decreasing
			if prev > next {
				return false
			}
		}
		prev = next
	}

	return double
}
