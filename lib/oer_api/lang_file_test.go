package oer_api

import (
	"testing"

	"github.com/kr/pretty"
)

func Test_parseLangFile(t *testing.T) {
	langfile:=readLangFile("test/saved-langfile.txt")

	langdict:=parseLangFile(langfile)

	// pretty.Print(langdict)
	pretty.Print(langdict.Nested["Item"].Nested["Name"].Fields)
}