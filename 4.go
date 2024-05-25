package main

import (
	"fmt"
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

// range will go over map elements in an unspecified order

func main() {

}
