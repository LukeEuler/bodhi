package main

import (
	"encoding/json"
	"fmt"

	"github.com/LukeEuler/bodhi"
)

func main() {
	collects, err := bodhi.Collect()
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(collects)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
