FROM ubuntu:latest

WORKDIR promdemo

RUN apt update && \
apt install -y wget && \
wget https://github.com/juliusv/prometheus_demo_service/releases/download/0.0.4/prometheus_demo_service-0.0.4.linux.amd64.tar.gz && \
tar -xzvf prometheus_demo_service-0.0.4.linux.amd64.tar.gz

ENTRYPOINT ["./prometheus_demo_service"]
