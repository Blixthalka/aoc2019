package main

import (
	"fmt"
	"strconv"
)

func containsTwoIdenticalAdjacent(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			return true
		}
	}
	return false
}

func containsOnlyIncreasingDigits(input string) bool {
	lastDigit, _ := strconv.Atoi(string(input[0]))
	for i := 1; i < len(input); i++ {
		digit, _ := strconv.Atoi(string(input[i]))
		if lastDigit > digit {
			return false
		}
		lastDigit = digit
	}
	return true
}

func part1() {
	low := 387638
	high := 919123

	sum := 0
	for i := low + 1; i < high; i++ {
		input := strconv.Itoa(i)
		if containsOnlyIncreasingDigits(input) && containsTwoIdenticalAdjacent(input) {
			sum++
		}
	}
	fmt.Println(sum)
}

func containsOnlyTwoIdenticalAdjacent(input string) bool {
	combo := false
	comboNum := 0
	comboPoints := 0

	for i := 1; i < len(input); i++ {
		num, _ := strconv.Atoi(string(input[i]))

		if input[i-1] == input[i] {
			combo = true
			comboNum = num
			comboPoints++
			continue
		}

		if combo && comboNum != num {
			if comboPoints == 1 {
				return true
			}
			combo = false
			comboPoints = 0
		}

	}

	if comboPoints == 1 {
		return true
	}

	return false
}

func part2() {
	low := 387638
	high := 919123

	sum := 0
	for i := low + 1; i < high; i++ {
		input := strconv.Itoa(i)

		//fmt.Println(input, containsOnlyTwoIdenticalAdjacent(input))

		if containsOnlyIncreasingDigits(input) && containsOnlyTwoIdenticalAdjacent(input) {
			sum++
		}
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
