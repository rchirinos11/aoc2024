package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day2")
	var safe int
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		size := len(splits)
		nums := make([]int, size)
		for i, s := range splits {
			nums[i], _ = strconv.Atoi(s)
		}
		if isSafe(nums, size) {
			safe++
		}
	}
	fmt.Println("Safe:", safe)
}

func isSafe(nums []int, size int) bool {
	desc := nums[0] > nums[1]
	for i := 1; i < size; i++ {
		if unsafeRange(nums[i-1], nums[i], desc) {
			return false
		}
	}
	return true
}

func unsafeRange(n1, n2 int, desc bool) bool {
	if desc {
		return n1-n2 < 1 || n1-n2 > 3
	}
	return n2-n1 < 1 || n2-n1 > 3
}
