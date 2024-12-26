package day25

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day25")
	var keys, locks [][]int
	var fit int
	for {
		item := []string{}
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			item = append(item, line)
		}

		if len(item) == 0 {
			break
		}
		cols := getColumns(item)
		if item[0][0] == '#' {
			locks = append(locks, cols)
			fit += overlap(keys, cols)
			continue
		}
		keys = append(keys, cols)
		fit += overlap(locks, cols)
	}
	fmt.Println("Unique pairs:", fit)
}

func getColumns(item []string) (cols []int) {
	for i := range item[0] {
		size := 0
		for j := range item {
			if item[j][i] == '#' {
				size++
			}
		}
		cols = append(cols, size-1)
	}
	return
}

func overlap(group [][]int, added []int) (fit int) {
	for _, single := range group {
		overlap := false
		for i := range single {
			if single[i]+added[i] > 5 {
				overlap = true
				break
			}
		}
		if !overlap {
			fit++
		}
	}
	return
}
