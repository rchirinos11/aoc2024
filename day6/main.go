package day6

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

func Run() {
	solvePt1()
	solvePt2()
}

func solvePt2() {
	x, y, grid := fillGrid()
	fmt.Println()
	sum := 0
	for i, row := range grid {
		for j, char := range row {
			if char == '.' {
				a, b := x, y
				_, _, loop := processMoves(&a, &b, copyGrid(grid, i, j))
				if loop {
					sum++
				}
			}
		}
	}
	fmt.Println("Pt2 positions is", sum)
}

func copyGrid(grid []string, i, j int) []string {
	newGrid := make([]string, len(grid))
	copy(newGrid, grid)
	newGrid[i] = newGrid[i][:j] + string('#') + newGrid[i][j+1:]
	return newGrid
}

func solvePt1() {
	x, y, grid := fillGrid()
	sum, _, _ := processMoves(&x, &y, grid)
	fmt.Println("Pt1 sum is:", sum+1)
	fmt.Println("Guard is now at:", x, ",", y)
}

func fillGrid() (x, y int, grid []string) {
	scanner := util.Scanner("day6")
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, char := range line {
			if char == '^' {
				fmt.Println("Found guard! She was at", i, ",", j)
				x, y = i, j
			}
		}
		grid = append(grid, line)
	}
	return
}

// Very cursed
func moveOnce(evalX, evalY int, x, y, sum, facing *int, grid []string) {
	if grid[evalX][evalY] == '#' {
		*facing++
		if *facing > LEFT {
			*facing = UP
		}
		return
	}

	if grid[*x][*y] == '^' || grid[*x][*y] == '.' {
		*sum++
	}
	grid[*x] = grid[*x][:*y] + string('M') + grid[*x][*y+1:]
	*x, *y = evalX, evalY
}

func processMoves(x, y *int, grid []string) (sum, facing int, loop bool) {
	for i := 0; i < 9999; i++ {
		switch facing {
		case UP:
			if *x == 0 {
				return
			}
			moveOnce(*x-1, *y, x, y, &sum, &facing, grid)
		case RIGHT:
			if *y == len(grid)-1 {
				return
			}
			moveOnce(*x, *y+1, x, y, &sum, &facing, grid)
		case DOWN:
			if *x == len(grid)-1 {
				return
			}
			moveOnce(*x+1, *y, x, y, &sum, &facing, grid)
		case LEFT:
			if *y == 0 {
				return
			}
			moveOnce(*x, *y-1, x, y, &sum, &facing, grid)
		}
	}
	return sum, facing, true
}

// Completely unnecessary but at least you can see it
func printTrail(x, y, facing int, grid []string) {
	for i, row := range grid {
		for j, char := range row {
			if x == i && y == j {
				printFace(facing)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func printFace(facing int) {
	switch facing {
	case UP:
		fmt.Print(string("^"))
	case LEFT:
		fmt.Print(string("<"))
	case RIGHT:
		fmt.Print(string(">"))
	case DOWN:
		fmt.Print(string("v"))
	}
}
