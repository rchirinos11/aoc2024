package day15

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
}

func solvePt1() {
	grid := read()
	// printGrid(grid)
	sum := gpsSum(grid)
	fmt.Println("Gps sum:", sum)
}

func gpsSum(grid [][]byte) (sum int) {
	for i, row := range grid {
		for j, char := range row {
			if char == 'O' {
				sum += i*100 + j
			}
		}
	}
	return
}

func read() [][]byte {
	reader := util.Reader("day15")
	grid := [][]byte{}
	x, y := -1, -1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || len(line) == 1 {
			break
		}
		for i := 0; i < len(line) && x == -1; i++ {
			if line[i] == '@' {
				y, x = i, len(grid)
			}
		}
		grid = append(grid, line[:len(line)-1])
	}

	fmt.Println("Starting")
	fmt.Println("X:", x, "Y:", y)
	// printGrid(grid)
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		x, y = moveRobot(grid, char, x, y)
	}

	fmt.Println("X:", x, "Y:", y)
	return grid
}

func moveRobot(grid [][]byte, char byte, x, y int) (int, int) {
	switch char {
	case '^':
		if x == 0 || grid[x-1][y] == '#' {
			return x, y
		}
		i := 1
		for ; grid[x-i][y] == 'O'; i++ {
			if grid[x-1-i][y] == '#' {
				return x, y
			}
		}
		if grid[x-i][y] == '.' {
			grid[x-1][y] = '@'
			grid[x][y] = '.'
			for j := 1; j < i; j++ {
				grid[x-1-j][y] = 'O'
			}
			return x - 1, y
		}
	case 'v':
		if x == len(grid[0])-1 || grid[x+1][y] == '#' {
			break
		}
		i := 1
		for ; grid[x+i][y] == 'O'; i++ {
			if grid[x+1+i][y] == '#' {
				return x, y
			}
		}
		if grid[x+i][y] == '.' {
			grid[x+1][y] = '@'
			grid[x][y] = '.'
			for j := 1; j < i; j++ {
				grid[x+1+j][y] = 'O'
			}
			return x + 1, y
		}
	case '<':
		if y == 0 || grid[x][y-1] == '#' {
			break
		}
		i := 1
		for ; grid[x][y-i] == 'O'; i++ {
			if grid[x][y-1-i] == '#' {
				return x, y
			}
		}
		if grid[x][y-i] == '.' {
			grid[x][y-1] = '@'
			grid[x][y] = '.'
			for j := 1; j < i; j++ {
				grid[x][y-1-j] = 'O'
			}
			return x, y - 1
		}
	case '>':
		if y == len(grid)-1 || grid[x][y+1] == '#' {
			break
		}
		i := 1
		for ; grid[x][y+i] == 'O'; i++ {
			if grid[x][y+1+i] == '#' {
				return x, y
			}
		}
		if grid[x][y+i] == '.' {
			grid[x][y+1] = '@'
			grid[x][y] = '.'
			for j := 1; j < i; j++ {
				grid[x][y+1+j] = 'O'
			}
			return x, y + 1
		}
	}
	return x, y
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c", char)
		}

		fmt.Println()
	}
}
