package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
}

func solvePt1() {
	scanner := util.Scanner("day7")
	sum := 0
	for scanner.Scan() {
		result, nums := readLine(scanner.Text())
		sum += addCalibrations(result, nums)
	}
	fmt.Println("Pt1 calibration result:", sum)
}

func readLine(line string) (result int, nums []int) {
	splits := strings.Split(line, " ")
	fmt.Sscanf(splits[0], "%d:", &result)
	for i := 1; i < len(splits); i++ {
		n, _ := strconv.Atoi(splits[i])
		nums = append(nums, n)
	}
	return
}

func addCalibrations(result int, nums []int) int {
	truthTable := 1 << (len(nums) - 1)
	for i := 0; i < truthTable; i++ {
		if x := calculate(i, nums); x == result {
			return x
		}
	}
	return 0
}

func calculate(operands int, nums []int) int {
	calc := nums[0]
	for i := 1; i < len(nums); i++ {
		if operands>>(i-1)&1 == 1 {
			calc *= nums[i]
			continue
		}

		calc += nums[i]
	}
	return calc
}
