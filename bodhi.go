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
