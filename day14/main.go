package day14

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

const (
	WIDTH   = 101
	HEIGHT  = 103
	SECONDS = 100
)

var (
	xlim = WIDTH / 2
	ylim = HEIGHT / 2
)

func Run() {
	scanner := util.Scanner("day14")
	robots := [][4]int{}
	for scanner.Scan() {
		var pX, pY, vX, vY int
		line := scanner.Text()
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY)
		r := [4]int{pY, pX, vY, vX}
		for i := 0; i < SECONDS; i++ {
			r = moveBySecond(r)
		}
		robots = append(robots, r)
	}
	printGrid(robots)
}

func printGrid(robots [][4]int) {
	var q1, q2, q3, q4 int
	for i := range HEIGHT {
		for j := range WIDTH {
			val := 0
			for _, robot := range robots {
				if robot[0] == i && robot[1] == j {
					val++
					switch {
					case i < ylim && j < xlim:
						q1++
					case i < ylim && j > xlim:
						q2++
					case i > ylim && j < xlim:
						q3++
					case i > ylim && j > xlim:
						q4++
					}
				}
			}
			if val == 0 {
				fmt.Printf("%1c", '.')
			} else {
				fmt.Printf("%1d", val)
			}
		}
		fmt.Println()
	}
	fmt.Println("Q1:", q1, "Q2:", q2, "Q3:", q3, "Q4:", q4)
	fmt.Println("Safety factor:", q1*q2*q3*q4)
}

func moveBySecond(robot [4]int) [4]int {
	if robot[0]+robot[2] < 0 {
		robot[0] += HEIGHT + robot[2]
	} else if robot[0]+robot[2] >= HEIGHT {
		robot[0] += robot[2] - HEIGHT
	} else {
		robot[0] += robot[2]
	}
	if robot[1]+robot[3] < 0 {
		robot[1] += WIDTH + robot[3]
	} else if robot[1]+robot[3] >= WIDTH {
		robot[1] += robot[3] - WIDTH
	} else {
		robot[1] += robot[3]
	}
	return robot
}
