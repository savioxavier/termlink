package termlink_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
	"github.com/stretchr/testify/assert"
)

func testAll(input, expectedHyperlink, expectedNoHyperlink string) func(t *testing.T) {
	return func(t *testing.T) {
		if termlink.SupportsHyperlinks() {
			assert.Equal(t, input, expectedHyperlink)
		} else {
			assert.Equal(t, input, expectedNoHyperlink)
		}
	}
}

func TestBasicLink(t *testing.T) {
	input := termlink.Link("Hello", "https://google.com")
	expectedHyperlink := "\x1b]8;;https://google.com\aHello\x1b]8;;\a\x1b[m"
	expectedNoHyperlink := "Hello (https://google.com)\u001b[m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)
}

func TestColorLink(t *testing.T) {
	input := termlink.ColorLink("Hello", "https://google.com", "red")
	expectedHyperlink := "\x1b]8;;https://google.com\a\x1b[31mHello\x1b]8;;\a\x1b[m"
	expectedNoHyperlink := "\x1b[31mHello (https://google.com)\x1b[m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)
}

func TestColorsPackageLink(t *testing.T) {
	input := color.New(color.FgCyan).SprintFunc()(termlink.Link("Hello", "https://google.com"))
	expectedHyperlink := "\x1b[36m\x1b]8;;https://google.com\aHello\x1b]8;;\a\x1b[m\x1b[0m"
	expectedNoHyperlink := "\x1b[36mHello (https://google.com)\x1b[m\x1b[0m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)
}

func TestBoldLink(t *testing.T) {
	input := termlink.ColorLink("Hello", "https://google.com", "bold")
	expectedHyperlink := "\x1b]8;;https://google.com\a\x1b[1mHello\x1b]8;;\a\x1b[m"
	expectedNoHyperlink := "\x1b[1mHello (https://google.com)\x1b[m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)
}

func TestItalicsLink(t *testing.T) {
	input := termlink.ColorLink("Hello", "https://google.com", "italic")
	expectedHyperlink := "\x1b]8;;https://google.com\a\x1b[3mHello\x1b]8;;\a\x1b[m"
	expectedNoHyperlink := "\x1b[3mHello (https://google.com)\x1b[m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)

}

func TestItalicsBoldLink(t *testing.T) {
	input := termlink.ColorLink("Hello", "https://google.com", "bold italic")
	expectedHyperlink := "\x1b]8;;https://google.com\a\x1b[1;3mHello\x1b]8;;\a\x1b[m"
	expectedNoHyperlink := "\x1b[1;3mHello (https://google.com)\x1b[m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)
}

func TestColorItalicsBoldLink(t *testing.T) {
	input := termlink.ColorLink("Hello", "https://google.com", "red bold italic")
	expectedHyperlink := "\x1b]8;;https://google.com\a\x1b[31;1;3mHello\x1b]8;;\a\x1b[m"
	expectedNoHyperlink := "\x1b[31;1;3mHello (https://google.com)\x1b[m"

	testAll(input, expectedHyperlink, expectedNoHyperlink)(t)
}
