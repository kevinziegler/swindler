package adapter

import (
	"fmt"
)

type Dowop struct {}

func (a Dowop) Add(thing string) {
	fmt.Printf("dowop called: thing=%s", thing)
	fmt.Println("")
}
