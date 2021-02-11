// +build gofuzz

package metrics

func FuzzParseLine(data []byte) int {
	logLine, err := ParseLine(string(data))
	if err != nil {
		return 0
	}
	if logLine.Workers == 0 {
		return 0
	}
	if logLine.Corpus == 0 {
		return 0
	}
	if logLine.CorpusLastUpdated == 0 {
		return 0
	}
	if logLine.Crashers == 0 {
		return 0
	}
	if logLine.Execs == 0 {
		return 0
	}
	if logLine.ExecsPerSecond == 0 {
		return 0
	}
	if logLine.Coverage == 0 {
		return 0
	}
	return 1
}
