package main

import (
	"fmt"
	"math"
	"slices"
)

func dayTwoPartOne(data [][]int) {
  safeReports := 0

	for _, elements := range data {
		descending := "new"
		stayed := true
		withinJumpSize := true
		for index, element := range elements[1:] {
			currentlyDescending := element < elements[index]
			if descending == "new" {
				if currentlyDescending {
					descending = "desc"
				} else {
					descending = "asc"
				}
			}
			if descending != "new" && currentlyDescending != (descending == "desc") {
				stayed = false
			}

			diff := int(math.Abs(float64(element - elements[index])))

			if diff == 0 || diff > 3 {
				withinJumpSize = false
			}
		}

		if stayed && withinJumpSize {
			safeReports++
		}
	}

  fmt.Println("Safe Reports", safeReports)
}

func dayTwoPartTwo(data [][]int) {
	safeReports := 0
	
	for _, report := range data {
		isGood := false
		if testReport(report) {
			isGood = true
		} else {
			for i := 0; i < len(report) ; i++ {
				modReport := slices.Clone(report)
				modReport = slices.Delete(modReport, i, i + 1)
				if testReport(modReport) {
					isGood = true
					break
				}
			}
		}
		if isGood {
			fmt.Println("This one is good!", report)
			safeReports++
		} else {
			fmt.Println("This one NOT is good", report)

		}
	}

  fmt.Println("Damper Safe Reports", safeReports)
}

func testReport(report []int) (bool) {
	lastDifference :=  report[0] - report[1] 
	if math.Abs(float64(lastDifference)) > 3 || lastDifference == 0 {
		return false
	}
	lastDesc := lastDifference > 0
	for i := 1; i < len(report); i++ {
		currentDifference := report[i - 1] - report[i]
		currentDesc := currentDifference > 0

		if currentDesc != lastDesc || currentDifference == 0 ||  math.Abs(float64(currentDifference)) > 3 {
			return false
		} else {
			lastDesc = currentDesc
		}
	}
	return true
}

func dayTwoMain() {
	data := dayTwoData()
	dayTwoPartOne(data)
  dayTwoPartTwo(data)
}