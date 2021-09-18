package gohtml

import (
	"bytes"
	"strconv"
	"strings"
)

// Format parses the input HTML string, formats it and returns the result.
func Format(s string, preserveNewLines bool) string {
	return parse(strings.NewReader(s), preserveNewLines).html()
}

// FormatBytes parses input HTML as bytes, formats it and returns the result.
func FormatBytes(b []byte, preserveNewLines bool) []byte {
	return parse(bytes.NewReader(b), preserveNewLines).bytes()
}

// Format parses the input HTML string, formats it and returns the result with line no.
func FormatWithLineNo(s string, preserveEmptyLines bool) string {
	return AddLineNo(Format(s, preserveEmptyLines))
}

func AddLineNo(s string) string {
	lines := strings.Split(s, "\n")
	maxLineNoStrLen := len(strconv.Itoa(len(lines)))
	bf := &bytes.Buffer{}
	for i, line := range lines {
		lineNoStr := strconv.Itoa(i + 1)
		if i > 0 {
			bf.WriteString("\n")
		}
		bf.WriteString(strings.Repeat(" ", maxLineNoStrLen-len(lineNoStr)) + lineNoStr + "  " + line)
	}
	return bf.String()

}
