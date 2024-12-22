package day20

import (
	"fmt"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	grid, sX, sY, eX, eY := read()

	sol := dijkstra(grid, sX, sY, eX, eY, false, nil)
	fmt.Println("Base:", sol)
	solvePt1(grid, sX, sY, eX, eY, sol)
}

func solvePt1(grid []string, sX, sY, eX, eY, initial int) {
	var sum int
	attempts := make(map[[2]int]bool)
	for prev := 0; ; prev = len(attempts) {
		sol := dijkstra(grid, sX, sY, eX, eY, true, attempts)
		if initial-sol >= 100 {
			sum++
		}
		if prev == len(attempts) {
			fmt.Println("Prev:", prev)
			break
		}
	}
	fmt.Println("amount:", sum)
}

func read() (grid []string, sX, sY, eX, eY int) {
	reader := util.Reader("day20")
	sX, sY, eX, eY = -1, -1, -1, -1
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if sX == -1 && strings.ContainsRune(line, 'S') {
			sX, sY = len(grid), strings.IndexRune(line, 'S')
		}
		if eX == -1 && strings.ContainsRune(line, 'E') {
			eX, eY = len(grid), strings.IndexRune(line, 'E')
		}
		grid = append(grid, line[:len(line)-1])
	}
}

func dijkstra(grid []string, sX, sY, eX, eY int, cheats bool, cheated map[[2]int]bool) int {
	moves := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[[2]int]int)
	set := [][3]int{{sX, sY, 0}}

	for len(set) > 0 {
		node := set[0]
		set = set[1:]

		for _, move := range moves {
			mx, my := node[0]+move[0], node[1]+move[1]
			if mx == 0 || my == 0 || mx == len(grid)-1 || my == len(grid)-1 {
				continue
			}
			key := [2]int{mx, my}
			if grid[mx][my] == '#' {
				if !cheats || cheated[key] {
					continue
				}
				cheats = false
				cheated[key] = true
			}
			if visited[key] > 0 {
				continue
			}
			visited[key] = node[2] + 1
			nbor := [3]int{mx, my, node[2] + 1}
			set = append(set, nbor)
		}
	}
	return visited[[2]int{eX, eY}]
}
