# Lab 5

## Running prometheus
Running the command below will start up a prometheus container loading `./prometheus.yml` as the config file.
``` sh
docker run --name prometheus -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 127.0.0.1:9090:9090 prom/prometheus
```

## Shutdown prometheus
``` sh
docker container stop prometheus
```

## Clean environment
``` sh
docker container rm prometheus
```
