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
Benchmark_ParseNumberMax-12               283080              4548 ns/op            2601 B/op         42 allocs/op
Benchmark_ParseNumberMin-12               350337              3074 ns/op            2417 B/op         37 allocs/op
Benchmark_ParseNumberMixed-12              93307             13945 ns/op            9852 B/op        146 allocs/op
Benchmark_ParseNumberMixedLong-12          80961             15239 ns/op           11533 B/op        148 allocs/op
Benchmark_ParseMixed-12                    76900             15187 ns/op            1347 B/op         12 allocs/op
Benchmark_ParseMixedLong-12                21332             55825 ns/op            1344 B/op         12 allocs/op
Benchmark_ParseMax-12                     243813              4662 ns/op             337 B/op          3 allocs/op
Benchmark_ParseMin-12                     315446              3572 ns/op             337 B/op          3 allocs/op
```

## Contributing

Everyone wanting to contribute to this library is welcome. Add your ideas and bugs to the issue section. Open a pull request for additional features. 