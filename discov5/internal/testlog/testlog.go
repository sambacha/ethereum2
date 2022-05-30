package testlog

type Logger struct {
	lines []string
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Write(p []byte) (int, error) {
	l.lines = append(l.lines, string(p))
	return len(p), nil
}
func (l *Logger) Has(s string) bool {
	for _, line := range l.lines {
		if line == s {
			return true
		}
	}
	return false
}
func (l *Logger) Len() int {
	return len(l.lines)
}
