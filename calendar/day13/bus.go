package main

import (
	"strconv"
	"strings"
)

type bus struct {
	idx int
	id  int
}

func loadBusses(s string) []bus {
	var busses []bus

	for idx, strId := range strings.Split(s, ",") {
		if strId == "x" {
			continue
		}

		id, _ := strconv.Atoi(strId)
		busses = append(busses, bus{idx: idx, id: id})
	}
	return busses
}

func getEarliestBusRide(earliestDepartmentTime int, busses []bus) (firstBus bus, idleTime int) {
	shortestIdleTime := 10000000

	for _, bus := range busses {
		if bus.id == -1 {
			continue
		}
		nextDepartmentTime := earliestDepartmentTime + (bus.id - (earliestDepartmentTime % bus.id))
		idleTime := nextDepartmentTime - earliestDepartmentTime
		if idleTime < shortestIdleTime {
			shortestIdleTime = idleTime
			firstBus = bus
		}
	}

	return firstBus, shortestIdleTime
}

func getFirstCommonTimestamp(busses []bus) int {
	firstCommonTimestamp := 0

	step := 1
	for _, bus := range busses {
		for (firstCommonTimestamp+bus.idx)%bus.id != 0 {
			firstCommonTimestamp += step
		}
		step *= bus.id
	}
	return firstCommonTimestamp
}
