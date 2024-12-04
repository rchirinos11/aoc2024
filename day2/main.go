package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
	solvePt2()
}

func solvePt1() {
	scanner := util.Scanner("day2")
	var safe int
	for scanner.Scan() {
		nums, size := getNums(scanner.Text())
		if isSafe(nums, size) {
			safe++
		}
	}
	fmt.Println("Pt1 safe:", safe)
}

func solvePt2() {
	scanner := util.Scanner("day2")
	var safe int
	for scanner.Scan() {
		nums, size := getNums(scanner.Text())
		for i := 0; i < size; i++ {
			if isSafe(makeCopy(nums, i, size)) {
				safe++
				break
			}
		}
	}
	fmt.Println("Pt2 safe:", safe)
}

func getNums(text string) ([]int, int) {
	splits := strings.Split(text, " ")
	size := len(splits)
	nums := make([]int, size)
	for i, s := range splits {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums, size
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

func makeCopy(nums []int, index, size int) ([]int, int) {
	cp := make([]int, size)
	copy(cp, nums)
	return append(cp[:index], cp[index+1:]...), size - 1
}
