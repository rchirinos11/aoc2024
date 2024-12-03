package util

import (
	"bufio"
	"os"
)

func Reader(pack string) *bufio.Reader {
	return bufio.NewReader(getFile(pack))
}

func Scanner(pack string) *bufio.Scanner {
	return bufio.NewScanner(getFile(pack))
}

func getFile(pack string) *os.File {
	file, err := os.Open(pack + "/input")
	if err != nil {
		panic(err)
	}
	return file
}
