# Lab 12 - Relabeling

In this exercise, a relabeling rule is edited so that all metrics not starting by prometheus are dropped.
This applies only for the prometheus scrape job.

If the implementation is functional, all metrics exposed by prometheus start by prometheus. The following
promql command list all metrics exposed by prometheus which don't start by prometheus. If your implementation
is correct, only default special metrics (such as up) will be exposed :
``` promql
max({job="prometheus", __name__!~"prometheus.*"}) by(__name__)
```

## Run the exporter
```sh
curl -fLO https://github.com/prometheus/node_exporter/releases/download/v1.0.0/node_exporter-1.0.0.linux-amd64.tar.gz
tar -xzvf node_exporter-1.0.0.linux-amd64.tar.gz && rm -rf node_exporter-1.0.0.linux-amd64.tar.gz
./node_exporter-1.0.0.linux-amd64/node_exporter
```

## Run prometheus
Update the conf to your local IP wherever a hard coded IP is present.
``` sh
docker run --name prometheus -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 127.0.0.1:9090:9090 prom/prometheus
```

## Reload prom config
``` sh
docker exec prometheus /bin/killall -SIGHUP prometheus
```
