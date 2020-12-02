package helpers

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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
