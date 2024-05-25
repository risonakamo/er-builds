package oer_api

import (
	"os"
	"strings"
)

// map of oer lang dicts
type OerLangDictDict map[string]OerLangDict

// leaf fields of oer lang dict
type OerLangDictFields map[string]string

// oerlang dict. infinite nested dictionary
type OerLangDict struct {
	// contains additional levels
	Nested OerLangDictDict

	// contains leaf levels and their string values
	Fields OerLangDictFields
}

// parse langfile string into langfile dict
func parseLangFile(langfile string) OerLangDict {
	var lines []string=strings.Split(langfile,"\n")

	var langdict OerLangDict=newLangDict()

	for i := range lines {
		// splits line into 2 main parts, key and value
		lines[i]=strings.Trim(lines[i],"\r")
		var splitLine []string=strings.Split(lines[i],"┃")

		if len(splitLine)!=2 {
			continue
		}

		var lineKeyText string=splitLine[0]
		var lineValue string=splitLine[1]
		var lineKeys []string=strings.Split(lineKeyText,"/")

		setInLangDict(
			langdict,
			lineKeys,
			lineValue,
		)
	}

	return langdict
}

// set a value in langdict with a string path. MUTATES the dict
func setInLangDict(dict OerLangDict,keys []string,value string) {
	for i := range keys {
		// if we are on the last key, set the value in the dict's Fields dict
		if i==len(keys)-1 {
			dict.Fields[keys[i]]=value
			return

		// otherwise, try to index into the Nested field
		} else {

			// check if it exists first, and create if it doesnt
			_,in:=dict.Nested[keys[i]]

			if !in {
				dict.Nested[keys[i]]=newLangDict()
			}

			dict=dict.Nested[keys[i]]
		}
	}
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

// make new lang dict
func newLangDict() OerLangDict {
	return OerLangDict{
		Nested: make(OerLangDictDict),
		Fields: make(OerLangDictFields),
	}
}