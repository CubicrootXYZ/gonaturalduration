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
