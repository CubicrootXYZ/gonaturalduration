package gonaturalduration

import (
	"regexp"
	"strings"
	"time"

	"github.com/CubicrootXYZ/gonaturalduration/converter"
)

// Patterns with first match group beeing a number
var defaultSearchPatterns = map[time.Duration]*regexp.Regexp{
	8760 * time.Hour: regexp.MustCompile(`(\w+) (year|years)`),
	720 * time.Hour:  regexp.MustCompile(`(\w+) (month|months)`),
	168 * time.Hour:  regexp.MustCompile(`(\w+) (week|weeks)`),
	24 * time.Hour:   regexp.MustCompile(`(\w+) (day|days)`),
	time.Hour:        regexp.MustCompile(`(\w+) (hour|hours)`),
	time.Minute:      regexp.MustCompile(`(\w+) (minute|minutes)`),
	time.Second:      regexp.MustCompile(`(\w+) (second|seconds)`),
}

// Parse parses the first found duration in text and returns it as time.Duration. Can only handle digits not written numbers but is faster.
func Parse(text string) time.Duration {
	text = strings.ToLower(text)
	var duration time.Duration

	for unit, pattern := range defaultSearchPatterns {
		match := pattern.FindAllStringSubmatch(text, -1)
		if len(match) > 0 {
			duration += stringToDuration(match[0][1], unit)
		}
	}

	return duration
}

// ParseNumber parses the first found duration in the text and returns it as time.Duration. Can handle digits as well as written numbers.
func ParseNumber(text string) time.Duration {
	text = strings.ToLower(text)
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
