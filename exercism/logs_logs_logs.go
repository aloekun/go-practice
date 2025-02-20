package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation := '❗'
	serach := '🔍'
	weather := '☀'
	for _, char := range log {
		switch char {
		case recommendation:
			return "recommendation"
		case serach:
			return "search"
		case weather:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	resultLog := make([]byte, 0)
	for _, char := range log {
		if char == oldRune {
			resultLog = utf8.AppendRune(resultLog, newRune)
		} else {
			resultLog = utf8.AppendRune(resultLog, char)
		}
	}
	return string(resultLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
	if numberOfRunes <= limit {
		return true
	}
	return false
}
