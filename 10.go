package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Json handling. Reading a json from a string, reading a json from a file, writing
a json to a file.

As a separate topic, mention the  Unmarshaler and Marshaler interfaces.
*/

// Let's start with a sample string. Explain const. Explain multiline stirng.
const programmerJson = `
{
	"name": "Alex",
	"language": "Go",
	"yoe": 10
}
`

// Add the struct tags only after showing that it doesn't work.
type WorkshopProgrammer struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	YoE      int    `json:"yoe"`
}

func exercise10a() {
	var p WorkshopProgrammer
	// To read a json into a struct, we need to use the json.Unmarhsal function
	err := json.Unmarshal([]byte(programmerJson), &p)
	if err != nil {
		fmt.Printf("could not read programmer: %s\n", err.Error())
	}

	fmt.Println("programmer", p)

	// Alternatively, a nicer method is to use a Decoder. This takes in a byte,
	// so it means we need the byte data. But we don't really want to convert
	// the data into a byte, if it's large, then struct. We want to use the
	// data right away. So we can use io.Reader
	p = WorkshopProgrammer{}

	// This takes in an io.reader. So let's create one
	r := strings.NewReader(programmerJson)
	dec := json.NewDecoder(r)
	err = dec.Decode(&p)
	if err != nil {
		fmt.Printf("could not read programmer: %s\n", err.Error())
	}

	fmt.Println("programmer", p)

	// Now, we can use this exact pattern on files. Let's see:
	p = WorkshopProgrammer{}
	f, err := os.Open("10.json")
	if err != nil {
		panic(err)
	}

	// Modify the original code instead.
	dec = json.NewDecoder(f)
	err = dec.Decode(&p)
	if err != nil {
		fmt.Printf("could not read programmer: %s\n", err.Error())
	}
	fmt.Println("programmer", p)
}

// Same exercise. But writing.
func exercise10b() {
	p := WorkshopProgrammer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10,
	}

	// Handle errors at workshop
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	// Now something more stylish
	data, _ = json.MarshalIndent(p, "", "  ")
	fmt.Println(string(data))

	// Let's do the same, but with an io.Writer. Let's create an encoder.
	enc := json.NewEncoder(os.Stdout) // can do to file as well.
	enc.SetIndent("", "-->")

	enc.Encode(p)
}

func main() {
	exercise10a()
	exercise10b()
}
