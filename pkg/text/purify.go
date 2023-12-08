package text

import (
	"regexp"
	"strings"
)

var whiteSpaceRegex = regexp.MustCompile(`\s+`)

func Purify(text string) string {
	nbsps := []string{"\u00a0", "\u2007", "\u202f", "\u2060"}

	for _, nbsp := range nbsps {
		text = strings.ReplaceAll(text, nbsp, " ")
	}

	text = strings.TrimSpace(text)
	text = whiteSpaceRegex.ReplaceAllString(text, " ")

	return text
}
