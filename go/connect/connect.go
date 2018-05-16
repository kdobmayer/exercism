package connect

const X, O = 'X', 'O'

type Field struct{ x, y int }

func isConnect(board []string, visited []Field, f Field, stone byte) bool {
	if stone == X && f.y == len(board[0])-1 ||
		stone == O && f.x == len(board)-1 {
		return true
	}
	visited = append(visited, f)

	for i := 1; i < 8; i++ {
		next := Field{f.x + (i/3 - 1), f.y + (i%3 - 1)}
		if next.x < 0 || next.y < 0 || next.x > len(board)-1 || next.y > len(board[0])-1 ||
			board[next.x][next.y] != stone || isVisited(visited, next) {
			continue
		}

		return isConnect(board, visited, next, stone)
	}
	return false
}

func isVisited(visited []Field, f Field) bool {
	for _, v := range visited {
		if v.x == f.x && v.y == f.y {
			return true
		}
	}
	return false
}

func ResultOf(board []string) (string, error) {
	for i := range board {
		if board[i][0] == X && isConnect(board, []Field{}, Field{i, 0}, X) {
			return string(X), nil
		}
	}
	for i := range board[0] {
		if board[0][i] == O && isConnect(board, []Field{}, Field{0, i}, O) {
			return string(O), nil
		}
	}
	return "", nil
}