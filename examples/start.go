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
Here's an example to get you started:

	Hello I'm %s.
	I'm %d years old.
	Here's Twitter: %s.
	And here's Twitter's Twitter: %s.
		`,
		name,
		age,
		termlink.Link("Twitter", "https://twitter.com"),
		termlink.ColorLink("@twitter", "https://twitter.com/twitter", "italic blue"),
	)
}

func printHelpLinks() {
	fmt.Printf(
		`
Also, here are some other links to help you out:

	🐹 %s
	💻 %s

	🔗 %s
	💖 %s
		`,
		termlink.Link("Official Go Website", "https://go.dev"),
		termlink.ColorLink("A Tour of Go", "https://go.dev/tour", "cyan"),
		termlink.ColorLink("termlink's source code", "https://github.com/savioxavier/termlink", "italic magenta"),
		termlink.ColorLink("termlink's donate page", "https://www.buymeacoffee.com/savioxavier", "italic red"),
	)
}

func printGoodbye() {
	fmt.Println()
	fmt.Println("👋 Goodbye!")
}

func main() {
	printSupportsHyperlinks()
	printExample()
	printHelpLinks()
	printGoodbye()
}
