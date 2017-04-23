# qframe-filter-docker-stats
Qframe filter refining docker-stats messages to metrics.

```bash
$ docker run -ti --name qframe-collector-docker-events --rm -e SKIP_ENTRYPOINTS=1 \
             -v ${GOPATH}/src/github.com/qnib/qframe-filter-docker-stats:/usr/local/src/github.com/qnib/qframe-filter-docker-stats \
             -v ${GOPATH}/src/github.com/qnib/qframe-types:/usr/local/src/github.com/qnib/qframe-types \
             -v ${GOPATH}/src/github.com/qnib/qframe-utils:/usr/local/src/github.com/qnib/qframe-utils \
             -w /usr/local/src/github.com/qnib/qframe-filter-docker-stats \
              qnib/uplain-golang bash
root@869129c63e79# govendor update github.com/qnib/qframe-types github.com/qnib/qframe-utils
```
