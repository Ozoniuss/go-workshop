package main

import (
	"fmt"
	"slices"
	"strings"
)

// flow control, if and for. Also introduce range over slice or map.

// range will go over arrays inthe same order

/*
Build an array with the elements 100, 45, 24, 96 and 789. Go over each element,
add them to a variable and find their sum.
*/

func exercise4a() {
	// short initialization was explained by this point
	arr := []int{100, 45, 24, 96, 789}
	s := 0

	// len was also explained
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	fmt.Println(s)

	// at this point, you should introduce the range keyword. Do the exact same
	// exercise using range
}

/*
Build a string with the names Alex, Armand, Andrei, Alexandra and Mihai. Find
the number of names containing the "dr" combo.
*/

// This encourages them to use the strings library, which was already introduced.
// Give them some time, see if they manage to look up the right function in the
// strings library. When showing the solution, be sure to let them know again
// to go through the library's functions. Repetition is key!!
func exercise4b() {
	arr := []string{"Alex", "Armand", "Andrei", "Alexandra", "Mihai"}
	// same pattern
	s := 0

	// essential to repeat again the use of range here
	for _, name := range arr {
		// This is the function they should ideally find.
		if strings.Contains(name, "dr") {
			s += 1
		}
	}
	fmt.Println(s)
}

/*
Before we go to map, it's important to be aware of the slices package,
introduced in go 1.21. This contains useful slices functions. Tell them
about this, so they add it to their internal list of packages they should
be aware of.
*/

// Sort the items of the array 67, 11, 78, 23, 10. Ideally, find a useful
// function that does this for you. (encourage them to find it. searching
// for stuff without external guidance is an important skill)
func exercise4c() {
	numbers := []int{67, 11, 78, 23, 10}
	slices.Sort(numbers)

	fmt.Println(numbers)

	// Now tell them to sort in the reverse order. This is perhaps a point where
	// they will not be sure how to proceed, cause htere's no reverse parameter
	// like in Python.

	// Trick: negative means ascending (but I always try them out, really)
	slices.SortStableFunc(numbers, func(a, b int) int {
		return b - a
	})

	fmt.Println(numbers)
}

// range will go over map elements in an unspecified order, so if elements are
// the same, you can expect the same

// Use a map to track the number of occurences of each string in the following
// array. Then, print each string, alongside the number of occurences.

// Alex, Armand, Andrei, Andrei, Iulia, Andrei, Andrei

func exercise4d() {
	names := []string{"Alex", "Armand", "Andrei", "Andrei", "Iulia"}
	freq := make(map[string]int)

	for _, name := range names {
		freq[name]++
	}

	// Please make sure to mention unspecified order
	for key, val := range freq {
		fmt.Printf("%s appears %d times\n", key, val)
	}

}

func main() {
	// exercise4c()
	exercise4d()
}
