package main

import (
	"strconv"
	"strings"
)

type gameconsole struct {
	instructions          []instruction
	accumulator           int
	currentInstructionIdx int
}

func (original gameconsole) clone() gameconsole {
	var clone gameconsole
	clone.accumulator = original.accumulator
	clone.currentInstructionIdx = original.currentInstructionIdx
	clone.instructions = append(clone.instructions, original.instructions...)
	return clone
}

type instruction struct {
	operation string
	argument  int
	called    bool
}

func loadInstructions(inputValues []string) gameconsole {
	console := gameconsole{
		instructions: make([]instruction, 0),
	}

	for _, line := range inputValues {
		lineParts := strings.Split(line, " ")
		operation := lineParts[0]
		argument, _ := strconv.Atoi(strings.Replace(lineParts[1], "+", "", 1))

		instruction := instruction{
			operation: operation,
			argument:  argument,
		}
		console.instructions = append(console.instructions, instruction)
	}

	return console
}

func runInstructions(console gameconsole) (isCorrupted bool, accumulator int) {
	for {
		currentInstruction := &console.instructions[console.currentInstructionIdx]
		if currentInstruction.called {
			return true, console.accumulator
		}
		currentInstruction.called = true

		switch currentInstruction.operation {
		case "nop":
			{
				console.currentInstructionIdx += 1
				break
			}
		case "acc":
			{
				console.currentInstructionIdx += 1
				console.accumulator += currentInstruction.argument
				break
			}
		case "jmp":
			{
				console.currentInstructionIdx += currentInstruction.argument
				break
			}
		}

		if console.currentInstructionIdx == len(console.instructions) {
			return false, console.accumulator
		}
	}
}

func fixInstructions(console gameconsole) (isFixed bool, accumulator int) {
	for i := 0; i < len(console.instructions); i++ {
		clonedConsole := console.clone()

		// Switch nop and jmp
		oldOperation := clonedConsole.instructions[i].operation
		if oldOperation == "acc" {
			continue
		} else if oldOperation == "nop" {
			clonedConsole.instructions[i].operation = "jmp"
		} else if oldOperation == "jmp" {
			clonedConsole.instructions[i].operation = "nop"
		}

		// Run instructions with a single "nop" switched with "jmp"
		isCorrupted, accumulator := runInstructions(clonedConsole)
		if !isCorrupted {
			return true, accumulator
		}
	}

	return false, 0
}
