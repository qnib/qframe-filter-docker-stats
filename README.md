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

## Start main.go to showcase functionality

```bash
$ 

```


