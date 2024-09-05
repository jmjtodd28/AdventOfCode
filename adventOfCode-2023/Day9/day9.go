package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntSlice []int

func (arr IntSlice) is_all_zeros() bool {
	all_zeros := true
	for _, n := range arr {
		if n != 0 {
			all_zeros = false
		}
	}
	return all_zeros
}

func (nums IntSlice) sequence_total() int {
	total := 0
	nums_len := len(nums)
	for i := range nums {
		total = nums[nums_len-1-i] - total
	}
	return total
}

func solve(input []string) (int, int) {

	part1_total := 0
	part2_total := 0

	for _, seq := range input {
		//convert the array of strings into an array of integers for ease
		nums_strings := strings.Split(seq, " ")
		nums := make([]int, len(nums_strings))
		for i, s := range nums_strings {
			n, _ := strconv.Atoi(s)
			nums[i] = n
		}
		//this will be used to keep track of all the first numbers in each sequence so we calculate part 2
		var first_nums IntSlice
		next_seq_len := len(nums)
		seq_depth := 0
		part1_total += nums[len(nums)-1]
		first_nums = append(first_nums, nums[0])

		for {
			//calculate the next level of the sequence and add the last int to the part 1 total
			var next_level IntSlice
			for i := range next_seq_len {
				if i == next_seq_len-1 {
					break
				}
				diff := nums[seq_depth+i+1] - nums[seq_depth+i]
				next_level = append(next_level, diff)
			}
			part1_total += next_level[len(next_level)-1]
			//Add first number
			first_nums = append(first_nums, next_level[0])
			//Add next level to the number array
			nums = append(nums, next_level...)
			//check to see if we have reached the end of the sequence and all values in the next level are zero
			if next_level.is_all_zeros() {
				break
			}
			next_seq_len = len(next_level)
			seq_depth = len(nums) - len(next_level)
		}
		//Once we have collected all the first numbers from all the level, we can calculate the total for that sequence and add it to the total
		first_nums_total := first_nums.sequence_total()
		part2_total = part2_total + first_nums_total
	}
	return part1_total, part2_total
}

func main() {
	filename := "input.txt"
	file, err := os.ReadFile(filename)

	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1]

	if err != nil {
		fmt.Println(err)
	}

	part_one_total, part_two_total := solve(input)
	fmt.Println("Part 1 total: ", part_one_total)
	fmt.Println("Part 2 total: ", part_two_total)
}
