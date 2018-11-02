package bodhi

import (
	"fmt"

	"github.com/LukeEuler/bodhi/cpu"
	"github.com/LukeEuler/bodhi/filesystem"
	"github.com/LukeEuler/bodhi/memory"
	"github.com/LukeEuler/bodhi/network"
	"github.com/LukeEuler/bodhi/platform"
	"github.com/LukeEuler/bodhi/processes"
)

// Collector ...
type Collector interface {
	Name() string
	Collect() (interface{}, error)
}

var collectors = []Collector{
	&cpu.CPU{},
	&filesystem.FileSystem{},
	&memory.Memory{},
	&network.Network{},
	&platform.Platform{},
	&processes.Processes{},
}

// Collect ...
func Collect() (result map[string]interface{}, err error) {
	result = make(map[string]interface{})

	for _, item := range collectors {
		c, err := item.Collect()
		if err != nil {
			fmt.Printf("[%s] %s", item.Name(), err)
		}
		if c != nil {
			result[item.Name()] = c
		}
	}
	return
}
