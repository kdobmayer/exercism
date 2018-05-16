package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
				continue
			}
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}