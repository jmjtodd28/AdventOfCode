package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Coord struct {
	x int
	y int
}

func getStartPos(grid []string) Coord {

	start := Coord{}
	for i := range grid {
		for j := range len(grid[0]) {
			if grid[j][i] == 'S' {
				start = Coord{x: i, y: j}
				break
			}
		}
	}

	return start
}

func BFS(grid []string, start Coord, steps int) int {

	//Queue represents the nodes we need to explore
	queue := make([]Coord, 0)
	queue = append(queue, start)

	//Queue to represent the nodes to explore after this iteration
	nextQueue := make([]Coord, 0)

	//A queue to represent the visited nodes so we dont have to revisit and waste time
	visited := make([]Coord, 0)

	for i := 0; i < steps; i++ {
		for _, pos := range queue {
			if pos.x+1 < len(grid[0]) && grid[pos.y][pos.x+1] != '#' && !slices.Contains(visited, Coord{x: pos.x + 1, y: pos.y}) {
				nextQueue = append(nextQueue, Coord{x: pos.x + 1, y: pos.y})
				visited = append(visited, Coord{x: pos.x + 1, y: pos.y})
			}
			if pos.x-1 > 0 && grid[pos.y][pos.x-1] != '#' && !slices.Contains(visited, Coord{x: pos.x - 1, y: pos.y}) {
				nextQueue = append(nextQueue, Coord{x: pos.x - 1, y: pos.y})
				visited = append(visited, Coord{x: pos.x - 1, y: pos.y})
			}
			if pos.y+1 < len(grid) && grid[pos.y+1][pos.x] != '#' && !slices.Contains(visited, Coord{x: pos.x, y: pos.y + 1}) {
				nextQueue = append(nextQueue, Coord{x: pos.x, y: pos.y + 1})
				visited = append(visited, Coord{x: pos.x, y: pos.y + 1})
			}
			if pos.y-1 > 0 && grid[pos.y-1][pos.x] != '#' && !slices.Contains(visited, Coord{x: pos.x, y: pos.y - 1}) {
				nextQueue = append(nextQueue, Coord{x: pos.x, y: pos.y - 1})
				visited = append(visited, Coord{x: pos.x, y: pos.y - 1})
			}
		}
		queue = nextQueue
		visited = []Coord{}
		nextQueue = []Coord{}
	}

	return len(queue)
}

func main() {

	t0 := time.Now()

	filename := "input.txt"
	input, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	grid := strings.Split(string(input), "\n")
	//	grid = grid[:len(grid)-1]

	//	for i := range grid {
	//		fmt.Printf("%v %d\n", grid[i], len(grid[i]))
	//	}

	start := getStartPos(grid)

	fmt.Printf("start: %v\n", start)

	partOne := BFS(grid, start, 1000)
	fmt.Printf("Part One: %v Took: %s\n", partOne, time.Since(t0))

}
