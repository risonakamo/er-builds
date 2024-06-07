// funcs handling lang file data struct format

package oer_api

import (
	go_utils "er-builds/lib/utils"
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

// oerlang dict. infinite nested dictionary. top level of oer lang dict
// see newLangDict() to construct
type OerLangDict struct {
	// contains additional levels
	Nested OerLangDictDict

	// contains leaf levels and their string values
	Fields OerLangDictFields
}

// map of oer lang dicts. NOT the top level
type OerLangDictDict map[string]OerLangDict

// leaf fields of oer lang dict
type OerLangDictFields map[string]string

// read lang file into dict. if file not found, does not crash, returns empty dict
func ReadLangFileToDict(filename string) OerLangDict {
	var langfileText string
	var e error
	langfileText,e=readLangFile(filename)

	if e!=nil {
		// if lang file not found, return empty
		if errors.Is(e,os.ErrNotExist) {
			log.Warn().Msg("lang file did not exist")
			return newLangDict()
		}

		panic(e)
	}

	return parseLangFile(langfileText)
}

// write langfile string to txt file
func WriteLangFile(filename string,langfile string) {
	var e error=go_utils.WriteStringToCompressedFile(filename,langfile)

	if e!=nil {
		panic(e)
	}
}

// get langfile string from file
func readLangFile(filename string) (string,error) {
	var result string
	var e error
	result,e=go_utils.ReadCompressedStringFile(filename)

	return result,e
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

// make new lang dict
func newLangDict() OerLangDict {
	return OerLangDict{
		Nested: make(OerLangDictDict),
		Fields: make(OerLangDictFields),
	}
}