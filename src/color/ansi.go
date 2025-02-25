package color

import (
	"fmt"
	"oh-my-posh/regex"
	"oh-my-posh/shell"
	"strings"
)

const (
	AnsiRegex = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

	OSC99 string = "osc99"
	OSC7  string = "osc7"
)

type Ansi struct {
	title                 string
	shell                 string
	linechange            string
	left                  string
	right                 string
	creset                string
	clearBelow            string
	clearLine             string
	saveCursorPosition    string
	restoreCursorPosition string
	colorSingle           string
	colorFull             string
	colorTransparent      string
	escapeLeft            string
	escapeRight           string
	hyperlink             string
	hyperlinkRegex        string
	osc99                 string
	osc7                  string
	bold                  string
	italic                string
	underline             string
	overline              string
	strikethrough         string
	blink                 string
	reverse               string
	dimmed                string
	format                string
}

func (a *Ansi) Init(shellName string) {
	a.shell = shellName
	switch shellName {
	case shell.ZSH:
		a.format = "%%{%s%%}"
		a.linechange = "%%{\x1b[%d%s%%}"
		a.right = "%%{\x1b[%dC%%}"
		a.left = "%%{\x1b[%dD%%}"
		a.creset = "%{\x1b[0m%}"
		a.clearBelow = "%{\x1b[0J%}"
		a.clearLine = "%{\x1b[K%}"
		a.saveCursorPosition = "%{\x1b7%}"
		a.restoreCursorPosition = "%{\x1b8%}"
		a.title = "%%{\x1b]0;%s\007%%}"
		a.colorSingle = "%%{\x1b[%sm%%}%s%%{\x1b[0m%%}"
		a.colorFull = "%%{\x1b[%sm\x1b[%sm%%}%s%%{\x1b[0m%%}"
		a.colorTransparent = "%%{\x1b[%s;49m\x1b[7m%%}%s%%{\x1b[0m%%}"
		a.escapeLeft = "%{"
		a.escapeRight = "%}"
		a.hyperlink = "%%{\x1b]8;;%s\x1b\\%%}%s%%{\x1b]8;;\x1b\\%%}"
		a.hyperlinkRegex = `(?P<STR>%{\x1b]8;;(.+)\x1b\\%}(?P<TEXT>.+)%{\x1b]8;;\x1b\\%})`
		a.osc99 = "%%{\x1b]9;9;\"%s\"\x1b\\%%}"
		a.osc7 = "%%{\x1b]7;file:\"//%s/%s\"\x1b\\%%}"
		a.bold = "%%{\x1b[1m%%}%s%%{\x1b[22m%%}"
		a.italic = "%%{\x1b[3m%%}%s%%{\x1b[23m%%}"
		a.underline = "%%{\x1b[4m%%}%s%%{\x1b[24m%%}"
		a.overline = "%%{\x1b[53m%%}%s%%{\x1b[55m%%}"
		a.blink = "%%{\x1b[5m%%}%s%%{\x1b[25m%%}"
		a.reverse = "%%{\x1b[7m%%}%s%%{\x1b[27m%%}"
		a.dimmed = "%%{\x1b[2m%%}%s%%{\x1b[22m%%}"
		a.strikethrough = "%%{\x1b[9m%%}%s%%{\x1b[29m%%}"
	case shell.BASH:
		a.format = "\001%s\002"
		a.linechange = "\001\x1b[%d%s\002"
		a.right = "\001\x1b[%dC\002"
		a.left = "\001\x1b[%dD\002"
		a.creset = "\001\x1b[0m\002"
		a.clearBelow = "\001\x1b[0J\002"
		a.clearLine = "\001\x1b[K\002"
		a.saveCursorPosition = "\001\x1b7\002"
		a.restoreCursorPosition = "\001\x1b8\002"
		a.title = "\001\x1b]0;%s\007\002"
		a.colorSingle = "\001\x1b[%sm\002%s\001\x1b[0m\002"
		a.colorFull = "\001\x1b[%sm\x1b[%sm\002%s\001\x1b[0m\002"
		a.colorTransparent = "\001\x1b[%s;49m\x1b[7m\002%s\001\x1b[0m\002"
		a.escapeLeft = "\001"
		a.escapeRight = "\002"
		a.hyperlink = "\001\x1b]8;;%s\x1b\\\\\002%s\001\x1b]8;;\x1b\\\\\002"
		a.hyperlinkRegex = `(?P<STR>\001\x1b\]8;;(.+)\x1b\\\\\002(?P<TEXT>.+)\001\x1b\]8;;\x1b\\\\\002)`
		a.osc99 = "\001\x1b]9;9;\"%s\"\x1b\\\\\002"
		a.osc7 = "\001\x1b]7;\"file://%s/%s\"\x1b\\\\\002"
		a.bold = "\001\x1b[1m\002%s\001\x1b[22m\002"
		a.italic = "\001\x1b[3m\002%s\001\x1b[23m\002"
		a.underline = "\001\x1b[4m\002%s\001\x1b[24m\002"
		a.overline = "\001\x1b[53m\002%s\001\x1b[55m\002"
		a.blink = "\001\x1b[5m\002%s\001\x1b[25m\002"
		a.reverse = "\001\x1b[7m\002%s\001\x1b[27m\002"
		a.dimmed = "\001\x1b[2m\002%s\001\x1b[22m\002"
		a.strikethrough = "\001\x1b[9m\002%s\001\x1b[29m\002"
	default:
		a.format = "%s"
		a.linechange = "\x1b[%d%s"
		a.right = "\x1b[%dC"
		a.left = "\x1b[%dD"
		a.creset = "\x1b[0m"
		a.clearBelow = "\x1b[0J"
		a.clearLine = "\x1b[K"
		a.saveCursorPosition = "\x1b7"
		a.restoreCursorPosition = "\x1b8"
		a.title = "\x1b]0;%s\007"
		a.colorSingle = "\x1b[%sm%s\x1b[0m"
		a.colorFull = "\x1b[%sm\x1b[%sm%s\x1b[0m"
		a.colorTransparent = "\x1b[%s;49m\x1b[7m%s\x1b[0m"
		a.escapeLeft = ""
		a.escapeRight = ""
		a.hyperlink = "\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\"
		a.hyperlinkRegex = "(?P<STR>\x1b]8;;(.+)\x1b\\\\\\\\?(?P<TEXT>.+)\x1b]8;;\x1b\\\\)"
		a.osc99 = "\x1b]9;9;\"%s\"\x1b\\"
		a.osc7 = "\x1b]7;\"file://%s/%s\"\x1b\\"
		a.bold = "\x1b[1m%s\x1b[22m"
		a.italic = "\x1b[3m%s\x1b[23m"
		a.underline = "\x1b[4m%s\x1b[24m"
		a.overline = "\x1b[53m%s\x1b[55m"
		a.blink = "\x1b[5m%s\x1b[25m"
		a.reverse = "\x1b[7m%s\x1b[27m"
		a.dimmed = "\x1b[2m%s\x1b[22m"
		a.strikethrough = "\x1b[9m%s\x1b[29m"
	}
}

func (a *Ansi) InitPlain() {
	a.Init(shell.PLAIN)
}

func (a *Ansi) GenerateHyperlink(text string) string {
	// hyperlink matching
	results := regex.FindNamedRegexMatch("(?P<ALL>(?:\\[(?P<TEXT>.+)\\])(?:\\((?P<URL>.*)\\)))", text)
	if len(results) != 3 {
		return text
	}
	linkText := a.escapeLinkTextForFishShell(results["TEXT"])
	// build hyperlink ansi
	hyperlink := fmt.Sprintf(a.hyperlink, results["URL"], linkText)
	// replace original text by the new onex
	return strings.Replace(text, results["ALL"], hyperlink, 1)
}

func (a *Ansi) escapeLinkTextForFishShell(text string) string {
	if a.shell != shell.FISH {
		return text
	}
	escapeChars := map[string]string{
		`c`: `\c`,
		`a`: `\a`,
		`b`: `\b`,
		`e`: `\e`,
		`f`: `\f`,
		`n`: `\n`,
		`r`: `\r`,
		`t`: `\t`,
		`v`: `\v`,
		`$`: `\$`,
		`*`: `\*`,
		`?`: `\?`,
		`~`: `\~`,
		`%`: `\%`,
		`#`: `\#`,
		`(`: `\(`,
		`)`: `\)`,
		`{`: `\{`,
		`}`: `\}`,
		`[`: `\[`,
		`]`: `\]`,
		`<`: `\<`,
		`>`: `\>`,
		`^`: `\^`,
		`&`: `\&`,
		`;`: `\;`,
		`"`: `\"`,
		`'`: `\'`,
		`x`: `\x`,
		`X`: `\X`,
		`0`: `\0`,
		`u`: `\u`,
		`U`: `\U`,
	}
	if val, ok := escapeChars[text[0:1]]; ok {
		return val + text[1:]
	}
	return text
}

func (a *Ansi) formatText(text string) string {
	replaceFormats := func(results []map[string]string) {
		for _, result := range results {
			var formatted string
			switch result["format"] {
			case "b":
				formatted = fmt.Sprintf(a.bold, result["text"])
			case "u":
				formatted = fmt.Sprintf(a.underline, result["text"])
			case "o":
				formatted = fmt.Sprintf(a.overline, result["text"])
			case "i":
				formatted = fmt.Sprintf(a.italic, result["text"])
			case "s":
				formatted = fmt.Sprintf(a.strikethrough, result["text"])
			case "d":
				formatted = fmt.Sprintf(a.dimmed, result["text"])
			case "f":
				formatted = fmt.Sprintf(a.blink, result["text"])
			case "r":
				formatted = fmt.Sprintf(a.reverse, result["text"])
			}
			text = strings.Replace(text, result["context"], formatted, 1)
		}
	}
	rgx := "(?P<context><(?P<format>[buisrdfo])>(?P<text>[^<]+)</[buisrdfo]>)"
	for results := regex.FindAllNamedRegexMatch(rgx, text); len(results) != 0; results = regex.FindAllNamedRegexMatch(rgx, text) {
		replaceFormats(results)
	}
	return text
}

func (a *Ansi) CarriageForward() string {
	return fmt.Sprintf(a.right, 1000)
}

func (a *Ansi) GetCursorForRightWrite(length, offset int) string {
	strippedLen := length + (-offset)
	return fmt.Sprintf(a.left, strippedLen)
}

func (a *Ansi) ChangeLine(numberOfLines int) string {
	position := "B"
	if numberOfLines < 0 {
		position = "F"
		numberOfLines = -numberOfLines
	}
	return fmt.Sprintf(a.linechange, numberOfLines, position)
}

func (a *Ansi) ConsolePwd(pwdType, hostName, pwd string) string {
	if strings.HasSuffix(pwd, ":") {
		pwd += "\\"
	}
	switch pwdType {
	case OSC7:
		return fmt.Sprintf(a.osc7, hostName, pwd)
	case OSC99:
		fallthrough
	default:
		return fmt.Sprintf(a.osc99, pwd)
	}
}

func (a *Ansi) ClearAfter() string {
	return a.clearLine + a.clearBelow
}

func (a *Ansi) Title(title string) string {
	return fmt.Sprintf(a.title, title)
}

func (a *Ansi) ColorReset() string {
	return a.creset
}

func (a *Ansi) FormatText(text string) string {
	return fmt.Sprintf(a.format, text)
}

func (a *Ansi) SaveCursorPosition() string {
	return a.saveCursorPosition
}

func (a *Ansi) RestoreCursorPosition() string {
	return a.restoreCursorPosition
}
