package cpu

type Cpu struct{}

const name = "cpu"

func (c *Cpu) Name() string {
	return name
}

func (c *Cpu) Collect() (result interface{}, err error) {
	result, err = getCpuInfo()
	return
}