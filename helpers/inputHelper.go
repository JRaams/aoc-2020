package helpers

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// GetInputValues reads the file at the specified input path, strips the last line (if needed) and returns the content as string array
func GetInputValues(absFilePath string) (values []string) {
	txt, err := ioutil.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("Input file '%s' not found...", absFilePath))
	}

	strValues := strings.Split(string(txt), "\n")

	// Remove empty last line
	if len(strValues[len(strValues)-1]) == 0 {
		strValues = strValues[:len(strValues)-1]
	}

	return strValues
}

// MustAtoi converts a string that is SURE to be an int to an int without error
func MustAtoi(input string) int {
	intVal, _ := strconv.Atoi(input)
	return intVal
}
