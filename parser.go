package gonaturalduration

import (
	"regexp"
	"time"

	"github.com/CubicrootXYZ/gonaturalduration/converter"
)

// Parse parses the first found duration in text and returns it as time.duration
func Parse(text string) time.Duration {
	var duration time.Duration

	// Patterns with first match group beeing a number
	defaultSearchPatterns := make(map[time.Duration]string)
	defaultSearchPatterns[24*time.Hour] = `(\w+) (day|days)`
	defaultSearchPatterns[time.Hour] = `(\w+) (hour|hours)`
	defaultSearchPatterns[time.Minute] = `(\w+) (minute|minutes)`
	defaultSearchPatterns[time.Second] = `(\w+) (second|seconds)`

	for unit, pattern := range defaultSearchPatterns {
		reg := regexp.MustCompile(pattern)
		match := reg.FindAllStringSubmatch(text, -1)
		if len(match) > 0 {
			duration += toDuration(match[0][1], unit)
		}
	}

	return duration
}

func toDuration(number string, unit time.Duration) time.Duration {
	multiplier, err := converter.StringToInt(number)
	if err != nil {
		return 0
	}
	return unit * time.Duration(multiplier)
}
