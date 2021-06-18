package alggo

func highestHourglass(matrix [6][6]int) int {
	var highest int
	var indexes = []int{1, 2, 3, 4}
	for _, i := range indexes {
		for _, j := range indexes {
			sum := matrix[i][j] + matrix[i-1][j-1] + matrix[i-1][j] + matrix[i-1][j+1] + matrix[i+1][j-1] + matrix[i+1][j] + matrix[i+1][j+1]
			if highest < sum {
				highest = sum
			}
		}
	}

	return highest
}
