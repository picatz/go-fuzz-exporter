default: build

fuzz/build:
	@cd ./pkg/metrics && go-fuzz-build .

fuzz/run:
	@cd ./pkg/metrics && mkdir -p workdir && go-fuzz -workdir workdir -procs 1 metrics-fuzz.zip

fuzz/cleanup:
	@cd ./pkg/metrics && rm -rf workdir && rm metrics-fuzz.zip

example:
	@make fuzz/run 2>&1 | go run main.go example

docker/example:
	@cd ./pkg/metrics && go-fuzz -workdir workdir -procs 1 metrics-fuzz.zip 2>&1 | docker run --rm --name go-fuzz-exporter-example -i -p 9090:9090 -p 3000:3000 -p 4022:4022 go-fuzz-expoter:latest example

docker/example/metrics-stack: 
	@docker exec go-fuzz-exporter-example /usr/bin/make metrics

docker/build:
	@docker build -t go-fuzz-expoter .

prometheus:
	@/etc/prometheus/prometheus --config.file="prometheus.yml" 2>&1 > prometheus.log &

grafana:
	@/etc/grafana/bin/grafana-server -homepath="/etc/grafana" 2>&1 > grafana.log &

metrics: prometheus grafana

build:
	@go build -o go-fuzz-exporter .