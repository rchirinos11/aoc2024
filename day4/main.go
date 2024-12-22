package day4

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	puzzle := read()
	var pt1Sum, pt2Sum int
	for i := range puzzle {
		for j, char := range puzzle[i] {
			if char == 'X' {
				pt1Sum += solvePt1(puzzle, i, j)
			} else if char == 'A' {
				pt2Sum += solvePt2(puzzle, i, j)
			}
		}
	}
	fmt.Println("Pt1 sum:", pt1Sum)
	fmt.Println("Pt2 sum:", pt2Sum)
}

func solvePt2(puzzle []string, i, j int) (sum int) {
	if i == 0 || i == len(puzzle)-1 || j == 0 || j == len(puzzle[0])-1 {
		return
	}
	down := [][]int{{1, -1}, {1, 1}, {-1, -1}, {-1, 1}}
	up := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	right := [][]int{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
	left := [][]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
	combinations := [][][]int{down, up, right, left}
	for _, comb := range combinations {
		func() {
			for n, side := range comb {
				if n >= 2 && puzzle[i+side[0]][j+side[1]] != 'S' || n < 2 && puzzle[i+side[0]][j+side[1]] != 'M' {
					return
				}
			}
			sum++
		}()
	}

	return
}

func read() (puzzle []string) {
	scanner := util.Scanner("day4")
	for scanner.Scan() {
		puzzle = append(puzzle, scanner.Text())
	}
	return
}

func solvePt1(puzzle []string, i, j int) (sum int) {
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
	return
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
