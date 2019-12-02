package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func get_input() []string {
	data_byte, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	return strings.Split(string(data_byte), "\n")
}

func calc_fuel(weight float64) float64 {
	return math.Floor(weight/3.0) - 2.0
}

func part1() {
	data := get_input()
	sum := 0.0
	for _, s := range data {
		weight, _ := strconv.ParseFloat(s, 64)
		sum += calc_fuel(weight)
	}
	fmt.Printf("part1: %.0f\n", sum)
}

func calc_fuel_rec(weight float64) float64 {
	fuel := calc_fuel(weight)
	if fuel > 0.0 {
		return fuel + calc_fuel_rec(fuel)
	} else {
		return 0.0
	}
}

func part2() {
	data := get_input()
	sum := 0.0
	for _, s := range data {
		weight, _ := strconv.ParseFloat(s, 64)
		sum += calc_fuel_rec(weight)
	}
	fmt.Printf("part2: %.0f\n", sum)
}

func main() {
	part1()
	part2()
}
