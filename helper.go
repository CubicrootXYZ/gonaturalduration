package gonaturalduration

import (
	"strings"
	"time"

	"github.com/CubicrootXYZ/gonaturalduration/converter"
)

// toDuration converts a string to a time.Duration. The integer will be multiolied with the unit
func stringToDuration(number string, unit time.Duration) time.Duration {
	multiplier, err := converter.StringToInt(number)
	if err != nil {
		return 0
	}
	return unit * time.Duration(multiplier)
}

// splitIntoWords splits the string into an array of words
func splitIntoWords(text string) []string {
	return strings.FieldsFunc(text, numberSpliter)
}

// numberSplitter determines the runes to split words on
func numberSpliter(r rune) bool {
	return r == '-' || r == ' ' || r == ',' || r == '\n'
}

// notIn checks if a unit is already in a list of units
func notIn(unit time.Duration, units *[]time.Duration) bool {
	for _, knownUnit := range *units {
		if knownUnit == unit {
			return false
		}
	}

	return true
}
