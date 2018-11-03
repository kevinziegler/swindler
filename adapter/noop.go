package adapter

import (
	"fmt"
)

type Noop struct {
	Name string
	Thing string
}

func (a Noop) Show() {
	fmt.Printf("add called: name=%s, thing=%s", a.Name, a.Thing)
	fmt.Println("")
}
