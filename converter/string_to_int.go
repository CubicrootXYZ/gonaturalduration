package converter

import (
	"regexp"
	"strconv"
	"strings"
)

// WordToInt maps words to integers
var wordToInt = map[string]int{
	"zero":      0,
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighten":   18,
	"nineteen":  19,
	"twenty":    20,
	"thirty":    30,
	"thourty":   40,
	"fifty":     50,
	"sixty":     60,
	"seventy":   70,
	"eighty":    80,
	"ninety":    90,
}

// Units indicates that the next word needs to me multiplied with the given factor
var units = map[string]int{
	"hundred":  100,
	"thousand": 1000,
}

// StringToInt parses the given string into an integer
func StringToInt(text string) (int, error) {
	// TODO also parse written numbers
	return strconv.Atoi(text)
}

// WordsToInt reads a written number or digits from a list of words. It stops after more than one non-numeric word. It can handle summed expressions like "onehundred and ten".
func WordsToInt(words []string) (int, error) {
	if len(words) < 1 {
		return 0, nil
	}

	// Check if it is a digit
	reg := regexp.MustCompile("^[0-9]+$")
	if reg.Match([]byte(words[len(words)-1])) {
		return StringToInt(words[len(words)-1])
	}

	// Not a digit so parse written numbers
	numbers := make(map[int][]int)
	number := 0
	nan := 0
	unitMultiplicator := 1

	for i := len(words) - 1; i >= 0; i-- {
		word := strings.ToLower(words[i])
		if value, ok := wordToInt[word]; ok {
			// it is a number
			numbers[unitMultiplicator] = append(numbers[unitMultiplicator], value)
			nan = 0
		} else if value, ok := units[word]; ok {
			// it is a multiplier
			if unitMultiplicator > value {
				unitMultiplicator = value * unitMultiplicator
			} else {
				unitMultiplicator = value
			}
			nan = 0
		} else {
			// it is neither
			nan++
		}

		if nan > 1 {
			break
		}
	}

	// Sum up all numbers with their respective multiplicators
	for multiplicator, nums := range numbers {
		number += multiplicator * sum(nums)
	}

	return number, nil
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
