package connections

import (
	"encoding/json"
	"fmt"
)

type FileRepository struct {
	File string
}

func (repo FileRepository) Create() {
	repo.loadConnections()
}

func (repo FileRepository) loadConnections() {
	data := []byte(` {"name": "test", "adapter": "dowop", "configuration": {} }`)

	var connections ConnectionProfile
	err := json.Unmarshal([]byte(data), &connections)

	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Println(connections)
}
