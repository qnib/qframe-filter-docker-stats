package qframe_filter_docker_stats

import (
	"C"
	"fmt"
	"github.com/zpatrick/go-config"
	"github.com/qnib/qframe-types"
)

const (
	version = "0.1.1"
	pluginTyp = "filter"
)

type Plugin struct {
	qtypes.Plugin
}

func New(qChan qtypes.QChan, cfg config.Config, name string) (p Plugin, err error) {
	p = Plugin{
		Plugin: qtypes.NewNamedPlugin(qChan, cfg, pluginTyp, name, version),
	}
	return p, err
}

// Run fetches everything from the Data channel and flushes it to stdout
func (p *Plugin) Run() {
	p.Log("info", fmt.Sprintf("Start docker-stats filter v%s", p.Version))
	dc := p.QChan.Data.Join()
	inputs := p.GetInputs()
	p.Log("info", fmt.Sprintf("%v", inputs))
	srcSuccess := p.CfgBoolOr("source-success", true)
	for {
		select {
		case val := <- dc.Read:
			switch val.(type) {
			case qtypes.ContainerStats:
				qcs := val.(qtypes.ContainerStats)
				if qcs.IsLastSource(p.Name) {
					p.Log("debug", "IsLastSource() = true")
					continue
				}
				if len(inputs) != 0 && ! qcs.InputsMatch(inputs) {
					p.Log("debug", fmt.Sprintf("InputsMatch(%v) = false", inputs))
					continue
				}
				if qcs.SourceSuccess != srcSuccess {
					p.Log("debug", "qcs.SourceSuccess != srcSuccess")
					continue
				}
				// Process ContainerStats and create send multiple qtypes.Metrics
				//// CPUStats
				cstat := qcs.GetCpuStats()
				for _, m := range cstat.ToMetrics(p.Name) {
					p.QChan.Data.Send(m)
				}
				//// MemoryStats
				mstat := qcs.GetMemStats()
				for _, m := range mstat.ToMetrics(p.Name) {
					p.QChan.Data.Send(m)
				}
				//// MemoryStats
				nstat := qcs.GetNetStats()
				for _, m := range nstat.ToMetrics(p.Name) {
					p.QChan.Data.Send(m)
				}
			}
		}
	}
}
