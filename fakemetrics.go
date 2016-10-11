package main

import (
	"flag"

	"gopkg.in/raintank/schema.v1"

	"github.com/OOM-Killer/fakemetrics_ng/datagen"
	"github.com/OOM-Killer/fakemetrics_ng/out"
	"github.com/OOM-Killer/fakemetrics_ng/timer"
)

var (
	confFile = flag.String("config",
		"fakemetrics.ini",
		"configuration file path")
)

func main() {
	flag.Parse()

	setupConfig()

	timer := timer.Get(timerMod)
	dataGen := datagen.Get(dataGenMod)
	out := out.Get(outMod)

	out.Start()
	outChan := out.GetChan()

	tick := timer.GetTicker()
	for range tick.C {
		go doTick(dataGen, outChan, timer.GetTimestamp())
	}
}

func doTick(dg datagen.DataGen, outChan chan *schema.MetricData, ts int64) {
	metrics := dg.GetData(ts)
	for _, m := range metrics {
		outChan <- m
	}
}
