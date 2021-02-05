# go-fuzz-exporter

Custom prometheus exporter for [`go-fuzz`](https://github.com/dvyukov/go-fuzz/).

<p align="center">
  <img alt="Logo" src="screenshot.png?raw=true" />
</p>

## Install

```console
$ go get -u github.com/picatz/go-fuzz-exporter
```

## Example Usage

Using the harness `metrics-fuzz.zip` from the `github.com/picatz/go-fuzz-exporter/pkg/metrics` package.

```console
$ ./pkg/metrics
$ go-fuzz-build .
$ mkdir -p workdir 
$ go-fuzz -workdir workdir -procs 1 metrics-fuzz.zip 2>&1 | go-fuzz-exporter example
2021/02/04 19:08:26 starting go-fuzz-exporter server for "example" available at 127.0.0.1:4022
2021/02/04 19:08:27 workers: 1, corpus: 198 (0s ago), crashers: 0, restarts: 1/0, execs: 0 (0/sec), cover: 0, uptime: 3s
2021/02/04 19:08:30 workers: 1, corpus: 199 (1s ago), crashers: 0, restarts: 1/0, execs: 0 (0/sec), cover: 386, uptime: 6s
2021/02/04 19:08:33 workers: 1, corpus: 199 (4s ago), crashers: 0, restarts: 1/6875, execs: 41253 (4575/sec), cover: 386, uptime: 9s
2021/02/04 19:08:36 workers: 1, corpus: 199 (7s ago), crashers: 0, restarts: 1/8827, execs: 97103 (8081/sec), cover: 386, uptime: 12s
2021/02/04 19:08:39 workers: 1, corpus: 208 (0s ago), crashers: 0, restarts: 1/8989, execs: 152813 (10176/sec), cover: 386, uptime: 15s
...
^C
```

### Docker

```console
$ make docker/build
...
Successfully tagged go-fuzz-expoter:latest
$ go-fuzz -workdir workdir -procs 1 metrics-fuzz.zip 2>&1 | docker run -i go-fuzz-expoter:latest example
...
```
