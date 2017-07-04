package main

import (
	"bytes"
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

type SelectedCollectors map[string]struct{}

var collectors = []Collector{
	&cpu.Cpu{},
	&filesystem.FileSystem{},
	&memory.Memory{},
	&network.Network{},
	&platform.Platform{},
	&processes.Processes{},
}

var options struct {
	only     SelectedCollectors
	exclude  SelectedCollectors
	logLevel string
	version  bool
}

// version information filled in at build time
var (
	buildDate string
	gitCommit string
	gitBranch string
	goVersion string
)

func Collect() (result map[string]interface{}, err error) {
	result = make(map[string]interface{})

	for _, collector := range collectors {
		if shouldCollect(collector) {
			c, err := collector.Collect()
			if err != nil {
				fmt.Printf("[%s] %s", collector.Name(), err)
			}
			if c != nil {
				result[collector.Name()] = c
			}
		}
	}

	result["bodhi"] = versionMap()

	return
}

func versionMap() (result map[string]interface{}) {
	result = make(map[string]interface{})

	result["git_hash"] = gitCommit
	result["git_branch"] = gitBranch
	result["build_date"] = buildDate
	result["go_version"] = goVersion

	return
}

func versionString() string {
	var buf bytes.Buffer

	if gitCommit != "" {
		fmt.Fprintf(&buf, "Git hash: %s\n", gitCommit)
	}
	if gitBranch != "" {
		fmt.Fprintf(&buf, "Git branch: %s\n", gitBranch)
	}
	if buildDate != "" {
		fmt.Fprintf(&buf, "Build date: %s\n", buildDate)
	}
	if goVersion != "" {
		fmt.Fprintf(&buf, "Go Version: %s\n", goVersion)
	}

	return buf.String()
}

// Return whether we should collect on a given collector, depending on the parsed flags
func shouldCollect(collector Collector) bool {
	if _, ok := options.only[collector.Name()]; len(options.only) > 0 && !ok {
		return false
	}

	if _, ok := options.exclude[collector.Name()]; ok {
		return false
	}

	return true
}

func main() {
	if options.version {
		fmt.Printf("%s", versionString())
		os.Exit(0)
	}

	gohai, err := Collect()

	if err != nil {
		panic(err)
	}

	buf, err := json.Marshal(gohai)

	if err != nil {
		panic(err)
	}

	os.Stdout.Write(buf)
}
