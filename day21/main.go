package day21

import (
	"container/list"
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

var (
	nPad = makeNumPad()
	dPad = makeDirPad()
)

func Run() {
	scanner := util.Scanner("day21")
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		pad1 := getPath(line, nPad)
		pad2 := getPath(pad1, dPad)
		pad3 := getPath(pad2, dPad)
		sum += calcComplex(pad3, line)
	}
	fmt.Println("Complexity:", sum)
}

func getPath(line string, pad [][]int) string {
	start := find('A', pad)
	path := ""
	for _, char := range line {
		target := find(char, pad)
		_, p := bfs(pad, start, target)
		start = target
		path += p
	}
	return path
}

func find(char rune, pad [][]int) [2]int {
	for i, line := range pad {
		for j, val := range line {
			if int(char) == val {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func bfs(pad [][]int, start, target [2]int) (int, string) {
	moves := [][]int{{-1, 0, '^'}, {0, 1, '>'}, {1, 0, 'v'}, {0, -1, '<'}}
	visited := make(map[[2]int]int)
	queue := list.New()
	queue.PushBack(start)
	queue.PushBack("")
	visited[start] = 0
	paths := []string{}
	shortest := 9999
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).([2]int)
		path := queue.Remove(queue.Front()).(string)
		if node == target && visited[node] <= shortest {
			shortest = visited[node]
			paths = append(paths, path+"A")
		}
		for _, dir := range moves {
			mx, my := node[0]+dir[0], node[1]+dir[1]
			next := [2]int{mx, my}
			p := path
			dist, _ := visited[next]
			if mx < 0 || my < 0 || mx == len(pad) || my == len(pad[0]) || pad[mx][my] == 0 || dist > shortest {
				continue
			}
			visited[next] = visited[node] + 1
			queue.PushBack(next)
			p += string(dir[2])
			queue.PushBack(p)
		}
	}
	fmt.Println(len(paths))
	return shortest, paths[len(paths)-1]
}

func calcComplex(pad3, line string) int {
	var str string
	var num int
	fmt.Sscanf(line, "%d%s", &num, &str)
	fmt.Println(len(pad3), num)
	return len(pad3) * num
}

func makeNumPad() [][]int {
	nPad := [][]int{}
	nPad = append(nPad, []int{'7', '8', '9'})
	nPad = append(nPad, []int{'4', '5', '6'})
	nPad = append(nPad, []int{'1', '2', '3'})
	nPad = append(nPad, []int{0, '0', 'A'})
	return nPad
}

func makeDirPad() [][]int {
	dPad := [][]int{}
	dPad = append(dPad, []int{0, '^', 'A'})
	dPad = append(dPad, []int{'<', 'v', '>'})
	return dPad
}
