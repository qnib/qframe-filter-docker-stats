package qframe_filter_docker_stats

import (
	"C"
	"fmt"
	"github.com/zpatrick/go-config"
	"github.com/elastic/beats/metricbeat/module/docker/cpu"

	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-utils"
)

const (
	version = "0.0.0"
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
	myId := qutils.GetGID()
	bg := p.QChan.Data.Join()
	inputs := p.GetInputs()
	srcSuccess := p.CfgBoolOr("source-success", true)
	for {
		val := bg.Recv()
		switch val.(type) {
		case qtypes.QMsg:
			qm := val.(qtypes.QMsg)
			if qm.SourceID == myId {
				continue
			}
			if len(inputs) != 0 && !qutils.IsInput(inputs, qm.Source) {
				continue
			}
			if qm.SourceSuccess != srcSuccess {
				continue
			}
			switch qm.Data.(type) {
			case qtypes.ContainerStats:
				qm.Type = "filter"
				qm.Source = p.Name
				qm.SourceID = myId
				qm.SourcePath = append(qm.SourcePath, p.Name)
				cs := qm.Data.(qtypes.ContainerStats)
				cstat := cs.GetCpuStats()
				for _, m := range cstat.ToMetrics() {
					p.QChan.Data.Send(m)
				}
			}
		}
	}
}
