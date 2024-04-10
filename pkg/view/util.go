package view

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.AmericanEnglish)

func prittyName(s string) string {
	return caser.String(strings.ReplaceAll(s, "_", " "))
}
