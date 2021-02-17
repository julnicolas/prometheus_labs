# Lab 7 - Node exporter

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
If prometheus was running locally, it would be possible to reload its config sending the `SIGHUP` signal to the prometheus process.
