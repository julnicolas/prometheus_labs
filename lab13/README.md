# Chapter 13

## Lab 1
Run the demo services from chapter 5. Documentation is provided in `chapter5/README.md`.

Download consul :
``` sh
curl -o consul.zip https://releases.hashicorp.com/consul/1.9.3/consul_1.9.3_darwin_amd64.zip
unzip consul.zip
chmod 744 consul
mv consul consul_1.9.3
```

Run consul in development mode :
```sh
./consul_1.9.3/consul agent -dev -client='192.168.1.94 127.0.0.1' -config-file=./consul_1.9.3/demo_services.json
```
Note : call `killall -SIGHUP consul` to reload consul's config.

Run prometheus :
``` sh
docker run --name prometheus -d -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml -p 127.0.0.1:9090:9090 prom/prometheus
```
Note : to reload prom's config run `docker exec prometheus /bin/killall -HUP prometheus`

### Check services are available on consul
Browse to `http://192.168.1.94:8500`

### Check that consul's services are scrapped
Browse to `http://localhost:9090/targets`

## Lab 2
Target detection will be done reading /etc/prometheus/targets.yml.

Run prometheus :
``` sh
docker run --name prometheus -d -v $(pwd)/file_sd/:/etc/prometheus/ -p 127.0.0.1:9090:9090 prom/prometheus
```
