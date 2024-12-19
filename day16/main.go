package day16

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	reader := util.Reader("day16")
	grid := [][]byte{}
	sX, sY, eX, eY := -1, -1, -1, -1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil || len(line) == 1 {
			break
		}
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				sX, sY = len(grid), i
			}
			if line[i] == 'E' {
				eX, eY = len(grid), i
			}
		}
		grid = append(grid, line[:len(line)-1])
	}

	dijkstra(grid, sX, sY, eX, eY)
}

func dijkstra(grid [][]byte, sX, sY, eX, eY int) {
	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[[2]int]int)
	set := [][4]int{{sX, sY, 0, 1}}

	for len(set) > 0 {
		node := set[0]
		set = set[1:]

		for f, move := range moves {
			mx, my := node[0]+move[0], node[1]+move[1]
			key := [2]int{mx, my}
			if grid[mx][my] == '#' {
				continue
			}
			dist := 1000*abs(node[3], f) + 1 + node[2]
			val, ok := visited[key]
			if ok && val < dist {
				continue
			}
			visited[key] = dist
			nbor := [4]int{mx, my, dist, f}
			set = append(set, nbor)
		}
	}
	fmt.Println("Pt1 solution is:", visited[[2]int{eX, eY}])
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c", char)
		}

		fmt.Println()
	}
}

func abs(n1, n2 int) int {
	if n1 == 0 && n2 == 3 || n1 == 3 && n2 == 0 {
		return 1
	}
	if n1 > n2 {
		return n1 - n2
	}
	return n2 - n1
}
