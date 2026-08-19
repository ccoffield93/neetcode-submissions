func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	// start in the middle 
	row := m/2
	leftRowPtr := 0
	rightRowPtr := m-1
	for leftRowPtr <= rightRowPtr {
		// check first and last values of this row
		first := matrix[row][0]
		last := matrix[row][n-1]

		if target == first || target == last {
			// stop processing, we're done
			return true
		}

		// if target is between first and last, it has to be
		// in this row, or not in the matrix at all 
		if target >= first && target <= last {
			// SECOND binary search 
			leftColumnPtr := 0
			rightColumnPtr := n -1
			column := n/2

			for leftColumnPtr <= rightColumnPtr {
				val := matrix[row][column]

				if val == target {
					return true
				} else if val < target {
					leftColumnPtr = column + 1
					column = (leftColumnPtr + rightColumnPtr) / 2
				} else {
					rightColumnPtr = column - 1
					column = (leftColumnPtr + rightColumnPtr) / 2
				}
			}

			// it HAD to be in this row, and we didn't find it
			// so it isn't here
			return false

		} else if target < first {
			// binary search on the "left"
			rightRowPtr = row - 1 
			row = (leftRowPtr + rightRowPtr )/2
		} else {
			// binary search on the "right"
			leftRowPtr = row + 1 
			row = (leftRowPtr + rightRowPtr )/2
		}
	}

	return false
}
