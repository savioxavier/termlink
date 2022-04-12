package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
)

func main() {
	name := "John"
	age := 21

	// With regular termlink package
	fmt.Printf(
		`
		Hello I'm %s.
		I'm %d years old.
		Here's Twitter: %s.
		And here's Twitter's Twitter: %s.
		Anyways, see ya!
		`,
		name,
		age,
		termlink.Link("Twitter", "https://twitter.com"),
		termlink.ColorLink("@twitter", "https://twitter.com/twitter", "italic green"),
	)

	// With fatih/color package
	color.Cyan(termlink.Link("Example link using the colors package!", "https://example.com"))
}
