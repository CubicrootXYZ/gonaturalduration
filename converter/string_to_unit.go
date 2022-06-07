package converter

import "time"

var stringToUnit = map[string]time.Duration{
	"year":    time.Hour * 8760, // Ignoring leap years
	"years":   time.Hour * 8760,
	"month":   time.Hour * 720, // This is not really an exact time unit since months have varying lengths. Assuming the average to be 30.42 days we use 30 days to avoid shifting the time of day.
	"months":  time.Hour * 720,
	"week":    time.Hour * 168,
	"weeks":   time.Hour * 168,
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
