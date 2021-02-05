package metrics

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseLine(line string) (*LogLine, error) {
	line = strings.TrimSpace(line)

	if line == "" {
		return nil, fmt.Errorf("log line is empty")
	}

	fields := strings.Fields(line)

	if len(fields) != 19 {
		return nil, fmt.Errorf("unexpected number of fields found on log line: %d", len(fields))
	}

	// ------------------------------------------------------------------------------------------------------------------------------------------------
	// Log line fields for go-fuzz
	// ---------- -------- -------- ---- ------- --- ---- ----- --------- -- --------- ------- ------ -------- ------------- ------ ----- ------- -----
	// 0          1        2        3    4       5   6    7     8         9  10        11      12     13       14            15     16    17      18
	// ---------- -------- -------- ---- ------- --- ---- ----- --------- -- --------- ------- ------ -------- ------------- ------ ----- ------- -----
	// 2015/04/25 12:39:53 workers: 500, corpus: 186 (42s ago), crashers: 3, restarts: 1/8027, execs: 12009519 (121224/sec), cover: 2746, uptime: 1m39s

	var (
		// Ignored
		// dateStampStr         = fields[0]
		// timeStampStr         = fields[1]
		workersStr           = fields[3]
		corpusStr            = fields[5]
		corpusLastUpdatedStr = fields[6]
		crashersStr          = fields[9]
		// TODO(kent): restarts ratio
		// restartsStr          = fields[11]
		execsStr          = fields[13]
		execsPerSecondStr = fields[14]
		coverageStr       = fields[16]
		// uptimeStr         = fields[18]
	)

	var (
		err               error
		workers           int64
		corpus            int64
		corpusLastUpdated time.Duration
		crashers          int64
		execs             int64
		execsPerSecond    int64
		coverage          int64
	)

	workers, err = strconv.ParseInt(strings.TrimRight(workersStr, ","), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse workers: %w", err)
	}

	corpus, err = strconv.ParseInt(corpusStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse corpus: %w", err)
	}

	corpusLastUpdated, err = time.ParseDuration(strings.TrimLeft(corpusLastUpdatedStr, "("))
	if err != nil {
		return nil, fmt.Errorf("failed to parse corpus last updated duration: %w", err)
	}

	crashers, err = strconv.ParseInt(strings.TrimRight(crashersStr, ","), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse crashers: %w", err)
	}

	execs, err = strconv.ParseInt(execsStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse execs: %w", err)
	}

	execsPerSecond, err = strconv.ParseInt(strings.TrimRight(strings.TrimLeft(execsPerSecondStr, "("), "/sec),"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse execs per second: %w", err)
	}

	coverage, err = strconv.ParseInt(strings.TrimRight(coverageStr, ","), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse coverage: %w", err)
	}

	logLine := &LogLine{
		Workers:           float64(workers),
		Corpus:            float64(corpus),
		CorpusLastUpdated: corpusLastUpdated,
		Crashers:          float64(crashers),
		Execs:             float64(execs),
		ExecsPerSecond:    float64(execsPerSecond),
		Coverage:          float64(coverage),
	}

	return logLine, nil
}
