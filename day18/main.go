package day18

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

const size = 71

func Run() {
	scanner := util.Scanner("day18")
	grid := [size][size]int{}
	for i := 0; scanner.Scan(); i++ {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		grid[y][x] = '#' - '.' // for printing purposes
		res := dijkstra(grid)
		if i == 1024 {
			fmt.Println("Pt1 sol:", res)
		}
		if res == 0 {
			fmt.Println("Pt2 sol:")
			fmt.Printf("%d,%d\n", x, y)
			break
		}
	}
}

func dijkstra(grid [size][size]int) int {
	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[[2]int]int)
	set := [][3]int{{0, 0, 0}}

	for len(set) > 0 {
		node := set[0]
		set = set[1:]

		for _, move := range moves {
			mx, my := node[0]+move[0], node[1]+move[1]
			key := [2]int{mx, my}
			if mx < 0 || my < 0 || mx == size || my == size || grid[mx][my] == '#'-'.' {
				continue
			}
			dist := 1 + node[2]
			_, ok := visited[key]
			if ok {
				continue
			}
			visited[key] = dist
			nbor := [3]int{mx, my, dist}
			set = append(set, nbor)
		}
	}
	return visited[[2]int{size - 1, size - 1}]
}
