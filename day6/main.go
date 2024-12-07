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
}

func solvePt1() {
	scanner := util.Scanner("day6")
	grid := []string{}
	var x, y int
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

	printTrail(x, y, 0, grid)
	fmt.Println()
	sum, facing := processMoves(&x, &y, grid)
	printTrail(x, y, facing, grid)
	fmt.Println("Sum is:", sum+1)
	fmt.Println("Guard is now at:", x, ",", y)
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

func processMoves(x, y *int, grid []string) (sum int, facing int) {
	for {
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
