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
				maskedValueStr := getMaskedValueStr(instruction_set.mask, instruction.value, false)
				maskedValue, _ := strconv.ParseInt(maskedValueStr, 2, 64)
				p.memory[instruction.address] = int(maskedValue)
			} else if version == 2 {
				maskedFloatingAddress := getMaskedValueStr(instruction_set.mask, instruction.address, true)
				addressCombinations := getAddressCombinations("", maskedFloatingAddress, fmt.Sprintf("%036b", instruction.address))
				for _, addresCombi := range addressCombinations {
					intAddr, _ := strconv.ParseInt(addresCombi, 2, 64)
					p.memory[int(intAddr)] = instruction.value
				}
			}
		}
	}
}

func getMaskedValueStr(mask string, intValue int, isAddress bool) string {
	if isAddress {
		base2 := strconv.FormatInt(int64(intValue), 2)
		paddedBase2Value := fmt.Sprintf("%036s", base2)
		result := []rune(paddedBase2Value)

		for i := 0; i < len(mask); i++ {
			char := string(mask[i])
			switch string(char) {
			case "0":
				{
					break
				}
			case "1":
				{
					result[i] = '1'
					break
				}
			case "X":
				{
					if isAddress {
						result[i] = 'X'
					} else {
						result[i] = rune(paddedBase2Value[i])
					}
					break
				}
			default:
				{
					panic(fmt.Sprintf("Unknown mask value: %s", char))
				}
			}
		}

		return string(result)
	} else {
		base2 := strconv.FormatInt(int64(intValue), 2)
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
		return result
	}
}

func getAddressCombinations(mask, remaining, addr string) []string {
	if len(remaining) == 0 {
		return []string{mask}
	}

	switch remaining[0] {
	case '0':
		return getAddressCombinations(mask+string(addr[len(mask)]), remaining[1:], addr)
	case '1':
		return getAddressCombinations(mask+"1", remaining[1:], addr)
	case 'X':
		return append(getAddressCombinations(mask+"0", remaining[1:], addr), getAddressCombinations(mask+"1", remaining[1:], addr)...)
	default:
		panic(fmt.Errorf("Unknown bitmask type %s", string(remaining[0])))
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
