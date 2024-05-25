package main

import (
	"fmt"
)

/*
Next, let's look a bit into some of go's data types. There are many extensive
guides on Go's typesystem, but at the moment we'll be looking only at a few
of them: int, float64, bool and string.
*/

/*
Find the result of the following computation:
153963 + 3212 - 3435 - 11132
*/
func exercise2a() {
	fmt.Println(153963 + 3212 - 3435 - 11132)

	x1 := 153963
	x2 := 3212
	x3 := -3435
	x4 := -11132

	fmt.Println(x1 + x2 + x3 + x4)
}

/*
Now they should be able to do this exercise and get the correct answer:
34253 + 23 - 11242 - 9990

But this time around, explain the var notation.
*/
func exercise2b() {
	fmt.Println(34253 + 23 - 11242 - 9990)

	var x1, x2, x3, x4 int

	x1 = 34253
	x2 = 23
	x3 = -11242
	x4 = -9990

	fmt.Println(x1 + x2 + x3 + x4)
}

/*
Let's do a similar ones with float:
11.5 * 3 + 4.5 / 10
*/
func exercise2c() {

	// They should write this to get used to the notation and practice it.
	var x1, x2, x3, x4 float64
	x1 = 11.5
	x2 = 3
	x3 = 4.5
	x4 = 10

	fmt.Println(x1*x2 + x3/x4)
}

/*
Let's do the same, but with the short notation.
*/
func exercise2d() {

	// link them this https://go101.org/article/basic-types-and-value-literals.html
	// and tell them to look for Floating point value literals.

	x1 := 11.5
	x2 := 3e0
	x3 := 4.5
	x4 := 10e0

	fmt.Println(x1*x2 + x3/x4)
}

// boolean exercise, with some and, or, not. This is the point to introduce
// things like ==, <, >, etc. Perhaps store these in variables.

// string exercise, with concatenation
// maybe a conversion between string and int here

// link the strings standard library. At this point you should show them how
// to go inside a function to read its definition

// maybe explain string builder?

func exercise2e() {
}

func main() {
	exercise2a()
	exercise2b()
	exercise2c()
	exercise2d()
}
