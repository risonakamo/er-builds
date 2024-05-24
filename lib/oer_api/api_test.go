package oer_api

import (
	"fmt"
	"testing"
)

func Test_langApi(t *testing.T) {
	configFile:=readOerConfig("test/oer-config.yml")

	if len(configFile.ApiKey)==0 {
		t.Fatal("failed to read api key")
	}

	fmt.Println("what",configFile)
	getLanguageFile(configFile.ApiKey,ErLang_english)
}