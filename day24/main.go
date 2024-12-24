package day24

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input, _ := os.ReadFile("day24/input")
	section := strings.Split(string(input), "\n\n")
	lines := strings.Split(section[0], "\n")

	values := make(map[string]int)
	for _, line := range lines {
		splits := strings.Split(line, ": ")
		val, _ := strconv.Atoi(splits[1])
		values[splits[0]] = val
	}

	lines = strings.Split(section[1], "\n")
	for len(lines) > 1 {
		for i, line := range lines {
			splits := strings.Split(line, " ")

			result := operate(splits, values)
			if result == -1 {
				continue
			}
			cp := lines[:i]
			lines = append(cp, lines[i+1:]...)
			values[splits[4]] = result
		}
	}

	var sum int
	for i := 0; ; i++ {
		key := "z"
		if i < 10 {
			key += "0"
		}

		val, ok := values[key+strconv.Itoa(i)]
		if !ok {
			break
		}

		val = val << i
		sum += val
	}
	fmt.Printf("%b\n", sum)
	fmt.Println(sum)
}

func operate(splits []string, values map[string]int) int {
	if _, ok := values[splits[0]]; !ok || len(splits) != 5 {
		return -1
	}
	if _, ok := values[splits[2]]; !ok {
		return -1
	}
	switch splits[1] {
	case "AND":
		return values[splits[0]] & values[splits[2]]
	case "XOR":
		return values[splits[0]] ^ values[splits[2]]
	default:
		return values[splits[0]] | values[splits[2]]
	}
}
