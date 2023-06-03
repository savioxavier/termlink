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

var ValueSpecificEnvironmentVariables = map[string][]string{
	"TERM_PROGRAM": []string{
		"iTerm.app",
		"terminology",
		"WezTerm",
		"Hyper",
	},
	"TERM": []string{
		"xterm-kitty",
	},
}

func parseVersion(version string) (int, int, int) {
	var major, minor, patch int
	fmt.Sscanf(version, "%d.%d.%d", &major, &minor, &patch)
	return major, minor, patch
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

func matchesEnv(name string, subNames []string) bool {
	if hasEnv(name) {
		for _, subName := range subNames {
			if getEnv(name) == subName {
				return true
			}
		}
	}
	return false
}

func matchAllEnvs(envList map[string][]string) bool {
	for key, value := range envList {
		if matchesEnv(key, value) {
			return true
		}
	}
	return false
}

func supportsHyperlinks() bool {
	if hasEnv("VTE_VERSION") {
		// VTE-based terminals above v0.50 (Gnome Terminal, Guake, ROXTerm, etc)
		major, minor, patch := parseVersion(getEnv("VTE_VERSION"))
		// 0.50.0 was supposed to support hyperlinks, but throws a segfault
		if major >= 0 && minor >= 50 && patch > 0 {
			return true
		}
	}

	if matchesEnv("TERM_PROGRAM", []string{"vscode"}) {
		major, minor, _ := parseVersion(getEnv("TERM_PROGRAM_VERSION"))
		if major > 1 || (major == 1 && minor >= 72) {
			return true
		}
	}

	if checkAllEnvs(EnvironmentVariables) || matchAllEnvs(ValueSpecificEnvironmentVariables) {
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
