package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day19")
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	// rgx := getRegexes(patterns)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		val := isPossible(line, patterns)
		sum += val
	}
	fmt.Println("Pt1 sol:", sum)
}

func isPossible(line string, patterns []string) int {
	cp := line
	for {
		if cp == "" {
			return 1
		}
		for _, p := range patterns {
			if strings.HasPrefix(cp, p) {
				cp = cp[len(p)-1:]
				break
			}
		}
		break
	}
	return 0
}

func isPossibleRegex(line string, rgx []*regexp.Regexp) int {
	match := make([]bool, len(line))
	for _, r := range rgx {
		regexCheck(line, r, match)
	}
	for i := range len(line) {
		if !match[i] {
			return 0
		}
	}
	fmt.Println(line)
	return 1
}

func regexCheck(line string, r *regexp.Regexp, match []bool) {
	matches := r.FindAllStringIndex(line, -1)
	for _, m := range matches {
		for i := m[0]; i < m[1]; i++ {
			match[i] = true
		}
	}
}

func getRegexes(patterns []string) (arr []*regexp.Regexp) {
	for _, p := range patterns {
		r := regexp.MustCompile(p)
		arr = append(arr, r)
	}
	return
}
