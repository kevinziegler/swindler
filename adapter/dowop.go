package adapter

import (
	"fmt"
	"swindler/connections"
)

type Dowop struct {}

func (a Dowop) Add(thing string) {
	fmt.Printf("dowop called: thing=%s", thing)
	fmt.Println("")

	repo := connections.FileRepository{File: "test.json"}
	repo.Create()
}
