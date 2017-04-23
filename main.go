package main

import (
	"log"
	"fmt"
	"time"

	"github.com/zpatrick/go-config"
	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-filter-docker-stats/lib"
	"github.com/fsouza/go-dockerclient"
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
		"filter.test.inputs": "dstats-in",
	}

	cfg := config.NewConfig(
		[]config.Provider{
			config.NewStatic(cfgMap),
		},
	)
	p, err := qframe_filter_docker_stats.New(qChan, *cfg, "test")
	if err != nil {
		log.Printf("[EE] Failed to create filter: %v", err)
		return
	}
	go p.Run()
	time.Sleep(2*time.Second)
	bg := qChan.Data.Join()
	qm := qtypes.NewQMsg("test", "dstats-in")
	qm.Msg = "Send Metrics of TestContainer1"
	oldPerCpuValuesTest := [][]uint64{{1, 9, 9, 5}, {1, 2, 3, 4}, {0, 0, 0, 0}}
	newPerCpuValuesTest := [][]uint64{{100000001, 900000009, 900000009, 500000005}, {101, 202, 303, 404}, {0, 0, 0, 0}}
	for index := range statsList {
		statsList[index].PreCPUStats.CPUUsage.PercpuUsage = oldPerCpuValuesTest[index]
		statsList[index].CPUStats.CPUUsage.PercpuUsage = newPerCpuValuesTest[index]
	}
	qm.Data = qtypes.ContainerStats{
		Container: docker.APIContainers{
			ID: "ContainerID",
			Names: []string{"ContainerName"},
			Command: "echo Huhu",
			Created: 0,
			Image: "debian:latest",
		},
		Stats: *statsList[0],
	}
	qChan.Data.Send(qm)
	for {
		qm = bg.Recv().(qtypes.QMsg)
		if qm.Source == "test" {
			continue
		}
		fmt.Printf("#### Received result filter for input: %s\n", qm.Msg)
		for k, v := range qm.KV {
			fmt.Printf("%+15s: %s\n", k, v)
		}
		break

	}
}
