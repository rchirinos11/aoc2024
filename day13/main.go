package day13

import (
	"bufio"
	"fmt"
	"math"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	var solved1, token1, solved2, token2 int
	scanner := util.Scanner("day13")
	offset := 10000000000000.0

	for scanner.Scan() {
		lines := readLines(scanner)
		solve(lines, 0, &solved1, &token1)
		solve(lines, offset, &solved2, &token2)
	}
	fmt.Println("Pt1 Prizes won:", solved1, "Tokens spent:", token1)
	fmt.Println("Pt2 Prizes won:", solved2, "Tokens spent:", token2)
}

func readLines(scanner *bufio.Scanner) [3]string {
	lineA := scanner.Text()
	scanner.Scan()
	lineB := scanner.Text()
	scanner.Scan()
	lineP := scanner.Text()
	scanner.Scan()
	return [3]string{lineA, lineB, lineP}
}

func solve(lines [3]string, offset float64, solved, tokens *int) {
	var aX, bX, aY, bY, pX, pY float64
	fmt.Sscanf(lines[0], "Button A: X+%f, Y+%f", &aX, &aY)
	fmt.Sscanf(lines[1], "Button B: X+%f, Y+%f", &bX, &bY)
	fmt.Sscanf(lines[2], "Prize: X=%f, Y=%f", &pX, &pY)
	pX += offset
	pY += offset
	valB := math.Round((pY - aY*pX/aX) / (bY - aY*bX/aX))
	valA := math.Round((pX - bX*valB) / aX)

	if pX != valA*aX+valB*bX || pY != valA*aY+valB*bY {
		return
	}

	*solved++
	*tokens += int(valA)*3 + int(valB)
}
