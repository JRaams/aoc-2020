package main

type instruction_set struct {
	mask         string
	instructions []instruction
}

type instruction struct {
	address int
	value   int
}
