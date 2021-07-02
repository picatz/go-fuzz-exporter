package metrics

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

func Collect(caseName string, src io.Reader) error {
	scanner := bufio.NewScanner(src)

	labels := prometheus.Labels{"case": caseName}

	for scanner.Scan() {
		line := scanner.Text()
		Lines.With(labels).Inc()
		logInterface, style, err := ParseLine(line)
		if err != nil {
			LinesSkipped.With(labels).Inc()
		} else {
			switch style {
			case OldStyle:
				logLine := logInterface.(*OldLogLine)
				Workers.With(labels).Set(logLine.Workers)
				Corpus.With(labels).Set(logLine.Corpus)
				CorpusLastUpdated.With(labels).Set(logLine.CorpusLastUpdated.Seconds())
				Crashers.With(labels).Set(logLine.Crashers)
				Execs.With(labels).Set(logLine.Execs)
				ExecsPerSecond.With(labels).Set(logLine.ExecsPerSecond)
				Coverage.With(labels).Set(logLine.Coverage)
			case NewStyle:
				logLine := logInterface.(*NewLogLine)
				Workers.With(labels).Set(logLine.Workers)
				Interesting.With(labels).Set(logLine.Interesting)
				Execs.With(labels).Set(logLine.Execs)
				ExecsPerSecond.With(labels).Set(logLine.ExecsPerSecond)
			}
		}
		fmt.Fprintln(os.Stderr, line)
	}

	return scanner.Err()
}
