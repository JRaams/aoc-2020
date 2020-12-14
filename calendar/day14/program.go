package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type program struct {
	memory           map[int]int
	instruction_sets []instruction_set
}

func (p *program) run(version int) {
	for _, instruction_set := range p.instruction_sets {
		for _, instruction := range instruction_set.instructions {
			if version == 1 {
				p.memory[instruction.address] = instruction.getMaskedValue(instruction_set.mask)
			} else if version == 2 {
				maskedFloatingAddress := getMaskedFloatingAddress(instruction_set.mask, instruction.address)
				addressCombinations := getAddressCombinations("", maskedFloatingAddress, fmt.Sprintf("%036b", instruction.address))
				for _, addresCombi := range addressCombinations {
					intAddr, _ := strconv.ParseInt(addresCombi, 2, 64)
					p.memory[int(intAddr)] = instruction.value
				}
			}
		}
	}
}

func (p *program) getSumOfAllValues() int {
	sum := 0
	for _, v := range p.memory {
		sum += v
	}
	return sum
}

func loadProgram(inputValues []string) program {
	isets := []instruction_set{}

	var iset instruction_set
	for _, line := range inputValues {
		lineParts := strings.Split(line, " = ")

		if funk.Contains(line, "mask") {
			isets = append(isets, iset)
			iset = instruction_set{
				mask: lineParts[1],
			}
		} else {
			address, _ := strconv.Atoi(strings.Replace(strings.Replace(lineParts[0], "mem[", "", 1), "]", "", 1))
			value, _ := strconv.Atoi(lineParts[1])
			instruction := instruction{
				address: address,
				value:   value,
			}
			iset.instructions = append(iset.instructions, instruction)
		}
	}
	isets = append(isets, iset)
	isets = isets[1:]

	program := program{
		memory:           make(map[int]int),
		instruction_sets: isets,
	}
	return program
}
