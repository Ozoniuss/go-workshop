/*
Let's set up the go main folder. The most basic debugging tool we have is
printing data on the screen, so let's begin with that to build us our
playground. Our first task is writing a program that prints hello.

- explain basic file structure: package, import, entrypoint
- explain basic go toolchain: go run 1.go
*/
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
