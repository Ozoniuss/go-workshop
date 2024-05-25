package main

import "fmt"

/*
Pointers. No need to go in depth about pointers. Perhaps the only really
relevant rules here for using pointers:
- use pointers if your struct is huge
- use pointers if you want to mutate variables inside functions
- use pointers if you want to share the same variable accross functions
- DON'T use pointers for things that are pointers under the hood, like slices,
interfaces, channels

Note that functions are a prerequitsite!!!

https://tour.ardanlabs.com/tour/eng/pointers/1
https://go101.org/article/pointer.html
*/

// Write a function that adds 1 to an integer, without returning the original
// integer. (below are personal notes)

// Also write a function that does this by value to show them the difference.
// Write these functions outside the main/exercise function so they understand
// the difference better when calling them.
func changeByValue(n int) {
	n += 1
}
func ChangeByPointer(n *int) {
	*n += 1
}

func exercise6a() {
	n := 100

	changeByValue(n)
	fmt.Println(n)
	ChangeByPointer(&n)
	fmt.Println(n)
}

func main() {
	exercise6a()
}
