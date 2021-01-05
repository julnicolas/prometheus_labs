# Lab 5

## Start lab env
### Run prometheus
Running the command below will start up a prometheus container loading `./prometheus.yml` as the config file.
``` sh
docker run --name prometheus -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 127.0.0.1:9090:9090 prom/prometheus
```

### Runthe test services
``` sh
docker build -t promdemo .
docker run -d --name promdemo1 -p 10000:10000 promdemo -listen-address=:10000
docker run -d --name promdemo2 -p 10001:10001 promdemo -listen-address=:10001
docker run -d --name promdemo3 -p 10002:10002 promdemo -listen-address=:10002
```

## Shutdown the environment
## Shutdown prometheus
``` sh
docker container stop prometheus
docker container stop promedemo1
docker container stop promedemo2
docker container stop promedemo3
```

## Clean environment
``` sh
docker container rm prometheus
docker container rm  promedemo1
docker container rm  promedemo2
docker container rm  promedemo3
```
