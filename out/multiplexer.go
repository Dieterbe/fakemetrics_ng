package out

import (
	"gopkg.in/raintank/schema.v1"
)

type Multiplexer struct {
	in       chan *schema.MetricData
	outs     []Out
	outChans []chan *schema.MetricData
}

func (m *Multiplexer) Start() {
	m.in = make(chan *schema.MetricData)
	for _, out := range m.outs {
		out.Start()
		m.outChans = append(m.outChans, out.GetChan())
	}

	go m.loop()
}

func (m *Multiplexer) loop() {
	for {
		metric := <-m.in
		for _, c := range m.outChans {
			c <- metric
		}
	}
}

func (m *Multiplexer) AddOut(out *Out) {
	m.outs = append(m.outs, *out)
}

func (m *Multiplexer) GetChan() chan *schema.MetricData {
	if m.in == nil {
		panic("can't provide channel before starting")
	}
	return m.in
}

func (m *Multiplexer) String() string {
	return "multiplexer"
}

func (m *Multiplexer) RegisterFlagSet() {}
