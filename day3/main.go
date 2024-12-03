package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
	solvePt2()
}

func solvePt1() {
	var sum int
	scanner := util.Scanner("day3")
	rx := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := rx.FindAllString(line, -1)
		for _, str := range matches {
			sum += multiply(str)
		}
	}
	fmt.Println("Pt1 sum is:", sum)
}

func solvePt2() {
	active := true
	var sum int
	scanner := util.Scanner("day3")
	rx := regexp.MustCompile(`(mul\([\d]{1,3},[\d]{1,3}\))|(don't\(\))|(do\(\))`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := rx.FindAllString(line, -1)
		for _, str := range matches {
			active = !(str == "don't()") && (!(str == "do()") && active || (str == "do()"))
			if active && str != "do()" {
				sum += multiply(str)
			}
		}
	}
	fmt.Println("Pt2 sum is:", sum)
}

func multiply(mul string) int {
	rx := regexp.MustCompile(`[\d]{1,3}`)
	numsStr := rx.FindAllString(mul, -1)
	n1, _ := strconv.Atoi(numsStr[0])
	n2, _ := strconv.Atoi(numsStr[1])
	return n1 * n2
}
