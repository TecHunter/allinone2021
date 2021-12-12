package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
)

const Debug = false

func main() {
	var prefix string
	if len(os.Args) > 1 {
		prefix = os.Args[1]
	} else {
		prefix = "./2021/7/"
	}
	FirstCase(prefix + ".input")
	SecondCase(prefix + ".input")
}

func ProcessInput(file string) [][]byte {
	input, _ := ioutil.ReadFile(file)
	return bytes.Split(input[:bytes.Index(input, []byte("\n"))], []byte(","))
}

func FirstCase(file string) int {
	if Debug {
		wd, _ := os.Getwd()
		println("Workdir is ", wd)
		println("DAY-7 -- 1st case -- input: ", file)
	}
	input := ProcessInput(file)
	var crabs []int
	for _, s := range input {
		pos, _ := strconv.Atoi(string(s))
		crabs = append(crabs, pos)
	}
	sort.Ints(crabs)
	center := crabs[int(len(crabs)/2)]

	fuelCost := 0
	for _, position := range crabs {
		fuelCost += int(math.Abs(float64(position - center)))
	}
	println("\tResult: ", fuelCost)

	return fuelCost
}
func SecondCase(file string) int {
	if Debug {
		wd, _ := os.Getwd()
		println("Workdir is ", wd)
		println("DAY-7 -- 2nd case -- input: ", file)
	}

	input := ProcessInput(file)
	crabs, flooredMean := Mean(input)
	fuelCost := 0
	for _, pos := range crabs {
		//realignment cost
		fuelCost += Sum(int(math.Abs(float64(pos - flooredMean))))
	}
	println("\tResult: ", fuelCost)

	return fuelCost
}

// Sum
// sum from n-1 to k = nk − (k/2) * (k+1)
// https://math.stackexchange.com/questions/556807/what-is-the-sum-of-n-1n-2-n-k
//
func Sum(n int) int {
	return (n + 1) * n / 2
}

func Mean(input [][]byte) ([]int, int) {
	total := 0.0
	var crabs = make([]int, len(input))
	for i, s := range input {
		pos, _ := strconv.Atoi(string(s))
		//println(pos)
		crabs[i] = pos
		total += float64(pos)
	}
	//find the reference point

	flooredMean := total / float64(len(input))
	fmt.Printf("%.2f / %d = %d\n", total, len(input), int(flooredMean))
	//tried to use math.Round here but ended up with a weird sum
	return crabs, int(flooredMean)
}
