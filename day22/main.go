package day22

import (
	"fmt"
	"strconv"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	scanner := util.Scanner("day22")
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		secret, _ := strconv.Atoi(line)
		for range 2000 {
			secret = nextSecret(secret)
		}
		sum += secret
	}
	fmt.Println(sum)
}

func nextSecret(secret int) (result int) {
	result = secret ^ secret*64
	result %= 16777216

	result ^= result / 32
	result %= 16777216

	result ^= result * 2048
	result %= 16777216
	return
}
