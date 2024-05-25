package oer_api

import (
	"testing"
)

func Test_langApi(t *testing.T) {
	configFile:=readOerConfig("test/oer-config.yml")

	if len(configFile.ApiKey)==0 {
		t.Fatal("failed to read api key")
	}

	var langfile string=getLanguageFile(configFile.ApiKey,ErLang_english)
	// fmt.Println(langfile)

	writeLangFile("test/saved-langfile.txt",langfile)
}