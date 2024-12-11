package day4

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
}

func solvePt1() {
	var puzzle []string
	scanner := util.Scanner("day4")
	for scanner.Scan() {
		puzzle = append(puzzle, scanner.Text())
	}

	var sum int
	for i, line := range puzzle {
		for j, char := range line {
			if char != 'X' {
				continue
			}

			up := [][]int{{1, 0}, {2, 0}, {3, 0}}
			down := [][]int{{-1, 0}, {-2, 0}, {-3, 0}}
			left := [][]int{{0, 1}, {0, 2}, {0, 3}}
			right := [][]int{{0, -1}, {0, -2}, {0, -3}}
			dgUp := [][]int{{1, 1}, {2, 2}, {3, 3}}
			dgLeft := [][]int{{1, -1}, {2, -2}, {3, -3}}
			dgDown := [][]int{{-1, -1}, {-2, -2}, {-3, -3}}
			dgRight := [][]int{{-1, 1}, {-2, 2}, {-3, 3}}
			paths := [][][]int{up, down, left, right, dgLeft, dgRight, dgUp, dgDown}
			for _, path := range paths {
				sum += isMas(i, j, path, puzzle)
			}
		}
	}
	fmt.Println("Word count:", sum)
}

func isMas(i, j int, path [][]int, puzzle []string) int {
	str := "MAS"
	for n, move := range path {
		x := i + move[0]
		y := j + move[1]
		if x < 0 || x >= len(puzzle) || y < 0 || y >= len(puzzle[0]) || str[n] != puzzle[x][y] {
			return 0
		}
	}

	return 1
}
