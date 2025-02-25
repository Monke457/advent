package grid

var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1}, 
	{1, -1}, {1, 0}, {1, 1},
}
var dirsCont = [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func oob[T comparable](data [][]T, y, x int) bool {
	if y < 0 || y >= len(data) {
		return true 
	}
	if x < 0 || x >= len(data[y]) {
		return true 
	}
	return false
}
