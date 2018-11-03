package adapter

import (
	"fmt"
)

type Dowop struct {}

func (a Dowop) Show(thing string) {
	fmt.Printf("dowop called: thing=%s", thing)
	fmt.Println("")
}
