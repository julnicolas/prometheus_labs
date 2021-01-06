# Lab 5

## Start lab env
### Set up the configuration file
In order to run the lab, you need to generate your `prometheus.yml`. This file needs your IP address.

To get your IP address please use one of the following commands :
``` sh
# Linux
ip a

# OS X
ifconfig
```

Then execute the command below to replace the token `<YOUR_IP>` with the address obtained from the previous step.
``` sh
sed 's/<YOUR_IP>/<WRITE YOUR IP HERE>/g' default_prometheus.yml > prometheus.yml
```
Note : `<WRITE YOUR IP HERE>` must be replaced by your IP address.

### Run prometheus
Running the command below will start up a prometheus container loading `./prometheus.yml` as the config file.
``` sh
docker run --name prometheus -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 127.0.0.1:9090:9090 prom/prometheus
```

### Run the test services
``` sh
docker build -t promdemo .
docker run --name promdemo1 -d -p 10000:10000 promdemo -listen-address=:10000
docker run --name promdemo2 -d -p 10001:10001 promdemo -listen-address=:10001
docker run --name promdemo3 -d -p 10002:10002 promdemo -listen-address=:10002
```

### Check everything is running
Wait a few seconds (15 seconds should be enough) then open `http://localhost:9090/targets` in your web browser. All targets should be up.
If you opened the page to early it is possible that some targets are not up yet. Wait a bit, then refresh the page.

## Shutdown the environment
``` sh
docker container stop prometheus
docker container stop promdemo1
docker container stop promdemo2
docker container stop promdemo3
```

## Clean environment
``` sh
docker container rm prometheus
docker container rm  promdemo1
docker container rm  promdemo2
docker container rm  promdemo3
```
