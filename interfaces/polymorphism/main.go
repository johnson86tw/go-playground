package main

import "fmt"

// source: https://go101.org/article/interface.html

// Filter ...
type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter ...
type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "remove duplicate numbers"
}

func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)

	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}

	return outs
}

type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))

	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}

	return outs
}

// polymorphism makes it unnecessary to write one filterAndPrint function for each filter types.
func filterAndPrint(fltr Filter, unfiltered []int) []int {
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filtering:\n\t", numbers)

	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	for _, fltr := range filters {
		numbers = filterAndPrint(fltr, numbers)
	}

	fmt.Println("Done! result: ", numbers)
}
