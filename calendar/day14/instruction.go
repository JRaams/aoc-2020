package main

import (
	"fmt"
	"strconv"
)

type instruction_set struct {
	mask         string
	instructions []instruction
}

type instruction struct {
	address int
	value   int
}

func (i *instruction) getMaskedValue(mask string) int {
	base2 := strconv.FormatInt(int64(i.value), 2)
	paddedBase2Value := fmt.Sprintf("%036s", base2)
	result := ""

	for i := 0; i < len(mask); i++ {
		char := string(mask[i])
		switch string(char) {
		case "0":
			{
				result += "0"
				break
			}
		case "1":
			{
				result += "1"
				break
			}
		case "X":
			{
				result += string(paddedBase2Value[i])
				break
			}
		default:
			{
				panic(fmt.Sprintf("Unknown mask value: %s", char))
			}
		}
	}

	base10, _ := strconv.ParseInt(result, 2, 64)
	return int(base10)
}
