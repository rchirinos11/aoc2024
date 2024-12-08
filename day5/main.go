package day5

import (
	"bufio"
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
	scanner := util.Scanner("day5")
	rules := readRules(scanner)

	var sum int
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		ocurrences, lineNums := readUpdate(line)
		if correct, _, _ := evaluateUpdate(lineNums, ocurrences, rules); correct {
			mid := lineNums[len(lineNums)/2]
			sum += mid
		}
	}

	fmt.Println("Pt1 sum is:", sum)
}

func solvePt2() {
	scanner := util.Scanner("day5")
	rules := readRules(scanner)

	var sum int
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		ocurrences, lineNums := readUpdate(line)
		if correct, posEval, errVal := evaluateUpdate(lineNums, ocurrences, rules); !correct {
			sum += swap(lineNums, rules, ocurrences, posEval, errVal)
		}
	}

	fmt.Println("Pt2 sum is:", sum)
}

func swap(nums []int, rules map[int][]int, ocurrences map[int]int, posEval, errVal int) int {
	for correct := false; !correct; correct, posEval, errVal = evaluateUpdate(nums, ocurrences, rules) {
		cp := ocurrences[errVal]
		if cp == -1 {
			cp = 0
		}

		ocurrences[errVal], ocurrences[nums[posEval]] = ocurrences[nums[posEval]], ocurrences[errVal]
		nums[cp], nums[posEval] = nums[posEval], nums[cp]
	}
	return nums[len(nums)/2]
}

func copyMap(ocurrences map[int]int) map[int]int {
	cp := make(map[int]int)
	for key, val := range ocurrences {
		cp[key] = val
	}
	return cp
}

func readRules(scanner *bufio.Scanner) (rules map[int][]int) {
	rules = make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var X, Y int
		fmt.Sscanf(line, "%d|%d", &X, &Y)
		rules[X] = append(rules[X], Y)
	}
	return
}

func readUpdate(line string) (ocurrences map[int]int, lineNums []int) {
	ocurrences = make(map[int]int)
	for i, val := range strings.Split(line, ",") {
		Y, _ := strconv.Atoi(val)
		lineNums = append(lineNums, Y)
		if i == 0 {
			ocurrences[Y] = -1
			continue
		}
		ocurrences[Y] = i
	}
	return
}

func evaluateUpdate(lineNums []int, ocurrences map[int]int, rules map[int][]int) (bool, int, int) {
	for i, val := range lineNums {
		require := rules[val]
		for _, r := range require {
			if ocurrences[r] < i && ocurrences[r] != 0 {
				return false, i, r
			}
		}
	}
	return true, 0, 0
}
