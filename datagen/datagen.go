package datagen

import (
	"fmt"
	"gopkg.in/raintank/schema.v1"
)

var modules map[string]Datagen

type Datagen interface {
	RegisterFlagSet()
	String() string
	GetData(int64) []*schema.MetricData
}

func Get(name string) Datagen {
	d, ok := modules[name]
	if !ok {
		panic(fmt.Sprintf("could not find datagen %q", name))
	}
	return d
}

func RegisterFlagSet() {
	for _, mod := range modules {
		mod.RegisterFlagSet()
	}
}
