package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getInput() []int {
	dataByte, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}

	stringInput := strings.Split(string(dataByte), ",")
	intInput := make([]int, len(stringInput))

	for i, value := range stringInput {
		intValue, _ := strconv.Atoi(value)
		intInput[i] = intValue
	}
	return intInput
}

func calc(data []int) int {
	i := 0
Loop:
	for true {
		input1 := data[data[i+1]]
		input2 := data[data[i+2]]
		result := 0

		switch data[i] {
		case 1:
			result = input1 + input2
		case 2:
			result = input1 * input2
		case 99:
			break Loop
		default:
			os.Exit(1)
		}
		data[data[i+3]] = result
		i = i + 4
	}
	return data[0]
}

func part1() {
	data := getInput()
	fmt.Println(calc(data))
}

func part2() {
	dataOriginal := getInput()

All:
	for noun := 1; noun < 100; noun++ {
		for verb := 1; verb < 100; verb++ {
			data := make([]int, len(dataOriginal))
			copy(data, dataOriginal)
			data[1] = noun
			data[2] = verb

			result := calc(data)

			if result == 19690720 {
				fmt.Println(100*noun + verb)
				break All
			}
		}
	}

}

func main() {
	part1()
	part2()
}
