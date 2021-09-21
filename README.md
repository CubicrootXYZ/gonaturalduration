# go-natural-duration

[![GitHub license](https://img.shields.io/github/license/CubicrootXYZ/gonaturalduration)](https://github.com/CubicrootXYZ/gonaturalduration/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/CubicrootXYZ/gonaturalduration)](https://github.com/CubicrootXYZ/gonaturalduration/issues)
[![Actions Status](https://github.com/CubicrootXYZ/gonaturalduration/workflows/Main/badge.svg?branch=main)](https://github.com/CubicrootXYZ/gonaturalduration/actions)

Library for parsing time durations in natural language.

## Features

* Supports different kinds of durations:
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
	"fmt"
	"log"

	"github.com/CubicrootXYZ/gonaturalduration"
)

func main() {
	// Parse can only handle digits
	duration := gonaturalduration.Parse("in 1 hour")
	log.Print(fmt.Sprintf("%v", duration))

	// ParseNumber can handle digits and written numbers
	duration = gonaturalduration.ParseNumber("let's do that in two hundred twenty two thousand and two hundred twenty two hours")
	log.Print(fmt.Sprintf("%v", duration))
}
```

## Parser examples

* 1d (1 day)
    * "1 day"
    * "there once was an attempt to generate a duration of 1 day"
    * "this is 01 day"
    * "in one day"
* 1h (1 hour)
    * "1 hour"
    * "01 hour"
    * "please remind me in 1 hour"
    * "let's meet in one hour"
* -1s (-1 second)
    * "this happend -1 second ago"
    * "-1 second"
    * "one second"
* 1d2m (1 day and 2 minutes)
    * "1 day and 2 minutes"
    * "we have 2 minutes and 1 day left"
    * "please to that in one day and two minutes"
    * "1 day and two minutes"


## Contributing

Everyone wanting to contribute to this library is welcome. Add your ideas and bugs to the issue section. Open a pull request for additional features. 