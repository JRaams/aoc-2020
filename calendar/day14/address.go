package main

import (
	"fmt"
	"strconv"
)

func getMaskedFloatingAddress(mask string, intValue int) string {
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
				result[i] = 'X'
				break
			}
		default:
			{
				panic(fmt.Sprintf("Unknown mask value: %s", char))
			}
		}
	}

	return string(result)
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
