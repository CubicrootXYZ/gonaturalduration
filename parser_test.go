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
	// Days
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "1 day"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "0 day"})
	testCases = append(testCases, TestCase{Duration: 24 * 24 * time.Hour, Text: "24 day"})
	testCases = append(testCases, TestCase{Duration: 99999 * 24 * time.Hour, Text: "99999 day"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-48) * time.Hour, Text: "-2 day"})
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "there once was an attempt to 01 days generate a test"})
	// Hours
	testCases = append(testCases, TestCase{Duration: 99 * time.Hour, Text: "äää 99 hours"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Hour, Text: "later the day we will meet or remind me in 1 hour"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-2) * time.Hour, Text: "remind me in -2 hours"})
	// Minutes
	testCases = append(testCases, TestCase{Duration: 10 * time.Minute, Text: "In 10 minutes"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Minute, Text: "In 1 minute"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-4578) * time.Minute, Text: "let's meet -4578 minute"})
	// Seconds
	testCases = append(testCases, TestCase{Duration: 10 * time.Second, Text: "In 10 seconds"})
	testCases = append(testCases, TestCase{Duration: 60 * time.Second, Text: "In 60 seconds"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-60) * time.Second, Text: "In -60 seconds"})
	// Failures
	testCases = append(testCases, TestCase{Duration: 0, Text: "the quick brown fox"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "the 012 da lacy"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "second minute day year"})
	// MIXED
	// Day + Hour + Minute + Second
	testCases = append(testCases, TestCase{Duration: 25*time.Hour + time.Minute + time.Second, Text: "We need this in 1 day, 1 hour, 1 minute and 1 second"})
	// Hour + Minute + Second
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute + time.Second, Text: "We need this in 1 hour, 1 second and 1 minute"})
	// Hour + Second
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Second, Text: "We need this in 1 hour and 1 second"})

	for _, testCase := range testCases {
		duration := Parse(testCase.Text)
		assert.Equal(duration, testCase.Duration, testCase.Text+" failed.")
	}
}
