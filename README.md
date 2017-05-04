# qframe-filter-docker-stats
Qframe filter refining docker-stats messages to metrics.

## Start Dev Container

```bash
$ docker run -ti --name qframe-filter-docker-stats --rm -e SKIP_ENTRYPOINTS=1 \
             -v ${GOPATH}/src/github.com/qnib/qframe-collector-docker-events:/usr/local/src/github.com/qnib/qframe-collector-docker-events \
             -v ${GOPATH}/src/github.com/qnib/qframe-collector-docker-stats:/usr/local/src/github.com/qnib/qframe-collector-docker-stats \
             -v ${GOPATH}/src/github.com/qnib/qframe-filter-docker-stats:/usr/local/src/github.com/qnib/qframe-filter-docker-stats \
             -v ${GOPATH}/src/github.com/qnib/qframe-types:/usr/local/src/github.com/qnib/qframe-types \
             -v ${GOPATH}/src/github.com/qnib/qframe-utils:/usr/local/src/github.com/qnib/qframe-utils \
             -v /var/run/docker.sock:/var/run/docker.sock \
             -w /usr/local/src/github.com/qnib/qframe-filter-docker-stats \
              qnib/uplain-golang bash
$ govendor update github.com/qnib/qframe-collector-docker-events/lib \
                  github.com/qnib/qframe-collector-docker-stats/lib \
                  github.com/qnib/qframe-filter-docker-stats/lib \
                  github.com/qnib/qframe-types \
                  github.com/qnib/qframe-utils
$ govendor fetch +m
```

## Start container to showcase functionality

The `Dockerfile` build a container that runs the example implementation as a standalone daemon.

```bash
$ docker run -ti -v /var/run/docker.sock:/var/run/docker.sock qnib/$(basename $(pwd))
 > execute CMD 'qframe-filter-docker-stats'
 2017/05/04 15:08:47 [II] Dispatch broadcast for Back, Data and Tick
 2017/05/04 15:08:47 [  INFO] container-stats >> Start docker-stats filter v0.1.0
 2017/05/04 15:08:47 [  INFO] docker-events >> Start docker-events collector v0.2.1
 2017/05/04 15:08:47 [  INFO] container-stats >> [docker-stats]
 2017/05/04 15:08:48 [  INFO] docker-events >> Connected to 'moby' / v'17.05.0-ce-rc1'
 2017/05/04 15:08:48 [  INFO] docker-stats >> Connected to 'moby' / v'17.05.0-ce-rc1' (SWARM: active)
 2017/05/04 15:08:48 [  INFO] docker-stats >> Currently running containers: 4
 2017/05/04 15:08:48 [II] Start listener for: 'elated_keller' [03f883aa96b5ea0433419b322328af40c44d54cfe8be49563f7f59b29946ad11]
 2017/05/04 15:08:48 [II] Start listener for: 'qframe-filter-docker-stats' [f8e94b5c3966766b3009e97b3ee737128fe413c87d2af83fb11336f3357f5917]
 2017/05/04 15:08:48 [II] Start listener for: 'gcollect_influxdb.1.ues4sdm7vzmhtzopkc08qb8fc' [9da2ac9a0db27f30bd913176c2177e588f0eae84c7cd885200a7391c9e6e6d72]
 2017/05/04 15:08:48 [II] Start listener for: 'gcollect_frontend.1.17u6xcak5rogijggdzebeouiu' [ea2ff849cbd72164fd965caa944b8e0f125afccf8605bd260fa08b65d10daf2d]
 2017/05/04 15:08:49 [  INFO] container-stats >> Received ContainerStats
 2017-05-04T15:08:49.640021+00:00 Metric usage_kernel_percent: 0.04 container_id=03f883aa96b5ea0433419b322328af40c44d54cfe8be49563f7f59b29946ad11,container_name=elated_keller,image_name=qnib/qframe-filter-docker-stats,command=/usr/local/bin/entrypoint.sh#qframe-filter-docker-stats,created=1493910527
 2017-05-04T15:08:49.640021+00:00 Metric usage_user_percent: 0.05 container_id=03f883aa96b5ea0433419b322328af40c44d54cfe8be49563f7f59b29946ad11,container_name=elated_keller,image_name=qnib/qframe-filter-docker-stats,command=/usr/local/bin/entrypoint.sh#qframe-filter-docker-stats,created=1493910527
 2017-05-04T15:08:49.640021+00:00 Metric system_usage_percent: 161199.21 image_name=qnib/qframe-filter-docker-stats,command=/usr/local/bin/entrypoint.sh#qframe-filter-docker-stats,created=1493910527,container_id=03f883aa96b5ea0433419b322328af40c44d54cfe8be49563f7f59b29946ad11,container_name=elated_keller
 2017/05/04 15:08:49 [  INFO] container-stats >> Received ContainerStats
 2017-05-04T15:08:49.65478+00:00 Metric usage_kernel_percent: 242.86 container_id=f8e94b5c3966766b3009e97b3ee737128fe413c87d2af83fb11336f3357f5917,container_name=qframe-filter-docker-stats,image_name=qnib/uplain-golang,command=/usr/local/bin/entrypoint.sh#bash,created=1493904515
 2017-05-04T15:08:49.65478+00:00 Metric usage_user_percent: 423.72 container_id=f8e94b5c3966766b3009e97b3ee737128fe413c87d2af83fb11336f3357f5917,container_name=qframe-filter-docker-stats,image_name=qnib/uplain-golang,command=/usr/local/bin/entrypoint.sh#bash,created=1493904515
```


