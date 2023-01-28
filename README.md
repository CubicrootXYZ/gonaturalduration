# go-natural-duration

[![GitHub license](https://img.shields.io/github/license/CubicrootXYZ/gonaturalduration)](https://github.com/CubicrootXYZ/gonaturalduration/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/CubicrootXYZ/gonaturalduration)](https://github.com/CubicrootXYZ/gonaturalduration/issues)
[![Actions Status](https://github.com/CubicrootXYZ/gonaturalduration/workflows/Main/badge.svg?branch=main)](https://github.com/CubicrootXYZ/gonaturalduration/actions)

Library for parsing time durations in natural language.

![mascot](images/mascot_100x100.png)

## Features

* Supports different kinds of durations:
    * years (365 days)
    * months (30 days)
    * weeks
    * days
    * hours
    * minutes
    * seconds
* Support for english language
* Parses from any string, even with context around the duration
* Digits up to at least 2.147.483.647
* Written numbers up to 999.999 (nine hundred ninety nine thousand nine hundred ninety nine)
* Automatic testing for code quality and functionality

## Installation

`go get github.com/CubicrootXYZ/gonaturalduration`

## Code-Examples

```
package main

import (
	"log"

	"github.com/CubicrootXYZ/gonaturalduration"
)

func main() {
	// Parse can only handle digits
	duration := gonaturalduration.Parse("in 1 hour")
	log.Printf("%v", duration)

	// ParseNumber can handle digits and written numbers
	duration = gonaturalduration.ParseNumber("let's do that in two hundred twenty two thousand and two hundred twenty two hours")
	log.Printf("%v", duration)
}
```

## Parser examples

* 1 day
    * "1 day"
    * "there once was an attempt to generate a duration of 1 day"
    * "this is 01 day"
    * "in one day"
* 1 hour
    * "1 hour"
    * "01 hour"
    * "please remind me in 1 hour"
    * "let's meet in one hour"
* -1 second
    * "this happend -1 second ago"
    * "-1 second"
    * "one second"
* 1 day and 2 minutes
    * "1 day and 2 minutes"
    * "we have 2 minutes and 1 day left"
    * "please to that in one day and two minutes"
    * "1 day and two minutes"
* 1 year and 2 months
    * "There will be an event in 1 year and 2 months"

## Benchmarks

```
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
Benchmark_ParseNumberMax-12               855056              1318 ns/op             232 B/op          8 allocs/op
Benchmark_ParseNumberMin-12              3179952               378.0 ns/op            48 B/op          3 allocs/op
Benchmark_ParseNumberMixed-12             706048              1660 ns/op             296 B/op          9 allocs/op
Benchmark_ParseNumberMixedLong-12         340768              3399 ns/op            1880 B/op         11 allocs/op
Benchmark_ParseMixed-12                    76251             15649 ns/op            1348 B/op         12 allocs/op
Benchmark_ParseMixedLong-12                20875             57201 ns/op            1510 B/op         13 allocs/op
Benchmark_ParseMax-12                     234073              4850 ns/op             336 B/op          3 allocs/op
Benchmark_ParseMin-12                     310310              3651 ns/op             336 B/op          3 allocs/op
```

## Contributing

Everyone wanting to contribute to this library is welcome. Add your ideas and bugs to the issue section. Open a pull request for additional features. 