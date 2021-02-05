FROM golang:latest

# install prometheus
RUN wget -nv https://github.com/prometheus/prometheus/releases/download/v2.23.0/prometheus-2.23.0.linux-amd64.tar.gz
RUN tar xvf prometheus-2.23.0.linux-amd64.tar.gz
RUN mv prometheus-2.23.0.linux-amd64 /etc/prometheus
RUN rm prometheus-2.23.0.linux-amd64.tar.gz 

# install grafana
RUN wget -nv https://dl.grafana.com/oss/release/grafana-7.0.0.linux-amd64.tar.gz
RUN tar xvf grafana-7.0.0.linux-amd64.tar.gz
RUN mv grafana-7.0.0 /etc/grafana
RUN rm grafana-7.0.0.linux-amd64.tar.gz

# install go-fuzz and go-fuzz-build
RUN go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

# setup working directory
RUN mkdir /go-fuzz-exporter
WORKDIR /go-fuzz-exporter

# copy git repositroy contents into working directory
COPY . .

# build go-fuzz-exporter
RUN make build
RUN mv go-fuzz-exporter /usr/local/bin

# go-fuzz-exporter
EXPOSE 4022

# prometheus
EXPOSE 9090

# grafana
EXPOSE 3000

ENTRYPOINT [ "/usr/local/bin/go-fuzz-exporter" ]