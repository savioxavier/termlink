<!-- markdownlint-disable MD010 MD033 -->

# termlink

> Clickable links in the terminal for Go

**Termlink is a Go package that allows you to craete fully customizable clickable links in the terminal. It is the Go version of Sindre Sorhus' popular [terminal-link](https://github.com/sindresorhus/terminal-link/) library.**

**It includes multiple features including dynamic links and fully customizable colored links in terminal.**

## 🛠️ Install

Using `go get`:

```go
go get github.com/savioxavier/termlink
```

---

## 🔗 Usage

- Basic regular link:

```go
import (
	"fmt"

	"github.com/savioxavier/termlink"
)

func main() {
	fmt.Println(termlink.Link("Example", "https://example.com"))
}
```

- Customizable colored link:

```go
import (
	"fmt"

	"github.com/savioxavier/termlink"
)

func main() {
	fmt.Println(termlink.ColorLink("Example", "https://example.com", "italic green"))
}
```

- You can also use this package in combination with another popular Go package [fatih/color](https://github.com/fatih/color)

```go
import (
	"fmt"

	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
)

func main() {
	// With fatih/color package
	color.Cyan(termlink.Link("Example link using the colors package!", "https://example.com"))
}
```

---

## 🍵 Examples

More examples can be found in the [`examples/`](examples/) directory.

---

## 🔮 Features

- **`termlink.Link(text, url)`**

  - Creates a regular, clickable link in the terminal
  - For unsupported terminals, the link will be printed in parentheses after the text: `Example Link (https://example.com)`.

- **`termlink.ColorLink(text, url, color)`**

  - Creates a clickable link in the terminal with custom color formatting
  - Examples of color options include:
    - Foreground only: `green`, `red`, `blue`, etc.
    - Background only: `bgGreen`, `bgRed`, `bgBlue`, etc.
    - Foreground and background: `green bgRed`, `bgBlue red`, etc.
    - With formatting: `green bold`, `red bgGreen italic`, `italic blue bgGreen`, etc.

- **`termlink.SupportsHyperlinks()`**:

- Return `true` if the terminal supports hyperlinks, `false` otherwise.

---

## ❤️ Support

You can support further development of this bot by **giving it a 🌟** and help me make even better stuff in the future by **buying me a ☕**

<a href="https://www.buymeacoffee.com/savioxavier">
<img src="https://cdn.buymeacoffee.com/buttons/v2/default-blue.png" height="50px">
</a>

<br>

**Also, if you liked this repo, consider checking out my other projects, that would be real cool!**

---

## 💫 Attributions and special thanks

- [terminal-link](https://github.com/sindresorhus/terminal-link) - Sindre Sorhus' original package for providing inspiration for this package.
- [go-supportscolor](https://github.com/jwalton/go-supportscolor) - A package for detecting terminal color support.
