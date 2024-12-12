package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

type Key struct {
	Val  string
	Iter int
}

func Run() {
	scanner := util.Scanner("day11")
	dp := make(map[Key]int)
	times := 75
	var sumPt1, sumPt2 int
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")
		for _, split := range splits {
			sumPt1 += blink(25, dp, split)
			sumPt2 += blink(times, dp, split)
		}
	}

	fmt.Printf("Part1: For 25 blinks, %d stones\n", sumPt1)
	fmt.Printf("Part2: For %d blinks, %d stones\n", times, sumPt2)
}

func blink(i int, dp map[Key]int, val string) (sum int) {
	digs := len(val)
	key := Key{Val: val, Iter: i}
	switch {
	case i == 0:
		sum++
	case dp[key] != 0:
		return dp[key]
	case val == "0":
		sum += blink(i-1, dp, "1")
	case digs%2 == 0:
		n1, n2 := val[:digs/2], reduce(val[digs/2:], 1)
		sum += blink(i-1, dp, n1)
		sum += blink(i-1, dp, n2)
	default:
		sum += blink(i-1, dp, reduce(val, 2024))
	}
	dp[key] = sum
	return sum
}

func reduce(val string, mult int) string {
	num, _ := strconv.Atoi(val)
	num *= mult
	return strconv.Itoa(num)
}
