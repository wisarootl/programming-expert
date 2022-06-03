package main

import "fmt"

// magic square (each sum row = each sum col = sun diagonal)
//    -------------
//    | 2 | 7 | 6 |  15
//    -------------
//    | 9 | 5 | 1 |  15
//    -------------
//    | 4 | 3 | 8 |  15
//    -------------
//  15  15  15  15   15


func sumSlice(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func DetectMagicSquare(square [][]int) bool {
	desiredSum := sumSlice(square[0])
	columnSums := make([]int, len(square))
	diagonalLeftRightSum := 0
	diagonalRightLeftSum := 0

	for row, rowNums := range square {
		rowSum := 0

		for col, num := range rowNums {
			rowSum += num
			columnSums[col] += num

			if row == col {
				diagonalLeftRightSum += num
			}
			if col == len(square)-1-row {
				diagonalRightLeftSum += num
			}

			if row == len(square)-1 && columnSums[col] != desiredSum {
				return false
			}
		}

		if rowSum != desiredSum {
			return false
		}
	}
	return diagonalRightLeftSum == desiredSum && diagonalLeftRightSum == desiredSum
}

func main() {
	square := [][]int{{5, 6, 19, 68}, {69, 18, 3, 8}, {4, 7, 70, 17}, {20, 67, 6, 5}}
	fmt.Println(DetectMagicSquare(square))
}