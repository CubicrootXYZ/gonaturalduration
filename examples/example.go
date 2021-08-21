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
