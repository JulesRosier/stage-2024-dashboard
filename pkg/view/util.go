package view

import (
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.AmericanEnglish)
var caserM = sync.Mutex{}

func prettyName(s string) string {
	caserM.Lock()
	r := caser.String(strings.ReplaceAll(s, "_", " "))
	caserM.Unlock()
	return r
}

func formatIndexName(n string) string {
	n, _ = strings.CutPrefix(n, "index_")
	r := caser.String(strings.ReplaceAll(n, "_", " "))
	return r
}

func shortenedName(in string) string {
	if len(in) == 0 {
		return ""
	}
	words := strings.Fields(in)
	var firstLetters strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			firstLetters.WriteByte(word[0]) // Append the first letter of the word
		}
	}
	return firstLetters.String()
}
