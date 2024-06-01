package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	ROWS_LENGTH := len(matrix)
	COLS_LENGTH := len(matrix[0])

	topIndex := 0
	bottomIndex := ROWS_LENGTH - 1

	for topIndex <= bottomIndex {
		rowIndex := (topIndex + bottomIndex) / 2
		currentRow := matrix[rowIndex]

		if target > currentRow[COLS_LENGTH-1] {
			topIndex = rowIndex + 1
			continue
		}

		if target < currentRow[0] {
			bottomIndex = rowIndex - 1
			continue
		}

		if topIndex > bottomIndex {
			return false
		}

        break
	}

	row := (topIndex + bottomIndex) / 2
	leftIndex := 0
	rightIndex := COLS_LENGTH - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2
		middleVal := matrix[row][middleIndex]

		if target == middleVal {
			return true
		}

		if target > middleVal {
			leftIndex = middleIndex + 1
			continue
		}

		rightIndex = middleIndex - 1
	}

	return false
}
