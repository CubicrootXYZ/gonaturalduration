package gonaturalduration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Duration time.Duration
	Text     string
}

func TestParser_Parse(t *testing.T) {
	assert := assert.New(t)
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "1 day"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "0 day"})
	testCases = append(testCases, TestCase{Duration: 24 * 24 * time.Hour, Text: "24 day"})
	testCases = append(testCases, TestCase{Duration: 99999 * 24 * time.Hour, Text: "99999 day"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-48) * time.Hour, Text: "-2 day"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "the quick brown fox"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "the 012 da lacy"})
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "there once was an attempt to 01 days generate a test"})

	for _, testCase := range testCases {
		duration := Parse(testCase.Text)
		assert.Equal(duration, testCase.Duration, testCase.Text+" failed.")
	}
}
