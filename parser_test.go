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
	testCases := getDigitTestCases()

	for _, testCase := range testCases {
		duration := Parse(testCase.Text)
		assert.Equal(testCase.Duration, duration, testCase.Text+" failed.")
	}
}

func TestParser_ParseWIP(t *testing.T) {
	assert := assert.New(t)

	duration := Parse("second minute day year")
	assert.Equal(time.Hour, duration)
}

func TestParser_ParseNumber(t *testing.T) {
	assert := assert.New(t)
	testCases := getDigitTestCases()
	testCases = append(testCases, getNumberTestCases()...)

	for _, testCase := range testCases {
		duration := ParseNumber(testCase.Text)
		assert.Equal(testCase.Duration, duration, testCase.Text+" failed.")
	}
}

func Benchmark_ParseNumberMax(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseNumber("nine hundred ninety nine thousand nine hundred ninety nine days")
	}
}

func Benchmark_ParseNumberMin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseNumber("one second")
	}
}

func Benchmark_ParseNumberMixed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseNumber("ninety nine days, six hour, fourty three minutes and fourty five second")
	}
}

func Benchmark_ParseNumberMixedLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseNumber("This is a random text where I am trying to get more words into it. So in ninety nine days, six hour, fourty three minutes and fourty five second something might happen. Or not.")
	}
}

func Benchmark_ParseMixed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("99 days, 6 hour, 43 minute and 45 second")
	}
}

func Benchmark_ParseMixedLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("This is a random text where I am trying to get more words into it. So in 99 days, 6 hour, 43 minute and 45 second something might happen. Or not.")
	}
}

func Benchmark_ParseMax(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("999999 days")
	}
}

func Benchmark_ParseMin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("1 second")
	}
}

func FuzzParser_Parse(f *testing.F) {
	f.Add(string("this is my random sentence, where I ask you to bring me: milk, bread and a cat "))
	f.Fuzz(func(t *testing.T, s string) {
		// We always add a proper duration to the test string so we know the parser should find at least that, thus characters and combinations breaking the parsing will result in a zero duration
		res := Parse(s + " in 22 seconds")
		assert.Greaterf(t, int(res/time.Millisecond), 0, "The result is a zero time, parsing failed. String was: ", s)
	})
}

func FuzzParser_ParseNumber(f *testing.F) {
	f.Add(string("this is my random sentence, where I ask you to bring me: milk, bread and a cat "))
	f.Fuzz(func(t *testing.T, s string) {
		// We always add a proper duration to the test string so we know the parser should find at least that, thus characters and combinations breaking the parsing will result in a zero duration
		res := ParseNumber(s + " in twenty two seconds")
		assert.Greaterf(t, int(res/time.Millisecond), 0, "The result is a zero time, parsing failed. String was: ", s)
	})
}

// ##### DEFINE TESTCASES HERE #####

func getDigitTestCases() []TestCase {
	testCases := make([]TestCase, 0)
	// Years
	testCases = append(testCases, TestCase{Duration: 8760 * time.Hour, Text: "year"})
	testCases = append(testCases, TestCase{Duration: 8760 * time.Hour, Text: "  year"})
	testCases = append(testCases, TestCase{Duration: 8760 * time.Hour, Text: "1 year"})
	testCases = append(testCases, TestCase{Duration: 17520 * time.Hour, Text: "seems like 2 years are great!"})
	// Months
	testCases = append(testCases, TestCase{Duration: 720 * time.Hour, Text: "month"})
	testCases = append(testCases, TestCase{Duration: 720 * time.Hour, Text: " month"})
	testCases = append(testCases, TestCase{Duration: 720 * time.Hour, Text: "1 month"})
	testCases = append(testCases, TestCase{Duration: 1440 * time.Hour, Text: "seems like 2 months are great!"})
	// Weeks
	testCases = append(testCases, TestCase{Duration: 168 * time.Hour, Text: "1 week"})
	testCases = append(testCases, TestCase{Duration: 336 * time.Hour, Text: "seems like 2 weeks are great!"})
	// Days
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "day"})
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: " day"})
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "1 day"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "0 day"})
	testCases = append(testCases, TestCase{Duration: 24 * 24 * time.Hour, Text: "24 day"})
	testCases = append(testCases, TestCase{Duration: 99999 * 24 * time.Hour, Text: "99999 day"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-48) * time.Hour, Text: "-2 day"})
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "there once was an attempt to 01 days generate a test"})
	// Hours
	testCases = append(testCases, TestCase{Duration: 1 * time.Hour, Text: "hour"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Hour, Text: "  hour   "})
	testCases = append(testCases, TestCase{Duration: 99 * time.Hour, Text: "äää 99 hours"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Hour, Text: "later the day we will meet or remind me in 1 hour"})
	testCases = append(testCases, TestCase{Duration: -1 * time.Duration(-2) * time.Hour, Text: "remind me in -2 hours"})
	// Minutes
	testCases = append(testCases, TestCase{Duration: 10 * time.Minute, Text: "In 10 minutes"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Minute, Text: "minute"})
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
	// Year + Month + Week + Day
	testCases = append(testCases, TestCase{Duration: 538 * 24 * time.Hour, Text: "Hurry up! There is a storm in 1 year, 4 months, 6 weeks and 11 days"})
	// Day + Hour + Minute + Second
	testCases = append(testCases, TestCase{Duration: 25*time.Hour + time.Minute + time.Second, Text: "We need this in 1 day, 1 hour, 1 minute and 1 second"})
	// Hour + Minute + Second
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute + time.Second, Text: "We need this in 1 hour, 1 second and 1 minute"})
	// Hour + Second
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Second, Text: "We need this in 1 hour and 1 second"})
	// Special characters
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "Remember A, B, c in 1 hour and 1 minute"})
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "Remember c. In 1 hour and 1 minute"})
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "Remember !§%$%&%&(&/)(=?)'*#+_.,.:;:><<²³¼½¼½¬¬. In 1 hour and 1 minute"})
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "every hour and minute"})

	return testCases
}

func getNumberTestCases() []TestCase {
	testCases := make([]TestCase, 0)
	// Years
	testCases = append(testCases, TestCase{Duration: 8760 * time.Hour, Text: "one year"})
	testCases = append(testCases, TestCase{Duration: 17520 * time.Hour, Text: "seems like two years are great!"})
	// Months
	testCases = append(testCases, TestCase{Duration: 720 * time.Hour, Text: "one month"})
	testCases = append(testCases, TestCase{Duration: 1440 * time.Hour, Text: "seems like two months are great!"})
	// Weeks
	testCases = append(testCases, TestCase{Duration: 168 * time.Hour, Text: "one week"})
	testCases = append(testCases, TestCase{Duration: 336 * time.Hour, Text: "seems like two weeks are great!"})
	// Days
	testCases = append(testCases, TestCase{Duration: 24 * time.Hour, Text: "one day"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "zero day"})
	testCases = append(testCases, TestCase{Duration: 24 * 24 * time.Hour, Text: "twenty-four day"})
	testCases = append(testCases, TestCase{Duration: 999 * 24 * time.Hour, Text: "nine hundred ninety nine days"})
	testCases = append(testCases, TestCase{Duration: 48 * time.Hour, Text: "two days"})
	// Hours
	testCases = append(testCases, TestCase{Duration: 99 * time.Hour, Text: "äää ninety-nine hours"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Hour, Text: "later the day we will meet or remind me in one hour"})
	testCases = append(testCases, TestCase{Duration: 22 * time.Hour, Text: "remind me in twenty two hours"})
	testCases = append(testCases, TestCase{Duration: 122067 * time.Hour, Text: "one hundred twenty two thousand sixty seven hours"})
	testCases = append(testCases, TestCase{Duration: 222222 * time.Hour, Text: "let's do that in two hundred twenty two thousand and two hundred twenty two hours"})
	// Minutes
	testCases = append(testCases, TestCase{Duration: 10 * time.Minute, Text: "In ten minutes"})
	testCases = append(testCases, TestCase{Duration: 1 * time.Minute, Text: "In one minute"})
	testCases = append(testCases, TestCase{Duration: 4094 * time.Minute, Text: "let's meet four thousand and ninety four minute"})
	testCases = append(testCases, TestCase{Duration: 999999 * time.Minute, Text: "nine hundred ninety nine thousand nine hundred ninety nine minutes"})
	// Seconds
	testCases = append(testCases, TestCase{Duration: 10 * time.Second, Text: "In ten seconds"})
	testCases = append(testCases, TestCase{Duration: 60 * time.Second, Text: "In sixty seconds"})
	// Failures
	testCases = append(testCases, TestCase{Duration: 0, Text: "the quick brown fox"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "the 012 da lacy"})
	testCases = append(testCases, TestCase{Duration: 0, Text: "second minute day year"})
	// MIXED
	// Year + Month + Week + Day
	testCases = append(testCases, TestCase{Duration: 538 * 24 * time.Hour, Text: "Hurry up! There is a storm in one year, 4 months, six weeks and eleven days"})
	// Day + Hour + Minute + Second
	testCases = append(testCases, TestCase{Duration: 25*time.Hour + time.Minute + time.Second, Text: "We need this in one day, one hour, one minute and one second"})
	// Hour + Minute + Second
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute + time.Second, Text: "We need this in one hour, one second and one minute"})
	// Hour + Second
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Second, Text: "We need this in one hour and one second"})
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Second, Text: "We need this in 1 hour and one second"})
	// Special characters
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "Remember A, B, c in one hour and 1 minute"})
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "Remember c. In one hour and 1 minute"})
	testCases = append(testCases, TestCase{Duration: time.Hour + time.Minute, Text: "Remember !§%$%&%&(&/)(=?)'*#+_.,.:;:><<²³¼½¼½¬¬. In one hour and 1 minute"})

	return testCases
}
