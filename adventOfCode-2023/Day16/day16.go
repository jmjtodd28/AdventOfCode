package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

type coord struct {
	x int
	y int
}

// Recursive function that returns a array of all the visited coordinates
// Appends newly visited squares to the visited coordinates
// there are checks to make sure there are no duplicates
// we only visit | and - once if there is splitting to avoid infinite loops because the outcome will always be the same
func energiseGrid(pos coord, direction string, grid []string, visited []coord) []coord {

	if pos.x < 0 || pos.x >= len(grid)-1 || pos.y < 0 || pos.y >= len(grid[0]) {
		return visited
	}

	if string(grid[pos.y][pos.x]) == "." {
		if !slices.Contains(visited, coord{pos.x, pos.y}) {
			visited = append(visited, coord{pos.x, pos.y})
		}
		switch direction {
		case "r":
			visited = energiseGrid(coord{pos.x + 1, pos.y}, "r", grid, visited)
		case "l":
			visited = energiseGrid(coord{pos.x - 1, pos.y}, "l", grid, visited)
		case "u":
			visited = energiseGrid(coord{pos.x, pos.y - 1}, "u", grid, visited)
		case "d":
			visited = energiseGrid(coord{pos.x, pos.y + 1}, "d", grid, visited)
		}
	}

	if string(grid[pos.y][pos.x]) == "|" {
		switch direction {

		case "r":
			if slices.Contains(visited, coord{pos.x, pos.y}) {
				return visited
			} else {
				visited = append(visited, coord{pos.x, pos.y})
				visited = energiseGrid(coord{pos.x, pos.y - 1}, "u", grid, visited)
				visited = energiseGrid(coord{pos.x, pos.y + 1}, "d", grid, visited)
			}
		case "l":
			if slices.Contains(visited, coord{pos.x, pos.y}) {
				return visited
			} else {
				visited = append(visited, coord{pos.x, pos.y})
				visited = energiseGrid(coord{pos.x, pos.y - 1}, "u", grid, visited)
				visited = energiseGrid(coord{pos.x, pos.y + 1}, "d", grid, visited)
			}
		case "u":
			if !slices.Contains(visited, coord{pos.x, pos.y}) {
				visited = append(visited, coord{pos.x, pos.y})
			}
			visited = energiseGrid(coord{pos.x, pos.y - 1}, "u", grid, visited)
		case "d":
			if !slices.Contains(visited, coord{pos.x, pos.y}) {
				visited = append(visited, coord{pos.x, pos.y})
			}
			visited = energiseGrid(coord{pos.x, pos.y + 1}, "d", grid, visited)
		}
	}

	if string(grid[pos.y][pos.x]) == "-" {
		switch direction {
		case "r":
			if !slices.Contains(visited, coord{pos.x, pos.y}) {
				visited = append(visited, coord{pos.x, pos.y})
			}
			visited = energiseGrid(coord{pos.x + 1, pos.y}, "r", grid, visited)
		case "l":
			if !slices.Contains(visited, coord{pos.x, pos.y}) {
				visited = append(visited, coord{pos.x, pos.y})
			}
			visited = energiseGrid(coord{pos.x - 1, pos.y}, "l", grid, visited)
		case "u":
			if slices.Contains(visited, coord{pos.x, pos.y}) {
				return visited
			} else {
				visited = append(visited, coord{pos.x, pos.y})
				visited = energiseGrid(coord{pos.x - 1, pos.y}, "l", grid, visited)
				visited = energiseGrid(coord{pos.x + 1, pos.y}, "r", grid, visited)
			}
		case "d":
			if slices.Contains(visited, coord{pos.x, pos.y}) {
				return visited
			} else {
				visited = append(visited, coord{pos.x, pos.y})
				visited = energiseGrid(coord{pos.x - 1, pos.y}, "l", grid, visited)
				visited = energiseGrid(coord{pos.x + 1, pos.y}, "r", grid, visited)
			}
		}
	}

	if string(grid[pos.y][pos.x]) == "/" {

		if !slices.Contains(visited, coord{pos.x, pos.y}) {
			visited = append(visited, coord{pos.x, pos.y})
		}
		switch direction {
		case "l":
			visited = energiseGrid(coord{pos.x, pos.y + 1}, "d", grid, visited)
		case "r":
			visited = energiseGrid(coord{pos.x, pos.y - 1}, "u", grid, visited)
		case "u":
			visited = energiseGrid(coord{pos.x + 1, pos.y}, "r", grid, visited)
		case "d":
			visited = energiseGrid(coord{pos.x - 1, pos.y}, "l", grid, visited)
		}
	}

	if string(grid[pos.y][pos.x]) == "\\" {
		if !slices.Contains(visited, coord{pos.x, pos.y}) {
			visited = append(visited, coord{pos.x, pos.y})
		}
		switch direction {
		case "l":
			visited = energiseGrid(coord{pos.x, pos.y - 1}, "u", grid, visited)
		case "r":
			visited = energiseGrid(coord{pos.x, pos.y + 1}, "d", grid, visited)
		case "u":
			visited = energiseGrid(coord{pos.x - 1, pos.y}, "l", grid, visited)
		case "d":
			visited = energiseGrid(coord{pos.x + 1, pos.y}, "r", grid, visited)
		}
	}

	return visited
}

func main() {

	t0 := time.Now()

	filename := "input.txt"
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	grid := strings.Split(string(data), "\n")

	//part1
	start := coord{x: 0, y: 0}
	var visited []coord

	visited = energiseGrid(start, "r", grid, visited)

	fmt.Println("Part 1: ", len(visited), " took ", time.Since(t0))

	//part 2
	//making use of go routines
	var wg = sync.WaitGroup{}
  //locks around the shared maxEnergise variable
  var mu sync.Mutex

	maxEnergise := 0

	//top row going down
	for i := 0; i < len(grid[0]); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			visited = []coord{}
			visited = energiseGrid(coord{i, 0}, "d", grid, visited)

      mu.Lock()
			maxEnergise = max(maxEnergise, len(visited))
      mu.Unlock()
		}(i)
	}
	//bottom row going up
	for i := 0; i < len(grid[0]); i++ {
		wg.Add(1)
		go func(i int) {
      defer wg.Done()

			visited = []coord{}
			visited = energiseGrid(coord{i, len(grid) - 2}, "u", grid, visited)

      mu.Lock()
			maxEnergise = max(maxEnergise, len(visited))
      mu.Unlock()
		}(i)
	}
	//left column going right
	for i := 0; i < len(grid)-1; i++ {
		wg.Add(1)
		go func(i int) {
      defer wg.Done()

			visited = []coord{}
			visited = energiseGrid(coord{0, i}, "r", grid, visited)

      mu.Lock()
			maxEnergise = max(maxEnergise, len(visited))
      mu.Unlock()
		}(i)
	}
	//right column going left
	for i := 0; i < len(grid)-1; i++ {
		wg.Add(1)
		go func(i int) {
      defer wg.Done()

			visited = []coord{}
			visited = energiseGrid(coord{len(grid[0]) - 1, i}, "l", grid, visited)

      mu.Lock()
			maxEnergise = max(maxEnergise, len(visited))
      mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println("Part 2: ", maxEnergise, " took ", time.Since(t0))

}
