package main

import (
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	argsWithoutProg := os.Args[1:]
	text := strings.Join(argsWithoutProg, " ")
	myFigure := figure.NewFigure(text, "starwars", true)
	myFigure.Print()
}
