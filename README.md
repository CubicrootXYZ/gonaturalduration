# go-natural-duration

[![GitHub license](https://img.shields.io/github/license/CubicrootXYZ/gonaturalduration)](https://github.com/CubicrootXYZ/gonaturalduration/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/CubicrootXYZ/gonaturalduration)](https://github.com/CubicrootXYZ/gonaturalduration/issues)
[![Actions Status](https://github.com/CubicrootXYZ/gonaturalduration/workflows/Main/badge.svg?branch=main)](https://github.com/CubicrootXYZ/gonaturalduration/workflows/actions)

Library for parsing time durations in natural language.

## Features

* Supports different kinds of durations:
    * days
    * hours
    * minutes
    * seconds
* Support for english language
* Parses from any string, even with context around the duration

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
	duration := gonaturalduration.Parse("in 1 hour")
	log.Print(fmt.Sprintf("%v", duration))
}
```

## Parser examples

* 1d (1 day)
    * "1 day"
    * "there once was an attempt to generate a duration of 1 day"
    * "this is 01 day"
* 1h (1 hour)
    * "1 hour"
    * "01 hour"
    * "please remind me in 1 hour"
* -1s (-1 second)
    * "this happend -1 second ago"
    * "-1 second"
* 1d2m (1 day and 2 minutes)
    * "1 day and 2 minutes"
    * "we have 2 minutes and 1 day left"
