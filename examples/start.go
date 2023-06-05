package main

import (
	"fmt"

	"github.com/savioxavier/termlink"
)

func printSupportsHyperlinks() {
	fmt.Printf("Does this terminal support hyperlinks: %t\n", termlink.SupportsHyperlinks())
}

func printExample() {
	name := "John"
	age := 21

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
		termlink.ColorLink("@twitter", "https://twitter.com/twitter", "italic cyan"),
	)
}

func main() {
	printSupportsHyperlinks()
	printExample()
}
