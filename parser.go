package gonaturalduration

import (
	"regexp"
	"time"

	"github.com/CubicrootXYZ/gonaturalduration/converter"
)

// Parse parses the first found duration in text and returns it as time.duration
func Parse(text string) time.Duration {

	regDay := regexp.MustCompile(`(\w+) (day|days)`)
	matchDay := regDay.FindAllStringSubmatch(text, -1)
	if len(matchDay) > 0 {
		return toDuration(matchDay[0][1], 24*time.Hour)
	}
	return 0
}

func toDuration(number string, unit time.Duration) time.Duration {
	multiplier, err := converter.StringToInt(number)
	if err != nil {
		return 0
	}
	return unit * time.Duration(multiplier)
}
