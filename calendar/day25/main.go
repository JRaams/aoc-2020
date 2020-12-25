package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/jraams/aoc-2020/helpers"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	publicKeys := helpers.TranslateStringArrToIntArr(lines)

	a := solveA(publicKeys)
	fmt.Printf("Solution day 25 part a: %d\n", a)
}

func findLoopSize(publicKey int, subject int) int {
	loopSize := 0
	for key := 1; key != publicKey; loopSize++ {
		key *= subject
		key = key % 20201227
	}
	return loopSize
}

func transformSubject(subject int, loopSize int) int {
	encryptionKey := 1
	for i := 0; i < loopSize; i++ {
		encryptionKey *= subject
		encryptionKey = encryptionKey % 20201227
	}
	return encryptionKey
}

func solveA(publicKeys []int) int {
	defer helpers.Measure(time.Now(), "")
	cardPU, doorPU := publicKeys[0], publicKeys[1]
	doorLS := findLoopSize(cardPU, 7)
	a := transformSubject(doorPU, doorLS)
	return a
}
