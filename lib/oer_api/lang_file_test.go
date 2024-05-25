package oer_api

import "testing"

func Test_parseLangFile(t *testing.T) {
	langfile:=readLangFile("test/saved-langfile.txt")

	parseLangFile(langfile)
}