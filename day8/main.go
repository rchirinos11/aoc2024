package day8

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
}

func solvePt1() {
	cityMap, antennaMap := readInput()
	unique := 0
	for i, line := range cityMap {
		for j, char := range line {
			if char == '.' {
				continue
			}
			unique += evalChar(i, j, char, cityMap, antennaMap)
		}
	}
	printMap(cityMap, antennaMap)
	fmt.Println("Unique:", unique)
}

func evalChar(i, j int, char rune, cityMap []string, antennaMap [][]bool) (sum int) {
	for x, line := range cityMap {
		if x == i {
			break
		}
		for y, val := range line {
			if char == val {
				sum += modByDiff(i, j, x, y, antennaMap)
			}
		}
	}
	return
}

func modByDiff(i, j, x, y int, antennaMap [][]bool) (sum int) {
	lineLen := len(antennaMap[0])
	if iDiff, jDiff := 2*i-x, 2*j-y; iDiff < len(antennaMap) && jDiff < lineLen && jDiff >= 0 {
		if !antennaMap[iDiff][jDiff] {
			sum++
		}
		antennaMap[iDiff][jDiff] = true
	}

	if xDiff, yDiff := 2*x-i, 2*y-j; xDiff >= 0 && yDiff >= 0 && yDiff < lineLen {
		if !antennaMap[xDiff][yDiff] {
			sum++
		}
		antennaMap[xDiff][yDiff] = true
	}
	return
}

func printMap(cityMap []string, antennaMap [][]bool) {
	for i, line := range cityMap {
		for y, val := range line {
			if antennaMap[i][y] {
				fmt.Print("#")
				continue
			}

			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func readInput() ([]string, [][]bool) {
	cityMap := []string{}
	size := 0
	scanner := util.Scanner("day8")

	for ; scanner.Scan(); size++ {
		line := scanner.Text()
		cityMap = append(cityMap, line)
	}

	antennaMap := make([][]bool, size)
	for i := 0; i < size; i++ {
		line := make([]bool, len(cityMap[0]))
		antennaMap[i] = line
	}
	return cityMap, antennaMap
}
