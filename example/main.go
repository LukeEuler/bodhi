package main

import (
	"fmt"

	"github.com/LukeEuler/bodhi"
)

func main() {
	result, err := bodhi.Collect()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	for key, value := range result {
		fmt.Println(key, value)
	}
}
