package main

import (
	"fmt"
	"strings"
)

func dayFourPartOne(data string) {
	rows := strings.Split(data, "\n")
	var grid [][]string = make([][]string, len(rows))
	for i, row := range rows {
		grid[i] = strings.Split(row, "")
	}
	totalXmas := 0
	for i, _ := range grid {
		for j, col := range grid[i] {
			if col == "X" {
				if checkHorizontal(grid, i, j) {
					totalXmas++
				}
				if checkBackwards(grid, i, j){
					totalXmas++
				}
				if checkVertical(grid, i, j){
					totalXmas++
				}
				if checkVerticalBackwards(grid, i, j) {
					totalXmas++
				}
				if checkDiagonalNE(grid, i, j){
					totalXmas++
				}
				if checkDiagonalNW(grid, i, j){
					totalXmas++
				}
				if checkDiagonalSE(grid, i, j){
					totalXmas++
				}
				if checkDiagonalSW(grid, i, j){
					totalXmas++
				}
			}
		}
	}

	fmt.Println("totalXmas", totalXmas)
}

func checkHorizontal(grid [][]string, i int, j int) (bool) {
	if j + 3 >= len(grid[i]) {
		return false
	}
	return (grid [i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S")
}

func checkBackwards(grid [][]string, i int, j int) (bool) {
	if j - 3 < 0 {
		return false
	}
	return (grid [i][j-1] == "M" && grid[i][j-2] == "A" && grid[i][j-3] == "S")
}

func checkVertical(grid [][]string, i int, j int) (bool) {
	if i + 3 >= len(grid) {
		return false
	}
	return (grid [i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S")
}

func checkVerticalBackwards(grid [][]string, i int, j int) (bool) {
	if i - 3 < 0 {
		return false
	}
	return (grid [i-1][j] == "M" && grid[i-2][j] == "A" && grid[i-3][j] == "S")
}

func checkDiagonalSE(grid [][]string, i int, j int) (bool) {
	if j + 3 >= len(grid[i]) || i + 3 >= len(grid){
		return false
	}
	return (grid [i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S")
}

func checkDiagonalNE(grid [][]string, i int, j int) (bool) {
	if j - 3 < 0 || i + 3 >= len(grid){
		return false
	}
	return (grid [i+1][j-1] == "M" && grid[i+2][j-2] == "A" && grid[i+3][j-3] == "S")
}

func checkDiagonalNW(grid [][]string, i int, j int) (bool) {
	if j - 3 < 0 || i - 3 < 0{
		return false
	}
	return (grid [i-1][j-1] == "M" && grid[i-2][j-2] == "A" && grid[i-3][j-3] == "S")
}

func checkDiagonalSW(grid [][]string, i int, j int) (bool) {
	if j + 3 >= len(grid[i]) || i - 3 < 0{
		return false
	}
	return (grid [i-1][j+1] == "M" && grid[i-2][j+2] == "A" && grid[i-3][j+3] == "S")
}

func dayFourPartTwo(data string) {
	rows := strings.Split(data, "\n")
	var grid [][]string = make([][]string, len(rows))
	for i, row := range rows {
		grid[i] = strings.Split(row, "")
	}
	totalXmas := 0
	for i, _ := range grid {
		for j, col := range grid[i] {
			if col == "A" {
				numMas := 0
				if checkSE(grid, i, j) {
					numMas++
				}
				if checkSW(grid, i, j) {
					numMas++
				}
				if checkNE(grid, i, j) {
					numMas++
				}
				if checkNW(grid, i, j) {
					numMas++
				}
				if numMas > 1 {
					totalXmas++
				}
			}
		}
	}
	fmt.Println("total X-MAS", totalXmas)
}

func checkSE(grid [][]string, i int, j int) (bool) {
	if j - 1 < 0 || i - 1 < 0 || i + 1 > len(grid) - 1 || j + 1 > len(grid[0]) - 1 {
		return false
	}
	return grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S"
}

func checkSW(grid [][]string, i int, j int) (bool) {
	if j - 1 < 0 || i - 1 < 0 || i + 1 > len(grid) - 1 || j + 1 > len(grid[0]) - 1 {
		return false
	}
	return grid[i+1][j-1] == "M" && grid[i-1][j+1] == "S"
}

func checkNE(grid [][]string, i int, j int) (bool) {
	if j - 1 < 0 || i - 1 < 0 || i + 1 > len(grid) - 1 || j + 1 > len(grid[0]) - 1 {
		return false
	}
	return grid[i-1][j+1] == "M" && grid[i+1][j-1] == "S"
}

func checkNW(grid [][]string, i int, j int) (bool) {
	if j - 1 < 0 || i - 1 < 0 || i + 1 > len(grid) - 1 || j + 1 > len(grid[0]) - 1 {
		return false
	}
	return grid[i+1][j+1] == "M" && grid[i-1][j-1] == "S"
}

func main() {
	data := dayFourData()
	dayFourPartOne(data)
	dayFourPartTwo(data)
}