package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day17")
	registers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			scanner.Scan()
			break
		}
		var value int
		fmt.Sscanf(strings.Split(line, ":")[1], " %d", &value)
		registers = append(registers, value)
	}
	vals := strings.Split(strings.Split(scanner.Text(), " ")[1], ",")
	output := operate(vals, registers)
	for i, num := range output {
		fmt.Print(num)
		if i < len(output)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

func operate(vals []string, registers []int) []int {
	output := []int{}
	for i := 0; i < len(vals); {
		opcode, _ := strconv.Atoi(vals[i])
		operand, _ := strconv.Atoi(vals[i+1])
		switch opcode {
		case 0:
			registers[0] /= pow(2, getCombo(operand, registers))
		case 1:
			registers[1] ^= operand
		case 2:
			registers[1] = getCombo(operand, registers) % 8
		case 3:
			if registers[0] == 0 {
				break
			}
			i = operand
			continue
		case 4:
			registers[1] ^= registers[2]
		case 5:
			output = append(output, getCombo(operand, registers)%8)
		case 6:
			registers[1] = registers[0] / pow(2, getCombo(operand, registers))
		case 7:
			registers[2] = registers[0] / pow(2, getCombo(operand, registers))
		}
		i += 2
	}
	return output
}

func pow(val, exponent int) int {
	mult := 1
	for range exponent {
		mult *= val
	}
	return mult
}

func getCombo(operand int, registers []int) int {
	switch operand {
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	default:
		return operand
	}
}
