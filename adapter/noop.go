package adapter

import (
	"fmt"
)

type Noop struct {}

func (a Noop) Show(thing string) {
	fmt.Printf("noop called: thing=%s", thing)
	fmt.Println("")
}
