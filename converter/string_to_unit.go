package converter

import "time"

var stringToUnit = map[string]time.Duration{
	"day":     time.Hour * 24,
	"days":    time.Hour * 24,
	"hour":    time.Hour,
	"hours":   time.Hour,
	"minute":  time.Minute,
	"minutes": time.Minute,
	"second":  time.Second,
	"seconds": time.Second,
}

// StringToUnit converts the given word to a known unit or 0
func StringToUnit(word string) time.Duration {
	if unit, ok := stringToUnit[word]; ok {
		return unit
	}

	return 0
}
