package day10

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	topoMap := readMap()
	var sum1, sum2 int
	for i, line := range topoMap {
		for j, val := range line {
			if val == 0 {
				sum1 += recurse(0, i, j, topoMap, make(map[[2]int]bool))
				sum2 += recurse(0, i, j, topoMap, nil)
			}
		}
	}
	fmt.Println("Pt1 rating is:", sum1)
	fmt.Println("Pt2 rating is:", sum2)
}

func recurse(val, i, j int, topoMap [][]int, solutions map[[2]int]bool) (sum int) {
	if i < 0 || i == len(topoMap) || j < 0 || j == len(topoMap[0]) || topoMap[i][j] != val {
		return
	}
	if val == 9 {
		return fromSolutions(i, j, solutions)
	}

	sum += recurse(val+1, i-1, j, topoMap, solutions)
	sum += recurse(val+1, i, j-1, topoMap, solutions)
	sum += recurse(val+1, i+1, j, topoMap, solutions)
	sum += recurse(val+1, i, j+1, topoMap, solutions)
	return
}

func fromSolutions(i, j int, solutions map[[2]int]bool) int {
	if solutions == nil {
		return 1
	}
	key := [2]int{i, j}
	if solutions[key] {
		return 0
	}
	solutions[key] = true
	return 1
}

func readMap() (topoMap [][]int) {
	scanner := util.Scanner("day10")
	for scanner.Scan() {
		line := scanner.Text()
		nums := []int{}
		for _, char := range line {
			n := int(char - '0')
			nums = append(nums, n)
		}
		topoMap = append(topoMap, nums)
	}
	return
}
