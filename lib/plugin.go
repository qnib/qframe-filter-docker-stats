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
				go p.GetCpuMetrics(qcs)
				go p.GetMemoryMetrics(qcs)
				go p.GetNetworkMetrics(qcs)
			}
		}
	}
}

func (p *Plugin) GetCpuMetrics(qcs qtypes.ContainerStats) {
	stat := qcs.GetCpuStats()
	for _, m := range stat.ToMetrics(p.Name) {
		p.QChan.Data.Send(m)
	}
}

func (p *Plugin) GetMemoryMetrics(qcs qtypes.ContainerStats) {
	stat := qcs.GetMemStats()
	for _, m := range stat.ToMetrics(p.Name) {
		p.QChan.Data.Send(m)
	}
}

func (p *Plugin) GetNetworkMetrics(qcs qtypes.ContainerStats) {
	stat := qcs.GetNetStats()
	for _, m := range stat.ToMetrics(p.Name) {
		p.QChan.Data.Send(m)
	}
	aggStats := qtypes.NewNetStats(qcs.Base, qcs.GetContainer())
	for iface, _ := range qcs.Stats.Networks {
		stats := qcs.GetNetPerIfaceStats(iface)
		for _, m := range stats.ToMetrics(p.Name) {
			p.QChan.Data.Send(m)
		}
		aggStats = qtypes.AggregateNetStats("total", aggStats, stats)
	}
	for _, m := range aggStats.ToMetrics(p.Name) {
		p.QChan.Data.Send(m)
	}
}
