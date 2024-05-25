package oer_api

import (
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp/v3"
)

// oerlang dict. infinite nested dictionary
type OerLangDict struct {
	// contains additional levels
	Nested map[string]OerLangDict

	// contains leaf levels and their string values
	Fields map[string]string
}

// parse langfile string into langfile dict
func parseLangFile(langfile string) OerLangDict {
	var lines []string=strings.Split(langfile,"\n")

	for i := range lines {
		// splits line into 2 main parts, key and value
		var splitLine []string=strings.Split(lines[i],"┃")

		if len(splitLine)!=2 {
			continue
		}

		var lineKeyText string=splitLine[0]
		var lineValue string=splitLine[1]
		var lineKeys []string=strings.Split(lineKeyText,"/")

		pp.Print(lineKeys)
		fmt.Println(lineValue)
	}

	return OerLangDict{}
}

// write langfile string to txt file
func writeLangFile(filename string,langfile string) {
	var wfile *os.File
	var e error
	wfile,e=os.Create(filename)

	if e!=nil {
		panic(e)
	}

	defer wfile.Close()

	wfile.Write([]byte(langfile))
}

// get langfile string from file
func readLangFile(filename string) string {
	var data []byte
	var e error
	data,e=os.ReadFile(filename)

	if e!=nil {
		panic(e)
	}

	return string(data)
}