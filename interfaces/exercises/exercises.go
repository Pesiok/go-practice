package exc

import (
	"fmt"
	"sort"
)

type people []string

// to meet requirements of the 'Interface' interface we need to implement following methods:
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func SortExc() {

	// 1 /\
	studyGroupOne := people{"Zeno", "Clash", "Al", "Barbara", "Gopher"}
	sort.Sort(studyGroupOne)

	// 2
	studyGroupTwo := []string{"Zeno", "Clash", "Al", "Barbara", "Gopher"}
	// sort.Strings(studyGroupTwo)
	// sort.Sort(sort.StringSlice(studyGroupTwo))
	// StringSlice type have already implemented sort methods required by 'Interface' interface
	sort.StringSlice(studyGroupTwo).Sort()

	// 3
	numbers := []int{1, 23, 5, 3, 32, 6, 7, 2, 5, 7}
	// sort.IntSlice(numbers).Sort()
	// sort.Ints(numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	fmt.Println(studyGroupOne, studyGroupTwo, numbers)
}
