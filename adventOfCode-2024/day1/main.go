package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
  t0 := time.Now()
	fmt.Println("Part 1 answer:", part1(), "took", time.Since(t0))
  t0 = time.Now()
	fmt.Println("Part 2 answer:", part2(), "took", time.Since(t0))

}

func part2() int {
  filename := "data.txt"

  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scan := bufio.NewScanner(file)

  rmap := make(map[int]int)
  l := []int{}

  for scan.Scan(){
    line := scan.Text()
    split := strings.Split(line, "   ")

    lv, _ := strconv.Atoi(split[0])
    rv, _ := strconv.Atoi(split[1])

    l = append(l, lv)
    rmap[rv]++
  }

  total := 0
  for _, l := range l {
    total += l * rmap[l]
  }

  return total
}

func part1() int {
	filename := "data.txt"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
  defer file.Close()

	scn := bufio.NewScanner(file)

	l := []int{}
	r := []int{}

	for scn.Scan() {
		line := scn.Text()
		split := strings.Split(line, "   ")
		lv, _ := strconv.Atoi(split[0])
		rv, _ := strconv.Atoi(split[1])
		l = append(l, lv)
		r = append(r, rv)
	}

	sort.Ints(l)
	sort.Ints(r)

	total := 0
	for i := range len(l) {
		total += intAbs(l[i], r[i])
	}

	return total

}

func intAbs(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	} else {
		return diff
	}
}
