// Package termlink implements a set of functions to create customizable, clickable hyperlinks in the terminal.
package termlink

import (
	"fmt"
	"os"
	"strings"

	"github.com/jwalton/go-supportscolor"
)

var EnvironmentVariables = []string{
	"DOMTERM",
	"WT_SESSION",
	"KONSOLE_VERSION",
}

type Version struct {
	major int
	minor int
	patch int
}

func parseVersion(version string) Version {
	var major, minor, patch int
	fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	return Version{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func hasEnv(name string) bool {
	_, envExists := os.LookupEnv(name)

	return envExists
}

func checkAllEnvs(vars []string) bool {
	for _, v := range vars {
		if hasEnv(v) {
			return true
		}
	}

	return false
}

func getEnv(name string) string {
	envValue, _ := os.LookupEnv(name)

	return envValue
}

func supportsHyperlinks() bool {
	if hasEnv("VTE_VERSION") {
		// VTE-based terminals above v0.50 (Gnome Terminal, Guake, ROXTerm, etc)
		v := parseVersion(getEnv("VTE_VERSION"))
		return v.major >= 0 && v.minor >= 50 && v.patch > 0
	}

	if hasEnv("TERM_PROGRAM") {
		v := parseVersion(getEnv("TERM_PROGRAM_VERSION"))

		switch term := getEnv("TERM_PROGRAM"); term {
		case "iTerm.app":
			if v.major == 3 {
				return v.minor >= 1
			}
			return v.major > 3
		case "WezTerm":
			return v.major >= 20200620
		case "vscode":
			return v.major > 1 || (v.major == 1 && v.minor >= 72)
		}
	}

	if checkAllEnvs(EnvironmentVariables) {
		return true
	}

	return false
}

var colorsList = map[string]int{
	"reset":     0,
	"bold":      1,
	"dim":       2,
	"italic":    3,
	"underline": 4,
	"blink":     5,
	"black":     30,
	"red":       31,
	"green":     32,
	"yellow":    33,
	"blue":      34,
	"magenta":   35,
	"cyan":      36,
	"white":     37,
	"bgBlack":   40,
	"bgRed":     41,
	"bgGreen":   42,
	"bgYellow":  43,
	"bgBlue":    44,
	"bgMagenta": 45,
	"bgCyan":    46,
	"bgWhite":   47,
}

func isInList(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

var colors []string

func addColor(value string) []string {
	colors = append(colors, fmt.Sprint(colorsList[value]))

	return colors
}

func isValidColor(color string) bool {
	// Create a slice with keys of the colorsList map
	keys := make([]string, len(colorsList))

	i := 0
	for k := range colorsList {
		keys[i] = k
		i++
	}

	// Check if the color is in the keys slice
	return isInList(keys, color)
}

func parseColor(color string) string {
	// If nothing is provided, return empty string
	if color == "" {
		return ""
	}

	for _, c := range strings.Split(color, " ") {
		// If the color doesn't exist, skip and go to the next word
		if !isValidColor(c) {
			continue
		}

		// Add the color, if present in colorsList
		addColor(c)
	}

	return "\u001b[" + strings.Join(colors, ";") + "m"
}

func supportsColor() bool {
	return supportscolor.Stdout().SupportsColor
}

// Function Link creates a clickable link in the terminal's stdout.
//
// The function takes two required parameters: text and url
// and one optional parameter: shouldForce
//
// The text parameter is the text to be displayed.
// The url parameter is the URL to be opened when the link is clicked.
// The shouldForce parameter indicates whether to force the non-hyperlink supported behavior (i.e., text (url))
//
// The function returns the clickable link.
func Link(text string, url string, shouldForce ...bool) string {
	shouldForceDefault := false

	if len(shouldForce) > 0 {
		shouldForceDefault = shouldForce[0]
	}

	if shouldForceDefault {
		return text + " (" + url + ")" + parseColor("reset")
	} else {
		if supportsHyperlinks() {
			return "\x1b]8;;" + url + "\x07" + text + "\x1b]8;;\x07" + parseColor("reset")
		}
		return text + " (" + url + ")" + parseColor("reset")
	}
}

// Function LinkColor creates a colored clickable link in the terminal's stdout.
//
// The function takes three required parameters: text, url and color
// and one optional parameter: shouldForce
//
// The text parameter is the text to be displayed.
// The url parameter is the URL to be opened when the link is clicked.
// The color parameter is the color of the link.
// The shouldForce parameter indicates whether to force the non-hyperlink supported behavior (i.e., text (url))
//
// The function returns the clickable link.
func ColorLink(text string, url string, color string, shouldForce ...bool) string {
	var textColor string

	if supportsColor() {
		textColor = parseColor(color)
	} else {
		textColor = ""
	}

	shouldForceDefault := false

	if len(shouldForce) > 0 {
		shouldForceDefault = shouldForce[0]
	}

	if shouldForceDefault {
		return textColor + text + " (" + url + ")" + parseColor("reset")
	} else {
		if supportsHyperlinks() {
			return "\x1b]8;;" + url + "\x07" + textColor + text + "\x1b]8;;\x07" + parseColor("reset")
		}
		return textColor + text + " (" + url + ")" + parseColor("reset")
	}

}

// Function SupportsHyperlinks returns true if the terminal supports hyperlinks.
//
// The function returns true if the terminal supports hyperlinks, false otherwise.
func SupportsHyperlinks() bool {
	return supportsHyperlinks()
}
