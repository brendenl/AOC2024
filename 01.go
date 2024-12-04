package main

import (
	"fmt"
	"math"
	"slices"
)

func partOne(left []int, right []int) {
    slices.Sort(left)
    slices.Sort(right)

    totalDistance := 0

    for index, element := range left {
        distance := element - right[index]
        totalDistance += int(math.Abs(float64(distance)))
    }

    fmt.Println("total distance", totalDistance)
}

func partTwo(left []int, right []int) {
    rightMap := make(map[int]int)
    for _, element := range right {
        rightMap[element] = rightMap[element] + 1
    }

	similarity := 0

    for _, element := range left {
        similarity += element * rightMap[element]
    }

    fmt.Println("Similarity Score", similarity)
}

func dayOneMain() {
    left, right := dayOneData()
    partOne(left, right)
    partTwo(left, right)
}