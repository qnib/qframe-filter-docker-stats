package main

import (
	"log"
	"fmt"

	"github.com/zpatrick/go-config"
	"github.com/fsouza/go-dockerclient"
	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-filter-docker-stats/lib"
	"github.com/qnib/qframe-collector-docker-events/lib"
	"github.com/qnib/qframe-collector-docker-stats/lib"
)


var statsList = make([]docker.Stats, 3)


func Run(qChan qtypes.QChan, cfg config.Config, name string) {
	p, _ := qframe_filter_docker_stats.New(qChan, cfg, name)
	p.Run()
}

func main() {
	qChan := qtypes.NewQChan()
	qChan.Broadcast()
	cfgMap := map[string]string{
		"collector.docker-events.docker-host": "unix:///var/run/docker.sock",
		"filter.container-stats.inputs": "docker-stats",
		"log.level": "info",
	}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)
	// Start filter
	pfc, err := qframe_filter_docker_stats.New(qChan, *cfg, "container-stats")
	if err != nil {
		log.Printf("[EE] Failed to docker-stats filter: %v", err)
		return
	}
	go pfc.Run()
	// start docker-events
	pe, err := qframe_collector_docker_events.New(qChan, *cfg, "docker-events")
	if err != nil {
		log.Printf("[EE] Failed to docker-event collector: %v", err)
		return
	}
	go pe.Run()
	// start docker-stats
	p, err := qframe_collector_docker_stats.New(qChan, *cfg, "docker-stats")
	if err != nil {
		log.Printf("[EE] Failed to docker-stats collector: %v", err)
		return
	}
	go p.Run()
	dc := qChan.Data.Join()
	done := 5
	for {
		select {
		case msg := <-dc.Read:
			switch msg.(type) {
			case qtypes.Metric:
				qm := msg.(qtypes.Metric)
				if qm.IsLastSource("container-stats") {
					fmt.Printf("%s Metric %s: %v %s\n", qm.GetTimeRFC(), qm.Name, qm.Value, qm.GetDimensionList())
					done -= 1
				}
			}
		}
		if done == 0 {
			break
		}
	}
}
