package day23

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day23")
	network := make(map[string][]string)
	for scanner.Scan() {
		pcs := strings.Split(scanner.Text(), "-")
		network[pcs[0]] = append(network[pcs[0]], pcs[1])
		network[pcs[1]] = append(network[pcs[1]], pcs[0])
	}

	set := make(map[string]bool)
	var sum int
	for pc1, val := range network {
		for _, pc2 := range val {
			for _, pc3 := range network[pc2] {
				if contains(network, pc1, pc3) {
					add := []string{pc1, pc2, pc3}
					sort.Strings(add)
					n := strings.Join(add, ",")
					if set[n] {
						continue
					}
					if pc1[0] == 't' || pc2[0] == 't' || pc3[0] == 't' {
						sum++
					}
					set[n] = true
				}
			}
		}
	}

	fmt.Println("Pt1:", sum)
}

func contains(network map[string][]string, eval, pc string) bool {
	for _, pc2 := range network[pc] {
		if eval == pc2 {
			return true
		}
	}
	return false
}
