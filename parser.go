package gonaturalduration

import (
	"regexp"
	"time"

	"github.com/CubicrootXYZ/gonaturalduration/converter"
)

// Parse parses the first found duration in text and returns it as time.Duration. Can only handle digits not written numbers but is faster.
func Parse(text string) time.Duration {
	var duration time.Duration

	// Patterns with first match group beeing a number
	defaultSearchPatterns := make(map[time.Duration]string)
	defaultSearchPatterns[8760*time.Hour] = `(\w+) (year|years)`
	defaultSearchPatterns[720*time.Hour] = `(\w+) (month|months)`
	defaultSearchPatterns[168*time.Hour] = `(\w+) (week|weeks)`
	defaultSearchPatterns[24*time.Hour] = `(\w+) (day|days)`
	defaultSearchPatterns[time.Hour] = `(\w+) (hour|hours)`
	defaultSearchPatterns[time.Minute] = `(\w+) (minute|minutes)`
	defaultSearchPatterns[time.Second] = `(\w+) (second|seconds)`

	for unit, pattern := range defaultSearchPatterns {
		reg := regexp.MustCompile(pattern)
		match := reg.FindAllStringSubmatch(text, -1)
		if len(match) > 0 {
			duration += stringToDuration(match[0][1], unit)
		}
	}

	return duration
}

// ParseNumber parses the first found duration in the text and returns it as time.Duration. Can handle digits as well as written numbers.
func ParseNumber(text string) time.Duration {
	var duration time.Duration
	words := splitIntoWords(text)

	lastMatch := 0
	parsedUnits := make([]time.Duration, 0)
	for i := 0; i < len(words); i++ {
		unit := converter.StringToUnit(words[i])
		if unit != 0 && notIn(unit, &parsedUnits) {
			multiplier, err := converter.WordsToInt(words[lastMatch:i])

			if err == nil {
				duration += unit * time.Duration(multiplier)
			}

			parsedUnits = append(parsedUnits, unit)
			lastMatch = i
		}
	}

	return duration
}
