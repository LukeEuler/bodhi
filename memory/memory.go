package memory

type Memory struct{}

const name = "memory"

func (m *Memory) Name() string {
	return name
}

func (m *Memory) Collect() (result interface{}, err error) {
	result, err = getMemoryInfo()
	return
}
