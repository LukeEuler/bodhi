package memory

// Memory ...
type Memory struct{}

const name = "memory"

// Name implement Collector's
func (m *Memory) Name() string {
	return name
}

// Collect implement Collector's
func (m *Memory) Collect() (result interface{}, err error) {
	result, err = getMemoryInfo()
	return
}
