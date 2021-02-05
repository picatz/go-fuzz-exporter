package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// 2015/04/25 12:39:53 workers: 500, corpus: 186 (42s ago), crashers: 3, restarts: 1/8027, execs: 12009519 (121224/sec), cover: 2746, uptime: 1m39s

func init() {
	prometheus.MustRegister(Workers)
	prometheus.MustRegister(Corpus)
	prometheus.MustRegister(Crashers)
	prometheus.MustRegister(Restarts)
	prometheus.MustRegister(Execs)
	prometheus.MustRegister(ExecsPerSecond)
	prometheus.MustRegister(Coverage)
	prometheus.MustRegister(Lines)
	prometheus.MustRegister(LinesSkipped)
}

var Workers = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "workers",
		Help:      "Total number of workers",
	},
	[]string{"case"},
)

var Corpus = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "corpus",
		Help:      "Total number of files in corpus",
	},
	[]string{"case"},
)

var CorpusLastUpdated = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "corpus_last_updated",
		Help:      "Duration in seconds since the corpus was last updated",
	},
	[]string{"case"},
)

var Crashers = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "crasers",
		Help:      "Total number of crashers",
	},
	[]string{"case"},
)

var Restarts = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "restart_rate",
		Help:      "Rate with which the fuzzer restarts test processes",
	},
	[]string{"case"},
)

var Execs = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "execs",
		Help:      "Total number of executions",
	},
	[]string{"case"},
)

var ExecsPerSecond = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "execs_per_second",
		Help:      "Total number of executions per second",
	},
	[]string{"case"},
)

var Coverage = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Name:      "coverage",
		Help:      "Total coverage",
	},
	[]string{"case"},
)

var Lines = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "lines",
		Help:      "Total log lines",
	},
	[]string{"case"},
)

var LinesSkipped = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "lines_skipped",
		Help:      "Total log lines skipped",
	},
	[]string{"case"},
)
