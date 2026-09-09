func isValidSudoku(board [][]byte) bool {
	for i := 0; i <= 8; i++ {
		// build 3 sets: row, column, group 
		// row is easy 
		rowSet := board[i]

		// column has to be built from each row
		columnSet := make([]byte, 0)
		for j := 0; j < 9; j++ {
			columnSet = append(columnSet, board[j][i])
		}

		// group requires math 
		groupSet := make([]byte, 0, 9)
		startRow := 3 * (i / 3) // 0, 0, 0, 3, 3, 3, 6, 6, 6
		startCol := 3 * (i % 3) // 0, 3, 6, 0, 3, 6, 0, 3, 6

		// add the next two rows/columns' worth
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				groupSet = append(groupSet, board[startRow+r][startCol+c])
			}
		}

		// validate that all three sets have no duplicate numbers
		if !noDuplicates(rowSet) || !noDuplicates(columnSet) || !noDuplicates(groupSet) {
			return false
		}
	}

	// no errors found so return true
	return true; 
}

func noDuplicates(check []byte) bool {
	found := make(map[byte]bool)

	for _, val := range check {
		if val != '.' { // if it's a . we don't care, it's blank
			_, hit := found[val]
			if hit {
				return false
			} else {
				found[val] = true
			}
		}
	}

	return true
}