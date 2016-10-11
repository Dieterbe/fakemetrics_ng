package out

import (
	"fmt"
	"gopkg.in/raintank/schema.v1"
)

var modules map[string]Out

type Out interface {
	RegisterFlagSet()
	String() string
	GetChan() chan *schema.MetricData
	Start()
}

func Get(name string) Out {
	o, ok := modules[name]
	if !ok {
		panic(fmt.Sprintf("could not find output %q", name))
	}
	return o
}

func RegisterFlagSet() {
	for _, mod := range modules {
		mod.RegisterFlagSet()
	}
}
