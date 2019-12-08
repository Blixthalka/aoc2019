package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Length    int
}

type Point struct {
	x, y int
}

type CordValue struct {
	val          int
	programSteps map[int]int
}

func (p Point) step(Direction string) Point {
	switch Direction {
	case "R":
		return Point{p.x + 1, p.y}
	case "L":
		return Point{p.x - 1, p.y}
	case "U":
		return Point{p.x, p.y + 1}
	case "D":
		return Point{p.x, p.y - 1}
	}
	os.Exit(1)
	return Point{0, 0}
}

func (p Point) manhattan(p2 Point) int64 {
	return int64(math.Abs(float64(p.x-p2.x)) + math.Abs(float64(p.y-p2.y)))
}

func getInput() [][]Instruction {
	databyte, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	programs := strings.Split(string(databyte), "\n")

	parsedPrograms := make([][]Instruction, len(programs))

	for j, program := range programs {
		Instructions := strings.Split(program, ",")
		parsedInstruction := make([]Instruction, len(Instructions))

		for i, InstructionString := range Instructions {
			Direction := string(InstructionString[0])
			Length, _ := strconv.Atoi(string(InstructionString[1:]))
			parsedInstruction[i] = Instruction{Direction, Length}
		}
		parsedPrograms[j] = parsedInstruction
	}

	return parsedPrograms
}

func part2() {
	START := -1
	INTERSECTION := -2

	cords := map[Point]CordValue{}
	cords[Point{0, 0}] = CordValue{START, map[int]int{}}

	programs := getInput()

	for programID, program := range programs {
		location := Point{0, 0}
		totalSteps := 0

		for _, instruction := range program {
			for step := 0; step < instruction.Length; step++ {
				location = location.step(instruction.Direction)
				totalSteps++

				if value, exists := cords[location]; exists {
					switch {
					case value.val == START || value.val == INTERSECTION:
						continue
					case value.val != programID:
						value.programSteps[programID] = totalSteps
						cords[location] = CordValue{INTERSECTION, value.programSteps}
					}
				} else {
					programSteps := map[int]int{}
					programSteps[programID] = totalSteps
					cords[location] = CordValue{programID, programSteps}
				}
			}
		}
	}

	var shortest int64 = math.MaxInt64
	for _, v := range cords {
		if v.val != INTERSECTION {
			continue
		}

		var dist int64 = 0
		for _, steps := range v.programSteps {
			dist += int64(steps)
		}

		if dist < shortest {
			shortest = dist
		}
	}

	fmt.Println(shortest)
}

func main() {
	part2()
}
