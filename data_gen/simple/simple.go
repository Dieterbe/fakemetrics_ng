package simple

import (
	"flag"
	"fmt"

	"gopkg.in/raintank/schema.v1"

	gc "github.com/rakyll/globalconf"
)

type Simple struct{}

var (
	keyCount  int
	keyPrefix string
)

func (s *Simple) RegisterFlagSet() {
	flags := flag.NewFlagSet(s.String(), flag.ExitOnError)
	flags.IntVar(&keyCount, "key_count", 100, "number of keys to generate")
	flags.StringVar(&keyPrefix, "key_prefix", "some.key.", "prefix for keys")
	gc.Register(s.String(), flags)
}

func (s *Simple) String() string {
	return "simple"
}

func (s *Simple) GetData(ts int64) []*schema.MetricData {
	metrics := make([]*schema.MetricData, keyCount)

	for i := 1; i <= keyCount; i++ {
		metrics[i-1] = &schema.MetricData{
			Name:   fmt.Sprintf(keyPrefix+"%d", i),
			Metric: fmt.Sprintf(keyPrefix+"%d", i),
			OrgId:  i,
			Value:  0,
			Unit:   "ms",
			Mtype:  "gauge",
			Tags:   []string{"some_tag", "ok", "k:2"},
			Time:   ts,
		}
	}
	return metrics
}
