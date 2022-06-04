package termlink_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
	"github.com/stretchr/testify/assert"
)

func TestBasicLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		assert.Equal(t, termlink.Link("Hello", "https://google.com"), "\x1b]8;;https://google.com\aHello\x1b]8;;\a\x1b[m")
	} else {
		assert.Equal(t, termlink.Link("Hello", "https://google.com"), "Hello (\u200Bhttps://google.com)\u001b[m")
	}
}

func TestColorLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "red"),
			"\x1b]8;;https://google.com\a\x1b[31mHello\x1b]8;;\a\x1b[m")
	} else {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "red"),
			"\x1b[31mHello (\u200bhttps://google.com)\x1b[m")

	}
}

func TestColorsPackageLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		// Check if color.Cyan(termlink.Link("Hello", "https://google.com")) will print the output to the terminal
		assert.Equal(t, color.New(color.FgCyan).SprintFunc()(termlink.Link("Hello", "https://google.com")), "\x1b[36m\x1b]8;;https://google.com\aHello\x1b]8;;\a\x1b[m\x1b[0m")
	} else {
		assert.Equal(t, color.New(color.FgCyan).SprintFunc()(termlink.Link("Hello", "https://google.com")), "\x1b[36mHello (\u200bhttps://google.com)\x1b[m\x1b[0m")
	}
}

func TestBoldLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "bold"),
			"\x1b]8;;https://google.com\a\x1b[1mHello\x1b]8;;\a\x1b[m")
	} else {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "bold"), "\x1b[1mHello (\u200bhttps://google.com)\x1b[m")
	}
}

func TestItalicsLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "italic"),
			"\x1b]8;;https://google.com\a\x1b[3mHello\x1b]8;;\a\x1b[m")
	} else {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "italic"), "\x1b[3mHello (\u200bhttps://google.com)\x1b[m")
	}
}

func TestItalicsBoldLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "bold italic"),
			"\x1b]8;;https://google.com\a\x1b[1;3mHello\x1b]8;;\a\x1b[m")
	} else {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "bold italic"), "\x1b[1;3mHello (\u200bhttps://google.com)\x1b[m")
	}
}

func TestColorItalicsBoldLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "red bold italic"),
			"\x1b]8;;https://google.com\a\x1b[31;1;3mHello\x1b]8;;\a\x1b[m")
	} else {
		assert.Equal(t, termlink.ColorLink("Hello", "https://google.com", "red bold italic"), "\x1b[31;1;3mHello (\u200bhttps://google.com)\x1b[m")
	}
}
