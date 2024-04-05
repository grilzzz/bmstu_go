package main

import (
	"fmt"
	"math"
)

func main() {
	var N, x int
	fmt.Scan(&N)

	grid := [][]int{}
	for i := 0; i < N; i++ {
		grid = append(grid, []int{})
		for j := 0; j < N; j++ {
			fmt.Scan(&x)
			grid[i] = append(grid[i], x)
		}
	}

	queue := [][]int{{0, 0}}

	dist := make([][]int, N)

	for i := range dist {
		dist[i] = make([]int, N)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = grid[0][0]

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		x, y := point[0], point[1]

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]

			if newX >= 0 && newX < N && newY >= 0 && newY < N {
				newDist := dist[x][y] + grid[newX][newY]
				if newDist < dist[newX][newY] {
					dist[newX][newY] = newDist
					queue = append(queue, []int{newX, newY})
				}
			}
		}
	}

	fmt.Println(dist[N-1][N-1])
}
