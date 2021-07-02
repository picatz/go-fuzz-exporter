package metrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseLineOld(t *testing.T) {
	exampleLine := `2015/04/25 12:39:53 workers: 500, corpus: 186 (42s ago), crashers: 3, restarts: 1/8027, execs: 12009519 (121224/sec), cover: 2746, uptime: 1m39s`

	logInterface, style, err := ParseLine(exampleLine)
	require.NoError(t, err)
	require.NotNil(t, logInterface)
	require.Equal(t, OldStyle, style)

	logLine, ok := logInterface.(*OldLogLine)
	require.True(t, ok)

	require.Equal(t, float64(500), logLine.Workers)
	require.Equal(t, float64(186), logLine.Corpus)
	require.Equal(t, time.Duration(42000000000), logLine.CorpusLastUpdated)
	require.Equal(t, float64(3), logLine.Crashers)
	require.Equal(t, float64(0), logLine.Restarts)
	require.Equal(t, 1.2009519e+07, logLine.Execs)
	require.Equal(t, float64(121224), logLine.ExecsPerSecond)
	require.Equal(t, float64(2746), logLine.Coverage)
}

func TestParseLineNew(t *testing.T) {
	exampleLine := `fuzzing, elapsed: 30.0s, execs: 3345 (111/sec), workers: 1, interesting: 10`

	logInterface, style, err := ParseLine(exampleLine)
	require.NoError(t, err)
	require.NotNil(t, logInterface)
	require.Equal(t, NewStyle, style)

	logLine, ok := logInterface.(*NewLogLine)
	require.True(t, ok)

	require.Equal(t, float64(1), logLine.Workers)
	require.Equal(t, float64(3345), logLine.Execs)
	require.Equal(t, float64(111), logLine.ExecsPerSecond)
	require.Equal(t, float64(10), logLine.Interesting)
}
