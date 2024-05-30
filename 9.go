package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
Interfaces. Part 2. Exploring standard library useful interfaces. Implementing
one or two of them.

Start by giving a speech about how in Go, a lot of things are done via interfaces.

Useful interfaces to know about: io.Reader, http.Handler (which we will implement
later), json.Marshaler/Unmarshaler, net.Conn, driver.Driver, Context
*/

// Let's explore some file operations. Let's create a file and write some text.

func exercise9a() {
	f, err := os.Open("9.txt")
	if err != nil {
		// printing the error
		fmt.Printf("could not open file: %s\n", err.Error())
		return
	}
	// Introduce defer. Main use case is cleanup of resources.
	defer f.Close()

	// Now let's read something from the file. Let's see the available methods.
	// This is the point to explain that file implements io.Reader, which is
	// the interface used for string operations in Go.
	// f.Read()

	// Let's read the entire contents at once. We have a helper for that.
	// Also make sure to walk them over the ReadAll implementation. It's
	// crucial to be in the habit of knowing to explore functions.
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("could not read file content: %s\n", err.Error())
	}

	fmt.Println(string(data))

	// Alright, now it's time to write some string to a file.
}

func exercise9b() {
	// Let's explore creating a file. Make sure to go into the function.

	f, err := os.Create("91.txt")
	if err != nil {
		fmt.Printf("could not create file: %s\n", err.Error())
	}
	defer f.Close()

	// Two ways of writing to this file. The second one will be interersting
	// to us.
	f.WriteString("hello?")
	io.WriteString(f, "hello")
}

// Last thing we'll do now is a simple advent of code exercise. Part a of
//https://adventofcode.com/2022/day/1

// In this we'll learn a nice pattern of reading a file line by line.
// ** based on time, either do this or max of numbers in a file. Introduce
// scanner, similar to before.

func exercise9c() {

}

func getCaloriesOfTopElf() (maxCalories int32, maxLine int32, err error) {

	/* Part 1 */

	// Loop through the entire list and use a counter to determine the number
	// of calories carried by every elf.

	// Keep the max in a variable, which is reset if the counter surpasses it.

	var calorieCounter int32 = 0
	var lineNumber int32 = 0

	maxCalories = 0
	maxLine = 0

	f, err := os.Open("input.txt")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	// Initialize a new scanner to read the file line by line.
	scanner := bufio.NewScanner(f)

	// ScanLines uses newline as the line splitter (should be the default
	// splitter set by NewScanner).
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineNumber += 1
		value := scanner.Text()

		// All food calories of the current Elf had been read.
		if value == "" {
			if calorieCounter > maxCalories {
				maxCalories = calorieCounter
				maxLine = lineNumber
			}
			calorieCounter = 0
		} else {
			// Still reading from current elf.
			calories, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return 0, 0, err
			}

			// Increase the calories counter for the current elf.
			calorieCounter += int32(calories)
		}
	}

	// Address EOF special case, calorieCounter stores the value of the last
	// elf
	if calorieCounter > maxCalories {
		maxCalories = calorieCounter
		maxLine = lineNumber
	}

	return maxCalories, maxLine, nil
}

func main() {
	exercise9a()
}
