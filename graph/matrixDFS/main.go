package main

import "fmt"

func uniquePaths(grid [][]int) int {
	return dfs(grid, 0, 0)
}

func dfs(grid [][]int, row, col int) int {
	r, c := len(grid), len(grid[0])

	if row < 0 || row >= r || col < 0 || col >= c || grid[row][col] == 1 {
		return 0
	}
	if row == r-1 && col == c-1 {
		return 1
	}
	grid[row][col] = 1 // Marking as visited

	paths := 0
	paths += dfs(grid, row-1, col)
	paths += dfs(grid, row+1, col)
	paths += dfs(grid, row, col-1)
	paths += dfs(grid, row, col+1)

	grid[row][col] = 0
	return paths

}

func main() {

	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}
	result := uniquePaths(grid)
	fmt.Println("Number of unique paths:", result)
	queue := [][]int{{1, 2}}
	fmt.Println(queue)

}
