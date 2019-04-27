package graphs

// Cordinate .
type Cordinate struct {
	X, Y int
}

// SearchMaze .
func SearchMaze(maze [][]bool, start, end Cordinate) []Cordinate {
	path := []Cordinate{}

	maze[start.X][start.Y] = false
	path = append(path, start)

	ok, path := searchMazeHelper(maze, start, end, path)
	if ok {
		path = path[:len(path)-1]
	}

	return path
}

var shifts = [][]int{
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
	[]int{-1, 0},
}

func searchMazeHelper(maze [][]bool, s, e Cordinate, path []Cordinate) (bool, []Cordinate) {
	if s.X == e.X && s.Y == e.Y {
		return true, path
	}

	for _, shift := range shifts {
		next := Cordinate{X: s.X + shift[0], Y: s.Y + shift[1]}
		if isFeasable(next, maze) {
			maze[next.X][next.Y] = false
			path = append(path, next)
			if ok, path := searchMazeHelper(maze, next, e, path); ok {
				return true, path
			}
			path = path[:len(path)-1]
		}
	}

	return false, path
}

func isFeasable(n Cordinate, maze [][]bool) bool {
	return n.X >= 0 && n.X < len(maze) &&
		n.Y >= 0 && n.Y < len(maze[n.X]) &&
		maze[n.X][n.Y] == true
}
