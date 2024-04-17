package view

import (
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.AmericanEnglish)
var caserM = sync.Mutex{}

func prittyName(s string) string {
	caserM.Lock()
	r := caser.String(strings.ReplaceAll(s, "_", " "))
	caserM.Unlock()
	return r
}
