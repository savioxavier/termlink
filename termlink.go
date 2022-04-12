package termlink

import (
	"fmt"
	"os"
	"strings"

	"github.com/jwalton/go-supportscolor"
)

func parseVersion(version string) (int, int, int) {
	var major, minor, patch int
	fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	return major, minor, patch
}

func supportsHyperlinks() bool {
	if os.Getenv("FORCE_HYPERLINK") != "" {
		return true
	}

	if os.Getenv("DOMTERM") != "" {
		// DomTerm
		return true
	}

	if os.Getenv("VTE_VERSION") != "" {
		// VTE-based terminals above v0.50 (Gnome Terminal, Guake, ROXTerm, etc)
		major, minor, patch := parseVersion(os.Getenv("VTE_VERSION"))
		if major >= 5000 && minor >= 50 && patch >= 50 {
			return true
		}
	}

	if os.Getenv("TERM_PROGRAM") != "" {
		if os.Getenv("TERM_PROGRAM") == "Hyper" ||
			os.Getenv("TERM_PROGRAM") == "iTerm.app" ||
			os.Getenv("TERM_PROGRAM") == "terminology" ||
			os.Getenv("TERM_PROGRAM") == "WezTerm" {
			return true
		}
	}

	if os.Getenv("TERM") != "" {
		// Kitty
		if os.Getenv("TERM") == "xterm-kitty" {
			return true
		}
	}

	// Windows Terminal and Konsole
	if os.Getenv("WT_SESSION") != "" || os.Getenv("KONSOLE_VERSION") != "" {
		return true
	}

	return false
}

var colorsList = map[string]string{
	"black":     "30",
	"red":       "31",
	"green":     "32",
	"yellow":    "33",
	"blue":      "34",
	"magenta":   "35",
	"cyan":      "36",
	"white":     "37",
	"bold":      "1",
	"italic":    "3",
	"bgBlack":   "40",
	"bgRed":     "41",
	"bgGreen":   "42",
	"bgYellow":  "43",
	"bgBlue":    "44",
	"bgMagenta": "45",
	"bgCyan":    "46",
	"bgWhite":   "47",
}

func isInList(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func parseColor(color string) string {
	acceptedForegroundColors := []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}

	acceptedBackgroundColors := []string{"bgBlack", "bgRed", "bgGreen", "bgYellow", "bgBlue", "bgMagenta", "bgCyan", "bgWhite"}

	if color == "" {
		return ""
	}

	var colors []string
	for _, c := range strings.Split(color, " ") {
		if c == "" {
			continue
		}
		if c == "bold" {
			colors = append(colors, colorsList["bold"])
		} else if c == "italic" {
			colors = append(colors, colorsList["italic"])
		} else if c == "underline" {
			colors = append(colors, colorsList["underline"])
		} else if c == "blink" {
			colors = append(colors, colorsList["blink"])
		} else if c == "reverse" {
			colors = append(colors, colorsList["reverse"])
		} else if c == "hidden" {
			colors = append(colors, colorsList["hidden"])
		} else if c == "strike" {
			colors = append(colors, colorsList["strike"])
		} else if isInList(acceptedForegroundColors, c) {
			colors = append(colors, colorsList[c])
		} else if isInList(acceptedBackgroundColors, c) {
			colors = append(colors, colorsList[c])
		} else if c == "reset" {
			colors = append(colors, colorsList["reset"])
		} else {
			return ""
		}
	}

	if len(colors) == 0 {
		return ""
	}

	return "\u001b[" + strings.Join(colors, ";") + "m"
}

func supportsColor() bool {
	return supportscolor.Stdout().SupportsColor
}

/**
 * @info Create a clickable link in the terminal's stdout.
 * @param {string} [text] The text to be displayed.
 * @param {string} [url] The URL to be opened when the link is clicked.
 * @returns {string} The clickable link.
 */
func Link(text string, url string) string {
	if supportsHyperlinks() {
		return "\x1b]8;;" + url + "\x07" + text + "\x1b]8;;\x07" + parseColor("reset")
	} else {
		return text + " (\u200B" + url + ")" + parseColor("reset")
	}
}

/**
 * @info Create a clickable link in the terminal's stdout.
 * @param {string} [text] The text to be displayed.
 * @param {string} [url] The URL to be opened when the link is clicked.
 * @param {string} [color] The color of the link.
 * @returns {string} The clickable link.
 */
func ColorLink(text string, url string, color string) string {
	var textColor string

	if supportsColor() {
		textColor = parseColor(color)
	} else {
		textColor = ""
	}
	if supportsHyperlinks() {
		return "\x1b]8;;" + url + "\x07" + textColor + text + "\x1b]8;;\x07" + parseColor("reset")
	} else {
		return textColor + text + " (\u200B" + url + ")" + parseColor("reset")
	}
}

// export supportsHyperlinks to be used in testing
func SupportsHyperlinks() bool {
	return supportsHyperlinks()
}
