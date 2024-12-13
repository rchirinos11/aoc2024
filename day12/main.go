package day12

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

type Key struct {
	Char byte
	PosI int
	PosJ int
}

func Run() {
	solvePt1()
}

func solvePt1() {
	scanner := util.Scanner("day12")
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	var sum int
	visited := make(map[Key]bool)
	for i, line := range grid {
		for j, char := range line {
			area, perim := findAll(byte(char), i, j, grid, visited)
			sum += area * perim
		}
	}
	fmt.Println("Final sum:", sum)
}

func findAll(char byte, i, j int, grid []string, visited map[Key]bool) (area, perim int) {
	key := Key{char, i, j}
	if visited[key] || i == len(grid) || i < 0 || j == len(grid[0]) || j < 0 {
		return
	}
	if grid[i][j] != char {
		return 0, 1
	}

	area++
	visited[key] = true
	perim = evalPerim(i, j, grid)

	moves := [4][2]int{{i + 1, j}, {i, j + 1}, {i - 1, j}, {i, j - 1}}
	for _, move := range moves {
		a, p := findAll(char, move[0], move[1], grid, visited)
		area += a
		perim += p
	}

	return
}

func evalPerim(i, j int, grid []string) (perim int) {
	perimConditions := [4]bool{j == 0, j == len(grid[i])-1, i == len(grid)-1, i == 0}
	for _, cond := range perimConditions {
		if cond {
			perim++
		}
	}

	return
}
