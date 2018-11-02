package cpu

// CPU ...
type CPU struct{}

const name = "cpu"

// Name implement Collector's
func (c *CPU) Name() string {
	return name
}

// Collect implement Collector's
func (c *CPU) Collect() (result interface{}, err error) {
	result, err = getCPUInfo()
	return
}
