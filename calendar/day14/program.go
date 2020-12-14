package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type program struct {
	memory          map[int]int
	instructionSets []instructionSet
}

func (p *program) run(version int) {
	for _, instructionSet := range p.instructionSets {
		for _, instruction := range instructionSet.instructions {
			if version == 1 {
				p.memory[instruction.address] = instruction.getMaskedValue(instructionSet.mask)
			} else if version == 2 {
				maskedFloatingAddress := getMaskedFloatingAddress(instructionSet.mask, instruction.address)
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
	isets := []instructionSet{}

	var iset instructionSet
	for _, line := range inputValues {
		lineParts := strings.Split(line, " = ")

		if funk.Contains(line, "mask") {
			isets = append(isets, iset)
			iset = instructionSet{
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
		memory:          make(map[int]int),
		instructionSets: isets,
	}
	return program
}
