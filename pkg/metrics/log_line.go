package metrics

import (
	"time"
)

type LogLine struct {
	Workers           float64
	Corpus            float64
	CorpusLastUpdated time.Duration
	Crashers          float64
	Restarts          float64
	Execs             float64
	ExecsPerSecond    float64
	Coverage          float64
}
