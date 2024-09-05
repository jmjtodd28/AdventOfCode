package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type coord struct {
	x int
	y int
}

func get_start_point(input []string) coord {
	for i, l := range input {
		for j, c := range l {
			if c == 'S' {
				start := coord{j, i}
				return start
			}
		}
	}
	return coord{0, 0}
}

func get_loop_coords(input []string, start coord) []coord {

	var loop_coords []coord
	current_pos := start
	previous_pos := coord{-1, -1}

	//search around the start position and choose first viable path
	if string(input[current_pos.y][current_pos.x+1]) == "J" || string(input[current_pos.y][current_pos.x+1]) == "7" || string(input[current_pos.y][current_pos.x+1]) == "-" {
		loop_coords = append(loop_coords, coord{current_pos.x + 1, current_pos.y})
		previous_pos = current_pos
		current_pos = coord{current_pos.x + 1, current_pos.y}
	} else if string(input[current_pos.y][current_pos.x-1]) == "L" || string(input[current_pos.y][current_pos.x+1]) == "F" || string(input[current_pos.y][current_pos.x+1]) == "-" {
		loop_coords = append(loop_coords, coord{current_pos.x - 1, current_pos.y})
		previous_pos = current_pos
		current_pos = coord{current_pos.x - 1, current_pos.y}
	} else if string(input[current_pos.y+1][current_pos.x]) == "J" || string(input[current_pos.y][current_pos.x+1]) == "|" || string(input[current_pos.y][current_pos.x+1]) == "L" {
		loop_coords = append(loop_coords, coord{current_pos.x, current_pos.y + 1})
		previous_pos = current_pos
		current_pos = coord{current_pos.x, current_pos.y + 1}
	} else if string(input[current_pos.y-1][current_pos.x]) == "F" || string(input[current_pos.y][current_pos.x+1]) == "|" || string(input[current_pos.y][current_pos.x+1]) == "F" {
		loop_coords = append(loop_coords, coord{current_pos.x, current_pos.y - 1})
		previous_pos = current_pos
		current_pos = coord{current_pos.x, current_pos.y - 1}
	}

Loop:
	for {
		switch input[current_pos.y][current_pos.x] {
		case 'S':
			break Loop
		case 'J':
			u := coord{current_pos.x, current_pos.y - 1}
			l := coord{current_pos.x - 1, current_pos.y}
			if l == previous_pos {
				loop_coords = append(loop_coords, u)
				previous_pos = current_pos
				current_pos = u
			} else {
				loop_coords = append(loop_coords, l)
				previous_pos = current_pos
				current_pos = l
			}
		case 'L':
			u := coord{current_pos.x, current_pos.y - 1}
			r := coord{current_pos.x + 1, current_pos.y}
			if r == previous_pos {
				loop_coords = append(loop_coords, u)
				previous_pos = current_pos
				current_pos = u
			} else {
				loop_coords = append(loop_coords, r)
				previous_pos = current_pos
				current_pos = r
			}
		case 'F':
			r := coord{current_pos.x + 1, current_pos.y}
			d := coord{current_pos.x, current_pos.y + 1}
			if r == previous_pos {
				loop_coords = append(loop_coords, d)
				previous_pos = current_pos
				current_pos = d
			} else {
				loop_coords = append(loop_coords, r)
				previous_pos = current_pos
				current_pos = r
			}
		case '7':
			d := coord{current_pos.x, current_pos.y + 1}
			l := coord{current_pos.x - 1, current_pos.y}
			if l == previous_pos {
				loop_coords = append(loop_coords, d)
				previous_pos = current_pos
				current_pos = d
			} else {
				loop_coords = append(loop_coords, l)
				previous_pos = current_pos
				current_pos = l
			}
		case '|':
			u := coord{current_pos.x, current_pos.y - 1}
			d := coord{current_pos.x, current_pos.y + 1}
			if u == previous_pos {
				loop_coords = append(loop_coords, d)
				previous_pos = current_pos
				current_pos = d
			} else {
				loop_coords = append(loop_coords, u)
				previous_pos = current_pos
				current_pos = u
			}
		case '-':
			l := coord{current_pos.x - 1, current_pos.y}
			r := coord{current_pos.x + 1, current_pos.y}
			if l == previous_pos {
				loop_coords = append(loop_coords, r)
				previous_pos = current_pos
				current_pos = r
			} else {
				loop_coords = append(loop_coords, l)
				previous_pos = current_pos
				current_pos = l
			}
		}
	}
	return loop_coords
}

func shoelace_area(boundary []coord) float64 {

	total_area := 0.0
	for i := range boundary {
		ax := boundary[i].x
		ay := boundary[i].y
		bx := boundary[(i+1)%len(boundary)].x
		by := boundary[(i+1)%len(boundary)].y
		A := ax*by - ay*bx
		total_area += float64(A) / 2
	}
	return math.Abs(float64(total_area))
}

func picks_theorem(area float64, boundary_points int) float64 {
	num_of_points := area - (float64(boundary_points) / 2) + 1

	return num_of_points
}

func main() {
	t0 := time.Now()

	filename := "input.txt"
	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	input := strings.Split(string(file), "\n")

	//part 1
	start := get_start_point(input)
	loop_coords := get_loop_coords(input, start)
	max_dist_idx := len(loop_coords) / 2
	fmt.Println("Part 1:", max_dist_idx, "Time:", time.Since(t0))

	//part2 using shoelace formula and picks theorem

	//first find the area using the shoelace formula
	area := shoelace_area(loop_coords)

	//use that area with Picks theorum to find the number of points
	num_of_points := picks_theorem(area, len(loop_coords))
	fmt.Println("Part 2:", num_of_points, "Time", time.Since(t0))
}
