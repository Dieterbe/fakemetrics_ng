package timer

import (
	"fmt"
	"time"
)

var modules map[string]Timer

type Timer interface {
	RegisterFlagSet()
	String() string
	GetTicker() *time.Ticker
	GetTimestamp() int64
}

func Get(name string) Timer {
	t, ok := modules[name]
	if !ok {
		panic(fmt.Sprintf("could not find timer %q", name))
	}
	return t
}

func RegisterFlagSet() {
	for _, mod := range modules {
		mod.RegisterFlagSet()
	}
}
