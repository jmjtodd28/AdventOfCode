package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ranges struct {
	dest   int
	source int
	length int
}

func partOne(seeds []string, mapping [][]Ranges) int {

	minSeed := int(math.Inf(1))

	for _, seed := range seeds {
		seed, _ := strconv.Atoi(seed)
		for _, m := range mapping {
			for _, v := range m {
				if seed >= v.source && seed < v.source+v.length {
					seed = v.dest + (seed - v.source)
					break
				}
			}
		}
		if seed < minSeed {
			minSeed = seed
		}
	}

	return minSeed
}

func partTwo(seeds []string, mapping [][]Ranges) int {
	minSeed := int(math.Inf(1))

	for i := 0; i < len(seeds); i = i + 2 {
		fmt.Printf("i: %v\n", i)
		start, _ := strconv.Atoi(seeds[i])
		sRange, _ := strconv.Atoi(seeds[i+1])
		for j := range sRange {
			seed := j + start
			for _, m := range mapping {
				for _, v := range m {
					if seed >= v.source && seed < v.source+v.length {
						seed = v.dest + (seed - v.source)
						break
					}
				}
			}
			if seed < minSeed {
				minSeed = seed
			}
		}

	}

	return minSeed
}

func parse(input []byte) ([]string, [][]Ranges) {

	data := strings.Split(string(input), "\n\n")

	seeds := strings.Split(data[0], " ")[1:]

	data = data[1:]

	var mapping [][]Ranges

	for i, d := range data {
		mapping = append(mapping, []Ranges{})
		x := strings.Split(d, "\n")
		x = x[1:]
		for _, y := range x {
			z := strings.Split(y, " ")
			dest, _ := strconv.Atoi(z[0])
			source, _ := strconv.Atoi(z[1])
			length, _ := strconv.Atoi(z[2])
			mapping[i] = append(mapping[i], Ranges{dest: dest, source: source, length: length})
		}

	}

	return seeds, mapping
}

func main() {
	filename := "input.txt"
	input, err := os.ReadFile(filename)

	t0 := time.Now()

	if err != nil {
		fmt.Println(err)
	}

	seeds, mapping := parse(input)

	minSeed := partOne(seeds, mapping)
	fmt.Println("Part 1:", minSeed, "took", time.Since(t0))

	minSeed = partTwo(seeds, mapping)
	fmt.Println("Part 2:", minSeed, "took", time.Since(t0))
}
