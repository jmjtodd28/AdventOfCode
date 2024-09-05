package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hashString(text string) int {
	total := 0

	for _, char := range text {
		if char != 44 && char != 10 {
			total = ((total + int(char)) * 17) % 256
		}
	}
	return total
}

func part1(input []byte) int {

	inputArray := strings.Split(string(input), ",")

	grandtotal := 0

	for _, elem := range inputArray {
		elemTotal := hashString(elem)
		grandtotal = grandtotal + elemTotal
	}

	return grandtotal
}

func part2(data []byte) int {

	boxes := make([][]string, 256)

	dataArray := strings.Split(string(data), ",")

	//iterate over the input data and add/remove from boxes
	for _, d := range dataArray {

		if d[len(d)-1] == '-' {

			label := strings.Split(d, "-")[0]
			boxNum := hashString(label)

			for i, b := range boxes[boxNum] {
				b_label := strings.Split(b, "=")[0]
				if label == b_label {
					boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
				}
			}
		} else {

			label := strings.Split(d, "=")[0]
			boxNum := hashString(label)

			contains := false
			//check if lens is in the box
			for i, b := range boxes[boxNum] {
				b_label := strings.Split(b, "=")[0]
				if label == b_label {
					boxes[boxNum][i] = d
					contains = true
				}
			}

			if contains == false {
				boxes[boxNum] = append(boxes[boxNum], d)
			}
		}
	}

	//calculate the final value based on the boxes
	total := 0

	for i, ob := range boxes {
		for j, ib := range ob {
			focalLength := strings.SplitAfter(ib, "=")[1]
			lensPower, _ := strconv.Atoi(focalLength)
			total += (i + 1) * (j + 1) * lensPower
		}
	}
	return total
}

func main() {
	filename := "input.txt"
	input, err := os.ReadFile(filename)
	//remove the trailing new line character
	input = input[:len(input)-1]

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1 total is: ", part1(input))
	fmt.Println("Part 2 total is: ", part2(input))
}
