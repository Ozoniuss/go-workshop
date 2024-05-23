package main

import "fmt"

/*
Next, let's look a bit into some of go's data types. There are many extensive
guides on Go's typesystem, but at the moment we'll be looking only at a few
of them: int, float64, bool and string.
*/

/*
Find the result of the following computation:
153963 + 3212 - 3435 - 11132
*/
func exercise1() {
	fmt.Println(153963 + 3212 - 3435 - 11132)

	x1 := 153963
	x2 := 3212
	x3 := -3435
	x4 := -11132

	fmt.Println(x1 + x2 + x3 + x4)
}

func main() {
	exercise1()
}
