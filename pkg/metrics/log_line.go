package metrics

import (
	"time"
)

type LogLine interface {
	IsLogLine()
}

type OldLogLine struct {
	LogLine
	Workers           float64
	Corpus            float64
	CorpusLastUpdated time.Duration
	Crashers          float64
	Restarts          float64
	Execs             float64
	ExecsPerSecond    float64
	Coverage          float64
}

// fuzzing, elapsed: 30.0s, execs: 3345 (111/sec), workers: 1, interesting: 7
type NewLogLine struct {
	LogLine
	Workers        float64
	Interesting    float64
	Execs          float64
	ExecsPerSecond float64
}
