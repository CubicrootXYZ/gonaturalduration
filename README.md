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

## Examples

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
