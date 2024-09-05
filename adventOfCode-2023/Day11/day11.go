package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Coords struct {
	x int
	y int
}

func getEmtpyRows(universe []string) []int {
	emptyRows := []int{}
	for i, line := range universe {
		emptySpace := true
		for _, c := range line {
			if c != '.' {
				emptySpace = false
				break
			}
		}
		if emptySpace {
			emptyRows = append(emptyRows, i)
		}
	}
	return emptyRows
}

func getEmptyColumns(universe []string) []int {
	emptyColumns := []int{}
	for i := range len(universe[0]) {
		emptySpace := true
		for j := range len(universe) {
			if universe[j][i] != '.' {
				emptySpace = false
				break
			}
		}
		if emptySpace == true {
			emptyColumns = append(emptyColumns, i)
		}
	}
	return emptyColumns
}

func findGalaxies(universe []string) []Coords {
	galaxyCoords := []Coords{}
	for i := range universe[0] {
		for j := range universe {
			if universe[j][i] == '#' {
				galaxyCoords = append(galaxyCoords, Coords{x: i, y: j})
			}
		}
	}
	return galaxyCoords
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sumOfDistances(galaxies []Coords) int {
	sum := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}
	return sum
}

// For every empty row that is above a galaxy, you add 1,000,000 to the y value
// For every empty column to the left of a galazy, you add 1,000,000 to the x value
func adjustGalaxies(galaxies []Coords, universe []string, factor int) {

	emptyRows := getEmtpyRows(universe)
	emptyCols := getEmptyColumns(universe)

	for i := range galaxies {
		x := galaxies[i].x
		for _, c := range emptyCols {
			if x > c {
				galaxies[i].x = galaxies[i].x + factor
			} else {
				break
			}
		}
	}

	for i := range galaxies {
		y := galaxies[i].y
		for _, r := range emptyRows {
			if y > r {
				galaxies[i].y = galaxies[i].y + factor
			} else {
				break
			}
		}
	}
}

func partOne(universe []string) int {

	galaxies := findGalaxies(universe)

	adjustGalaxies(galaxies, universe, 1)

	sum := sumOfDistances(galaxies)

	return sum
}

func partTwo(universe []string) int {

	galaxies := findGalaxies(universe)

	adjustGalaxies(galaxies, universe, 999999)

	sum := sumOfDistances(galaxies)

	return sum
}

func getInput() []string {

	filename := "input.txt"
	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1]

	return input
}

func main() {

	t0 := time.Now()

	sum := partOne(getInput())
	fmt.Println("Part 1:", sum, "took:", time.Since(t0))

	sum2 := partTwo(getInput())
	fmt.Println("Part 2:", sum2, "took:", time.Since(t0))
}
