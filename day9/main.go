package day9

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
}

func solvePt1() {
	fs, size := readFs()
	moveSingleBlocks(fs, size)
	checksum := hashFs(fs, size)
	fmt.Println("Checksum:", checksum)
}

func hashFs(fs []int, size int) (sum int) {
	for i := 0; i < size; i++ {
		if fs[i] == -1 {
			return
		}
		sum += i * fs[i]
	}
	return
}

func moveSingleBlocks(fs []int, size int) {
	for i := size; i > 0; i-- {
		if fs[i-1] == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			if fs[j] == -1 {
				fs[i-1], fs[j] = fs[j], fs[i-1]
				break
			}
		}
	}
}

func readFs() (fs []int, i int) {
	reader := util.Reader("day9")
	fs = []int{}
	for ; ; i++ {
		b, err := reader.ReadByte()
		if err != nil || b == '\n' {
			return fs, len(fs)
		}

		filler := arrayFiller(i, b)
		fs = append(fs, filler...)
	}
}

func arrayFiller(i int, b byte) []int {
	size := int(b - '0')
	arr := make([]int, size)
	fill := i / 2
	if i%2 == 1 {
		fill = -1
	}
	for i := 0; i < size; i++ {
		arr[i] = fill
	}
	return arr
}
