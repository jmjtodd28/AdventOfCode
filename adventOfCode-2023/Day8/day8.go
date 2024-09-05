package main

import (
	"fmt"
	"os"
	"strings"
)

func step_number(instructions string, path_map map[string][]string, current_pos string) int {

	i := 0
	steps := 0
	for {
		steps += 1
		if string(instructions[i]) == "L" {
			current_pos = path_map[current_pos][0]

		} else {
			current_pos = path_map[current_pos][1]
		}

		if current_pos == "ZZZ" {
			break
		}
		i = (i + 1) % len(instructions)
	}

	return steps
}

// get the cycles of the starting posistions where they find a positions that ends in z and then find the LCM of all the starting positions
func part_two(instructions string, path_map map[string][]string, current_pos []string) int {

	endpoints := []int{}

	for _, pos := range current_pos {
		steps := 0

		for {
			if strings.HasSuffix(pos, "Z") {
				endpoints = append(endpoints, steps)
				break
			}

			if string(instructions[steps%len(instructions)]) == "L" {
				pos = path_map[pos][0]
			} else {
				pos = path_map[pos][1]
			}
			steps++
		}
	}

	val := endpoints[0]
	for i := 1; i < len(endpoints); i++ {
		val = lcm(val, endpoints[i])
	}

	return val

}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x * (y / gcd(x, y))
}

func get_start_positions(path_map map[string][]string) []string {
	var starting_positions []string
	for key := range path_map {
		if string(key[2]) == "A" {
			starting_positions = append(starting_positions, key)
		}
	}

	return starting_positions

}

func main() {
	filename := "input.txt"
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	// extract all the instructions and create a dictionary with destinations and potential routes
	d := string(data)

	insructions := strings.Split(d, "\n")[0]

	paths := strings.Split(d, "\n")[2:]
	paths = paths[:len(paths)-1]

	path_map := make(map[string][]string)

	for _, p := range paths {
		key := strings.Split(p, "=")[0]
		key = strings.Trim(key, " ")

		left := strings.Split(p, "=")[1][2:5]
		right := strings.Split(p, "=")[1][7:10]

		path_map[key] = append(path_map[key], left)
		path_map[key] = append(path_map[key], right)
	}

	//part 1
	part_1 := step_number(insructions, path_map, "AAA")
	fmt.Println("Part 1: ", part_1)

	//part 2
	//get all the starting positions, ie, all places that end with an a
	start_positions := get_start_positions(path_map)

	total_steps := part_two(insructions, path_map, start_positions)

	fmt.Println("Part 2: ", total_steps)

}
