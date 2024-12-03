package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day1")

	var x, y []int
	var len int
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, "   ")
		a, _ := strconv.Atoi(splits[0])
		b, _ := strconv.Atoi(splits[1])
		x = append(x, int(a))
		y = append(y, int(b))
		len++
	}

	slices.Sort(x)
	slices.Sort(y)
	solveDistance(x, y, len)
	solveSimilarity(x, y, len)
}

func solveSimilarity(x []int, y []int, len int) {
	var sim int
	for i := 0; i < len; i++ {
		var find int
		for j := 0; j < len; j++ {
			if x[i] == y[j] {
				find++
			}
		}
		sim += find * x[i]
	}
	fmt.Println("Similarity is:", sim)
}

func solveDistance(x, y []int, len int) {
	var sum int
	for i := 0; i < len; i++ {
		sum += abs(x[i] - y[i])
	}
	fmt.Println("Total distance between the lists is:", sum)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
