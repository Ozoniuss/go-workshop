package main

import (
	"httpexercise/benchpool"
	"net/http"
)

// Let's define the following server. We're looking for the same
// characteristics. Particularly (ask them what they should think)

// Add a new programmer
// POST /programmers/

// List all programmers
// GET /programmers/

// Hire programmer for language
// POST /hire&language=Go&yoe=2

// Every request body will send json data.

/*
Start with ListenAndServer. Show them the difference using curl
Then, let's explore the handler interface.
*/

type AddProgrammerHandler struct {
	pool benchpool.BenchPool
}

func NewAddProgrammerHandler(pool benchpool.BenchPool) *AddProgrammerHandler {
	return &AddProgrammerHandler{
		pool: pool,
	}
}

func (a *AddProgrammerHandler) ServeHTTP() {

}

func main() {

	http.ListenAndServe(":6900", nil)
}
