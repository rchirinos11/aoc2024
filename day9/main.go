package day9

import (
	"fmt"

	"github.com/rchirinos11/aoc2024/util"
)

func Run() {
	solvePt1()
	solvePt2()
}

func solvePt2() {
	fs, size := readFs()
	moveWholeFileBlocks(fs, size)
	checksum := hashFs(fs)
	fmt.Println("Checksum:", checksum)
}

func solvePt1() {
	fs, size := readFs()
	moveSingleBlocks(fs, size)
	checksum := hashFs(fs)
	fmt.Println("Checksum:", checksum)
}

func hashFs(fs []int) (sum int) {
	for i, val := range fs {
		if val == -1 {
			continue
		}
		sum += i * val
	}
	return
}

func moveWholeFileBlocks(fs []int, size int) {
	blocksize := 0
	last := fs[size-1]
	for i := size - 1; i > 0; i-- {
		if last != fs[i] {
			moveFile(fs, blocksize, i)
			blocksize = 0
		}
		blocksize++
		last = fs[i]
	}
}

func moveFile(fs []int, blocksize, x int) {
	if blocksize == 0 {
		return
	}
	space := 0
	for i := range x {
		if fs[i] != -1 {
			space = 0
			continue
		}
		if space++; space == blocksize {
			for j := 1; j <= blocksize; j++ {
				fs[i-space+j], fs[x+j] = fs[x+j], fs[i-space+j]
			}
			return
		}
	}
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
