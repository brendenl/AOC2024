package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func dayThreePartOne(data string) {
	re := regexp.MustCompile(`mul\((?P<X>\d+)\,(?P<Y>\d+)\)`)
	foundRes := re.FindAllStringSubmatch(data, -1)
	total := 0

	for _, found := range foundRes {
		x, err := strconv.Atoi(found[1])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(found[2])
		if err != nil {
			continue
		}
		total += x * y
	}
	fmt.Println("total of nums is", total)
}

func dayThreePartTwo(data string) {
	re := regexp.MustCompile(`mul\((?P<X>\d+)\,(?P<Y>\d+)\)|do\(\)|don't\(\)`)
	sendIt := true
	foundRes := re.FindAllStringSubmatch(data, -1)
	total := 0

	for _, found := range foundRes {
		if found[0] == "don't()" {
			sendIt = false
		} else if found[0] == "do()" {
			sendIt = true
		}

		if !sendIt {
			continue
		}
		x, err := strconv.Atoi(found[1])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(found[2])
		if err != nil {
			continue
		}
		total += x * y
	}
	fmt.Println("total of nums is", total)
}

func dayThreeMain() {
	data := dayThreeData()
	dayThreePartOne(data)
	dayThreePartTwo(data)
}