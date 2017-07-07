package main

import (
	"encoding/json"
	"fmt"
	"os"

	"bodhi/cpu"
	"bodhi/filesystem"
	"bodhi/memory"
	"bodhi/network"
	"bodhi/platform"
	"bodhi/processes"
)

type Collector interface {
	Name() string
	Collect() (interface{}, error)
}

var collectors = []Collector{
	&cpu.Cpu{},
	&filesystem.FileSystem{},
	&memory.Memory{},
	&network.Network{},
	&platform.Platform{},
	&processes.Processes{},
}

func Collect() (result map[string]interface{}, err error) {
	result = make(map[string]interface{})

	for _, collector := range collectors {
		c, err := collector.Collect()
		if err != nil {
			fmt.Printf("[%s] %s", collector.Name(), err)
		}
		if c != nil {
			result[collector.Name()] = c
		}
	}
	return
}

func main() {
	result, err := Collect()

	if err != nil {
		panic(err)
	}

	buf, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	os.Stdout.Write(buf)
}
