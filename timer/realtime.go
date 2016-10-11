package timer

import (
	"flag"
	"time"

	gc "github.com/rakyll/globalconf"
)

type Realtime struct{}

var (
	interval int
)

func (r *Realtime) RegisterFlagSet() {
	flags := flag.NewFlagSet(r.String(), flag.ExitOnError)
	flags.IntVar(&interval, "interval", 100, "the metric interval")
	gc.Register(r.String(), flags)
}

func (r *Realtime) String() string {
	return "realtime"
}

func (r *Realtime) GetTicker() *time.Ticker {
	return time.NewTicker(time.Duration(interval) * time.Millisecond)
}

func (r *Realtime) GetTimestamp() int64 {
	return time.Now().Unix()
}
