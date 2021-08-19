package converter

import "strconv"

// StringToInt parses the given string into an integer
func StringToInt(text string) (int, error) {
	// TODO also parse written numbers
	return strconv.Atoi(text)
}
