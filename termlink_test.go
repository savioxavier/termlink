package termlink_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
	"github.com/stretchr/testify/assert"
)

func TestBasicLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := termlink.Link("Hello", "https://google.com")
		expected := "\x1b]8;;https://google.com\aHello\x1b]8;;\a\x1b[m"
		assert.Equal(t, input, expected)
	} else {
		input := termlink.Link("Hello", "https://google.com")
		expected := "Hello (\u200Bhttps://google.com)\u001b[m"
		assert.Equal(t, input, expected)
	}
}

func TestColorLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := termlink.ColorLink("Hello", "https://google.com", "red")
		expected := "\x1b]8;;https://google.com\a\x1b[31mHello\x1b]8;;\a\x1b[m"
		assert.Equal(t, input, expected)
	} else {
		input := termlink.ColorLink("Hello", "https://google.com", "red")
		expected := "\x1b[31mHello (\u200bhttps://google.com)\x1b[m"
		assert.Equal(t, input, expected)
	}
}

func TestColorsPackageLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := color.New(color.FgCyan).SprintFunc()(termlink.Link("Hello", "https://google.com"))
		expected := "\x1b[36m\x1b]8;;https://google.com\aHello\x1b]8;;\a\x1b[m\x1b[0m"
		assert.Equal(t, input, expected)
	} else {
		input := color.New(color.FgCyan).SprintFunc()(termlink.Link("Hello", "https://google.com"))
		expected := "\x1b[36mHello (\u200bhttps://google.com)\x1b[m\x1b[0m"
		assert.Equal(t, input, expected)
	}
}

func TestBoldLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := termlink.ColorLink("Hello", "https://google.com", "bold")
		expected := "\x1b]8;;https://google.com\a\x1b[1mHello\x1b]8;;\a\x1b[m"
		assert.Equal(t, input, expected)
	} else {
		input := termlink.ColorLink("Hello", "https://google.com", "bold")
		expected := "\x1b[1mHello (\u200bhttps://google.com)\x1b[m"
		assert.Equal(t, input, expected)
	}
}

func TestItalicsLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := termlink.ColorLink("Hello", "https://google.com", "italic")
		expected := "\x1b]8;;https://google.com\a\x1b[3mHello\x1b]8;;\a\x1b[m"
		assert.Equal(t, input, expected)
	} else {
		input := termlink.ColorLink("Hello", "https://google.com", "italic")
		expected := "\x1b[3mHello (\u200bhttps://google.com)\x1b[m"
		assert.Equal(t, input, expected)
	}
}

func TestItalicsBoldLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := termlink.ColorLink("Hello", "https://google.com", "bold italic")
		expected := "\x1b]8;;https://google.com\a\x1b[1;3mHello\x1b]8;;\a\x1b[m"
		assert.Equal(t, input, expected)
	} else {
		input := termlink.ColorLink("Hello", "https://google.com", "bold italic")
		expected := "\x1b[1;3mHello (\u200bhttps://google.com)\x1b[m"
		assert.Equal(t, input, expected)
	}
}

func TestColorItalicsBoldLink(t *testing.T) {
	if termlink.SupportsHyperlinks() {
		input := termlink.ColorLink("Hello", "https://google.com", "red bold italic")
		expected := "\x1b]8;;https://google.com\a\x1b[31;1;3mHello\x1b]8;;\a\x1b[m"
		assert.Equal(t, input, expected)
	} else {
		input := termlink.ColorLink("Hello", "https://google.com", "red bold italic")
		expected := "\x1b[31;1;3mHello (\u200bhttps://google.com)\x1b[m"
		assert.Equal(t, input, expected)
	}
}
